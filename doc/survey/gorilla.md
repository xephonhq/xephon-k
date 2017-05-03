# Gorilla (Beringei)

- https://github.com/facebookincubator/beringei

## TODO

- [ ] does gorilla use milliseconds?
- [ ] `PutDataRequest` in seems low efficient, why use list on `DataPoint` instead of `TimeValuePair`, though with zlib, the impact of duplication of key could be reduced a lot
- [ ] they seems to be using general purpose compression algorithm like ZLIB in both thrift and on disk structure, how much space does this save?
- [ ] are on disk blocks interleaved across series

## Meta

- Time interval: varies
- Time precision: Unix timestamp in milliseconds? (64bit)
  - 1491502273                    (second)
  - 1491502273000                 (millisecond)
  - 1491502273000000              (microsecond)
  - 1491502273000000000           (nanosecond)
  - 2147483647                    (max of int32 is not enough for millisecond, unit32 won't help)
  - 9223372036854775807           
- Value precision: 64 bit float
- Store Timestamp: Yes
- Time and value are compressed separately
  - [x] But they are stored in same bit stream? Yes, with variable length

## Architecture

### Entities

[beringei_data.thrift](https://github.com/facebookincubator/beringei/blob/master/beringei/if/beringei_data.thrift)

````
struct TimeValuePair {
  1: i64 unixTime,
  2: double value,
}

struct DataPoint {
  1: Key key,
  2: TimeValuePair value,
  3: i32 categoryId,
}

struct PutDataRequest {
  1: list<DataPoint> data,
}

struct PutDataResult {
  // return not owned data points
  1: list<DataPoint> data,
}
````

````
// getData structs
enum Compression {
  NONE,
  ZLIB,
}

// DO NOT interact with this struct directly. Feed it into TimeSeries.h.
struct TimeSeriesBlock {
  1: Compression compression,
  2: i32 count,
  3: binary data,
}
````

### Compression

- The paper has error https://github.com/facebookincubator/beringei/issues/17
- The compression in memory, disk and log are similar, time and value are stored in **same** bit stream
- [TimeSeriesStream.cpp](https://github.com/facebookincubator/beringei/blob/master/beringei/lib/TimeSeriesStream.cpp)
- [appendTimestamp](https://github.com/facebookincubator/beringei/blob/master/beringei/lib/TimeSeriesStream.cpp#L100)

````cpp
folly::fbstring data_;

bool TimeSeriesStream::appendTimestamp(
    int64_t timestamp,
    int64_t minTimestampDelta) {...}
void TimeSeriesStream::appendValue(double value) {...}

BitUtil::addValueToBitString(0, 1, data_, numBits_);
````

### In Memory

- [BucketMap.h](https://github.com/facebookincubator/beringei/blob/master/beringei/lib/BucketMap.h)
- https://github.com/facebookincubator/beringei/blob/master/beringei/lib/BucketMap.h#L89
  - [ ] TODO: why put is virtual

````cpp
virtual std::pair<int, int> put(
    const std::string& key,
    const TimeValuePair& value,
    uint16_t category,
    bool skipStateCheck = false);
````

````cpp
bool BucketMap::putDataPointWithId(
    BucketedTimeSeries* timeSeries,
    uint32_t timeSeriesId,
    const TimeValuePair& value,
    uint16_t category) {
  uint32_t b = bucket(value.unixTime);
  bool added = timeSeries->put(b, value, &storage_, timeSeriesId, &category);
  if (added) {
    logWriter_->logData(shardId_, timeSeriesId, value.unixTime, value.value);
  }
  return added;
}
````

- [BucketedTimeSeries.h](https://github.com/facebookincubator/beringei/blob/master/beringei/lib/BucketedTimeSeries.h)

### On Disk

- Section 4.3 in paper

#### Log

**only one append-only log per shard, so values within a shard are interleaved across time series**

- [DataLog.h](https://github.com/facebookincubator/beringei/blob/master/beringei/lib/DataLog.h)
- use 32-bit int to identify Time Value pair in log (thus **much larger** than in memory blocks)
- not WAL, use large buffer to increase performance
- [DataLog.cpp](https://github.com/facebookincubator/beringei/blob/master/beringei/lib/DataLog.cpp#L94)

````
void DataLogWriter::append(uint32_t id, int64_t unixTime, double value)

append(id, unixTime, value) {
  tmpBits
  tmpBits += id // NOTE: there is control bit for different length of id

  // calculate and store delta of time
  delta = unixTime - lastUnixTime
  controlBit = controlBitBasedOnLength(delta)
  tmpBits += controlBit
  tmpBits += delta

  // calculate and store delta of value
  xorResult = value ^ previousValues[id]
  // NOTE: this is a simplified in correct version of the code, but more intuitive
  tmpBits += leadingZeros
  tmpBits += xorResult

  lastUnixTime = unixTime
  previousValues[id] = id

  memcpy(buffer, tmpBits)
}
````

#### Compressed Block

> Every two hours, Gorilla copies the compressed block data
to disk, as this format is much smaller than the log files.
There is one complete block file for every two hours worth
of data. It has two sections: a set of consecutive 64kB slabs
of data blocks, copied directly as they appear in memory,
and a list of <time series ID, data block pointer> pairs.
Once a block file is complete, Gorilla touches a checkpoint
file and deletes the corresponding logs.

- [DataBlock.h](https://github.com/facebookincubator/beringei/blob/master/beringei/lib/DataBlock.h)
- 64KB Data Blocks
  - [ ] are blocks interleaved across series
- list<time series ID, databloc pointer>
- [BucketStorage::write](https://github.com/facebookincubator/beringei/blob/master/beringei/lib/BucketStorage.cpp#L357)
- use zlib for compression I suppose?

````cpp
const uint32_t kDataBlockSize = 65536;

struct DataBlock {
  char data[kDataBlockSize];
};
````
