package controllers

import (
	"moneyball/helpers"
	"moneyball/initializers"
	"moneyball/models"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
    var body struct {
        Name    string
        Balance float64
    }

    if err := c.Bind(&body); err != nil {
        helpers.HandleError(c, 400, "Invalid request body", err)
        return
    }

    account := models.Account{Name: body.Name, Balance: body.Balance}
    if err := initializers.DB.Create(&account).Error; err != nil {
        helpers.HandleError(c, 500, "Account creation failed", err)
        return
    }

    c.JSON(201, gin.H{
        "message": "Account created successfully",
        "account": account,
    })
}

func IndexAccounts(c *gin.Context) {
    var accounts []models.Account
    if err := initializers.DB.Find(&accounts).Error; err != nil {
        helpers.HandleError(c, 500, "Failed to fetch accounts", err)
        return
    }

    c.JSON(200, gin.H{
        "accounts": accounts,
    })
}

func ShowAccount(c *gin.Context) {
    id := c.Param("id")
    var account models.Account
    if err := initializers.DB.First(&account, id).Error; err != nil {
        helpers.HandleError(c, 404, "Account not found", err)
        return
    }

    c.JSON(200, gin.H{
        "account": account,
    })
}

func UpdateAccount(c *gin.Context) {
    id := c.Param("id")

    var account models.Account
    if err := initializers.DB.First(&account, id).Error; err != nil {
        helpers.HandleError(c, 404, "Account not found", err)
        return
    }

    var body struct {
        Name    string
        Balance float64
    }

    if err := c.Bind(&body); err != nil {
        helpers.HandleError(c, 400, "Invalid request body", err)
        return
    }

    if err := initializers.DB.Model(&account).Updates(models.Account{
        Name:    body.Name,
        Balance: body.Balance,
    }).Error; err != nil {
        helpers.HandleError(c, 500, "Failed to update account", err)
        return
    }

    c.JSON(200, gin.H{
        "message": "Account updated successfully",
        "account": account,
    })
}

func DeleteAccount(c *gin.Context) {
    id := c.Param("id")
    if err := initializers.DB.Delete(&models.Account{}, id).Error; err != nil {
        helpers.HandleError(c, 500, "Failed to delete account", err)
        return
    }

    c.Status(200)
}
