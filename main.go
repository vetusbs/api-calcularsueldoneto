package main

import (
	"calcularsueldoneto/application"

	"github.com/gin-gonic/gin"

	"net/http"
)

type salaryResponse struct {
	MonthlyPayroll float32 `json:"monthly_payroll"`
}

type SalaryRequest struct {
	GrossSalary float32 `json:"gross_salary"`
}

var calculateNetSalary application.CalculateNetSalary

func main() {
	calculateNetSalary = application.CalculateNetSalary{
		RegionRepository: application.RegionRepository{},
		StateRepository:  application.StateRepository{},
	}

	router := gin.Default()
	router.POST("/netSalary", calculateNetSalaryFucntion)

	router.Run("0.0.0.0:8080")
}

func calculateNetSalaryFucntion(c *gin.Context) {

	var salaryRequest SalaryRequest

	if err := c.BindJSON(&salaryRequest); err != nil {
		return
	}

	netSalary := calculateNetSalary.Execute(
		application.CalculateNetSalaryInput{
			GrossSalary: salaryRequest.GrossSalary,
			Region:      "whatever"})

	montlySalary := netSalary / 12
	c.IndentedJSON(http.StatusOK, salaryResponse{MonthlyPayroll: montlySalary})
}
