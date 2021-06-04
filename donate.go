package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/kazpot/Reentrancy_attack_analysys_go/config"
	"github.com/kazpot/Reentrancy_attack_analysys_go/contracts"
	"github.com/kazpot/Reentrancy_attack_analysys_go/utils"
	"log"
	"math/big"
	"strings"
)

func main() {
	utils.LoggingSettings(config.Config.LogFile)

	blockchain, err := ethclient.Dial(config.Config.Infura)
	if err != nil {
		log.Fatalf("Unable to connect to network:%v\n", err)
	}

	auth, err := bind.NewTransactorWithChainID(
		strings.NewReader(config.Config.Key),
		config.Config.Password,
		big.NewInt(config.Config.NetworkId))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	walletAddress1 := common.HexToAddress("0x58b64bb4de3bf3a2022252c9769df5dd9f290be0")
	walletAddress2 := common.HexToAddress("0x358642ae777fc38a1a6fd43e36eb43762465ce07")

	wallet1, err := contracts.NewWallet(walletAddress1, blockchain)
	if err != nil {
		log.Fatalf("Failed to create wallet1 instance: %v", err)
	}

	wallet2, err := contracts.NewWallet(walletAddress2, blockchain)
	if err != nil {
		log.Fatalf("Failed to create wallet2 instance: %v", err)
	}

	tx1, err := wallet1.Donate(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasPrice: big.NewInt(100e+9),
		GasLimit: 420000,
		Value:    nil,
	}, big.NewInt(1e+9))
	if err != nil {
		log.Fatalf("Failed to call donate1: %v", err)
	}
	if tx1 != nil {
		log.Printf("Donated from wallet1, tx1: %s", tx1.Hash().String())
	}

	tx2, err := wallet2.Donate(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasPrice: big.NewInt(100e+9),
		GasLimit: 420000,
		Value:    nil,
	}, big.NewInt(1e+9))
	if err != nil {
		log.Fatalf("Failed to call donate2: %v", err)
	}
	if tx2 != nil {
		log.Printf("Donated from wallet2, tx2: %s", tx2.Hash().String())
	}
}
