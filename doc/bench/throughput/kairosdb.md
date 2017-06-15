# KairosDB

10s 10c, using docker compose without mounting volume
564, 0.16

````
Total: 564
0.03889225063398691 	 ......................................................
0.06587120199259254 	 ...............................................
0.09460961393750005 	 ..................................
0.14096801581395355 	 ...............
0.1701104841111111 	 ...
0.19785774584210528 	 .............
0.23091559832142858 	 .........
0.2759964175000001 	 ...
0.3161347954285714 	 ..
0.382333963 	 ...
0.43071079 	 .
0.798383969 	 
0.880128464 	 
0.9219981572222223 	 ...
0.982492536 	 
1.2538152713750004 	 ..
1.3178153049999999 	 
1.346705595 	 
1.388585408 	 .
1.4288116889999998 	 ..
INFO[0013] total time 11.034879 s pkg=k.bench.reporter 
INFO[0013] total request 564 pkg=k.bench.reporter 
INFO[0013] fastest 0.012433 s pkg=k.bench.reporter 
INFO[0013] slowest 1.439205 s pkg=k.bench.reporter 
INFO[0013] average 0.165553 s pkg=k.bench.reporter 
INFO[0013] total request size 101556660 pkg=k.bench.reporter 
INFO[0013] toatl response size 0 pkg=k.bench.reporter 
INFO[0013] 204: 564 pkg=k.bench.reporter 
log:
  level: info
  color: true
  source: false
mode: local
loader:
  target: kairosdb
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

10s, 100c, docker compose, 562, 1.9s


```` 
INFO[0012] basic report finished by channel pkg=k.bench.reporter 
Total: 562
0.11147881128947372 	 .............
0.30663771190384626 	 ..................
0.5760724463999998 	 ................
0.8555596358260867 	 ................
1.116227279105263 	 ......
1.4393232677142858 	 ....
1.729560999976191 	 ..............
1.9684073728648648 	 .............
2.213225110339623 	 ..................
2.4243764266097574 	 ..............
2.635098603422222 	 ................
2.8825483789722224 	 ............
3.211152123066667 	 .....
3.6563970556000003 	 ........
3.945037518842106 	 ......
4.174839036615385 	 ....
4.440750438666665 	 ...
4.672741370200001 	 ...
6.6378540610000005 	 
7.017274448 	 
INFO[0013] total time 11.820552 s pkg=k.bench.reporter 
INFO[0013] total request 562 pkg=k.bench.reporter 
INFO[0013] fastest 0.019277 s pkg=k.bench.reporter 
INFO[0013] slowest 7.017274 s pkg=k.bench.reporter 
INFO[0013] average 1.934406 s pkg=k.bench.reporter 
INFO[0013] total request size 101196530 pkg=k.bench.reporter 
INFO[0013] toatl response size 0 pkg=k.bench.reporter 
INFO[0013] 204: 562 pkg=k.bench.reporter 
log:
  level: info
  color: true
  source: false
mode: local
loader:
  target: kairosdb
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
INFO[0013] bench finished pkg=k.cmd.bench 
````

10s (actually 48s ...) 1000c

````
INFO[0049] basic report finished by channel pkg=k.bench.reporter 
Total: 2220
0.7701286159892471 	 ........
1.839095116940298 	 ......
3.7437321259529392 	 .......
5.09696370994505 	 ................
6.855414177000002 	 .......
8.220800388983145 	 ................
9.725535745941173 	 .......
11.598194327134777 	 ....................
13.595598881139706 	 ............
15.088106019868263 	 ...............
16.393714322233627 	 ...................
17.71110649291667 	 ...........
18.953105772780024 	 ..................
20.229421693177777 	 ........
21.480499082274715 	 ........
23.071107117325575 	 ...
24.787787846696972 	 .....
26.4669317105 	 ...
27.974809438083327 	 .
29.81820356696774 	 ..
INFO[0050] total time 48.617459 s pkg=k.bench.reporter 
INFO[0050] total request 2220 pkg=k.bench.reporter 
INFO[0050] fastest 0.019157 s pkg=k.bench.reporter 
INFO[0050] slowest 30.000194 s pkg=k.bench.reporter 
INFO[0050] average 13.360634 s pkg=k.bench.reporter 
INFO[0050] total request size 399744300 pkg=k.bench.reporter 
INFO[0050] toatl response size 0 pkg=k.bench.reporter 
INFO[0050] 204: 2201 pkg=k.bench.reporter 
INFO[0050] 0: 19 pkg=k.bench.reporter 
log:
  level: info
  color: true
  source: false
mode: local
loader:
  target: kairosdb
  reporter: basic
  limitBy: time
  points: 100000000
  series: 100
  time: 10
  workerNum: 1000
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
INFO[0050] bench finished pkg=k.cmd.bench 
````