// Codinoc IDE
//
// Project by Group 35
// Computing Group Project
//
// Plymouth University UK
// NSBM Green University LK
//
// Database Connection

package main

import (
	"fmt"
	"database/sql"
)

func Connect_Database(db_conn *sql.DB, host string, port string, user string, password string, dbase string) *sql.DB {
	// TODO Add Sessions for Databases
	// https://github.com/antonlindstrom/pgstore

	psql_string := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
															host, port, user, password, dbase)

	db_conn, _ = sql.Open("postgres", psql_string)
	fmt.Println("> Database is Connected: " + dbase)

	return db_conn
}