// Package repository
package repository

import (
	"database/sql"
	"fmt"

	"github.com/AppBlitz/department_backend/internal/model"
)

type DepartmentRepository interface {
	Save(department *model.Department) error
	FindByID(id int64) (*model.Department, error)
}

type mysqlDepartmentRepo struct {
	db *sql.DB
}

func NewDepartmentRepository(db *sql.DB) DepartmentRepository {
	return &mysqlDepartmentRepo{db: db}
}

func (r *mysqlDepartmentRepo) Save(dept *model.Department) error {
	query := "INSERT INTO departments (id,name,description) VALUES (?,?,?)"
	_, err := r.db.Exec(query, dept.ID, dept.Name, dept.Description)
	return err
}

func (r *mysqlDepartmentRepo) FindByID(id int64) (*model.Department, error) {
	department := &model.Department{}
	query := "select depart.id,depart.name,depart.description from departments depar where depart.id=?"
	rows := r.db.QueryRow(query, id)
	if err := rows.Scan(department.ID, department.Name, department.Description); err != nil {
		if err == sql.ErrNoRows {
			return department, fmt.Errorf("%s %d", "Department not found with id", id)
		}
		return department, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return department, nil
}
