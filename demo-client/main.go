package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//DBConfig struct
type DBConfig struct {
	UserName string `yaml:"username"`
	Password string `yaml:"password"`
	Database string `yaml:"database"`
	Address  string `yaml:"address"`
	CertFile string `yami:"cert_file"`
	KeyFile  string `yami:"key_file"`
}

func loadConfig(filepath string, watcher chan *DBConfig) (*DBConfig, error) {
	log.Printf("Reading config from [%v]", filepath)
	viper.AddConfigPath(filepath)
	viper.SetConfigType("yml")
	viper.SetConfigName("config")
	viper.ReadInConfig()
	go func() {
		for {
			time.Sleep(time.Second * 30)
			viper.WatchConfig()
			viper.OnConfigChange(
				func(e fsnotify.Event) {
					config := &DBConfig{}
					viper.ReadInConfig()
					err := viper.Unmarshal(config)
					if err != nil {
						log.Printf("Error unmarshalling config")
					}
					watcher <- config
				},
			)
		}
	}()
	config := &DBConfig{}
	err := viper.Unmarshal(config)
	if err != nil {
		log.Printf("Error unmarshalling config")
		return nil, err
	}
	return config, nil

}

func main() {
	SignalChan := make(chan os.Signal)
	signal.Notify(SignalChan,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGKILL,
	)
	watcher := make(chan *DBConfig)
	filePath := os.Args[1:][0]
	config, err := loadConfig(filePath, watcher)
	if err != nil {
		log.Fatalf("Error while loading config %v", err)
		os.Exit(1)
	}
	fmt.Printf("Config %v\n", config)
	var client *mongo.Client
	var ctx context.Context
	var url string
	go func() {
		for {
			select {
			case config = <-watcher:
				fmt.Printf("Config Change event reloading connection %v", config)
				url = fmt.Sprintf("mongodb://%v:%v@%v:27017/%v", config.UserName, config.Password, config.Address, config.Database)
				fmt.Printf("URI %v\n", url)
				client, err = mongo.NewClient(options.Client().ApplyURI(url))
				if err != nil {
					log.Printf("Error while Creating new client for mongodb %v", err)
				}
				ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
				err = client.Connect(ctx)
				if err != nil {
					log.Printf("Error while connecting to db %v", err)
				}
				err = client.Ping(context.TODO(), nil)

				if err != nil {
					log.Printf("Error while pinging to db %v", err)
				} else {
					fmt.Println("Connected to MongoDB!")
				}
				client.Disconnect(ctx)
			default:
				url = fmt.Sprintf("mongodb://%v:%v@%v:27017/%v", config.UserName, config.Password, config.Address, config.Database)
				fmt.Printf("URI %v\n", url)
				client, err = mongo.NewClient(options.Client().ApplyURI(url))
				if err != nil {
					log.Printf("Error while Creating new client for mongodb %v", err)
				}
				ctx, _ = context.WithTimeout(context.Background(), 10*time.Second)
				err = client.Connect(ctx)
				if err != nil {
					log.Printf("Error while connecting to db %v", err)
				}
				err = client.Ping(context.TODO(), nil)
				if err != nil {
					log.Printf("Error while pinging to db %v", err)
				} else {
					fmt.Println("Connected to MongoDB!")
				}

				client.Disconnect(ctx)
			}
			time.Sleep(30 * time.Second)

		}
	}()
	<-SignalChan
	fmt.Println("Interrupt encountered, exiting")
	os.Exit(0)
}
