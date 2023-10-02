package helper

import (
	"regexp"

	"github.com/go-playground/validator/v10"
)

func ValidateValidContact(fl validator.FieldLevel) bool {
	contact := fl.Field().String()

	// Regular expression patterns for email and phone number
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	phonePattern := `^\+\d{1,3}\d{4,}$`

	// Check if the contact matches the email pattern
	isEmail, _ := regexp.MatchString(emailPattern, contact)

	// Check if the contact matches the phone number pattern
	isPhone, _ := regexp.MatchString(phonePattern, contact)

	return isEmail || isPhone
}

func ValidateValidEmail(fl validator.FieldLevel) bool {
	email := fl.Field().String()
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	match, _ := regexp.MatchString(emailPattern, email)
	return match

}

func ValidateGUID(fl validator.FieldLevel) bool {
	guid := fl.Field().String()

	// Define the regular expression pattern for a GUID-like string
	// Adjust the pattern according to the specific format you expect
	pattern := `^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$`

	// Match the GUID string against the regular expression pattern
	match, _ := regexp.MatchString(pattern, guid)

	return match
}

func ValidateIMEI(fl validator.FieldLevel) bool {
	imei := fl.Field().String()

	// Define the regular expression pattern for an IMEI number
	// Adjust the pattern according to the specific format you expect
	pattern := `^\d{15}$`

	// Match the IMEI number against the regular expression pattern
	match, _ := regexp.MatchString(pattern, imei)

	return match
}

func ValidateValidPhone(fl validator.FieldLevel) bool {
	contact := fl.Field().String()

	// Regular expression pattern for phone number
	phonePattern := `^\+\d{1,3}\d{4,}$`

	// Check if the contact matches the phone number pattern
	isPhone, _ := regexp.MatchString(phonePattern, contact)

	return isPhone
}

func IsValidPhone(contact string) bool {
	isPhone, _ := regexp.MatchString(`^\+\d{1,3}\d{4,}$`, contact)
	return isPhone
}

func IsValidEmail(contact string) bool {
	isPhone, _ := regexp.MatchString(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`, contact)
	return isPhone
}
