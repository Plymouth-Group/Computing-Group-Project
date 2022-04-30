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
	"os"
	"strings"
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

func Dashboard_Update_Wiki(content string) {
	var query_update_wcontent = fmt.Sprintf("UPDATE schema_server.tb_main SET dashboard_wiki='%s' WHERE server_id='1';", content)
	_, err_row_update_wcontent := db_server.Query(query_update_wcontent)
	Check_Error(err_row_update_wcontent)
}

func Dashboard_Update_User_Profile(session_type string, session_code string, first_name string, middle_name string, last_name string, password string) {
	if session_type == "admin" {
		var query_get_current = fmt.Sprintf("SELECT data_id FROM schema_server.tb_server WHERE server_code='%s';", session_code)
		row_get_current, err_row_get_current := db_site.Query(query_get_current)
		Check_Error(err_row_get_current)

		if row_get_current.Next() != true {
			fmt.Println("> Unable to get column id")
			return
		}

		var current_column_id string
		row_get_current.Scan(&current_column_id)

		var query_update_name = fmt.Sprintf("UPDATE schema_server.tb_info SET first_name='%s', middle_name='%s', last_name='%s' WHERE data_id='%s';", first_name, middle_name, last_name, current_column_id)
		_, err_row_update_name := db_site.Query(query_update_name)
		Check_Error(err_row_update_name)

		if password != "" {
			new_password := Hash_Password(password)

			var query_update_psw = fmt.Sprintf("UPDATE schema_server.tb_password SET hashed_psw='%s' WHERE data_id='%d';", new_password, current_column_id)
			_, err_row_update_psw := db_site.Query(query_update_psw)
			Check_Error(err_row_update_psw)
		}
	} else {
	}
}

func Dashboard_Update_Server_Profile(dashboard_name string, dashboard_description string) {
	var query_update_dname = fmt.Sprintf("UPDATE schema_server.tb_main SET dashboard_name='%s' WHERE server_id='1';", dashboard_name)
	_, err_row_update_dname := db_server.Query(query_update_dname)
	Check_Error(err_row_update_dname)

	var query_update_ddescription = fmt.Sprintf("UPDATE schema_server.tb_main SET dashboard_description='%s' WHERE server_id='1';", dashboard_description)
	_, err_row_update_ddescription := db_server.Query(query_update_ddescription)
	Check_Error(err_row_update_ddescription)
}

func Dashboard_Create_New_Team_Member(user_name string, password string) {
	var query_tb_member = fmt.Sprintf("INSERT INTO schema_team.tb_main (user_name, hashed_psw, member_name, status_current_active) VALUES ('%s', '%s', '%s', '%s');", user_name, Hash_Password(password), "Unknown Name", "FALSE")
	_, err_member := db_server.Query(query_tb_member)
	Check_Error(err_member)
}

func Dashboard_Remove_Team_Member(user_name string) {
	var query_tb_delete = fmt.Sprintf("DELETE FROM schema_team.tb_main WHERE user_name='%s';", user_name)
	_, err_delete := db_server.Query(query_tb_delete)
	Check_Error(err_delete)

	fmt.Println("> Team Member Removed: ", user_name)
}

func Dashboard_Get_Member_Availability(id int) string {
	var query_get_member = fmt.Sprintf("SELECT user_name FROM schema_team.tb_main WHERE member_id='%d';", id)
	row_get_member, err_row_get_member := db_server.Query(query_get_member)
	Check_Error(err_row_get_member)

	if row_get_member.Next() != true {
		return ""
	}

	return "available"
}

func Dashboard_Get_Member_Name(id int) string {
	var query_get_member = fmt.Sprintf("SELECT member_name FROM schema_team.tb_main WHERE member_id='%d';", id)
	row_get_member, err_row_get_member := db_server.Query(query_get_member)
	Check_Error(err_row_get_member)

	if row_get_member.Next() != true {
		return ""
	}

	var current_data string
	row_get_member.Scan(&current_data)

	return current_data
}

