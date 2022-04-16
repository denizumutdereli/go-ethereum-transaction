package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var (
	mainurl = "https://mainnet.infura.io/v3/...."
	url     = "https://kovan.infura.io/v3/....."
)

func main() {

	/*
	 */

	client, err := ethclient.Dial(url)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()

	a1 := common.HexToAddress("abababababababababababababababababababab")
	a2 := common.HexToAddress("cdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcdcd")

	// fmt.Println(a1)
	// fmt.Println(a2)

	b1, err := client.BalanceAt(context.Background(), a1, nil)
	if err != nil {
		log.Fatal(err)
	}

	b2, err := client.BalanceAt(context.Background(), a2, nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Balance one:", b1)
	fmt.Println("Balance two:", b2)

	value := big.NewInt(100000000000000)
	gasLimit := uint64(21000) // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	nonce, err := client.PendingNonceAt(context.Background(), a1)

	if err != nil {
		log.Fatal(err)
	}

	var data []byte

	tx := types.NewTransaction(nonce, a2, value, gasLimit, gasPrice, data)

	chainID, err := client.NetworkID(context.Background())

	fmt.Println(chainID)

	if err != nil {
		log.Fatal(err)
	}

	b, err := ioutil.ReadFile("./wallet/UTC--2022-04-12T22-16-02.825598600Z--llwwecsma...")

	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(b, "password")

	if err != nil {
		log.Fatal(err)
	}
	tx, err = types.SignTx(tx, types.NewEIP155Signer(chainID), key.PrivateKey)

	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), tx)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx has sent to: %s", tx.Hash().Hex())
}
