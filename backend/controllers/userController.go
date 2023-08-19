package controllers

import (
	"moneyball/helpers"
	"moneyball/initializers"
	"moneyball/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(c *gin.Context) {
    var user struct {
        Name     string
        Username string
        Password string
    }

    if err := c.Bind(&user); err != nil {
        helpers.HandleError(c, 400, "Invalid request body", err)
        return
    }

    hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
    if err != nil {
        helpers.HandleError(c, 400, "Unable to encrypt the given password", err)
        return
    }

    authToken := helpers.RandomString(10)

    newUser := models.User{
        Name:     user.Name,
        Username: user.Username,
        Password: string(hash),
        AuthToken: authToken,
    }

    if err := initializers.DB.Create(&newUser).Error; err != nil {
        helpers.HandleError(c, 500, "User creation failed", err)
        return
    }

    c.JSON(201, gin.H{
        "message": "User created successfully",
        "user":    newUser,
    })
}

func IndexUsers(c *gin.Context) {
    var users []models.User
    if err := initializers.DB.Find(&users).Error; err != nil {
        helpers.HandleError(c, 500, "Failed to fetch users", err)
        return
    }

    c.JSON(200, gin.H{
        "users": users,
    })
}

func ShowUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User
    if err := initializers.DB.First(&user, id).Error; err != nil {
        helpers.HandleError(c, 404, "User not found", err)
        return
    }

    c.JSON(200, gin.H{
        "user": user,
    })
}

func UpdateUser(c *gin.Context) {
    id := c.Param("id")
    var user models.User
    if err := initializers.DB.First(&user, id).Error; err != nil {
        helpers.HandleError(c, 404, "User not found", err)
        return
    }

    var userUpdate struct {
        Name     string
        Username string
        Password string
    }

    if err := c.Bind(&userUpdate); err != nil {
        helpers.HandleError(c, 400, "Invalid request body", err)
        return
    }

    if err := initializers.DB.Model(&user).Updates(models.User{
        Name:     userUpdate.Name,
        Username: userUpdate.Username,
        Password: userUpdate.Password,
    }).Error; err != nil {
        helpers.HandleError(c, 500, "User update failed", err)
        return
    }

    c.JSON(200, gin.H{
        "message": "User updated successfully",
        "user":    user,
    })
}

func DeleteUser(c *gin.Context) {
    id := c.Param("id")
    if err := initializers.DB.Delete(&models.User{}, id).Error; err != nil {
        helpers.HandleError(c, 500, "Failed to delete user", err)
        return
    }

    c.Status(200)
}

func LoginUser(c *gin.Context) {
    var credentials struct {
        Username string
        Password string
    }

    if err := c.Bind(&credentials); err != nil {
        helpers.HandleError(c, 400, "Invalid request body", err)
        return
    }

    var user models.User
    if err := initializers.DB.First(&user, "username = ?", credentials.Username).Error; err != nil {
        helpers.HandleError(c, 400, "Invalid Username or Password", err)
        return
    }

    if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(credentials.Password)); err != nil {
        helpers.HandleError(c, 400, "Username or Password Incorrect", err)
        return
    }

    c.JSON(200, gin.H{
        "message": "Login successful",
        "user":    user,
    })
}
