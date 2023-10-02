package services

import (
	"context"
	"errors"
	"fmt"
	"time"

	publisher "walls-user-service/internal/adapter/events/publisher"
	"walls-user-service/internal/core/domain/dto"
	"walls-user-service/internal/core/domain/entity"
	event "walls-user-service/internal/core/domain/event/eto"
	"walls-user-service/internal/core/domain/mapper"
	configuration "walls-user-service/internal/core/helper/configuration-helper"
	eto "walls-user-service/internal/core/helper/event-helper/eto"
	logger "walls-user-service/internal/core/helper/log-helper"
	validation "walls-user-service/internal/core/helper/validation-helper"
	ports "walls-user-service/internal/port"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var UserService = &userService{}

type userService struct {
	userRepository ports.UserRepository
	redisClient    *redis.Client
}

func NewUserService(userRepository ports.UserRepository, redisClient *redis.Client) *userService {
	UserService = &userService{
		userRepository: userRepository,
		redisClient:    redisClient,
	}

	return UserService
}

func (service *userService) CreateUser(ctx context.Context, createUserDto dto.CreateUserDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	logger.LogEvent("INFO", "Creating User")
	userData, _ := service.GetUserByReference(ctx, currentUserDto.UserReference)
	if userData != nil {
		logger.LogEvent("ERROR", "Sorry, user already exists")
		return nil, errors.New("sorry, user already exists")
	}

	if createUserDto.Phone != currentUserDto.Phone {
		logger.LogEvent("ERROR", "The current user phone number does not match the intended registration phone number")
		return nil, errors.New("the current user phone number does nto match the intended registration phone number")
	}

	user := mapper.CurrentUserDtoToUser(createUserDto, currentUserDto)

	result, err := service.userRepository.CreateUser(ctx, user)
	if err != nil {
		logger.LogEvent("ERROR", "Unable to create User")
		return nil, errors.New("unable to create User")
	}

	request := struct {
		UserReference string        `json:"user_reference"`
		Phone         string        `json:"phone"`
		Device        entity.Device `json:"device"`
	}{
		UserReference: user.UserReference,
		Phone:         user.UserProfile.Phone,
		Device:        user.Device,
	}

	userCreatedEvent := event.UserCreatedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "usercreatedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "usercreatedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          request,
		},
	}
	//publishing user created event

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishUserCreatedEvent(ctx, userCreatedEvent)

	// requestDto := dto.IdentityDto{
	// 	Phone:  user.UserProfile.Phone,
	// 	Device: dto.DeviceDto(user.Device),
	// }

	//publishing create identity request
	//service.CreateIdentityRequest(ctx, user.UserReference, requestDto, currentUserDto)

	return result, err
}

func (service *userService) CreateCompanyProfile(ctx context.Context, user_reference string, createCompanyProfileDto dto.CreateCompanyProfileDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	logger.LogEvent("INFO", "Creating Company Profile")
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}
	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}

	user = mapper.CreateCompanyProfileDtoToUser(user, createCompanyProfileDto)

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to create walls tag")
		return nil, errors.New("failed to create walls tag")
	}

	companyProfileCreatedEvent := event.CompanyProfileCreatedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "companyprofilecreatedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "companyprofilecreatedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishCompanyProfileCreatedEvent(ctx, companyProfileCreatedEvent)

	return result, nil
}

func (service *userService) CreateCompanyWallsBadge(ctx context.Context, user_reference string, companyWallsBadgeDto dto.CompanyWallsBadgeDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}

	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	wallsTagData, _ := service.GetUserByWallsTag(ctx, companyWallsBadgeDto.WallsTag)
	if wallsTagData != nil {
		logger.LogEvent("ERROR", "WallsTag already exists")
		return nil, errors.New("this wallstag is already in use")
	}

	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}

	if user.UserReference != currentUserDto.UserReference {
		logger.LogEvent("ERROR", "This user with phone number "+currentUserDto.Phone+" is not registered")
		return nil, errors.New("this user  with phone number " + currentUserDto.Phone + " is not registered")
	}

	user = mapper.CreateCompanyWallsBadgeDtoToUser(user, companyWallsBadgeDto)
	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to create walls tag")
		return nil, errors.New("failed to create walls tag")
	}

	companyWallsBadgeCreatedEvent := event.CompanyWallsBadgeCreatedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "companywallsbadgecreatedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "companywallsbadgecreatedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishCompanyWallsBadgeCreatedEvent(ctx, companyWallsBadgeCreatedEvent)

	return result, nil
}

func (service *userService) CreateUserWallsBadge(ctx context.Context, user_reference string, userWallsBadgeDto dto.UserWallsBadgeDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}

	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	wallsTagData, _ := service.GetUserByWallsTag(ctx, userWallsBadgeDto.WallsTag)
	if wallsTagData != nil {
		logger.LogEvent("ERROR", "WallsTag already exists")
		return nil, errors.New("this wallstag is already in use")
	}

	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}

	if user.UserReference != currentUserDto.UserReference {
		logger.LogEvent("ERROR", "This user with phone number "+currentUserDto.Phone+" is not registered")
		return nil, errors.New("this user  with phone number " + currentUserDto.Phone + " is not registered")
	}

	user = mapper.CreateUserWallsBadgeDtoToUser(user, userWallsBadgeDto)
	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to create walls tag")
		return nil, errors.New("failed to create walls tag")
	}

	request := struct {
		UserReference   string `json:"user_reference"`
		DeviceReference string `json:"device_reference"`
		Contact         string `json:"contact"`
		Channel         string `json:"channel"`
		Message         string `json:"message"`
	}{
		UserReference:   user.UserReference,
		DeviceReference: user.Device.DeviceReference,
		Contact:         user.UserProfile.Email,
		Channel:         "email",
		Message:         "A new wallsbadge has been successfully created for you",
	}

	userWallsBadgeCreatedEvent := event.UserWallsBadgeCreatedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "userwallsbadgecreatedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "userwallsbadgecreatedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          request,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishUserWallsBadgeCreatedEvent(ctx, userWallsBadgeCreatedEvent)

	return result, nil
}

