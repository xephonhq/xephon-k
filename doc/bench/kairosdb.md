# Benchmark Result: KairosDB

````bash
xkb --db k
````

````
INFO[0009] basic report finished via context pkg=k.b.reporter 
INFO[0009] total request 15561 pkg=k.b.reporter 
INFO[0009] fastest 189063 pkg=k.b.reporter 
INFO[0009] slowest 626411401 pkg=k.b.reporter 
INFO[0009] total request size 30639609 pkg=k.b.reporter 
INFO[0009] toatl response size 0 pkg=k.b.reporter 
INFO[0009] 204: 15561 pkg=k.b.reporter 
INFO[0009] loader finished pkg=k.b.loader 

Duration: 5
Worker number: 10
Batch size: 100
Timeout: 30
TargetDB: kairosdb
````

````
xephonhq-kairosdb-cp | 05:46:08.919 [WriteBuffer-data_points-2] ERROR [HThriftClient.java:132] - Could not flush transport (to be expected if the pool is shutting down) in close for client: CassandraClient<kairosdbcassandra:9160-14>
xephonhq-kairosdb-cp | org.apache.thrift.transport.TTransportException: java.net.SocketException: Broken pipe (Write failed)
xephonhq-kairosdb-cp | 	at org.apache.thrift.transport.TIOStreamTransport.write(TIOStreamTransport.java:147)
xephonhq-kairosdb-cp | 	at org.apache.thrift.transport.TFramedTransport.flush(TFramedTransport.java:156)
xephonhq-kairosdb-cp | 	at me.prettyprint.cassandra.connection.client.HThriftClient.close(HThriftClient.java:130)
xephonhq-kairosdb-cp | 	at me.prettyprint.cassandra.connection.client.HThriftClient.close(HThriftClient.java:40)
xephonhq-kairosdb-cp | 	at me.prettyprint.cassandra.connection.HConnectionManager.closeClient(HConnectionManager.java:320)
xephonhq-kairosdb-cp | 	at me.prettyprint.cassandra.connection.HConnectionManager.operateWithFailover(HConnectionManager.java:268)
xephonhq-kairosdb-cp | 	at me.prettyprint.cassandra.model.ExecutingKeyspace.doExecuteOperation(ExecutingKeyspace.java:113)
xephonhq-kairosdb-cp | 	at me.prettyprint.cassandra.model.MutatorImpl.execute(MutatorImpl.java:243)
xephonhq-kairosdb-cp | 	at org.kairosdb.datastore.cassandra.WriteBuffer$WriteDataJob.run(WriteBuffer.java:308)
xephonhq-kairosdb-cp | 	at java.util.concurrent.Executors$RunnableAdapter.call(Executors.java:511)
xephonhq-kairosdb-cp | 	at java.util.concurrent.FutureTask.run(FutureTask.java:266)
xephonhq-kairosdb-cp | 	at java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1142)
xephonhq-kairosdb-cp | 	at java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:617)
xephonhq-kairosdb-cp | 	at java.lang.Thread.run(Thread.java:745)
xephonhq-kairosdb-cp | Caused by: java.net.SocketException: Broken pipe (Write failed)
xephonhq-kairosdb-cp | 	at java.net.SocketOutputStream.socketWrite0(Native Method)
xephonhq-kairosdb-cp | 	at java.net.SocketOutputStream.socketWrite(SocketOutputStream.java:109)
xephonhq-kairosdb-cp | 	at java.net.SocketOutputStream.write(SocketOutputStream.java:153)
xephonhq-kairosdb-cp | 	at org.apache.thrift.transport.TIOStreamTransport.write(TIOStreamTransport.java:145)
xephonhq-kairosdb-cp | 	... 13 common frames omitted
xephonhq-kairosdb-cp | 05:46:08.919 [WriteBuffer-data_points-2] ERROR [HConnectionManager.java:418] - MARK HOST AS DOWN TRIGGERED for host kairosdbcassandra(172.18.0.2):9160
xephonhq-kairosdb-cp | 05:46:08.919 [WriteBuffer-data_points-2] ERROR [HConnectionManager.java:422] - Pool state on shutdown: <ConcurrentCassandraClientPoolByHost>:{kairosdbcassandra(172.18.0.2):9160}; IsActive?: true; Active: 1; Blocked: 0; Idle: 15; NumBeforeExhausted: 49
xephonhq-kairosdb-cp | 05:46:08.920 [WriteBuffer-data_points-2] INFO  [ConcurrentHClientPool.java:189] - Shutdown triggered on <ConcurrentCassandraClientPoolByHost>:{kairosdbcassandra(172.18.0.2):9160}
xephonhq-kairosdb-cp | 05:46:08.920 [WriteBuffer-data_points-2] INFO  [ConcurrentHClientPool.java:197] - Shutdown complete on <ConcurrentCassandraClientPoolByHost>:{kairosdbcassandra(172.18.0.2):9160}
xephonhq-kairosdb-cp | 05:46:08.920 [WriteBuffer-data_points-2] INFO  [CassandraHostRetryService.java:68] - Host detected as down was added to retry queue: kairosdbcassandra(172.18.0.2):9160
xephonhq-kairosdb-cp | 05:46:08.921 [WriteBuffer-data_points-2] WARN  [HConnectionManager.java:302] - Could not fullfill request on this host CassandraClient<kairosdbcassandra:9160-14>
xephonhq-kairosdb-cp | 05:46:08.921 [WriteBuffer-data_points-2] WARN  [HConnectionManager.java:303] - Exception: 
xephonhq-kairosdb-cp | me.prettyprint.hector.api.exceptions.HectorTransportException: org.apache.thrift.transport.TTransportException: java.net.SocketException: Connection reset
xephonhq-kairosdb-cp | 	at me.prettyprint.cassandra.service.ExceptionsTranslatorImpl.translate(ExceptionsTranslatorImpl.java:39)
xephonhq-kairosdb-cp | 	at me.prettyprint.cassandra.connection.HConnectionManager.operateWithFailover(HConnectionManager.java:260)
xephonhq-kairosdb-cp | 	at me.prettyprint.cassandra.model.ExecutingKeyspace.doExecuteOperation(ExecutingKeyspace.java:113)
xephonhq-kairosdb-cp | 	at me.prettyprint.cassandra.model.MutatorImpl.execute(MutatorImpl.java:243)
xephonhq-kairosdb-cp | 	at org.kairosdb.datastore.cassandra.WriteBuffer$WriteDataJob.run(WriteBuffer.java:308)
xephonhq-kairosdb-cp | 	at java.util.concurrent.Executors$RunnableAdapter.call(Executors.java:511)
xephonhq-kairosdb-cp | 	at java.util.concurrent.FutureTask.run(FutureTask.java:266)
xephonhq-kairosdb-cp | 	at java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1142)
xephonhq-kairosdb-cp | 	at java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:617)
xephonhq-kairosdb-cp | 	at java.lang.Thread.run(Thread.java:745)
xephonhq-kairosdb-cp | Caused by: org.apache.thrift.transport.TTransportException: java.net.SocketException: Connection reset
xephonhq-kairosdb-cp | 	at org.apache.thrift.transport.TIOStreamTransport.write(TIOStreamTransport.java:147)
xephonhq-kairosdb-cp | 	at org.apache.thrift.transport.TFramedTransport.flush(TFramedTransport.java:157)
xephonhq-kairosdb-cp | 	at org.apache.thrift.TServiceClient.sendBase(TServiceClient.java:65)
xephonhq-kairosdb-cp | 	at org.apache.cassandra.thrift.Cassandra$Client.send_batch_mutate(Cassandra.java:958)
xephonhq-kairosdb-cp | 	at org.apache.cassandra.thrift.Cassandra$Client.batch_mutate(Cassandra.java:949)
xephonhq-kairosdb-cp | 	at me.prettyprint.cassandra.model.MutatorImpl$3.execute(MutatorImpl.java:246)
xephonhq-kairosdb-cp | 	at me.prettyprint.cassandra.model.MutatorImpl$3.execute(MutatorImpl.java:243)
xephonhq-kairosdb-cp | 	at me.prettyprint.cassandra.service.Operation.executeAndSetResult(Operation.java:104)
xephonhq-kairosdb-cp | 	at me.prettyprint.cassandra.connection.HConnectionManager.operateWithFailover(HConnectionManager.java:253)
xephonhq-kairosdb-cp | 	... 8 common frames omitted
xephonhq-kairosdb-cp | Caused by: java.net.SocketException: Connection reset
xephonhq-kairosdb-cp | 	at java.net.SocketOutputStream.socketWrite(SocketOutputStream.java:113)
xephonhq-kairosdb-cp | 	at java.net.SocketOutputStream.write(SocketOutputStream.java:153)
xephonhq-kairosdb-cp | 	at org.apache.thrift.transport.TIOStreamTransport.write(TIOStreamTransport.java:145)
xephonhq-kairosdb-cp | 	... 16 common frames omitted
xephonhq-kairosdb-cp | 05:46:08.921 [WriteBuffer-data_points-2] INFO  [HConnectionManager.java:404] - Client CassandraClient<kairosdbcassandra:9160-14> released to inactive or dead pool. Closing.
xephonhq-kairosdb-cp | 05:46:08.921 [WriteBuffer-data_points-2] ERROR [WriteBuffer.java:315] - Error sending data to Cassandra (data_points)
xephonhq-kairosdb-cp | me.prettyprint.hector.api.exceptions.HectorException: All host pools marked down. Retry burden pushed out to client.
xephonhq-kairosdb-cp | 	at me.prettyprint.cassandra.connection.HConnectionManager.getClientFromLBPolicy(HConnectionManager.java:390)
xephonhq-kairosdb-cp | 	at me.prettyprint.cassandra.connection.HConnectionManager.operateWithFailover(HConnectionManager.java:244)
xephonhq-kairosdb-cp | 	at me.prettyprint.cassandra.model.ExecutingKeyspace.doExecuteOperation(ExecutingKeyspace.java:113)
xephonhq-kairosdb-cp | 	at me.prettyprint.cassandra.model.MutatorImpl.execute(MutatorImpl.java:243)
xephonhq-kairosdb-cp | 	at org.kairosdb.datastore.cassandra.WriteBuffer$WriteDataJob.run(WriteBuffer.java:308)
xephonhq-kairosdb-cp | 	at java.util.concurrent.Executors$RunnableAdapter.call(Executors.java:511)
xephonhq-kairosdb-cp | 	at java.util.concurrent.FutureTask.run(FutureTask.java:266)
xephonhq-kairosdb-cp | 	at java.util.concurrent.ThreadPoolExecutor.runWorker(ThreadPoolExecutor.java:1142)
xephonhq-kairosdb-cp | 	at java.util.concurrent.ThreadPoolExecutor$Worker.run(ThreadPoolExecutor.java:617)
xephonhq-kairosdb-cp | 	at java.lang.Thread.run(Thread.java:745)
xephonhq-kairosdb-cp | 05:46:08.921 [WriteBuffer-data_points-2] ERROR [WriteBuffer.java:319] - Reducing write buffer size to 375000.  You need to increase your cassandra capacity or change the kairosdb.datastore.cassandra.write_buffer_max_size property.
````

