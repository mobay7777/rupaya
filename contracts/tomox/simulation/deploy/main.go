package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/rupayaproject/go-rupaya/accounts/abi/bind"
	"github.com/rupayaproject/go-rupaya/common"
	"github.com/rupayaproject/go-rupaya/contracts/tomox"
	"github.com/rupayaproject/go-rupaya/contracts/tomox/simulation"
	"github.com/rupayaproject/go-rupaya/ethclient"
)

func main() {
	fmt.Println("========================")
	fmt.Println("mainAddr", simulation.MainAddr.Hex())
	fmt.Println("relayerAddr", simulation.RelayerCoinbaseAddr.Hex())
	fmt.Println("ownerRelayerAddr", simulation.OwnerRelayerAddr.Hex())
	fmt.Println("========================")
	client, err := ethclient.Dial(simulation.RpcEndpoint)
	if err != nil {
		fmt.Println(err, client)
	}
	nonce, _ := client.NonceAt(context.Background(), simulation.MainAddr, nil)
	fmt.Println(nonce)
	auth := bind.NewKeyedTransactor(simulation.MainKey)
	auth.Value = big.NewInt(0)      // in wei
	auth.GasLimit = uint64(4000000) // in units
	auth.GasPrice = big.NewInt(210000000000000)

	// init rrc21 issuer
	auth.Nonce = big.NewInt(int64(nonce))
	rrc21IssuerAddr, rrc21Issuer, err := tomox.DeployRRC21Issuer(auth, client, simulation.MinRRC21Apply)
	if err != nil {
		log.Fatal("DeployRRC21Issuer", err)
	}
	rrc21Issuer.TransactOpts.GasPrice = big.NewInt(210000000000000)

	fmt.Println("===> rrc21 issuer address", rrc21IssuerAddr.Hex())
	fmt.Println("wait 10s to execute init smart contract : TRC Issuer")
	time.Sleep(2 * time.Second)

	//init TOMOX Listing in
	auth.Nonce = big.NewInt(int64(nonce + 1))
	tomoxListtingAddr, tomoxListing, err := tomox.DeployTOMOXListing(auth, client)
	if err != nil {
		log.Fatal("DeployTOMOXListing", err)
	}
	tomoxListing.TransactOpts.GasPrice = big.NewInt(210000000000000)

	fmt.Println("===> tomox listing address", tomoxListtingAddr.Hex())
	fmt.Println("wait 10s to execute init smart contract : tomox listing ")
	time.Sleep(2 * time.Second)

	// init Relayer Registration
	auth.Nonce = big.NewInt(int64(nonce + 2))
	relayerRegistrationAddr, relayerRegistration, err := tomox.DeployRelayerRegistration(auth, client, tomoxListtingAddr, simulation.MaxRelayers, simulation.MaxTokenList, simulation.MinDeposit)
	if err != nil {
		log.Fatal("DeployRelayerRegistration", err)
	}
	relayerRegistration.TransactOpts.GasPrice = big.NewInt(210000000000000)

	fmt.Println("===> relayer registration address", relayerRegistrationAddr.Hex())
	fmt.Println("wait 10s to execute init smart contract : relayer registration ")
	time.Sleep(2 * time.Second)

	currentNonce := nonce + 3
	tokenList := initRRC21(auth, client, currentNonce, simulation.TokenNameList)

	currentNonce = currentNonce + uint64(len(simulation.TokenNameList)) // init smartcontract

	applyIssuer(rrc21Issuer, tokenList, currentNonce)

	currentNonce = currentNonce + uint64(len(simulation.TokenNameList))
	applyTomoXListing(tomoxListing, tokenList, currentNonce)

	currentNonce = currentNonce + uint64(len(simulation.TokenNameList))
	airdrop(auth, client, tokenList, simulation.TeamAddresses, currentNonce)

	// relayer registration
	ownerRelayer := bind.NewKeyedTransactor(simulation.OwnerRelayerKey)
	nonce, _ = client.NonceAt(context.Background(), simulation.OwnerRelayerAddr, nil)
	relayerRegistration, err = tomox.NewRelayerRegistration(ownerRelayer, relayerRegistrationAddr, client)
	if err != nil {
		log.Fatal("NewRelayerRegistration", err)
	}
	relayerRegistration.TransactOpts.Nonce = big.NewInt(int64(nonce))
	relayerRegistration.TransactOpts.Value = simulation.MinDeposit
	relayerRegistration.TransactOpts.GasPrice = big.NewInt(210000000000000)

	fromTokens := []common.Address{}
	toTokens := []common.Address{}

	/*
		for _, token := range tokenList {
			fromTokens = append(fromTokens, token["address"].(common.Address))
			toTokens = append(toTokens, simulation.TOMONative)
		}
	*/

	// TOMO/BTC
	fromTokens = append(fromTokens, simulation.TOMONative)
	toTokens = append(toTokens, tokenList[0]["address"].(common.Address))

	// TOMO/USD
	fromTokens = append(fromTokens, simulation.TOMONative)
	toTokens = append(toTokens, tokenList[9]["address"].(common.Address))

	// ETH/TOMO
	fromTokens = append(fromTokens, tokenList[1]["address"].(common.Address))
	toTokens = append(toTokens, simulation.TOMONative)

	fromTokens = append(fromTokens, tokenList[2]["address"].(common.Address))
	toTokens = append(toTokens, simulation.TOMONative)

	fromTokens = append(fromTokens, tokenList[3]["address"].(common.Address))
	toTokens = append(toTokens, simulation.TOMONative)

	fromTokens = append(fromTokens, tokenList[4]["address"].(common.Address))
	toTokens = append(toTokens, simulation.TOMONative)

	fromTokens = append(fromTokens, tokenList[5]["address"].(common.Address))
	toTokens = append(toTokens, simulation.TOMONative)

	fromTokens = append(fromTokens, tokenList[6]["address"].(common.Address))
	toTokens = append(toTokens, simulation.TOMONative)

	fromTokens = append(fromTokens, tokenList[7]["address"].(common.Address))
	toTokens = append(toTokens, simulation.TOMONative)

	fromTokens = append(fromTokens, tokenList[8]["address"].(common.Address))
	toTokens = append(toTokens, simulation.TOMONative)

	// ETH/BTC
	fromTokens = append(fromTokens, tokenList[1]["address"].(common.Address))
	toTokens = append(toTokens, tokenList[0]["address"].(common.Address))

	// XRP/BTC
	fromTokens = append(fromTokens, tokenList[2]["address"].(common.Address))
	toTokens = append(toTokens, tokenList[0]["address"].(common.Address))

	// BTC/USD
	fromTokens = append(fromTokens, tokenList[0]["address"].(common.Address))
	toTokens = append(toTokens, tokenList[9]["address"].(common.Address))

	// ETH/USD
	fromTokens = append(fromTokens, tokenList[1]["address"].(common.Address))
	toTokens = append(toTokens, tokenList[9]["address"].(common.Address))

	_, err = relayerRegistration.Register(simulation.RelayerCoinbaseAddr, simulation.TradeFee, fromTokens, toTokens)
	if err != nil {
		log.Fatal("relayerRegistration Register ", err)
	}
	fmt.Println("wait 10s to apply token to list tomox")
	time.Sleep(2 * time.Second)
}

