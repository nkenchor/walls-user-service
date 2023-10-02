package entity

type User struct {
	UserReference       string              `json:"user_reference" bson:"user_reference"`
	CreatedOn           string              `json:"created_on" bson:"created_on"`
	UpdatedOn           string              `json:"updated_on" bson:"updated_on"`
	IsActive            bool                `json:"is_active" bson:"is_active"`
	UserProfile         UserProfile         `json:"user_profile" bson:"user_profile"`
	CompanyProfile      []CompanyProfile    `json:"company_profile" bson:"company_profile"`
	Contacts            []Contact           `json:"contacts" bson:"contacts"`
	Security            Security            `json:"security" bson:"security"`
	Wallet              Wallet              `json:"wallet" bson:"wallet"`
	BankAccounts        []Bank              `json:"bank_accounts" bson:"bank_accounts"`
	Cards               []Card              `json:"cards" bson:"cards"`
	NotificationOptions NotificationOptions `json:"notification_options" bson:"notification_options"`
	Device              Device              `json:"device" bson:"device"`
	Kyc                 Kyc                 `json:"kyc" bson:"kyc"`
}

type UserProfile struct {
	UserProfileReference string       `json:"user_profile_reference" bson:"user_profile_reference"`
	FirstName            string       `json:"first_name" bson:"first_name"`
	LastName             string       `json:"last_name" bson:"last_name"`
	FullName             string       `json:"full_name" bson:"full_name"`
	DateOfBirth          string       `json:"date_of_birth" bson:"date_of_birth"`
	Phone                string       `json:"phone" bson:"phone"`
	Email                string       `json:"email" bson:"email"`
	Address              Address      `json:"address" bson:"address"`
	Photos               []Photo      `json:"photos" bson:"photos"`
	IsVerifiedEmail      bool         `json:"is_verified_email" bson:"is_verified_email"`
	IsVerifiedPhone      bool         `json:"is_verified_phone" bson:"is_verified_phone"`
	WallsBadge           []WallsBadge `json:"walls_badge" bson:"walls_badge"`
}
type CompanyProfile struct {
	CompanyProfileReference string       `json:"company_profile_reference" bson:"company_profile_reference"`
	CompanyName             string       `json:"company_name" bson:"company_name"`
	RegistrationNo          string       `json:"registration_no" bson:"registration_no"`
	RegistrationDate        string       `json:"registration_date" bson:"registration_date"`
	Phone                   string       `json:"phone" bson:"phone"`
	Email                   string       `json:"email" bson:"email"`
	Address                 Address      `json:"address" bson:"address"`
	Logo                    Photo        `json:"logo" bson:"logo"`
	IsVerifiedEmail         bool         `json:"is_verified_email" bson:"is_verified_email"`
	IsActive                bool         `json:"is_active" bson:"is_active"`
	WallsBadge              []WallsBadge `json:"walls_badge" bson:"walls_badge"`
}

type Country struct {
	CountryReference string `json:"country_reference" bson:"country_reference" validate:"required,guid"`
	CountryName      string `json:"country_name" bson:"country_name" validate:"required,alpha,min=3"`
	CountryCode      string `json:"country_code" bson:"country_code" validate:"required,len=3"`
	DialCode         string `json:"dial_code" bson:"dial_code" validate:"required,min=2,startswith=+"`
}

type Address struct {
	AddressReference string   `json:"address_reference" bson:"address_reference" validate:"required,guid"`
	Country          Country  `json:"country" bson:"country" validate:"required"`
	State            State    `json:"state" bson:"state" validate:"required"`
	City             string   `json:"city" bson:"city" validate:"required,min=3"`
	AddressLines     []string `json:"address_lines" bson:"address_lines" validate:"required,min=1"`
}

type State struct {
	StateReference string `json:"state_reference" bson:"state_reference" validate:"required,guid"`
	StateName      string `json:"state_name" bson:"state_name" validate:"required,min=3"`
}

type Photo struct {
	PhotoReference    string `json:"photo_reference" bson:"photo_reference" validate:"required,guid"`
	IsDefault         bool   `json:"is_default" bson:"is_default" validate:"boolean"`
	IsVerified        bool   `json:"is_verified" bson:"is_verified" validate:"boolean"`
	DocumentReference string `json:"document_reference" bson:"document_reference" validate:"required,guid"`
}

type Contact struct {
	ContactReference string `json:"contact_reference" bson:"contact_reference"`
	Phone            string `json:"phone" bson:"phone"`
	FullName         string `json:"full_name" bson:"full_name"`
	WallsTag         string `json:"walls_tag" bson:"walls_tag"`
	IsBeneficiary    bool   `json:"is_beneficiary" bson:"is_beneficiary"`
}

