package publisher

import (
	"context"
	helper "walls-user-service/internal/core/helper/event-helper"

	"github.com/redis/go-redis/v9"
)

type EventPublisher struct {
	redisClient *redis.Client
}

func NewPublisher(redisClient *redis.Client) *EventPublisher {
	return &EventPublisher{
		redisClient: redisClient,
	}
}

func (p *EventPublisher) PublishUserCreatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishWallsTagCreatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}
func (p *EventPublisher) PublishUserNameUpdatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishEmailUpdatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishDobUpdatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishAddressUpdatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishPhotoUpdatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishWalletUpdatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishBankAddedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishBankUpdatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishCardAddedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishCardUpdatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishNotificationOptionsUpdatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishDeviceUpdatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishDocumentationAddedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishDocumentationUpdatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishContactAddedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishCompanyProfileCreatedEvent(ctx context.Context, event interface{}, eventType ...string) error {

	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)

}

func (p *EventPublisher) PublishCompanyWallsBadgeCreatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}
func (p *EventPublisher) PublishUserWallsBadgeCreatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishCompanyProfileUpdatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishCompanyWallsBadgeDisabledEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishUserWallsBadgeDisabledEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishCompanyProfileDisabledEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishCompanyLogoUpdatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishUserProfileEmailStatusUpdatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishUserProfilePhoneStatusUpdatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishCompanyProfileEmailStatusUpdatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishKycStatusUpdatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishDefaultBankSetEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishDefaultCardSetEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}
func (p *EventPublisher) PublishUsernameUpdatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishDOBUpdatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishPhotosUpdatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishBalanceUpdatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishTierUpdatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishCouponAddedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishRewardsUpdatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishKycScoreUpdatedEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishUserEnabledEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishUserDisabledEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}
func (p *EventPublisher) PublishOnboardingCreateRequestEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishCreateIdentityRequestEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishValidateOtpRequestEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishCreateOtpRequestEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event, eventType...)
}

func (p *EventPublisher) PublishValidateIdentityRequestEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}

func (p *EventPublisher) PublishUpgradeTierRequestEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}
func (p *EventPublisher) PublishCreateTransactionRequestEvent(ctx context.Context, event interface{}, eventType ...string) error {
	redisHelper := helper.NewRedisClient(p.redisClient)
	return redisHelper.PublishEvent(ctx, event)
}
