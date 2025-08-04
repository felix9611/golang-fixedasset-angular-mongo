package services

import (
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/config"
	"time"
)



func CreateSysRole(role *models.SysRole) (interface{}, error) {
	collection := config.GetCollection("sys_roles")

	filter := bson.M{"status": 1, "name": role.Name}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	count, err := collection.CountDocuments(ctx, filter)

	if err != nil {
		return nil, err
	}

	if count == 0 {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if role.Status == 0 {
			role.Status = 1
		}

		if role.CreatedAt.IsZero() {
			role.CreatedAt = time.Now()
		}

		if role.UpdatedAt.IsZero() {
			role.UpdatedAt = time.Now()
		}

		result, err := collection.InsertOne(ctx, user)
		if err != nil {
			return nil, err
		}

		return result, nil
	} else {
		return "Role with the same name already exists", nil
	}
}

func GetAllRoles() (interface{}, error) {
	collection := config.GetCollection("sys_roles")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var roles []*models.SysRole	

	cursor, err := collection.Find(ctx, bson.M{"status": 1})
	if err != nil {
		return nil, err
	}

	for cursor.Next(ctx) {
		var role models.SysRole
		err := cursor.Decode(&role)
		if err != nil {
			return nil, err
		}
		roles = append(roles, &role)

	}

	return roles, nil
}

func GetOneSysRoleById(id string) (*models.SysUsers, error) {

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := config.GetCollection("sys_users")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	filter := bson.M{"_id": objectID, "status": 1}
	var role models.SysRole
	err = collection.FindOne(ctx, filter).Decode(&role)
	if err != nil {
		return nil, err
	}

	if role.Status == 0 {
		return nil, errors.New("Role is inactive")
	} else {
		return &role, nil
	}
}

func UpdateRoleById(id string, updateData *models.SysRole) (interface{}, error) {
	collection := config.GetCollection("sys_roles")
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID, "status": 1}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var checkSysRole models.SysRole
	err = collection.FindOne(ctx, filter).Decode(&checkSysRole)
	if err != nil {
		return nil, err
	}

	if checkDept.Status == 1 {

		updateData.UpdatedAt = time.Now()
		update := bson.M{"$set": updateData}

		result, err := collection.UpdateOne(ctx, filter, update)
		if err != nil {
			return nil, err
		}

		return result.ModifiedCount, nil
	} else {
		return "Role is already voided", nil
	}
}

func VoidRoleById(id string) (interface{}, error) {
	collection := config.GetCollection("sys_roles")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	var role models.SysRole

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, filter).Decode(&role)

	if err != nil {
		return nil, err
	}

	if role.Status == 1 {

		result, err := collection.UpdateOne(ctx, filter, bson.M{
			"$set": bson.M{
				"status": 0,
				"updated_at": time.Now(),
			},
		})

		if err != nil {
			return nil, err
		}
		return result.ModifiedCount, nil
	} else {
		return "Role is already voided", nil
	}
}

func RolesList(pageDto *dto.RolePageDto) (interface{}, error) {

	if pageDto.Page < 1 {
		pageDto.Page = 1
	}
	if pageDto.Limit < 1 {
		pageDto.Limit = 10
	}

	skip := (pageDto.Page - 1) * pageDto.Limit
	limit := pageDto.Limit

	collection := config.GetCollection("sys_roles")
	filter := bson.M{"status": 1}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if pageDto.Name != "" {
		filter["name"] = bson.M{"$regex": pageDto.Name, "$options": "i"}
	}

	if pageDto.Code != "" {
		filter["code"] = bson.M{"$regex": pageDto.Code, "$options": "i"}
	}

	count, errCount := collection.CountDocuments(ctx, filter)
	if errCount != nil {
		return nil, errCount
	}

	findOptions := options.Find()
	findOptions.SetSkip(int64(skip))
	findOptions.SetLimit(int64(limit))
	findOptions.SetSort(bson.D{{"created_at", -1}})
	cursor, err := collection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, err
	}
	var roles []models.SysRole
	for cursor.Next(ctx) {
		var role models.SysRole
		if err := cursor.Decode(&role); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if err := cursor.Err(); err != nil {
        return nil, err
    }


	return gin.H{"lists": roles, "total": count, "page": pageDto.Page, "limit": pageDto.Limit }, nil
}
