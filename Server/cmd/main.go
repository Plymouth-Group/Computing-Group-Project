// Codinoc IDE
//
// Project by Group 35
// Computing Group Project
//
// Plymouth University UK
// NSBM Green University LK
//
// Custom Services
// - Email Send: https://sendgrid.com/
//
// Starting File

package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// Main Database Connection

	site_db_connect()

	// Domain Routing

	router := mux.NewRouter()
	router.PathPrefix("/resources/").Handler(http.StripPrefix("/resources/", http.FileServer(http.Dir("thm/"))))

	// Non Logged Section

	router.HandleFunc("/", Non_logged_home).Methods("GET")
	router.HandleFunc("/home", Non_logged_home).Methods("GET")
	router.HandleFunc("/about", Non_logged_about).Methods("GET")
	router.HandleFunc("/contact", Non_logged_contact).Methods("GET")
	router.HandleFunc("/legal", Non_logged_legal).Methods("GET")
	router.HandleFunc("/site_map", Non_logged_site_map).Methods("GET")

	router.HandleFunc("/account_select", Non_logged_account_select).Methods("GET")
	router.HandleFunc("/admin_signin", Non_logged_admin_signin).Methods("GET")
	router.HandleFunc("/member_signin", Non_logged_member_signin).Methods("GET")

	router.HandleFunc("/create_server", Non_logged_create_server).Methods("GET")
	router.HandleFunc("/create_admin", Non_logged_create_admin).Methods("GET")
	router.HandleFunc("/create_admin_success", Non_logged_create_admin_success).Methods("GET")

	router.HandleFunc("/reset_password", Non_logged_reset_password).Methods("GET")
	router.HandleFunc("/reset_sent", Non_logged_reset_sent).Methods("GET")
	router.HandleFunc("/reset_new", Non_logged_reset_new).Methods("GET")
	router.HandleFunc("/reset_failed", Non_logged_reset_failed).Methods("GET")
	router.HandleFunc("/reset_success", Non_logged_reset_success).Methods("GET")

	// Logged Section

	router.HandleFunc("/dashboard", Logged_dashboard).Methods("GET")
	router.HandleFunc("/editor", Logged_editor).Methods("GET")

	fmt.Println("> To open Codinoc Server, URL is http://" + CONN_HOST + ":" + CONN_PORT)

	err := http.ListenAndServe(CONN_HOST + ":" + CONN_PORT, router)

	if err != nil {
		fmt.Println("> Unable to start codinoc server : ", err)
		return
	}
}