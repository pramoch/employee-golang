package Branches

import (
	"context"
	DB "employee-golang/db"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type branch struct {
	Id      primitive.ObjectID `json:"_id" bson:"_id"`
	Name    string             `json:"name" bson:"name"`
	TelNo   string             `json:"telNo" bson:"telNo"`
	Address string             `json:"address" bson:"address"`
	Map     string             `json:"map" bson:"map"`
}

type branchesResultStatus struct {
	Success bool   `json:"success"`
	Desc    string `json:"desc"`
}

type branchesResultData struct {
	Branches []branch `json:"branches"`
}

type branchesResult struct {
	Status branchesResultStatus `json:"status"`
	Data   branchesResultData   `json:"data"`
}

func GetBranches(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := DB.Client.Database("employee").Collection("branches")

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	var branches []branch
	if err := cur.All(ctx, &branches); err != nil {
		log.Fatal(err)
	}

	result := branchesResult{
		Status: branchesResultStatus{
			Success: true,
			Desc:    "Success",
		},
		Data: branchesResultData{
			Branches: branches,
		},
	}
	c.JSON(200, result)

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
