// Codinoc IDE
//
// Project by Group 35
// Computing Group Project
//
// Plymouth University UK
// NSBM Green University LK
//
// Process Functions

package main

import (
	"fmt"
	"strings"
	_ "github.com/lib/pq"
)

func Process_Is_Server_Available(server_name string) bool {
	server_code := strings.ReplaceAll(server_name, " ", "")

	var query_get_scode = fmt.Sprintf("SELECT data_id FROM schema_server.tb_server WHERE server_code='%s';", server_code)
	row_get_scode, err_row_get_scode := db_site.Query(query_get_scode)
	Check_Error(err_row_get_scode)

	if row_get_scode.Next() == true {
		return true
	}

	return false
}

func Process_Is_Server_Available_From_Code(server_code string) bool {
	var query_get_scode = fmt.Sprintf("SELECT data_id FROM schema_server.tb_server WHERE server_code='%s';", server_code)
	row_get_scode, err_row_get_scode := db_site.Query(query_get_scode)
	Check_Error(err_row_get_scode)

	if row_get_scode.Next() == true {
		return true
	}

	return false
}

func Process_Is_Admin_Available_In_Server(server_code string, admin_email string) bool {
	var query_get_scode = fmt.Sprintf("SELECT data_id FROM schema_server.tb_server WHERE server_code='%s';", server_code)
	row_get_scode, err_row_get_scode := db_site.Query(query_get_scode)
	Check_Error(err_row_get_scode)

	if row_get_scode.Next() == true {
		return true
	}

	var current_scode_id string
	row_get_scode.Scan(&current_scode_id)

	var query_get_email = fmt.Sprintf("SELECT admin_email FROM schema_server.tb_security WHERE data_id='%s';", current_scode_id)
	row_get_email, err_row_get_email := db_site.Query(query_get_email)
	Check_Error(err_row_get_email)

	if row_get_email.Next() == true {
		return true
	}

	return false
}