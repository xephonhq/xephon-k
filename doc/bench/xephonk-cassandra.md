# Benchmark Result: Xephon-K Cassandra

````bash
xkb
````

````
INFO[0022] basic report finished via context pkg=k.b.reporter 
INFO[0022] total request 7931 pkg=k.b.reporter 
INFO[0022] fastest 1071335 pkg=k.b.reporter 
INFO[0022] slowest 67254788 pkg=k.b.reporter 
INFO[0022] total request size 15584415 pkg=k.b.reporter 
INFO[0022] toatl response size 245861 pkg=k.b.reporter 
INFO[0022] 200: 7931 pkg=k.b.reporter 
INFO[0022] loader finished pkg=k.b.loader 

Duration: 5
Worker number: 10
Batch size: 100
Timeout: 30
TargetDB: xephonk
````

````bash
xkb -c 100
````

````
INFO[0007] basic report finished via context pkg=k.b.reporter 
INFO[0007] total request 11336 pkg=k.b.reporter 
INFO[0007] fastest 1802954 pkg=k.b.reporter 
INFO[0007] slowest 171326825 pkg=k.b.reporter 
INFO[0007] total request size 22275240 pkg=k.b.reporter 
INFO[0007] toatl response size 351416 pkg=k.b.reporter 
INFO[0007] 200: 11336 pkg=k.b.reporter 
INFO[0007] loader finished pkg=k.b.loader 

Duration: 5
Worker number: 100
Batch size: 100
Timeout: 30
TargetDB: xephonk
````

````bash
xkb -c 1000
````

````
INFO[0007] basic report finished via context pkg=k.b.reporter 
INFO[0007] worker finished pkg=k.b.loader 
INFO[0007] total request 18498 pkg=k.b.reporter 
INFO[0007] fastest 24885 pkg=k.b.reporter 
INFO[0007] slowest 1265079140 pkg=k.b.reporter 
INFO[0007] total request size 36348570 pkg=k.b.reporter 
INFO[0007] toatl response size 452290 pkg=k.b.reporter 
INFO[0007] 200: 14590 pkg=k.b.reporter 
INFO[0007] 0: 3908 pkg=k.b.reporter 
INFO[0007] loader finished pkg=k.b.loader 

Duration: 5
Worker number: 1000
Batch size: 100
Timeout: 30
TargetDB: xephonk
````

when xephon k runs in docker (a lot better)

````
INFO[0008] basic report finished via context pkg=k.b.reporter 
INFO[0008] total request 10297 pkg=k.b.reporter 
INFO[0008] fastest 40014 pkg=k.b.reporter 
INFO[0008] slowest 2513650429 pkg=k.b.reporter 
INFO[0008] total request size 20233605 pkg=k.b.reporter 
INFO[0008] toatl response size 306962 pkg=k.b.reporter 
INFO[0008] 0: 395 pkg=k.b.reporter 
INFO[0008] 200: 9902 pkg=k.b.reporter 
INFO[0008] loader finished pkg=k.b.loader 

Duration: 5
Worker number: 1000
Batch size: 100
Timeout: 30
TargetDB: xephonk
````

````bash
xkb -c 5000
````

````
INFO[0006] basic report finished via context pkg=k.b.reporter 
WARN[0032] Post http://localhost:23333/write: net/http: request canceled (Client.Timeout exceeded while awaiting headers) pkg=k.b.loader 
INFO[0032] worker finished pkg=k.b.loader 
WARN[0032] Post http://localhost:23333/write: net/http: request canceled (Client.Timeout exceeded while awaiting headers) pkg=k.b.loader 
INFO[0032] worker finished pkg=k.b.loader 
WARN[0032] Post http://localhost:23333/write: net/http: request canceled (Client.Timeout exceeded while awaiting headers) pkg=k.b.loader 
INFO[0032] worker finished pkg=k.b.loader 
WARN[0032] Post http://localhost:23333/write: net/http: request canceled (Client.Timeout exceeded while awaiting headers) pkg=k.b.loader 
INFO[0032] worker finished pkg=k.b.loader 
WARN[0032] Post http://localhost:23333/write: net/http: request canceled (Client.Timeout exceeded while awaiting headers) pkg=k.b.loader 
INFO[0032] worker finished pkg=k.b.loader 
WARN[0032] Post http://localhost:23333/write: net/http: request canceled (Client.Timeout exceeded while awaiting headers) pkg=k.b.loader 
INFO[0032] worker finished pkg=k.b.loader 
INFO[0032] total request 42345 pkg=k.b.reporter 
INFO[0032] fastest 21755 pkg=k.b.reporter 
INFO[0032] slowest 3064984043 pkg=k.b.reporter 
INFO[0032] total request size 83207925 pkg=k.b.reporter 
INFO[0032] toatl response size 269793 pkg=k.b.reporter 
INFO[0032] 200: 8703 pkg=k.b.reporter 
INFO[0032] 0: 33642 pkg=k.b.reporter 
INFO[0032] loader finished pkg=k.b.loader 

Duration: 5
Worker number: 5000
Batch size: 100
Timeout: 30
TargetDB: xephonk
````

````bash
xkb -c 100 -d 60
````

- it seems Cassandra is using memory, 8GB from docker stats

````
INFO[0062] basic report finished via context pkg=k.b.reporter 
INFO[0062] total request 149627 pkg=k.b.reporter 
INFO[0062] fastest 1264924 pkg=k.b.reporter 
INFO[0062] slowest 376510175 pkg=k.b.reporter 
INFO[0062] total request size 294017055 pkg=k.b.reporter 
INFO[0062] toatl response size 4638437 pkg=k.b.reporter 
INFO[0062] 200: 149627 pkg=k.b.reporter 
INFO[0062] loader finished pkg=k.b.loader 

Duration: 60
Worker number: 100
Batch size: 100
Timeout: 30
TargetDB: xephonk
````

