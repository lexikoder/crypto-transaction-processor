package models

import (
	"time"

	"github.com/google/uuid"
)

type FiatWallet struct {
    ID              uuid.UUID   `gorm:"type:uuid;primaryKey" json:"id"`
    UserID          *uuid.UUID         `gorm:"type:uuid;uniqueIndex:userid_currency;not null" json:"user_id"`
	User            User          `gorm:"foreignKey:UserID;" json:"-"`
    Currency        *Currency     `gorm:"uniqueIndex:userid_currency;not null;check:currency IN ('NGN','USD')" json:"currency"`
    Balance         *int64       `gorm:"not null" json:"balance"`
    CreatedAt       time.Time     `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt       time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
}


// this is for withdrawal and deposit 
type FiatTransaction struct {
    ID               uuid.UUID   `gorm:"type:uuid;primaryKey" json:"id"`
    UserID            *uuid.UUID                  `gorm:"type:uuid;not null;" json:"user_id"`
	User              User                   `gorm:"foreignKey:UserID;" json:"-"`
    TransactionType  *FiatTransactionType    `gorm:"not null;check:transaction_type IN ('WITHDRAWAL','DEPOSIT')" json:"transaction_type"`
    BankName          *string                 `gorm:"not null;" json:"bank_name" `
    AccountNo         *string                 `gorm:"not null;" json:"account_no" `
    Amount            *int64          `gorm:"not null;" json:"amount" `
	FiatCurrency          *Currency       `gorm:"not null;check:fiat_currency IN ('NGN','USD')" json:"fiat_currency"`
    // Status          string    // pending, confirmed, failed
    CreatedAt         time.Time     `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt         time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
}

// internal transfer to other users 
type FiatInternalTransaction struct {
    ID              uuid.UUID   `gorm:"type:uuid;primaryKey" json:"id"`
    TransactionType *FiatInternalTransactionType    `gorm:"not null;check:transaction_type IN ('TRANSFER')" json:"transaction_type"`
    FromUserID      *uuid.UUID         `gorm:"type:uuid;not null" json:"from_user_id"`   
	FromUser             User          `gorm:"foreignKey:FromUserID" json:"-"`
	ToUserID             *uuid.UUID         `gorm:"type:uuid;not null" json:"to_user_id"`   
    ToUser               User          `gorm:"foreignKey:ToUserID" json:"-"`
    Amount          *int64            `gorm:"not null;" json:"amount" `
	FiatCurrency        *Currency       `gorm:"not null;check:fiat_currency IN ('NGN','USD')" json:"fiat_currency"`
    CreatedAt       time.Time     `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt       time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
}