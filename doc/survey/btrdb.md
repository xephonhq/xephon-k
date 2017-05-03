# BTrDB

- https://www.usenix.org/sites/default/files/conference/protected-files/fast16_sildes_andersen.pdf

## TODO

- [ ] No timestamp in internal nodes?
  - it's time partitioned, so each block has fixed time range?
  - [ ] does it mentioned hole?
- [ ] How the version number is set?
- [ ] How is the out of order handled?
- [ ] Crash recovery? seems not mentioned at all?

## Meta

- 53 million values / second insert (four-node)
- 119 million queried values / second read (four-node)

## Special of telemetry

- arrives out of order
- delayed and duplicated
- nanosecond timestamp

## Time partitioned Tree

- time partitioning
- multi resolution
- version annotated
- [ ] copy on write
- k-ary tree

Internal Node

- Each internal node holds scalar summaries of the sub-trees below it, along with the links to the subtrees.
  - min
  - mean
  - max
  - count
- When querying a stream for statistical records, the tree need only be traversed to the depth corresponding to the desired resolution.
- 2 (address + version) * 8 (64 bit) * K (K = 64) = 1024 KB
- 4 (min,max,mean,count) *  8 (64 bit) * K (K = 64) = 2048 KB
- total 3 KB
- [ ] No timestamp in internal nodes?

Leaf

- 16KB (1024 points)

Block Store

- [ ] Fields such as a block’s address, UUID, resolution (tree depth) and time extent are useful for traversing the
tree, but can be deduced from context when a block is read from disk, so are stripped before the block enters the compression engine
- a pool of pre-created initial addresses

Delta Delta compression
