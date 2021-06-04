package main

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
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

	// 100 GWei
	auth.GasPrice = big.NewInt(100e+9)
	auth.GasLimit = 420000

	charityAddress, tx, _, err := contracts.DeployCharity(auth, blockchain)
	if tx != nil {
		log.Printf("charityAddress: 0x%x\ntx: %s", charityAddress, tx.Hash().String())
	}
}
