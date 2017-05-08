# Collector

For Linux systems, all use `/proc`

## Issues

- [System Metrics Collector #21](https://github.com/xephonhq/xephon-k/issues/21)
- [Process Metrics Collector #22](https://github.com/xephonhq/xephon-k/issues/22)
- [Container metrics collector #31](https://github.com/xephonhq/xephon-k/issues/31) 
  - using cAdvisor
- [ ] Handle the many duplicate code, really annoying

## Ref 

- [gopsutil](https://github.com/shirou/gopsutil)
- [Prometheus node exporter system](https://github.com/prometheus/node_exporter/blob/master/collector/stat_linux.go#L87)
- [Elastic Search Metric Beat System](https://github.com/elastic/beats/tree/master/metricbeat/module/system)
  - it is also using gopsutil like InfluxDB's telegraf