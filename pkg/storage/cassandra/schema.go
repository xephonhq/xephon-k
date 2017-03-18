package cassandra

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/pkg/errors"
)

var createNaiveKeyspaceTmpl = `
CREATE KEYSPACE IF NOT EXISTS "%s"
  WITH REPLICATION = {
    'class' : 'SimpleStrategy',
    'replication_factor' : 1
  };
`

var createNaiveMetircTableTmpl = `
CREATE TABLE IF NOT EXISTS "%s".metrics (
  metric_name text,
  metric_timestamp timestamp,
  tags frozen<map<text, text>>,
  value int,
  PRIMARY KEY (metric_name, metric_timestamp, tags)
);
`

// CreateSchema use naive with tag schema
// TODO: allow passing different configuration
func CreateSchema() {
	err := CreateKeyspace()
	if err != nil {
		// TODO: print the trace?
		log.Fatal(err)
		return
	}
	err = CreateMetricTable()
	if err != nil {
		log.Fatal(err)
		return
	}
}

// CreateKeyspace creates naivewithtag keyspace
func CreateKeyspace() error {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "system"
	session, err := cluster.CreateSession()
	defer session.Close()
	if err != nil {
		return errors.Wrap(err, "can't connect using system keyspace")
	}
	log.Info("connected using system namespace")
	createKeyspaceStmt := fmt.Sprintf(createNaiveKeyspaceTmpl, naiveKeySpace)
	err = session.Query(createKeyspaceStmt).Exec()
	if err != nil {
		return errors.Wrap(err, "can't create naivewithtag keyspace")
	}
	log.Infof("keyspace %s created", naiveKeySpace)
	return nil
}

// CreateMetricTable creates naive mertic table that has tag but not bucket
func CreateMetricTable() error {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = naiveKeySpace
	session, err := cluster.CreateSession()
	defer session.Close()
	if err != nil {
		return errors.Wrap(err, "can't connect using naivewithtag keyspace")
	}
	log.Info("connected using naivewithtag namespace")
	createMetricTableStmt := fmt.Sprintf(createNaiveMetircTableTmpl, naiveKeySpace)
	err = session.Query(createMetricTableStmt).Exec()
	if err != nil {
		return errors.Wrap(err, "can't create metric table in naivewithtag keyspace")
	}
	log.Infof("metrics table created")
	return nil
}
