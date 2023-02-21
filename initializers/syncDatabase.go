package initializers

import "github.com/rudiath95/JWT-training/models"

func SyncDatabases() {
	DB.AutoMigrate(models.User{})
}
