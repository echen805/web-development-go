package controllers

import (
	"html/template"
	"net/http"

	"github.com/echen805/web-development-go/views"
)

type Static struct {
	Template views.Template
}

func (static Static) ServeHttp(w http.ResponseWriter, r *http.Request) {
	static.Template.Execute(w, nil)
}

func StaticHandler(tpl views.Template) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, nil)
	}
}

func FAQ(tpl views.Template) http.HandlerFunc {
	questions := []struct {
		Question string
		Answer   template.HTML
	}{
		{
			Question: "Is there a free version?",
			Answer:   "Sure why not!",
		},
		{
			Question: "How do I contact you?",
			Answer:   "Email me <a href='mailto:edwardchen109@gmail.com'>here</a>",
		},
	}
	return func(w http.ResponseWriter, r *http.Request) {
		tpl.Execute(w, questions)
	}
}
