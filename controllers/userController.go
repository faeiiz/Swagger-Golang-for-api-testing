package controllers

import (
	"log"
	"net/http"
	"swag/initializers"
	"swag/models"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

// CreateUser godoc
// @Summary Create a new user
// @Description Add a new user to the database
// @Tags users
// @Accept json
// @Produce json
// @Param user body models.User true "User details"
// @Success 201 {object} models.User
// @Failure 400 {object} models.ErrorResponse
// @Router /users [post]
func CreateUser(c *gin.Context) {
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": err.Error()})
		return
	}
	res := initializers.DB.Create(&user)
	if res.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"Error": res.Error})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}

// GetUsers godoc
// @Summary Get a list of users
// @Description Retrieve all users from the database
// @Tags users
// @Produce json
// @Success 200 {array} models.User
// @Failure 500 {object} models.ErrorResponse
// @Router /users [get]

var cac = cache.New(5*time.Minute, 10*time.Minute)

func GetUsers(c *gin.Context) {
	var users []*models.User
	startTime := time.Now()

	cachedUsers, found := cac.Get("all_users")
	if found {
		users = cachedUsers.([]*models.User)
		log.Println("[CACHE HIT] Fetched all users from cache")
		return
	}
	// c.Set("all_users", users, cache.DefaultExpiration)
	elapsedTime := time.Since(startTime)
	log.Printf("[SUCCESS] Users fetched successfully in %v", elapsedTime)

	initializers.DB.Find(&users)
	cac.Set("all_users", users, cache.DefaultExpiration)
	// c.Set("all_users", users)
	c.JSON(http.StatusOK, gin.H{"data": users})

}