func (service *userService) UpdateCompanyProfile(ctx context.Context, user_reference string, companyProfileReference string, updateCompanyProfileDto dto.UpdateCompanyProfileDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	logger.LogEvent("INFO", "Updating Company Profile")
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}
	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}

	for i, c := range user.CompanyProfile {
		if c.CompanyProfileReference == companyProfileReference {
			user.CompanyProfile[i].Email = updateCompanyProfileDto.Email
			user.CompanyProfile[i].Phone = updateCompanyProfileDto.Phone
			user.CompanyProfile[i].Address = updateCompanyProfileDto.Address
			// this overrides the company profile which is not supposed to
			// user.CompanyProfile[i] = mapper.UpdateCompanyProfileDtoToCompanyProfile(updateCompanyProfileDto)
			break
		}
	}

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to create walls tag")
		return nil, errors.New("failed to create walls tag")
	}

	companyProfileUpdatedEvent := event.CompanyProfileUpdatedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "companyprofilecreatedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "companyprofilecreatedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishCompanyProfileUpdatedEvent(ctx, companyProfileUpdatedEvent)

	return result, nil
}

func (service *userService) DisableCompanyWallsBadge(ctx context.Context, user_reference string, companyProfileReference string, companyWallsBadgeReference string, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}
	user := mapper.UserDtoToUser(userData.(dto.UserDto))
	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}

OUTER:
	for index, companyProfile := range user.CompanyProfile {
		if companyProfile.CompanyProfileReference == companyProfileReference {
			for index2, wallsBadge := range companyProfile.WallsBadge {
				if wallsBadge.WallsBadgeReference == companyWallsBadgeReference {
					user.CompanyProfile[index].WallsBadge[index2].IsActive = false
					break OUTER
				}
			}
		}
	}

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to update user's name")
		return nil, errors.New("failed to update user's name")
	}

	companyWallsBadgeDisabledEvent := event.CompanyWallsBadgeDisabledEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "companywallsbadgedisabledevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "companywallsbadgedisabledevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishCompanyWallsBadgeDisabledEvent(ctx, companyWallsBadgeDisabledEvent)

	return result, nil
}

func (service *userService) DisableUserWallsBadge(ctx context.Context, user_reference string, userWallsBadgeReference string, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}
	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}

	for index, userWallsBadge := range user.UserProfile.WallsBadge {
		if userWallsBadge.WallsBadgeReference == userWallsBadgeReference {
			user.UserProfile.WallsBadge[index].IsActive = false
			break
		}
	}

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to update user's name")
		return nil, errors.New("failed to update user's name")
	}

	userWallsBadgeDisabledEvent := event.UserWallsBadgeDisabledEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "userwallsbadgedisabledEvent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "userwallsbadgedisabledEvent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishUserWallsBadgeDisabledEvent(ctx, userWallsBadgeDisabledEvent)

	return result, nil
}

func (service *userService) GetCompanyWallsBadgeList(ctx context.Context, user_reference string, companyProfileReference string) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}
	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	result := []entity.WallsBadge{}
	for _, companyProfile := range user.CompanyProfile {
		if companyProfile.CompanyProfileReference == companyProfileReference {
			result = companyProfile.WallsBadge
		}
	}

	return result, nil
}

func (service *userService) GetUserWallsBadgeList(ctx context.Context, user_reference string) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}
	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	result := user.UserProfile.WallsBadge

	return result, nil
}

func (service *userService) GetDefaultCompanyWallsBadge(ctx context.Context, user_reference string, companyProfileReference string) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}
	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	result := entity.WallsBadge{}
OUTER:
	for _, companyProfile := range user.CompanyProfile {
		if companyProfile.CompanyProfileReference == companyProfileReference {
			for _, wallsBadge := range companyProfile.WallsBadge {
				if wallsBadge.IsDefault {
					result = wallsBadge
					break OUTER
				}
			}
		}
	}

	return result, nil
}

func (service *userService) GetDefaultUserWallsBadge(ctx context.Context, user_reference string) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}
	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	// result := entity.WallsBadge{}
	for _, wallsBagde := range user.UserProfile.WallsBadge {
		if wallsBagde.IsDefault {
			return wallsBagde, nil
		}
	}

	return nil, errors.New("no default wallsbadge found")
}

func (service *userService) DisableCompanyProfile(ctx context.Context, user_reference string, companyProfileReference string) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}
	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	for index, companyProfile := range user.CompanyProfile {
		if companyProfile.CompanyProfileReference == companyProfileReference {
			user.CompanyProfile[index].IsActive = false
			break
		}
	}

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to update user's name")
		return nil, errors.New("failed to update user's name")
	}

	companyProfileDisabledEvent := event.CompanyProfileDisabledEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "companyprofiledisabledevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "companyprofiledisabledevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishCompanyProfileDisabledEvent(ctx, companyProfileDisabledEvent)

	return result, nil
}

func (service *userService) UpdateCompanyLogo(ctx context.Context, user_reference string, companyProfileReference string, updateCompanyLogoDto dto.UpdateCompanyLogo, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}
	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}

	for index, companyProfile := range user.CompanyProfile {
		if companyProfile.CompanyProfileReference == companyProfileReference {
			user.CompanyProfile[index].Logo = updateCompanyLogoDto.Logo
		}
	}

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to update user's name")
		return nil, errors.New("failed to update user's name")
	}

	companyLogoUpdatedEvent := event.CompanyLogoUpdatedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "companylogoupdatedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "companylogoupdatedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishCompanyLogoUpdatedEvent(ctx, companyLogoUpdatedEvent)

	return result, nil
}

func (service *userService) UpdateUserProfileEmailStatus(ctx context.Context, user_reference string) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}
	// user := mapper.UserDtoToUser(userData.(dto.UserDto))
	user := userData.(entity.User)

	user.UserProfile.IsVerifiedEmail = true

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to update user's name")
		return nil, errors.New("failed to update user's name")
	}

	etoRequestData := struct {
		UserReference string        `json:"user_reference"`
		Contact       string        `json:"contact"`
		Channel       string        `json:"channel"`
		Device        entity.Device `json:"device"`
	}{
		UserReference: user.UserReference,
		Contact:       user.UserProfile.Email,
		Channel:       "email",
		Device:        user.Device,
	}

	userProfileEmailStatusUpdatedEvent := event.UserProfileEmailStatusUpdatedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "userprofileemailstatusupdatedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "userprofileemailstatusupdatedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          etoRequestData,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishUserProfileEmailStatusUpdatedEvent(ctx, userProfileEmailStatusUpdatedEvent)

	return result, nil
}

