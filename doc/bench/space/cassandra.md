# Cassandra

1,000,000 5.7M

````
⇒  xkb --limit points                  
log:
level: info
color: true
source: false
mode: local
loader:
target: xephonk
reporter: basic
limitBy: points
points: 1000000
series: 100
time: 10
workerNum: 10
workerTimeout: 30
generator:
timeInterval: 1
timeNoise: false
pointsPerSeries: 100
numSeries: 10
targets:
influxdb:
 host: localhost
 port: 8086
 url: write?db=xb
 timeout: 30
xephonk:
 host: localhost
 port: 2333
 url: write
 timeout: 30
kairosdb:
 host: localhost
 port: 8080
 url: api/v1/datapoints
 timeout: 30
Do you want to proceed? [Y/N]y
INFO[0010] worker started pkg=k.bench.worker 
INFO[0010] worker started pkg=k.bench.worker 
INFO[0010] worker started pkg=k.bench.worker 
INFO[0010] worker started pkg=k.bench.worker 
INFO[0010] worker started pkg=k.bench.worker 
INFO[0010] worker started pkg=k.bench.worker 
INFO[0010] worker started pkg=k.bench.worker 
INFO[0010] worker started pkg=k.bench.worker 
INFO[0010] worker started pkg=k.bench.worker 
INFO[0010] worker started pkg=k.bench.worker 
INFO[0016] generator stopped after 1000000 points pkg=k.bench 
INFO[0016] close data channel pkg=k.bench 
INFO[0016] worker finished by input channel pkg=k.bench.worker 
INFO[0016] worker finished by input channel pkg=k.bench.worker 
INFO[0016] worker finished by input channel pkg=k.bench.worker 
INFO[0016] worker finished by input channel pkg=k.bench.worker 
INFO[0016] worker finished by input channel pkg=k.bench.worker 
INFO[0016] worker finished by input channel pkg=k.bench.worker 
INFO[0016] worker finished by input channel pkg=k.bench.worker 
INFO[0016] worker finished by input channel pkg=k.bench.worker 
INFO[0016] worker finished by input channel pkg=k.bench.worker 
INFO[0016] worker finished by input channel pkg=k.bench.worker 
INFO[0016] basic report finished by channel pkg=k.bench.reporter 
Total: 10000
0.0022296591022727273 	 .
0.0034435686221429653 	 ...........................................................................................................................
0.00586344693700787 	 ......................................
0.008388722092348271 	 ......................
0.011266087155844153 	 ....
0.013116105266666667 	 ...
0.01564649299224806 	 ..
0.018326106750000012 	 .
0.020744893818181814 	 
0.023335760400000002 	 
0.026754299000000002 	 
0.044450648599999996 	 
0.0462308615 	 
0.04903623783333333 	 
0.051487911500000004 	 
0.05425395715384614 	 
0.057249247705882356 	 
0.060334449111111105 	 
0.06393049499999999 	 
0.06815353833333333 	 
INFO[0017] run time 5.733265 s pkg=k.bench.reporter 
INFO[0017] total request 10000 pkg=k.bench.reporter 
INFO[0017] fastest 895463 pkg=k.bench.reporter 
INFO[0017] slowest 68439493 pkg=k.bench.reporter 
INFO[0017] total request size 18990000 pkg=k.bench.reporter 
INFO[0017] toatl response size 30000 pkg=k.bench.reporter 
INFO[0017] 200: 10000 pkg=k.bench.reporter 
INFO[0017] bench finished pkg=k.cmd.bench 

````

````
root@af2b4b0898ef:/var/lib/cassandra/data/xephon/metrics_int-061feae051f211e798e26d2c86545d91# du -sh * 
4.0K	backups
4.0K	mc-1-big-CompressionInfo.db
5.7M	mc-1-big-Data.db
4.0K	mc-1-big-Digest.crc32
4.0K	mc-1-big-Filter.db
12K	mc-1-big-Index.db
8.0K	mc-1-big-Statistics.db
4.0K	mc-1-big-Summary.db
4.0K	mc-1-big-TOC.txt
````

10, 000, 000 58M, NOTE: it is using 100 points per series, which I think might be why it is so slow ... I will increase that 
to 1,000 to see if it speeds up on other targets

