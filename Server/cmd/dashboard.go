// Codinoc IDE
//
// Project by Group 35
// Computing Group Project
//
// Plymouth University UK
// NSBM Green University LK
//
// Dashboard

package main

import (
	"fmt"
)

func Dashboard_Get_Server_Name() string {
	var query_get_sname = fmt.Sprintf("SELECT server_name FROM schema_server.tb_main WHERE server_id='1';")
	row_get_sname, err_row_get_sname := db_server.Query(query_get_sname)
	Check_Error(err_row_get_sname)

	if row_get_sname.Next() != true {
	}

	var current_sname string
	row_get_sname.Scan(&current_sname)

	return current_sname
}

func Dashboard_Get_Server_Code() string {
	var query_get_scode = fmt.Sprintf("SELECT server_code FROM schema_server.tb_main WHERE server_id='1';")
	row_get_scode, err_row_get_scode := db_server.Query(query_get_scode)
	Check_Error(err_row_get_scode)

	if row_get_scode.Next() != true {
	}

	var current_scode string
	row_get_scode.Scan(&current_scode)

	return current_scode
}

func Dashboard_Get_Dashboard_Name() string {
	var query_get_dname = fmt.Sprintf("SELECT dashboard_name FROM schema_server.tb_main WHERE server_id='1';")
	row_get_dname, err_row_get_dname := db_server.Query(query_get_dname)
	Check_Error(err_row_get_dname)

	if row_get_dname.Next() != true {
	}

	var current_dname string
	row_get_dname.Scan(&current_dname)

	return current_dname
}

func Dashboard_Get_Dashboard_Description() string {
	var query_get_ddescription = fmt.Sprintf("SELECT dashboard_description FROM schema_server.tb_main WHERE server_id='1';")
	row_get_ddescription, err_row_get_ddescription := db_server.Query(query_get_ddescription)
	Check_Error(err_row_get_ddescription)

	if row_get_ddescription.Next() != true {
	}

	var current_ddescription string
	row_get_ddescription.Scan(&current_ddescription)

	return current_ddescription
}

func Dashboard_Get_WikiContent() string {
	var query_get_wcontent = fmt.Sprintf("SELECT dashboard_wiki FROM schema_server.tb_main WHERE server_id='1';")
	row_get_wcontent, err_row_get_wcontent := db_server.Query(query_get_wcontent)
	Check_Error(err_row_get_wcontent)

	if row_get_wcontent.Next() != true {
	}

	var current_wcontent string
	row_get_wcontent.Scan(&current_wcontent)

	return current_wcontent
}

func Dashboard_Get_Account_Type(session_type string) string {
	if session_type == "admin" {
		return "Server Administrator"
	}

	return "Server team member"
}

func Dashboard_Get_First_Name(session_type string, server_code string) string {
	if session_type == "admin" {
		var query_get_current = fmt.Sprintf("SELECT data_id FROM schema_server.tb_server WHERE server_code='%s';", server_code)
		row_get_current, err_row_get_current := db_site.Query(query_get_current)
		Check_Error(err_row_get_current)

		if row_get_current.Next() != true {
			fmt.Println("> Unable to get column id: First Name")
			return ""
		}

		var current_column_id string
		row_get_current.Scan(&current_column_id)

		var query_get = fmt.Sprintf("SELECT first_name FROM schema_server.tb_info WHERE data_id='%s';", current_column_id)
		row_get, err_row_get := db_site.Query(query_get)
		Check_Error(err_row_get)

		row_get.Next()

		var current string
		row_get.Scan(&current)

		return current
	} else {
		// TODO Write code for server team member
	}

	return ""
}

func Dashboard_Get_Middle_Name(session_type string, server_code string) string {
	if session_type == "admin" {
		var query_get_current = fmt.Sprintf("SELECT data_id FROM schema_server.tb_server WHERE server_code='%s';", server_code)
		row_get_current, err_row_get_current := db_site.Query(query_get_current)
		Check_Error(err_row_get_current)

		if row_get_current.Next() != true {
			fmt.Println("> Unable to get column id: Middle Name")
			return ""
		}

		var current_column_id string
		row_get_current.Scan(&current_column_id)

		var query_get = fmt.Sprintf("SELECT middle_name FROM schema_server.tb_info WHERE data_id='%s';", current_column_id)
		row_get, err_row_get := db_site.Query(query_get)
		Check_Error(err_row_get)

		row_get.Next()

		var current string
		row_get.Scan(&current)

		return current
	} else {
		// TODO Write code for server team member
	}

	return ""
}

func Dashboard_Get_Last_Name(session_type string, server_code string) string {
	if session_type == "admin" {
		var query_get_current = fmt.Sprintf("SELECT data_id FROM schema_server.tb_server WHERE server_code='%s';", server_code)
		row_get_current, err_row_get_current := db_site.Query(query_get_current)
		Check_Error(err_row_get_current)

		if row_get_current.Next() != true {
			fmt.Println("> Unable to get column id: Last Name")
			return ""
		}

		var current_column_id string
		row_get_current.Scan(&current_column_id)

		var query_get = fmt.Sprintf("SELECT last_name FROM schema_server.tb_info WHERE data_id='%s';", current_column_id)
		row_get, err_row_get := db_site.Query(query_get)
		Check_Error(err_row_get)

		row_get.Next()

		var current string
		row_get.Scan(&current)

		return current
	} else {
		// TODO Write code for server team member
	}

	return ""
}

func Dashboard_Get_User_Name(session_type string, server_code string) string {
	// If logged as an administrator, the user name is user's email address
	// nor user name

	if session_type == "admin" {
		var query_get_current = fmt.Sprintf("SELECT data_id FROM schema_server.tb_server WHERE server_code='%s';", server_code)
		row_get_current, err_row_get_current := db_site.Query(query_get_current)
		Check_Error(err_row_get_current)

		if row_get_current.Next() != true {
			fmt.Println("> Unable to get column id: Email Address")
			return ""
		}

		var current_column_id string
		row_get_current.Scan(&current_column_id)

		var query_get = fmt.Sprintf("SELECT admin_email FROM schema_server.tb_security WHERE data_id='%s';", current_column_id)
		row_get, err_row_get := db_site.Query(query_get)
		Check_Error(err_row_get)

		row_get.Next()

		var current string
		row_get.Scan(&current)

		return current
	} else {
		// TODO Write code for server team member
	}

	return ""
}

func Update_Wiki(content string) {
	var query_update_wcontent = fmt.Sprintf("UPDATE schema_server.tb_main SET dashboard_wiki='%s' WHERE server_id='1';", content)
	_, err_row_update_wcontent := db_server.Query(query_update_wcontent)
	CheckError(err_row_update_wcontent)
}