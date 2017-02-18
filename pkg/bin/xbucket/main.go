package main

import (
	"log"

	"fmt"

	"github.com/gocql/gocql"
)

// This demostrate how to use the bucket schema

var defaultKeyspace = "system"
var keyspace = "xephonbucket"

func main() {
	log.Print("Let's use bucket schema")
	cluster := gocql.NewCluster("127.0.0.1")
	// https://github.com/gocql/gocql/blob/master/common_test.go#L98
	// It seems you can connect to system keyspace to create new keyspace
	log.Print("Connect use system namespace")
	cluster.Keyspace = defaultKeyspace
	session, err := cluster.CreateSession()
	defer session.Close()
	if err != nil {
		log.Fatal(err)
	}
	log.Print("connected using system namespace")

	createKeyspaceStmt := fmt.Sprintf("CREATE KEYSPACE IF NOT EXISTS %s WITH REPLICATION = {'class':'SimpleStrategy', 'replication_factor':1}", keyspace)
	if err = session.Query(createKeyspaceStmt).Exec(); err != nil {
		log.Fatal(err)
	}
	log.Printf("keyspace %s created", keyspace)

	// TODO: can I create a new session using existing cluster variable, or do I need to create a new one
}
