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
	"os"
	"fmt"
	"time"
	"strings"
	"net/http"
	"html/template"
	"encoding/json"
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

	error_server := r.URL.Query().Get("error")

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/signin_member.tmpl",
		"thm/page/_part/non_logged/head.tmpl",
		"thm/page/_part/non_logged/bottom.tmpl",
		"thm/page/_part/non_logged/footer_form.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)

		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": "Member Sign In", "Error_Server": error_server})
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

	var session_string = fmt.Sprintf("%v", session_login.Values["session_string"])
	var session_type = fmt.Sprintf("%v", session_login.Values["session_type"])
	var session_code = fmt.Sprintf("%v", session_login.Values["session_code"])

	parsedTemplate, err := template.ParseFiles(
		"thm/page/logged/dashboard.tmpl",
		"thm/page/_part/logged/head.tmpl",
		"thm/page/_part/logged/bottom.tmpl",
		"thm/page/_part/logged/navbar.tmpl",
		"thm/page/_part/logged/footer.tmpl",
		"thm/page/_part/logged/panel_team.tmpl",
		"thm/page/_part/logged/panel_chat.tmpl",
		"thm/page/_part/logged/modal.tmpl")

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
	// Get_User_Name := Dashboard_Get_User_Name(session_type, session_code)
	Get_User_Name := session_string
	Get_Server_Code := Dashboard_Get_Server_Code()
	Get_Server_Name := Dashboard_Get_Server_Name()

	parsedTemplate.Execute(w, map[string]string{
		"Page_Title": "Dashboard",
		"Session_Type": session_type,
		"Session_Code": session_code,
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
		"Server_Name": Get_Server_Name,
		"Is_Member_01": Dashboard_Get_Member_Availability(1),
		"Is_Member_02": Dashboard_Get_Member_Availability(2),
		"Is_Member_03": Dashboard_Get_Member_Availability(3),
		"Is_Member_04": Dashboard_Get_Member_Availability(4),
		"Is_Member_05": Dashboard_Get_Member_Availability(5),
		"Is_Member_06": Dashboard_Get_Member_Availability(6),
		"Is_Member_07": Dashboard_Get_Member_Availability(7),
		"Is_Member_08": Dashboard_Get_Member_Availability(8),
		"Is_Member_09": Dashboard_Get_Member_Availability(9),
		"Is_Member_10": Dashboard_Get_Member_Availability(10),
		"Member_01_Name": Dashboard_Get_Member_Name(1),
		"Member_02_Name": Dashboard_Get_Member_Name(2),
		"Member_03_Name": Dashboard_Get_Member_Name(3),
		"Member_04_Name": Dashboard_Get_Member_Name(4),
		"Member_05_Name": Dashboard_Get_Member_Name(5),
		"Member_06_Name": Dashboard_Get_Member_Name(6),
		"Member_07_Name": Dashboard_Get_Member_Name(7),
		"Member_08_Name": Dashboard_Get_Member_Name(8),
		"Member_09_Name": Dashboard_Get_Member_Name(9),
		"Member_10_Name": Dashboard_Get_Member_Name(10),
		"Member_01_User_Name": Dashboard_Get_Member_User_Name(1),
		"Member_02_User_Name": Dashboard_Get_Member_User_Name(2),
		"Member_03_User_Name": Dashboard_Get_Member_User_Name(3),
		"Member_04_User_Name": Dashboard_Get_Member_User_Name(4),
		"Member_05_User_Name": Dashboard_Get_Member_User_Name(5),
		"Member_06_User_Name": Dashboard_Get_Member_User_Name(6),
		"Member_07_User_Name": Dashboard_Get_Member_User_Name(7),
		"Member_08_User_Name": Dashboard_Get_Member_User_Name(8),
		"Member_09_User_Name": Dashboard_Get_Member_User_Name(9),
		"Member_10_User_Name": Dashboard_Get_Member_User_Name(10),
		"Is_Project_01": Dashboard_Get_Project_Availability(1),
		"Is_Project_02": Dashboard_Get_Project_Availability(2),
		"Is_Project_03": Dashboard_Get_Project_Availability(3),
		"Is_Project_04": Dashboard_Get_Project_Availability(4),
		"Is_Project_05": Dashboard_Get_Project_Availability(5),
		"Is_Project_06": Dashboard_Get_Project_Availability(6),
		"Is_Project_07": Dashboard_Get_Project_Availability(7),
		"Is_Project_08": Dashboard_Get_Project_Availability(8),
		"Is_Project_09": Dashboard_Get_Project_Availability(9),
		"Is_Project_10": Dashboard_Get_Project_Availability(10),
		"Project_01_Name": Dashboard_Get_Project_Name(1),
		"Project_02_Name": Dashboard_Get_Project_Name(2),
		"Project_03_Name": Dashboard_Get_Project_Name(3),
		"Project_04_Name": Dashboard_Get_Project_Name(4),
		"Project_05_Name": Dashboard_Get_Project_Name(5),
		"Project_06_Name": Dashboard_Get_Project_Name(6),
		"Project_07_Name": Dashboard_Get_Project_Name(7),
		"Project_08_Name": Dashboard_Get_Project_Name(8),
		"Project_09_Name": Dashboard_Get_Project_Name(9),
		"Project_10_Name": Dashboard_Get_Project_Name(10),
		"Project_01_Code": Dashboard_Get_Project_Code(1),
		"Project_02_Code": Dashboard_Get_Project_Code(2),
		"Project_03_Code": Dashboard_Get_Project_Code(3),
		"Project_04_Code": Dashboard_Get_Project_Code(4),
		"Project_05_Code": Dashboard_Get_Project_Code(5),
		"Project_06_Code": Dashboard_Get_Project_Code(6),
		"Project_07_Code": Dashboard_Get_Project_Code(7),
		"Project_08_Code": Dashboard_Get_Project_Code(8),
		"Project_09_Code": Dashboard_Get_Project_Code(9),
		"Project_10_Code": Dashboard_Get_Project_Code(10),
	})
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

	parsedTemplate, err := template.ParseFiles(
		"thm/page/logged/editor.tmpl",
		"thm/page/_part/logged/head.tmpl",
		"thm/page/_part/logged/bottom.tmpl",
		"thm/page/_part/logged/navbar.tmpl",
		"thm/page/_part/logged/footer.tmpl",
		"thm/page/_part/logged/panel_team.tmpl",
		"thm/page/_part/logged/panel_chat.tmpl",
		"thm/page/_part/logged/modal.tmpl")

	if err != nil {
		fmt.Println("> Unable to parse html file : ", err)

		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	}

	// var session_string = fmt.Sprintf("%v", session_login.Values["session_string"])
	var session_type = fmt.Sprintf("%v", session_login.Values["session_type"])
	var session_code = fmt.Sprintf("%v", session_login.Values["session_code"])

	project_code := r.URL.Query().Get("project")

	path_html := fmt.Sprintf("cdn/%s/%s/%s", session_code, project_code, "index.html")
	path_js := fmt.Sprintf("cdn/%s/%s/%s", session_code, project_code, "source_js.js")
	path_css := fmt.Sprintf("cdn/%s/%s/%s", session_code, project_code, "source_css.css")
	path_scss := fmt.Sprintf("cdn/%s/%s/%s", session_code, project_code, "source_scss.scss")

	Get_Navbar_Name := project_code

	// Get_Server_Name := Dashboard_Get_Server_Name()

	fmt.Println()

	// fmt.Println(path_html)
	// fmt.Println(path_js)
	// fmt.Println(path_css)
	// fmt.Println(path_scss)

	parsedTemplate.Execute(w, map[string]string{
		"Page_Title": "Editor",
		"Session_Type": session_type,
		"Navbar_Name": Get_Navbar_Name,
		"Server_Code": session_code,
		"HTML_PATH": path_html,
		"JS_PATH": path_js,
		"CSS_PATH": path_css,
		"SCSS_PATH": path_scss,
	})
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
		user_name := r.FormValue("input_name")
		user_password := r.FormValue("input_password")

		if Process_Is_Server_Available_From_Code(server_code) == false {
			http.Redirect(w, r, fmt.Sprintf("/member_signin?error=%s", "No Server Available"), 302)
			return
		}

		if Process_Is_Member_Available_In_Server(server_code, user_name) == false {
			http.Redirect(w, r, fmt.Sprintf("/member_signin?error=%s", "Invalid Member User Name"), 302)
			return
		}

		if Process_Is_Member_Password_Match(server_code, user_name, user_password) == false {
			http.Redirect(w, r, fmt.Sprintf("/member_signin?error=%s", "Password does not match"), 302)
			return
		}

		// Sign In Session Store Process

		SignIn_Session_Enable(w, r, user_name, "member", server_code)

		// Connect Server Database

		SignIn_Database_Connect(server_code)

		http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
	}
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
	if r.Method == "GET" {
		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	} else {
		parse_error := r.ParseForm()

		if parse_error != nil {
			fmt.Println(parse_error)
			return
		}

		new_content := r.FormValue("wiki_content")

		Dashboard_Update_Wiki(new_content)

		http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
	}
}

