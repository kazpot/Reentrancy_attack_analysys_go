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

	charityAddress := common.HexToAddress("0x082537af866061fb35dab18e18c8ececaea92db7")

	// 150 GWei
	auth.Value = big.NewInt(150e+9)

	walletAddress1, tx1, _, err := contracts.DeployWallet(auth, blockchain, charityAddress)
	if tx1 != nil {
		log.Printf("walletAddress1: 0x%x\ntx: %s", walletAddress1, tx1.Hash().String())
	}
}
