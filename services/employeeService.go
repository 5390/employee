package services

import (
	"employee/dao"
	"employee/models"
	"employee/services/common"
	"fmt"
)

type employeeService struct{}

type EmployeeServiceIF interface {
	AddEmployee(employee models.Employee) (string, error)
	SearchEmployee(key string) (string, error)
	GetAllEmployee(orderByKey string) (string, error)
	UpdateEmployee(employee models.Employee) (string, error)
	DeleteEmployee(id int) (string, error)
	RestoreEmployee(id int) (string, error)
}

func EmployeeService() EmployeeServiceIF {
	return &employeeService{}
}

func (self *employeeService) AddEmployee(employee models.Employee) (string, error) {
	if employee.Name == nil {
		_, rs := common.FormatResult(nil, false, common.New("Employee Name can't be Empty"))
		return common.MarshalJson(rs), common.New("Employee Name can't be Empty")
	}
	employeeId, EmployeeError := dao.EmployeeDao().AddEmployee(employee)
	var responseMessage string
	if employeeId > 0 {
		responseMessage = "Employee is added succesfuly"
	} else {
		responseMessage = "Failed to add employee details"
	}
	_, rs := common.FormatResult(responseMessage, true, EmployeeError)
	return common.MarshalJson(rs), common.New("Employee Name can't be Empty")
}

func (self *employeeService) SearchEmployee(key string) (string, error) {
	empList, EmployeeError := dao.EmployeeDao().SearchEmployee(key)
	_, rs := common.FormatResult(empList, true, EmployeeError)
	return common.MarshalJson(rs), common.New("Employee Name can't be Empty")
}

func (self *employeeService) GetAllEmployee(orderByKey string) (string, error) {
	empList, EmployeeError := dao.EmployeeDao().GetAllEmployee(orderByKey)
	_, rs := common.FormatResult(empList, true, EmployeeError)
	return common.MarshalJson(rs), common.New("Employee Name can't be Empty")
}

func (self *employeeService) UpdateEmployee(employee models.Employee) (string, error) {
	var responseMessage string
	var employeeError error
	if employee.Id > 0 {
		employeeError = dao.EmployeeDao().UpdateEmployee(employee)
		if employeeError != nil {
			responseMessage = "Failed to update employee details"
		} else {
			responseMessage = "Updated employee details successfully"
		}
	} else {
		responseMessage = "Please provide Employee Id to update employee details"
	}
	_, rs := common.FormatResult(responseMessage, true, employeeError)
	return common.MarshalJson(rs), employeeError
}

func (self *employeeService) DeleteEmployee(id int) (string, error) {
	id, empError := dao.EmployeeDao().DeleteEmployee(id)
	if empError != nil {
		fmt.Println(empError)
		_, rs := common.FormatResult(nil, false, empError)
		return common.MarshalJson(rs), empError

	}
	_, rs := common.FormatResult("Employee is Deleted", true, empError)
	return common.MarshalJson(rs), empError
}

func (self *employeeService) RestoreEmployee(id int) (string, error) {
	id, empError := dao.EmployeeDao().RestoreEmployee(id)
	if empError != nil {
		fmt.Println(empError)
		_, rs := common.FormatResult(nil, false, empError)
		return common.MarshalJson(rs), empError

	}
	_, rs := common.FormatResult("Employee is Restored", true, empError)
	return common.MarshalJson(rs), empError
}
