// Codinoc IDE
//
// Project by Group 35
// Computing Group Project
//
// Plymouth University UK
// NSBM Green University LK
//
// Sign Up Process

package main

import (
	"fmt"
	"strings"
	"database/sql"
	_ "github.com/lib/pq"
)

func SignUp_Add_New_Server(server_name string, admin_email string, admin_password string) {
	first_name := "Codinoc"
	middle_name := "Server"
	last_name := "Administrator"
	server_code := strings.ReplaceAll(server_name, " ", "")
	tenant_db := strings.ToLower(fmt.Sprintf("db_%s", server_code))

	query_tb_info := fmt.Sprintf("INSERT INTO schema_server.tb_info (first_name, middle_name, last_name) VALUES ('%s', '%s', '%s');", first_name, middle_name, last_name)
	query_tb_password := fmt.Sprintf("INSERT INTO schema_server.tb_password (hashed_psw) VALUES ('%s');", Hash_Password(admin_password))
	query_tb_security := fmt.Sprintf("INSERT INTO schema_server.tb_security (admin_email) VALUES ('%s');", admin_email)
	query_tb_server := fmt.Sprintf("INSERT INTO schema_server.tb_server (server_code, server_name) VALUES ('%s', '%s');", server_code, server_name)
	query_tb_status := fmt.Sprintf("INSERT INTO schema_server.tb_status (status_logged) VALUES (%s);", "FALSE")
	query_tb_tenant := fmt.Sprintf("INSERT INTO schema_server.tb_tenant (name_user_db) VALUES ('%s');", tenant_db)

	db_site.Query(query_tb_info)
	db_site.Query(query_tb_password)
	db_site.Query(query_tb_security)
	db_site.Query(query_tb_server)
	db_site.Query(query_tb_status)
	db_site.Query(query_tb_tenant)

	// https://go.dev/doc/database/querying
	// https://blog.logrocket.com/building-simple-app-go-postgresql/

	var query_get_current = fmt.Sprintf("SELECT data_id FROM schema_server.tb_tenant WHERE name_user_db='%s';", tenant_db)
	row_get_current, err_row_get_current := db_site.Query(query_get_current)
	Check_Error(err_row_get_current)

	if row_get_current.Next() != true {
		fmt.Println("> Unable to get \"data_id\" column id from \"schema_server.tb_tenant\"")
		return
	}

	var current_column_id string
	row_get_current.Scan(&current_column_id)

	var query_tb_main = fmt.Sprintf("INSERT INTO schema_server.tb_main (id_tenant, id_info, id_server, id_security, id_password, id_status) VALUES ('%s', '%s', '%s', '%s', '%s', '%s');", current_column_id, current_column_id, current_column_id, current_column_id, current_column_id, current_column_id)
	_, err_main := db_site.Query(query_tb_main)
	Check_Error(err_main)

	// Reset Table Values
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

	// Create tenant database
	//
	// https://stackoverflow.com/questions/55555836/is-it-possible-to-create-postgresql-databases-with-dynamic-names-with-the-help-o

	psql_string := fmt.Sprintf("user=postgres password=%s host=%s sslmode=disable", DB_PASS, DB_HOST)
	psql_conn, err_psql := sql.Open("postgres", psql_string)
	Check_Error(err_psql)

	query_psql_tenant := fmt.Sprintf("CREATE DATABASE %s WITH TEMPLATE %s OWNER admin_codinoc;", tenant_db, "db_server")
	_, err_psql_tenant := psql_conn.Exec(query_psql_tenant)
	Check_Error(err_psql_tenant)

	psql_conn.Close()

	// Update New Database Main Server Table

	db_server = Connect_Database(db_server, DB_HOST, DB_PORT, DB_USER, DB_PASS, tenant_db)

	query_new_update := fmt.Sprintf("UPDATE schema_server.tb_main SET server_code='%s', server_name='%s' WHERE server_id='1';", server_code, server_name)
	_, err_query_new_update := db_server.Query(query_new_update)
	Check_Error(err_query_new_update)

	db_server.Close()

	// FIXME Send Registation Successful Email

	// from_email := "no-reply@codinoc.com"
	// from_name := "Codinoc Team"
	// to_name := "Server Administrator"
	// mail_subject := "Welcome to Codinoc!"
	// mail_content := fmt.Sprintf("We\'re so happy you\'re here. Please use \"%s\" as your Server's code for Sign In", server_code)

	// Sent_Email(from_email, from_name, admin_email, to_name, mail_subject, mail_content)

	fmt.Println("> New server created successfully: " + server_name)
}