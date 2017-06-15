# Disk size

- 1,000,000
- 10,000,000
- 100,000,000
- 1000,000,000

1,000, 000, 19M?? wtf, even bigger than raw encoding? might due to difference between real disk and tmpfs? 
see https://github.com/xephonhq/xephon-k/issues/59, used limit by time instead of limit by points

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

1,000, 000   2.3M when using time: delta-rle, int: rle, should use delta, because we are generating constant, rle is too
good for this unrealistic workload

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
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0010] generator stopped after 1000000 points pkg=k.bench 
INFO[0010] close data channel pkg=k.bench 
INFO[0010] worker finished by input channel pkg=k.bench.worker 
INFO[0010] worker finished by input channel pkg=k.bench.worker 
INFO[0010] worker finished by input channel pkg=k.bench.worker 
INFO[0010] worker finished by input channel pkg=k.bench.worker 
INFO[0010] worker finished by input channel pkg=k.bench.worker 
INFO[0010] worker finished by input channel pkg=k.bench.worker 
INFO[0010] worker finished by input channel pkg=k.bench.worker 
INFO[0010] worker finished by input channel pkg=k.bench.worker 
INFO[0010] worker finished by input channel pkg=k.bench.worker 
INFO[0010] worker finished by input channel pkg=k.bench.worker 
INFO[0010] basic report finished by channel pkg=k.bench.reporter 
Total: 10000
0.0006102571154192971 	 ........................................................................................
0.0012749141321081784 	 .........................................................
0.0018576486859872598 	 ...............................
0.002437606095041328 	 ..............
0.003035432190909087 	 ....
0.003775578472222223 	 .
0.004373778457142859 	 
0.005180335 	 
0.005956733 	 
0.0067009719999999995 	 
0.00746055425 	 
0.008226141666666666 	 
0.0090028815 	 
0.009619168 	 
0.0112141125 	 
0.012202475 	 
0.016019723 	 
0.020154118 	 
0.02318547 	 
0.024463787 	 
INFO[0011] run time 1.432487 s pkg=k.bench.reporter 
INFO[0011] total request 10000 pkg=k.bench.reporter 
INFO[0011] fastest 232410 pkg=k.bench.reporter 
INFO[0011] slowest 24463787 pkg=k.bench.reporter 
INFO[0011] total request size 18990000 pkg=k.bench.reporter 
INFO[0011] toatl response size 30000 pkg=k.bench.reporter 
INFO[0011] 200: 10000 pkg=k.bench.reporter 
INFO[0011] null reporter has nothing to say pkg=k.bench.reporter 
INFO[0011] bench finished pkg=k.cmd.bench 
````

1,000,000 3.2M, time: delta rle, int: delta

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
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0002] generator stopped after 1000000 points pkg=k.bench 
INFO[0002] close data channel pkg=k.bench 
INFO[0002] worker finished by input channel pkg=k.bench.worker 
INFO[0002] worker finished by input channel pkg=k.bench.worker 
INFO[0002] worker finished by input channel pkg=k.bench.worker 
INFO[0002] worker finished by input channel pkg=k.bench.worker 
INFO[0002] worker finished by input channel pkg=k.bench.worker 
INFO[0002] worker finished by input channel pkg=k.bench.worker 
INFO[0002] worker finished by input channel pkg=k.bench.worker 
INFO[0002] worker finished by input channel pkg=k.bench.worker 
INFO[0002] worker finished by input channel pkg=k.bench.worker 
INFO[0002] worker finished by input channel pkg=k.bench.worker 
INFO[0002] basic report finished by channel pkg=k.bench.reporter 
Total: 10000
0.00045218181326116426 	 ............................................
0.0007082348347457625 	 ............................
0.0009681919509866333 	 ...............................
0.0012405022393617 	 ..........................
0.0014915103282149723 	 ....................
0.0017311016346153855 	 ...............
0.00196930296085409 	 ...........
0.00220873830049261 	 ........
0.0024708952166064987 	 .....
0.0027174648333333335 	 ...
0.003025671896226416 	 ..
0.0033968824200000004 	 .
0.0037254968666666664 	 
0.004061329739130433 	 
0.004394342461538461 	 
0.00470427875 	 
0.0050012042 	 
0.005654899499999999 	 
0.006169457 	 
0.0070461612000000005 	 
INFO[0003] run time 1.368823 s pkg=k.bench.reporter 
INFO[0003] total request 10000 pkg=k.bench.reporter 
INFO[0003] fastest 226438 pkg=k.bench.reporter 
INFO[0003] slowest 7236968 pkg=k.bench.reporter 
INFO[0003] total request size 18990000 pkg=k.bench.reporter 
INFO[0003] toatl response size 30000 pkg=k.bench.reporter 
INFO[0003] 200: 10000 pkg=k.bench.reporter 
INFO[0003] bench finished pkg=k.cmd.bench 
````

