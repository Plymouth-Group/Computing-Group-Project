// Codinoc IDE
//
// Project by Group 35
// Computing Group Project
//
// Plymouth University UK
// NSBM Green University LK
//
// Logged Dashboard

package main

import (
	"fmt"
)

func Get_ServerCode() string {
	var query_get_scode = fmt.Sprintf("SELECT server_code FROM schema_server.tb_main WHERE server_id='1';")
	row_get_scode, err_row_get_scode := db_server.Query(query_get_scode)
	CheckError(err_row_get_scode)

	if row_get_scode.Next() != true {
		return ""
	}

	var current_scode string
	row_get_scode.Scan(&current_scode)

	return current_scode
}

func Get_DashboardName() string {
	var query_get_dname = fmt.Sprintf("SELECT dashboard_name FROM schema_server.tb_main WHERE server_id='1';")
	row_get_dname, err_row_get_dname := db_server.Query(query_get_dname)
	CheckError(err_row_get_dname)

	if row_get_dname.Next() != true {
		return ""
	}

	var current_dname string
	row_get_dname.Scan(&current_dname)

	return current_dname
}

func Get_DashboardDescription() string {
	var query_get_ddescription = fmt.Sprintf("SELECT dashboard_description FROM schema_server.tb_main WHERE server_id='1';")
	row_get_ddescription, err_row_get_ddescription := db_server.Query(query_get_ddescription)
	CheckError(err_row_get_ddescription)

	if row_get_ddescription.Next() != true {
		return ""
	}

	var current_ddescription string
	row_get_ddescription.Scan(&current_ddescription)

	return current_ddescription
}

func Get_WikiContent() string {
	var query_get_wcontent = fmt.Sprintf("SELECT dashboard_wiki FROM schema_server.tb_main WHERE server_id='1';")
	row_get_wcontent, err_row_get_wcontent := db_server.Query(query_get_wcontent)
	CheckError(err_row_get_wcontent)

	if row_get_wcontent.Next() != true {
		return ""
	}

	var current_wcontent string
	row_get_wcontent.Scan(&current_wcontent)

	return current_wcontent
}