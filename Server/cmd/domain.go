// Codinoc IDE
//
// Project by Group 35
// Computing Group Project
//
// Plymouth University UK
// NSBM Green University LK
//
// Domain Routing

package main

import (
	"fmt"
	"strings"
	"net/http"
	"html/template"
)

func Check_Error(err error) {
	if err != nil {
			fmt.Println("> Error Found")
			panic(err)
	}
}

func Route_404(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := template.ParseFiles("thm/page/server/404.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
		return
	}

	parsedTemplate.Execute(w, nil)
}

func Route_Server(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := template.ParseFiles("thm/page/server/error.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, nil)
}

func Route_Maintenance(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := template.ParseFiles("thm/page/server/maintenance.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, nil)
}

func Route_Mobile(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := template.ParseFiles("thm/page/server/mobile.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, nil)
}

func Route_Home(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/home.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/navbar.tmpl",
		"thm/page/_part/non_logged/footer.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)

		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": "Welcome to Codinoc"})
}

func Route_About(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/about.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/navbar.tmpl",
		"thm/page/_part/non_logged/footer.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)

		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": "About Codinoc"})
}

func Route_Contact(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/contact.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/navbar.tmpl",
		"thm/page/_part/non_logged/footer.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)

		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": "Contact Codinoc"})
}

func Route_Legal(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/legal.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/navbar.tmpl",
		"thm/page/_part/non_logged/footer.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)

		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": "Terms and Conditions"})
}

func Route_Site_map(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/site_map.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/navbar.tmpl",
		"thm/page/_part/non_logged/footer.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)

		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": "Site Map"})
}

func Route_Account_Select(w http.ResponseWriter, r *http.Request) {
	// Redirect to Dashboard when User is logged

	session_login, _ := store_login.Get(r, "LOGIN_SESSION")

	if session_login != nil {
		if session_login.Values["session_string"] != nil {
			http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
			return
		}
	}

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/signin_account_select.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/footer_form.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)

		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": "Select Account"})
}

func Route_Admin_Signin(w http.ResponseWriter, r *http.Request) {
	// Redirect to Dashboard when User is logged

	session_login, _ := store_login.Get(r, "LOGIN_SESSION")

	if session_login != nil {
		if session_login.Values["session_string"] != nil {
			http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
			return
		}
	}

	error_server := r.URL.Query().Get("error")

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/signin_admin.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/footer_form.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)

		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": "Administrator Sign In", "Error_Server": error_server})
}

func Route_Member_Signin(w http.ResponseWriter, r *http.Request) {
	// Redirect to Dashboard when User is logged

	session_login, _ := store_login.Get(r, "LOGIN_SESSION")

	if session_login != nil {
		if session_login.Values["session_string"] != nil {
			http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
			return
		}
	}

	// ...
}

func Route_Create_Server(w http.ResponseWriter, r *http.Request) {
	// Redirect to Dashboard when User is logged

	session_login, _ := store_login.Get(r, "LOGIN_SESSION")

	if session_login != nil {
		if session_login.Values["session_string"] != nil {
			http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
			return
		}
	}

	error_server := r.URL.Query().Get("serror")

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/create_server.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/footer_form.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)

		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": "Create Server", "Server_Error": error_server})
}

func Route_Create_Admin(w http.ResponseWriter, r *http.Request) {
	// Redirect to Dashboard when User is logged

	session_login, _ := store_login.Get(r, "LOGIN_SESSION")

	if session_login != nil {
		if session_login.Values["session_string"] != nil {
			http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
			return
		}
	}

	server_name := r.URL.Query().Get("sname")
	creation_error := r.URL.Query().Get("error")

	if server_name == "" {
		http.Redirect(w, r, fmt.Sprintf("/404"), 302)
		return
	}

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/create_admin.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/footer_form.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)

		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": "Create Administrator Account", "Server_Name": server_name, "Error_Message": creation_error})
}

func Route_Create_Admin_Success(w http.ResponseWriter, r *http.Request) {
	// Redirect to Dashboard when User is logged

	session_login, _ := store_login.Get(r, "LOGIN_SESSION")

	if session_login != nil {
		if session_login.Values["session_string"] != nil {
			http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
			return
		}
	}

	server_name := r.URL.Query().Get("name")
	server_code := r.URL.Query().Get("code")
	server_email := r.URL.Query().Get("email")

	if server_name == "" || server_code == "" || server_email == "" {
		http.Redirect(w, r, fmt.Sprintf("/404"), 302)
		return
	}

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/create_admin_success.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/footer_form.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)

		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": "Congratulations !", "Server_Name": server_name, "Server_Code": server_code, "Server_Email": server_email})
}

