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
	"github.com/gorilla/mux"
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
	var Page_Title = "Sign in As Administrator"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/signin_admin.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/footer_form.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_member_signin(w http.ResponseWriter, r *http.Request) {
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
	if r.Method == "GET" {
		var Page_Title = "Create Server"

		parsedTemplate, err := template.ParseFiles(
			"thm/page/non_logged/create_server.tmpl",
			"thm/page/_part/non_logged/head.tmpl",
			"thm/page/_part/non_logged/bottom.tmpl",
			"thm/page/_part/non_logged/footer_form.tmpl")

		if err != nil {
			fmt.Println("> Unable to parse html file : ", err)
		}

		parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
	} else {
		parse_error := r.ParseForm()

		if parse_error != nil {
			fmt.Println(parse_error)
			return
		}

		glob_sign_up.server_name = r.FormValue("input_sname")
		http.Redirect(w, r, fmt.Sprintf("/create_admin/%s", glob_sign_up.server_name), 302)
	}
}

func Non_logged_create_admin(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		vars := mux.Vars(r)
		server_name := vars["server_name"]

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
		// TODO
	}
}

func Non_logged_create_admin_success(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Congratulations"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/create_admin_success.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/footer_form.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_reset_password(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Reset Password"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/reset_password.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/footer_form.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_reset_sent(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Check your Email"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/reset_sent.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/footer_form.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_reset_new(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Set New Password"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/reset_new.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/footer_form.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_reset_failed(w http.ResponseWriter, r *http.Request) {
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

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Logged_editor(w http.ResponseWriter, r *http.Request) {
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