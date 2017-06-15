# InfluxDB

1,000,000 6.1M WAL, data folder is empty

````
⇒  xkb --limit points --target influxdb
log:
  level: info
  color: true
  source: false
mode: local
loader:
  target: influxdb
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
INFO[0011] worker started pkg=k.bench.worker 
INFO[0011] worker started pkg=k.bench.worker 
INFO[0011] worker started pkg=k.bench.worker 
INFO[0011] worker started pkg=k.bench.worker 
INFO[0011] worker started pkg=k.bench.worker 
INFO[0011] worker started pkg=k.bench.worker 
INFO[0011] worker started pkg=k.bench.worker 
INFO[0011] worker started pkg=k.bench.worker 
INFO[0011] worker started pkg=k.bench.worker 
INFO[0011] worker started pkg=k.bench.worker 
INFO[0029] generator stopped after 1000000 points pkg=k.bench 
INFO[0029] close data channel pkg=k.bench 
INFO[0029] worker finished by input channel pkg=k.bench.worker 
INFO[0029] worker finished by input channel pkg=k.bench.worker 
INFO[0029] worker finished by input channel pkg=k.bench.worker 
INFO[0029] worker finished by input channel pkg=k.bench.worker 
INFO[0029] worker finished by input channel pkg=k.bench.worker 
INFO[0029] worker finished by input channel pkg=k.bench.worker 
INFO[0029] worker finished by input channel pkg=k.bench.worker 
INFO[0029] worker finished by input channel pkg=k.bench.worker 
INFO[0029] worker finished by input channel pkg=k.bench.worker 
INFO[0029] worker finished by input channel pkg=k.bench.worker 
INFO[0029] basic report finished by channel pkg=k.bench.reporter 
Total: 10000
0.0058959335000000005 	 
0.009159969666666668 	 
0.011229679 	 
0.013282145 	 
0.015270097005449566 	 ..............
0.016677509277120688 	 ...........................................................
0.018215746064766215 	 ...............................................................................
0.019876091 	 
0.02079301199631902 	 ................................
0.023975321691131493 	 ......
0.02617547595833334 	 .
0.027930318166666666 	 
0.030007455722222225 	 
0.0320365560506329 	 .
0.033529532788235276 	 .
0.035604549571428584 	 
0.037761762250000004 	 
0.03995495111764706 	 
0.04229246 	 
0.055023617 	 
INFO[0030] run time 18.733643 s pkg=k.bench.reporter 
INFO[0030] total request 10000 pkg=k.bench.reporter 
INFO[0030] fastest 5892809 pkg=k.bench.reporter 
INFO[0030] slowest 55023617 pkg=k.bench.reporter 
INFO[0030] total request size 47000000 pkg=k.bench.reporter 
INFO[0030] toatl response size 0 pkg=k.bench.reporter 
INFO[0030] 204: 10000 pkg=k.bench.reporter 
INFO[0030] bench finished pkg=k.cmd.bench 
````

10,000,000  1.8 M 


````
root@329e05187d49:/var/lib/influxdb/data/xb/autogen/2# du -sh *
1.5M	000000004-000000003.tsm
332K	000000005-000000001.tsm
````

```` 
⇒  xkb --limit points --target influxdb
log:
  level: info
  color: true
  source: false
mode: local
loader:
  target: influxdb
  reporter: basic
  limitBy: points
  points: 10000000
  series: 100
  time: 10
  workerNum: 10
  workerTimeout: 30
generator:
  timeInterval: 1
  timeNoise: false
  pointsPerSeries: 1000
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
INFO[0003] worker started pkg=k.bench.worker 
INFO[0003] worker started pkg=k.bench.worker 
INFO[0003] worker started pkg=k.bench.worker 
INFO[0003] worker started pkg=k.bench.worker 
INFO[0003] worker started pkg=k.bench.worker 
INFO[0003] worker started pkg=k.bench.worker 
INFO[0003] worker started pkg=k.bench.worker 
INFO[0003] worker started pkg=k.bench.worker 
INFO[0003] worker started pkg=k.bench.worker 
INFO[0003] worker started pkg=k.bench.worker 
INFO[0023] generator stopped after 10000000 points pkg=k.bench 
INFO[0023] close data channel pkg=k.bench 
INFO[0023] worker finished by input channel pkg=k.bench.worker 
INFO[0023] worker finished by input channel pkg=k.bench.worker 
INFO[0023] worker finished by input channel pkg=k.bench.worker 
INFO[0023] worker finished by input channel pkg=k.bench.worker 
INFO[0023] worker finished by input channel pkg=k.bench.worker 
INFO[0023] worker finished by input channel pkg=k.bench.worker 
INFO[0023] worker finished by input channel pkg=k.bench.worker 
INFO[0023] worker finished by input channel pkg=k.bench.worker 
INFO[0023] worker finished by input channel pkg=k.bench.worker 
INFO[0023] worker finished by input channel pkg=k.bench.worker 
INFO[0023] basic report finished by channel pkg=k.bench.reporter 
Total: 10000
0.006332456714285712 	 
0.011334941534883722 	 
0.0152038685 	 
0.01617709979157388 	 ..................................................................................................
0.019085289336168654 	 ......................................................
0.021819765044131394 	 .....................
0.025140076313765188 	 .........
0.028374251141078836 	 ....
0.031259074407035146 	 ...
0.03442091589719627 	 ..
0.03699015820895522 	 .
0.0399347663125 	 
0.04243660850000001 	 
0.045809186999999994 	 
0.050121871307692314 	 
0.05558600208333333 	 
0.061915671000000005 	 
0.06706112425 	 
0.076331226 	 
0.0790071195 	 
INFO[0024] run time 19.655630 s pkg=k.bench.reporter 
INFO[0024] total request 10000 pkg=k.bench.reporter 
INFO[0024] fastest 4459598 pkg=k.bench.reporter 
INFO[0024] slowest 79678882 pkg=k.bench.reporter 
INFO[0024] total request size 470000000 pkg=k.bench.reporter 
INFO[0024] toatl response size 0 pkg=k.bench.reporter 
INFO[0024] 204: 10000 pkg=k.bench.reporter 
INFO[0024] bench finished pkg=k.cmd.bench 
````