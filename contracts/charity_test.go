package contracts

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)

func TestDeployCharity(t *testing.T) {
	key, _ := crypto.GenerateKey()

	// geth private chainID = 1337
	auth, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	alloc := make(core.GenesisAlloc)

	// 1000000000 ETH
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, 5000000)

	address, _, contract, err := DeployCharity(auth, blockchain)
	if err != nil {
		t.Fatalf("Failed to deploy the inbox contract: %v", err)
	}
	blockchain.Commit()

	if len(address.Bytes()) == 0 {
		t.Error("Expected a valid address")
	}

	_, err = contract.Donate(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		Value:  big.NewInt(2e+18),
	})

	_, err = contract.GetBalance(nil)
	if err != nil {
		t.Fatalf("Failed to get blance from contract: %v", err)
	}
}