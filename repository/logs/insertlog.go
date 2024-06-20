package logs

import (
	"context"
    "go.mongodb.org/mongo-driver/bson"
    "log"
	"os"
    "time"
	"strconv"
	"goborrow/repository/db"
)

func InsertLog(action string){
	coll := (db.Cli).Database("online_library_development_logs").Collection("logs")
	now := time.Now()
	id := os.Getenv("USERID")
	user_id, err := strconv.Atoi(id)
	//fmt.Println(intVar, err, reflect.TypeOf(intVar))

    doc := bson.D{
        {"action", action},
        {"user_id", user_id},
        {"timestamp", now},
    }
	_, err = coll.InsertOne(context.Background(), doc)
    if err != nil {
        log.Fatalf("Error inserting log: %v", err)
    }
}