func Route_Reset_Password(w http.ResponseWriter, r *http.Request) {
	// Redirect to Dashboard when User is logged

	session_login, _ := store_login.Get(r, "LOGIN_SESSION")

	if session_login != nil {
		if session_login.Values["session_string"] != nil {
			http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
			return
		}
	}

	error_server := r.URL.Query().Get("serror")
	error_email := r.URL.Query().Get("merror")

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/reset_password.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/footer_form.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)

		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": "Reset Password", "Server_Error": error_server, "Email_Error": error_email})
}

func Route_Reset_sent(w http.ResponseWriter, r *http.Request) {
	// Redirect to Dashboard when User is logged

	session_login, _ := store_login.Get(r, "LOGIN_SESSION")

	if session_login != nil {
		if session_login.Values["session_string"] != nil {
			http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
			return
		}
	}

	server_code := r.URL.Query().Get("scode")
	admin_email := r.URL.Query().Get("semail")

	if server_code == "" || admin_email == "" {
		http.Redirect(w, r, fmt.Sprintf("/404"), 302)
		return
	}

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/reset_sent.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/footer_form.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)

		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": "Instructions are sent", "Server_Code": server_code, "Server_Email": admin_email})
}

func Route_Reset_New(w http.ResponseWriter, r *http.Request) {
	// Redirect to Dashboard when User is logged

	session_login, _ := store_login.Get(r, "LOGIN_SESSION")

	if session_login != nil {
		if session_login.Values["session_string"] != nil {
			http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
			return
		}
	}

	server_code := r.URL.Query().Get("scode")
	admin_email := r.URL.Query().Get("semail")

	if Process_Is_Server_Available_From_Code(server_code) == false {
		http.Redirect(w, r, fmt.Sprintf("/404"), 302)
		return
	}

	if Process_Is_Admin_Available_In_Server(server_code, admin_email) == false {
		http.Redirect(w, r, fmt.Sprintf("/404"), 302)
		return
	}

	password_error := r.URL.Query().Get("perror")

	if server_code == "" || admin_email == "" {
		http.Redirect(w, r, fmt.Sprintf("/reset_failed"), 302)
		return
	}

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/reset_new.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/footer_form.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)

		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": "Setup New Password", "Server_Code": server_code, "Server_Email": admin_email, "Password_Error": password_error})
}

func Route_Reset_Failed(w http.ResponseWriter, r *http.Request) {
	// Redirect to Dashboard when User is logged

	session_login, _ := store_login.Get(r, "LOGIN_SESSION")

	if session_login != nil {
		if session_login.Values["session_string"] != nil {
			http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
			return
		}
	}

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/reset_failed.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/footer_form.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)

		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": "Password Reset Failed"})
}

func Route_Reset_Success(w http.ResponseWriter, r *http.Request) {
	// Redirect to Dashboard when User is logged

	session_login, _ := store_login.Get(r, "LOGIN_SESSION")

	if session_login != nil {
		if session_login.Values["session_string"] != nil {
			http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
			return
		}
	}

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/reset_success.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/footer_form.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)

		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": "Reset Successful"})
}

