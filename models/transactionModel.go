package models

import (
	"time"

	"github.com/google/uuid"
)

type FiatCryptoTransaction struct {
	ID              uuid.UUID                  `gorm:"type:uuid;primaryKey" json:"id"`
	UserID          *uuid.UUID                 `gorm:"type:uuid;not null;"`
	User            User                       `gorm:"foreignKey:UserID;" json:"-"`
	TransactionType *FiatCryptoTransactionType `gorm:"not null;check:transaction_type IN ('FIAT_TO_CRYPTO','CRYPTO_TO_FIAT')" json:"transaction_type"`
	// Crypto side
	CryptoAsset  *Assets `gorm:"not null;check:crypto_asset IN ('ETH','APT','POL','BNB','SOL','USDC','USDT')" json:"crypto_asset"`
	CryptoAmount *string  `gorm:"type:numeric;not null;check:crypto_amount > 0" json:"crypto_amount"` // string to  avoid float precision issues
	// Fiat side
	FiatCurrency *Currency `gorm:"not null;check:fiat_currency IN ('NGN','USD')" json:"fiat_currency"`
	FiatAmount   *string    `gorm:"type:numeric;not null;check:fiat_amount> 0" json:"fiat_amount"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
}