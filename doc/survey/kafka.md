# Kafka

Kafka is used by Druid, Spotify Heroic to improve performance.

## Design

- use mmap
  - they quote an article from Varnish author
- use sendfile to avoid extra copy
- use own binary format
