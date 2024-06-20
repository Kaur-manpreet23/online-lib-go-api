package db
import (
	"os"
	"fmt"
	"database/sql"
	"log"
	"github.com/go-sql-driver/mysql"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
    "context"
    "time"

)
var stmt *sql.Stmt
var Db *sql.DB
var Cli *mongo.Client

func Configure(){
	dbUser := os.Getenv("DBUSER")
    dbPassword := os.Getenv("DBPASS")

    if dbUser == "" && dbPassword == "" {
	    log.Fatal("Environment variables are not set")
	    os.Exit(1)
    }

    cfg := mysql.Config{
        User:   dbUser,
        Passwd: dbPassword,
        Net:    "tcp",
        Addr:   "127.0.0.1:3306",
        DBName: "online_library_development",
    }

    // Get a database handle.
    var err error
    Db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }
    if Db == nil {
	    fmt.Println("error")
		return
    }
    pingErr := Db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
	return 	
    }
    //coll := client.Database("online_library_development_logs").Collection("logs")
    uri := "mongodb://localhost:27017/?timeoutMS=5000"
    Cli, err = mongo.Connect(context.Background(), options.Client().ApplyURI(uri))
    if err != nil {
        log.Fatal("Error connecting to MongoDB: %v", err)
        return
    }
    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()
    if err := Cli.Ping(ctx, nil); err != nil {
        log.Fatalf("Error pinging MongoDB server: %v", err)
        return
    }

    log.Println("Connected to MongoDB!")

}