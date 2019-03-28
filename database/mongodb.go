package database

import (
	"kim/database/schema"
	"kim/util"

	"kim/config"

	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
)

type dialInfo struct {
	Addrs    []string
	Timeout  int
	Database string
	Username string
	Password string
}

func getDialInfo(config config.Configstructure) mgo.DialInfo {
	infraDb := config.DefaultDatabaseName
	mongoAddress := config.DatabaseConfig["HOST"] + ":" + config.DatabaseConfig["PORT"]

	return mgo.DialInfo{
		Addrs:    []string{mongoAddress},
		Database: infraDb,
		Username: config.DatabaseConfig["USERNAME"],
		Password: config.DatabaseConfig["PASSWORD"],
	}
}

func saveMicroserviceDetail(config config.Configstructure, msDetail *schema.MicroserviceDetail) bool {
	dialInfo := getDialInfo(config)
	colName := "microservice_details"

	session, error := mgo.DialWithInfo(&dialInfo)
	defer session.Close()

	col := session.DB(config.DefaultDatabaseName).C(colName)

	var query schema.MicroserviceDetail

	error = col.Find(bson.M{"microservicename": msDetail.MicroserviceName}).One(&query)

	if (query == schema.MicroserviceDetail{}) {
		error = col.Insert(&msDetail)

		if error != nil {
			util.Throw("Unable to insert in the database")
		}

		return true
	}

	util.FatalPrint("Microservice already exists with name\t" + query.MicroserviceName)

	return false
}
