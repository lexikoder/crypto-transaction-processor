package models

type Role string 
const ( 
	RoleUser Role = "USER" 
	RoleAdmin Role = "ADMIN" )


type VerificationType string 
const ( 
	VerificationNin VerificationType = "NIN" 
	VerificationDriversLicense VerificationType = "DRIVERS_LICENSE"
	VerificationIntlPassport VerificationType = "INTL_PASSPORT"
)


type NetworkType string 
const ( 
	NetworkEvmType NetworkType = "EVM" 
	NetworkSolanaType NetworkType = "SOLANA"
	NetworkAptosType NetworkType = "APTOS"   
	NetworkUnknownType NetworkType = "UNKNOWN"
)


type Network string 
const ( 
	NetworkSEPOLIA Network =    "SEPOLIA"
	NetworkBASESEPOLIA Network= "BASESEPOLIA"
	NetworkEth Network =        "ETHEREUM" 
	NetworkBase Network =       "BASE"
	NetworkBsc Network =        "BSC"
	NetworkPolygon Network =    "POLYGON"
	NetworkSolana Network = "SOLANA"
	NetworkAPTOS Network =   "APTOS"   
)

type Assets string 
const (
	AssetEth Assets =  "ETH"
	AssetApt Assets =  "APT"
	AssetPol Assets =  "POL"
	AssetBnb Assets =  "BNB"
	AssetSol Assets =  "SOL" 
	AssetUsdc Assets = "USDC" 
	AssetUsdt Assets = "USDT" 
)


// Transaction table for crypto 
type TransactionType string 
const (
	TransactionDeposit TransactionType =     "DEPOSIT"
	TransactionWithdrawal TransactionType =  "WITHDRAWAL"
	// TransactionTransfer TransactionType =    "TRANSFER"
	// TransactionSwap TransactionType =        "SWAP"
)


type StatusType string 
const (
	StatusPending StatusType =    "PENDING"
	StatusConfirmed StatusType =  "CONFIRMED"
	StatusFailed StatusType =     "FAILED"
)


type FiatCryptoTransactionType string
const (
    FiatToCrypto FiatCryptoTransactionType = "FIAT_TO_CRYPTO"
    CryptoToFiat FiatCryptoTransactionType = "CRYPTO_TO_FIAT"
)


type Currency string 
const (
	CurrencyNgn Currency =    "NGN"
	CurrencyUsd Currency =    "USD"
)

type FiatTransactionType string 
const (
	// DEPOSIT not supported yet
	FiatTransactionTypeDEPOSIT FiatTransactionType =    "DEPOSIT"
	FiatTransactionTypeWITHDRAWAL FiatTransactionType =    "WITHDRAWAL"	
)


type FiatInternalTransactionType string 
const (

	FiatTransactionTypeTRANSFER FiatInternalTransactionType =    "TRANSFER"
	
)





// always store 
// naira -> kobo
// eth wei 
// dollar Current

// library to use shopspring/decimal

// model
// balance int64
// amount int64

// remember to avoid storing - balance  use atomicity 

// // User deposits ₦10,000
// depositNaira := int64(10000 * 100) // 10000 Naira → 1,000,000 kobo

// // Convert to decimal for safe calculations
// balanceDec := decimal.NewFromInt(wallet.Balance).Div(decimal.NewFromInt(100)) // back to Naira
// depositDec := decimal.NewFromInt(depositNaira).Div(decimal.NewFromInt(100))
// newBalance := balanceDec.Add(depositDec)
// return fixed to 2 decimal point to user 
// fmt.Println("New balance:", newBalance.StringFixed(2))



// // for crypto
// type Wallet struct {
//     Balance int64 // smallest unit, e.g., wei for ETH
// }

// // Suppose user deposits 0.75 ETH
// func main() {
//     wallet := Wallet{
//         Balance: 1500000000000000000, // 1.5 ETH in wei
//     }

//     // 0.75 ETH deposit in wei
//     depositETH := int64(750000000000000000) // 0.75 ETH * 10^18 wei

//     // Convert to decimal for calculations
//     balanceDec := decimal.NewFromInt(wallet.Balance).Div(decimal.NewFromInt(1e18)) // back to ETH
//     depositDec := decimal.NewFromInt(depositETH).Div(decimal.NewFromInt(1e18))

//     newBalance := balanceDec.Add(depositDec)
//     fmt.Println("New balance in ETH:", newBalance.StringFixed(18))
// }



// fiat to crypto 
// func main() {
//     // User deposits ₦10,000 (stored in kobo)
//     fiatDeposit := int64(10000 * 100) // 10,000 Naira → 1,000,000 kobo

//     // Current rate: 1 USDT = ₦650
//     rate := decimal.NewFromFloat(650) // Naira per USDT

//     // Convert fiat to crypto (USDT)
//     fiatDec := decimal.NewFromInt(fiatDeposit)          // 1,000,000 kobo
//     fiatInNaira := fiatDec.Div(decimal.NewFromInt(100)) // back to Naira
//     cryptoAmount := fiatInNaira.Div(rate)              // amount in USDT

//     fmt.Println("USDT to credit:", cryptoAmount.StringFixed(6))
// }