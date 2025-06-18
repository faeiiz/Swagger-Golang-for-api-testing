package migrate

import (
	"swag/initializers"
	"swag/models"
)

func Migrate() {
	// Add your migration code here
	initializers.DB.AutoMigrate(&models.User{})
}
