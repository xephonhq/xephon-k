# KairosDB

10s 10c, using docker compose without mounting volume

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