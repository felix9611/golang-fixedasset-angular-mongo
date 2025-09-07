package services

import (
	"context"
	"golang-fixedasset-mongo-backend/backend/config"
	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/dto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	
	"go.mongodb.org/mongo-driver/mongo"
	"time"
	"fmt"
	"sort"
)

func CreateSysMenu(data *models.SysMenus) (interface{}, error) {
	filter := bson.M{"status": 1, "name": data.Name, "path": data.Path}

	collection := config.GetCollection("sys_menus")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)

	defer cancel()

	count, err := collection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if data.Status == 0 {
			data.Status = 1
		}

		if data.CreatedAt.IsZero() {
			data.CreatedAt = time.Now()
		}

		if data.UpdatedAt.IsZero() {		
			data.UpdatedAt = time.Now()
		}
		result, err := collection.InsertOne(ctx, data)
		if err != nil {
			return nil, err
		}

		CreateActionRecord("Menu Item Create", "POST", "System Menu", data, "Success")

		return result.InsertedID, nil
	} else {
		CreateActionRecord("Menu Item Create", "POST", "System Menu", data, "Failed")
		return "Menu with the same name already exists", nil
	}
	
}

func GetOneMenuItemById(id string) (interface{}, error) {
	var menuItem models.SysMenus
	collection := config.GetCollection("sys_menus")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return menuItem, err
	}

	filter := bson.M{"_id": objectID, "status": 1}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err = collection.FindOne(ctx, filter).Decode(&menuItem)
	if err != nil {
		return nil, err
	}

	return menuItem, nil
}

func VoidMenuItemById(id string) (interface{}, error) {
	collection := config.GetCollection("sys_menus")

	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}

	filter := bson.M{"_id": objectID}

	var menuItem models.SysMenus

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = collection.FindOne(ctx, filter).Decode(&menuItem)

	if err != nil {
		return nil, err
	}

	if menuItem.Status == 0 {
		CreateActionRecord("Menu Item Void", "DELETE", "System Menu", menuItem, "Failed")
		return "Menu item already voided", nil
	} else {

		menuItem.Status = 0
		menuItem.UpdatedAt = time.Now()

		result, err := collection.UpdateOne(ctx, filter, bson.M{
			"$set": menuItem,
		})

		if err != nil {
			return nil, err
		}

		CreateActionRecord("Menu Item Void", "DELETE", "System Menu", menuItem, "Success")
		
		return result, nil
	}
}

func UpdateMenuItem(data *models.SysMenus) (interface{}, error) {
	collection := config.GetCollection("sys_menus")

	filter := bson.M{"_id": data.ID}

	var menuItem models.SysMenus
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err := collection.FindOne(ctx, filter).Decode(&menuItem)
	if err != nil {
		return nil, err
	}

	if menuItem.Status == 0 {
		CreateActionRecord("Menu Item Update", "POST", "System Menu", data, "Failed")
		return "This Menu Item is inactive", nil
	} else {
		data.UpdatedAt = time.Now()

		result, err := collection.UpdateOne(ctx, filter, bson.M{"$set": data})

		CreateActionRecord("Menu Item Update", "POST", "System Menu", data, "Success")

		if err != nil {
			return nil, err
		}
		return result, nil
	}
}

func ListAllMainIdMenu() (interface{}, error) {
	collection := config.GetCollection("sys_menus")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Aggregation pipeline
	pipeline := mongo.Pipeline{
		// $match: mainId = ""
		{{Key: "$match", Value: bson.M{"mainId": ""}}},
		// $group by mainId + name, keep first _id
		{{Key: "$group", Value: bson.M{
			"_id": bson.M{
				"mainId": "$mainId",
				"name":   "$name",
			},
			"_idValue": bson.M{"$first": "$_id"},
		}}},
		// $project final fields
		{{Key: "$project", Value: bson.M{
			"_id":    "$_idValue",
			"mainId": "$_id.mainId",
			"name":   "$_id.name",
		}}},
	}

	cursor, err := collection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var results []dto.SysMenuMainId
	if err := cursor.All(ctx, &results); err != nil {
		return nil, err
	}

	return results, nil
}

// ListAllMenu fetch menus by name and build hierarchical tree
func ListAllMenu(query *dto.SysMenuList) (interface{}, error) {
	collection := config.GetCollection("sys_menus")

	filter := bson.M{}
	if query.Name != "" {
		filter["name"] = bson.M{"$regex": primitive.Regex{Pattern: fmt.Sprintf(".*%s.*", query.Name), Options: "i"}}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	cursor, err := collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var menus []dto.SysMenuChildrens
	if err := cursor.All(ctx, &menus); err != nil {
		return nil, err
	}

	// Build hierarchical tree
	tree := BuildSortedTree(menus)
	return tree, nil
}

// BuildSortedTree builds a hierarchical and sorted tree from flat menu list
func BuildSortedTree(data []dto.SysMenuChildrens) []*dto.SysMenuChildrens {
	nodeMap := make(map[string]*dto.SysMenuChildrens)
	tree := []*dto.SysMenuChildrens{}

	// 將資料全部轉成 pointer，初始化 Childrens
	for i := range data {
		item := &data[i]
		if item.Childrens == nil {
			item.Childrens = []dto.SysMenuChildrens{}
		}
		nodeMap[item.ID.Hex()] = item
	}

	// Link children
	for _, item := range data {
		current := nodeMap[item.ID.Hex()]
		if item.MainId != "" {
			if parent, exists := nodeMap[item.MainId]; exists {
				parent.Childrens = append(parent.Childrens, *current) // append copy 但已經正確 build
			} else {
				tree = append(tree, current)
			}
		} else {
			tree = append(tree, current)
		}
	}

	// Recursive sort
	var sortTree func(nodes []*dto.SysMenuChildrens)
	sortTree = func(nodes []*dto.SysMenuChildrens) {
		sort.Slice(nodes, func(i, j int) bool {
			return nodes[i].Sort < nodes[j].Sort
		})
		for _, node := range nodes {
			if node.Childrens == nil {
				node.Childrens = []dto.SysMenuChildrens{}
			}
			// 將子節點轉 pointer 方便遞迴
			childPtrs := make([]*dto.SysMenuChildrens, len(node.Childrens))
			for i := range node.Childrens {
				childPtrs[i] = &node.Childrens[i]
			}
			sortTree(childPtrs)
		}
	}

	sortTree(tree)

	return tree
}















