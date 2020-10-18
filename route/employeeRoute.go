package route

import (
	"employee/controllers"

	"github.com/labstack/echo/v4"
)

func EmployeeRouteService(e *echo.Echo) {
	e.POST("/employee/add", controllers.AddEmployee)
	e.GET("/employee/search", controllers.SearchEmployee)
	e.GET("/employee/list/:orderByKey", controllers.GetAllEmployee)
	e.POST("/employee/update", controllers.UpdateEmployee)
	e.DELETE("/employee/delete/:id", controllers.DeleteEmployee)
	e.PUT("/employee/restore/:id", controllers.RestoreEmployee)
}
