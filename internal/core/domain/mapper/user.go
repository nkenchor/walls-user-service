package mapper

import (
	"fmt"
	"time"
	"walls-user-service/internal/core/domain/dto"
	"walls-user-service/internal/core/domain/entity"

	"github.com/google/uuid"
)

func CurrentUserDtoToUser(createUserDto dto.CreateUserDto, currentUser dto.CurrentUserDto) entity.User {
	user := entity.User{
		UserReference: currentUser.UserReference,
		CreatedOn:     time.Now().Format(time.RFC3339),
		UserProfile: entity.UserProfile{
			UserProfileReference: uuid.New().String(),
			Phone:                createUserDto.Phone,
		},
		Wallet: entity.Wallet{
			WalletReference: uuid.New().String(),
			Balance: entity.Balance{
				Currency: "NGN",
			},
		},
		Device: entity.Device(currentUser.Device),
	}
	return user
}

func CreateCompanyWallsBadgeDtoToUser(user entity.User, dto dto.CompanyWallsBadgeDto) entity.User {
	wallsBadge := entity.WallsBadge{
		WallsBadgeReference: uuid.New().String(),
		WalletReference:     dto.WalletReference,
		UserReference:       dto.UserReference,
		WallsTag:            dto.WallsTag,
		IsActive:            false,
		IsDefault:           false,
		ActiveFor:           time.Now().Add(time.Hour * 24 * 365).Format(time.RFC3339),
	}

	for index, company := range user.CompanyProfile {
		if company.CompanyProfileReference == dto.CompanyProfileReference {
			user.CompanyProfile[index].WallsBadge = append(user.CompanyProfile[index].WallsBadge, wallsBadge)
			break
		}
	}
	return user
}

func CreateUserWallsBadgeDtoToUser(user entity.User, dto dto.UserWallsBadgeDto) entity.User {
	wallsBadge := entity.WallsBadge{
		WallsBadgeReference: uuid.New().String(),
		WallsTag:            dto.WallsTag,
		WalletReference:     dto.WalletReference,
		UserReference:       dto.UserReference,
		IsActive:            false,
		IsDefault:           false,
		ActiveFor:           time.Now().Add(time.Hour * 24 * 365).Format(time.RFC3339),
	}

	// if user does not have a walls badge prior, make the walls badge the default walls badge
	if len(user.UserProfile.WallsBadge) == 0 {
		wallsBadge.IsDefault = true
	}

	user.UserProfile.WallsBadge = append(user.UserProfile.WallsBadge, wallsBadge)

	return user
}

func UpdateUserNameDtoToUser(user entity.User, dto dto.UserNameDto) entity.User {
	user.UserProfile.FirstName = dto.FirstName
	user.UserProfile.LastName = dto.LastName
	user.UserProfile.FullName = fmt.Sprintf("%v, %v", dto.FirstName, dto.LastName)
	return user
}
func UpdateEmailDtoToUser(user entity.User, dto dto.EmailDto) entity.User {
	user.UserProfile.Email = dto.Email
	user.UserProfile.IsVerifiedEmail = false
	return user
}

func UpdateDobDtoToUser(user entity.User, dto dto.DobDto) entity.User {
	user.UserProfile.DateOfBirth = dto.DateOfBirth
	return user
}
func UpdateAddressDtoToUser(user entity.User, dto dto.AddressDto) entity.User {
	user.UserProfile.Address = dto.Address
	return user
}

