package main

// import (
// 	"fmt"
// 	"log"
// 	"net/http"
	
// 	"yourJube_API/controllers"

// 	"github.com/gin-gonic/gin"
// 	_ "github.com/gin-gonic/gin"
// 	"github.com/jinzhu/gorm"
// 	_ "github.com/jinzhu/gorm/dialects/postgres"
// 	_ "github.com/lib/pq"
// )

// const dbConnection = "postgres://postgres:alvin123@localhost:5432/youRJube?sslmode=disable"

// func getUserByID(c *gin.Context) {
// 	// validasi if id valid
// 	userID := c.Params.ByName("id")
// 	fmt.Println(userID)
// 	db, err := gorm.Open("postgres", dbConnection)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, err.Error())
// 		return
// 	}
// 	defer db.Close()
// 	data := models.User{}

// 	db.Where(&models.User{ID: userID}).Find(&data)
// 	c.JSON(200, data)
// 	c.JSON(http.StatusOK, gin.H{"status": "ok"})
// }

// func main() {
// 	db, err := gorm.Open("postgres", dbConnection)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	defer db.Close()

// 	fmt.Println("Listening at Port : 3000")

// 	router := gin.Default()

// 	RegisterController := controllers.RegisterRepository{Db: db}

// 	router.POST("/user", RegisterController.RegisterUser)
// 	router.GET("/user/:id", getUserByID)

// 	err = http.ListenAndServe(":3000", router)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// }