10,000,000 29M

````
⇒  xkb --limit points --target xephonk 
log:
  level: info
  color: true
  source: false
mode: local
loader:
  target: xephonk
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
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0011] generator stopped after 10000000 points pkg=k.bench 
INFO[0011] close data channel pkg=k.bench 
INFO[0011] worker finished by input channel pkg=k.bench.worker 
INFO[0011] worker finished by input channel pkg=k.bench.worker 
INFO[0011] worker finished by input channel pkg=k.bench.worker 
INFO[0011] worker finished by input channel pkg=k.bench.worker 
INFO[0011] worker finished by input channel pkg=k.bench.worker 
INFO[0011] worker finished by input channel pkg=k.bench.worker 
INFO[0011] worker finished by input channel pkg=k.bench.worker 
INFO[0011] worker finished by input channel pkg=k.bench.worker 
INFO[0011] worker finished by input channel pkg=k.bench.worker 
INFO[0011] worker finished by input channel pkg=k.bench.worker 
INFO[0011] basic report finished by channel pkg=k.bench.reporter 
Total: 10000
0.0020281875526315784 	 
0.003175934055555555 	 ...
0.004552226834090912 	 .................
0.006030054187468917 	 ........................................
0.007512146075441425 	 .....................................
0.008701753154525383 	 ...........................
0.009732645569049939 	 ....................
0.010847382959956723 	 ..................
0.01198236113880126 	 ............
0.013331892762183236 	 ..........
0.014807544319702604 	 .....
0.016060423599999988 	 ..
0.0173818688313253 	 .
0.018705331800000002 	 
0.02003336238461538 	 
0.0214789623 	 
0.02297275325 	 
0.0242951098 	 
0.027566877 	 
0.029418582 	 
INFO[0012] run time 10.341061 s pkg=k.bench.reporter 
INFO[0012] total request 10000 pkg=k.bench.reporter 
INFO[0012] fastest 1613087 pkg=k.bench.reporter 
INFO[0012] slowest 29418582 pkg=k.bench.reporter 
INFO[0012] total request size 180990000 pkg=k.bench.reporter 
INFO[0012] toatl response size 30000 pkg=k.bench.reporter 
INFO[0012] 200: 10000 pkg=k.bench.reporter 
INFO[0012] bench finished pkg=k.cmd.bench
````

