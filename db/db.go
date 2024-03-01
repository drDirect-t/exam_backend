package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// กำหนดข้อมูลการเชื่อมต่อ MongoDB
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// เชื่อมต่อกับ MongoDB
	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// ตรวจสอบการเชื่อมต่อ MongoDB
	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to MongoDB!")

	// // ยกเลิกการเชื่อมต่อ
	// err = client.Disconnect(context.Background())
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println("Disconnected from MongoDB!")
}
