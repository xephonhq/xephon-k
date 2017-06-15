# Xephon-K Disk

10s 10c 1472, 0.05s 

````
Total: 1472
0.014978058 	 
0.0214171595 	 
0.025039542999999997 	 
0.028800024500000004 	 .
0.03258977300000001 	 ..
0.03835688349999999 	 ..............
0.04490087190361447 	 .................................
0.04965365493296092 	 ........................
0.05311992185029937 	 ......................
0.05748738035156252 	 ..................................
0.061714163086092705 	 ....................
0.06566945549999997 	 ..............
0.06938242887951805 	 ...........
0.07363191050000005 	 ........
0.0781949465588235 	 ....
0.08293006346666669 	 ..
0.0865497556923077 	 .
0.09027460271428571 	 
0.09611980475000001 	 
0.10110105250000001 	 
INFO[0014] total time 10.147098 s pkg=k.bench.reporter 
INFO[0014] total request 1472 pkg=k.bench.reporter 
INFO[0014] fastest 0.014978 s pkg=k.bench.reporter 
INFO[0014] slowest 0.102321 s pkg=k.bench.reporter 
INFO[0014] average 0.055538 s pkg=k.bench.reporter 
INFO[0014] total request size 265105728 pkg=k.bench.reporter 
INFO[0014] toatl response size 4416 pkg=k.bench.reporter 
INFO[0014] 200: 1472 pkg=k.bench.reporter 
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
INFO[0014] bench finished pkg=k.cmd.bench
````