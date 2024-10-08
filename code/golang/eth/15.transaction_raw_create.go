package main

//func main() {
//	client, err := ethclient.Dial("https://rinkeby.infura.io")
//	if err != nil {
//		klog.Fatal(err)
//	}
//
//	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
//	if err != nil {
//		klog.Fatal(err)
//	}
//
//	publicKey := privateKey.Public()
//	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
//	if !ok {
//		klog.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
//	}
//
//	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
//	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
//	if err != nil {
//		klog.Fatal(err)
//	}
//
//	value := big.NewInt(1000000000000000000) // in wei (1 eth)
//	gasLimit := uint64(21000)                // in units
//	gasPrice, err := client.SuggestGasPrice(context.Background())
//	if err != nil {
//		klog.Fatal(err)
//	}
//
//	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
//	var data []byte
//	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
//
//	chainID, err := client.NetworkID(context.Background())
//	if err != nil {
//		klog.Fatal(err)
//	}
//
//	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
//	if err != nil {
//		klog.Fatal(err)
//	}
//
//	ts := types.Transactions{signedTx}
//	rawTxBytes := ts.GetRlp(0)
//	rawTxHex := hex.EncodeToString(rawTxBytes)
//
//	fmt.Printf(rawTxHex) // f86...772
//}
