package main

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog/log"
)

var ctx = context.Background()
var client *redis.Client


func init() {
	log.Info().Msg("init started")

	client = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
		Password: "",
		DB: 0,
	})
}

// @title							Receipt Processor
// @version						1.0
// @description				The Receipt Processor API provides endpoints to...
// @termsOfService		http://swagger.io/terms/
//
// @contact.name			Zachary Broskey
// @contact.email			zbroskey@me.com
//
// @host							localhost:8892
//
//

func main() {
	log.Info().Msg("setup started")

}