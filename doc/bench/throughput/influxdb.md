# InfluxDB

10s 10c 1623

````
INFO[0012] basic report finished by channel pkg=k.bench.reporter 
Total: 1623
0.02054284 	 
0.025581650416666667 	 ..
0.03061365166101694 	 .......
0.035206574351351355 	 .............
0.03955302647058824 	 ................
0.04397608606206901 	 .................
0.04851016481456952 	 ..................
0.052955175656249964 	 .......................
0.05966477013620074 	 ..................................
0.06708291470833337 	 .......................
0.07339053885576925 	 ............
0.07767607813043481 	 ........
0.08259220787499998 	 .....
0.08677535819444444 	 ....
0.09314709521739128 	 .....
0.10052046128571428 	 .
0.10723624899999999 	 
0.11200186366666666 	 
0.120475895 	 
0.167624053 	 
INFO[0013] total time 10.131469 s pkg=k.bench.reporter 
INFO[0013] total request 1623 pkg=k.bench.reporter 
INFO[0013] fastest 0.020543 s pkg=k.bench.reporter 
INFO[0013] slowest 0.167624 s pkg=k.bench.reporter 
INFO[0013] average 0.057133 s pkg=k.bench.reporter 
INFO[0013] total request size 762810000 pkg=k.bench.reporter 
INFO[0013] toatl response size 0 pkg=k.bench.reporter 
INFO[0013] 204: 1623 pkg=k.bench.reporter 
log:
  level: info
  color: true
  source: false
mode: local
loader:
  target: influxdb
  reporter: basic
  limitBy: time
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
INFO[0013] bench finished pkg=k.cmd.bench
````

10s, 100c, 1718, 0.6s

````
INFO[0029] basic report finished by channel pkg=k.bench.reporter 
Total: 1718
0.05990231680555553 	 ....
0.1280447141458334 	 .....
0.22172316906451608 	 .......
0.2902368711842106 	 ....
0.3523123217142857 	 ....
0.43115505891304345 	 .............
0.49646052039655125 	 .................................
0.569806209795919 	 .......................................
0.6449458539999998 	 ..........................
0.7190059840899475 	 ......................
0.8077944264931504 	 ................
0.8994903978829787 	 ..........
1.0315257449999997 	 ..
1.1065272299615385 	 ...
1.1620830149999999 	 
1.2258270083999998 	 .
1.3398150854000002 	 
1.4133713003333332 	 
1.4806555065000002 	 
1.569190699375 	 
INFO[0030] total time 10.851907 s pkg=k.bench.reporter 
INFO[0030] total request 1718 pkg=k.bench.reporter 
INFO[0030] fastest 0.018583 s pkg=k.bench.reporter 
INFO[0030] slowest 1.600794 s pkg=k.bench.reporter 
INFO[0030] average 0.601662 s pkg=k.bench.reporter 
INFO[0030] total request size 807460000 pkg=k.bench.reporter 
INFO[0030] toatl response size 0 pkg=k.bench.reporter 
INFO[0030] 204: 1718 pkg=k.bench.reporter 
log:
  level: info
  color: true
  source: false
mode: local
loader:
  target: influxdb
  reporter: basic
  limitBy: time
  points: 100000000
  series: 100
  time: 10
  workerNum: 100
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
INFO[0030] bench finished pkg=k.cmd.bench 
````