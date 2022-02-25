package app

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Car struct {
	Make  string
	Model string
	Trim  string
	Color string
	Year  int
}

var allCars = []Car{
	{
		Make:  "Toyota",
		Model: "4Runner",
		Trim:  "TRD Pro",
		Color: "Super White",
		Year:  2022,
	},
	{
		Make:  "Toyota",
		Model: "Rav4",
		Trim:  "Adventure",
		Color: "Lunar Rock",
		Year:  2019,
	},
}

func index(c *gin.Context) {
	c.String(200, "Web Application!")
}

func cars(c *gin.Context) {
	c.JSON(0, allCars)
}

func getCar(c *gin.Context) {
	var foundCar = Car{}
	var httpStatus int
	model := c.Param("make")
	fmt.Println(c.Params)
	for i := 0; i < len(allCars); i++ {
		car := allCars[i]
		if car.Model == model {
			fmt.Println(car)
			foundCar = car
			httpStatus = 200
			break
		} else {
			httpStatus = 404
		}
	}
	c.JSON(httpStatus, foundCar)
}

func StartApp() {
	gin.SetMode(gin.DebugMode)

	router := gin.New()
	router.GET("/", index)
	router.GET("/cars", cars)
	router.GET("/cars/:make", getCar)
	router.Run(":8000")
}
