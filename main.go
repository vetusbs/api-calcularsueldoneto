package main

import (
	"calcularsueldoneto/application"
	"calcularsueldoneto/repository"

	"github.com/gin-gonic/gin"

	"net/http"
)

type salaryResponse struct {
	MonthlyPayroll float32 `json:"monthly_payroll"`
}

type SalaryRequest struct {
	GrossSalary float32 `json:"gross_salary"`
	Region      string  `json:"region"`
	Children    int     `json:"children"`
	Babies      int     `json:"babies"`
}

var calculateNetSalary application.CalculateNetSalary

func main() {
	calculateNetSalary = application.CalculateNetSalary{
		RegionRepository: repository.NewRegionRepository(),
		StateRepository:  repository.NewStateRepository(),
	}

	router := gin.Default()
	router.POST("/netSalary", calculateNetSalaryFucntion)
	router.GET("/", index)

	router.Run("0.0.0.0:8080")
}

func index(c *gin.Context) {
	c.Data(200, "application/json; charset=utf-8", []byte("Good to see you"))
}

func calculateNetSalaryFucntion(c *gin.Context) {

	var salaryRequest SalaryRequest

	if err := c.BindJSON(&salaryRequest); err != nil {
		return
	}

	netSalary := calculateNetSalary.Execute(
		application.CalculateNetSalaryInput{
			GrossSalary: salaryRequest.GrossSalary,
			Region:      salaryRequest.Region,
			Children:    salaryRequest.Children,
			Babies:      salaryRequest.Babies,
		})

	montlySalary := netSalary / 12
	c.IndentedJSON(http.StatusOK, salaryResponse{MonthlyPayroll: montlySalary})
}
