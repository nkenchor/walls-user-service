package event

import (
	"walls-user-service/internal/core/helper/event-helper/eto"
)

type UserCreatedEvent struct {
	eto.Event
}

type UserWallsBadgeCreatedEvent struct {
	eto.Event
}

type CompanyWallsBadgeCreatedEvent struct {
	eto.Event
}
type CompanyProfileCreatedEvent struct {
	eto.Event
}

type CompanyProfileUpdatedEvent struct {
	eto.Event
}

type CompanyWallsBadgeDisabledEvent struct {
	eto.Event
}

type UserWallsBadgeDisabledEvent struct {
	eto.Event
}

type CompanyProfileDisabledEvent struct {
	eto.Event
}

type CompanyLogoUpdatedEvent struct {
	eto.Event
}

type UserProfileEmailStatusUpdatedEvent struct {
	eto.Event
}

type UserProfilePhoneStatusUpdatedEvent struct {
	eto.Event
}

type CompanyProfileEmailStatusUpdatedEvent struct {
	eto.Event
}

type UsernameUpdatedEvent struct {
	eto.Event
}

type KycStatusUpdatedEvent struct {
	eto.Event
}

type DefaultBankSetEvent struct {
	eto.Event
}

type DefaultCardSetEvent struct {
	eto.Event
}

type EmailUpdatedEvent struct {
	eto.Event
}

type DOBUpdatedEvent struct {
	eto.Event
}

type AddressUpdatedEvent struct {
	eto.Event
}

type PhotosUpdatedEvent struct {
	eto.Event
}

type WalletUpdatedEvent struct {
	eto.Event
}

type BankAddedEvent struct {
	eto.Event
}

type BankUpdatedEvent struct {
	eto.Event
}

type CardAddedEvent struct {
	eto.Event
}

type CardUpdatedEvent struct {
	eto.Event
}

type NotificationOptionsUpdatedEvent struct {
	eto.Event
}

type DeviceUpdatedEvent struct {
	eto.Event
}

type DocumentationAddedEvent struct {
	eto.Event
}

type DocumentationUpdatedEvent struct {
	eto.Event
}

type ContactAddedEvent struct {
	eto.Event
}

type BookBalanceUpdatedEvent struct {
	eto.Event
}

type BalanceUpdatedEvent struct {
	eto.Event
}

type BookTransferredtoAvailableEvent struct {
	eto.Event
}

type TierUpdatedEvent struct {
	eto.Event
}

type CouponAddedEvent struct {
	eto.Event
}

type RewardsUpdatedEvent struct {
	eto.Event
}

type KycScoreUpdatedEvent struct {
	eto.Event
}

type UserEnabledEvent struct {
	eto.Event
}

type UserDisabledEvent struct {
	eto.Event
}
type OtpRequestCreatedEvent struct {
	eto.Event
}
type ValidateOtpRequestEvent struct {
	eto.Event
}
type CreateIdentityRequestEvent struct {
	eto.Event
}
type TierUpgradeRequestEvent struct {
	eto.Event
}
type TransactionCreateRequest struct {
	eto.Event
}
