// Codinoc IDE
//
// Project by Group 35
// Computing Group Project
//
// Plymouth University UK
// NSBM Green University LK
//
// Domain Routing Structure

package main

import (
	"fmt"
	"net/http"
	"html/template"
	"strings"
)

// Server Section Domain Routing

func Server_404(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := template.ParseFiles("thm/page/server/404.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, nil)
}

func Server_error(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := template.ParseFiles("thm/page/server/error.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, nil)
}

func Server_maintenance(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := template.ParseFiles("thm/page/server/maintenance.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, nil)
}

func Server_mobile(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := template.ParseFiles("thm/page/server/mobile.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, nil)
}

// Non Logged Section Domain Routing

func Non_logged_home(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Welcome to Codinoc"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/home.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/navbar.tmpl",
		"thm/page/_part/non_logged/footer.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_about(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "About Codinoc"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/about.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/navbar.tmpl",
		"thm/page/_part/non_logged/footer.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_contact(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Contact Codinoc"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/contact.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/navbar.tmpl",
		"thm/page/_part/non_logged/footer.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_legal(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Legal Notes"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/legal.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/navbar.tmpl",
		"thm/page/_part/non_logged/footer.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_site_map(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Site Map"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/site_map.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/navbar.tmpl",
		"thm/page/_part/non_logged/footer.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_account_select(w http.ResponseWriter, r *http.Request) {
	// redirect to dashboard if user is already logged

	if glob_sign_in.is_logged == true {
		http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
	}

	var Page_Title = "Select account for Sign In"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/signin_account_select.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/footer_form.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_admin_signin(w http.ResponseWriter, r *http.Request) {
	// redirect to dashboard if user is already logged

	if glob_sign_in.is_logged == true {
		http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
	}

	if r.Method == "GET" {
		error_server := r.URL.Query().Get("serror")

		var Page_Title = "Sign in As Administrator"

		parsedTemplate, err := template.ParseFiles(
			"thm/page/non_logged/signin_admin.tmpl",
			"thm/page/_part/non_logged/head.tmpl",
			"thm/page/_part/non_logged/bottom.tmpl",
			"thm/page/_part/non_logged/footer_form.tmpl")

		if err != nil {
			fmt.Println("> Unable to parse html file : ", err)
		}

		parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title, "Error_Server": error_server})
	} else {
		parse_error := r.ParseForm()

		if parse_error != nil {
			fmt.Println(parse_error)
			return
		}

		server_code := r.FormValue("input_scode")
		admin_email := r.FormValue("input_email")
		admin_password := r.FormValue("input_psw")

		if Signin_admin_check(server_code, admin_email, admin_password) == false {
			http.Redirect(w, r, fmt.Sprintf("/admin_signin?serror=%s", "Check Login Details Again"), 302)
			return
		}

		// Store session

		// sign_in_session, _ := sign_in_store.Get(r, "login_session")
		// sign_in_session.Values["email"] = admin_email
		// sign_in_session.Values["password"] = admin_password
		// sign_in_session.Save(r, w)

		// Connect server database

		server_db_connect(fmt.Sprintf("db_%s", server_code))

		glob_sign_in.is_logged = true
		http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
	}
}

func Non_logged_member_signin(w http.ResponseWriter, r *http.Request) {
	// redirect to dashboard if user is already logged

	if glob_sign_in.is_logged == true {
		http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
	}

	var Page_Title = "Sign in As Member"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/signin_member.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/footer_form.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_create_server(w http.ResponseWriter, r *http.Request) {
	// redirect to dashboard if user is already logged

	if glob_sign_in.is_logged == true {
		http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
	}

	if r.Method == "GET" {
		error_server := r.URL.Query().Get("serror")

		var Page_Title = "Create Server"

		parsedTemplate, err := template.ParseFiles(
			"thm/page/non_logged/create_server.tmpl",
			"thm/page/_part/non_logged/head.tmpl",
			"thm/page/_part/non_logged/bottom.tmpl",
			"thm/page/_part/non_logged/footer_form.tmpl")

		if err != nil {
			fmt.Println("> Unable to parse html file : ", err)
		}

		parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title, "Server_Error": error_server})
	} else {
		parse_error := r.ParseForm()

		if parse_error != nil {
			fmt.Println(parse_error)
			return
		}

		glob_sign_up.server_name = r.FormValue("input_sname")

		// Check server is already available or not

		var Server_Code = strings.ToLower(Create_ServerCode(glob_sign_up.server_name))

		if Check_Server(Server_Code) == false {
			http.Redirect(w, r, fmt.Sprintf("/create_server?serror=%s", "Server Already Available"), 302)
			return
		}

		http.Redirect(w, r, fmt.Sprintf("/create_admin?server_name=%s", glob_sign_up.server_name), 302)
	}
}

func Non_logged_create_admin(w http.ResponseWriter, r *http.Request) {
	// redirect to dashboard if user is already logged

	if glob_sign_in.is_logged == true {
		http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
	}

	if r.Method == "GET" {
		server_name := r.URL.Query().Get("server_name")

		if server_name == "" {
			http.Redirect(w, r, fmt.Sprintf("/404"), 302)
			return
		}

		var Page_Title = "Create Server Administrator"

		parsedTemplate, err := template.ParseFiles(
			"thm/page/non_logged/create_admin.tmpl",
			"thm/page/_part/non_logged/head.tmpl",
			"thm/page/_part/non_logged/bottom.tmpl",
			"thm/page/_part/non_logged/footer_form.tmpl")

		if err != nil {
			fmt.Println("> Unable to parse html file : ", err)
		}

		parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title, "Server_name": server_name})
	} else {
		parse_error := r.ParseForm()

		if parse_error != nil {
			fmt.Println(parse_error)
			return
		}

		glob_sign_up.admin_email = r.FormValue("admin_email")

		// Password hashing

		glob_sign_up.admin_password = Hash_Password(r.FormValue("admin_password"))

		glob_sign_up.is_account_created = true
		http.Redirect(w, r, fmt.Sprintf("/create_admin_success"), 302)
	}
}

