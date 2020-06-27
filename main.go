package main

import (
	"github.com/programmer-richa/form_template/models"
	"html/template"
	"log"
	"os"
)
var tpl *template.Template
// Initializer
func init() {
	tpl = template.Must(template.New("").ParseGlob("templates/*.gohtml"))
}

//GetRegistrationForm returns a signup form
func GetRegistrationForm() *models.Form {
	name := models.NewFormField("Name", "Name", "",
		"text", "Enter your name.", true, 2, 50)
	email := models.NewFormField("Email address", "Email", "",
		"email", "Enter your email.", true, 2, 50)

	password := models.NewFormField("Password", "Password", "",
		"password", "Enter your password.", true, 8, 50)

	cPassword := models.NewFormField("Confirm Password", "CPassword", "", "password",
		"Enter your confirm password.", true, 8, 50)
	subscribe := models.NewFormField("Subscribe to E-mail alerts", "Subscribe", "on",
		"checkbox", "", false, -1, -1)
	btn := models.NewFormField("Register", "Register", "Register",
		"submit", "", false, -1, -1)
	return models.NewForm("signup", "signup", "/register", "post",
		[]*models.FormField{name,email,password,cPassword,subscribe,btn},
		"register", "")


}

// GetLoginForm returns a login form for Login/Registration Page
func GetLoginForm() *models.Form {
	userName := models.NewFormField("Username", "Username", "",
		"text", "Username or email address", true, 2, 50)

	password := models.NewFormField("Password", "Password", "",
		"password", "Enter your password.", true, 8, 50)

	rememberMe := models.NewFormField("Remember Me on this browser", "Remember", "on",
		"checkbox", "", false, -1, -1)
	btn := models.NewFormField("Login", "Login", "Login",
		"submit", "", false, -1, -1)
	return models.NewForm("login", "login", "/login", "post",
		[]*models.FormField{ userName,  password,
			 rememberMe,  btn},
		"login", "")

}

// Entry point of a program
func main(){
	signupFrm:=GetRegistrationForm()
	loginFrm:=GetLoginForm()
	// Passing struct to a template
	page:=models.NewPage("Home","index", map[string]*models.Form{"signup":signupFrm,"login":loginFrm})
	err := tpl.ExecuteTemplate(os.Stdout, "index.gohtml", page)
	if err != nil {
		log.Fatal(err)
	}
}
