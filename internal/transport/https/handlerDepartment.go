// Package https
package https

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/AppBlitz/department_backend/configs"
	"github.com/AppBlitz/department_backend/internal/model"
	"github.com/AppBlitz/department_backend/internal/service"
)

type DepartmentHandler struct {
	service *service.DepartmentService
}
type ErrorResponse struct {
	Error string `json:"error"`
}

func NewDepartmentHandler(s *service.DepartmentService) *DepartmentHandler {
	return &DepartmentHandler{service: s}
}

func (serviceDepart *DepartmentHandler) SaveDepartments(w http.ResponseWriter, r *http.Request) {
	department_model := model.Department{}
	defer func() {
		if erro := r.Body.Close(); erro != nil {
			log.Fatal(erro)
		}
	}()
	configs.EnableCors(w)
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodPost {
		http.Error(w, "Method  not allowed", http.StatusMethodNotAllowed)
	} else {
		data, erro := io.ReadAll(r.Body)
		if erro != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: "Bad request"})
		}
		erro = json.Unmarshal(data, &department_model)
		if erro != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: "Bad request"})
		}
		w.WriteHeader(http.StatusCreated)
		_, erro = serviceDepart.service.SaveDepartment(&department_model)
		if erro != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: "Error save department"})
		}
		_, err := w.Write([]byte("Department save success"))
		if err != nil {
			http.Error(w, "Internal server error with response", http.StatusInternalServerError)
		}
	}
}

func (serviceDepart *DepartmentHandler) DepartmentID(w http.ResponseWriter, r *http.Request) {
	configs.EnableCors(w)
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(ErrorResponse{Error: "Method not allowed"})
	} else {
		id, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: "ID no valid"})
			return
		}
		department, err := serviceDepart.service.SearchDepartmentID(int64(id))
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: "Department not found"})
			return
		}
		data, err := json.Marshal(department)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: "Error internal server"})
			return
		}
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: "Error internal server"})
			return
		}
	}
}

func (serviceDepart *DepartmentHandler) FindAllDepartments(w http.ResponseWriter, r *http.Request) {
	configs.EnableCors(w)
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		_ = json.NewEncoder(w).Encode(ErrorResponse{Error: "Method not allowed"})
		return
	} else {
		departmens, _ := serviceDepart.service.FinAllDepartments()
		data, err := json.Marshal(departmens)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: "Server error creating response"})
			return
		}
		w.WriteHeader(http.StatusOK)
		_, err = w.Write(data)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			_ = json.NewEncoder(w).Encode(ErrorResponse{Error: "Server error creating response"})
			return
		}
	}
}
