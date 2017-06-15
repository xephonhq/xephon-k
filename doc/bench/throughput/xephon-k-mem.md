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