package global

import (
	"html/template"
	"net/http"
)

func ExecuteTemplate(w http.ResponseWriter, r *http.Request, pages []string, data any) {
	err := Error{Code: http.StatusInternalServerError, Message: "Server Error!"}
	basic_pages := []string{
		"template/base.html",
		"template/components/navigation.html",
		"template/components/footer.html",
	}
	pages = append(pages, basic_pages...)
	tmpl, err_tmpl := template.ParseFiles(pages...)
	if err_tmpl != nil {
		HandleError(w, r, err)
		return
	}
	err_tmpl = tmpl.ExecuteTemplate(w, "base", data)
	if err_tmpl != nil {
		HandleError(w, r, err)
		return
	}
}
