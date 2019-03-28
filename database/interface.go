package database

import (
	"kim/config"
	"kim/database/schema"
)

type db interface {
	SaveMicroserviceDetail(schema.MicroserviceDetail) bool
}

// SaveMicroserviceDetail saves the new microservice and its information
// in the database
func SaveMicroserviceDetail(config config.Configstructure, msDetail *schema.MicroserviceDetail) bool {
	switch dbEngine := config.DatabaseEngine; dbEngine {
	case "mongo":
		return saveMicroserviceDetail(config, msDetail)
	}

	return true
}
