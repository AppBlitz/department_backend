package test

import (
	"testing"

	"github.com/AppBlitz/department_backend/internal/database/mysqls"
	"github.com/AppBlitz/department_backend/internal/repository"
	"github.com/AppBlitz/department_backend/internal/service"
)

func TestServiceGetIDdepartment(t *testing.T) {
	db, _ := mysqls.ConnectionDatabaseMysql()
	repo := repository.NewDepartmentRepository(db)
	servi := service.NewDepartmentService(repo)
	_, err := servi.SearchDepartmentID(1234)
	if err != nil {
		t.Error(err)
	}
}