// func (service *userService) UpdateUserProfilePhoneStatus(ctx context.Context, user_reference string) (interface{}, error) {
// 	userData, err := service.GetUserByReference(ctx, user_reference)
// 	if err != nil {
// 		logger.LogEvent("ERROR", "Failed to retrieve user")
// 		return nil, errors.New("failed to retrieve user")
// 	}
// 	// user := mapper.UserDtoToUser(userData.(dto.UserDto))
// 	user := userData.(entity.User)

// 	user.UserProfile.IsVerifiedPhone = true

// 	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
// 	if err != nil {
// 		logger.LogEvent("ERROR", "Failed to update user's name")
// 		return nil, errors.New("failed to update user's name")
// 	}

// 	userProfilePhoneStatusUpdatedEvent := event.UserProfilePhoneStatusUpdatedEvent{
// 		Event: eto.Event{
// 			EventReference:     uuid.New().String(),
// 			EventName:          "updateuserprofilephonestatusevent",
// 			EventDate:          time.Now().Format(time.RFC3339),
// 			EventType:          "updateuserprofilephonestatusevent",
// 			EventSource:        configuration.ServiceConfiguration.ServiceName,
// 			EventUserReference: user.UserReference,
// 			EventData:          user,
// 		},
// 	}

// 	eventPublisher := publisher.NewPublisher(service.redisClient)
// 	eventPublisher.PublishUserProfilePhoneStatusUpdatedEvent(ctx, userProfilePhoneStatusUpdatedEvent)

// 	return result, nil
// }

func (service *userService) UpdateCompanyProfileEmailStatus(ctx context.Context, user_reference string, companyProfileReference string) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}
	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	for index, companyProfile := range user.CompanyProfile {
		if companyProfile.CompanyProfileReference == companyProfileReference {
			user.CompanyProfile[index].IsVerifiedEmail = true
		}
	}

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to update user's name")
		return nil, errors.New("failed to update user's name")
	}

	companyProfileEmailStatusUpdatedEvent := event.CompanyProfileEmailStatusUpdatedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "companyprofileemailstatusupdatedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "companyprofileemailstatusupdatedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishCompanyProfileEmailStatusUpdatedEvent(ctx, companyProfileEmailStatusUpdatedEvent)

	return result, nil
}



func (service *userService) SetDefaultBank(ctx context.Context, user_reference string, bankReference string, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}
	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}

	for index, bankAccount := range user.BankAccounts {
		if bankAccount.BankReference == bankReference {
			user.BankAccounts[index].IsDefault = true
			break
		}
	}

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to update user's name")
		return nil, errors.New("failed to update user's name")
	}

	defaultBankSetEvent := event.DefaultBankSetEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "defaultbanksetevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "defaultbanksetevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishDefaultBankSetEvent(ctx, defaultBankSetEvent)

	return result, nil
}

func (service *userService) SetDefaultCard(ctx context.Context, user_reference string, cardReference string, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}
	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}

	for index, card := range user.Cards {
		if card.CardReference == cardReference {
			user.Cards[index].IsDefault = true
		}
	}

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to update user's name")
		return nil, errors.New("failed to update user's name")
	}

	defaultCardSetEvent := event.DefaultCardSetEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "defaultcardsetevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "defaultcardsetevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishDefaultCardSetEvent(ctx, defaultCardSetEvent)

	return result, nil
}

func (service *userService) UpdateUserName(ctx context.Context, user_reference string, usernameDto dto.UserNameDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}
	user := userData.(entity.User)

	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}

	user = mapper.UpdateUserNameDtoToUser(user, usernameDto)

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to update user's name")
		return nil, errors.New("failed to update user's name")
	}

	etoRequestData := struct {
		UserReference string        `json:"userReference"`
		Contact       string        `json:"contact"`
		Channel       string        `json:"channel"`
		Device        entity.Device `json:"device"`
	}{
		UserReference: user.UserReference,
		Contact:       user.UserProfile.Phone,
		Channel:       "sms",
		Device:        user.Device,
	}

	usernameUpdatedEvent := event.UsernameUpdatedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "usernameupdatedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "usernameupdatedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          etoRequestData,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishUsernameUpdatedEvent(ctx, usernameUpdatedEvent)

	return result, nil
}

func (service *userService) UpdateEmail(ctx context.Context, user_reference string, emailDto dto.EmailDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}
	// user := mapper.UserDtoToUser(userData.(dto.UserDto))
	user := userData.(entity.User)

	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}

	user = mapper.UpdateEmailDtoToUser(user, emailDto)

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to update user's email")
		return nil, errors.New("failed to update user's email")
	}

	// emailUpdatedEvent := event.EmailUpdatedEvent{
	// 	Event: eto.Event{
	// 		EventReference:     uuid.New().String(),
	// 		EventName:          "emailupdatedevent",
	// 		EventDate:          time.Now().Format(time.RFC3339),
	// 		EventType:          "emailupdatedevent",
	// 		EventSource:        configuration.ServiceConfiguration.ServiceName,
	// 		EventUserReference: user.UserReference,
	// 		EventData:          user,
	// 	},
	// }

	// eventPublisher := publisher.NewPublisher(service.redisClient)
	// eventPublisher.PublishEmailUpdatedEvent(ctx, emailUpdatedEvent)

	return result, nil
}

func (service *userService) UpdateDateOfBirth(ctx context.Context, user_reference string, dobDto dto.DobDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}

	user := mapper.UserDtoToUser(userData.(dto.UserDto))
	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}
	user = mapper.UpdateDobDtoToUser(user, dobDto)

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to update user's date of birth")
		return nil, errors.New("failed to update user's date of birth")
	}

	dobUpdatedEvent := event.DOBUpdatedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "dobupdatedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "dobupdatedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishDOBUpdatedEvent(ctx, dobUpdatedEvent)

	return result, nil
}