func Non_logged_create_admin_success(w http.ResponseWriter, r *http.Request) {
	// redirect to dashboard if user is already logged

	if glob_sign_in.is_logged == true {
		http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
	}

	if glob_sign_up.is_account_created == false {
		http.Redirect(w, r, fmt.Sprintf("/404"), 302)
		return
	}

	var Page_Title = "Congratulations"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/create_admin_success.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/footer_form.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
	}

	var Server_Name = glob_sign_up.server_name
	var Server_Code = strings.ToLower(Create_ServerCode(Server_Name))

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title, "Server_Name": Server_Name, "Server_Code": Server_Code})

	glob_sign_up.server_code = Server_Code
	Create_Server(glob_sign_up.server_name,glob_sign_up.server_code,glob_sign_up.admin_email,glob_sign_up.admin_password)

	// TODO Send Congratulation email
	// main_server_created(glob_sign_up.admin_email, glob_sign_up.server_name)

	glob_sign_up.is_account_created = false
}

func Non_logged_reset_password(w http.ResponseWriter, r *http.Request) {
	// redirect to dashboard if user is already logged

	if glob_sign_in.is_logged == true {
		http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
	}

	if r.Method == "GET" {
		error_server := r.URL.Query().Get("serror")
		error_email := r.URL.Query().Get("merror")

		var Page_Title = "Reset Password"

		parsedTemplate, err := template.ParseFiles(
			"thm/page/non_logged/reset_password.tmpl",
			"thm/page/_part/non_logged/head.tmpl",
			"thm/page/_part/non_logged/bottom.tmpl",
			"thm/page/_part/non_logged/footer_form.tmpl")

		if err != nil {
			fmt.Println("> Unable to parse html file : ", err)
		}

		parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title, "ServerError": error_server, "EmailError": error_email})
	} else {
		parse_error := r.ParseForm()

		if parse_error != nil {
			fmt.Println(parse_error)
			return
		}

		glob_forgot_password.server_code = strings.ToLower(Create_ServerCode(r.FormValue("input_scode")))
		glob_forgot_password.admin_email = r.FormValue("input_email")

		if CheckServerfor_PasswordReset(glob_forgot_password.server_code) == false {
			http.Redirect(w, r, fmt.Sprintf("/reset_password?serror=%s", "No Server Found"), 302)
			return
		}

		glob_forgot_password.is_password_reset_available = true
		http.Redirect(w, r, fmt.Sprintf("/reset_sent?server_email=%s", glob_forgot_password.admin_email), 302)
	}
}

