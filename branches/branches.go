package Branches

import (
	"context"
	DB "employee-golang/db"
	"fmt"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type branch struct {
	Id      primitive.ObjectID `json:"_id" bson:"_id"`
	Name    string             `json:"name" bson:"name" binding:"required"`
	TelNo   string             `json:"telNo" bson:"telNo"`
	Address string             `json:"address" bson:"address"`
	Map     string             `json:"map" bson:"map"`
}

type status struct {
	Success bool   `json:"success"`
	Desc    string `json:"desc"`
}

// GetBranches
type branchesResultData struct {
	Branches []branch `json:"branches"`
}

type branchesResult struct {
	Status status             `json:"status"`
	Data   branchesResultData `json:"data"`
}

// GetBranchById
type branchResultData struct {
	Branch branch `json:"branch"`
}

type branchResult struct {
	Status status            `json:"status"`
	Data   *branchResultData `json:"data,omitempty"`
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

	branches := []branch{}
	if err := cur.All(ctx, &branches); err != nil {
		log.Fatal(err)
	}

	result := branchesResult{
		Status: status{
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

func GetBranchById(c *gin.Context) {
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := DB.Client.Database("employee").Collection("branches")
	docID, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{"_id": docID}
	var branch branch

	err := collection.FindOne(ctx, filter).Decode(&branch)

	if err == mongo.ErrNoDocuments {
		result := branchResult{
			Status: status{
				Success: false,
				Desc:    "Branch not found",
			},
		}
		c.JSON(404, result)
		return
	} else if err != nil {
		log.Fatal(err)
	}

	data := branchResultData{
		Branch: branch,
	}

	result := branchResult{
		Status: status{
			Success: true,
			Desc:    "Success",
		},
		Data: &data,
	}
	c.JSON(200, result)
}

func AddBranch(c *gin.Context) {
	var branch branch

	if err := c.ShouldBindJSON(&branch); err != nil {
		errMsg := "Invalid input"

		if branch.Name == "" {
			errMsg = "name is required"
		}

		result := branchResult{
			Status: status{
				Success: false,
				Desc:    errMsg,
			},
		}

		c.JSON(400, result)
		return
	}

	fmt.Println("== Branch ==")
	fmt.Println(branch)
	fmt.Println("Id: ", branch.Id)
	fmt.Println("Name: ", branch.Name)
	fmt.Println("TelNo: ", branch.TelNo)
	fmt.Println("Address: ", branch.Address)
	fmt.Println("Map: ", branch.Map)

	result := branchResult{
		Status: status{
			Success: true,
			Desc:    "Success",
		},
	}

	c.JSON(200, result)
}