func (service *userService) UpdateAddress(ctx context.Context, user_reference string, addressDto dto.AddressDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}

	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}

	user = mapper.UpdateAddressDtoToUser(user, addressDto)

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to update user's address")
		return nil, errors.New("failed to update user's address")
	}

	addressUpdatedEvent := event.AddressUpdatedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "addressupdatedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "addressupdatedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishAddressUpdatedEvent(ctx, addressUpdatedEvent)

	return result, nil
}

func (service *userService) UpdatePhoto(ctx context.Context, user_reference string, photoDto dto.PhotoDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}

	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}

	user = mapper.UpdatePhotoDtoToUser(user, photoDto)

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to update user's photos")
		return nil, errors.New("failed to update user's photos")
	}

	photosUpdatedEvent := event.PhotosUpdatedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "photoupdatedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "photoupdatedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishPhotosUpdatedEvent(ctx, photosUpdatedEvent)

	return result, nil
}

func (service *userService) UpdateWallet(ctx context.Context, user_reference string, walletDto dto.UpdateWalletDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}

	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}

	user = mapper.UpdateWalletDtoToWallet(user, walletDto)

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to update user's wallet")
		return nil, errors.New("failed to update user's wallet")
	}

	walletUpdatedEvent := event.WalletUpdatedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "walletupdatedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "walletupdatedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishWalletUpdatedEvent(ctx, walletUpdatedEvent)

	return result, nil
}

func (service *userService) AddBank(ctx context.Context, user_reference string, bankDto dto.BankDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}

	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}

	user = mapper.AddBankDtoToBank(user, bankDto)

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to add bank for the user")
		return nil, errors.New("failed to add bank for the user")
	}

	bankAddedEvent := event.BankAddedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "bankaddedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "bankaddedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishBankAddedEvent(ctx, bankAddedEvent)

	return result, nil
}

func (service *userService) UpdateBank(ctx context.Context, user_reference string, bank_reference string, bankDto dto.BankDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}

	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}

	for i, b := range user.BankAccounts {
		if b.BankReference == bank_reference {
			user.BankAccounts[i] = mapper.UpdateBankDtoToBank(bankDto, bank_reference)
			break
		}
	}

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to update bank for the user")
		return nil, errors.New("failed to update bank for the user")
	}

	bankUpdatedEvent := event.BankUpdatedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "bankupdatedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "bankupdatedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishBankUpdatedEvent(ctx, bankUpdatedEvent)

	return result, nil
}

func (service *userService) AddCard(ctx context.Context, user_reference string, cardDto dto.CardDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}

	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}
	//check if expry month and year is valid
	if !isValidExpiryMonthAndYear(cardDto.ExpiryMonth, cardDto.ExpiryYear) {
		logger.LogEvent("ERROR", "Invalid expiry month or year")
		return nil, errors.New("invalid expiry month or year")
	}

	user = mapper.AddCardDtoToCard(user, cardDto)

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to add card for the user")
		return nil, errors.New("failed to add card for the user")
	}

	cardAddedEvent := event.CardAddedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "cardaddedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "cardaddedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishCardAddedEvent(ctx, cardAddedEvent)

	return result, nil
}

func (service *userService) UpdateCard(ctx context.Context, user_reference string, card_reference string, cardDto dto.CardDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}

	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}
	//check if expry month and year is valid
	if !isValidExpiryMonthAndYear(cardDto.ExpiryMonth, cardDto.ExpiryYear) {
		logger.LogEvent("ERROR", "Invalid expiry month or year")
		return nil, errors.New("invalid expiry month or year")
	}

	for i, c := range user.Cards {
		if c.CardReference == card_reference {
			user.Cards[i] = mapper.UpdateCardDtoToCard(cardDto, card_reference)
			break
		}
	}

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to update card for the user")
		return nil, errors.New("failed to update card for the user")
	}

	cardUpdatedEvent := event.CardUpdatedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "cardupdatedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "cardupdatedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishCardUpdatedEvent(ctx, cardUpdatedEvent)

	return result, nil
}

func (service *userService) UpdateNotificationOptions(ctx context.Context, user_reference string, optionsDto dto.UpdateNotificationOptionsDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}

	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}

	user = mapper.UpdateNotificationOptionsDtoToNotificationOptions(user, optionsDto)

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to update user's notification options")
		return nil, errors.New("failed to update user's notification options")
	}

	notificationOptionsUpdatedEvent := event.NotificationOptionsUpdatedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "notificationoptionsupdatedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "notificationoptionsupdatedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishNotificationOptionsUpdatedEvent(ctx, notificationOptionsUpdatedEvent)

	return result, nil
}

func (service *userService) UpdateDevice(ctx context.Context, user_reference string, deviceDto dto.UpdateDeviceDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}

	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}

	//check if the new device is registered at all here, if it is registered stop
	registeredUserwithDevice, _ := service.GetUserByDevice(ctx, dto.DeviceDto(deviceDto.NewDevice))
	if registeredUserwithDevice != nil {
		logger.LogEvent("ERROR", "This device is already registered")
		return nil, errors.New("this device is already registered")
	}

	user = mapper.UpdateDeviceDtoToDevice(user, deviceDto)
	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to update device for the user")
		return nil, errors.New("failed to update device for the user")
	}

	deviceUpdatedEvent := event.DeviceUpdatedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "deviceupdatedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "deviceupdatedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishDeviceUpdatedEvent(ctx, deviceUpdatedEvent)

	return result, nil
}

func (service *userService) AddDocumentation(ctx context.Context, user_reference string, documentationDto dto.AddDocumentationDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}

	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}

	user = mapper.AddDocumentationDtoToDocumentation(user, documentationDto)

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to add documentation for the user")
		return nil, errors.New("failed to add documentation for the user")
	}

	documentationAddedEvent := event.DocumentationAddedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "documentationaddedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "documentationaddedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishDocumentationAddedEvent(ctx, documentationAddedEvent)

	return result, nil
}

