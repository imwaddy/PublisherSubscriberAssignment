package dbhelpers

import (
	"PublisherSubscriberAssignment/ConsumerModules/helpers/confighelpers"
	"PublisherSubscriberAssignment/ConsumerModules/models"
	"sync"

	"gorm.io/driver/mysql"

	"gorm.io/gorm"
)

// Hold a single global connection (pooling provided by ORM driver)
var sqlConnection *gorm.DB
var connectionError error
var sqlOnce sync.Once

//GetConnection to db
func GetConnection() (*gorm.DB, error) {

	sqlOnce.Do(func() {
		connection, err := gorm.Open(mysql.Open(confighelpers.GetConfig("DataSource")), &gorm.Config{})
		if err != nil {
			connectionError = err
		}

		connection.AutoMigrate(&models.Hotel{})
		connection.AutoMigrate(&models.Room{})
		connection.AutoMigrate(&models.RatePlan{})

		sqlConnection = connection
	})

	return sqlConnection, connectionError
}
