package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"
)
import "github.com/jackc/pgx/v4/pgxpool"

func main() {
	host := flag.String("host", "127.0.0.1", "PG host")
	port := flag.Int("port", 5436, "PG port")
	user := flag.String("user", "postgres", "PG username")
	pass := flag.String("pass", "postgres", "PG password")
	sysdb := flag.String("sysdb", "postgres", "PG sys table nasme")
	cleanPrefix := flag.String("prefix", "test_", "Prefix for databases to clean")
	listOnly := flag.Bool("listonly", false, "Only list no drop")

	flag.Parse()

	if len(*cleanPrefix) == 0 {
		fmt.Printf("empty prefix\n")
		os.Exit(10)
	}

	if !strings.Contains(*cleanPrefix, "_") {
		fmt.Printf("invalid dbase prefix: `%s`\n", *cleanPrefix)
		os.Exit(10)
	}

	dbUrl := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", *user, *pass, *host, *port, *sysdb)

	connection, err := pgxpool.Connect(context.Background(), dbUrl)

	if err != nil {
		fmt.Printf("connection error\n%v\n", err)
		os.Exit(1)
	}

	defer connection.Close()

	databasesToDeleteQuery := fmt.Sprintf("SELECT datname FROM pg_catalog.pg_database pd WHERE datname LIKE '%s%%'", *cleanPrefix)

	databaseNamesRowSet, err := connection.Query(context.Background(), databasesToDeleteQuery)

	if err != nil {
		fmt.Printf("database list query\n%v\n", err)
		os.Exit(2)
	}

	for databaseNamesRowSet.Next() {
		if err := databaseNamesRowSet.Err(); err != nil {
			fmt.Printf("database row error\n%v\n", err)
			os.Exit(3)
		}
		var databaseName string
		if err := databaseNamesRowSet.Scan(&databaseName); err != nil {
			fmt.Printf("row scan error\n%v\n", err)
			os.Exit(4)
		}
		if *listOnly {
			fmt.Println(databaseName)
		} else {
			fmt.Printf("begin drop database `%s` terminate ...\n", databaseName)
			terminateConnectionsForDb := fmt.Sprintf(
				"SELECT pg_terminate_backend(pid) FROM pg_catalog.pg_stat_activity  WHERE datname = '%s';", databaseName,
			)
			if _, err := connection.Exec(context.Background(), terminateConnectionsForDb); err != nil {
				fmt.Printf("cannot terminate sessions\n%v\n", err)
				os.Exit(5)
			}
			fmt.Println("terminated, dropping ...")
			dropDatabaseQuery := fmt.Sprintf("DROP DATABASE %s;", databaseName)
			if _, err := connection.Exec(context.Background(), dropDatabaseQuery); err != nil {
				fmt.Printf("cannot drop database\n%v\n", err)
				os.Exit(6)
			}
			fmt.Println("successfully dropped")
			fmt.Println()
		}
	}
}