func (service *userService) UpdateDocumentation(ctx context.Context, user_reference string, documentation_reference string, documentationDto dto.AddDocumentationDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}

	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}

	for i, id := range user.Kyc.Documentations {
		if id.DocumentationReference == documentation_reference {
			user.Kyc.Documentations[i] = mapper.UpdateDocumentationDtoToDocumentation(documentationDto, documentation_reference)
			break
		}
	}

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to update documentation for the user")
		return nil, errors.New("failed to update documentation for the user")
	}

	documentationUpdatedEvent := event.DocumentationUpdatedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "documentationupdatedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "documentationupdatedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishDocumentationUpdatedEvent(ctx, documentationUpdatedEvent)

	return result, nil
}

func (service *userService) AddContact(ctx context.Context, user_reference string, contactDto dto.ContactDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}

	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}

	user = mapper.AddContactDtoToContact(user, contactDto)

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to add contact for the user")
		return nil, errors.New("failed to add contact for the user")
	}

	contactAddedEvent := event.ContactAddedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "contactaddedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "contactaddedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishContactAddedEvent(ctx, contactAddedEvent)

	return result, nil
}

func (service *userService) GetUserByReference(ctx context.Context, user_reference string) (interface{}, error) {
	logger.LogEvent("INFO", "Fetching User by reference: "+user_reference)

	user, err := service.userRepository.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to fetch User by reference: "+user_reference)
		return nil, errors.New("failed to retrieve user")
	}

	if user == nil {
		logger.LogEvent("ERROR", "User not found for reference: "+user_reference)
		return nil, errors.New("user not found")
	}

	logger.LogEvent("INFO", "User fetched successfully by reference: "+user_reference)

	return user, nil
}
func (service *userService) GetUserDefaultBadge(ctx context.Context, user_reference string) (interface{}, error) {
	logger.LogEvent("INFO", "Fetching User by reference: "+user_reference)

	badge, err := service.userRepository.GetUserDefaultWallsBadge(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to fetch default user wallsbadge by reference: "+user_reference)
		return nil, errors.New("failed to retrieve user")
	}

	if badge == nil {
		logger.LogEvent("ERROR", "badge not found for reference: "+user_reference)
		return nil, errors.New("badge not found")
	}

	logger.LogEvent("INFO", "Badge fetched successfully by reference: "+user_reference)

	return badge, nil
}
func (service *userService) GetUserByPhone(ctx context.Context, phone string) (interface{}, error) {
	user, err := service.userRepository.GetUserByPhone(ctx, phone)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to fetch User by phone: "+phone)
		return nil, errors.New("failed to retrieve user")
	}

	if user == nil {
		logger.LogEvent("ERROR", "User not found for phone: "+phone)
		return nil, errors.New("user not found")
	}

	logger.LogEvent("INFO", "User fetched successfully by phone: "+phone)

	return user, nil
}

func (service *userService) GetUserByWallsTag(ctx context.Context, wallsTag string) (interface{}, error) {
	user, err := service.userRepository.GetUserByWallsTag(ctx, wallsTag)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to fetch User by reference: "+wallsTag)
		return nil, errors.New("failed to retrieve user")
	}

	if user == nil {
		logger.LogEvent("ERROR", "User not found for reference: "+wallsTag)
		return nil, errors.New("user not found")
	}

	logger.LogEvent("INFO", "User fetched successfully by reference: "+wallsTag)

	return user, nil
}

func (service *userService) GetUserByWallsBagdeReference(ctx context.Context, wallsBagdeReference string) (interface{}, error) {
	user, err := service.userRepository.GetUserByWallsBadgeReference(ctx, wallsBagdeReference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to fetch User by reference: "+wallsBagdeReference)
		return nil, errors.New("failed to retrieve user")
	}

	if user == nil {
		logger.LogEvent("ERROR", "User not found for reference: "+wallsBagdeReference)
		return nil, errors.New("user not found")
	}

	logger.LogEvent("INFO", "User fetched successfully by reference: "+wallsBagdeReference)

	return user, nil
}

func (service *userService) GetUserByDevice(ctx context.Context, device dto.DeviceDto) (interface{}, error) {
	user, err := service.userRepository.GetUserByDevice(ctx, entity.Device(device))
	if err != nil {
		logger.LogEvent("ERROR", "Failed to fetch User by device")
		return nil, errors.New("failed to retrieve user")
	}

	return user, nil
}

func (service *userService) UpdateBalance(ctx context.Context, user_reference string, balanceDto dto.BalanceDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}

	user := mapper.UserDtoToUser(userData.(dto.UserDto))
	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}

	user.Wallet.Balance.PendingIncomingAmount = balanceDto.BookAmount
	user.Wallet.Balance.AvailableAmount = balanceDto.AvailableAmount
	user.Wallet.Balance.IsSynced = true
	user.Wallet.Balance.LastSyncedOn = time.Now().Format(time.RFC3339)

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to update book balance for the user")
		return nil, errors.New("failed to update book balance for the user")
	}

	balanceUpdatedEvent := event.BalanceUpdatedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "balanceupdatedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "balanceupdatedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishBalanceUpdatedEvent(ctx, balanceUpdatedEvent)

	return result, nil
}

func (service *userService) UpdateTier(ctx context.Context, user_reference string, tierDto dto.TierDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}
	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	//check if device is registered
	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "This device is not registered to this user")
		return nil, errors.New("this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "The phone number is not registered to this user")
		return nil, errors.New("the phone number is not registered to this user")
	}

	user.Wallet.Tier.TierReference = tierDto.TierReference
	user.Wallet.Tier.TierName = tierDto.TierName
	user.Wallet.Tier.SendingLimit = tierDto.SendingLimit
	user.Wallet.Tier.DailyTransactionLimit = tierDto.DailyTransactionLimit
	user.Wallet.Tier.MinimumBalance = tierDto.MinimumBalance
	user.Wallet.Tier.DailyTransactionLimit = tierDto.DailyTransactionLimit
	user.Wallet.Tier.UpgradeOptions = tierDto.UpgradeOptions

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to update tier for the user")
		return nil, errors.New("failed to update tier for the user")
	}

	tierUpdatedEvent := event.TierUpdatedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "tierupdatedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "tierupdatedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishTierUpdatedEvent(ctx, tierUpdatedEvent)

	return result, nil
}

