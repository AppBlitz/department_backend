// Package app
package app

import (
	"database/sql"
	"net/http"

	"github.com/AppBlitz/department_backend/internal/repository"
	"github.com/AppBlitz/department_backend/internal/service"
	"github.com/AppBlitz/department_backend/internal/transport/https"
)

func Run(db *sql.DB) {
	repo := repository.NewDepartmentRepository(db)
	servi := service.NewDepartmentService(repo)
	trans := https.NewDepartmentHandler(servi)
	erro := http.ListenAndServe("localhost:4567", https.AllHandlers(trans))
	if erro != nil {
		panic("Error with server")
	}
}
