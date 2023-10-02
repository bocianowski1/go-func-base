package main

import (
	"context"
	"fmt"
	f "go-base/function"
	"log"
	"os"
)

func main() {
	listenAddr := ":8080"
	if val, ok := os.LookupEnv("FUNCTIONS_CUSTOMHANDLER_PORT"); ok {
		listenAddr = ":" + val
	}

	poolSize := 8

	mongoConfig := f.MongoDBConfig{
		URI:      "mongodb://localhost:27017",
		Database: "db",
	}

	db, err := f.NewMongoDB(mongoConfig)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Client.Disconnect(context.Background())

	if err := db.Init(); err != nil {
		log.Fatal("Initialization failed:", err)
	}

	fmt.Println("Connected to MongoDB")

	server := f.NewServer(listenAddr, db)
	workerPool := f.NewWorkerPool(poolSize)
	defer workerPool.Shutdown()

	server.Start([]f.HandlerFuncPair{
		{
			Route:   fmt.Sprintf("/api/%v", functionFolder()),
			Handler: f.MakeHTTPHandler(server.HandleUsers, workerPool),
		},
	})
}

func functionFolder() string {
	folders, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	if len(folders) == 0 {
		log.Fatal("no function folder found")
	}

	var name string
	for _, f := range folders {
		if f.IsDir() && f.Name() != "terraform" {
			name = f.Name()
		}
	}

	return name
}
