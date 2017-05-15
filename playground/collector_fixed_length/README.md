# Collector have fixed length series

https://github.com/xephonhq/xephon-k/issues/33

- [mem-1.json](mem-1.json) is the first payload after the collector is started
- [mem-2.json](mem-2.json) shows the serializer is not reset between each send
  - [ ] but the server does not report any error? Yeah ... it seems they drop that part silently