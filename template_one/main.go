package main

import (
	"fmt"
	"html/template"
	"os"
)


func main() {
	greetings := `Hello {{.FirstName}} {{.LastName}},
	Thank you for registering. Your account is created and must be activated before you can use it. To activate the account click on the following link or copy-paste it on your browser:
	{{.ActivationURL}}
	After activation you may login to {{.LoginURL}} using the following url.
	Thanks
	Sabirul Islam`

	data := struct {
		FirstName     string
		LastName      string
		ActivationURL string
		LoginURL      string
	}{
		"sabirul",
		"islam",
		"htps://mdsabirulislam.blogspot.com",
		"htps://mdsabirulislamshimul@blogspot.com",
	}

	tpl := template.New("email")
	email, err := tpl.Parse(greetings)
	if err != nil {
		fmt.Println(err)
	}
	email.Execute(os.Stdout, data)
}