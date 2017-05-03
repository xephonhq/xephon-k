# Round Robin Database

## TODO

- [x] fixed interval and predefined size means only need to store the start timestamp?
  - 'timestamp being inferred from its position in the archive' found this in whisper doc, I think that's how RRD Tool works
- [ ] If I have 3 RRA with interval of 1 hour, 1 minutes and 1 second, when and where is the consolidation function executed?
- [x] Is there any requirement for the size of multiple RRAs, i.e. the RRA using one second should be larger than 60 so the values to be dropped can be consolidated and put into another more coarse grained RRA.
  - answered by graphite whisper http://graphite.readthedocs.io/en/latest/whisper.html#archives-retention-and-precision


## Ref

- rrdtool https://github.com/oetiker/rrdtool-1.x
  - http://oss.oetiker.ch/rrdtool/doc/rrdtool.en.html
- http://www.loriotpro.com/Products/On-line_Documentation_V5/LoriotProDoc_EN/V22-RRD_Collector_RRD_Manager/V22-A1_Introduction_RRD_EN.htm
- http://cuddletech.com/articles/rrd/ar01s02.html

## Meta

- On Disk Size: fixed
- Time Interval: constant interval

## Terms

- RRD: Round Robin Database
- RRA: Round Robin Archive
- CF: Consolidation functions (MAX, MIN, AVG, LAST)
  - [ ] I think is called `roll-up` or `aggregation` in other TSDBs
- PDP: Primary data points

## Architecture

- **One RRD can contains multiple RRA**
  - [ ] If I have 3 RRA with interval of 1 hour, 1 minutes and 1 second, when and where is the consolidation function executed? - [ ] Is there any requirement
  for the size of multiple RRAs, i.e. the RRA using one second should be larger than 60 so the values to be dropped can be consolidated and put into another more coarse grained RRA.
- Data values of the same consolidation setup are stored into Round Robin Archives (RRA)
- 'Unknown' is used when data point is not available when they should be

Command for creating RRD file

````
--start N --step 300 \
DS:probe1-temp:GAUGE:600:55:95 \
DS:probe2-temp:GAUGE:600:55:95 \
DS:probe3-temp:GAUGE:600:55:95 \
DS:probe4-temp:GAUGE:600:55:95 \
RRA:MIN:0.5:12:1440 \
RRA:MAX:0.5:12:1440 \
RRA:AVERAGE:0.5:1:1440
````

- step 'Specifies the base interval in seconds with which data will be fed into the RRD.'
- `DS:ds-name:{GAUGE | COUNTER | DERIVE | DCOUNTER | DDERIVE | ABSOLUTE}:heartbeat:min:max`
- `RRA:{AVERAGE | MIN | MAX | LAST}:xff:steps:rows`
  - steps defines how many of these primary data points are used to build a consolidated data point which then goes into the archive
  - if step = 1, means no consolidation function is actually needed
