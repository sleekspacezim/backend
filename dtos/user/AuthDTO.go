package dtos

type NewPasswordDTO struct {
	Password string `json:"password" validate:"required"`
	UserId   int    `json:"userId" validate:"required"`
}

type NativeUserRegistrationDTO struct {
	GivenName  string `json:"givenName" validate:"required,min=2,max=50"`
	FamilyName string `json:"familyName" validate:"required,min=2,max=50"`
	Email      string `json:"email" validate:"email,required"`
	Password   string `json:"password" validate:"required"`
}

type LoginDTO struct {
	Email    string `json:"email" validate:"email,required"`
	Password string `json:"password" validate:"required"`
}

type VerificationDTO struct {
	UserId           int `json:"userId" validate:"required"`
	VerificationCode int `json:"verificationCode" validate:"required,min=100000,max=999999"`
}

type CreateVerificationCodeDTO struct {
	Email string `json:"email" validate:"email,required"`
}
