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