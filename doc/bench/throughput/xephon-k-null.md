# Xephon-K Null

10s, 10c, 1407, 0.059

````
Total: 1407
0.015683903 	 
0.0253293515 	 
0.03133023744444444 	 .
0.036211688423076926 	 ...
0.040174712999999994 	 ......
0.04557385726315788 	 ........................
0.04997236222962964 	 ...................
0.0533814260448718 	 ......................
0.057669672732793525 	 ...................................
0.06191504331007749 	 ..................
0.06644933333936652 	 ...............................
0.07135827299999999 	 ..............
0.07485128701754387 	 ........
0.0785662916 	 ....
0.08195361674193548 	 ....
0.08741449133333332 	 ....
0.09515455137500001 	 .
0.10306686749999999 	 
0.1087344955 	 
0.124013587 	 
INFO[0013] total time 10.118765 s pkg=k.bench.reporter 
INFO[0013] total request 1407 pkg=k.bench.reporter 
INFO[0013] fastest 0.015145 s pkg=k.bench.reporter 
INFO[0013] slowest 0.124014 s pkg=k.bench.reporter 
INFO[0013] average 0.059257 s pkg=k.bench.reporter 
INFO[0013] total request size 253399293 pkg=k.bench.reporter 
INFO[0013] toatl response size 4221 pkg=k.bench.reporter 
INFO[0013] 200: 1407 pkg=k.bench.reporter 
log:
  level: info
  color: true
  source: false
mode: local
loader:
  target: xephonk
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

10s, 100c, 1945, 0.513

````
Total: 1945
0.07662645130894306 	 ............
0.198655085172956 	 ................................
0.2991982508578679 	 ....................
0.38388590561596975 	 ...........................
0.486382264061644 	 ..............................
0.5956476913557691 	 .....................
0.6866650859017857 	 ...........
0.7642428613084111 	 ...........
0.8685184986220472 	 .............
0.9776862855373135 	 ......
1.0711245849459456 	 ...
1.1798269245853656 	 ....
1.272346432214286 	 .
1.3722082778666664 	 .
1.4912131262500001 	 .
1.57994103825 	 
1.680621274 	 
1.86014916075 	 
1.953064937 	 
2.292784733 	 
INFO[0015] total time 10.803580 s pkg=k.bench.reporter 
INFO[0015] total request 1945 pkg=k.bench.reporter 
INFO[0015] fastest 0.014745 s pkg=k.bench.reporter 
INFO[0015] slowest 2.292785 s pkg=k.bench.reporter 
INFO[0015] average 0.513403 s pkg=k.bench.reporter 
INFO[0015] total request size 350292555 pkg=k.bench.reporter 
INFO[0015] toatl response size 5835 pkg=k.bench.reporter 
INFO[0015] 200: 1945 pkg=k.bench.reporter 
log:
  level: info
  color: true
  source: false
mode: local
loader:
  target: xephonk
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
INFO[0015] bench finished pkg=k.cmd.bench
````

10s, 1000c, meet too many open files again, last time, using docker seems to solve this problem?
2.50

````
Total: 4886
0.025048396133365645 	 ....................................................................................
0.9255612686530176 	 ..................
1.7314230289446257 	 ............
2.354237470726458 	 .........
2.992571925317343 	 ...........
3.576663612965988 	 ......
4.449020903450674 	 ..................
5.34087896286636 	 ........
6.020238054726185 	 ......
6.695190262861539 	 .....
7.519384754915154 	 ......
8.43046741831818 	 ...
9.489908734290907 	 ..
10.452952540709678 	 .
11.15994354252174 	 
11.98495257577143 	 .
12.963767908699998 	 
13.727257348 	 
15.449447533296293 	 .
16.237691126999998 	 
INFO[0023] total time 18.504740 s pkg=k.bench.reporter 
INFO[0023] total request 4886 pkg=k.bench.reporter 
INFO[0023] fastest 0.000044 s pkg=k.bench.reporter 
INFO[0023] slowest 16.654811 s pkg=k.bench.reporter 
INFO[0023] average 2.500904 s pkg=k.bench.reporter 
INFO[0023] total request size 879963714 pkg=k.bench.reporter 
INFO[0023] toatl response size 8544 pkg=k.bench.reporter 
INFO[0023] 0: 2038 pkg=k.bench.reporter 
INFO[0023] 200: 2848 pkg=k.bench.reporter 
log:
  level: info
  color: true
  source: false
mode: local
loader:
  target: xephonk
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
INFO[0023] bench finished pkg=k.cmd.bench 
````