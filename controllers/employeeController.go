package controllers

import (
	"employee/models"
	"employee/services"
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

func AddEmployee(c echo.Context) error {
	// searchKey := c.Param("SEARCHKEY")
	employee := new(models.Employee)
	bindError := c.Bind(employee)
	if bindError != nil {
		fmt.Println(bindError)
		return c.String(http.StatusBadRequest, bindError.Error())
	}
	response, addEmployeeError := services.EmployeeService().AddEmployee(*employee)
	if addEmployeeError != nil {
		return c.String(http.StatusBadRequest, response)
	}

	return c.String(http.StatusOK, response)
}

func SearchEmployee(c echo.Context) error {
	searchKey := new(models.Search)
	bindError := c.Bind(searchKey)
	if bindError != nil {
		fmt.Println(bindError)
		return c.String(http.StatusBadRequest, bindError.Error())
	}
	response, addEmployeeError := services.EmployeeService().SearchEmployee(searchKey.SearchKey)
	if addEmployeeError != nil {
		return c.String(http.StatusBadRequest, response)
	}
	return c.String(http.StatusOK, response)
}

func GetAllEmployee(c echo.Context) error {
	orderByKey := c.Param("orderByKey")
	response, addEmployeeError := services.EmployeeService().GetAllEmployee(orderByKey)
	if addEmployeeError != nil {
		return c.String(http.StatusBadRequest, response)
	}
	return c.String(http.StatusOK, response)
}

func UpdateEmployee(c echo.Context) error {
	// searchKey := c.Param("SEARCHKEY")
	employee := new(models.Employee)
	bindError := c.Bind(employee)
	if bindError != nil {
		fmt.Println(bindError)
		return c.String(http.StatusBadRequest, bindError.Error())
	}
	response, addEmployeeError := services.EmployeeService().UpdateEmployee(*employee)
	if addEmployeeError != nil {
		return c.String(http.StatusBadRequest, response)
	}
	return c.String(http.StatusOK, response)
}

func DeleteEmployee(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	response, addEmployeeError := services.EmployeeService().DeleteEmployee(id)
	if addEmployeeError != nil {
		return c.String(http.StatusBadRequest, response)
	}
	return c.String(http.StatusOK, response)
}
func RestoreEmployee(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	response, addEmployeeError := services.EmployeeService().RestoreEmployee(id)
	if addEmployeeError != nil {
		return c.String(http.StatusBadRequest, response)
	}
	return c.String(http.StatusOK, response)
}
