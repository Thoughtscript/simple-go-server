package handlers

import (
	"html/template"
	"log"
	"net/http"
	m "simplegoserver/models"
	tmp "simplegoserver/templates"
)

func WriteOut(w http.ResponseWriter, r *http.Request) {
	person := m.PersonModel{
		Name:   "Satish",
		Emails: []string{"satish@rubylearning.org", "satishtalim@gmail.com"},
	}

	t := template.New("Person template")

	t, err := t.Parse(tmp.Template)
	if err != nil {
		log.Fatal("Parse: ", err)
		return
	}

	err = t.Execute(w, person)
	if err != nil {
		log.Fatal("Execute: ", err)
		return
	}
}