func (service *userService) AddCoupon(ctx context.Context, user_reference string, couponDto dto.CouponDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}
	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	user.Wallet.Coupons = append(user.Wallet.Coupons, entity.Coupon(couponDto))

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to update coupons for the user")
		return nil, errors.New("failed to update coupons for the user")
	}

	couponAddedEvent := event.CouponAddedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "couponaddedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "couponaddedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishCouponAddedEvent(ctx, couponAddedEvent)

	return result, nil
}

func (service *userService) UpdateRewards(ctx context.Context, user_reference string, rewardDto dto.RewardDto) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}
	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	// user.Wallet.Reward.Points = rewardDto.Points
	user.Wallet.Reward.Points += rewardDto.Points

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to update rewards for the user")
		return nil, errors.New("failed to update rewards for the user")
	}

	rewardsUpdatedEvent := event.RewardsUpdatedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "rewardsupdatedevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "rewardsupdatedevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishRewardsUpdatedEvent(ctx, rewardsUpdatedEvent)

	return result, nil
}


func (service *userService) EnableUser(ctx context.Context, user_reference string) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}

	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	user.IsActive = true

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to enable user")
		return nil, errors.New("failed to enable user")
	}

	userEnabledEvent := event.UserEnabledEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "userenabledevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "userenabledevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishUserEnabledEvent(ctx, userEnabledEvent)

	return result, nil
}

func (service *userService) DisableUser(ctx context.Context, user_reference string) (interface{}, error) {
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}
	user := mapper.UserDtoToUser(userData.(dto.UserDto))

	user.IsActive = false

	result, err := service.userRepository.UpdateUser(ctx, user_reference, user)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to disable user")
		return nil, errors.New("failed to disable user")
	}

	userDisabledEvent := event.UserDisabledEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "userdisabledevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "userdisabledevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: user.UserReference,
			EventData:          user,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishUserDisabledEvent(ctx, userDisabledEvent)

	return result, nil
}

func (service *userService) CreateOtpRequest(ctx context.Context, requestOtpDto dto.CreateOtpDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	logger.LogEvent("INFO", "Requesting OTP Creation")

	// Fetch the identity by user reference
	if requestOtpDto.OtpType != "create_user" {
		userData, err := service.GetUserByReference(ctx, currentUserDto.UserReference)
		if err != nil {
			logger.LogEvent("ERROR", "Unauthorized User: Failed to fetch identity for the user with reference: "+currentUserDto.UserReference)
			return nil, errors.New("unauthorized user: failed to retrieve identity")
		}

		// Convert identityData to entity.Identity type
		user := userData.(entity.User)

		//check if device is registered

		if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
			logger.LogEvent("ERROR", "Unauthorized Device: This device is not registered to this user")
			return nil, errors.New("unauthorized device: this device is not registered to this user")
		}
		//check if the phone number is registered
		if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
			logger.LogEvent("ERROR", "Unauthorized Phone: The phone number is not registered to this user")
			return nil, errors.New("unauthorized phone: the phone number is not registered to this user")
		}
	}

	// new user must request otp with phone number as contact
	if requestOtpDto.OtpType == "create_user" && !validation.IsValidPhone(requestOtpDto.Contact) {
		fmt.Println("contact for new user otp must be a phone number")
		logger.LogEvent("ERROR", "contact for new user otp must be a phone number")
		return nil, errors.New("contact for new user otp must be a phone number")
	}

	// email verification must have an email as contact
	if requestOtpDto.OtpType == "verify_email" && !validation.IsValidEmail(requestOtpDto.Contact) {
		fmt.Println("otp for email verification requires an email contact")
		logger.LogEvent("ERROR", "otp for email verification requires an email as contact")
		return nil, errors.New("otp for email verification requires an email as contact")
	}

	// phone number verification must have a phone number as contact
	if requestOtpDto.OtpType == "verify_phone" && !validation.IsValidPhone(requestOtpDto.Contact) {
		fmt.Println("otp for phone number verification requires an email contact")
		logger.LogEvent("ERROR", "otp for phone number verification requires a phone number as contact")
		return nil, errors.New("otp for email verification requires a phone number as contact")
	}

	request := struct {
		UserReference string        `json:"userReference"`
		Contact       string        `json:"contact"`
		Channel       string        `json:"channel"`
		Device        dto.DeviceDto `json:"device"`
	}{
		currentUserDto.UserReference,
		requestOtpDto.Contact,
		requestOtpDto.Channel,
		requestOtpDto.Device,
	}

	createOtpRequestEvent := event.OtpRequestCreatedEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "CREATEOTPREQUESTEVENT",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          requestOtpDto.OtpType,
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: currentUserDto.UserReference,
			EventData:          request,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishCreateOtpRequestEvent(ctx, createOtpRequestEvent, createOtpRequestEvent.EventType)
	return request.UserReference, nil
}