100,000,000 287 M, 1.5G when using raw-binary

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
Do you want to proceed? [Y/N]y
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0001] worker started pkg=k.bench.worker 
INFO[0072] generator stopped after 100000000 points pkg=k.bench 
INFO[0072] close data channel pkg=k.bench 
INFO[0072] worker finished by input channel pkg=k.bench.worker 
INFO[0072] worker finished by input channel pkg=k.bench.worker 
INFO[0072] worker finished by input channel pkg=k.bench.worker 
INFO[0072] worker finished by input channel pkg=k.bench.worker 
INFO[0072] worker finished by input channel pkg=k.bench.worker 
INFO[0072] worker finished by input channel pkg=k.bench.worker 
INFO[0072] worker finished by input channel pkg=k.bench.worker 
INFO[0072] worker finished by input channel pkg=k.bench.worker 
INFO[0072] worker finished by input channel pkg=k.bench.worker 
INFO[0072] worker finished by input channel pkg=k.bench.worker 
INFO[0072] basic report finished by channel pkg=k.bench.reporter 
Total: 10000
0.026522103499999995 	 
0.032191093134751755 	 ..
0.03801784936561268 	 ..........
0.04339379608096082 	 ......................
0.04989259830796464 	 .............................................
0.0559716648253574 	 ................................
0.060730263785116916 	 .......................
0.06509780295828989 	 ...................
0.07002791309801762 	 ..................
0.07535481996280075 	 .........
0.07988493372222216 	 ......
0.08501465965533982 	 ....
0.09101157888888886 	 ...
0.09750006949999998 	 .
0.1029330303125 	 
0.1105206458076923 	 
0.11855614436363636 	 
0.1263750395 	 
0.140397975 	 
0.14645729049999998 	 
INFO[0073] run time 71.274927 s pkg=k.bench.reporter 
INFO[0073] total request 10000 pkg=k.bench.reporter 
INFO[0073] fastest 21844465 pkg=k.bench.reporter 
INFO[0073] slowest 147451601 pkg=k.bench.reporter 
INFO[0073] total request size 1800990000 pkg=k.bench.reporter 
INFO[0073] toatl response size 30000 pkg=k.bench.reporter 
INFO[0073] 200: 10000 pkg=k.bench.reporter 
INFO[0073] bench finished pkg=k.cmd.bench 
````

287 M after compressed is 

- zip result in 582.8 kb ....?
- lzma 181.7 kb
- 7z 181.8 kb

1.5G using raw-big. gzip is 283.7 MB

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
Do you want to proceed? [Y/N]y
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0008] worker started pkg=k.bench.worker 
INFO[0079] generator stopped after 100000000 points pkg=k.bench 
INFO[0079] close data channel pkg=k.bench 
INFO[0079] worker finished by input channel pkg=k.bench.worker 
INFO[0079] worker finished by input channel pkg=k.bench.worker 
INFO[0079] worker finished by input channel pkg=k.bench.worker 
INFO[0079] worker finished by input channel pkg=k.bench.worker 
INFO[0079] worker finished by input channel pkg=k.bench.worker 
INFO[0079] worker finished by input channel pkg=k.bench.worker 
INFO[0079] worker finished by input channel pkg=k.bench.worker 
INFO[0079] worker finished by input channel pkg=k.bench.worker 
INFO[0079] worker finished by input channel pkg=k.bench.worker 
INFO[0079] worker finished by input channel pkg=k.bench.worker 
INFO[0079] basic report finished by channel pkg=k.bench.reporter 
Total: 10000
0.029133290621621625 	 
0.03402392273988434 	 ...
0.0401161710662791 	 .................
0.045676627330795236 	 .......................
0.05067933385255645 	 .................................
0.05561866520184795 	 ............................
0.06005642726605504 	 ..........................
0.06435995484747493 	 ...................
0.06869983804575164 	 ...............
0.07390380865286628 	 ............
0.0792305336590909 	 .......
0.08519451690573766 	 ....
0.09164639187704916 	 ..
0.09776819165789473 	 .
0.10467487340909094 	 
0.11232691065625001 	 
0.12048861888461539 	 
0.12880395666666666 	 
0.134796374 	 
0.14027290633333334 	 
INFO[0080] run time 71.095215 s pkg=k.bench.reporter 
INFO[0080] total request 10000 pkg=k.bench.reporter 
INFO[0080] fastest 26767856 pkg=k.bench.reporter 
INFO[0080] slowest 143867098 pkg=k.bench.reporter 
INFO[0080] total request size 1800990000 pkg=k.bench.reporter 
INFO[0080] toatl response size 30000 pkg=k.bench.reporter 
INFO[0080] 200: 10000 pkg=k.bench.reporter 
INFO[0080] bench finished pkg=k.cmd.bench 
````