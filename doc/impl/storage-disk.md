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

Compression

- delta + run length
  - http://www.dspguide.com/ch27/4.htm signal processing, it mentioned Linear Predictive Coding, which seems to be what Akumli is doing
- it pretty like **signal**, sound wave etc.
- AWS Redshift http://docs.aws.amazon.com/redshift/latest/dg/c_Runlength_encoding.html
  - We do not recommend applying runlength encoding on any column that is designated as a sort key. Range-restricted scans perform better when blocks contain similar numbers of row
- A time-series compression technique and its application to the smart grid
- https://julien.danjou.info/blog/2016/gnocchi-carbonara-timeseries-compression
