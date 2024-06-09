package lostpapyrus

import (
	"fmt"
	"net/http"
	"html/template"
)

type LostPapyrus struct {}

func (lib *LostPapyrus) Get(returntype string, returnitem string, route string) {
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodGet {
			if returntype == "html" {
				t, err := template.New("response").Parse(returnitem)
				if err != nil {
					http.Error(w, "Error parsing template", http.StatusInternalServerError)
					return
				}
				if err := t.Execute(w, nil); err != nil {
					http.Error(w, "Error executing template", http.StatusInternalServerError)
				}
			} else if returntype == "text" {
				fmt.Fprintf(w, returnitem)
			} else {
				http.Error(w, "The return type doesn't match any available on LostPapyrus. You must use 'text' or 'html' as arguments, but you used '"+returntype+"'", http.StatusBadRequest)
			}
		} else {
			http.Error(w, "LostPapyrus 627: The method used in the request is not allowed.", http.StatusMethodNotAllowed)
		}
	})
}

func (lib *LostPapyrus) Post(returntype string, returnitem string, route string) {
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			if returntype == "html" {
				t, err := template.New("response").Parse(returnitem)
				if err != nil {
					http.Error(w, "Error parsing template", http.StatusInternalServerError)
					return
				}
				if err := t.Execute(w, nil); err != nil {
					http.Error(w, "Error executing template", http.StatusInternalServerError)
				}
			} else if returntype == "text" {
				fmt.Fprintf(w, returnitem)
			} else {
				http.Error(w, "The return type doesn't match any available on LostPapyrus. You must use 'text' or 'html' as arguments, but you used '"+returntype+"'", http.StatusBadRequest)
			}
		} else {
			http.Error(w, "LostPapyrus 627: The method used in the request is not allowed.", http.StatusMethodNotAllowed)
		}
	})
}

func New() *LostPapyrus {
	return &LostPapyrus{}
}