func Route_Process_Update_User(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	} else {
		parse_error := r.ParseForm()

		if parse_error != nil {
			fmt.Println(parse_error)
			return
		}

		session_type := r.FormValue("session_type")
		session_code := r.FormValue("session_code")

		first_name := r.FormValue("first_name")
		middle_name := r.FormValue("middle_name")
		last_name := r.FormValue("last_name")
		password := r.FormValue("password")

		Dashboard_Update_User_Profile(session_type, session_code, first_name, middle_name, last_name, password)

		http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
	}
}

func Route_Process_Update_Server(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	} else {
		parse_error := r.ParseForm()

		if parse_error != nil {
			fmt.Println(parse_error)
			return
		}

		dashboard_name := r.FormValue("dashboard_name")
		dashboard_description := r.FormValue("dashboard_description")

		Dashboard_Update_Server_Profile(dashboard_name, dashboard_description)
		http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
	}
}

func Route_Process_Add_Team_Member(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	} else {
		parse_error := r.ParseForm()

		if parse_error != nil {
			fmt.Println(parse_error)
			return
		}

		user_name := r.FormValue("user_name")
		password := r.FormValue("password")

		Dashboard_Create_New_Team_Member(user_name, password)

		http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
	}
}

func Route_Process_Remove_Team_Member(w http.ResponseWriter, r *http.Request) {
	user_name := r.URL.Query().Get("uname")

	if user_name == "" {
		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	}

	Dashboard_Remove_Team_Member(user_name)

	http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
}

