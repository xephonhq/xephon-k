# Query Process

a6c543039763c0f08253d71a43aefe3b570ecf37

- services/httpd/handler.go
  - L431 `results := h.QueryExecutor.ExecuteQuery(query, opts, closing)`
- influxql/query_executor.go
  - L328 `err := e.StatementExecutor.ExecuteStatement(stmt, ctx)`
- influxdb/coordinator/statement_executor.go
  - L54 `func (e *StatementExecutor) ExecuteStatement(stmt influxql.Statement, ctx influxql.ExecutionContext) error`
    - L57 `e.executeSelectStatement(stmt, &ctx)`
  - L432 `func (e *StatementExecutor) executeSelectStatement(stmt *influxql.SelectStatement, ctx *influxql.ExecutionContext)`
    - L433 `itrs, stmt, err := e.createIterators(stmt, ctx)`
    - L439 `em := influxql.NewEmitter(itrs, stmt.TimeAscending(), ctx.ChunkSize)`
    - L457 `row, partial, err := em.Emit()`
- [ ] the emitter queries from iterator, does the iterator query from cache first and then query from disk?
