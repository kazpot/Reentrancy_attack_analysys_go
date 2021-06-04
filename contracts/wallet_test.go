package contracts

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/abi/bind/backends"
	"github.com/ethereum/go-ethereum/core"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
)

func TestDeployWallet(t *testing.T) {
	key, _ := crypto.GenerateKey()
	auth, _ := bind.NewKeyedTransactorWithChainID(key, big.NewInt(1337))
	alloc := make(core.GenesisAlloc)
	alloc[auth.From] = core.GenesisAccount{Balance: big.NewInt(1000000000)}
	blockchain := backends.NewSimulatedBackend(alloc, 5000000)

	charityAddress, _, _, err := DeployCharity(auth, blockchain)
	if err != nil {
		t.Fatalf("Failed to deploy the inbox contract: %v", err)
	}
	blockchain.Commit()

	if len(charityAddress.Bytes()) == 0 {
		t.Error("Expected a valid deployment address. Received empty.")
	}

	walletAddress, _, wallet, err := DeployWallet(auth, blockchain, charityAddress)
	if err != nil {
		t.Fatalf("Failed to deploy the inbox contract: %v", err)
	}
	blockchain.Commit()

	if len(walletAddress.Bytes()) == 0 {
		t.Error("Expected a valid deployment address. Received empty.")
	}

	gasPrice, err := blockchain.SuggestGasPrice(nil)
	if err != nil {
		t.Error("Failed to get SuggestGasPrice")
	}

	_, err = wallet.Donate(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		GasPrice: gasPrice,
		GasLimit: 420000,
		Value:  nil,
	}, big.NewInt(2e+18))
	if err != nil {
		t.Fatalf("Failed to call donate: %v", err)
	}

	_, err = wallet.Withdraw(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		GasPrice: gasPrice,
		GasLimit: 420000,
		Value:  nil,
	})
	if err != nil {
		t.Fatalf("Failed to call withdraw: %v", err)
	}

	_, err = wallet.GetBalance(nil)
	if err != nil {
		t.Fatalf("Failed to get blance from contract: %v", err)
	}

	_, err = wallet.GetDonationBalance(nil)
	if err != nil {
		t.Fatalf("Failed to get donation blance from contract: %v", err)
	}


}