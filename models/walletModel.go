package models

import (
	"time"
	"github.com/google/uuid"
)

// deposit and withdraw of crypto
type WalletNetwork struct {
    ID               uuid.UUID   `gorm:"type:uuid;primaryKey" json:"id"`
    NetworkType      *string     `gorm:"not null;uniqueIndex:networktype_userid;check:network_type IN ('EVM','SOLANA','APTOS')" json:"network_type"`
    WalletAddress    *string     `gorm:"unique;not null" json:"wallet_address"`
    EncPrivKey       *string     `gorm:"unique;not null" json:"enc_priv_key"`           // AES-encrypted private key (never raw)
	UserID           *uuid.UUID       `gorm:"type:uuid;not null;uniqueIndex:networktype_userid" json:"user_id"`
	User             User        `gorm:"foreignKey:UserID;" json:"-"`
	Wallet           []Wallet    `json:"wallet"`
	CreatedAt     time.Time      `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt     time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}

// deposit and withdraw of crypto
type Wallet struct {
    ID              uuid.UUID      `gorm:"type:uuid;primaryKey" json:"id"`
	// Network         *Network       `gorm:"uniqueIndex:network_asset_walletid;not null;check:network IN ('ETHEREUM','BASE','BSC','POLYGON','SOLANA','APTOS')" json:"network"`
	Balance         *int64         `gorm:"not null;" json:"balance"`
	Decimal         *uint          `gorm:"not null;" json:"decimal"`
	Asset           *Assets        `gorm:"uniqueIndex:network_asset_walletid;not null;check:asset IN ('ETH','APT','POL','BNB','SOL','USDC','USDT')" json:"asset"` // also add surpoorted assets 
	WalletNetworkId *uuid.UUID     `gorm:"type:uuid;uniqueIndex:network_asset_walletid;not null;" json:"wallet_network_id"`
    WalletNetwork   WalletNetwork  `gorm:"foreignKey:WalletNetworkId;" json:"-"`
	CreatedAt       time.Time      `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt       time.Time      `gorm:"autoUpdateTime" json:"updated_at"`
}


// deposit and withdraw of crypto 
type TransactionOnchain struct {
    ID               uuid.UUID        `gorm:"type:uuid;primaryKey" json:"id"`
    UserID           *uuid.UUID       `gorm:"type:uuid;not null;" json:"user_id"`
	User             User             `gorm:"foreignKey:UserID;" json:"-"`
    // WalletID       uint           // Which wallet (Solana / EVM / Sui / Aptos)
    Network          *Network         `gorm:"not null;check:network IN ('ETHEREUM','BASE','BSC','POLYGON','SOLANA','APTOS')" json:"network"`
    TxHash           *string          `gorm:"not null;unique" json:"tx_hash"`
    FromAddress      *string          `gorm:"not null;" json:"from_address"`
    ToAddress        *string          `gorm:"not null;" json:"to_address"`
    Amount           *float64         `gorm:"not null;" json:"amount"` 
    Asset            *Assets          `gorm:"not null;check:asset IN ('ETH','APT','POL','BNB','SOL','USDC','USDT')" json:"asset"` // also add surpoorted assets    // "ETH", "SOL", "APT", "USDC"
    TransactionType  *TransactionType `gorm:"not null;check:transaction_type IN ('DEPOSIT','WITHDRAWAL')" json:"transaction_type"`// "deposit", "withdrawal", "transfer", "swap"
    Status           *StatusType      `gorm:"not null;check:status IN ('PENDING','CONFIRMED','FAILED')" json:"status"`// "pending", "confirmed", "failed"
    CreatedAt        time.Time        `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt        time.Time        `gorm:"autoUpdateTime" json:"updated_at"`
}

// transfer of crypto internally
type CryptoInternalTransaction struct {
    ID                   uuid.UUID   `gorm:"type:uuid;primaryKey" json:"id"` 
	// one to many twice with same table  
	FromUserID           *uuid.UUID         `gorm:"type:uuid;not null" json:"from_user_id"`   
	FromUser             User          `gorm:"foreignKey:FromUserID" json:"-"`
	ToUserID             *uuid.UUID         `gorm:"type:uuid;not null" json:"to_user_id"`   
    ToUser               User          `gorm:"foreignKey:ToUserID" json:"-"`
    FromWalletaddress    *string       `gorm:"not null" json:"from_wallet_address"`  
    ToWalletaddress      *string       `gorm:"not null" json:"to_wallet_address"`  
    Asset                *Assets       `gorm:"not null;check:asset IN ('ETH','APT','POL','BNB','SOL','USDC','USDT')" json:"asset"` 
    Amount               *int64     `gorm:"not null;check:amount > 0" json:"amount"`
    // Type            string          `gorm:"not null;check:type IN ('TRANSFER')"`
    CreatedAt            time.Time     `gorm:"autoCreateTime" json:"created_at"`
    UpdatedAt            time.Time     `gorm:"autoUpdateTime" json:"updated_at"`
}

