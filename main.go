package main

import (
	"github.com/micro/go-micro"
	"log"
	"github.com/micro/go-micro/broker"
	pb "github.com/chauhanr/shipcon-user-service/proto/user"
	"encoding/json"
)

const topic = "user.created"

func main(){
	srv := micro.NewService(
		micro.Name("go.micro.srv.email"),
		micro.Version("lastest"),
	)
	srv.Init()

	pubsub := srv.Server().Options().Broker
	if err := pubsub.Connect(); err != nil {
		log.Fatal(err)
	}

	// subscribe now
	_, err := pubsub.Subscribe(topic, func(p broker.Publication) error{
		var user *pb.User
		if err := json.Unmarshal(p.Message().Body, &user); err != nil{
			return err
		}
		log.Println(user)
		go sendEmail(user)
		return nil
	})

	if err != nil{
		log.Println(err)
	}
	if err := srv.Run(); err != nil{
		log.Println(err)
	}
}

func sendEmail(user *pb.User) error{
	log.Println("Sending email to: ", user.Name)
	return nil
}