func (service *userService) ValidateOtpRequest(ctx context.Context, user_reference string, validateOtpDto dto.ValidateOtpDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {

	logger.LogEvent("INFO", "Validating OTP for "+user_reference)

	if validateOtpDto.OtpType != "create_user" {
		userData, err := service.GetUserByReference(ctx, currentUserDto.UserReference)
		if err != nil {
			logger.LogEvent("ERROR", "Unauthorized User: Failed to fetch identity for the user with reference: "+user_reference)
			return nil, errors.New("unauthorized user: failed to retrieve identity")
		}

		// Convert identityData to entity.Identity type
		user := userData.(entity.User)

		//check if device is registered

		if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
			logger.LogEvent("ERROR", "Unauthorized Device: This device is not registered to this user")
			return nil, errors.New("unauthorized device: this device is not registered to this user")
		}
		//check if the phone number is registered
		if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
			logger.LogEvent("ERROR", "Unauthorized Phone: The phone number is not registered to this user")
			return nil, errors.New("unauthorized phone: the phone number is not registered to this user")
		}
	}

	// new user must request otp validation with phone number as contact
	if validateOtpDto.OtpType == "create_user" && !validation.IsValidPhone(validateOtpDto.Contact) {
		fmt.Println("validating otp for new user otp requires a phone number as contact")
		logger.LogEvent("ERROR", "contact field for validating otp for new user otp requires a phone number as contact")
		return nil, errors.New("contact for new user otp must be a phone number")
	}

	// check if contact is an email
	if validateOtpDto.OtpType == "verify_email" && !validation.IsValidEmail(validateOtpDto.Contact) {
		fmt.Println("otp verification for email email must have an email in contact field")
		logger.LogEvent("ERROR", "otp verification for email email must have an email in contact field")
		return nil, errors.New("otp verification for email email must have an email in contact field")
	}

	// check if contact is a phone number
	if validateOtpDto.OtpType == "verify_phone" && !validation.IsValidPhone(validateOtpDto.Contact) {
		fmt.Println("otp verification for phone number must have a phone number in contact field")
		logger.LogEvent("ERROR", "otp verification for phone number must have a phone number in contact field")
		return nil, errors.New("otp verification for phone number must have a phone number in contact field")
	}

	etoRequestData := struct {
		UserReference string        `json:"userReference"`
		Contact       string        `json:"contact"`
		Otp           string        `json:"otp"`
		Device        dto.DeviceDto `json:"device"`
	}{
		UserReference: currentUserDto.UserReference,
		Contact:       validateOtpDto.Contact,
		Otp:           validateOtpDto.Otp,
		Device:        validateOtpDto.Device,
	}

	validateOtpRequestEvent := event.ValidateOtpRequestEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "OTPREQUESTCREATEDEVENT",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          validateOtpDto.OtpType,
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: currentUserDto.UserReference,
			EventData:          etoRequestData,
		},
	}
	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishValidateOtpRequestEvent(ctx, validateOtpRequestEvent, validateOtpDto.OtpType)

	result := "OTP sent for validation"

	return result, nil
}

func (service *userService) CreateIdentityRequest(ctx context.Context, user_reference string, requestIdentityDto dto.IdentityDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {

	logger.LogEvent("INFO", "Requesting Identity Creation")

	// Fetch the identity by user reference
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Unauthorized User: Failed to fetch identity for the user with reference: "+user_reference)
		return nil, errors.New("unauthorized user: failed to retrieve identity")
	}

	// Convert identityData to entity.Identity type
	user := userData.(entity.User)

	//check if device is registered

	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "Unauthorized Device: This device is not registered to this user")
		return nil, errors.New("unauthorized device: this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "Unauthorized Phone: The phone number is not registered to this user")
		return nil, errors.New("unauthorized phone: the phone number is not registered to this user")
	}

	request := struct {
		UserReference string        `json:"userReference"`
		Phone         string        `json:"phone"`
		Device        dto.DeviceDto `json:"device"`
	}{
		currentUserDto.UserReference,
		requestIdentityDto.Phone,
		requestIdentityDto.Device,
	}
	createIdentityRequestEvent := event.CreateIdentityRequestEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "createidentityrequestevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "createidentityrequestevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: currentUserDto.UserReference,
			EventData:          request,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishCreateIdentityRequestEvent(ctx, createIdentityRequestEvent)
	return request.UserReference, nil

}

