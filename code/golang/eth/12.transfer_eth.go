package main

//func main() {
//	// 创建连接
//	client, err := ethclient.Dial("https://rinkeby.infura.io")
//	if err != nil {
//		log.Fatal("连接失败")
//		log.Fatal(err)
//	}
//	// 加载私钥
//	privateKey, err := crypto.HexToECDSA("fad9c8855b740a0b7ed4c221dbad0f33a83a49cad6b3fe8d5817ac83d38b6a19")
//	if err != nil {
//		log.Fatal("加载私钥失败")
//		log.Fatal(err)
//	}
//	// 获取公钥
//	publicKey := privateKey.Public()
//	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
//	if !ok {
//		log.Fatal("获取公钥失败")
//		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
//	}
//	// 得到地址
//	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
//	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
//	if err != nil {
//		log.Println("公钥解析地址失败")
//		log.Fatal(err)
//	}
//	// 传输的价格
//	value := big.NewInt(1_00000_00000_0000_0000) // in wei (1 eth)
//	gasLimit := uint64(21000)                    // in units
//	gasPrice, err := client.SuggestGasPrice(context.Background())
//	if err != nil {
//		log.Fatal("获取平均价格失败")
//		log.Fatal(err)
//	}
//
//	toAddress := common.HexToAddress("0x4592d8f8d7b001e72cb26a73e4fa1806a51ac79d")
//	var data []byte
//	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
//
//	chainID, err := client.NetworkID(context.Background())
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	err = client.SendTransaction(context.Background(), signedTx)
//	if err != nil {
//		log.Fatal(err)
//	}
//
//	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
//}
