# Re-entrancy attack example in Go

This is an example of re-entrancy attack on single function which is written in Go.

## Install (MacOS Big Sur 11.4)
```
$ brew tap ethereum/ethereum
$ brew install ethereum
```

## Install geth
```
$ go get -d github.com/ethereum/go-ethereum
$ cd $GOPATH/src/github.com/ethereum
$ go install ./...
```

## Version
- Solidity - 0.5.0 (solc-js)
- go version go1.16.4 darwin/amd64

## Dependencies
```
$ go mod download
```

## Test

```
$ cd contracts
$ go test -v
```

## Deploy to Rinkeby Testnet

```
// modify config.ini
[rinkeby]
key = {json account key}
password = {account password}
infura = {infura key}
network_id = 4 // rinkeby network id
contract_address = {contract address}
log_file = {log file path}

// deploy command
$ cd deploy

// deploy Charity contract
$ go run deploy1.go

// deploy Wallet1 contract using Charity contract address
$ go run deploy2.go

// deploy Wallet2 contract using Charity contract address
$ go run deploy2.go
```

## Execute re-entrancy

```
// wallet1 and wallet2 both donate 1 Gwei to Charity Contract
$ go run donate.go

// Run wallet2 withdraw 1 time, but it runs 2 times due to re-entrancy
// So wallet2 maliciously can get 1 GWei * 2 = 2 GWei
$ go run exploit.go
```

## Prevention

ref: [https://quantstamp.com/blog/what-is-a-re-entrancy-attack](https://quantstamp.com/blog/what-is-a-re-entrancy-attack)

- Use send() or transfer() instead of call.value()
- Update the user balance before the ether transfer

```
contract Charity {
    mapping(address => uint256) public donations;

    function withdraw() public {
        uint256 amount = donations[msg.sender];
        donations[msg.sender] = 0;
        msg.sender.transfer(share);
    }
}
```