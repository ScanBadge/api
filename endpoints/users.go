package endpoints

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql" // implement MySQL driver
	"github.com/scanbadge/api/configuration"
	"github.com/scanbadge/api/models"
	"github.com/scanbadge/api/utility"
)

// GetUsers gets all users.
func GetUsers(c *gin.Context) {
	var users []models.User
	_, err := configuration.Dbmap.Select(&users, "select * from users")

	if err == nil && len(users) > 0 {
		// Omit password.
		for _, user := range users {
			user.Password = ""
		}

		c.JSON(200, users)
	} else {
		log.Println("error when selecting users from database:", err)
		c.JSON(404, gin.H{"error": "no user(s) found"})
	}
}

// GetUser gets a user based on the provided identifier.
func GetUser(c *gin.Context) {
	id := c.Params.ByName("id")

	if id != "" {
		var user models.User
		err := configuration.Dbmap.SelectOne(&user, "select id,username,firstname,lastname,admin from users where id=?", id)

		if err == nil {
			c.JSON(200, user)
		} else {
			c.JSON(404, gin.H{"error": "user not found"})
		}
	} else {
		c.JSON(422, gin.H{"error": "no identifier provided"})
	}
}

// AddUser adds a new user.
func AddUser(c *gin.Context) {
	var user models.User
	err := c.Bind(&user)

	if err == nil {
		if user.Username != "" && user.Email != "" && user.Password != "" && user.FirstName != "" && user.LastName != "" {
			if len(user.Password) >= 12 {
				// Always hash the password when saving to the database.
				user.Password = utility.HashPassword(user.Password)

				err := configuration.Dbmap.Insert(&user)

				if err == nil {
					showResult(c, 201, user)
				} else {
					fmt.Println("adding new user failed", err)
					c.JSON(400, gin.H{"error": "adding new user failed"})
				}
			} else {
				c.JSON(400, gin.H{"error": "provided password must be at least 12 characters long"})
			}
		} else {
			c.JSON(422, gin.H{"error": "field(s) are empty"})
		}
	} else {
		log.Println("adding new user failed", err)
		c.JSON(400, gin.H{"error": "adding new user failed"})
	}
}

// UpdateUser updates a user based on the identifer.
func UpdateUser(c *gin.Context) {
	c.JSON(403, gin.H{"error": "PUT is not supported"})
}

// DeleteUser deletes a user based on the identifier.
func DeleteUser(c *gin.Context) {
	c.JSON(403, gin.H{"error": "DELETE is not supported"})
}