# Xephon-K Cassandra

10s 10c 318, 0.29

````
Total: 318
0.0797981915 	 ..
0.11915894063157895 	 ...........
0.15643768999999996 	 ...................
0.19272900803333334 	 ..................
0.2348698889714286 	 ......................
0.269988428106383 	 .............................
0.30895350044318204 	 .......................................................
0.34247679645454543 	 ......
0.3769694775 	 ......
0.41616924066666666 	 ...
0.453921547 	 .
0.5083765958333333 	 ...
0.554463771 	 .....
0.6257069283333333 	 .
0.67679772925 	 ..
0.706570204 	 
0.7356947545 	 .
0.7681020169999999 	 ..
0.8083812667499999 	 ..
0.848361017 	 .
INFO[0012] total time 10.554138 s pkg=k.bench.reporter 
INFO[0012] total request 318 pkg=k.bench.reporter 
INFO[0012] fastest 0.070755 s pkg=k.bench.reporter 
INFO[0012] slowest 0.858360 s pkg=k.bench.reporter 
INFO[0012] average 0.298613 s pkg=k.bench.reporter 
INFO[0012] total request size 57271482 pkg=k.bench.reporter 
INFO[0012] toatl response size 954 pkg=k.bench.reporter 
INFO[0012] 200: 318 pkg=k.bench.reporter 
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

10s 100c 627, 1.9s

````
INFO[0014] basic report finished by channel pkg=k.bench.reporter 
Total: 627
0.26272336955555553 	 ........
0.5567736055737706 	 ...................
0.7894176150882355 	 .....................
1.0373669888529409 	 .....................
1.3154160888372095 	 .............
1.57516289704 	 .......
1.8121432002258064 	 .........
2.0164037380769244 	 ........
2.188213587800001 	 ...........
2.4304624050285715 	 ...........
2.623882539045455 	 ..............
2.838263624354838 	 .........
3.0132867326190476 	 ......
3.1948451251562493 	 ..........
3.3762267595000006 	 ..........
3.5541917688888884 	 .....
3.860322758538462 	 ....
4.102779997857143 	 ..
4.273254579 	 
4.4537129502857145 	 ..
INFO[0015] total time 13.736805 s pkg=k.bench.reporter 
INFO[0015] total request 627 pkg=k.bench.reporter 
INFO[0015] fastest 0.118767 s pkg=k.bench.reporter 
INFO[0015] slowest 4.591250 s pkg=k.bench.reporter 
INFO[0015] average 1.903585 s pkg=k.bench.reporter 
INFO[0015] total request size 112922073 pkg=k.bench.reporter 
INFO[0015] toatl response size 1881 pkg=k.bench.reporter 
INFO[0015] 200: 627 pkg=k.bench.reporter 
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

10s (not really...), 1000c, 

````
INFO[0041] basic report finished by channel pkg=k.bench.reporter 
Total: 2189
0.0383672968368421 	 .................
0.813273105850746 	 ......
2.431096911587302 	 ...........
3.6952212878227852 	 .......
4.731130248885136 	 .............
6.060364972951389 	 .............
7.390589468971014 	 ............
8.685756356074076 	 .......
9.931236914206895 	 .......
11.084379300064516 	 .....
12.352632542195403 	 .......
13.843497839441442 	 ..........
15.654074393137098 	 ...........
17.129080950986847 	 ......
19.263259802813177 	 ................
20.747514336697485 	 ..........
22.00281510594703 	 .............
23.031888942984732 	 ...........
24.31982404686076 	 .......
25.801402818142854 	 
INFO[0042] total time 40.593780 s pkg=k.bench.reporter 
INFO[0042] total request 2189 pkg=k.bench.reporter 
INFO[0042] fastest 0.000063 s pkg=k.bench.reporter 
INFO[0042] slowest 26.222244 s pkg=k.bench.reporter 
INFO[0042] average 11.776238 s pkg=k.bench.reporter 
INFO[0042] total request size 394236711 pkg=k.bench.reporter 
INFO[0042] toatl response size 6036 pkg=k.bench.reporter 
INFO[0042] 200: 2012 pkg=k.bench.reporter 
INFO[0042] 0: 177 pkg=k.bench.reporter 
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
INFO[0042] bench finished pkg=k.cmd.bench 
````