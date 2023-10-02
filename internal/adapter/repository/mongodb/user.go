package repository

import (
	"context"
	"fmt"
	"time"
	"walls-user-service/internal/core/domain/entity"
	logger "walls-user-service/internal/core/helper/log-helper"
	ports "walls-user-service/internal/port"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserInfra struct {
	Collection *mongo.Collection
}

func NewUser(Collection *mongo.Collection) *UserInfra {
	return &UserInfra{Collection}
}

// UserRepo implements the repository.UserRepository interface
var _ ports.UserRepository = &UserInfra{}

func (r *UserInfra) CreateUser(ctx context.Context, user entity.User) (interface{}, error) {
	logger.LogEvent("INFO", "Persisting user with reference: "+user.UserReference)

	_, err := r.Collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	logger.LogEvent("INFO", "Persisting user with reference: "+user.UserReference+" completed successfully...")
	return user.UserReference, nil
}

func (r *UserInfra) GetUserByReference(ctx context.Context, user_reference string) (interface{}, error) {

	user := entity.User{}
	filter := bson.M{"user_reference": user_reference}

	err := r.Collection.FindOne(ctx, filter).Decode(&user)

	if err != nil {
		return nil, err
	}

	logger.LogEvent("INFO", "Retrieving user with user reference: "+user_reference+" completed successfully. ")

	return user, nil
}
func (r *UserInfra) GetUserByPhone(ctx context.Context, phone string) (interface{}, error) {
	filter := bson.M{"user_profile.phone": phone}

	user := entity.User{}
	err := r.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	logger.LogEvent("INFO", "Retrieving user with user phone: "+phone+" completed successfully. ")
	return user, nil
}

func (r *UserInfra) GetUserByWallsTag(ctx context.Context, wallsTag string) (interface{}, error) {
	filter := bson.M{
		"$or": []bson.M{
			{
				"user_profile.walls_badge.walls_tag": wallsTag,
			},
			{
				"company_profile": bson.M{
					"$elemMatch": bson.M{
						"walls_badge.walls_tag": wallsTag,
					},
				},
			},
		},
	}

	user := entity.User{}
	err := r.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	logger.LogEvent("INFO", "Retrieving user with walls tag: "+wallsTag+" completed successfully. ")
	return user, nil
}

func (r *UserInfra) GetUserByWallsBadgeReference(ctx context.Context, wallsBadgeReference string) (interface{}, error) {
	filter := bson.M{
		"$or": []bson.M{
			{
				"user_profile.walls_badge.walls_badge_reference": wallsBadgeReference,
			},
			{
				"company_profile": bson.M{
					"$elemMatch": bson.M{
						"walls_badge.walls_badge_reference": wallsBadgeReference,
					},
				},
			},
		},
	}

	user := entity.User{}
	err := r.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	logger.LogEvent("INFO", "Retrieving user with walls tag: "+wallsBadgeReference+" completed successfully. ")
	return user, nil
}

func (r *UserInfra) GetUserDefaultWallsBadge(ctx context.Context, userReference string) (interface{}, error) {
	// Define a filter to get the default walls badge of a specific user.
	filter := bson.M{
		"user_reference":                      userReference,
		"user_profile.walls_badge.is_default": true,
	}

	// Define the user entity to store the fetched data.
	user := entity.User{}
	err := r.Collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		return nil, err
	}

	defaultBadge := user.UserProfile.WallsBadge

	logger.LogEvent("INFO", fmt.Sprintf("Retrieving default walls badge for user %s completed successfully.", userReference))
	return defaultBadge, nil
}

func (r *UserInfra) GetUserByDevice(ctx context.Context, device entity.Device) (interface{}, error) {
	filter := bson.M{"device.type": device.Type, "device.imei": device.Imei, "device.brand": device.Brand, "device.model": device.Model}

	user := entity.User{}
	err := r.Collection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	logger.LogEvent("INFO", "Retrieving user with device reference: "+device.DeviceReference+" completed successfully. ")
	return user, nil
}

func (r *UserInfra) UpdateUser(ctx context.Context, user_reference string, user entity.User) (interface{}, error) {
	logger.LogEvent("INFO", "Updating user with reference: "+user_reference)

	filter := bson.M{"user_reference": user_reference}
	update := bson.M{"$set": bson.M{
		"user_profile":         user.UserProfile,
		"wallet":               user.Wallet,
		"bank_accounts":        user.BankAccounts,
		"cards":                user.Cards,
		"kyc.documentations":  user.Kyc.Documentations,
		"enabled":              user.IsActive,
		"notification_options": user.NotificationOptions,
		"device":               user.Device,
		"contacts":             user.Contacts,
		"updated_on":           time.Now().Format(time.RFC3339),
		"company_profile":      user.CompanyProfile,
	}}

	_, err := r.Collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	logger.LogEvent("INFO", "User with reference "+user_reference+" updated successfully")
	return user_reference, nil
}
