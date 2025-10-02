package service

import (
	"context"
	"crypto-transaction-processor/database"
	"crypto-transaction-processor/dto"
	"crypto-transaction-processor/models"
	"crypto-transaction-processor/utils"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
	"log"
	"net/http"
	"os"
	// "time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type WalletService struct {
	Db *database.Repository
}

func (walletService *WalletService) CreateWalletService(ctx context.Context, dto dto.CreateWalletDTO, apiresponsedto dto.CreateWalletApiResponseDTO) error {
	WALLETAPIKEY := os.Getenv("WALLETAPIKEY")
	var wallet models.WalletNetwork

	err := walletService.Db.DB.WithContext(ctx).Where("user_id = ? AND network_type = ? ", dto.UserID, dto.NetworkType).First(&wallet).Error
	log.Println("error:", err)
	if err == nil {
		return utils.NewAppError("you already created a wallet", http.StatusBadRequest)
	}

	data := map[string]string{
		"networktype": dto.NetworkType,
	}

	body, err := utils.ApiPostRequest("https://wallet-infracstructure-backend-dkvj.onrender.com/api/wallet/createwallet", data, WALLETAPIKEY)
	if err != nil {
		fmt.Println("Captured error1:", err)
		return utils.NewAppError("something went wrong", http.StatusInternalServerError)
	}

	responsedto := apiresponsedto
	if err := json.Unmarshal(body, &responsedto); err != nil {
		fmt.Println("Captured error2:", err)
		return utils.NewAppError("something went wrong", http.StatusInternalServerError)
	}
	fmt.Println("value:", responsedto)
	if !responsedto.Success {
		return utils.NewAppError("could not create wallet", http.StatusInternalServerError)
	}

	userUUID, err := uuid.Parse(dto.UserID)
	if err != nil {
		fmt.Println("invalid uuid:", err)
		return utils.NewAppError("something went wrong", http.StatusInternalServerError)
	}
	key := []byte(os.Getenv("ENCRYPTION_KEY"))
	plaintext := []byte(responsedto.Data.Mnemonic)
	encyptedkey, err := utils.Encrypt(plaintext, key)
	if err != nil {
		fmt.Println("Captured error3:", err)
		return utils.NewAppError("something went wrong", http.StatusInternalServerError)
	}

	wallet = models.WalletNetwork{
		ID:               uuid.New(),
		NetworkType:      &dto.NetworkType,
		WalletAddress:    &responsedto.Data.WalletAddress,
		EncPrivKey:       &encyptedkey,
		ExternalWalletId: &responsedto.Data.WalletID,
		UserID:           &userUUID,
	}

	fmt.Println("result:", responsedto.Success, responsedto.Data, responsedto.Message)
	err = walletService.Db.DB.WithContext(ctx).Create(&wallet).Error
	fmt.Println("Captured error4:", err)
	if err != nil {
		return utils.NewAppError("something went wrong", http.StatusInternalServerError)
	}

	return nil
}

