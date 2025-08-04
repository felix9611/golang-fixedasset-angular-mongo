package services

import (
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/config"
	"golang-fixedasset-mongo-backend/backend/tools"
	"time"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"log"
)

func CreateSysUser(user *models.SysUsers) (interface{}, error) {
	log.Printf("Insert SysUser: %+v\n", user)

	filter := bson.M{"status": 1, "username": user.Username }

	collection := config.GetCollection("sys_users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := collection.CountDocuments(ctx, filter)

	if err != nil {
		return nil, err
	}

	if count == 0 {

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if user.Status == 0 {
			user.Status = 1
		}

		if user.CreatedAt.IsZero() {
			user.CreatedAt = time.Now()
		}

		if user.UpdatedAt.IsZero() {
			user.UpdatedAt = time.Now()
		}

		if user.Password == "" {
			user.Password = tools.HashPassword("888888", tools.Salt)
		}

		result, err := collection.InsertOne(ctx, user)
		if err != nil {
			return nil, err
		}

		return result, nil
	} else {
		return "User with the same username already exists", nil
	}
}

func UpdateSysUserByID(id string, updateData *models.SysUsers) (interface{}, error) {
	collection := config.GetCollection("sys_users")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID, "status": 1}
	
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var checkSysUser models.SysUsers
	err = collection.FindOne(ctx, filter).Decode(&checkSysUser)
	if err != nil {
		return nil, err
	}

	if checkSysUser.Status == 1 {

		if updateData.UpdatedAt.IsZero() {
			updateData.UpdatedAt = time.Now()
		}

		result, err := collection.UpdateOne(ctx, filter, bson.M{"$set": updateData})
		if err != nil {
			return nil, err
		}
		return result, nil
	} else {
		return "User not found or inactive", nil
	}
}

func InactiveUserByID(id string) (interface{}, error) {
	collection := config.GetCollection("sys_users")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID, "status": 1}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.SysUsers

	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	if user.Status == 1 {
		result, err := collection.UpdateOne(ctx, filter, bson.M{
			"$set": bson.M{
				"status": 0,
				"updated_at": time.Now(),
			},
		})

		if err != nil {
			return nil, err
		}
		return result, nil
	} else {
		return "User is already inactive", nil
	}
}

func GetOneSysUserById(id string) (interface{}, error) {
	collection := config.GetCollection("sys_users")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{ "_id": objectID, "status": 1 }
	var user models.SysUsers
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	
	if user.Status == 0 {
		return "User is inactive", nil
	} else {
		return &user, nil
	}
}

func GetUserByUsername(username string) (*models.SysUsers, error) {
	collection := config.GetCollection("sys_users")
	filter := bson.M{"username": username, "status": 1}

	var user models.SysUsers
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func UpdateUserPassword(id string, newPassword string) (interface{}, error) {
	collection := config.GetCollection("sys_users")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID, "status": 1}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var user models.SysUsers
	err = collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	if user.Status == 1 {
		result, err := collection.UpdateOne(ctx, filter, bson.M{
			"$set": bson.M{
				"password": tools.HashPassword(newPassword, tools.Salt),
				"updated_at": time.Now(),
			},
		})

		if err != nil {
			return nil, err
		}
		return result.ModifiedCount, nil
	} else {
		return "User not found or inactive", nil
	}
}