````
⇒  xkb --limit points
log:
  level: info
  color: true
  source: false
mode: local
loader:
  target: xephonk
  reporter: basic
  limitBy: points
  points: 10000000
  series: 100
  time: 10
  workerNum: 10
  workerTimeout: 30
generator:
  timeInterval: 1
  timeNoise: false
  pointsPerSeries: 100
  numSeries: 10
targets:
  influxdb:
    host: localhost
    port: 8086
    url: write?db=xb
    timeout: 30
  xephonk:
    host: localhost
    port: 2333
    url: write
    timeout: 30
  kairosdb:
    host: localhost
    port: 8080
    url: api/v1/datapoints
    timeout: 30
Do you want to proceed? [Y/N]y
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0052] generator stopped after 10000000 points pkg=k.bench 
INFO[0052] close data channel pkg=k.bench 
INFO[0052] worker finished by input channel pkg=k.bench.worker 
INFO[0052] worker finished by input channel pkg=k.bench.worker 
INFO[0052] worker finished by input channel pkg=k.bench.worker 
INFO[0052] worker finished by input channel pkg=k.bench.worker 
INFO[0052] worker finished by input channel pkg=k.bench.worker 
INFO[0052] worker finished by input channel pkg=k.bench.worker 
INFO[0052] worker finished by input channel pkg=k.bench.worker 
INFO[0052] worker finished by input channel pkg=k.bench.worker 
INFO[0052] worker finished by input channel pkg=k.bench.worker 
INFO[0052] worker finished by input channel pkg=k.bench.worker 
INFO[0052] basic report finished by channel pkg=k.bench.reporter 
Total: 100000
0.0033398144862595205 	 ................................................................................................................................................................
0.005965555984695432 	 ................................
0.009283370952029522 	 ...
0.012445071919896631 	 
0.015580992948979591 	 
0.019303649494117652 	 
0.023690358423076922 	 
0.028695261 	 
0.032111190222222225 	 
0.035243050892307684 	 
0.03944567668421053 	 
0.04325521556020942 	 
0.0463956313248175 	 
0.05032273020588234 	 
0.05347614522222222 	 
0.05774266392857143 	 
0.06214909592307693 	 
0.0677891166 	 
0.07089879975 	 
0.07406825633333333 	 
INFO[0053] run time 44.595655 s pkg=k.bench.reporter 
INFO[0053] total request 100000 pkg=k.bench.reporter 
INFO[0053] fastest 960517 pkg=k.bench.reporter 
INFO[0053] slowest 75119183 pkg=k.bench.reporter 
INFO[0053] total request size 189900000 pkg=k.bench.reporter 
INFO[0053] toatl response size 300000 pkg=k.bench.reporter 
INFO[0053] 200: 100000 pkg=k.bench.reporter 
INFO[0053] bench finished pkg=k.cmd.bench 
````

````
4.0K	backups
28K	mc-5-big-CompressionInfo.db
58M	mc-5-big-Data.db
4.0K	mc-5-big-Digest.crc32
4.0K	mc-5-big-Filter.db
96K	mc-5-big-Index.db
8.0K	mc-5-big-Statistics.db
4.0K	mc-5-big-Summary.db
4.0K	mc-5-big-TOC.txt
````

100,000,000 528 M, peak memory 11GB

- `tar czf a.tar.gz *.db`, compressed file is 133 MB

````
root@5f850cc0b142:/var/lib/cassandra/data/xephon/metrics_int-50a04e9051f511e7a59d6d2c86545d91# du -sh *
133M	a.tar.gz
4.0K	backups
112K	mc-23-big-CompressionInfo.db
229M	mc-23-big-Data.db
4.0K	mc-23-big-Digest.crc32
4.0K	mc-23-big-Filter.db
412K	mc-23-big-Index.db
8.0K	mc-23-big-Statistics.db
4.0K	mc-23-big-Summary.db
4.0K	mc-23-big-TOC.txt
112K	mc-44-big-CompressionInfo.db
229M	mc-44-big-Data.db
4.0K	mc-44-big-Digest.crc32
4.0K	mc-44-big-Filter.db
408K	mc-44-big-Index.db
8.0K	mc-44-big-Statistics.db
4.0K	mc-44-big-Summary.db
4.0K	mc-44-big-TOC.txt
28K	mc-47-big-CompressionInfo.db
57M	mc-47-big-Data.db
4.0K	mc-47-big-Digest.crc32
4.0K	mc-47-big-Filter.db
104K	mc-47-big-Index.db
8.0K	mc-47-big-Statistics.db
4.0K	mc-47-big-Summary.db
4.0K	mc-47-big-TOC.txt
8.0K	mc-48-big-CompressionInfo.db
12M	mc-48-big-Data.db
4.0K	mc-48-big-Digest.crc32
4.0K	mc-48-big-Filter.db
24K	mc-48-big-Index.db
8.0K	mc-48-big-Statistics.db
4.0K	mc-48-big-Summary.db
4.0K	mc-48-big-TOC.txt
````

