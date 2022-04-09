// Codinoc IDE
//
// Domain Routing Functions

package route

import (
	"html/template"
	"log"
	"net/http"
)

// Server Section Domain Routing

func Server_404(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := template.ParseFiles("thm/page/server/404.html")

	if err != nil {
		log.Fatal("Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, nil)
}

func Server_error(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := template.ParseFiles("thm/page/server/error.html")

	if err != nil {
		log.Fatal("Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, nil)
}

func Server_maintenance(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := template.ParseFiles("thm/page/server/maintenance.html")

	if err != nil {
		log.Fatal("Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, nil)
}

func Server_mobile(w http.ResponseWriter, r *http.Request) {
	parsedTemplate, err := template.ParseFiles("thm/page/server/mobile.html")

	if err != nil {
		log.Fatal("Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, nil)
}

// Non Logged Section Domain Routing

func Non_logged_home(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Welcome to Codinoc"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/home.html",
		"thm/page/_part/non_logged/head.html",
		"thm/page/_part/non_logged/bottom.html",
		"thm/page/_part/non_logged/navbar.html",
		"thm/page/_part/non_logged/footer.html")

	if err != nil {
		log.Fatal("Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_about(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "About Codinoc"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/about.html",
		"thm/page/_part/non_logged/head.html",
		"thm/page/_part/non_logged/bottom.html",
		"thm/page/_part/non_logged/navbar.html",
		"thm/page/_part/non_logged/footer.html")

	if err != nil {
		log.Fatal("Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_contact(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Contact Codinoc"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/contact.html",
		"thm/page/_part/non_logged/head.html",
		"thm/page/_part/non_logged/bottom.html",
		"thm/page/_part/non_logged/navbar.html",
		"thm/page/_part/non_logged/footer.html")

	if err != nil {
		log.Fatal("Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_legal(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Legal Notes"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/legal.html",
		"thm/page/_part/non_logged/head.html",
		"thm/page/_part/non_logged/bottom.html",
		"thm/page/_part/non_logged/navbar.html",
		"thm/page/_part/non_logged/footer.html")

	if err != nil {
		log.Fatal("Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_site_map(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Site Map"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/site_map.html",
		"thm/page/_part/non_logged/head.html",
		"thm/page/_part/non_logged/bottom.html",
		"thm/page/_part/non_logged/navbar.html",
		"thm/page/_part/non_logged/footer.html")

	if err != nil {
		log.Fatal("Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_account_select(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Select account for Sign In"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/signin_account_select.html",
		"thm/page/_part/non_logged/head.html",
		"thm/page/_part/non_logged/bottom.html",
		"thm/page/_part/non_logged/footer_form.html")

	if err != nil {
		log.Fatal("Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_admin_signin(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Sign in As Administrator"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/signin_admin.html",
		"thm/page/_part/non_logged/head.html",
		"thm/page/_part/non_logged/bottom.html",
		"thm/page/_part/non_logged/footer_form.html")

	if err != nil {
		log.Fatal("Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_member_signin(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Sign in As Member"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/signin_member.html",
		"thm/page/_part/non_logged/head.html",
		"thm/page/_part/non_logged/bottom.html",
		"thm/page/_part/non_logged/footer_form.html")

	if err != nil {
		log.Fatal("Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_create_server(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Create Server"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/create_server.html",
		"thm/page/_part/non_logged/head.html",
		"thm/page/_part/non_logged/bottom.html",
		"thm/page/_part/non_logged/footer_form.html")

	if err != nil {
		log.Fatal("Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_create_admin(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Create Server Administrator"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/create_admin.html",
		"thm/page/_part/non_logged/head.html",
		"thm/page/_part/non_logged/bottom.html",
		"thm/page/_part/non_logged/footer_form.html")

	if err != nil {
		log.Fatal("Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_create_admin_success(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Congratulations"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/create_admin_success.html",
		"thm/page/_part/non_logged/head.html",
		"thm/page/_part/non_logged/bottom.html",
		"thm/page/_part/non_logged/footer_form.html")

	if err != nil {
		log.Fatal("Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_reset_password(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Reser Password"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/reset_password.html",
		"thm/page/_part/non_logged/head.html",
		"thm/page/_part/non_logged/bottom.html",
		"thm/page/_part/non_logged/footer_form.html")

	if err != nil {
		log.Fatal("Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_reset_sent(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Check your Email"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/reset_sent.html",
		"thm/page/_part/non_logged/head.html",
		"thm/page/_part/non_logged/bottom.html",
		"thm/page/_part/non_logged/footer_form.html")

	if err != nil {
		log.Fatal("Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_reset_new(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Set New Password"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/reset_new.html",
		"thm/page/_part/non_logged/head.html",
		"thm/page/_part/non_logged/bottom.html",
		"thm/page/_part/non_logged/footer_form.html")

	if err != nil {
		log.Fatal("Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_reset_failed(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Password Reset Failed"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/reset_failed.html",
		"thm/page/_part/non_logged/head.html",
		"thm/page/_part/non_logged/bottom.html",
		"thm/page/_part/non_logged/footer_form.html")

	if err != nil {
		log.Fatal("Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Non_logged_reset_success(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Password Reset Success"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/non_logged/reset_success.html",
		"thm/page/_part/non_logged/head.html",
		"thm/page/_part/non_logged/bottom.html",
		"thm/page/_part/non_logged/footer_form.html")

	if err != nil {
		log.Fatal("Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

// Logged Section Domain Routing

func Logged_dashboard(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Server Dashboard"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/logged/dashboard.html",
		"thm/page/_part/logged/head.html",
		"thm/page/_part/logged/bottom.html",
		"thm/page/_part/logged/navbar.html",
		"thm/page/_part/logged/footer.html",
		"thm/page/_part/logged/panel_team.html",
		"thm/page/_part/logged/panel_chat.html")

	if err != nil {
		log.Fatal("Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}

func Logged_editor(w http.ResponseWriter, r *http.Request) {
	var Page_Title = "Codinoc Editor"

	parsedTemplate, err := template.ParseFiles(
		"thm/page/logged/editor.html",
		"thm/page/_part/logged/head.html",
		"thm/page/_part/logged/bottom.html",
		"thm/page/_part/logged/navbar.html",
		"thm/page/_part/logged/footer.html",
		"thm/page/_part/logged/panel_team.html",
		"thm/page/_part/logged/panel_chat.html")

	if err != nil {
		log.Fatal("Unable to parse html file : ", err)
	}

	parsedTemplate.Execute(w, map[string]string{"Page_Title": Page_Title})
}
