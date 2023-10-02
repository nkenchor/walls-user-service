package dto

import (
	"walls-user-service/internal/core/domain/entity"
)

type CreateUserDto struct {
	Phone string `json:"phone" bson:"phone" validate:"required,valid_phone"`
}

type BookBalanceDto struct {
	BookAmount float64 `json:"book_amount" bson:"book_amount" validate:"required,number"`
}

type BalanceDto struct {
	BookAmount      float64 `json:"book_amount" bson:"book_amount" validate:"required,number"`
	AvailableAmount float64 `json:"available_amount" bson:"available_amount" validate:"required,number"`
}

type KycScoreDto struct {
	Score float64 `json:"kyc_score" bson:"kyc_score" validate:"required,number"`
}
type CompanyWallsBadgeDto struct {
	WallsTag                string `json:"walls_tag" bson:"walls_tag" validate:"required,len=8"`
	WalletReference         string `json:"wallet_reference" bson:"wallet_reference" validate:"required,uuid4"`
	UserReference           string `json:"user_reference" bson:"user_reference" validate:"required,uuid4"`
	CompanyProfileReference string `json:"company_profile_reference" bson:"company_profile_reference" validate:"required,guid,min=32,max=38"`
}
type UserWallsBadgeDto struct {
	WallsTag        string `json:"walls_tag" bson:"walls_tag" validate:"required,len=8"`
	WalletReference string `json:"wallet_reference" bson:"wallet_reference" validate:"required,uuid4"`
	UserReference   string `json:"user_reference" bson:"user_reference" validate:"required,uuid4"`
}

type WallsAccountNoDto struct {
	WallsAccountNo int16 `json:"walls_account_no" bson:"walls_account_no" validate:"required,len=10"`
}

type UserNameDto struct {
	FirstName string `json:"first_name" bson:"first_name" validate:"required,min=3,max=20"`
	LastName  string `json:"last_name" bson:"last_name" validate:"required,min=3,max=20"`
}
type EmailDto struct {
	Email string `json:"email" bson:"email" validate:"required,valid_email"`
}

type DobDto struct {
	DateOfBirth string `json:"date_of_birth" bson:"date_of_birth" validate:"required"`
}

type AddressDto struct {
	Address entity.Address `json:"address" bson:"address" validate:"required"`
}
type PhotoDto struct {
	Photo entity.Photo `json:"photo" bson:"photo" validate:"required"`
}

type ContactDto struct {
	Phone         string `json:"phone" bson:"phone" validate:"required,valid_contact"`
	FullName      string `json:"full_name" bson:"full_name" validate:"required,min=3"`
	WallsTag      string `json:"walls_tag" bson:"walls_tag" validate:"required,len=8"`
	IsBeneficiary bool   `json:"is_beneficiary" bson:"is_beneficiary" validate:"boolean"`
}

type UpdateWalletDto struct {
	AutoFund            bool    `json:"auto_fund" bson:"auto_fund" validate:"boolean"`
	AutoFundLevel       float64 `json:"auto_fund_level" bson:"auto_fund_level" validate:"number"`
	AutoFundLimit       float64 `json:"auto_fund_limit" bson:"auto_fund_limit" validate:"number"`
	AutoWithdrawal      bool    `json:"auto_withdrawal" bson:"auto_withdrawal" validate:"boolean"`
	AutoWithdrawalLevel float64 `json:"auto_withdrawal_level" bson:"auto_withdrawal_level" validate:"number"`
	AutoWithdrawalLimit float64 `json:"auto_withdrawal_limit" bson:"auto_withdrawal_limit" validate:"number"`
}

type BankDto struct {
	BankName      string `json:"bank_name" bson:"bank_name" validate:"required,min=3"`
	AccountNumber int64  `json:"account_number" bson:"account_number" validate:"required,number,min=1000000000"`
	AccountName   string `json:"account_name" bson:"account_name" validate:"required,min=3"`
	IsDefault     bool   `json:"is_default" bson:"is_default" validate:"boolean"`
}

type CardDto struct {
	CardName    string `json:"card_name" bson:"card_name" validate:"required,min=3"`
	Pan         string `json:"pan" bson:"pan" validate:"required,len=16"`
	ExpiryMonth int    `json:"expiry_month" bson:"expiry_month" validate:"required,number,min=1,max=12"`
	ExpiryYear  int    `json:"expiry_year" bson:"expiry_year" validate:"required,number"`
	IsDefault   bool   `json:"is_default" bson:"is_default" validate:"boolean"`
}

type UpdateNotificationOptionsDto struct {
	PushNotificationEnabled bool   `json:"push_notification_enabled" bson:"push_notification_enabled" validate:"boolean"`
	NotificationType        string `json:"notification_type" bson:"notification_type" validate:"eq=instant|eq=scheduled"`
	OtpType                 string `json:"otp_type" bson:"otp_type" validate:"required,eq=sms|eq=email|in_app"`
}

type UpdateDeviceDto struct {
	NewDevice entity.Device `json:"new_device" bson:"new_device" validate:"required,dive"`
}

