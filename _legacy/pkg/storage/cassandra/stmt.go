package cassandra

// read statements

var selectIntStmt = `
	SELECT metric_timestamp, value FROM metrics_int WHERE metric_name = ? AND tags = ?
	`

var selectIntByStartEndTimeStmt = `
	SELECT metric_timestamp, value FROM metrics_int WHERE metric_name = ? AND tags = ? AND metric_timestamp >= ? AND metric_timestamp <= ?
	`

// write statements

var insertIntStmt = `
	INSERT INTO metrics_int (metric_name, metric_timestamp, tags, value) VALUES (?, ?, ?, ?)
	`

var insertDoubleStmt = `
	INSERT INTO metrics_double (metric_name, metric_timestamp, tags, value) VALUES (?, ?, ?, ?)
	`