func Route_Dashboard(w http.ResponseWriter, r *http.Request) {
	// Redirect to Sign In Selection if User not Logged

	session_login, err := store_login.Get(r, "LOGIN_SESSION")

	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/account_select"), 302)
		return
	}

	if session_login != nil {
		if session_login.Values["session_type"] == nil {
			http.Redirect(w, r, fmt.Sprintf("/account_select"), 302)
			return
		}
	}

	var session_type = fmt.Sprintf("%v", session_login.Values["session_type"])
	var session_code = fmt.Sprintf("%v", session_login.Values["session_code"])

	parsedTemplate, err := template.ParseFiles(
		"thm/page/logged/dashboard.tmpl",
		"thm/page/_part/logged/head.tmpl",
		"thm/page/_part/logged/bottom.tmpl",
		"thm/page/_part/logged/navbar.tmpl",
		"thm/page/_part/logged/footer.tmpl",
		"thm/page/_part/logged/panel_team.tmpl",
		"thm/page/_part/logged/panel_chat.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)

		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	}

	Get_Navbar_Name := Dashboard_Get_Server_Name()
	Get_Dashboard_Name := Dashboard_Get_Dashboard_Name()
	Get_Dashboard_Description := Dashboard_Get_Dashboard_Description()
	Get_Wiki_Content := Dashboard_Get_WikiContent()
	Get_Account_Type := Dashboard_Get_Account_Type(session_type)
	Get_First_Name := Dashboard_Get_First_Name(session_type, session_code)
	Get_Middle_Name := Dashboard_Get_Middle_Name(session_type, session_code)
	Get_Last_Name := Dashboard_Get_Last_Name(session_type, session_code)
	Get_User_Name := Dashboard_Get_User_Name(session_type, session_code)
	Get_Server_Code := Dashboard_Get_Server_Code()
	Get_Server_Name := Dashboard_Get_Server_Name()

	parsedTemplate.Execute(w, map[string]string{
		"Page_Title": "Dashboard",
		"Navbar_Name": Get_Navbar_Name,
		"Dashboard_Name": Get_Dashboard_Name,
		"Dashboard_Description": Get_Dashboard_Description,
		"Wiki_Content": Get_Wiki_Content,
		"Account_Type": Get_Account_Type,
		"First_Name": Get_First_Name,
		"Middle_Name": Get_Middle_Name,
		"Last_Name": Get_Last_Name,
		"User_Name": Get_User_Name,
		"Server_Code": Get_Server_Code,
		"Server_Name": Get_Server_Name})
}

func Route_Editor(w http.ResponseWriter, r *http.Request) {
	// Redirect to Sign In Selection if User not Logged

	session_login, err := store_login.Get(r, "LOGIN_SESSION")

	if err != nil {
		http.Redirect(w, r, fmt.Sprintf("/account_select"), 302)
		return
	}

	if session_login != nil {
		if session_login.Values["session_type"] == nil {
			http.Redirect(w, r, fmt.Sprintf("/account_select"), 302)
			return
		}
	}

	// ...
}

func Route_Account_Sign_Out(w http.ResponseWriter, r *http.Request) {
	SignIn_Session_Disable(w, r)
	http.Redirect(w, r, fmt.Sprintf("/home"), 302)
}

func Route_Process_SignIn_Admin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	} else {
		parse_error := r.ParseForm()

		if parse_error != nil {
			fmt.Println(parse_error)

			http.Redirect(w, r, fmt.Sprintf("/server"), 302)
			return
		}

		server_code := r.FormValue("input_scode")
		admin_email := r.FormValue("input_email")
		admin_password := r.FormValue("input_psw")

		if Process_Is_Server_Available_From_Code(server_code) == false {
			http.Redirect(w, r, fmt.Sprintf("/admin_signin?error=%s", "No Server Available"), 302)
			return
		}

		if Process_Is_Admin_Available_In_Server(server_code, admin_email) == false {
			http.Redirect(w, r, fmt.Sprintf("/admin_signin?error=%s", "Invalid Administrator Email Address"), 302)
			return
		}

		if Process_Is_Admin_Password_Match(server_code, admin_email, admin_password) == false {
			http.Redirect(w, r, fmt.Sprintf("/admin_signin?error=%s", "Password does not match"), 302)
			return
		}

		// Sign In Session Store Process

		SignIn_Session_Enable(w, r, admin_email, "admin", server_code)

		// Connect Server Database

		SignIn_Database_Connect(server_code)

		// fmt.Println("> Current Session Details:")
		// fmt.Println(" - String -> '", session_login.Values["session_string"], "'")
		// fmt.Println(" - Type   -> '", session_login.Values["session_type"], "'")
		// fmt.Println(" - Code   -> '", session_login.Values["session_code"], "'")

		http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
	}
}

func Route_Process_SignIn_Member(w http.ResponseWriter, r *http.Request) {
	// ...
}

func Route_Process_Create_Server(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	} else {
		parse_error := r.ParseForm()

		if parse_error != nil {
			fmt.Println(parse_error)

			http.Redirect(w, r, fmt.Sprintf("/server"), 302)
			return
		}

		server_name := r.FormValue("input_sname")

		if Process_Is_Server_Available(server_name) == true {
			http.Redirect(w, r, fmt.Sprintf("/create_server?serror=%s", "Server Already Available"), 302)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/create_admin?sname=%s", server_name), 302)
	}
}

