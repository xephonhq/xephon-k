# Write Process

Deprecated, see [write-path](write-path.md)

a6c543039763c0f08253d71a43aefe3b570ecf37

- services/httpd/handler.go
  - L581 `func (h *Handler) serveWrite`
    - L673 `if err := h.PointsWriter.WritePoints(database, r.URL.Query().Get("rp"), consistency, points)`
- coordinator/points_writer.go
  - L288 `func (w *PointsWriter) WritePoints(database, retentionPolicy string, consistencyLevel models.ConsistencyLevel, points []models.Point`
    - L309 `ch <- w.writeToShard(shard, database, retentionPolicy, points)`
  - L353 `func (w *PointsWriter) writeToShard(shard *meta.ShardInfo, database, retentionPolicy string, points []models.Point)`
- tsdb/store.go
  - L888 `func (s *Store) WriteToShard(shardID uint64, points []models.Point)`
- tsdb/shard.go
  - L433 `func (s *Shard) WritePoints(points []models.Point) error`
    - L462 `if err := s.engine.WritePoints(points);`
- tsdb/engine/tsm1/engine.go
  - L804 `func (e *Engine) WritePoints(points []models.Point) error`
    - L854 `err := e.Cache.WriteMulti(values)`
    - L859 `_, err = e.WAL.WriteMulti(values)`
- tsdb/engine/tsm1/cache.go
  - L258 `func (c *Cache) Write(key string, values []Value) error`
    - L270 `if err := c.store.write(key, values); `
  - L178 `type storer interface`
- tsdb/engine/tsm1/ring.go` implements storer
  -
