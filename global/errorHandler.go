package global

import (
	"html/template"
	"net/http"
)

type Error struct {
	Message string
	Code    int
}

// HandleError sends an HTTP response with the specified error code and renders an error page using the provided error details.
func HandleError(w http.ResponseWriter, r *http.Request, errType Error) {
	pages := []string{
		"template/base.html",
		"template/pages/error.html",
		"template/components/navigation.html",
		"template/components/footer.html",
	}
	tmpl, err := template.ParseFiles(pages...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	err = tmpl.ExecuteTemplate(w, "base", errType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
