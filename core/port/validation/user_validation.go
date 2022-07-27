package validation

type PhoneNumberValidatorInterface interface {
	Validate(phoneNumber string, countryCode string) error
}
