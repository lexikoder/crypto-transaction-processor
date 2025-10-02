package models

import "gorm.io/gorm"

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(
		&User{},
		&WalletNetwork{},
		&Wallet{},
		&TransactionOnchain{},
		&CryptoInternalTransaction{},
		&FiatCryptoTransaction{},
		&FiatWallet{},
		&FiatTransaction{},
		&FiatInternalTransaction{},
		&RefreshToken{},
	)
	return err
}