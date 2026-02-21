package https

import (
	"net/http"
)

func AllHandlers(han *DepartmentHandler) *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/save/departments/", han.SaveDepartments)
	return mux
}
