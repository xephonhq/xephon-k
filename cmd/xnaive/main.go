package main

// This demostrate how to use the naive schema
// Take away
// Cqlsh have different behaviour than CQL API
// - `SELECT * FROM metrics` works in CQL, but since it does not specify row key, it involve scanning all the nodes, so
//    CQL API will return error: Cannot execute this query as it might involve data filtering and thus may have unpredictable performance.
//    If you want to execute this query despite the performance unpredictability, use ALLOW FILTERING
// - use single quote for string value and double quote for keyspace, column etc.
// gocql
// - C* use millisecond for timestamp, so when inserting using number instead string for timestamp, * 1000 for time.Now().Unix()
// - when SELECT *, you must have enough variables when Scan, otherwise specify what you want to select in the select statement

import (
	"log"
	"time"

	"fmt"

	"github.com/gocql/gocql"
)

var keyspace = "xephonnaive"

func main() {
	log.Print("Let's use naive schema")

	cluster := gocql.NewCluster("127.0.0.1")
	// TODO: will gocql create Keyspace automatically?
	cluster.Keyspace = keyspace
	session, err := cluster.CreateSession()
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("connected")

	// create the naive table
	createTableStmt := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s.metrics (metric_name text, metric_timestamp timestamp, value int, PRIMARY KEY (metric_name, metric_timestamp))", keyspace)

	if err = session.Query(createTableStmt).Exec(); err != nil {
		log.Fatal(err)
	}

	// insert into the naive table
	insertStmt := fmt.Sprintf("INSERT INTO %s.metrics (metric_name, metric_timestamp, value) VALUES (?, ?, ?)", keyspace)
	// NOTE: should use millisecond instead of second
	// if err = session.Query(insertStmt, "cpu.load", time.Now().Unix(), 10).Exec(); err != nil {
	if err = session.Query(insertStmt, "cpu.load", time.Now().Unix()*1000, 10).Exec(); err != nil {
		log.Fatal(err)
	}

	var selectAll = func() error {
		// fetch them all back
		// FIXED: should be < current timestamp
		// selectStmt := fmt.Sprintf("SELECT * FROM %s.metrics WHERE metric_timestamp > ?", keyspace)
		selectStmt := fmt.Sprintf("SELECT * FROM %s.metrics WHERE metric_timestamp < ?", keyspace)

		// FIXED: how to know if the iter works, since err is not returned?
		// we have iter.Close()
		iter := session.Query(selectStmt, time.Now().Unix()).Iter()
		var metricName string
		var metricTimestamp time.Time
		for iter.Scan(&metricName, &metricTimestamp) {
			log.Println(metricName, metricTimestamp)
		}
		// FIXED: Cannot execute this query as it might involve data filtering and thus may have unpredictable performance.
		// If you want to execute this query despite the performance unpredictability, use ALLOW FILTERING
		if err := iter.Close(); err != nil {
			return err
		}
		return nil
	}

	err = selectAll()
	// NOTE: there should be error
	log.Println(err)

	var selectByRowKey = func() error {
		selectCPU := fmt.Sprintf("SELECT * FROM %s.metrics WHERE metric_name = ? AND metric_timestamp < ?", keyspace)
		// NOTE: 1486320101000 works but 1486320101 not, ... it should be millisecond
		// iter := session.Query(selectCPU, "cpu.load", time.Now().Unix()).Iter()
		iter := session.Query(selectCPU, "cpu.load", time.Now().Unix()*1000).Iter()
		var metricName string
		var metricTimestamp time.Time
		var value int
		// FIXED: 2017/02/05 11:11:45 gocql: not enough columns to scan into: have 2 want 3
		// 	for iter.Scan(&metricName, &metricTimestamp) {
		for iter.Scan(&metricName, &metricTimestamp, &value) {
			log.Println(metricName, metricTimestamp, value)
		}
		// FIXED: the result time is pretty strange, don't know who to blame, gocql?
		// 2017/02/05 11:13:48 cpu.load 1970-01-18 04:52:00.432 +0000 UTC 10
		// 2017/02/05 11:13:48 cpu.load 1970-01-18 04:52:00.896 +0000 UTC 10
		// 2017/02/05 11:13:48 cpu.load 1970-01-18 04:52:01.615 +0000 UTC 10
		// 2017/02/05 11:13:48 cpu.load 1970-01-18 04:52:01.853 +0000 UTC 10
		// 2017/02/05 11:13:48 cpu.load 1970-01-18 04:52:01.905 +0000 UTC 10
		// NOTE: 1486320101000 works but 1486320101 not, ... it should be millisecond

		if err := iter.Close(); err != nil {
			return err
		}
		return nil
	}

	err = selectByRowKey()
	// NOTE: should be nil
	log.Println(err)
}