type Security struct {
	TwoFactorAuthEnabled  bool   `json:"two_factor_auth_enabled" bson:"two_factor_auth_enabled"`
	TwoFactorAuthMethod   string `json:"two_factor_auth_method" bson:"two_factor_auth_method"`
	FailedLoginAttempts   int    `json:"failed_login_attempts" bson:"failed_login_attempts"`
	LastFailedLoginOn     string `json:"last_failed_login_on" bson:"last_failed_login_on"`
	PasswordLastUpdatedOn string `json:"password_last_updated_on" bson:"password_last_updated_on"`
	TransactionPinEnabled bool   `json:"transaction_pin_enabled" bson:"transaction_pin_enabled"`
	FingerprintEnabled    bool   `json:"fingerprint_enabled" bson:"fingerprint_enabled"`
	FaceIdEnabled         bool   `json:"face_id_enabled" bson:"face_id_enabled"`
}

type Coupon struct {
	CouponReference     string  `json:"coupon_reference" bson:"coupon_reference" validate:"required,guid"`
	CouponId            string  `json:"coupon_id" bson:"coupon_id" validate:"required"`
	Discount_Percentage float64 `json:"discount_percentage" bson:"discount_percentage" validate:"required,number,min=1"`
	ExpiryDate          string  `json:"expiry_date" bson:"expiry_date" validate:"required"`
}

type Reward struct {
	Points int `json:"points" bson:"points" validate:"required,min=1"`
}

type WallsBadge struct {
	WallsBadgeReference string `json:"walls_badge_reference" bson:"walls_badge_reference"`
	WalletReference     string `json:"wallet_reference" bson:"wallet_reference"`
	UserReference       string `json:"user_reference" bson:"user_reference"`
	WallsTag            string `json:"walls_tag" bson:"walls_tag"`
	IsDefault           bool   `json:"is_default" bson:"is_default"`
	IsActive            bool   `json:"is_enabled" bson:"is_enabled"`
	ActiveFor           string `json:"active_for" bson:"active_for"`
}

type Wallet struct {
	WalletReference     string   `json:"wallet_reference" bson:"wallet_reference"`
	WallsAccountNo      int16    `json:"wallet_account_no" bson:"wallet_account_no"`
	AutoFund            bool     `json:"auto_fund" bson:"auto_fund"`
	AutoFundLevel       float64  `json:"auto_fund_level" bson:"auto_fund_level"`
	AutoFundLimit       float64  `json:"auto_fund_limit" bson:"auto_fund_limit" validate:"gte=0"`
	AutoWithdrawal      bool     `json:"auto_withdrawal" bson:"auto_withdrawal"`
	AutoWithdrawalLevel float64  `json:"auto_withdrawal_level" bson:"auto_withdrawal_level"`
	AutoWithdrawalLimit float64  `json:"auto_withdrawal_limit" bson:"auto_withdrawal_limit" validate:"gte=0"`
	Tier                Tier     `json:"tier" bson:"tier"`
	Coupons             []Coupon `json:"coupons" bson:"coupons"`
	Reward              Reward   `json:"reward" bson:"reward"`
	Balance             Balance  `json:"balance" bson:"balance"`
}

type Tier struct {
	TierReference         string   `json:"reference" bson:"reference" validate:"required,uuid"`
	TierName              string   `json:"name" bson:"name" validate:"required"`
	SendingLimit          float64  `json:"sending_limit" bson:"sending_limit" validate:"gte=0"`
	ReceivingLimit        float64  `json:"receiving_limit" bson:"receiving_limit" validate:"gte=0"`
	WalletLimit           float64  `json:"wallet_limit" bson:"wallet_limit" validate:"gte=0"`
	MinimumBalance        float64  `json:"minimum_balance" bson:"minimum_balance" validate:"gte=0"`
	DailyTransactionLimit float64  `json:"transaction_limit" bson:"transaction_limit" validate:"gt=0"`
	UpgradeOptions        []string `json:"upgrade_options" bson:"upgrade_options"`
}

type Balance struct {
	AvailableAmount       float64 `json:"available_amount" bson:"available_amount"`
	Currency              string  `json:"currency" bson:"currency"`
	PendingIncomingAmount float64 `json:"pending_incoming_amount" bson:"pending_incoming_amount"`
	IsSynced              bool    `json:"is_synced" bson:"is_synced"`
	LastSyncedOn          string  `json:"last_synced_on" bson:"last_synced_on"`
}

