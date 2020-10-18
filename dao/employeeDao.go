package dao

import (
	"employee/db"
	"employee/models"
	"fmt"
)

type employeeDao struct{}

type EmployeeDaoIF interface {
	AddEmployee(employee models.Employee) (int, error)
	SearchEmployee(key string) ([]models.Employee, error)
	GetAllEmployee(orderByKey string) ([]models.Employee, error)
	UpdateEmployee(employee models.Employee) error
	DeleteEmployee(id int) (int, error)
	RestoreEmployee(id int) (int, error)
}

func EmployeeDao() EmployeeDaoIF {
	return &employeeDao{}
}

func (self *employeeDao) AddEmployee(employee models.Employee) (int, error) {
	var employeeId int
	db, ConnectionErrs := db.SqlxConnect()
	if ConnectionErrs != nil {
		return employeeId, ConnectionErrs
	}
	sqlStatement := `INSERT INTO employee(
		name, address, department, skills)
		VALUES ($1, $2, $3, $4) RETURNING id`
	errMainSql := db.Get(&employeeId, sqlStatement, employee.Name, employee.Address, employee.Department, employee.Skills)
	if errMainSql != nil {
		fmt.Println(errMainSql)
		return employeeId, errMainSql
	}
	defer db.Close()
	return employeeId, nil
}

func (self *employeeDao) SearchEmployee(searchKey string) ([]models.Employee, error) {
	employeeList := []models.Employee{}
	db, ConnectionErrs := db.SqlxConnect()
	if ConnectionErrs != nil {
		return employeeList, ConnectionErrs
	}
	sqlStatement := `SELECT name, address, department, skills
	FROM employee where is_deleted=0 and name like $1 or address like $1 or department like $1 
	or skills like $1`
	errMainSql := db.Select(&employeeList, sqlStatement, searchKey)
	if errMainSql != nil {
		fmt.Println(errMainSql)
		return employeeList, errMainSql
	}
	defer db.Close()
	return employeeList, nil
}

func (self *employeeDao) GetAllEmployee(orderByKey string) ([]models.Employee, error) {
	employeeList := []models.Employee{}
	db, ConnectionErrs := db.SqlxConnect()
	if ConnectionErrs != nil {
		return employeeList, ConnectionErrs
	}
	sqlStatement := `SELECT id,name, address, department, skills
	FROM employee where is_deleted=0 order by $1`
	errMainSql := db.Select(&employeeList, sqlStatement, orderByKey)
	if errMainSql != nil {
		fmt.Println(errMainSql)
		return employeeList, errMainSql
	}
	defer db.Close()
	return employeeList, nil
}

func (self *employeeDao) UpdateEmployee(employee models.Employee) error {
	var empId int
	db, ConnectionErrs := db.SqlxConnect()
	if ConnectionErrs != nil {
		return ConnectionErrs
	}
	var sqlStatement string
	var errMainSql error
	if employee.Address != "" && employee.Department != nil && employee.Skills != nil {
		sqlStatement = `UPDATE employee SET address=$1, department=$2, skills=$3 WHERE id=$4 RETURNING id`
		errMainSql = db.Get(&empId, sqlStatement, employee.Address, employee.Department, employee.Skills, employee.Id)
	} else if employee.Department != nil && employee.Skills != nil {
		sqlStatement = `UPDATE employee SET department=$1, skills=$2 WHERE id=$3 RETURNING id`
		errMainSql = db.Get(&empId, sqlStatement, employee.Department, employee.Skills, employee.Id)
	} else if employee.Address != "" && employee.Department != nil {
		sqlStatement = `UPDATE employee SET address=$1, department=$2 WHERE id=$3 RETURNING id`
		errMainSql = db.Get(&empId, sqlStatement, employee.Address, employee.Department, employee.Id)
	} else if employee.Address != "" && employee.Skills != nil {
		sqlStatement = `UPDATE employee SET address=$1, skills=$2 WHERE id=$3 RETURNING id`
		errMainSql = db.Get(&empId, sqlStatement, employee.Address, employee.Skills, employee.Id)
	} else if employee.Address != "" {
		sqlStatement = `UPDATE employee SET address=$1 WHERE id=$2 RETURNING id`
		errMainSql = db.Get(&empId, sqlStatement, employee.Address, employee.Id)
	} else if employee.Department != nil {
		sqlStatement = `UPDATE employee SET department=$1 WHERE id=$2 RETURNING id`
		errMainSql = db.Get(&empId, sqlStatement, employee.Department, employee.Id)
	} else if employee.Skills != nil {
		sqlStatement = `UPDATE employee SET skills=$1 WHERE id=$2;`
		errMainSql = db.Get(&empId, sqlStatement, employee.Skills, employee.Id)
	}
	if errMainSql != nil {
		fmt.Println(errMainSql)
		return errMainSql
	}
	defer db.Close()
	return errMainSql
}

func (self *employeeDao) DeleteEmployee(id int) (int, error) {
	var empId int
	db, ConnectionErrs := db.SqlxConnect()
	if ConnectionErrs != nil {
		return empId, ConnectionErrs
	}
	sqlStatement := `UPDATE employee SET is_deleted=1 WHERE id=$1 RETURNING id `
	errMainSql := db.Get(&empId, sqlStatement, id)

	if errMainSql != nil {
		fmt.Println(errMainSql)
		return empId, errMainSql
	}
	defer db.Close()
	return empId, errMainSql
}

func (self *employeeDao) RestoreEmployee(id int) (int, error) {
	var empId int
	db, ConnectionErrs := db.SqlxConnect()
	if ConnectionErrs != nil {
		return empId, ConnectionErrs
	}
	sqlStatement := `UPDATE employee SET is_deleted=0 WHERE id=$1 RETURNING id `
	errMainSql := db.Get(&empId, sqlStatement, id)

	if errMainSql != nil {
		fmt.Println(errMainSql)
		return empId, errMainSql
	}
	defer db.Close()
	return empId, errMainSql
}
