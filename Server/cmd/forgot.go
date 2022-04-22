// Codinoc IDE
//
// Project by Group 35
// Computing Group Project
//
// Plymouth University UK
// NSBM Green University LK
//
// Forgot Password Functions

package main

import (
	"fmt"
)

func Reset_password(server_code string, admin_email string, new_password string) {
	// Get raw id using server code

	var query_get_scode = fmt.Sprintf("SELECT data_id FROM schema_server.tb_server WHERE server_code='%s';", server_code)
	row_get_scode, err_row_get_scode := db_site.Query(query_get_scode)
	CheckError(err_row_get_scode)

	if row_get_scode.Next() != true {
		fmt.Println("> Unable to get column id: " + query_get_scode)
		return
	}

	var current_scode_id string
	row_get_scode.Scan(&current_scode_id)

	// get main table's raw id using that

	var query_get_tcode = fmt.Sprintf("SELECT user_id FROM schema_server.tb_main WHERE id_server='%s';", current_scode_id)
	row_get_tcode, err_row_get_tcode := db_site.Query(query_get_tcode)
	CheckError(err_row_get_tcode)

	if row_get_tcode.Next() != true {
		fmt.Println("> Unable to get column id: " + query_get_tcode)
		return
	}

	var current_tcode_id string
	row_get_tcode.Scan(&current_tcode_id)

	// get password tables raw id

	var query_get_pcode = fmt.Sprintf("SELECT id_password FROM schema_server.tb_main WHERE user_id='%s';", current_tcode_id)
	row_get_pcode, err_row_get_pcode := db_site.Query(query_get_pcode)
	CheckError(err_row_get_pcode)

	if row_get_pcode.Next() != true {
		fmt.Println("> Unable to get column id: " + query_get_pcode)
		return
	}

	var current_pcode_id string
	row_get_pcode.Scan(&current_pcode_id)

	// reset password for that email

	var query_update_psw = fmt.Sprintf("UPDATE schema_server.tb_password SET hashed_psw='%s'WHERE data_id='%s';", new_password, current_pcode_id)
	_, err_update_psq := db_site.Query(query_update_psw)
	CheckError(err_update_psq)
}