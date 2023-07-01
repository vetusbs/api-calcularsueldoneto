package main

import (
	"github.com/gin-gonic/gin"

	"net/http"
)

type netSalary struct {
	MonthlyPayroll float32 `json:"monthly_payroll"`
}

type SalaryRequest struct {
	GrossSalary int `json:"gross_salary"`
}

func main() {
	router := gin.Default()
	router.POST("/netSalary", calculateNetSalary)

	router.Run("localhost:8080")
}

func calculateNetSalary(c *gin.Context) {

	var salaryRequest SalaryRequest

	if err := c.BindJSON(&salaryRequest); err != nil {
		return
	}

	montlySalary := float32(salaryRequest.GrossSalary/12) * float32(0.7)
	c.IndentedJSON(http.StatusOK, netSalary{MonthlyPayroll: montlySalary})
}
