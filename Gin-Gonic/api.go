package main

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

//here we use slice as a DB
var users []User

func main() {
	//Seeding sample datas
	users = append(users, User{Id: uuid.New().String(), Name: "Aswath", Age: 21})
	users = append(users, User{Id: uuid.New().String(), Name: "Rock", Age: 24})

	r := gin.Default()
	r.GET("/users", func(c *gin.Context) { //getting all users
		c.JSON(200, users)
	})

	r.GET("/ping", func(c *gin.Context) { //checking whether it is working
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/useradd", func(c *gin.Context) { //user adding
		var reqbody User
		if err := c.ShouldBindJSON(&reqbody); err != nil {
			c.JSON(422, gin.H{
				"message": "Invalid reqest body",
			})
			return
		}
		reqbody.Id = uuid.New().String()
		users = append(users, reqbody)
		c.JSON(200, gin.H{
			"Message": "User Added",
		})

	})

	r.PUT("/:id", func(c *gin.Context) { //user update
		id := c.Param("id")
		var reqbody User
		if err := c.ShouldBindJSON(&reqbody); err != nil {
			c.JSON(422, gin.H{
				"message": "Invalid Reqest Body",
			})
			return
		}
		for i, u := range users {
			if u.Id == id {
				users[i].Name = reqbody.Name
				users[i].Age = reqbody.Age

				c.JSON(200, gin.H{
					"Message": "User Updated",
				})
				return

			}
		}
		c.JSON(404, gin.H{
			"Message": "InValid UserId",
		})

	})

	r.DELETE("/:id", func(c *gin.Context) { //user deletion
		id := c.Param("id")
		for i, u := range users {
			if u.Id == id {
				users = append(users[:i], users[i+1:]...)

				c.JSON(200, gin.H{
					"Message": "User Deleted",
				})
				return
			}
		}
		c.JSON(404, gin.H{
			"Message": "InValid Userid",
		})
	})

	r.Run(":8051") //Running on port 8051

}
