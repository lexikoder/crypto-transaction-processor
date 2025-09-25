package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID               uuid.UUID `gorm:"type:uuid;primaryKey" json:"id"`

	Email            *string           `gorm:"not null;unique" json:"email"`
	Phone            *string           `gorm:"not null;unique" json:"phone"`
	Age              *uint             `json:"age"`
	Password         *string           `gorm:"not null;" json:"password"`
	Role             Role              `gorm:"not null;check:role IN ('USER','ADMIN');default:USER" json:"role"`
	Verified         bool              `gorm:"default:false" json:"verified"`
	VerificationType *VerificationType `gorm:"check:verification_type IN ('NIN','DRIVERS_LICENSE','INTL_PASSPORT')" json:"verification_type"`
	// encrypted
	VerificationId                     *string                     `gorm:"unique;" json:"verification_id"`
	WalletNetwork                      []WalletNetwork             `json:"wallet_network"`
	TransactionOnchain                 []TransactionOnchain        `json:"transaction_onchain"`
	CryptoInternalSenderTransactions   []CryptoInternalTransaction `gorm:"foreignKey:FromUserID"`
	CryptoInternalReceiverTransactions []CryptoInternalTransaction `gorm:"foreignKey:ToUserID"`
	FiatInternalSenderTransactions     []FiatInternalTransaction   `gorm:"foreignKey:FromUserID"`
	FiatInternalRecieverTransactions   []FiatInternalTransaction   `gorm:"foreignKey:FromUserID"`
	RefreshToken                       []RefreshToken              `gorm:"foreignKey:UserID"`
	CreatedAt                          time.Time                   `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt                          time.Time                   `gorm:"autoUpdateTime" json:"updated_at"`
}