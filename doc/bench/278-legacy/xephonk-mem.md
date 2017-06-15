# Benchmark Result: Xephon-K Memory

````bash
xkb
````

````
INFO[0007] basic report finished via context pkg=k.b.reporter 
INFO[0007] total request 12327 pkg=k.b.reporter 
INFO[0007] fastest 280812 pkg=k.b.reporter 
INFO[0007] slowest 20812536 pkg=k.b.reporter 
INFO[0007] total request size 24222555 pkg=k.b.reporter 
INFO[0007] toatl response size 382137 pkg=k.b.reporter 
INFO[0007] 200: 12327 pkg=k.b.reporter 
INFO[0007] loader finished pkg=k.b.loader 

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
INFO[0008] basic report finished via context pkg=k.b.reporter 
INFO[0008] total request 21099 pkg=k.b.reporter 
INFO[0008] fastest 261566 pkg=k.b.reporter 
INFO[0008] slowest 167451326 pkg=k.b.reporter 
INFO[0008] total request size 41459535 pkg=k.b.reporter 
INFO[0008] toatl response size 654069 pkg=k.b.reporter 
INFO[0008] 200: 21099 pkg=k.b.reporter 
INFO[0008] loader finished pkg=k.b.loader 

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
INFO[0007] total request 37385 pkg=k.b.reporter 
INFO[0007] fastest 22723 pkg=k.b.reporter 
INFO[0007] slowest 4386436638 pkg=k.b.reporter 
INFO[0007] total request size 73461525 pkg=k.b.reporter 
INFO[0007] toatl response size 985521 pkg=k.b.reporter 
INFO[0007] 200: 31791 pkg=k.b.reporter 
INFO[0007] 0: 5594 pkg=k.b.reporter 
INFO[0007] loader finished pkg=k.b.loader 

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
INFO[0008] basic report finished via context pkg=k.b.reporter 
WARN[0035] Post http://localhost:23333/write: net/http: request canceled (Client.Timeout exceeded while awaiting headers) pkg=k.b.loader 
INFO[0035] worker finished pkg=k.b.loader 
INFO[0035] total request 53679 pkg=k.b.reporter 
INFO[0035] fastest 22528 pkg=k.b.reporter 
INFO[0035] slowest 2231910891 pkg=k.b.reporter 
INFO[0035] total request size 105479235 pkg=k.b.reporter 
INFO[0035] toatl response size 380649 pkg=k.b.reporter 
INFO[0035] 200: 12279 pkg=k.b.reporter 
INFO[0035] 0: 41400 pkg=k.b.reporter 
INFO[0035] loader finished pkg=k.b.loader 

Duration: 5
Worker number: 5000
Batch size: 100
Timeout: 30
TargetDB: xephonk
````

````
2017/03/20 05:32:13 http: Accept error: accept tcp [::]:23333: accept4: too many open files; retrying in 1s
2017/03/20 05:32:14 http: Accept error: accept tcp [::]:23333: accept4: too many open files; retrying in 1s
````

````bash
xkb -c 100 -d 60
````

````
INFO[0062] basic report finished via context pkg=k.b.reporter 
INFO[0062] total request 122791 pkg=k.b.reporter 
INFO[0062] fastest 276946 pkg=k.b.reporter 
INFO[0062] slowest 485515668 pkg=k.b.reporter 
INFO[0062] total request size 241284315 pkg=k.b.reporter 
INFO[0062] toatl response size 3806521 pkg=k.b.reporter 
INFO[0062] 200: 122791 pkg=k.b.reporter 
INFO[0062] loader finished pkg=k.b.loader 

Duration: 60
Worker number: 100
Batch size: 100
Timeout: 30
TargetDB: xephonk
````