````bash
xkb --db k -c 100
````

````
INFO[0007] basic report finished via context pkg=k.b.reporter 
INFO[0007] total request 26154 pkg=k.b.reporter 
INFO[0007] fastest 250782 pkg=k.b.reporter 
INFO[0007] slowest 1067137018 pkg=k.b.reporter 
INFO[0007] total request size 51497226 pkg=k.b.reporter 
INFO[0007] toatl response size 0 pkg=k.b.reporter 
INFO[0007] 204: 26154 pkg=k.b.reporter 
INFO[0007] loader finished pkg=k.b.loader 

Duration: 5
Worker number: 100
Batch size: 100
Timeout: 30
TargetDB: kairosdb
````

````bash
xkb --db k -c 1000
````

````
INFO[0010] basic report finished via context pkg=k.b.reporter 
INFO[0010] total request 26945 pkg=k.b.reporter 
INFO[0010] fastest 3329169 pkg=k.b.reporter 
INFO[0010] slowest 1827577178 pkg=k.b.reporter 
INFO[0010] total request size 53054705 pkg=k.b.reporter 
INFO[0010] toatl response size 0 pkg=k.b.reporter 
INFO[0010] 204: 26939 pkg=k.b.reporter 
INFO[0010] 0: 6 pkg=k.b.reporter 
INFO[0010] loader finished pkg=k.b.loader 

