# Respawn

- https://users.ece.cmu.edu/~agr/resources/publications/respawn-rtss-13.pdf

## Take away

- http://www.cs.cmu.edu/~ishafer/ is a guy interested in time series
- 'The availability of fresh (recent) data often conflicts with the throughput gains achievable by batching data.'
- https://github.com/BodyTrack/datastore
- Dataseries: An efficient, flexible data format for structured serial data
- https://github.com/stevedh/readingdb (BTDS) (by the same group of guys from UCB, software defined buildings)
- Typical operations performed on time series data
  - plotting
  - zooming
  - correlation
  - clustering
  - prediction
  - pattern matching
  - summarization
- Predictive caching
  - Periodic Migration: when browsing, clients will frequently request tiles with the lowest resolution first
  - Proactive Migration: based on standard deviation, especially effective at accelerating access time for sparse data feeds
- SQLite > BTDS > MySQL > OpenTSDB

## Meta

- edge device (i.e. ARM)
- cloud-to-edge partitioning
- multi-resolution storage
  - high resolution at edge
  - low resolution at cloud backend
  - dispatcher + bloom filter
- Low-level (higher resolution) tiles are migrated based on both client access patterns and based on data metrics like standard deviation.
