// Codinoc IDE
//
// Project by Group 35
// Computing Group Project
//
// Plymouth University UK
// NSBM Green University LK
//
// Sign Up Functions

package main

import (
	"fmt"
	"strings"
	"database/sql"
	_ "github.com/lib/pq"
)

// https://stackoverflow.com/questions/61726581/how-to-retrieve-values-in-golang-database-sql-when-my-structure-in-dbhere-post

func Create_ServerCode(cur_code string) string {
	var updated_code = strings.ReplaceAll(cur_code, " ", "")
	return updated_code
}

// Create new server account

func Create_Server(server_name string, server_code string, admin_email string, admin_password string) {
	var first_name = "Codinoc"
	var middle_name = "Administrative"
	var last_name = "User"
	var gender = "Unknown"
	var mobile_number = "123456789"
	var home_address = "Unknown"
	var tenant_db = strings.ToLower(fmt.Sprintf("db_%s", Create_ServerCode(server_name)))

	var query_tb_info = fmt.Sprintf("INSERT INTO schema_server.tb_info (first_name, middle_name, last_name, gender, mobile_number, home_address) VALUES ('%s', '%s', '%s', '%s', '%s', '%s');", first_name, middle_name, last_name, gender, mobile_number, home_address)
	_, err_info := db_site.Query(query_tb_info)
	CheckError(err_info)

	var query_tb_password = fmt.Sprintf("INSERT INTO schema_server.tb_password (hashed_psw) VALUES ('%s');", admin_password)
	_, err_password := db_site.Query(query_tb_password)
	CheckError(err_password)

	var query_tb_security = fmt.Sprintf("INSERT INTO schema_server.tb_security (admin_email) VALUES ('%s');", admin_email)
	_, err_security := db_site.Query(query_tb_security)
	CheckError(err_security)

	var query_tb_server = fmt.Sprintf("INSERT INTO schema_server.tb_server (server_code, server_name) VALUES ('%s', '%s');", server_code, server_name)
	_, err_server := db_site.Query(query_tb_server)
	CheckError(err_server)

	var query_tb_status = fmt.Sprintf("INSERT INTO schema_server.tb_status (status_logged) VALUES (%s);", "FALSE")
	_, err_status := db_site.Query(query_tb_status)
	CheckError(err_status)

	var query_tb_tenant = fmt.Sprintf("INSERT INTO schema_server.tb_tenant (name_user_db) VALUES ('%s');", tenant_db)
	_, err_tenant := db_site.Query(query_tb_tenant)
	CheckError(err_tenant)

	// https://go.dev/doc/database/querying
	// https://blog.logrocket.com/building-simple-app-go-postgresql/

	var query_get_current = fmt.Sprintf("SELECT data_id FROM schema_server.tb_tenant WHERE name_user_db='%s';", tenant_db)
	row_get_current, err_row_get_current := db_site.Query(query_get_current)
	CheckError(err_row_get_current)

	if row_get_current.Next() != true {
		fmt.Println("> Unable to get column id")
		return
	}

	var current_column_id string
	row_get_current.Scan(&current_column_id)

	var query_tb_main = fmt.Sprintf("INSERT INTO schema_server.tb_main (id_tenant, id_info, id_server, id_security, id_password, id_status) VALUES ('%s', '%s', '%s', '%s', '%s', '%s');", current_column_id, current_column_id, current_column_id, current_column_id, current_column_id, current_column_id)
	_, err_main := db_site.Query(query_tb_main)
	CheckError(err_main)

	// Reset Values
	//
	// https://stackoverflow.com/questions/16362930/change-the-starting-value-of-a-serial-postgresql
	//
	// ALTER SEQUENCE schema_server.tb_main_user_id_seq RESTART WITH 1;
	// ALTER SEQUENCE schema_server.tb_info_data_id_seq RESTART WITH 1;
	// ALTER SEQUENCE schema_server.tb_password_data_id_seq RESTART WITH 1;
	// ALTER SEQUENCE schema_server.tb_security_data_id_seq RESTART WITH 1;
	// ALTER SEQUENCE schema_server.tb_server_data_id_seq RESTART WITH 1;
	// ALTER SEQUENCE schema_server.tb_status_data_id_seq RESTART WITH 1;
	// ALTER SEQUENCE schema_server.tb_tenant_data_id_seq RESTART WITH 1;

	// Check Queries
	//
	// fmt.Println(query_tb_info)
	// fmt.Println(query_tb_password)
	// fmt.Println(query_tb_security)
	// fmt.Println(query_tb_server)
	// fmt.Println(query_tb_status)
	// fmt.Println(query_tb_tenant)
	// fmt.Println(query_get_current)
	// fmt.Println(query_tb_main)

	// Create tenant database
	//
	// https://stackoverflow.com/questions/55555836/is-it-possible-to-create-postgresql-databases-with-dynamic-names-with-the-help-o

	postgre_conn_info := fmt.Sprintf("user=postgres password=%s host=%s sslmode=disable", DB_SITE_PASSWORD, DB_SITE_HOST)
	postgres_db, err_postgre := sql.Open("postgres", postgre_conn_info)
	CheckError(err_postgre)

	query_tenant_db := fmt.Sprintf("CREATE DATABASE %s WITH TEMPLATE %s OWNER admin_codinoc;", tenant_db, "db_server")
	_, err_tenant_create := postgres_db.Exec(query_tenant_db)
	CheckError(err_tenant_create)
}