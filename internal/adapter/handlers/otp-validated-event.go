package handlers

import (
	"context"
	// "strings"

	"fmt"
	extraction "walls-user-service/internal/adapter/handlers/extraction"
	"walls-user-service/internal/core/domain/dto"
	// "walls-user-service/internal/core/domain/event/channel"
	events "walls-user-service/internal/core/domain/event/data"
	"walls-user-service/internal/core/helper/event-helper/eto"
	logger "walls-user-service/internal/core/helper/log-helper"
	"walls-user-service/internal/core/services"
)

func OtpValidatedEventHandler(ctx context.Context, event interface{}) {
	event, data, err := extraction.ExtractEventData(event, events.OtpValidatedEventData{})
	if err != nil {
		fmt.Println("extracting event:", err)
		return
	}

	// iEvent := event.(eto.Event)
	// iEventData := data.(events.OtpValidatedEventData)
	// currentUserDto := dto.CurrentUserDto{
	// 	UserReference: iEventData.UserReference,
	// 	Phone:         iEventData.Contact,
	// 	Device: dto.DeviceDto{
	// 		DeviceReference: iEventData.Device.DeviceReference,
	// 		Imei:            iEventData.Device.Imei,
	// 		Brand:           iEventData.Device.Brand,
	// 		Model:           iEventData.Device.Model,
	// 		Type:            iEventData.Device.Type,
	// 	},
	// }

	eventType := event.(eto.Event).EventType
	iEventData := data.(map[string]interface{})
	device := iEventData["device"].(map[string]interface{})
	userReference := iEventData["user_reference"].(string)
	contact := iEventData["contact"].(string)

	switch eventType {
	case "create_user":

		currentUserDto := dto.CurrentUserDto{
			UserReference: userReference,
			Phone:         contact,
			Device: dto.DeviceDto{
				DeviceReference: device["device_reference"].(string),
				Imei:            device["imei"].(string),
				Brand:           device["brand"].(string),
				Model:           device["model"].(string),
				Type:            device["type"].(string),
			},
		}

		user, _ := services.UserService.GetUserByReference(ctx, userReference)

		if user == nil {
			createUserDto := dto.CreateUserDto{
				Phone: contact,
			}
			// Create an instance of the UserService
			services.UserService.CreateUser(ctx, createUserDto, currentUserDto)
		}

	case "verify_email":
		user, _ := services.UserService.GetUserByReference(ctx, userReference)

		if user != nil {
			services.UserService.UpdateUserProfileEmailStatus(ctx, userReference)
		}

	// case "verify_phone":
	// 	user, _ := services.UserService.GetUserByReference(ctx, userReference)

	// 	if user != nil {
	// 		services.UserService.UpdateUserProfilePhoneStatus(ctx, userReference)
	// 	}
	default:
		logger.LogEvent("ERROR:", fmt.Sprintf("invalid otp even type: %v", eventType))
	}

	// userReference := iEventData["user_reference"].(string)
	// contact := iEventData["contact"].(string)

	// currentUserDto := dto.CurrentUserDto{
	// 	UserReference: userReference,
	// 	Phone:       contact,
	// 	Device: dto.DeviceDto{
	// 		DeviceReference: device["device_reference"].(string),
	// 		Imei:            device["imei"].(string),
	// 		Brand:           device["brand"].(string),
	// 		Model:           device["model"].(string),
	// 		Type:            device["type"].(string),
	// 	},
	// }

	// user, _ := services.UserService.GetUserByReference(ctx, userReference)

	// if eventType == channel.AcceptedChannels[eventType] && user == nil {
	// 	createUserDto := dto.CreateUserDto{
	// 		Phone: contact,
	// 	}
	// 	// Create an instance of the UserService
	// 	services.UserService.CreateUser(ctx, createUserDto, currentUserDto)
	// } else {
	// 	return

	// }

}