func Route_Process_Create_Project(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		http.Redirect(w, r, fmt.Sprintf("/server"), 302)
		return
	} else {
		parse_error := r.ParseForm()

		if parse_error != nil {
			fmt.Println(parse_error)
			return
		}

		project_name := r.FormValue("project_name")
		server_code := r.FormValue("server_code")

		Dashboard_Create_New_Project(server_code, project_name)

		http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
	}
}

func Route_Process_Remove_Project(w http.ResponseWriter, r *http.Request) {
	project_code := r.URL.Query().Get("project_code")

	// Drop data from table

	var query_tb_delete = fmt.Sprintf("DELETE FROM schema_project.tb_main WHERE project_code='%s';", project_code)
	_, err_delete := db_server.Query(query_tb_delete)
	Check_Error(err_delete)

	fmt.Println("> Project Removed: ", project_code)

	http.Redirect(w, r, fmt.Sprintf("/dashboard"), 302)
}

func Route_Update_Source(w http.ResponseWriter, r *http.Request) {
	type Get_Data struct {
		Content string
		Type string
		Path string
		Project string
		Server string
	}

	var data Get_Data

	decoder := json.NewDecoder(r.Body)
	err := decoder.Decode(&data)

	if err != nil {
		panic(err)
	}

	get_Content := fmt.Sprintf("%+v", data.Content)
	get_Type := fmt.Sprintf("%+v", data.Type)
	get_Path := fmt.Sprintf("%+v", data.Path)
	get_Project := fmt.Sprintf("%+v", data.Project)
	get_Server := fmt.Sprintf("%+v", data.Server)

	// get_Time := time.Now()
	// get_Date := "Unknown"

	// fmt.Println("File Content :", get_Content)
	// fmt.Println("File Type    :", get_Type)
	// fmt.Println("File Path    :", get_Path)

	file, err := os.Create(get_Path)

	// if err != nil {
	// 	return
	// }

	defer file.Close()

	file.WriteString(get_Content)

	// Create File History

	// query_tb_insert := fmt.Sprintf("INSERT INTO schema_file.tb_main ( project_code, file_type, file_name, file_conetent, updated_time, updated_date, is_last_save) VALUES ('%s', '%s', '%s', '%s', '%s', '%s', '%s');", get_Project, get_Type, get_Path, get_Content, get_Time, get_Date, "FALSE")
	// db_server.Query(query_tb_insert)

	if get_Type == "HTML" {
		currentTime := time.Now()

		html_name := fmt.Sprintf("history_%s.html", currentTime.Format("2006-01-02_3-4-5-pm"))
		html_path := fmt.Sprintf("cdn/%s/%s/history_html/%s", get_Server, get_Project, html_name)

		file_html, _ := os.Create(html_path)
		defer file_html.Close()
		file_html.WriteString(get_Content)
	} else if get_Type == "JS" {
		currentTime := time.Now()

		js_name := fmt.Sprintf("history_%s.js", currentTime.Format("2006-01-02_3-4-5-pm"))
		js_path := fmt.Sprintf("cdn/%s/%s/history_js/%s", get_Server, get_Project, js_name)

		file_js, _ := os.Create(js_path)
		defer file_js.Close()
		file_js.WriteString(get_Content)
	} else if get_Type == "CSS" {
		currentTime := time.Now()

		css_name := fmt.Sprintf("history_%s.css", currentTime.Format("2006-01-02_3-4-5-pm"))
		css_path := fmt.Sprintf("cdn/%s/%s/history_css/%s", get_Server, get_Project, css_name)

		file_css, _ := os.Create(css_path)
		defer file_css.Close()
		file_css.WriteString(get_Content)
	} else if get_Type == "SCSS" {
		currentTime := time.Now()

		scss_name := fmt.Sprintf("history_%s.scss", currentTime.Format("2006-01-02_3-4-5-pm"))
		scss_path := fmt.Sprintf("cdn/%s/%s/history_scss/%s", get_Server, get_Project, scss_name)

		file_scss, _ := os.Create(scss_path)
		defer file_scss.Close()
		file_scss.WriteString(get_Content)
	}
}