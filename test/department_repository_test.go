package test

import (
	"testing"

	"github.com/AppBlitz/department_backend/internal/database/mysqls"
	"github.com/AppBlitz/department_backend/internal/repository"
)

func TestGetDepartmentID(t *testing.T) {
	db, _ := mysqls.ConnectionDatabaseMysql()
	repo := repository.NewDepartmentRepository(db)
	_, err := repo.FindByID(1234)
	if err != nil {
		t.Error(err)
	}
}

func TestFindAllDepartments(t *testing.T) {
	db, _ := mysqls.ConnectionDatabaseMysql()
	repo := repository.NewDepartmentRepository(db)
	_, err := repo.FindAll()
	if err != nil {
		t.Error(err)
	}
}