func Non_logged_reset_sent(w http.ResponseWriter, r *http.Request) {
	// redirect to dashboard if user is already logged

	if glob_sign_in.is_logged == true {
		http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
	}

	if glob_forgot_password.is_password_reset_available == false {
		http.Redirect(w, r, fmt.Sprintf("/404"), 302)
		return
	}

	var Page_Title = "Check your Email"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/reset_sent.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/footer_form.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title, "ServerName": glob_forgot_password.server_code, "Email_Sent": glob_forgot_password.admin_email})

	// TODO send password reset instructions
	// password_reset_instruction_sent(glob_forgot_password.server_code, glob_forgot_password.admin_email)

	glob_forgot_password.is_password_reset_available = false
}

func Non_logged_reset_new(w http.ResponseWriter, r *http.Request) {
	// redirect to dashboard if user is already logged

	if glob_sign_in.is_logged == true {
		http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
	}

	server_code := r.URL.Query().Get("scode")
	server_email := r.URL.Query().Get("semail")

	if r.Method == "GET" {
		var Page_Title = "Set New Password"

		if server_code == "" {
			http.Redirect(w, r, fmt.Sprintf("/reset_failed"), 302)
			return
		}

		if server_email == "" {
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
		}

		parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title, "ServerEmail": server_email, "ServerCode": server_code})
	} else {
		parse_error := r.ParseForm()

		if parse_error != nil {
			fmt.Println(parse_error)
			return
		}

		glob_forgot_password.new_password = Hash_Password(r.FormValue("input_psw"))
		Reset_password(glob_forgot_password.server_code, glob_forgot_password.admin_email, glob_forgot_password.new_password)

		http.Redirect(w, r, fmt.Sprintf("/reset_success"), 302)
	}
}

func Non_logged_reset_failed(w http.ResponseWriter, r *http.Request) {
	// redirect to dashboard if user is already logged

	if glob_sign_in.is_logged == true {
		http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
	}

	var Page_Title = "Password Reset Failed"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/reset_failed.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/footer_form.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_reset_success(w http.ResponseWriter, r *http.Request) {
	// redirect to dashboard if user is already logged

	if glob_sign_in.is_logged == true {
		http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
	}

	var Page_Title = "Password Reset Success"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/reset_success.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/footer_form.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

// Logged Section Domain Routing

func Logged_dashboard(w http.ResponseWriter, r *http.Request) {
	// redirect to sign in selection page if user not logged in

	if glob_sign_in.is_logged == false {
		http.Redirect(w, r, fmt.Sprintf("/account_select"), 302)
	}

	var Page_Title = "Server Dashboard"

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
	}

	parsedTemplate.Execute(w, map[string]string{
		"Page_Title": Page_Title,
		"Server_Code": Get_ServerCode(),
		"Dashboard_Name": Get_DashboardName(),
		"Dashboard_Description": Get_DashboardDescription(),
		"Wiki_Content": Get_WikiContent()})
}

func Logged_editor(w http.ResponseWriter, r *http.Request) {
	// redirect to sign in selection page if user not logged in

	if glob_sign_in.is_logged == false {
		http.Redirect(w, r, fmt.Sprintf("/account_select"), 302)
	}

	var Page_Title = "Codinoc Editor"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/logged/editor.tmpl",
		"thm/page/_part/logged/head.tmpl",
		"thm/page/_part/logged/bottom.tmpl",
		"thm/page/_part/logged/navbar.tmpl",
		"thm/page/_part/logged/footer.tmpl",
		"thm/page/_part/logged/panel_team.tmpl",
		"thm/page/_part/logged/panel_chat.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}