func Dashboard_Get_Member_User_Name(id int) string {
	var query_get_member = fmt.Sprintf("SELECT user_name FROM schema_team.tb_main WHERE member_id='%d';", id)
	row_get_member, err_row_get_member := db_server.Query(query_get_member)
	Check_Error(err_row_get_member)

	if row_get_member.Next() != true {
		return ""
	}

	var current_data string
	row_get_member.Scan(&current_data)

	return current_data
}

func Dashboard_Create_New_Project(server_code string, project_name string) {
	project_code := strings.ToLower(strings.ReplaceAll(project_name, " ", ""))
	project_directory := fmt.Sprintf("/cdn/%s", project_code)

	// Create Directory and Base HTML, SCSS, CSS and JS Files

	mkdir_base := fmt.Sprintf("cdn/%s", server_code)

	mkdir_path := fmt.Sprintf("%s/%s", mkdir_base, project_code)
	mkdir_path_history_html := fmt.Sprintf("%s/%s", mkdir_path, "history_html")
	mkdir_path_history_js := fmt.Sprintf("%s/%s", mkdir_path, "history_js")
	mkdir_path_history_css := fmt.Sprintf("%s/%s", mkdir_path, "history_css")
	mkdir_path_history_scss := fmt.Sprintf("%s/%s", mkdir_path, "history_scss")

	mkfle_html := fmt.Sprintf("%s/index.html", mkdir_path)
	mkfle_js := fmt.Sprintf("%s/source_js.js", mkdir_path)
	mkfle_css := fmt.Sprintf("%s/source_css.css", mkdir_path)
	mkfle_scss := fmt.Sprintf("%s/source_scss.scss", mkdir_path)

	os.Mkdir(mkdir_base, os.ModePerm)
	os.Mkdir(mkdir_path, os.ModePerm)

	os.Create(mkfle_html)
	os.Create(mkfle_js)
	os.Create(mkfle_css)
	os.Create(mkfle_scss)

	// Create directories for save old files

	os.Mkdir(mkdir_path_history_html, os.ModePerm)
	os.Mkdir(mkdir_path_history_js, os.ModePerm)
	os.Mkdir(mkdir_path_history_css, os.ModePerm)
	os.Mkdir(mkdir_path_history_scss, os.ModePerm)

	var query_tb_project = fmt.Sprintf("INSERT INTO schema_project.tb_main (project_code, server_directory, project_title, project_description, project_icon, status_pinned) VALUES ('%s', '%s', '%s', '%s', '%s', '%s');", project_code, project_directory, project_name, "Unknown", "Unknown", "FALSE")
	_, err_project := db_server.Query(query_tb_project)
	Check_Error(err_project)

	fmt.Println("> New Project Created: ", project_name)
	fmt.Println("  + Directory : ", mkdir_path)
	fmt.Println("  + HTML File : ", mkfle_html)
	fmt.Println("  + JS File   : ", mkfle_js)
	fmt.Println("  + CSS File  : ", mkfle_css)
	fmt.Println("  + SCSS File : ", mkfle_scss)
}

func Dashboard_Get_Project_Availability(id int) string {
	var query_get_project = fmt.Sprintf("SELECT project_code FROM schema_project.tb_main WHERE project_id='%d';", id)
	row_get_project, err_row_get_project := db_server.Query(query_get_project)
	Check_Error(err_row_get_project)

	if row_get_project.Next() != true {
		return ""
	}

	return "available"
}

func Dashboard_Get_Project_Name(id int) string {
	var query_get_name = fmt.Sprintf("SELECT project_title FROM schema_project.tb_main WHERE project_id='%d';", id)
	row_get_name, err_row_get_name := db_server.Query(query_get_name)
	Check_Error(err_row_get_name)

	if row_get_name.Next() != true {
		return ""
	}

	var current_data string
	row_get_name.Scan(&current_data)

	return current_data
}

func Dashboard_Get_Project_Code(id int) string {
	var query_get_code = fmt.Sprintf("SELECT project_code FROM schema_project.tb_main WHERE project_id='%d';", id)
	row_get_code, err_row_get_code := db_server.Query(query_get_code)
	Check_Error(err_row_get_code)

	if row_get_code.Next() != true {
		return ""
	}

	var current_data string
	row_get_code.Scan(&current_data)

	return current_data
}