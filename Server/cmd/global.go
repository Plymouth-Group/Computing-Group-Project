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
	"github.com/gorilla/securecookie"
	_ "github.com/gorilla/sessions"
)

const (
	CONN_HOST = "127.0.0.1"
	CONN_PORT = "8080"
)

const (
	DB_SITE_HOST = "127.0.0.1"
	DB_SITE_PORT = "5432"
	DB_SITE_USER = "admin_codinoc"
	DB_SITE_PASSWORD = "1234"
	DB_SITE_NAME = "db_site"
)

const (
	DB_SERVER_HOST = "127.0.0.1"
	DB_SERVER_PORT = "5432"
	DB_SERVER_USER = "admin_codinoc"
	DB_SERVER_PASSWORD = "1234"
)

var db_site *sql.DB
var db_server *sql.DB

type Sign_up struct {
	server_name string
	server_code string
	admin_email string
	admin_password string

	is_account_created bool
}

type Sign_in struct {
	server_code string

	admin_email string
	admin_password string

	member_uname string
	member_password string

	is_logged bool
}

type Forgot_password struct {
	server_code string
	admin_email string
	new_password string

	is_password_reset_available bool
}

var glob_sign_up Sign_up
var glob_sign_in Sign_in
var glob_forgot_password Forgot_password

// var sign_in_store = sessions.NewCookieStore([byte]("login_session"))

var cookieHandler = securecookie.New(
	securecookie.GenerateRandomKey(64),
	securecookie.GenerateRandomKey(32))