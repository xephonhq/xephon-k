# ScyllaDB

- http://www.scylladb.com/
- https://github.com/scylladb/scylla
- https://github.com/scylladb/scylla-grafana-monitoring it's quite funny they use prometheus for monitoring

## Usage

- `docker run --name tsdb-scylla -p 9042:9042 -d scylladb/scylla:1.6.0`
  - `docker stop tsdb-scylla`
  - `docker start tsdb-scylla`
