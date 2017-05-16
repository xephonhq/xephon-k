# Time Structured Merge Tree

- https://docs.influxdata.com/influxdb/v1.2/concepts/storage_engine/
- https://github.com/influxdata/influxdb/blob/master/tsdb/engine/tsm1/DESIGN.md

- [ ] The index is composed of a sequence of index entries ordered lexicographically by key and then by time
- If we we limit file sizes to 4GB, we would use 4 bytes for each pointer