func (walletService *WalletService) DepositTokenService(ctx context.Context, dto dto.DepositTokenDTO) error {

	// userid asset  network depositamount decimal
	var walletNetwork models.WalletNetwork
	var wallet models.Wallet
	var TransactionOnchain models.TransactionOnchain

	amount, greaterthanzeroerr := decimal.NewFromString(dto.DepositAmount)
	if greaterthanzeroerr != nil {
		return utils.NewAppError("invalid deposit amount", http.StatusBadRequest)
	}

	if !amount.GreaterThan(decimal.Zero) {
		return utils.NewAppError("deposit amount must be greater than 0", http.StatusBadRequest)
	}

	checknetworktype := utils.CheckNetworkType(models.Network(dto.Network))
	if checknetworktype == models.NetworkUnknownType {
		return utils.NewAppError("network not supported", http.StatusBadRequest)
	}

	SupportedAssetNetwork := utils.SupportedAssetNetwork(models.Network(dto.Network), models.Assets(dto.Asset))
	if !SupportedAssetNetwork {
		return utils.NewAppError("asset not supported on this network", http.StatusBadRequest)
	}

	err := walletService.Db.DB.WithContext(ctx).Where("user_id = ? AND network_type = ? ", dto.UserId, checknetworktype).First(&walletNetwork).Error
	log.Println("error:", err)
	if err != nil {
		return utils.NewAppError("no wallet created yet", http.StatusBadRequest)
	}

	// this allows rollback if on transaction doesnt go through
	transactionerr := walletService.Db.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {

		walleterr := walletService.Db.DB.WithContext(ctx).Where("wallet_network_id = ? AND asset = ? ", walletNetwork.ID, dto.Asset).First(&wallet).Error
		fmt.Println("walleterr:", walleterr)
		if errors.Is(walleterr, gorm.ErrRecordNotFound) {
			decimal, ok := utils.Assettodecimal(models.Network(dto.Network), models.Assets(dto.Asset))
			if !ok {
				return utils.NewAppError("asset not supported on this network", http.StatusBadRequest)
			}
			var balance string = "0"
			var available_balance string = "0"
			wallet = models.Wallet{
				ID:               uuid.New(),
				AvailableBalance: &available_balance,
				Balance:          &balance,
				Decimal:          &decimal,
				Asset:            (*models.Assets)(&dto.Asset),
				WalletNetworkId:  &walletNetwork.ID,
			}
			fmt.Println("check1:", err)
			err = tx.Create(&wallet).Error
			if err != nil {
				return errors.New("error creating wallet")
			}
		}
		fmt.Println("check2:", err)
		if err != nil {
			return utils.NewAppError("database error", http.StatusBadRequest)
		}

		fmt.Println("check3:", err, wallet.ID)
		// when this has been confirmed on the blockchain we wait for 13 blocks confirmation then we update available_balance on a background worker
		result := tx.Model(&models.Wallet{}).Where("id = ?", wallet.ID).Update("balance", gorm.Expr("balance + ?", dto.DepositAmount))
		if result.Error != nil {
			fmt.Println("checkhere2:", err)
			return result.Error
		}

		if result.RowsAffected == 0 {
			fmt.Println("checkhere:", err)
			return errors.New("could not update balance")
		}
		fmt.Println("check4:", result.Error)
		userUUID, err := uuid.Parse(dto.UserId)
		if err != nil {
			fmt.Println("invalid uuid:", err)
			return err
		}
		if models.TransactionType(dto.TransactionType) != models.TransactionDeposit {
			fmt.Println(models.TransactionType(dto.TransactionType), models.TransactionDeposit)
			return errors.New("invalid transaction type")
		}

		tt := models.TransactionType(dto.TransactionType)
		network := models.Network(dto.Network)
		fmt.Println(network, "tt")
		status := models.StatusPending
		amountStr := dto.DepositAmount

		TransactionOnchain = models.TransactionOnchain{
			ID:              uuid.New(),
			UserID:          &userUUID,
			Network:         &network,
			TxHash:          &dto.Txhash,
			FromAddress:     &dto.Fromaddress,
			ToAddress:       walletNetwork.WalletAddress,
			Amount:          &amountStr,
			Asset:           (*models.Assets)(&dto.Asset),
			TransactionType: &tt,
			Status:          &status,
		}

		if err := tx.Create(&TransactionOnchain).Error; err != nil {
			return err // rollback happens automatically
		}
		fmt.Println("check5", err)
		return nil // commit happens automatically
	})

	if transactionerr != nil {
		fmt.Println("check3:", transactionerr)
		return utils.NewAppError("could not deposit", http.StatusInternalServerError)
	}

	return nil

}

