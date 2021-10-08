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

// bson:"_id,omitempty" - omitempty to let MongoDB automatically generate _ID
type branch struct {
	Id      primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Name    string             `json:"name,omitempty" bson:"name" binding:"required"`
	TelNo   string             `json:"telNo,omitempty" bson:"telNo"`
	Address string             `json:"address,omitempty" bson:"address"`
	Map     string             `json:"map,omitempty" bson:"map"`
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

// Use pointer for branchResultData,
//   when data is not provided (Branch not found)
//   data field will be omitted from JSON
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
		log.Println(err)
	}
	defer cur.Close(ctx)

	branches := []branch{}
	if err := cur.All(ctx, &branches); err != nil {
		log.Println(err)
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
		log.Println(err)
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
		log.Println(err)
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
	var b branch

	if err := c.ShouldBindJSON(&b); err != nil {
		log.Println(err)

		errMsg := "Invalid input"
		if b.Name == "" {
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

	fmt.Println(b)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := DB.Client.Database("employee").Collection("branches")
	res, err := collection.InsertOne(ctx, b)

	if err != nil {
		log.Println(err)

		result := branchResult{
			Status: status{
				Success: false,
				Desc:    "Cannot add new branch",
			},
		}

		c.JSON(500, result)
		return
	}

	result := branchResult{
		Status: status{
			Success: true,
			Desc:    "Success",
		},
		Data: &branchResultData{
			Branch: branch{
				Id: res.InsertedID.(primitive.ObjectID),
			},
		},
	}

	c.JSON(200, result)
}
