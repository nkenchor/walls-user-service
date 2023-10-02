package ports

import (
	"context"
	"walls-user-service/internal/core/domain/dto"
)

type UserService interface {
	// USER MANAGEMENT
	//--------------------------------------------------------------------------

	// User creation and status
	CreateUser(ctx context.Context, createUserDto dto.CreateUserDto, currentUserDto dto.CurrentUserDto) (interface{}, error)
	EnableUser(ctx context.Context, user_reference string) (interface{}, error)
	DisableUser(ctx context.Context, user_reference string) (interface{}, error)

	// User retrieval by various identifiers
	GetUserByReference(ctx context.Context, user_reference string) (interface{}, error)
	GetUserByWallsTag(ctx context.Context, wallsTag string) (interface{}, error)
	GetUserByWallsBagdeReference(ctx context.Context, wallsBadgeReference string) (interface{}, error)
	GetUserByDevice(ctx context.Context, device dto.DeviceDto) (interface{}, error)

	// User details updates
	UpdateUserName(ctx context.Context, user_reference string, updateUserNameDto dto.UserNameDto, currentUser dto.CurrentUserDto) (interface{}, error)
	UpdateEmail(ctx context.Context, user_reference string, updateEmailDto dto.EmailDto, currentUser dto.CurrentUserDto) (interface{}, error)
	UpdateDateOfBirth(ctx context.Context, user_reference string, updateDobDto dto.DobDto, currentUser dto.CurrentUserDto) (interface{}, error)
	UpdateAddress(ctx context.Context, user_reference string, updateAddressDto dto.AddressDto, currentUser dto.CurrentUserDto) (interface{}, error)
	UpdatePhoto(ctx context.Context, user_reference string, updatePhotoDto dto.PhotoDto, currentUser dto.CurrentUserDto) (interface{}, error)

	// User communication statuses
	UpdateUserProfileEmailStatus(ctx context.Context, user_reference string) (interface{}, error)
	// UpdateUserProfilePhoneStatus(ctx context.Context, user_reference string) (interface{}, error)

	// User finance & rewards management
	UpdateBalance(ctx context.Context, user_reference string, balance dto.BalanceDto, currentUserDto dto.CurrentUserDto) (interface{}, error)
	UpdateWallet(ctx context.Context, user_reference string, updateWalletDto dto.UpdateWalletDto, currentUser dto.CurrentUserDto) (interface{}, error)
	UpdateTier(ctx context.Context, user_reference string, tierDto dto.TierDto, currentUserDto dto.CurrentUserDto) (interface{}, error)
	AddCoupon(ctx context.Context, user_reference string, couponDto dto.CouponDto) (interface{}, error)
	UpdateRewards(ctx context.Context, user_reference string, rewardDto dto.RewardDto) (interface{}, error)

	// KYC and ID related
	AddDocumentation(ctx context.Context, user_reference string, updateDocumentationDto dto.AddDocumentationDto, currentUser dto.CurrentUserDto) (interface{}, error)
	UpdateDocumentation(ctx context.Context, user_reference string, documentation_reference string, updateDocumentationDto dto.AddDocumentationDto, currentUser dto.CurrentUserDto) (interface{}, error)

	// DEVICE & NOTIFICATION MANAGEMENT
	//---------------------------------------------------------------------------

	UpdateDevice(ctx context.Context, user_reference string, updateDeviceDto dto.UpdateDeviceDto, currentUser dto.CurrentUserDto) (interface{}, error)
	UpdateNotificationOptions(ctx context.Context, user_reference string, updateNotificationOptionsDto dto.UpdateNotificationOptionsDto, currentUser dto.CurrentUserDto) (interface{}, error)

	// COMPANY MANAGEMENT
	//---------------------------------------------------------------------------

	// Company profile creation & updates
	CreateCompanyProfile(ctx context.Context, user_reference string, createCompanyProfileDto dto.CreateCompanyProfileDto, currentUser dto.CurrentUserDto) (interface{}, error)
	UpdateCompanyProfile(ctx context.Context, user_reference string, companyProfileReference string, updateCompanyProfileDto dto.UpdateCompanyProfileDto, currentUser dto.CurrentUserDto) (interface{}, error)
	UpdateCompanyLogo(ctx context.Context, user_reference string, companyProfileReference string, updateCompanyLogoDto dto.UpdateCompanyLogo, currentUser dto.CurrentUserDto) (interface{}, error)
	DisableCompanyProfile(ctx context.Context, user_reference string, companyProfileReference string) (interface{}, error)

	// Company communication statuses
	UpdateCompanyProfileEmailStatus(ctx context.Context, user_reference string, companyProfileReference string) (interface{}, error)

	// COMPANY BADGES
	//---------------------------------------------------------------------------

	// Company walls badge management
	CreateCompanyWallsBadge(ctx context.Context, user_reference string, companyWallsBadgeDto dto.CompanyWallsBadgeDto, currentUser dto.CurrentUserDto) (interface{}, error)
	DisableCompanyWallsBadge(ctx context.Context, user_reference string, companyProfileReference string, companyWallsBadgeReference string, currentUserDto dto.CurrentUserDto) (interface{}, error)
	GetCompanyWallsBadgeList(ctx context.Context, user_reference string, companyProfileReference string) (interface{}, error)
	GetDefaultCompanyWallsBadge(ctx context.Context, user_reference string, companyProfileReference string) (interface{}, error)

	// USER BADGES
	//---------------------------------------------------------------------------

	// User walls badge management
	CreateUserWallsBadge(ctx context.Context, user_reference string, userWallsBadgeDto dto.UserWallsBadgeDto, currentUser dto.CurrentUserDto) (interface{}, error)
	DisableUserWallsBadge(ctx context.Context, user_reference string, userWallsBadgeReference string, currentUserDto dto.CurrentUserDto) (interface{}, error)
	GetUserWallsBadgeList(ctx context.Context, user_reference string) (interface{}, error)
	GetDefaultUserWallsBadge(ctx context.Context, user_reference string) (interface{}, error)

	// PAYMENT METHODS
	//---------------------------------------------------------------------------

	// Bank management
	AddBank(ctx context.Context, user_reference string, updateBankDto dto.BankDto, currentUser dto.CurrentUserDto) (interface{}, error)
	UpdateBank(ctx context.Context, user_reference string, bank_reference string, updateBankDto dto.BankDto, currentUser dto.CurrentUserDto) (interface{}, error)
	SetDefaultBank(ctx context.Context, user_reference string, bankReference string, currentUserDto dto.CurrentUserDto) (interface{}, error)

	// Card management
	AddCard(ctx context.Context, user_reference string, updateCardDto dto.CardDto, currentUser dto.CurrentUserDto) (interface{}, error)
	UpdateCard(ctx context.Context, user_reference string, card_reference string, updateCardDto dto.CardDto, currentUser dto.CurrentUserDto) (interface{}, error)
	SetDefaultCard(ctx context.Context, user_reference string, cardReference string, currentUserDto dto.CurrentUserDto) (interface{}, error)

	// CONTACT MANAGEMENT
	//---------------------------------------------------------------------------

	AddContact(ctx context.Context, user_reference string, updateContactDto dto.ContactDto, currentUser dto.CurrentUserDto) (interface{}, error)

	// REQUESTS MANAGEMENT
	//-----------------------------------------------------------------------------

	CreateIdentityRequest(ctx context.Context, user_reference string, requestIdentityDto dto.IdentityDto, currentUser dto.CurrentUserDto) (interface{}, error)
	CreateOtpRequest(ctx context.Context, requestOtpDto dto.CreateOtpDto, currentUser dto.CurrentUserDto) (interface{}, error)
	ValidateOtpRequest(ctx context.Context, user_reference string, validateOtpDto dto.ValidateOtpDto, currentUserDto dto.CurrentUserDto) (interface{}, error)
	UpgradeTierRequest(ctx context.Context, user_reference string, requestTierDto dto.TierUpgradeRequestDto, currentUser dto.CurrentUserDto) (interface{}, error)

	// Implement these
	GetUserByPhone(ctx context.Context, phone string) (interface{}, error)
	CreateUserReference(ctx context.Context) (interface{}, error)
	CreateDocumentReference(ctx context.Context) (interface{}, error)
	CreateTransactionRequest(ctx context.Context, user_reference string, createTransactionDto dto.CreateTransactionDto, currentUser dto.CurrentUserDto) (interface{}, error)
	CreateDocumentation(ctx context.Context, user_reference string, documentReferenceDto dto.DocumentReferenceDto)(interface{}, error)
}
