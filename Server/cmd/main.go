// Codinoc IDE
//
// Project by Group 35
// Computing Group Project
//
// Plymouth University UK
// NSBM Green University LK
//
// Starting Function

package main

import (
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {
	db_site = Connect_Database(db_site, DB_HOST, DB_PORT, DB_USER, DB_PASS, "db_site")

	router := mux.NewRouter()
	router.PathPrefix("/resources/").Handler(http.StripPrefix("/resources/", http.FileServer(http.Dir("thm/"))))

	router.HandleFunc("/", Route_404).Methods("GET")
	router.HandleFunc("/404", Route_404).Methods("GET")
	router.HandleFunc("/server", Route_Server).Methods("GET")
	router.HandleFunc("/maintenance", Route_Maintenance).Methods("GET")
	router.HandleFunc("/mobile", Route_Mobile).Methods("GET")

	router.HandleFunc("/home", Route_Home).Methods("GET")
	router.HandleFunc("/about", Route_About).Methods("GET")
	router.HandleFunc("/contact", Route_Contact).Methods("GET")
	router.HandleFunc("/legal", Route_Legal).Methods("GET")
	router.HandleFunc("/site_map", Route_Site_map).Methods("GET")

	router.HandleFunc("/account_select", Route_Account_Select).Methods("GET")
	router.HandleFunc("/admin_signin", Route_Admin_Signin).Methods("GET")
	router.HandleFunc("/member_signin", Route_Member_Signin).Methods("GET")

	router.HandleFunc("/create_server", Route_Create_Server).Methods("GET")
	router.HandleFunc("/create_admin", Route_Create_Admin).Methods("GET")
	router.HandleFunc("/create_admin_success", Route_Create_Admin_Success).Methods("GET")

	router.HandleFunc("/reset_password", Route_Reset_Password).Methods("GET")
	router.HandleFunc("/reset_sent", Route_Reset_sent).Methods("GET")
	router.HandleFunc("/reset_new", Route_Reset_New).Methods("GET")
	router.HandleFunc("/reset_failed", Route_Reset_Failed).Methods("GET")
	router.HandleFunc("/reset_success", Route_Reset_Success).Methods("GET")

	router.HandleFunc("/dashboard", Route_Dashboard).Methods("GET")
	router.HandleFunc("/editor", Route_Editor).Methods("GET")

	router.HandleFunc("/process/common/signin_admin", Route_Process_SignIn_Admin).Methods("GET", "POST")
	router.HandleFunc("/process/common/signin_member", Route_Process_SignIn_Member).Methods("GET", "POST")
	router.HandleFunc("/process/common/create_server", Route_Process_Create_Server).Methods("GET", "POST")
	router.HandleFunc("/process/common/create_admin", Route_Process_Create_Admin).Methods("GET", "POST")
	router.HandleFunc("/process/common/reset_password", Route_Process_Reset_Password).Methods("GET", "POST")
	router.HandleFunc("/process/common/new_password", Route_Process_New_Password).Methods("GET", "POST")
	router.HandleFunc("/process/dashboard/update_wiki", Route_Process_Update_Wiki).Methods("GET", "POST")
	router.HandleFunc("/process/dashboard/update_user", Route_Process_Update_User).Methods("GET", "POST")
	router.HandleFunc("/process/dashboard/update_server", Route_Process_Update_Server).Methods("GET", "POST")
	router.HandleFunc("/process/dashboard/add_team_member", Route_Process_Add_Team_Member).Methods("GET", "POST")
	router.HandleFunc("/process/dashboard/remove_team_member", Route_Process_Remove_Team_Member).Methods("GET", "POST")
	router.HandleFunc("/process/dashboard/create_project", Route_Process_Create_Project).Methods("GET", "POST")
	router.HandleFunc("/process/dashboard/remove_project", Route_Process_Remove_Project).Methods("GET", "POST")

	fmt.Println("> To open Codinoc Server, URL is http://" + CONN_HOST + ":" + CONN_PORT)
	router_error := http.ListenAndServe(CONN_HOST + ":" + CONN_PORT, router)

	if router_error != nil {
		fmt.Println("> Unable to start codinoc server : ", router_error)
		return
	}
}