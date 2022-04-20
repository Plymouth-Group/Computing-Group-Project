// Codinoc IDE
//
// Project by Group 35
// Computing Group Project
//
// Plymouth University UK
// NSBM Green University LK
//
// Database Connections

package main

import (
	"fmt"
	"database/sql"

	_ "github.com/lib/pq"
)

// Site Database Connect

func site_db_connect() {
	psql_site_db_info := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	DB_SITE_HOST, DB_SITE_PORT, DB_SITE_USER, DB_SITE_PASSWORD, DB_SITE_NAME)

	db_site, db_site_error := sql.Open("postgres", psql_site_db_info)

	if db_site_error != nil {
		panic(db_site_error)
		return
	}

	defer db_site.Close()

	db_site_error = db_site.Ping()

	if db_site_error != nil {
		panic(db_site_error)
		return
	}

	fmt.Println("> Site Database \"" + DB_SITE_NAME + "\" Connected")
}

// Server Database Connect

func server_db_connect(DB_SERVER_NAME string) {
	psql_site_db_info := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
	DB_SERVER_HOST, DB_SERVER_PORT, DB_SERVER_USER, DB_SERVER_PASSWORD, DB_SERVER_NAME)

	db_site, db_site_error := sql.Open("postgres", psql_site_db_info)

	if db_site_error != nil {
		panic(db_site_error)
		return
	}

	defer db_site.Close()

	db_site_error = db_site.Ping()

	if db_site_error != nil {
		panic(db_site_error)
		return
	}

	fmt.Println("> User / Server Database \"" + DB_SERVER_NAME + "\" Connected")
}