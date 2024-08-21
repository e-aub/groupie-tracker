package global

import (
	"fmt"
	"html/template"
	"net/http"
)

func ExecuteTemplate(w http.ResponseWriter, r *http.Request, pages []string, data any) {
	basic_pages := []string{
		"template/base.html",
		"template/components/navigation.html",
	}
	pages = append(pages, basic_pages...)
	tmpl, err_tmpl := template.ParseFiles(pages...)
	if err_tmpl != nil {
		fmt.Println("err", err_tmpl)
		return
	}
	err := tmpl.ExecuteTemplate(w, "base", data)
	if err != nil {
		fmt.Println("err", err.Error())
		return
	}
}