type AddDocumentationDto struct {
	DocumentationType   string `json:"documentation_type" bson:"documentation_type" validate:"required,eq=international_passport|eq=bvn|eq=nin|drivers_license|voters_card"`
	DocumentationNumber string `json:"documentation_number" bson:"documentation_number" validate:"required,alpha"`
	Expiry               string `json:"expiry" bson:"expiry" validate:"required"`
	DocumentReference    string `json:"document_reference" bson:"document_reference" validate:"required,guid"`
	TierReference        string `json:"tier_reference" bson:"tier_reference" validate:"required,uuid4"`
}

type CurrentUserDto struct {
	UserReference string    `json:"user_reference" bson:"user_reference" validate:"required,guid,min=32,max=38"`
	Phone         string    `json:"phone" bson:"phone" validate:"required,valid_contact"`
	Device        DeviceDto `json:"device" bson:"device" validate:"required,dive"`
}

type DeviceDto struct {
	DeviceReference string `json:"device_reference" bson:"device_reference" validate:"required,guid,min=32,max=38"`
	Imei            string `json:"imei" bson:"imei" validate:"required,imei,min=10,max=50"`
	Type            string `json:"type" bson:"type" validate:"required,eq=mobile|eq=tablet|eq=desktop|eq=phablet|eq=smart_watch"`
	Brand           string `json:"brand" bson:"brand" validate:"required,alpha"`
	Model           string `json:"model" bson:"model" validate:"required,alpha"`
}

type UserDto struct {
	UserReference       string                     `json:"user_reference" bson:"user_reference"`
	CreatedOn           string                     `json:"created_on" bson:"created_on"`
	UpdatedOn           string                     `json:"updated_on" bson:"updated_on"`
	IsActive            bool                       `json:"is_active" bson:"is_active"`
	UserProfile         entity.UserProfile         `json:"user_profile" bson:"user_profile"`
	CompanyProfile      []entity.CompanyProfile    `json:"company_profile" bson:"company_profile"`
	Contacts            []entity.Contact           `json:"contacts" bson:"contacts"`
	Security            entity.Security            `json:"security" bson:"security"`
	Wallet              WalletDto                  `json:"wallet" bson:"wallet"`
	BankAccounts        []entity.Bank              `json:"bank_accounts" bson:"bank_accounts"`
	Cards               []entity.Card              `json:"cards" bson:"cards"`
	NotificationOptions entity.NotificationOptions `json:"notification_options" bson:"notification_options"`
	Device              entity.Device              `json:"device" bson:"device"`
	Kyc                 KycDto                     `json:"kyc" bson:"kyc"`
	LastSyncedOn        string                     `json:"last_synced_on" bson:"last_synced_on"`
}

type WalletDto struct {
	WalletReference      string          `json:"wallet_reference" bson:"wallet_reference" validate:"required,guid,min=32,max=38"`
	WallsAccountNo       int16           `json:"wallet_account_no" bson:"wallet_account_no" validate:"required,number,min=1000000000"`
	AutoFund             bool            `json:"auto_fund" bson:"auto_fund" validate:"boolean"`
	AutoFundLevel        float64         `json:"auto_fund_level" bson:"auto_fund_level" validate:"required,number"`
	AutoFundAmount       float64         `json:"auto_fund_amount" bson:"auto_fund_amount" validate:"required,number"`
	AutoWithdrawal       bool            `json:"auto_withdrawal" bson:"auto_withdrawal" validate:"boolean"`
	AutoWithdrawalLevel  float64         `json:"auto_withdrawal_level" bson:"auto_withdrawal_level" validate:"required,number"`
	AutoWithdrawalAmount float64         `json:"auto_withdrawal_amount" bson:"auto_withdrawal_amount" validate:"required,number"`
	Tier                 TierDto         `json:"tier" bson:"tier" validate:"required"`
	Coupons              []entity.Coupon `json:"coupons" bson:"coupons" validate:"required,min=1"`
	Reward               entity.Reward   `json:"reward" bson:"reward" validate:"required"`
	Balance              entity.Balance  `json:"balance" bson:"balance" validate:"required"`
}

type TierDto struct {
	TierReference         string   `json:"reference" bson:"reference" validate:"required,uuid"`
	TierName              string   `json:"name" bson:"name" validate:"required"`
	SendingLimit          float64  `json:"sending_limit" bson:"sending_limit" validate:"gte=0"`
	MinimumBalance        float64  `json:"minimum_balance" bson:"minimum_balance" validate:"gte=0"`
	DailyTransactionLimit float64  `json:"transaction_limit" bson:"transaction_limit" validate:"gt=0"`
	UpgradeOptions        []string `json:"upgrade_options" bson:"upgrade_options"`
	

}

type KycDto struct {
	Documentations []DocumentationDto `json:"documentations" bson:"documentations" validate:"required,min=1,dive"`
}

