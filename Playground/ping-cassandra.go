package main

import (
	"fmt"
	"github.com/gocql/gocql"
	"log"
	"time"
)

const ()

func main() {
	start := time.Now()
	// Connect to the cluster
	cluster := gocql.NewCluster("192.168.y.y", "192.168.y.y")
	cluster.Keyspace = "Testing_Keyspace"
	cluster.Consistency = gocql.Quorum
	session, err := cluster.CreateSession()
	defer session.Close()
	if err != nil {
		fmt.Println("Something wrong in connecting to the Cassandra cluster..\n", err)
	}

	var empid int64
	var empname string
	var cofounder bool

	if err := session.Query(`SELECT * FROM Testing_Keyspace.emptable LIMIT 1;`).Consistency(gocql.One).Scan(&empid, &empname, &cofounder); err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	fmt.Println(time.Now(), " Response: ", elapsed.Seconds(), "ms", "> 100ms?: ", elapsed.Seconds() > 0.1)
}
