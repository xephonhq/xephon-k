package cassandra

import (
	"fmt"
	"github.com/gocql/gocql"
	"github.com/pkg/errors"
	"time"
)

var createKeyspaceTmpl = `
CREATE KEYSPACE IF NOT EXISTS "%s"
  WITH REPLICATION = {
    'class' : 'SimpleStrategy',
    'replication_factor' : 1
  };
`

var createNaiveMetricIntTable = `
CREATE TABLE IF NOT EXISTS metrics_int (
  metric_name text,
  metric_timestamp timestamp,
  tags frozen<map<text, text>>,
  value int,
  PRIMARY KEY ((metric_name, tags), metric_timestamp)
);
`

var createNaiveMetricDoubleTable = `
CREATE TABLE IF NOT EXISTS metrics_double (
  metric_name text,
  metric_timestamp timestamp,
  tags frozen<map<text, text>>,
  value double,
  PRIMARY KEY ((metric_name, tags), metric_timestamp)
);
`

// CreateSchema use naive with tag schema
// TODO: allow passing different configuration
func CreateSchema() {
	err := createKeyspace()
	if err != nil {
		// TODO: print the trace?
		log.Fatal(err)
		return
	}

	err = createMetricTables()
	if err != nil {
		// TODO: have better retry policies
		log.Info("need to sleep for 10 seconds to wait for cassandra to settle down")
		time.Sleep(10 * time.Second)
		log.Info("try to do it again")
		err = createMetricTables()
		if err != nil {
			log.Fatal(err)
			return
		}
	}
}

// createKeyspace creates default keyspace
func createKeyspace() error {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = "system"
	session, err := cluster.CreateSession()
	defer session.Close()
	if err != nil {
		return errors.Wrap(err, "can't connect using system keyspace")
	}
	log.Info("connected using system namespace")
	createKeyspaceStmt := fmt.Sprintf(createKeyspaceTmpl, defaultKeySpace)
	err = session.Query(createKeyspaceStmt).Exec()
	if err != nil {
		return errors.Wrapf(err, "can't create %s keyspace", defaultKeySpace)
	}
	log.Infof("keyspace %s created", defaultKeySpace)
	return nil
}

// createMetricTables creates metric tables that has tag but no bucket
func createMetricTables() error {
	cluster := gocql.NewCluster("127.0.0.1")
	cluster.Keyspace = defaultKeySpace
	session, err := cluster.CreateSession()
	defer session.Close()
	if err != nil {
		return errors.Wrapf(err, "can't connect using %s keyspace", defaultKeySpace)
	}
	log.Infof("connected using %s namespace", defaultKeySpace)
	// FIXME: it seems there are also timeout problem when creating table, just ignore it for now
	err = session.Query(createNaiveMetricIntTable).Exec()
	if err != nil {
		return errors.Wrapf(err, "can't create metrics_int table in %s keyspace", defaultKeySpace)
	}
	err = session.Query(createNaiveMetricDoubleTable).Exec()
	if err != nil {
		return errors.Wrapf(err, "can't create metrics_double table in %s keyspace", defaultKeySpace)
	}
	log.Infof("metrics tables created")
	return nil
}
