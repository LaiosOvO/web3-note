package main

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

// LogTransfer ..
type LogTransfer struct {
	From   common.Address
	To     common.Address
	Tokens *big.Int
}

// LogApproval ..
type LogApproval struct {
	TokenOwner common.Address
	Spender    common.Address
	Tokens     *big.Int
}

//func main() {
//	// 连接网络
//	client, err := ethclient.Dial("https://cloudflare-eth.com")
//	if err != nil {
//		klog.Fatal(err)
//	}
//
//	// 指定合约
//	// 0x Protocol (ZRX) token address
//	contractAddress := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
//	query := ethereum.FilterQuery{
//		FromBlock: big.NewInt(6383820),
//		ToBlock:   big.NewInt(6383840),
//		Addresses: []common.Address{
//			contractAddress,
//		},
//	}
//	// 查询过滤参数
//	logs, err := client.FilterLogs(context.Background(), query)
//	if err != nil {
//		klog.Fatal(err)
//	}
//
//	contractAbi, err := abi.JSON(strings.NewReader(string(token.TokenABI)))
//	if err != nil {
//		klog.Fatal(err)
//	}
//
//	logTransferSig := []byte("Transfer(address,address,uint256)")
//	LogApprovalSig := []byte("Approval(address,address,uint256)")
//	logTransferSigHash := crypto.Keccak256Hash(logTransferSig)
//	logApprovalSigHash := crypto.Keccak256Hash(LogApprovalSig)
//
//	for _, vLog := range logs {
//		fmt.Printf("Log Block Number: %d\n", vLog.BlockNumber)
//		fmt.Printf("Log Index: %d\n", vLog.Index)
//
//		switch vLog.Topics[0].Hex() {
//		case logTransferSigHash.Hex():
//			fmt.Printf("Log Name: Transfer\n")
//
//			var transferEvent LogTransfer
//
//			a, err := contractAbi.Unpack("Transfer", vLog.Data)
//			if err != nil {
//				klog.Fatal(err)
//			}
//
//			transferEvent.From = common.HexToAddress(vLog.Topics[1].Hex())
//			transferEvent.To = common.HexToAddress(vLog.Topics[2].Hex())
//			fmt.Println(a[0])
//			//fmt.Println(a[1])
//			//transferEvent.Tokens = a[2].(*big.Int)
//			fmt.Printf("From: %s\n", transferEvent.From.Hex())
//			fmt.Printf("To: %s\n", transferEvent.To.Hex())
//			fmt.Printf("Tokens: %s\n", transferEvent.Tokens.String())
//
//		case logApprovalSigHash.Hex():
//			fmt.Printf("Log Name: Approval\n")
//
//			var approvalEvent LogApproval
//
//			_, err := contractAbi.Unpack("Approval", vLog.Data)
//			if err != nil {
//				klog.Fatal(err)
//			}
//
//			approvalEvent.TokenOwner = common.HexToAddress(vLog.Topics[1].Hex())
//			approvalEvent.Spender = common.HexToAddress(vLog.Topics[2].Hex())
//
//			fmt.Printf("Token Owner: %s\n", approvalEvent.TokenOwner.Hex())
//			fmt.Printf("Spender: %s\n", approvalEvent.Spender.Hex())
//			fmt.Printf("Tokens: %s\n", approvalEvent.Tokens.String())
//		}
//
//		fmt.Printf("\n\n")
//	}
//}
