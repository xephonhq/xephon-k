# Stories from the Trenches – The Challenges of Building an Analytics Stack

- https://www.youtube.com/watch?v=Sz4w75xRrYM

## Take away

- https://dev.twitter.com/streaming/public

## Notes

![druid](druid.png)

- Immutable data
- In Memory is Overrated
  - mmap + SSD
  - cost of scaling CPU << cost of adding RAM
  - decompression on the fly (LZF, Snappy, LZ4)
- Low Latency vs High throughput
  - combine batch + streaming
  - immutable made it easy to combine the two ingestion methods
  - Makes for easy backfill and re-processing
  - Historical Node
  - Real-time Node
- Not All Data is Created Equal
  - user really care about recent data
  - user still want to run quarterly report
  - large queries create bottlenecks and resource contention
- Smarter Rebalancing
- Create Data Tiers
- Addressing Multitenancy
  - HyperLogLog sketches
  - Approximate top-k
  - Approximate histograms (monitoring)
- Monitoring
  - Use Druid to monitor Druid
- **Use cases should define engineering**