func (walletService *WalletService) WithdrawTokenService(ctx context.Context, dto dto.WithdrawTokenDTO,apiresponsedto dto.TransfertokenApiResponseDTO) error {
    MNENOMIC := os.Getenv("MNENOMIC")
	WALLETADDRESS :=os.Getenv("WALLETADDRESS")
	// userid asset  network depositamount decimal
	var walletNetwork models.WalletNetwork
	var wallet models.Wallet
	var TransactionOnchain models.TransactionOnchain

	amount, greaterthanzeroerr := decimal.NewFromString(dto.WithdrawAmount)
	if greaterthanzeroerr != nil {
		return utils.NewAppError("invalid deposit amount", http.StatusBadRequest)
	}

	if !amount.GreaterThan(decimal.Zero) {
		return utils.NewAppError("deposit amount must be greater than 0", http.StatusBadRequest)
	}

	checknetworktype := utils.CheckNetworkType(models.Network(dto.Network))
	if checknetworktype == models.NetworkUnknownType {
		return utils.NewAppError("network not supported", http.StatusBadRequest)
	}

	SupportedAssetNetwork := utils.SupportedAssetNetwork(models.Network(dto.Network), models.Assets(dto.Asset))
	if !SupportedAssetNetwork {
		return utils.NewAppError("asset not supported on this network", http.StatusBadRequest)
	}

	if models.TransactionType(dto.TransactionType) != models.TransactionWithdrawal {
		fmt.Println(models.TransactionType(dto.TransactionType), models.TransactionDeposit)
		return errors.New("invalid transaction type")
	}
	tt := models.TransactionType(dto.TransactionType)

	err := walletService.Db.DB.WithContext(ctx).Where("user_id = ? AND network_type = ? ", dto.UserId, checknetworktype).First(&walletNetwork).Error
	log.Println("error:", err)
	if err != nil {
		return utils.NewAppError("no wallet created yet", http.StatusBadRequest)
	}
	network := models.Network(dto.Network)

	walleterr := walletService.Db.DB.WithContext(ctx).Where("wallet_network_id = ? AND asset = ? ", walletNetwork.ID, dto.Asset).First(&wallet).Error
	fmt.Println("walleterr:", walleterr)
	if errors.Is(walleterr, gorm.ErrRecordNotFound) {
		return utils.NewAppError("you have no wallet or balace to complete this transaction", http.StatusBadRequest)
	}
	fmt.Println("check2:", err)
	if err != nil {
		return utils.NewAppError("database error", http.StatusBadRequest)
	}

	fmt.Println("check3:", err, wallet.ID)

	result := walletService.Db.DB.WithContext(ctx).Model(&models.Wallet{}).Where("id = ? AND available_balance >= ?", wallet.ID, dto.WithdrawAmount).Update("available_balance", gorm.Expr("available_balance - ?", dto.WithdrawAmount))

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("insufficient balance")
	}

	userUUID, err := uuid.Parse(dto.UserId)
	if err != nil {
		fmt.Println("invalid uuid:", err)
		return err
	}

	// fromaddress := "rrrrrrrr"
	// Txhash := "rrrrrrrr"
	transactionerr := walletService.Db.DB.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		// network := models.Network(dto.Network)
		fmt.Println(network, "tt")
		status := models.StatusPending
		amountStr := dto.WithdrawAmount
		TransactionOnchain = models.TransactionOnchain{
			ID:      uuid.New(),
			UserID:  &userUUID,
			Network: &network,
			// TxHash:          &Txhash,
			FromAddress:     &WALLETADDRESS,
			ToAddress:       &dto.Toaddress,
			Amount:          &amountStr,
			Asset:           (*models.Assets)(&dto.Asset),
			TransactionType: &tt,
			Status:          &status,
		}

		if err := tx.Create(&TransactionOnchain).Error; err != nil {
			return err // rollback happens automatically
		}
		return nil // commit happens automatically
	})

	if transactionerr != nil {
		fmt.Println("check3:", transactionerr)
		return utils.NewAppError("could not deposit", http.StatusInternalServerError)
	}
    
    data := map[string]string{
 "mnemonic":MNENOMIC,
  "token":dto.Asset,
  "fromaddress":WALLETADDRESS,
  "toaddress":dto.Toaddress,
  "amount":dto.WithdrawAmount,
  "networktype":string(checknetworktype),
  "_network":dto.Network,
	}

	body, err := utils.ApiPostRequest("https://wallet-infracstructure-backend-dkvj.onrender.com/api/wallet/createwallet", data, "")
	// if err != nil {
	// 	fmt.Println("Captured error1:", err)
	// 	return utils.NewAppError("something went wrong", http.StatusInternalServerError)
	// }
	// txHash, err := blockchain.SendTransaction(wallet, amount, toAddress)
	if err != nil {
		//     // try to mark withdrawal as failed
		status := models.StatusFailed
		// stat := ""
		updateErr := walletService.Db.DB.WithContext(ctx).Model(&models.TransactionOnchain{}).
			Where("id = ?", TransactionOnchain.ID).
			Updates(models.TransactionOnchain{
				Status: &status,
				//  TxHash: &stat,
			}).Error

			
		if updateErr != nil {
			// log this separately, don’t lose it
			log.Printf("failed to mark withdrawal %s as failed: %v", TransactionOnchain.ID, updateErr)

			// optionally push into retry queue / outbox
		}

		result := walletService.Db.DB.WithContext(ctx).Model(&models.Wallet{}).Where("id = ?", wallet.ID).Update("available_balance", gorm.Expr("available_balance + ?", dto.WithdrawAmount))
		if result.Error != nil {
			log.Printf("failed to mark withdrawal %s as failed: %v", TransactionOnchain.ID, updateErr)

		}

		if result.RowsAffected == 0 {
			log.Printf("failed to mark withdrawal %s as failed: %v", TransactionOnchain.ID, updateErr)

		}

		return utils.NewAppError("you have no wallet or balace to complete this transaction", http.StatusBadRequest)

	}

    
	responsedto := apiresponsedto
	if err := json.Unmarshal(body, &responsedto); err != nil {
		fmt.Println("Captured error2:", err)
		return utils.NewAppError("something went wrong", http.StatusInternalServerError)
	}
	fmt.Println("value:", responsedto)
	if !responsedto.Success {
		return utils.NewAppError("could not create wallet", http.StatusInternalServerError)
	}

	// this should run as a background worker to complete the status to confirmed when 3 blocks has been confirmed
	status := models.StatusConfirmed
	updateErr := walletService.Db.DB.WithContext(ctx).Model(&models.TransactionOnchain{}).
		Where("id = ?", TransactionOnchain.ID).
		Updates(models.TransactionOnchain{
				Status: &status,
			    TxHash: &responsedto.Data.Txhash}).Error

	if updateErr != nil {
		log.Printf("failed to mark withdrawal %s as failed: %v", TransactionOnchain.ID, updateErr)
	}

	return nil

}


func (authController *WalletService) TransactionHistoryService() gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
