package services

import (
	"golang-fixedasset-mongo-backend/backend/config"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"time"
)

func AddLoginRecord(username string, ip string, loginStatus string) (interface{}, error) {
	collection := config.GetCollection("login_records")

	saveData := bson.M{
		"username":  username,
		"ipAddress": ip,
		"loginStatus": loginStatus,
		"loginTime":  time.Now(),
	}

	res, err := collection.InsertOne(context.Background(), saveData)


	if err != nil {		
		return nil, err
	} else {
		return res, nil
	}
	
}