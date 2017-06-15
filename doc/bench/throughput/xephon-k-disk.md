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

10s, 100c, 1904, 0.5s, too many open files, 29MB, note, the request size is much smaller than 100, 000, 000 points

````
INFO[0012] basic report finished by channel pkg=k.bench.reporter 
Total: 1904
0.024488726363636363 	 .
0.09511561451572323 	 ................
0.2021935584056604 	 ......................
0.31791729478095254 	 ......................
0.40921500762559226 	 ......................
0.501802244524844 	 .................................
0.5873197214069769 	 ..................
0.657748936430168 	 ..................
0.7252306062093024 	 .............
0.7938819955871557 	 ...........
0.8697194897916669 	 .......
0.9414576774444443 	 ..
1.0061535505217392 	 ..
1.0848924998857143 	 ...
1.1614845276190477 	 ..
1.236269681 	 
1.327587829 	 
1.480140467 	 
1.598941585 	 
1.6701778335000002 	 
INFO[0013] total time 10.676639 s pkg=k.bench.reporter 
INFO[0013] total request 1904 pkg=k.bench.reporter 
INFO[0013] fastest 0.014954 s pkg=k.bench.reporter 
INFO[0013] slowest 1.678570 s pkg=k.bench.reporter 
INFO[0013] average 0.505368 s pkg=k.bench.reporter 
INFO[0013] total request size 342908496 pkg=k.bench.reporter 
INFO[0013] toatl response size 5712 pkg=k.bench.reporter 
INFO[0013] 200: 1904 pkg=k.bench.reporter 
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

10s, 1000c, 2.4s too many open files, 2100, 37MB

````
INFO[0017] basic report finished by channel pkg=k.bench.reporter 
Total: 4500
0.06503710464872517 	 .............................................................................................................
0.7460046388979591 	 .............
1.451260495848361 	 ..........
2.1357716998633554 	 .......
2.8905308846896545 	 ...
3.6090073469813944 	 .........
4.1757641951477265 	 ...
4.9283410136956505 	 .
5.57605821540625 	 .
6.401577796608696 	 .
6.975071746681819 	 
7.55564702317143 	 .
8.158142469945947 	 ....
8.944331800109453 	 ........
9.875095056100001 	 ..............
10.63689198375 	 ...
11.244940673999993 	 .
11.834832136965515 	 .
12.369477097285715 	 
13.060684307 	 
INFO[0018] total time 15.925720 s pkg=k.bench.reporter 
INFO[0018] total request 4500 pkg=k.bench.reporter 
INFO[0018] fastest 0.000041 s pkg=k.bench.reporter 
INFO[0018] slowest 13.324191 s pkg=k.bench.reporter 
INFO[0018] average 2.468008 s pkg=k.bench.reporter 
INFO[0018] total request size 810445500 pkg=k.bench.reporter 
INFO[0018] toatl response size 7200 pkg=k.bench.reporter 
INFO[0018] 200: 2400 pkg=k.bench.reporter 
INFO[0018] 0: 2100 pkg=k.bench.reporter 
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
INFO[0018] bench finished pkg=k.cmd.bench 
````