# Storage: Disk

## Ref

mmap

- https://github.com/google/codesearch/tree/master/index
- http://beej.us/guide/bgipc/output/html/multipage/mmap.html
- https://github.com/arpith/mmapd

Prometheus

- https://fabxc.org/blog/2017-04-10-writing-a-tsdb/

InfluxDB

- https://docs.influxdata.com/influxdb/v1.2/concepts/storage_engine/
- https://github.com/influxdata/influxdb/blob/master/tsdb/engine/tsm1/DESIGN.md

## Design

- [ ] TODO: sync with the observation in [Xephon-S](https://github.com/xephonhq/xephon-s/issues/4)

DO NOT

- store a series with different granularity in a single file
  - **people don't query with mixed granularity**
  - different from tree structure in Akumuli and BTrDB, which upper level node store pre-aggregated value
  - we should let user specify what granularity they want, or provide a default rule to them
    - [ ] figure out [InfluxDB's retention policy](https://docs.influxdata.com/influxdb/v1.2/query_language/database_management/#retention-policy-management)
    - https://github.com/influxdata/influxdb/issues/7198 need to create a continuous query
    - auto rollup is not implemented
- support update and delete
  - memory cache is large enough to deal with data arrive out of order
  - bulk load backfill can be treated differently

DO

- store different series in a single file
  - run out of inode if one file per series, as mentioned by [prometheus guys](https://fabxc.org/blog/2017-04-10-writing-a-tsdb/)
- store common aggregation in meta to speed up scan
  - min value & time, max value & time, start time & value, end time & value

TODO

- [ ] how many space for meta data in a file, or create a separate file for series name and pointer to data blocks
  - [ ] you can know how many series in a file when you want to use fixed file size?
- [ ] when to dump memory into file, by size, by time etc.
- [ ] how to organize file
  - prometheus use folder to specify different time range
- [ ] bulk load historical data
- [ ] store data in time increasing order or decreasing order
- [ ] index for time stamp when we use delta + run length encoding
- [ ] separate data and meta into two different files?
 
Compression

- delta + run length
  - http://www.dspguide.com/ch27/4.htm signal processing, it mentioned Linear Predictive Coding, which seems to be what Akumli is doing
- it pretty like **signal**, sound wave etc.
- AWS Redshift http://docs.aws.amazon.com/redshift/latest/dg/c_Runlength_encoding.html
  - We do not recommend applying runlength encoding on any column that is designated as a sort key. Range-restricted scans perform better when blocks contain similar numbers of row
- A time-series compression technique and its application to the smart grid
- https://julien.danjou.info/blog/2016/gnocchi-carbonara-timeseries-compression
- https://golang.org/pkg/compress/
  - https://github.com/klauspost/compress a drop in replace that claims to be faster
- https://github.com/golang/snappy
  
## Implementation

- early code can be found on `playground/disk/nocompress_test.go`
- currently we store both data and indexes in one file
  - NOTE: the indexes is not the inverted index.

````
| header | blocks | indexes | footer |
````

Header stores magic number and format version, which is used for identifying file without extension

````
| magic | version |
````

Footer stores the offset of the indexes

````
| indexes offset | magic |
````

Blocks is a list of blocks, written by the order they comes in

- NOTE： there is no number of blocks like in indexes because we write index after we have wrote all the blocks

````
| encoding | number of points | compressed times | compressed values |
````

Indexes is a list of indexes, sorted by SeriesID

````
| number of indexes | index 1 | index 2 |
````

Index contains the meta of series and the position of all blocks of this series

- series name is stored as `__name__` in tags
- NOTE: we can denote the entries count using (total len - tags len) / size(entries)

````
| total len | tags len | json encoded tags | entries |
````

Entries contains the position and aggregation of each blocks

````
| offset | min time | min time's value | max time | max time's value | min value's time | min value | max value's time | max value | 
````