func initRRC21(auth *bind.TransactOpts, client *ethclient.Client, nonce uint64, tokenNameList []string) []map[string]interface{} {
	tokenListResult := []map[string]interface{}{}
	for _, tokenName := range tokenNameList {
		auth.Nonce = big.NewInt(int64(nonce))
		d := uint8(18)
		if tokenName == "USD" {
			d = 8
		}
		tokenAddr, _, err := tomox.DeployRRC21(auth, client, tokenName, tokenName, d, simulation.RRC21TokenCap, simulation.RRC21TokenFee)
		if err != nil {
			log.Fatal("DeployRRC21 ", tokenName, err)
		}

		fmt.Println(tokenName+" token address", tokenAddr.Hex(), "cap", simulation.RRC21TokenCap)
		fmt.Println("wait 10s to execute init smart contract", tokenName)

		tokenListResult = append(tokenListResult, map[string]interface{}{
			"name":    tokenName,
			"address": tokenAddr,
		})
		nonce = nonce + 1
	}
	time.Sleep(5 * time.Second)
	return tokenListResult
}

func applyIssuer(rrc21Issuer *tomox.RRC21Issuer, tokenList []map[string]interface{}, nonce uint64) {
	for _, token := range tokenList {
		rrc21Issuer.TransactOpts.Nonce = big.NewInt(int64(nonce))
		rrc21Issuer.TransactOpts.Value = simulation.MinRRC21Apply
		_, err := rrc21Issuer.Apply(token["address"].(common.Address))
		if err != nil {
			log.Fatal("rrc21Issuer Apply  ", token["name"].(string), err)
		}
		fmt.Println("wait 10s to applyIssuer ", token["name"].(string))
		nonce = nonce + 1
	}
	time.Sleep(5 * time.Second)
}

func applyTomoXListing(tomoxListing *tomox.TOMOXListing, tokenList []map[string]interface{}, nonce uint64) {
	for _, token := range tokenList {
		tomoxListing.TransactOpts.Nonce = big.NewInt(int64(nonce))
		tomoxListing.TransactOpts.Value = simulation.TomoXListingFee
		_, err := tomoxListing.Apply(token["address"].(common.Address))
		if err != nil {
			log.Fatal("tomoxListing Apply ", token["name"].(string), err)
		}
		fmt.Println("wait 10s to applyTomoXListing ", token["name"].(string))
		nonce = nonce + 1
	}
	time.Sleep(5 * time.Second)
}

func airdrop(auth *bind.TransactOpts, client *ethclient.Client, tokenList []map[string]interface{}, addresses []common.Address, nonce uint64) {
	for _, token := range tokenList {
		for _, address := range addresses {
			rrc21Contract, _ := tomox.NewRRC21(auth, token["address"].(common.Address), client)
			rrc21Contract.TransactOpts.Nonce = big.NewInt(int64(nonce))
			_, err := rrc21Contract.Transfer(address, big.NewInt(0).Mul(common.BasePrice, big.NewInt(1000000)))
			if err == nil {
				fmt.Printf("Transfer %v to %v successfully", token["name"].(string), address.String())
				fmt.Println()
			} else {
				fmt.Printf("Transfer %v to %v failed!", token["name"].(string), address.String())
				fmt.Println()
			}
			nonce = nonce + 1
		}
	}
	time.Sleep(5 * time.Second)
}
