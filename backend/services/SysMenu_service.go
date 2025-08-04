package services

import (
	//"context"
//	"log"

//	"golang-fixedasset-mongo-backend/backend/models"
	"golang-fixedasset-mongo-backend/backend/tools"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/mongo"
//	"go.mongodb.org/mongo-driver/bson/primitive"
	"sort"
)

func SortTree(nodes []*tools.TreeNode) {
	sort.Slice(nodes, func(i, j int) bool {
		return nodes[i].Sort < nodes[j].Sort
	})
	for _, node := range nodes {
		if len(node.Childrens) > 0 {
			SortTree(node.Childrens)
		}
	}
}
/* 
func GetTreeAllMenuByIds(collection *mongo.Collection, ids []string) ([]*tools.TreeNode, error) {
	ctx := context.Background()

	// 將 string 轉換成 ObjectID
	var objIDs []primitive.ObjectID
	for _, id := range ids {
		oid, err := primitive.ObjectIDFromHex(id)
		if err == nil {
			objIDs = append(objIDs, oid)
		}
	}

	// 查找 id in ids 或 mainId in ids
	filter := bson.M{
		"$or": []bson.M{
			{"_id": bson.M{"$in": objIDs}},
			{"mainId": bson.M{"$in": objIDs}},
		},
		"status": 1,
	}

	cur, err := collection.Find(ctx, filter)
	if err != nil {
		log.Println("Find error:", err)
		return nil, err
	}
	defer cur.Close(ctx)

	var menus []models.SysMenu
	if err := cur.All(ctx, &menus); err != nil {
		return nil, err
	}

	nodeMap := make(map[string]*tools.TreeNode)
	for _, m := range menus {
		node := &tools.TreeNode{
			ID:       m.ID, // ObjectID -> string
			Name:     m.Name,
			MainID:   m.MainId, // ObjectID -> string
			Childrens: []*tools.TreeNode{},
		}
		nodeMap[node.ID] = node
	}

	var rootNodes []*tools.TreeNode
	for _, node := range nodeMap {
		if node.MainID == "" || nodeMap[node.MainID] == nil {
			rootNodes = append(rootNodes, node)
		} else {
			parent := nodeMap[node.MainID]
			parent.Childrens = append(parent.Childrens, node)
		}
	}

	SortTree(rootNodes)
	return rootNodes, nil
}
*/