func Route_Process_Create_Admin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	} else {
		parse_error := r.ParseForm()

		if parse_error != nil {
			fmt.Println(parse_error)

			http.Redirect(w, r, fmt.Sprintf("/server"), 302)
			return
		}

		server_code := strings.ReplaceAll(r.FormValue("server_name"), " ", "")
		server_name := r.FormValue("server_name")
		admin_email := r.FormValue("admin_email")
		admin_password := r.FormValue("admin_password")
		admin_password_confirmed := r.FormValue("admin_password_confirmed")

		if admin_password != admin_password_confirmed {
			http.Redirect(w, r, fmt.Sprintf("/create_admin?sname=%s&error=%s", server_name, "Password Does not match"), 302)
			return
		}

		fmt.Println("> Creating Administrator Account for: " + server_code + " - " + server_name + " - " + admin_email + " - " + admin_password + " - " + admin_password_confirmed)

		SignUp_Add_New_Server(server_name, admin_email, admin_password)
		http.Redirect(w, r, fmt.Sprintf("/create_admin_success?name=%s&code=%s&email=%s", server_name, server_code, admin_email), 302)
	}
}

func Route_Process_Reset_Password(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	} else {
		parse_error := r.ParseForm()

		if parse_error != nil {
			fmt.Println(parse_error)

			http.Redirect(w, r, fmt.Sprintf("/server"), 302)
			return
		}

		server_code := strings.ToLower(strings.ReplaceAll(r.FormValue("input_scode"), " ", ""))
		admin_email := r.FormValue("input_email")

		if Process_Is_Server_Available_From_Code(server_code) == false {
			http.Redirect(w, r, fmt.Sprintf("/reset_password?serror=%s", "No Server Found"), 302)
			return
		}

		if Process_Is_Admin_Available_In_Server(server_code, admin_email) == false {
			http.Redirect(w, r, fmt.Sprintf("/reset_password?merror=%s", "No Email found for given Server"), 302)
			return
		}

		// FIXME Forgot Password Link Email

		// from_email := "no-reply@codinoc.com"
		// from_name := "Codinoc Team"
		// to_name := "Server Administrator"
		// mail_subject := "Server Password Reset"
		// mail_content := fmt.Sprintf("Follow this link to reset your server's password: http://127.0.0.1:8080/reset_new?scode=%s&semail=%s", server_code, admin_email)

		// Sent_Email(from_email, from_name, admin_email, to_name, mail_subject, mail_content)

		http.Redirect(w, r, fmt.Sprintf("/reset_sent?scode=%s&semail=%s", server_code, admin_email), 302)
	}
}

func Route_Process_New_Password(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	} else {
		parse_error := r.ParseForm()

		if parse_error != nil {
			fmt.Println(parse_error)

			http.Redirect(w, r, fmt.Sprintf("/server"), 302)
			return
		}

		server_code := r.FormValue("server_code")
		server_email := r.FormValue("server_email")
		new_password := r.FormValue("password")
		new_password_conf := r.FormValue("password_conf")

		if new_password != new_password_conf {
			http.Redirect(w, r, fmt.Sprintf("/reset_new?scode=%s%semail=%s&perror=%s", server_code, server_email, "Password Does not match"), 302)
			return
		}

		Forgot_Update_Password(server_code, new_password)
		http.Redirect(w, r, fmt.Sprintf("/reset_success"), 302)
	}
}

func Route_Process_Update_Wiki(w http.ResponseWriter, r *http.Request) {
	// ...
}

func Route_Process_Update_User(w http.ResponseWriter, r *http.Request) {
	// ...
}

func Route_Process_Update_Server(w http.ResponseWriter, r *http.Request) {
	// ...
}

func Route_Process_Add_Team_Member(w http.ResponseWriter, r *http.Request) {
	// ...
}

func Route_Process_Remove_Team_Member(w http.ResponseWriter, r *http.Request) {
	// ...
}

func Route_Process_Create_Project(w http.ResponseWriter, r *http.Request) {
	// ...
}

func Route_Process_Remove_Project(w http.ResponseWriter, r *http.Request) {
	// ...
}