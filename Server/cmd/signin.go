// Codinoc IDE
//
// Project by Group 35
// Computing Group Project
//
// Plymouth University UK
// NSBM Green University LK
//
// Sign In

package main

import (
	"fmt"
	"net/http"
)

func SignIn_Database_Connect(server_code string) {
	tenant_server := fmt.Sprintf("db_%s", server_code)

	db_server = Connect_Database(db_server, DB_HOST, DB_PORT, DB_USER, DB_PASS, tenant_server)
}

func SignIn_Session_Enable(w http.ResponseWriter, r *http.Request, session_string string, session_type string, session_code string) {
	session_login, _ := store_login.Get(r, "LOGIN_SESSION")

	// There are three types of Session Types
	//
	// 1. " main "  : Full Site Administrator
	// 2. " admin " : Codinoc Server Administrator
	// 3. " member ": Codinoc Server Member
	//
	// And Session Code means current logged server's code
	//
	// When compare or check session as a string, convert it
	// to a string before using

	session_login.Values["session_string"] = session_string
	session_login.Values["session_type"] = session_type
	session_login.Values["session_code"] = session_code
	session_login.Save(r, w)

	fmt.Println("> Session Opened: ", session_login.Values["session_string"])
}

func SignIn_Session_Disable(w http.ResponseWriter, r *http.Request) {
	session_login, _ := store_login.Get(r, "LOGIN_SESSION")

	if session_login != nil {
		session_login.Values["session_string"] = nil
		session_login.Values["session_type"] = nil
		session_login.Values["session_code"] = nil

		session_login.Options.MaxAge = -1
		session_login.Save(r, w)

		fmt.Println("> Session Closed: ", session_login.Values["session_string"])

		session_login = nil
	}

	if db_server != nil {
		db_server.Close()
	}
}