type Bank struct {
	BankReference        string `json:"bank_reference" bson:"bank_reference"`
	IntegrationType      string `json:"integration_type" bson:"integration_type"`
	IntegrationReference string `json:"integration_reference" bson:"integration_reference"`
	BankName             string `json:"bank_name" bson:"bank_name"`
	AccountNumber        int64  `json:"account_number" bson:"account_number"`
	AccountName          string `json:"account_name" bson:"account_name"`
	IsDefault            bool   `json:"is_default" bson:"is_default"`
}

type Card struct {
	CardReference        string `json:"card_reference" bson:"card_reference"`
	IntegrationType      string `json:"integration_type" bson:"integration_type"`
	IntegrationReference string `json:"integration_reference" bson:"integration_reference"`
	CardName             string `json:"card_name" bson:"card_name"`
	Pan                  string `json:"pan" bson:"pan"`
	ExpiryMonth          int    `json:"expiry_month" bson:"expiry_month"`
	ExpiryYear           int    `json:"expiry_year" bson:"expiry_year"`
	IsDefault            bool   `json:"is_default" bson:"is_default"`
}

type NotificationOptions struct {
	PushNotificationEnabled bool   `json:"push_notification_enabled" bson:"push_notification_enabled"`
	NotificationType        string `json:"notification_type" bson:"notification_type" validate:"eq=instant|eq=scheduled"`
	OtpChannel              string `json:"otp_channel" bson:"otp_channel" validate:"required,eq=sms|eq=email"`
}

type Device struct {
	DeviceReference string `json:"device_reference" bson:"device_reference" validate:"required,guid"`
	Imei            string `json:"imei" bson:"imei" validate:"required,imei"`
	Type            string `json:"type" bson:"type" validate:"required,eq=mobile|eq=tablet|eq=desktop|eq=phablet|eq=smart_watch"`
	Brand           string `json:"brand" bson:"brand" validate:"required,alpha"`
	Model           string `json:"model" bson:"model" validate:"required"`
}

type Documentation struct {
	DocumentationReference string `json:"documentation_reference" bson:"documentation_reference"`
	DocumentationType      string `json:"documentation_type" bson:"documentation_type" validate:"required,eq=international_passport|eq=bvn|eq=nin|drivers_license|voters_card"`
	DocumentationNumber    string `json:"documentation_number" bson:"documentation_number"`
	Expiry                  string `json:"expiry" bson:"expiry"`
	DocumentReference       string `json:"document_reference" bson:"document_reference"`
	TierReference           string `json:"tier_reference" bson:"tier_reference" validate:"required,uuid4"`
	IsVerified              bool   `json:"is_verified" bson:"is_verified"`
	VerifiedOn              string `json:"verified_on" bson:"verified_on"`
	VerificationMethod      string `json:"verification_method" bson:"verification_method"`
}

type Kyc struct {
	Documentations []Documentation `json:"documentations" bson:"documentations"`
	ProfileType     string           `json:"profile_type" bson:"profile_type" validate:"required,eq=user|eq=company"`
}
type Location struct {
	Longitude float64 `json:"longitude" bson:"longitude" validate:"required"`
	Latitude  float64 `json:"latitude" bson:"latitude" validate:"required"`
}
type TierRequest struct {
	RequestReference string `json:"request_reference" bson:"request_reference" validate:"required,uuid"`
	UserReference    string `json:"user_reference" bson:"user_reference" validate:"required,uuid4"`
	CurrentTier      Tier   `json:"current_tier" bson:"current_tier" validate:"required,dive"`
	RequestedTier    Tier   `json:"requested_tier" bson:"requested_tier" validate:"required,dive"`
}
type Sender struct {
	Type          string `json:"type" bson:"type" validate:"required,eq=wallet|eq=bank_account|eq=card"`
	UserReference string `json:"user_reference" bson:"user_reference" validate:"required,uuid4"`
	Reference     string `json:"reference" bson:"reference" validate:"required,uuid4"`
	Device        Device `json:"device" bson:"device" validate:"required,dive"`
}

type Receiver struct {
	Type                string `json:"type" bson:"type" validate:"required,eq=wallet|eq=bank_account|eq=card"`
	UserReference       string `json:"user_reference" bson:"user_reference" validate:"required,uuid4"`
	Reference           string `json:"reference" bson:"reference" validate:"required,uuid4"`
	WallsBadgeReference string `json:"wallsbadge_reference" bson:"wallsbadge_reference" validate:"required,uuid4"`
}
type Metadata struct {
	Note      string    `json:"note" bson:"note" validate:"required,max=200"`
	RelatedTx []string  `json:"related_tx,omitempty" bson:"related_tx,omitempty" validate:"omitempty,dive,uuid4"`
	Location  *Location `json:"location,omitempty" bson:"location,omitempty" validate:"omitempty,dive"`
}
