// Package https
package https

import (
	"encoding/json"
	"net/http"

	"github.com/AppBlitz/department_backend/internal/service"
)

type DepartmentHandler struct {
	ser *service.DepartmentService
}

func NewDepartmentHandler(s *service.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{ser: s}
}

func (serviceDepart *DepartmentHandler) SaveDepartments(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method "+r.Method+" not allowed", http.StatusMethodNotAllowed)
	} else {
		w.WriteHeader(http.StatusCreated)
		w.Header().Set("Content-Type", "application/json")
		_, err := w.Write([]byte("Hello world"))
		if err != nil {
			http.Error(w, "Internal server error with response", http.StatusInternalServerError)
		}
	}
}

func (serviceDepart *DepartmentHandler) Receiver(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method  not allowed", http.StatusMethodNotAllowed)
	} else {
		w.WriteHeader(http.StatusOK)
		w.Header().Set("Content-Type", "application/json")
	}
}

func (serviceDepart *DepartmentHandler) DepartmentID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method ", http.StatusMethodNotAllowed)
	} else {
		department, err := serviceDepart.ser.SearchDepartmentID(1234)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		data, err := json.Marshal(department)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(data)
		if err != nil {
			http.Error(w, "", http.StatusInternalServerError)
			return
		}
	}
}

func (serviceDepart *DepartmentHandler) FindAllDepartments(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	} else {
		departmens, _ := serviceDepart.ser.FinAllDepartments()
		data, err := json.Marshal(departmens)
		if err != nil {
			http.Error(w, "Error internal server", http.StatusInternalServerError)
			return
		}
		_, err = w.Write(data)
		if err != nil {
			http.Error(w, "Error internal server", http.StatusInternalServerError)
			return
		}
	}
}
