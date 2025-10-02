package dto
// import "github.com/go-playground/validator/v10"

type BaseEmailDTO struct {
    Email    string `json:"email" validate:"required,email"`       // required + valid email
}

type LoginRequestDTO struct {
	BaseEmailDTO
    Password string `json:"password" validate:"required,min=8"`
}

type VerifyOtpDTO struct {
	BaseEmailDTO
    Otp string `json:"otp" validate:"required,min=6"`
}

type SignupRequestDTO struct {
    LoginRequestDTO
    Phone    string `json:"phone" validate:"required"`            // required, can add regex if needed
    Age      uint   `json:"age" validate:"required,gte=15"`      // required + age >= 18
    Role     string `json:"role,omitempty"`
	Verified bool   `json:"verified,omitempty"`
	VerificationType string `json:"verification_type,omitempty"`
	// encrypted
	VerificationId    string   `json:"verification_id,omitempty"`
}

type RefreshTokenDTO struct {
    RefreshToken    string `json:"refresh_token" validate:"required"`       // required + valid email
}

