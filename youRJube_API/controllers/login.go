package controllers

// import (
// 	"net/http"

// 	"golang.org/x/crypto/bcrypt"

// 	"github.com/jinzhu/gorm"

// 	"github.com/gin-gonic/gin"
// )

// // LoginRepository -> ini Repo untuk Login
// type LoginRepository struct {
// 	Db *gorm.DB
// }

// // LoginUser -> for Login user
// func (t *LoginRepository) LoginUser(c *gin.Context) {
// 	// logic Login
// 	var u struct {
// 		Email    string `json:"email" binding:"required"`
// 		Password string `json:"password" binding:"required"`
// 	}

// 	if err := c.ShouldBind(&u); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Please Input email/password"})
// 		return // exit on first error
// 	}
// 	var userData models.User
// 	if err := t.Db.Where(&models.User{Email: u.Email}).Find(&userData).Error; err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Email/Password Salah"})
// 		return
// 	}
// 	if err := bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(u.Password)); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"message": "Email/Password Salah"})
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "Register success"})
// }
