# Storage: Disk

## Ref 

mmap

- https://github.com/google/codesearch/tree/master/index
- http://beej.us/guide/bgipc/output/html/multipage/mmap.html
- https://github.com/arpith/mmapd

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