// Codinoc IDE
//
// Project by Group 35
// Computing Group Project
//
// Plymouth University UK
// NSBM Green University LK
//
// Sign In Functions

package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
)

// user sign in

func Session_SignIn(email_address string, response http.ResponseWriter) {
	value := map[string]string {
		"username": userName,
	}

	encoded, err := cookieHandler.Encode("session", value)

	if err == nil {
		cookie := &http.Cookie {
			Name: "session",
			Value: encoded,
			Path: "/",
		}

		http.SetCookie(response, cookie)
	}
}

// logged user sign out

func Session_SignOut(response http.ResponseWriter) {
	cookie := &http.Cookie {
		Name: "session",
		Value: "",
		Path: "/",
		MaxAge: -1,
	}

	http.SetCookie(response, cookie)
}

// sign in as server administrator

func Signin_admin_check(server_code string, admin_email string, admin_password string) bool {
	// Check server code

	var query_get_scode = fmt.Sprintf("SELECT data_id FROM schema_server.tb_server WHERE server_code='%s';", server_code)
	row_get_scode, err_row_get_scode := db_site.Query(query_get_scode)
	CheckError(err_row_get_scode)

	if row_get_scode.Next() != true {
		return false
	}

	var current_scode_id string
	row_get_scode.Scan(&current_scode_id)

	// Get main table column id

	var query_get_tcode = fmt.Sprintf("SELECT id_security FROM schema_server.tb_main WHERE id_server='%s';", current_scode_id)
	row_get_tcode, err_row_get_tcode := db_site.Query(query_get_tcode)
	CheckError(err_row_get_tcode)

	if row_get_tcode.Next() != true {
		return false
	}

	var current_tcode_id string
	row_get_tcode.Scan(&current_tcode_id)

	// Check email address

	var query_get_mcode = fmt.Sprintf("SELECT admin_email FROM schema_server.tb_security WHERE data_id='%s';", current_tcode_id)
	row_get_mcode, err_row_get_mcode := db_site.Query(query_get_mcode)
	CheckError(err_row_get_mcode)

	if row_get_mcode.Next() != true {
		return false
	}

	var current_mail string
	row_get_mcode.Scan(&current_mail)

	// Check given email address is same or not

	if current_mail != admin_email {
		return false
	}

	// Check password, we check this used by using same data_id used by email address check

	var query_get_pcode = fmt.Sprintf("SELECT hashed_psw FROM schema_server.tb_password WHERE data_id='%s';", current_tcode_id)
	row_get_pcode, err_row_get_pcode := db_site.Query(query_get_pcode)
	CheckError(err_row_get_pcode)

	if row_get_pcode.Next() != true {
		return false
	}

	var current_hashed_password string
	row_get_pcode.Scan(&current_hashed_password)

	// check given password is same or note

	if Check_Password_Hash(admin_password, current_hashed_password) == false {
		return false
	}

	return true
}