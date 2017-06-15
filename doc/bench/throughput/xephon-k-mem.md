# Xephon-K Mem

- 10s, 10 client, disable index TODO: it has concurrent map access again??

NOTE: this result is when generator yield 0 points per series, the default config is zero .... 
https://github.com/xephonhq/xephon-k/issues/57

113608, avg: 0.000853

````
Total: 113608
0.0005293175084013952 	 ............................................................................................................................................................................................
0.0014072892387543261 	 
0.002857744191028228 	 ......
0.004784186681341721 	 
0.00605415579649123 	 
0.007377884203319499 	 
0.009599022158995819 	 
0.011469762726190475 	 
0.013272907014084508 	 
0.014387950718749999 	 
0.015569014336842104 	 
0.01670671983333332 	 
0.017910694215517255 	 
0.01923185354609929 	 
0.02058375948091603 	 
0.021864585696721325 	 
0.023684358263157904 	 
0.024892185388888887 	 
0.03280773377777778 	 
0.03520795 	 
INFO[0012] total time 10.000744 s pkg=k.bench.reporter 
INFO[0012] total request 113608 pkg=k.bench.reporter 
INFO[0012] fastest 0.000072 s pkg=k.bench.reporter 
INFO[0012] slowest 0.035208 s pkg=k.bench.reporter 
INFO[0012] average 0.000853 s pkg=k.bench.reporter 
INFO[0012] total request size 11817730 pkg=k.bench.reporter 
INFO[0012] toatl response size 340824 pkg=k.bench.reporter 
INFO[0012] 200: 113608 pkg=k.bench.reporter 
log:
  level: info
  color: true
  source: false
mode: local
loader:
  target: xephonk
  reporter: basic
  limitBy: time
  points: 1000000
  series: 1
  time: 10
  workerNum: 10
  workerTimeout: 30
generator:
  timeInterval: 0
  timeNoise: false
  pointsPerSeries: 0
  numSeries: 0
targets:
  influxdb:
    host: localhost
    port: 8086
    url: write?db=sb
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
INFO[0012] bench finished pkg=k.cmd.bench 
````

10s, 10c, right 1535

````
INFO[0012] total time 10.120981 s pkg=k.bench.reporter 
INFO[0012] total request 1535 pkg=k.bench.reporter 
INFO[0012] fastest 0.017193 s pkg=k.bench.reporter 
INFO[0012] slowest 0.099905 s pkg=k.bench.reporter 
INFO[0012] average 0.047307 s pkg=k.bench.reporter 
INFO[0012] total request size 276451965 pkg=k.bench.reporter 
INFO[0012] toatl response size 4605 pkg=k.bench.reporter 
INFO[0012] 200: 1535 pkg=k.bench.reporter 
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
INFO[0012] bench finished pkg=k.cmd.bench 
````

10s, 100c, 1753, 0.58

````
INFO[0012] basic report finished by channel pkg=k.bench.reporter 
Total: 1753
0.05252847954545454 	 .
0.1497975479757281 	 .......................
0.3080670730866937 	 ........................................................
0.47522930056571455 	 ...................
0.6251182133883791 	 .....................................
0.7688089279078946 	 .................
0.8920961607454546 	 ..................
1.0644694612527477 	 ..........
1.1876952337692308 	 ....
1.2930189929117646 	 ...
1.411249842590909 	 ..
1.5703640440909086 	 .
1.673480292 	 
1.790571128111111 	 .
1.9030379400000001 	 
2.0365130396666666 	 
2.161930483 	 
2.500472522 	 
2.65827995 	 
3.0063469785 	 
INFO[0013] total time 11.228774 s pkg=k.bench.reporter 
INFO[0013] total request 1753 pkg=k.bench.reporter 
INFO[0013] fastest 0.027413 s pkg=k.bench.reporter 
INFO[0013] slowest 3.049857 s pkg=k.bench.reporter 
INFO[0013] average 0.581593 s pkg=k.bench.reporter 
INFO[0013] total request size 315713547 pkg=k.bench.reporter 
INFO[0013] toatl response size 5259 pkg=k.bench.reporter 
INFO[0013] 200: 1753 pkg=k.bench.reporter 
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
INFO[0013] bench finished pkg=k.cmd.bench 
````

10s, 1000c, too many open file, 3.7s, 2594 ... 

````
INFO[0022] basic report finished by channel pkg=k.bench.reporter 
Total: 3581
0.10855596030293917 	 ..........................................................................
1.2050846354786724 	 ...........
2.076351425080645 	 ......
2.923477089874345 	 ..........
3.972779766271817 	 ......................
4.906069959877864 	 ..............
5.671541675458902 	 ........
6.56601188530216 	 .......
7.501207431262407 	 .......
8.28591844262626 	 .....
8.89988854341611 	 ........
9.48988844035052 	 .....
10.25442430036429 	 .......
11.332283565628208 	 ....
12.244584550791664 	 .
13.060022016105263 	 .
14.034195672421053 	 .
15.082105837625 	 
16.105758850199997 	 
16.765923756 	 
INFO[0023] total time 20.691676 s pkg=k.bench.reporter 
INFO[0023] total request 3581 pkg=k.bench.reporter 
INFO[0023] fastest 0.000040 s pkg=k.bench.reporter 
INFO[0023] slowest 16.765924 s pkg=k.bench.reporter 
INFO[0023] average 3.715192 s pkg=k.bench.reporter 
INFO[0023] total request size 644934519 pkg=k.bench.reporter 
INFO[0023] toatl response size 7782 pkg=k.bench.reporter 
INFO[0023] 0: 987 pkg=k.bench.reporter 
INFO[0023] 200: 2594 pkg=k.bench.reporter 
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