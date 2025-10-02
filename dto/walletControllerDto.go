package dto

// import "github.com/shopspring/decimal"

// import (
// 	"github.com/go-playground/validator/v10"
// 	"github.com/google/uuid"
// )

type CreateWalletDTO struct {
    UserID    string    `json:"user_id" validate:"required"`
	NetworkType  string  `json:"networktype" validate:"required"`

}

type CreateWalletApiResponseDTOData struct {
    WalletID    string    `json:"walletId" `
	WalletAddress  string  `json:"walletAddress"`
    Network  string  `json:"network"`
    Mnemonic string    `json:"mnemonic"`

}

type CreateWalletApiResponseDTO struct {
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    CreateWalletApiResponseDTOData `json:"data"`
}

type TransfertokenApiResponseDTOData struct {
    Txhash string   `json:"hash"`

}
type TransfertokenApiResponseDTO struct {
    Success bool        `json:"success"`
    Message string      `json:"message"`
    Data    TransfertokenApiResponseDTOData `json:"data"`
}


type DepositTokenDTO struct {
    Network        string      `json:"network" validate:"required"`
    Asset          string      `json:"asset" validate:"required"`
    UserId         string      `json:"user_id" validate:"required"`
    DepositAmount  string  `json:"deposit_amount" validate:"required,gt=0" ` 
    Fromaddress   string `json:"from_address" validate:"required"`
    Txhash   string `json:"tx_hash" validate:"required"`
    // Status   string `json:"status" validate:"required"`
    TransactionType string `json:"transaction_type" validate:"required"`
}

type WithdrawTokenDTO struct {
    Network        string      `json:"network" validate:"required"`
    Asset          string      `json:"asset" validate:"required"`
    UserId         string      `json:"user_id" validate:"required"`
    WithdrawAmount  string  `json:"withdraw_amount" validate:"required" ` 
    Toaddress   string `json:"to_address" validate:"required"`
    // Status   string `json:"status" validate:"required"`
    TransactionType string `json:"transaction_type" validate:"required"`
}

// type LoginRequestDTO struct {
// 	BaseEmailDTO
//     Password string `json:"password" validate:"required,min=8"`
// }

// type VerifyOtpDTO struct {
// 	BaseEmailDTO
//     Otp string `json:"otp" validate:"required,min=6"`
// }

// type SignupRequestDTO struct {
//     LoginRequestDTO
//     Phone    string `json:"phone" validate:"required"`            // required, can add regex if needed
//     Age      uint   `json:"age" validate:"required,gte=15"`      // required + age >= 18
//     Role     string `json:"role,omitempty"`
// 	Verified bool   `json:"verified,omitempty"`
// 	VerificationType string `json:"verification_type,omitempty"`
// 	// encrypted
// 	VerificationId    string   `json:"verification_id,omitempty"`
// }

// type RefreshTokenDTO struct {
//     RefreshToken    string `json:"refresh_token" validate:"required"`       // required + valid email
// }

