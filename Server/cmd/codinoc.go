// Codinoc IDE
// Main Go File
//
// Written by N.L.M. Nishshake
//            M.P.M. Abeyrathne

package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"

	C_Conn "./connection"
	C_DB "./database"
	C_Route "./route"
)

func main() {
	// Database Connection

	C_DB.Connect_Site_DB()

	// Domain Routing

	router := mux.NewRouter()
	router.PathPrefix("/resources/").Handler(http.StripPrefix("/resources/", http.FileServer(http.Dir("thm/"))))

	// Non Logged Section

	router.HandleFunc("/", C_Route.Non_logged_home).Methods("GET")
	router.HandleFunc("/home", C_Route.Non_logged_home).Methods("GET")
	router.HandleFunc("/about", C_Route.Non_logged_about).Methods("GET")
	router.HandleFunc("/contact", C_Route.Non_logged_contact).Methods("GET")
	router.HandleFunc("/legal", C_Route.Non_logged_legal).Methods("GET")
	router.HandleFunc("/site_map", C_Route.Non_logged_site_map).Methods("GET")

	router.HandleFunc("/account_select", C_Route.Non_logged_account_select).Methods("GET")
	router.HandleFunc("/admin_signin", C_Route.Non_logged_admin_signin).Methods("GET")
	router.HandleFunc("/member_signin", C_Route.Non_logged_member_signin).Methods("GET")

	router.HandleFunc("/create_server", C_Route.Non_logged_create_server).Methods("GET")
	router.HandleFunc("/create_admin", C_Route.Non_logged_create_admin).Methods("GET")
	router.HandleFunc("/create_admin_success", C_Route.Non_logged_create_admin_success).Methods("GET")

	router.HandleFunc("/reset_password", C_Route.Non_logged_reset_password).Methods("GET")
	router.HandleFunc("/reset_sent", C_Route.Non_logged_reset_sent).Methods("GET")
	router.HandleFunc("/reset_new", C_Route.Non_logged_reset_new).Methods("GET")
	router.HandleFunc("/reset_failed", C_Route.Non_logged_reset_failed).Methods("GET")
	router.HandleFunc("/reset_success", C_Route.Non_logged_reset_success).Methods("GET")

	// Logged Section

	router.HandleFunc("/dashboard", C_Route.Logged_dashboard).Methods("GET")
	router.HandleFunc("/editor", C_Route.Logged_editor).Methods("GET")

	err := http.ListenAndServe(C_Conn.CONN_HOST+":"+C_Conn.CONN_PORT, router)

	if err != nil {
		log.Fatal("Unable to start codinoc server : ", err)
		return
	}
}
