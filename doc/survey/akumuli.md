# Akumuli

- http://akumuli.org/akumuli/2016/09/14/next/
  - https://docs.google.com/document/d/1jFK8E3CZSqR5IPsMGojm2LknkNyUZA7tY51N6IgzW_g/pub
- https://github.com/at15/papers-i-read/issues/39

## TODO

- [ ] How many level would be in memory
- [ ] If want to be distributed, how to handle the address (pointer), file offset?
partition a large series? what's the physical limit of a single series?
- [ ] interleave different series would cause little performance penalty when read on SSD, but would
it cause trouble on spin drive? The block size is pretty small.
- [ ] In leaf node, how are the compressed data points stored? `t1, v1, t2, v2` or `t1, t2, v1, v2`

## Meta



## Vocabulary

Bear with my English

- fan-out: the number of inputs that can be connected to a specified output.

## NB+ tree

https://github.com/akumuli/Akumuli/blob/master/libakumuli/storage_engine/nbtree.h

- [ ] the topmost superblock is called rightmost https://github.com/akumuli/Akumuli/blob/master/libakumuli/storage_engine/nbtree.h#L62
which might explain the graph I didn't understand in the Google doc

Inner Node

- 4KB (2^12 bytes)
- 32 links, each contains aggregate (each link is 2^7 bytes)
  - first & last point timestamp and value (2^5 bytes = 64 bit * 2 * 2)
  - small & largest timestamp and value (2^5 bytes)
  - sum of all values in subtree (2^3 bytes)
    - [ ] what if it overflow
  - number of points in the subtree (2^3 bytes)
  - [ ] when to update the aggregate values, when the subtree is full?

Links

- inner node -> inner node
- inner node -> leaf node
- leaf node -> leaf node
- back references during crash recovery
- [ ] no link from leaf node to inner node?
