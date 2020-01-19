# KairosDB

- http://kairosdb.github.io/

## Usage

Use docker-compose to run KairosDB with Cassandra

- `./up.sh` or `docker-compose build`, `docker-compose up`

Build and run a single KairosDB node with H2

- rename `Dockerfile.h2` to `Dockerfile` in `node` folder
- `docker build -t xephonhq/kairosdb ./node`
- `docker run -p 8080:8080 --name xephonhq-kairosdb xephonhq/kairosdb`

## NOTE

- [wait-for-it](https://github.com/vishnubob/wait-for-it) does not support alpine, some pr do help https://github.com/vishnubob/wait-for-it/pull/6

## Requirement

- [official docker support for Kairosdb?](https://github.com/kairosdb/kairosdb/issues/288)
- JDK7/8
- Cassandra 2.2 with Thrift enabled
- Cassnadra start and listen on thrift port before KairosDB start

## TODO

- [ ] switch to Oracle JDK for Cassandra
- [ ] switch to Oracle JDK for KairosDB
- [ ] support 3 Node Cassandra cluster (TODO: does C* has 2n+1 requirement?)
- [ ] mount volume

## Docker images

They are listed for reference but only Cassandra image is used

- https://github.com/cit-lab/kairosdb/tree/feature/alpine
