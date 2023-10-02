package main

import (
	"context"
	"walls-user-service/internal/adapter/events/subscriber"
	extensions "walls-user-service/internal/adapter/extensions"
	mongoRepository "walls-user-service/internal/adapter/repository/mongodb"

	"fmt"
	"walls-user-service/internal/adapter/routes"
	channel "walls-user-service/internal/core/domain/event/channel"
	configuration "walls-user-service/internal/core/helper/configuration-helper"
	logger "walls-user-service/internal/core/helper/log-helper"
	message "walls-user-service/internal/core/helper/message-helper"
)

func main() {
	//Initialize request Log
	logger.InitializeLog()
	//Start DB Connection
	mongoRepo := extensions.StartDatabase("mongodb")

	logger.LogEvent("INFO", "MongoDB Connected and Initialized!")

	logger.LogEvent("INFO", message.StartingRedis)
	redisClient := extensions.StartEventBus("redis")
	ctx := context.Background()

	//Set up routes
	router := routes.SetupRouter(mongoRepo.(mongoRepository.MongoRepositories).User, redisClient)

	config := configuration.ServiceConfiguration

	go func() {
		logger.LogEvent("INFO", message.StartingServer)
		err := router.Run(":" + config.ServicePort)
		//api.SetConfiguration
		if err != nil {
			fmt.Println(err)
			logger.LogEvent("ERROR", "Error Starting Server : "+err.Error())
		}
	}()

	// Initialize the event subscriber
	eventSubscriber := subscriber.NewEventSubscriber(redisClient)
	// Run the subscription code in a Goroutine
	go func() {
		eventSubscriber.SubscribeToOtpValidatedEvent(ctx, channel.OtpValidatedEvent)
	}()

	select {}
}