func UpdatePhotoDtoToUser(user entity.User, dto dto.PhotoDto) entity.User {

	// update photo if there already is one with the specified reference
	for index, photo := range user.UserProfile.Photos {
		if photo.PhotoReference == dto.Photo.PhotoReference {
			user.UserProfile.Photos[index] = dto.Photo
			return user
		}
	}

	// add photo if it is a new photo
	user.UserProfile.Photos = append(user.UserProfile.Photos, dto.Photo)

	return user
}
func UpdateWalletDtoToWallet(user entity.User, dto dto.UpdateWalletDto) entity.User {
	wallet := entity.Wallet{
		AutoFund:            dto.AutoFund,
		AutoFundLevel:       dto.AutoFundLevel,
		AutoFundLimit:       dto.AutoFundLimit,
		AutoWithdrawal:      dto.AutoWithdrawal,
		AutoWithdrawalLevel: dto.AutoWithdrawalLevel,
		AutoWithdrawalLimit: dto.AutoWithdrawalLimit,
	}
	user.Wallet = wallet
	return user
}
func AddBankDtoToBank(user entity.User, dto dto.BankDto) entity.User {
	bank := entity.Bank{
		BankReference: uuid.New().String(),
		BankName:      dto.BankName,
		AccountNumber: dto.AccountNumber,
		AccountName:   dto.AccountName,
		IsDefault:     dto.IsDefault,
	}
	user.BankAccounts = append(user.BankAccounts, bank)
	return user
}
func UpdateBankDtoToBank(dto dto.BankDto, bank_reference string) entity.Bank {
	bank := entity.Bank{
		BankReference: bank_reference,
		BankName:      dto.BankName,
		AccountNumber: dto.AccountNumber,
		AccountName:   dto.AccountName,
		IsDefault:     dto.IsDefault,
	}
	return bank
}
func AddCardDtoToCard(user entity.User, dto dto.CardDto) entity.User {
	card := entity.Card{
		CardReference: uuid.New().String(),
		CardName:      dto.CardName,
		Pan:           dto.Pan,
		ExpiryMonth:   dto.ExpiryMonth,
		ExpiryYear:    dto.ExpiryYear,
		IsDefault:     dto.IsDefault,
	}
	user.Cards = append(user.Cards, card)
	return user
}
func UpdateCardDtoToCard(dto dto.CardDto, card_reference string) entity.Card {
	card := entity.Card{
		CardReference: card_reference,
		CardName:      dto.CardName,
		Pan:           dto.Pan,
		ExpiryMonth:   dto.ExpiryMonth,
		ExpiryYear:    dto.ExpiryYear,
		IsDefault:     dto.IsDefault,
	}
	return card
}
func UpdateNotificationOptionsDtoToNotificationOptions(user entity.User, dto dto.UpdateNotificationOptionsDto) entity.User {
	notificationOptions := entity.NotificationOptions{
		PushNotificationEnabled: dto.PushNotificationEnabled,
		NotificationType:        dto.NotificationType,
		OtpChannel:              dto.OtpType,
	}
	user.NotificationOptions = notificationOptions
	return user
}
func UpdateDeviceDtoToDevice(user entity.User, dto dto.UpdateDeviceDto) entity.User {
	user.Device = dto.NewDevice
	return user
}
func AddDocumentationDtoToDocumentation(user entity.User, dto dto.AddDocumentationDto) entity.User {
	documentation := entity.Documentation{
		DocumentationReference: uuid.New().String(),
		DocumentationType:      dto.DocumentationType,
		DocumentationNumber:    dto.DocumentationNumber,
		Expiry:                  dto.Expiry,
		DocumentReference:       dto.DocumentReference,
		TierReference:           dto.TierReference,
	}
	user.Kyc.Documentations = append(user.Kyc.Documentations, documentation)
	return user
}
func UpdateDocumentationDtoToDocumentation(dto dto.AddDocumentationDto, documentation_reference string) entity.Documentation {
	documentation := entity.Documentation{
		DocumentationReference: documentation_reference,
		DocumentationType:      dto.DocumentationType,
		DocumentationNumber:    dto.DocumentationNumber,
		Expiry:                  dto.Expiry,
		DocumentReference:       dto.DocumentReference,
		TierReference:           dto.TierReference,
	}
	return documentation
}
func AddContactDtoToContact(user entity.User, dto dto.ContactDto) entity.User {
	contact := entity.Contact{
		ContactReference: uuid.New().String(),
		Phone:            dto.Phone,
		FullName:         dto.FullName,
		WallsTag:         dto.WallsTag,
		IsBeneficiary:    dto.IsBeneficiary,
	}
	user.Contacts = append(user.Contacts, contact)

	return user
}

