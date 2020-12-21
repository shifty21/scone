package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/yaml.v2"
)

//DBConfig struct
type DBConfig struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
}

func loadConfig() (*DBConfig, error) {
	yamlFile, err := ioutil.ReadFile("resources/consul-template/templates/config.yml")
	if err != nil {
		return nil, fmt.Errorf("Error getting db config %w ", err)
	}
	var config DBConfig
	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		return nil, fmt.Errorf("Error Unmarshalling config %w", err)
	}
	return &config, nil

}

func main() {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("Error while loading config %v", err)
		os.Exit(1)
	}
	fmt.Printf("Config %v\n", config)
	url := fmt.Sprintf("mongodb://%v:%v@localhost:27017/%v", config.UserName, config.Password, config.Database)
	// url := "mongodb://localhost:27017/"
	fmt.Printf("URI %v\n", url)
	client, err := mongo.NewClient(options.Client().ApplyURI(url))
	if err != nil {
		log.Printf("Error while Creating new client for mongodb %v", err)
		os.Exit(1)
	}
	fmt.Printf("Created client\n")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatalf("Error while connecting to db %v", err)
		os.Exit(1)
	}
	err = client.Ping(context.TODO(), nil)

	if err != nil {
		log.Fatalf("Error while pinging to db %v", err)
	}

	fmt.Println("Connected to MongoDB!")

	defer client.Disconnect(ctx)
	os.Exit(0)
}