Duration: 5
Worker number: 1000
Batch size: 100
Timeout: 30
TargetDB: kairosdb
````

````bash
xkb --db k -c 5000
````

````
INFO[0008] basic report finished via context pkg=k.b.reporter 
INFO[0008] total request 42340 pkg=k.b.reporter 
INFO[0008] fastest 24020 pkg=k.b.reporter 
INFO[0008] slowest 2439599473 pkg=k.b.reporter 
INFO[0008] total request size 83367460 pkg=k.b.reporter 
INFO[0008] toatl response size 0 pkg=k.b.reporter 
INFO[0008] 0: 25834 pkg=k.b.reporter 
INFO[0008] 204: 16506 pkg=k.b.reporter 
INFO[0008] loader finished pkg=k.b.loader 

Duration: 5
Worker number: 5000
Batch size: 100
Timeout: 30
TargetDB: kairosdb
````

````bash
xkb --db k -c 100 -d 60
````

````
INFO[0062] basic report finished via context pkg=k.b.reporter 
INFO[0062] total request 307341 pkg=k.b.reporter 
INFO[0062] fastest 240496 pkg=k.b.reporter 
INFO[0062] slowest 236205198 pkg=k.b.reporter 
INFO[0062] total request size 605154429 pkg=k.b.reporter 
INFO[0062] toatl response size 0 pkg=k.b.reporter 
INFO[0062] 204: 307341 pkg=k.b.reporter 
INFO[0062] loader finished pkg=k.b.loader 

Duration: 60
Worker number: 100
Batch size: 100
Timeout: 30
TargetDB: kairosdb
````