````
⇒  xkb --limit points
log:
  level: info
  color: true
  source: false
mode: local
loader:
  target: xephonk
  reporter: basic
  limitBy: points
  points: 100000000
  series: 100
  time: 10
  workerNum: 10
  workerTimeout: 30
generator:
  timeInterval: 1
  timeNoise: false
  pointsPerSeries: 10000
  numSeries: 10
targets:
  influxdb:
    host: localhost
    port: 8086
    url: write?db=xb
    timeout: 30
  xephonk:
    host: localhost
    port: 2333
    url: write
    timeout: 30
  kairosdb:
    host: localhost
    port: 8080
    url: api/v1/datapoints
    timeout: 30
Do you want to proceed? [Y/N]y
INFO[0028] worker started pkg=k.bench.worker 
INFO[0028] worker started pkg=k.bench.worker 
INFO[0028] worker started pkg=k.bench.worker 
INFO[0028] worker started pkg=k.bench.worker 
INFO[0028] worker started pkg=k.bench.worker 
INFO[0028] worker started pkg=k.bench.worker 
INFO[0028] worker started pkg=k.bench.worker 
INFO[0028] worker started pkg=k.bench.worker 
INFO[0028] worker started pkg=k.bench.worker 
INFO[0028] worker started pkg=k.bench.worker 
INFO[0314] generator stopped after 100000000 points pkg=k.bench 
INFO[0314] close data channel pkg=k.bench 
INFO[0314] worker finished by input channel pkg=k.bench.worker 
INFO[0314] worker finished by input channel pkg=k.bench.worker 
INFO[0314] worker finished by input channel pkg=k.bench.worker 
INFO[0314] worker finished by input channel pkg=k.bench.worker 
INFO[0314] worker finished by input channel pkg=k.bench.worker 
INFO[0314] worker finished by input channel pkg=k.bench.worker 
INFO[0314] worker finished by input channel pkg=k.bench.worker 
INFO[0314] worker finished by input channel pkg=k.bench.worker 
INFO[0314] worker finished by input channel pkg=k.bench.worker 
INFO[0314] worker finished by input channel pkg=k.bench.worker 
INFO[0314] basic report finished by channel pkg=k.bench.reporter 
Total: 10000
0.10139872650000001 	 
0.1250772921203208 	 .......
0.15888912001463396 	 ............
0.18788544288105727 	 ..................
0.21725582411316371 	 .........................
0.2505242037726742 	 ..................................
0.2854833870690365 	 ................................................
0.3192575442346488 	 ...........................
0.3543575349254566 	 ..............
0.39478698483950575 	 ......
0.4324088514531249 	 ..
0.47090578787234033 	 
0.5150659387142857 	 
0.5527264906666666 	 
0.5937416911875001 	 
0.6501657372499999 	 
0.6940864644285715 	 
0.758047297 	 
0.804213103 	 
0.84117449225 	 
INFO[0315] run time 286.120063 s pkg=k.bench.reporter 
INFO[0315] total request 10000 pkg=k.bench.reporter 
INFO[0315] fastest 77673081 pkg=k.bench.reporter 
INFO[0315] slowest 855101446 pkg=k.bench.reporter 
INFO[0315] total request size 1800990000 pkg=k.bench.reporter 
INFO[0315] toatl response size 30000 pkg=k.bench.reporter 
INFO[0315] 200: 10000 pkg=k.bench.reporter 
INFO[0315] bench finished pkg=k.cmd.bench
````

