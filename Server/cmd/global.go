// Codinoc IDE
//
// Project by Group 35
// Computing Group Project
//
// Plymouth University UK
// NSBM Green University LK
//
// Global Variables

package main

import (
	"database/sql"
)

const (
	CONN_HOST = "127.0.0.1"
	CONN_PORT = "8080"

	DB_HOST   = "127.0.0.1"
	DB_PORT   = "5432"
	DB_USER   = "admin_codinoc"
	DB_PASS   = "1234"
)

var db_site *sql.DB
var db_server *sql.DB
