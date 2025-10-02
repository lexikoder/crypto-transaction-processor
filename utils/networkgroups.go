package utils

import "crypto-transaction-processor/models"

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






 
// var NativeToNetwork = map[models.Network]models.Assets{
// 	models.NetworkSEPOLIA:models.AssetEth,
// 	models.NetworkBASESEPOLIA:models.AssetEth,
// 	models.NetworkEth:models.AssetEth,
// 	models.NetworkBase:models.AssetEth,
// 	models.NetworkBsc:models.AssetBnb,
// 	models.NetworkPolygon:models.AssetPol,
// 	models.NetworkSolana:models.AssetSol,
// 	models.NetworkMove:models.AssetApt,
// }

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

var AssetDecimal = map[models.Assets]uint{
	models.AssetEth:18, 
	models.AssetApt:8, 
	models.AssetPol:18,
	models.AssetSol:9,
	models.AssetBnb:18,
	models.AssetUsdc:6,
	models.AssetUsdt:6,
}

func Assettodecimal(network models.Network,asset models.Assets) (uint,bool) {
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
