package controllers

// import (
// 	"fmt"
// 	"net/http"
// 	"youRJube_API/models"

// 	"golang.org/x/crypto/bcrypt"

// 	"github.com/jinzhu/gorm"

// 	"github.com/gin-gonic/gin"
// 	"github.com/go-playground/validator"
// )

// // RegisterRepository -> ini Repo untuk Register
// type RegisterRepository struct {
// 	Db *gorm.DB
// }

// // RegisterUser -> for register user
// func (t *RegisterRepository) RegisterUser(c *gin.Context) {
// 	// logic register
// 	var u models.User

// 	if err := c.ShouldBind(&u); err != nil {
// 		for _, fieldErr := range err.(validator.ValidationErrors) {
// 			c.JSON(http.StatusBadRequest, fmt.Sprint(fieldErr))
// 			return // exit on first error
// 		}
// 	}

// 	cost, _ := bcrypt.Cost([]byte("salt for hash !"))

// 	hashedPass, _ := bcrypt.GenerateFromPassword([]byte(u.Password), cost)

// 	u.Password = string(hashedPass)

// 	if err := t.Db.Create(&u).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, err.Error())
// 	}

// 	c.JSON(http.StatusOK, gin.H{"message": "Register success"})
// }