type DocumentationDto struct {
	DocumentationReference string `json:"documentation_reference" bson:"documentation_reference" validate:"required,guid"`
	DocumentationType      string `json:"documentation_type" bson:"documentation_type" validate:"required,eq=international_passport|eq=bvn|eq=nin|drivers_license|voters_card"`
	DocumentationNumber    string `json:"documentation_number" bson:"documentation_number" validate:"required,alpha"`
	Expiry                 string `json:"expiry" bson:"expiry" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	DocumentReference      string `json:"document_reference" bson:"document_reference" validate:"required,guid"`
	TierReference          string `json:"tier_reference" bson:"tier_reference" validate:"required,uuid4"`
	IsVerified             bool   `json:"is_verified" bson:"is_verified" validate:"boolean"`
	VerifiedOn             string `json:"verified_on" bson:"verified_on" validate:"required,datetime=2006-01-02T15:04:05Z07:00"`
	VerificationMethod     string `json:"verification_method" bson:"verification_method" validate:"eq=manual|eq=auto"`
}

type RewardDto struct {
	Points int `json:"points" bson:"points" validate:"required,min=1"`
}

type CouponDto struct {
	CouponReference     string  `json:"coupon_reference" bson:"coupon_reference" validate:"required,guid"`
	CouponId            string  `json:"coupon_id" bson:"coupon_id" validate:"required"`
	Discount_Percentage float64 `json:"discount_percentage" bson:"discount_percentage" validate:"required,number,min=1"`
	ExpiryDate          string  `json:"expiry_date" bson:"expiry_date" validate:"required"`
}

type CreateCompanyProfileDto struct {
	CompanyName      string         `json:"company_name" bson:"company_name" validate:"required,min=5,max=50"`
	RegistrationNo   string         `json:"registration_no" bson:"registration_no" validate:"required,min=5,max=50"`
	RegistrationDate string         `json:"registration_date" bson:"registration_date" validate:"required"`
	Phone            string         `json:"phone" bson:"phone" validate:"required,valid_contact"`
	Email            string         `json:"email" bson:"email" validate:"required,valid_email"`
	Address          entity.Address `json:"address" bson:"address" validate:"required"`
	Logo             entity.Photo   `json:"logo" bson:"logo" validate:"required"`
	IsVerifiedEmail  bool           `json:"is_verified_email" bson:"is_verified_email" validate:"boolean"`
}

type UpdateCompanyProfileDto struct {
	Phone   string         `json:"phone" bson:"phone" validate:"required,valid_contact"`
	Email   string         `json:"email" bson:"email" validate:"required,valid_email"`
	Address entity.Address `json:"address" bson:"address" validate:"required"`
}

type UpdateCompanyLogo struct {
	Logo entity.Photo `json:"logo" bson:"logo" validate:"required"`
}

type IdentityDto struct {
	Phone  string    `json:"phone" bson:"phone" validate:"required,valid_contact"` // Must be a valid E.164 format phone number
	Device DeviceDto `json:"device" bson:"device" validate:"required"`
}
type CreateOtpDto struct {
	OtpType string    `json:"otp_type" bson:"otp_type" validate:"required,eq=create_user|eq=create_company|eq=verify_email|eq=verify_phone"`
	Contact string    `json:"contact" bson:"contact" validate:"required,valid_contact"`
	Channel string    `json:"channel" bson:"channel" validate:"eq=sms|eq=email|eq=in_app"`
	Device  DeviceDto `json:"device" bson:"device" validate:"required,dive"`
}

type ValidateOtpDto struct {
	Otp     string    `json:"otp" bson:"otp" validate:"required,len=6"`
	OtpType string    `json:"otp_type" bson:"otp_type" validate:"required,eq=create_user|eq=create_company|eq=verify_email|eq=verify_phone"`
	Contact string    `json:"contact" bson:"contact" validate:"valid_contact"`
	Device  DeviceDto `json:"device" bson:"device" validate:"required,dive"`
}

type TierUpgradeRequestDto struct {
	RequestReference string             `json:"request_reference" bson:"request_reference" validate:"required,uuid"`
	User             UserDto            `json:"user" bson:"user" validate:"required,dive"`
	RequestedTier    TierDto            `json:"requested_tier" bson:"requested_tier" validate:"required,dive"`
	TierDocuments    []DocumentationDto `json:"tier_document" bson:"tier_document" validate:"required,dive"`
}

type CreateTransactionDto struct {
	TransactionType             string          `json:"transaction_type" bson:"transaction_type" validate:"required,eq=wallet_wallet|eq=bank_wallet|eq=card_wallet|eq=wallet_bank"`
	Amount                      float64         `json:"amount" bson:"amount" validate:"required,gte=0"`
	SenderReference             string          `json:"sender_reference" bson:"sender_reference" validate:"required"`
	OutReference                string          `json:"out_reference" bson:"out_reference" validate:"required"`
	ReceiverReference           string          `json:"receiver_reference" bson:"receiver_reference" validate:"required"`
	InReference                 string          `json:"in_reference" bson:"in_reference" validate:"required"`
	ReceiverWallsBadgeReference string          `json:"receiver_walls_badge_reference" bson:"receiver_walls_badge_reference"`
	Metadata                    entity.Metadata `json:"metadata" bson:"metadata" validate:"required,dive"`
}

type DocumentReferenceDto struct {
	DocumentReference      []string `json:"document_references" bson:"document_references" validate:"required,dive,uuid4"`
}
