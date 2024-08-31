package global

import (
	"fmt"
	"html/template"
	"net/http"
)

var basicPages = []string{
	"template/base.html",
	"template/components/navigation.html",
	"template/components/footer.html",
}

func ExecuteTemplate(w http.ResponseWriter, r *http.Request, pages []string, data any) {
	pages = append(pages, basicPages...)
	tmpl, parseErr := template.ParseFiles(pages...)
	if parseErr != nil {
		fmt.Println(parseErr)
		HandleError(w, r, InternalServerErr)
		return
	}
	executeErr := tmpl.ExecuteTemplate(w, "base", data)
	if executeErr != nil {
		fmt.Println(executeErr)
		HandleError(w, r, InternalServerErr)
		return
	}
}
