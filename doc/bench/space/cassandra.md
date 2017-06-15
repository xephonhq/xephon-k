# Cassandra

1,000,000 5.7M

````
⇒  xkb --limit points                  
log:
level: info
color: true
source: false
mode: local
loader:
target: xephonk
reporter: basic
limitBy: points
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
Do you want to proceed? [Y/N]y
INFO[0010] worker started pkg=k.bench.worker 
INFO[0010] worker started pkg=k.bench.worker 
INFO[0010] worker started pkg=k.bench.worker 
INFO[0010] worker started pkg=k.bench.worker 
INFO[0010] worker started pkg=k.bench.worker 
INFO[0010] worker started pkg=k.bench.worker 
INFO[0010] worker started pkg=k.bench.worker 
INFO[0010] worker started pkg=k.bench.worker 
INFO[0010] worker started pkg=k.bench.worker 
INFO[0010] worker started pkg=k.bench.worker 
INFO[0016] generator stopped after 1000000 points pkg=k.bench 
INFO[0016] close data channel pkg=k.bench 
INFO[0016] worker finished by input channel pkg=k.bench.worker 
INFO[0016] worker finished by input channel pkg=k.bench.worker 
INFO[0016] worker finished by input channel pkg=k.bench.worker 
INFO[0016] worker finished by input channel pkg=k.bench.worker 
INFO[0016] worker finished by input channel pkg=k.bench.worker 
INFO[0016] worker finished by input channel pkg=k.bench.worker 
INFO[0016] worker finished by input channel pkg=k.bench.worker 
INFO[0016] worker finished by input channel pkg=k.bench.worker 
INFO[0016] worker finished by input channel pkg=k.bench.worker 
INFO[0016] worker finished by input channel pkg=k.bench.worker 
INFO[0016] basic report finished by channel pkg=k.bench.reporter 
Total: 10000
0.0022296591022727273 	 .
0.0034435686221429653 	 ...........................................................................................................................
0.00586344693700787 	 ......................................
0.008388722092348271 	 ......................
0.011266087155844153 	 ....
0.013116105266666667 	 ...
0.01564649299224806 	 ..
0.018326106750000012 	 .
0.020744893818181814 	 
0.023335760400000002 	 
0.026754299000000002 	 
0.044450648599999996 	 
0.0462308615 	 
0.04903623783333333 	 
0.051487911500000004 	 
0.05425395715384614 	 
0.057249247705882356 	 
0.060334449111111105 	 
0.06393049499999999 	 
0.06815353833333333 	 
INFO[0017] run time 5.733265 s pkg=k.bench.reporter 
INFO[0017] total request 10000 pkg=k.bench.reporter 
INFO[0017] fastest 895463 pkg=k.bench.reporter 
INFO[0017] slowest 68439493 pkg=k.bench.reporter 
INFO[0017] total request size 18990000 pkg=k.bench.reporter 
INFO[0017] toatl response size 30000 pkg=k.bench.reporter 
INFO[0017] 200: 10000 pkg=k.bench.reporter 
INFO[0017] bench finished pkg=k.cmd.bench 

````

````
root@af2b4b0898ef:/var/lib/cassandra/data/xephon/metrics_int-061feae051f211e798e26d2c86545d91# du -sh * 
4.0K	backups
4.0K	mc-1-big-CompressionInfo.db
5.7M	mc-1-big-Data.db
4.0K	mc-1-big-Digest.crc32
4.0K	mc-1-big-Filter.db
12K	mc-1-big-Index.db
8.0K	mc-1-big-Statistics.db
4.0K	mc-1-big-Summary.db
4.0K	mc-1-big-TOC.txt
````