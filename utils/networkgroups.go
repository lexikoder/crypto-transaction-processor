package utils

import (
	"crypto-transaction-processor/models"
	"fmt"

	"github.com/shopspring/decimal"
)

var EVMNetworks = map[models.Network]bool{

	models.NetworkSEPOLIA:true,
	models.NetworkBASESEPOLIA:true,
	models.NetworkEth:true,
	models.NetworkBase:true,
	models.NetworkBsc:true,
	models.NetworkPolygon:true,
	
}
func IsEVM(network models.Network) (bool) {
	return EVMNetworks[network]
}

var SOLANANetworks = map[models.Network]bool{
	models.NetworkSolana:true,	
}
func IsSOLANA(network models.Network) (bool) {
	return SOLANANetworks[network]
}

var APTOSNetworks = map[models.Network]bool{
	models.NetworkAPTOS:true,	
}
func IsAPTOS(network models.Network) (bool) {
	return APTOSNetworks[network]
}

func CheckNetworkType(network models.Network) (models.NetworkType) {
	if(IsEVM(network)){
		return  models.NetworkEvmType
	}
	if(IsSOLANA(network)){
		return  models.NetworkSolanaType
	} 

	if(IsAPTOS(network)){
		return  models.NetworkAptosType
	}

    return models.NetworkUnknownType
}


var SepoliaAssetAddress = map[models.Assets]string{
	"ETH":"ETH",
	"USDC":"0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238",
	"USDT":"0xaA8E23Fb1079EA71e0a56F48a2aA51851D8433D0",
}

var BaseSepoliaAssetAddress = map[models.Assets]string{
	"ETH":"ETH",
	"USDC":" 0x036CbD53842c5426634e7929541eC2318f3dCF7e",
	"USDT":"",
}

func NetworkAssetAddress(network models.Network,asset models.Assets) (string) {
	if (network ==  models.NetworkSEPOLIA ){
       return  SepoliaAssetAddress[asset]
	}

	if (network ==  models.NetworkBASESEPOLIA ){
       return  BaseSepoliaAssetAddress[asset]
	}

	return ""
}

var SupportedAssetSepoliaNetwork = map[models.Assets]bool{
	models.AssetEth:true,
	models.AssetUsdc:true,
	models.AssetUsdt:true,
}
var SupportedAssetBaseSepoliaNetwork = map[models.Assets]bool{
	models.AssetEth:true,
	models.AssetUsdc:true,
	models.AssetUsdt:true,
}

func SupportedAssetNetwork(network models.Network,asset models.Assets) bool {
	if (network ==  models.NetworkSEPOLIA ){
       return  SupportedAssetSepoliaNetwork[asset]
	}

	if (network ==  models.NetworkBASESEPOLIA ){
       return  SupportedAssetBaseSepoliaNetwork[asset]
	}

	return false
}

var NativeAsset = map[models.Assets]bool{
	models.AssetEth:true, 
	models.AssetApt:true, 
	models.AssetPol:true,
	models.AssetBnb:true,
}

func Isnative(asset models.Assets) bool {
  return NativeAsset[asset]
}

var AssetDecimal = map[models.Assets]int32{
	models.AssetEth:18, 
	models.AssetApt:8, 
	models.AssetPol:18,
	models.AssetSol:9,
	models.AssetBnb:18,
	models.AssetUsdc:6,
	models.AssetUsdt:6,
}

func Assettodecimal(network models.Network,asset models.Assets) (int32,bool) {
if (network ==  models.NetworkBsc ){
       AssetDecimal[models.AssetUsdc]=18
	    AssetDecimal[models.AssetUsdt]=18
	}

	if val, ok := AssetDecimal[asset]; ok {
     return val,ok
} else {
    return 0,ok
}
 
}


func WeiToTokenAmount(amount string,network models.Network,asset models.Assets)(string,bool){

	decimals ,ok := Assettodecimal(network,asset)
numerator, err := decimal.NewFromString(amount)
	if err != nil {
		panic(err)
	}

	// Create denominator 10^18
	denominator := decimal.New(1, decimals) // 1 * 10^18

	// Divide
	result := numerator.Div(denominator)

	return result.String() ,ok
}

func SafeInt32ToUint(x int32) (uint, error) {
    if x < 0 {
        return 0, fmt.Errorf("cannot convert negative int32 (%d) to uint", x)
    }
    return uint(x), nil
}