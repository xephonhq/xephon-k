# Cassandra

NOTE: moved to [separated folder](cassandra)

## Usage

For Fedora `sudo systemctl start docker.service` might be needed

- `docker run --name tsdb-cassandra -p 9042:9042 -d cassandra:3.9`
  - `docker stop tsdb-cassandra`
  - `docker start tsdb-cassandra`
- `docker exec -it tsdb-cassandra bash`

## TODO:

- [ ] keys http://stackoverflow.com/questions/24949676/difference-between-partition-key-composite-key-and-clustering-key-in-cassandra

## Compact storage

- https://docs.datastax.com/en/cql/3.3/cql/cql_reference/cqlCreateTable.html#refClstrOrdr__cql-compact-storage
- **For Cassandra 3.0 and later, the storage engine is much more efficient at storing data, and compact storage is not necessary.**
  - [ ] TODO: http://www.datastax.com/2015/12/storage-engine-30

## Wide row

- [ ] TODO: underlying storage
- [ ] TODO: in old thrift way
- [ ] TODO: in CQL