func UserDtoToUser(userDto dto.UserDto) entity.User {

	userDocumentations := []entity.Documentation{}
	for _, id := range userDto.Kyc.Documentations {
		_id := entity.Documentation{
			DocumentationReference: id.DocumentationReference,
			DocumentationType:      id.DocumentationType,
			DocumentationNumber:    id.DocumentationNumber,
			Expiry:                  id.Expiry,
			DocumentReference:       id.DocumentReference,
			IsVerified:              id.IsVerified,
			VerifiedOn:              id.VerifiedOn,
			VerificationMethod:      id.VerificationMethod,
		}
		userDocumentations = append(userDocumentations, _id)
	}

	result := entity.User{
		UserReference:  userDto.UserReference,
		CreatedOn:      userDto.CreatedOn,
		UpdatedOn:      userDto.UpdatedOn,
		IsActive:       userDto.IsActive,
		UserProfile:    userDto.UserProfile,
		CompanyProfile: userDto.CompanyProfile,
		Contacts:       userDto.Contacts,
		Security:       userDto.Security,
		Wallet: entity.Wallet{
			WalletReference:     userDto.Wallet.WalletReference,
			WallsAccountNo:      userDto.Wallet.WallsAccountNo,
			AutoFund:            userDto.Wallet.AutoFund,
			AutoFundLevel:       userDto.Wallet.AutoFundLevel,
			AutoFundLimit:       userDto.Wallet.AutoFundAmount,
			AutoWithdrawal:      userDto.Wallet.AutoWithdrawal,
			AutoWithdrawalLevel: userDto.Wallet.AutoFundAmount,
			AutoWithdrawalLimit: userDto.Wallet.AutoFundAmount,
			Tier: entity.Tier{
				TierReference:         userDto.Wallet.Tier.TierReference,
				TierName:              userDto.Wallet.Tier.TierName,
				SendingLimit:          userDto.Wallet.Tier.SendingLimit,
				DailyTransactionLimit: userDto.Wallet.Tier.DailyTransactionLimit,
				MinimumBalance:        userDto.Wallet.Tier.MinimumBalance,
				UpgradeOptions:        userDto.Wallet.Tier.UpgradeOptions,
			},
			Coupons: userDto.Wallet.Coupons,
			Reward:  userDto.Wallet.Reward,
			Balance: userDto.Wallet.Balance,
		},
		BankAccounts:        userDto.BankAccounts,
		Cards:               userDto.Cards,
		NotificationOptions: userDto.NotificationOptions,
		Device:              userDto.Device,
		Kyc: entity.Kyc{
			Documentations: userDocumentations,
		},
	}

	return result
}

func CreateCompanyProfileDtoToUser(user entity.User, dto dto.CreateCompanyProfileDto) entity.User {
	companyProfile := entity.CompanyProfile{
		CompanyProfileReference: uuid.New().String(),
		CompanyName:             dto.CompanyName,
		RegistrationNo:          dto.RegistrationNo,
		RegistrationDate:        dto.RegistrationDate,
		Phone:                   dto.Phone,
		Email:                   dto.Email,
		Address:                 dto.Address,
		Logo:                    dto.Logo,
		IsVerifiedEmail:         false,
	}

	user.CompanyProfile = append(user.CompanyProfile, companyProfile)

	return user
}

func UpdateCompanyProfileDtoToCompanyProfile(dto dto.UpdateCompanyProfileDto) entity.CompanyProfile {
	companyProfile := entity.CompanyProfile{
		Phone:   dto.Phone,
		Email:   dto.Email,
		Address: dto.Address,
	}

	return companyProfile
}