func (service *userService) UpgradeTierRequest(ctx context.Context, user_reference string, requestTierDto dto.TierUpgradeRequestDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	logger.LogEvent("INFO", "Requesting Tier Upgrade")

	// Fetch the identity by user reference
	userData, err := service.GetUserByReference(ctx, user_reference)
	if err != nil {
		logger.LogEvent("ERROR", "Unauthorized User: Failed to fetch identity for the user with reference: "+user_reference)
		return nil, errors.New("unauthorized user: failed to retrieve identity")
	}

	// Convert identityData to entity.Identity type
	user := userData.(entity.User)

	//check if device is registered

	if !isRegisteredDevice(entity.Device(currentUserDto.Device), user.Device) {
		logger.LogEvent("ERROR", "Unauthorized Device: This device is not registered to this user")
		return nil, errors.New("unauthorized device: this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, user.UserProfile.Phone) {
		logger.LogEvent("ERROR", "Unauthorized Phone: The phone number is not registered to this user")
		return nil, errors.New("unauthorized phone: the phone number is not registered to this user")
	}

	if len(requestTierDto.TierDocuments) == 0 {
		logger.LogEvent("ERROR", "No uploaded documents: No uploaded documents found for this request. Kindly uploaded documents and try again.")
		return nil, errors.New("no uploaded documents: No uploaded documents found for this request. Kindly uploaded documents and try again")
	}

	request := struct {
		RequestReference string                  `json:"request_reference"`
		User             dto.UserDto             `json:"user"`
		CurrentTier      entity.Tier             `json:"current_tier"`
		RequestedTier    dto.TierDto             `json:"requested_tier"`
		KycDocuments     []entity.Documentation `json:"kyc_documents"`
		TierDocuments []dto.DocumentationDto 	`json:"tier_documents"`
		RequestStatus   string                  `json:"request_status"`
	}{
		uuid.New().String(),
		requestTierDto.User,
		user.Wallet.Tier,
		requestTierDto.RequestedTier,
		user.Kyc.Documentations,
		requestTierDto.TierDocuments,
		"pending",

	}
	upgradeTierRequest := event.TierUpgradeRequestEvent{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "upgradetierrequestevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "upgradetierrequestevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: currentUserDto.UserReference,
			EventData:          request,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishUpgradeTierRequestEvent(ctx, upgradeTierRequest)
	return request.RequestReference, nil
}

func (service *userService) CreateTransactionRequest(ctx context.Context, user_reference string, transactionDto dto.CreateTransactionDto, currentUserDto dto.CurrentUserDto) (interface{}, error) {
	logger.LogEvent("INFO", "Requesting Tier Upgrade")

	senderData, err := service.GetUserByReference(ctx, currentUserDto.UserReference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}
	sender := mapper.UserDtoToUser(senderData.(dto.UserDto))

	receiverData, err := service.GetUserByReference(ctx, currentUserDto.UserReference)
	if err != nil {
		logger.LogEvent("ERROR", "Failed to retrieve user")
		return nil, errors.New("failed to retrieve user")
	}
	receiver := mapper.UserDtoToUser(receiverData.(dto.UserDto))

	if !isRegisteredDevice(entity.Device(currentUserDto.Device), sender.Device) {
		logger.LogEvent("ERROR", "Unauthorized Device: This device is not registered to this user")
		return nil, errors.New("unauthorized device: this device is not registered to this user")
	}
	//check if the phone number is registered
	if !isRegisteredPhoneNumber(currentUserDto.Phone, sender.UserProfile.Phone) {
		logger.LogEvent("ERROR", "Unauthorized Phone: The phone number is not registered to this user")
		return nil, errors.New("unauthorized phone: the phone number is not registered to this user")
	}

	if sender.UserReference != currentUserDto.UserReference {
		logger.LogEvent("ERROR", "Unauthorized transaction")
		return nil, errors.New("unauthorized transaction")
	}

	if sender.Wallet.Balance.AvailableAmount < transactionDto.Amount {
		logger.LogEvent("ERROR", "Insufficient funds in sender's wallet")
		return nil, errors.New("insufficient funds in sender's wallet")
	}

	// Check if the transaction amount exceeds the sender's sending limit
	if transactionDto.Amount > sender.Wallet.Tier.SendingLimit {
		logger.LogEvent("ERROR", "The transaction amount exceeds the sender's sending limit")
		return nil, errors.New("the transaction amount exceeds the sender's sending limit")
	}

	// Check if receiver's wallet can accept more funds
	if receiver.Wallet.Balance.AvailableAmount+transactionDto.Amount > receiver.Wallet.Tier.WalletLimit {
		logger.LogEvent("ERROR", "Receiver's wallet limit exceeded")
		return nil, errors.New("receiver's wallet limit exceeded")
	}

	// Check if receiver's wallet can accept these funds at once
	if receiver.Wallet.Balance.AvailableAmount > receiver.Wallet.Tier.ReceivingLimit {
		logger.LogEvent("ERROR", "Receiver's receiving limit exceeded")
		return nil, errors.New("receiver's receiving limit exceeded")
	}

	requestSender := entity.Sender{
		Type:          transactionDto.TransactionType,
		UserReference: sender.UserReference,
		Reference:     transactionDto.OutReference,
		Device:        entity.Device(currentUserDto.Device),
	}
	var requestReceiverDefaultWallsBadge interface{}
	if transactionDto.ReceiverWallsBadgeReference == "" {
		requestReceiverDefaultWallsBadge, err = service.userRepository.GetUserDefaultWallsBadge(ctx, transactionDto.ReceiverReference)
		if err != nil {
			logger.LogEvent("ERROR", "Failed to retrieve user default badge. Aborting transaction.")
			return nil, errors.New("failed to retrieve user. aborting transaction")
		}
		transactionDto.ReceiverWallsBadgeReference = requestReceiverDefaultWallsBadge.(entity.WallsBadge).WallsBadgeReference
	}

	requestReceiver := entity.Receiver{
		Type:                transactionDto.TransactionType,
		UserReference:       receiver.UserReference,
		Reference:           transactionDto.InReference,
		WallsBadgeReference: transactionDto.ReceiverWallsBadgeReference,
	}

	request := struct {
		RequestReference string          `json:"request_reference"`
		TransactionType  string          `json:"transaction_type"`
		Amount           float64         `json:"amount"`
		Sender           entity.Sender   `json:"sender"`
		Receiver         entity.Receiver `json:"receiver"`
		Metadata         entity.Metadata `json:"metadata"`
	}{
		uuid.New().String(),
		transactionDto.TransactionType,
		transactionDto.Amount,
		requestSender,
		requestReceiver,
		transactionDto.Metadata,
	}

	transactionRequest := event.TransactionCreateRequest{
		Event: eto.Event{
			EventReference:     uuid.New().String(),
			EventName:          "createtransactionrequestevent",
			EventDate:          time.Now().Format(time.RFC3339),
			EventType:          "createtransactionrequestevent",
			EventSource:        configuration.ServiceConfiguration.ServiceName,
			EventUserReference: currentUserDto.UserReference,
			EventData:          request,
		},
	}

	eventPublisher := publisher.NewPublisher(service.redisClient)
	eventPublisher.PublishCreateTransactionRequestEvent(ctx, transactionRequest)
	return request.RequestReference, nil
}
func (service *userService) CreateUserReference(ctx context.Context) (interface{}, error) {
	logger.LogEvent("INFO", "Creating User Reference")

	newUserReference := uuid.New().String()

	return newUserReference, nil
}

func (service *userService) CreateDocumentReference(ctx context.Context) (interface{}, error) {
	logger.LogEvent("INFO", "Creating Document Reference")

	newUserReference := uuid.New().String()

	return newUserReference, nil
}

func (service *userService) CreateDocumentation(ctx context.Context, user_reference string, documentReferenceDto dto.DocumentReferenceDto ) (interface{}, error) {
	//return []dto.DocumentationDto
	return nil, nil
}

func isRegisteredDevice(currentDevice entity.Device, registeredDevice entity.Device) bool {

	return currentDevice.Imei == registeredDevice.Imei &&
		currentDevice.Type == registeredDevice.Type &&
		currentDevice.Brand == registeredDevice.Brand &&
		currentDevice.Model == registeredDevice.Model &&
		currentDevice.DeviceReference == registeredDevice.DeviceReference

}

func isRegisteredPhoneNumber(currentPhone string, registeredPhone string) bool {
	return currentPhone == registeredPhone
}

func isValidExpiryMonthAndYear(month, year int) bool {
	// previous year
	if year < time.Now().Year() {
		return false
	}

	// current year previous month
	if year == time.Now().Year() && month < int(time.Now().Month()) {
		return false
	}

	// future year
	if month < 0 || month > 12 {
		return false
	}

	return true
}
