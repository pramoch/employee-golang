package Employees

import (
	"context"
	DB "employee-golang/db"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type employee struct {
	Id         string             `json:"id" bson:"id"`
	Name       string             `json:"name" bson:"name"`
	Surname    string             `json:"surname" bson:"surname"`
	MobileNo   string             `json:"mobileNo" bson:"mobileNo"`
	Salary     int                `json:"salary" bson:"salary"`
	JoinedDate primitive.DateTime `json:"joinedDate" bson:"joinedDate"`
	PositionId primitive.ObjectID `json:"positionId" bson:"positionId"`
	BranchId   primitive.ObjectID `json:"branchId" bson:"branchId"`
}

type position struct {
	Id   primitive.ObjectID `json:"_id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}

// var e []employee

func Init() {
	// e = []employee{
	// 	{
	// 		Id:       "700101",
	// 		Name:     "Emma",
	// 		Surname:  "Cartner",
	// 		MobileNo: "0815467890",
	// 		Salary:   15000,
	// 		JoinDate: "2017-06-01",
	// 		Position: "Officer",
	// 		Branch:   "London - Oxford Street East",
	// 	},
	// 	{
	// 		Id:       "700102",
	// 		Name:     "Quinn",
	// 		Surname:  "Rivers",
	// 		MobileNo: "0896552310",
	// 		Salary:   30000,
	// 		JoinDate: "2015-10-01",
	// 		Position: "Branch Manager",
	// 		Branch:   "London - Oxford Street East",
	// 	},
	// 	{
	// 		Id:       "700103",
	// 		Name:     "Amelia",
	// 		Surname:  "Burrows",
	// 		MobileNo: "0823335859",
	// 		Salary:   18000,
	// 		JoinDate: "2018-02-01",
	// 		Position: "Officer",
	// 		Branch:   "Manchester - Trafford Centre",
	// 	},
	// 	{
	// 		Id:       "700104",
	// 		Name:     "Jacob",
	// 		Surname:  "Walsh",
	// 		MobileNo: "0864659874",
	// 		Salary:   17000,
	// 		JoinDate: "2015-07-01",
	// 		Position: "Officer",
	// 		Branch:   "Manchester - Trafford Centre",
	// 	},
	// 	{
	// 		Id:       "700105",
	// 		Name:     "Martha",
	// 		Surname:  "Hills",
	// 		MobileNo: "0895557890",
	// 		Salary:   33000,
	// 		JoinDate: "2015-04-01",
	// 		Position: "Branch Manager",
	// 		Branch:   "Manchester - Trafford Centre",
	// 	},
	// 	{
	// 		Id:       "700106",
	// 		Name:     "Tracy",
	// 		Surname:  "Robertson",
	// 		MobileNo: "0831129966",
	// 		Salary:   20000,
	// 		JoinDate: "2016-02-01",
	// 		Position: "Officer",
	// 		Branch:   "York - Coppergate",
	// 	},
	// 	{
	// 		Id:       "700107",
	// 		Name:     "Harvey",
	// 		Surname:  "Longbottom",
	// 		MobileNo: "0845671122",
	// 		Salary:   19000,
	// 		JoinDate: "2017-11-01",
	// 		Position: "Officer",
	// 		Branch:   "York - Coppergate",
	// 	},
	// 	{
	// 		Id:       "700108",
	// 		Name:     "Rinzee",
	// 		Surname:  "Wilma",
	// 		MobileNo: "0894453798",
	// 		Salary:   15000,
	// 		JoinDate: "2014-04-01",
	// 		Position: "Officer",
	// 		Branch:   "York - Coppergate",
	// 	},
	// 	{
	// 		Id:       "700109",
	// 		Name:     "Clark",
	// 		Surname:  "Elsie",
	// 		MobileNo: "0816962525",
	// 		Salary:   18000,
	// 		JoinDate: "2019-12-01",
	// 		Position: "Officer",
	// 		Branch:   "York - Coppergate",
	// 	},
	// 	{
	// 		Id:       "700110",
	// 		Name:     "Walker",
	// 		Surname:  "Kristen",
	// 		MobileNo: "0884662020",
	// 		Salary:   25000,
	// 		JoinDate: "2019-01-01",
	// 		Position: "Branch Manager",
	// 		Branch:   "York - Coppergate",
	// 	},
	// 	{
	// 		Id:       "700111",
	// 		Name:     "Christine",
	// 		Surname:  "Mccarty",
	// 		MobileNo: "0897750223",
	// 		Salary:   26000,
	// 		JoinDate: "2015-03-01",
	// 		Position: "Branch Manager",
	// 		Branch:   "Leeds - Trinity",
	// 	},
	// 	{
	// 		Id:       "700112",
	// 		Name:     "Lacey",
	// 		Surname:  "O\"Ryan",
	// 		MobileNo: "0819896547",
	// 		Salary:   15000,
	// 		JoinDate: "2016-04-01",
	// 		Position: "Officer",
	// 		Branch:   "Leeds - Trinity",
	// 	},
	// 	{
	// 		Id:       "700113",
	// 		Name:     "Troy",
	// 		Surname:  "Bryant",
	// 		MobileNo: "0861131447",
	// 		Salary:   16000,
	// 		JoinDate: "2018-01-01",
	// 		Position: "Officer",
	// 		Branch:   "Leeds - Trinity",
	// 	},
	// 	{
	// 		Id:       "700114",
	// 		Name:     "Ellie",
	// 		Surname:  "Stewart",
	// 		MobileNo: "0640509987",
	// 		Salary:   17000,
	// 		JoinDate: "2018-08-01",
	// 		Position: "Officer",
	// 		Branch:   "Leeds - Trinity",
	// 	},
	// 	{
	// 		Id:       "700115",
	// 		Name:     "Velma",
	// 		Surname:  "Farmer",
	// 		MobileNo: "0663456877",
	// 		Salary:   27000,
	// 		JoinDate: "2017-09-01",
	// 		Position: "Branch Manager",
	// 		Branch:   "Birmingham - Fort Parkway",
	// 	},
	// 	{
	// 		Id:       "700116",
	// 		Name:     "Nadia",
	// 		Surname:  "Pugh",
	// 		MobileNo: "0692524467",
	// 		Salary:   17000,
	// 		JoinDate: "2019-10-01",
	// 		Position: "Officer",
	// 		Branch:   "Birmingham - Fort Parkway",
	// 	},
	// 	{
	// 		Id:       "700117",
	// 		Name:     "Charley",
	// 		Surname:  "Pratt",
	// 		MobileNo: "0816450080",
	// 		Salary:   16000,
	// 		JoinDate: "2019-02-01",
	// 		Position: "Officer",
	// 		Branch:   "Birmingham - Fort Parkway",
	// 	},
	// }
}

func GetEmployees(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := DB.Client.Database("employee").Collection("employees")
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	var employees []employee
	if err := cur.All(ctx, &employees); err != nil {
		log.Fatal(err)
	}

	c.JSON(200, employees)

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}

func GetEmployeeById(c *gin.Context) {
	id := c.Param("id")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := DB.Client.Database("employee").Collection("employees")
	filter := bson.M{"id": id}
	var employee employee
	err := collection.FindOne(ctx, filter).Decode(&employee)

	if err == mongo.ErrNoDocuments {
		c.JSON(404, gin.H{"desc": "Employee not found"})
		return
	} else if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, employee)
}

func GetPositions(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collection := DB.Client.Database("employee").Collection("positions")

	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	var positions []position
	if err := cur.All(ctx, &positions); err != nil {
		log.Fatal(err)
	}

	c.JSON(200, positions)

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}
}
