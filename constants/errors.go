package constants


// Server error constants
const (
	InternalServerError = "InternalServerError"
	InvalidRequestData  = "InvalidRequestData"
	Unauthorized        = "Unauthorized"
	NotFound            = "NotFound"
	Conflict            = "Conflict"
	BadRequest          = "Bad Request"

	EmailNotVerified           = "EmailNotVerified"
	EmailNotFound              = "EmailNotFound"
	EmailNotAttached           = "EmailNotAttached"
	ExpiredToken               = "TokenExpired"
	InvalidUserCreateRequest   = "InvalidUserCreateRequest"
	InvalidUserLoginRequest    = "InvalidUserLoginRequest"
	ITEM_ALREADY = "ITEM_ALREADY"
	WrongPassword              = "WrongPassword"
	InvalidResetPasswordToken  = "InvalidResetPasswordToken"
	InvalidUserInviteToken     = "InvalidUserInviteToken"
	NoInviteExists             = "NoInviteExists"
	InvalidForgotPasswordToken = "InvalidForgotPasswordToken"
	InvalidServiceAccountToken = "InvalidServiceAccountToken"
	InvalidPaginationInput     = "InvalidPaginationInput"
	OTPExpired                 = "OTPExpired"
	InvalidImageURL            = "InvalidImageURL"
	InvalidObjectName          = "InvalidObjectName"
	InvalidOTP                 = "InvalidOTP"
	InvalidEmailVerifyToken    = "InvalidEmailVerifyToken"
	InvalidJWTToken            = "InvalidJWTToken"
	MismatchFBInfo             = "MismatchFBInfo"
	UserDoesNotBelongToTenant  = "UserDoesNotBelongToTenant"
	SignupEmailFail            = "SignupEmailFail"
	MinPasswordLengthError     = "MinPasswordLengthError"
	TeamDoesNotExist           = "TeamDoesNotExit"
	InvalidFile                = "InvalidFile"
	TenantNotSupplied          = "TenantNotSupplied"
	EmailNotPresent            = "EmailNotPresent"
	NameInvalid                = "NameInvalid"
	NoUserExists               = "NoUserExists"
	NoUserSupplied             = "NoUserSupplied"
	UserAlreadyPartOfTeam      = "UserAlreadyPartOfTeam"
	FailedToSendSMS            = "FailedToSendSMS"
	FailedToCreateTenant       = "FailedToCreateTenant"
	UserAlreadyPartOfTheTenant = "UserAlreadyPartOfTheTenant"
	UserAlreadyPartOfConduct   = "UserAlreadyPartOfConduct"
)

// ErrorString returns the string version of the error which is sent to the user
var ErrorString = map[string]string{
	InternalServerError:        "We're sorry! Looks like something went wrong",
	InvalidRequestData:         "The request failed because it contained an invalid value",
	Unauthorized:               "We're sorry! You are not authorized to perform this action",
	NotFound:                   "This request could not be found",
	EmailNotFound:              "This email address is not registered",
	EmailNotAttached:           "Email id not attached with your facebook account",
	Conflict:                   "An item already exists with this name",
	InvalidUserCreateRequest:   "Please enter valid information",
	InvalidUserLoginRequest:    "Your username and/or password do not match",
	WrongPassword:              "You've entered a wrong password. Please try again or reset your password",
	InvalidResetPasswordToken:  "Invalid reset token",
	InvalidUserInviteToken:     "Invalid invite token",
	NoInviteExists:             "No such user has been invited",
	InvalidForgotPasswordToken: "Invalid forgot password token",
	InvalidServiceAccountToken: "Invalid service account token",
	InvalidPaginationInput:     "Invalid page value",
	OTPExpired:                 "The OTP has expired. Generate a new one",
	InvalidImageURL:            "The Image URL is not valid",
	InvalidObjectName:          "Invalid object name",
	InvalidOTP:                 "You've entered a wrong OTP. Please try again",
	InvalidEmailVerifyToken:    "Invalid verification token",
	InvalidJWTToken:            "Invalid authentication token",
	MismatchFBInfo:             "Mismatch in facebook information",
	UserDoesNotBelongToTenant:  "The user does not belong to the organization specified",
	EmailNotVerified:           "Your email is not verified. Please check your email for instructions",
	SignupEmailFail:            "Failed to send signup email",
	//UserEmailAlreadyRegistered: "This email is already taken. Please enter another email id",
	// MinPasswordLengthError:     fmt.Sprintf("Your password should have %v characters minimum", MinPasswor),
	TeamDoesNotExist:           "Team doesn't exist",
	InvalidFile:                "Invalid file uploaded",
	TenantNotSupplied:          "Tenant is mandatory",
	EmailNotPresent:            "Email is not present in one or more rows",
	NameInvalid:                "Name cannot be empty",
	NoUserExists:               "We could not find a user registered with this email address",
	NoUserSupplied:             "No user supplied.",
	UserAlreadyPartOfTeam:      "The user is already a part of the team",
	FailedToSendSMS:            "SMS sending failed.",
	FailedToCreateTenant:       "Tenant creation failed.",
	UserAlreadyPartOfTheTenant: "The user is already a part of the tenant",
	UserAlreadyPartOfConduct:   "The user is already part of conduct.",
}
