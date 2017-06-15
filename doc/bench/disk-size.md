# Disk size

- 1,000,000
- 10,000,000
- 100,000,000
- 1000,000,000

1,000, 000, 19M?? wtf, even bigger than raw encoding? might due to difference between real disk and tmpfs?

````
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

Total: 82180
0.0005146889933370763 	 ..................................................................................
0.0008697091775663236 	 ..........................................
0.0012561395031003477 	 ..................................
0.0016568676530133073 	 ..................
0.00207345217950582 	 .............
0.002609313659885009 	 .....
0.00291582425 	 
0.003223850977853494 	 .
0.003644772858407078 	 
0.0039871978904109555 	 
0.004368673791044774 	 
0.004781278690140846 	 
0.005158234000000003 	 
0.005579800076923074 	 
0.006111763235294117 	 
0.007065272375 	 
0.007436614000000001 	 
0.007799303 	 
0.00813409225 	 
0.008477714 	 
INFO[0014] run time 10.001976 s pkg=k.bench.reporter 
INFO[0014] total request 82180 pkg=k.bench.reporter 
INFO[0014] fastest 223901 pkg=k.bench.reporter 
INFO[0014] slowest 8477714 pkg=k.bench.reporter 
INFO[0014] total request size 156059820 pkg=k.bench.reporter 
INFO[0014] toatl response size 246540 pkg=k.bench.reporter 
INFO[0014] 200: 82180 pkg=k.bench.reporter 
INFO[0014] null reporter has nothing to say pkg=k.bench.reporter 
INFO[0014] bench finished pkg=k.cmd.bench 
````