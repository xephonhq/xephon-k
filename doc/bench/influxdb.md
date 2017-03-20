# Benchmark Result: InfluxDB

version: 1.2.2

````bash
xkb --db i
````

- mem: 70 MB

````
INFO[0033] basic report finished via context pkg=k.b.reporter 
INFO[0033] total request 118 pkg=k.b.reporter 
INFO[0033] fastest 96604875 pkg=k.b.reporter 
INFO[0033] slowest 572493221 pkg=k.b.reporter 
INFO[0033] total request size 613600 pkg=k.b.reporter 
INFO[0033] toatl response size 0 pkg=k.b.reporter 
INFO[0033] 204: 118 pkg=k.b.reporter 
INFO[0033] loader finished pkg=k.b.loader 

Duration: 5
Worker number: 10
Batch size: 100
Timeout: 30
TargetDB: influxdb
````

````bash
xkb --db i -c 100
````

- mem: 90 MB

````
INFO[0015] total request 139 pkg=k.b.reporter 
INFO[0015] fastest 102784459 pkg=k.b.reporter 
INFO[0015] slowest 4299014783 pkg=k.b.reporter 
INFO[0015] total request size 722800 pkg=k.b.reporter 
INFO[0015] toatl response size 0 pkg=k.b.reporter 
INFO[0015] 204: 139 pkg=k.b.reporter 
INFO[0015] loader finished pkg=k.b.loader 

Duration: 5
Worker number: 100
Batch size: 100
Timeout: 30
TargetDB: influxdb
````

````bash
xkb --db i -c 1000
````

- mem: 450 MB

````
INFO[0017] total request 1026 pkg=k.b.reporter 
INFO[0017] fastest 26741 pkg=k.b.reporter 
INFO[0017] slowest 5936632875 pkg=k.b.reporter 
INFO[0017] total request size 5335200 pkg=k.b.reporter 
INFO[0017] toatl response size 0 pkg=k.b.reporter 
INFO[0017] 204: 131 pkg=k.b.reporter 
INFO[0017] 0: 895 pkg=k.b.reporter 
INFO[0017] loader finished pkg=k.b.loader 

Duration: 5
Worker number: 1000
Batch size: 100
Timeout: 30
TargetDB: influxdb
````

````bash
xkb --db i -c 100 -d 60
````

- mem: 121 MB
- [ ] TODO: virtual size is the disk size? 790 MB?

````
INFO[0068] total request 1332 pkg=k.b.reporter 
INFO[0068] fastest 87702586 pkg=k.b.reporter 
INFO[0068] slowest 5231069450 pkg=k.b.reporter 
INFO[0068] total request size 6926400 pkg=k.b.reporter 
INFO[0068] toatl response size 0 pkg=k.b.reporter 
INFO[0068] 204: 1332 pkg=k.b.reporter 
INFO[0068] loader finished pkg=k.b.loader 

Duration: 60
Worker number: 100
Batch size: 100
Timeout: 30
TargetDB: influxdb
````