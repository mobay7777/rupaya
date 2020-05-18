package main

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"time"

	"github.com/rupayaproject/rupaya/accounts/abi/bind"
	"github.com/rupayaproject/rupaya/common"
	"github.com/rupayaproject/rupaya/contracts/rupx"
	"github.com/rupayaproject/rupaya/contracts/rupx/simulation"
	"github.com/rupayaproject/rupaya/ethclient"
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
	auth := bind.NewKeyedTransactor(simulation.MainKey)
	auth.Value = big.NewInt(0)       // in wei
	auth.GasLimit = uint64(10000000) // in units
	auth.GasPrice = big.NewInt(250000000000000)

	// init rrc21 issuer
	auth.Nonce = big.NewInt(int64(nonce))
	rrc21IssuerAddr, rrc21Issuer, err := rupx.DeployRRC21Issuer(auth, client, simulation.MinRRC21Apply)
	if err != nil {
		log.Fatal("DeployRRC21Issuer", err)
	}
	rrc21Issuer.TransactOpts.GasPrice = big.NewInt(250000000000000)

	fmt.Println("===> rrc21 issuer address", rrc21IssuerAddr.Hex())
	fmt.Println("wait 10s to execute init smart contract : TRC Issuer")
	time.Sleep(2 * time.Second)

	//init RUPX Listing in
	auth.Nonce = big.NewInt(int64(nonce + 1))
	rupxListtingAddr, rupxListing, err := rupx.DeployRUPXListing(auth, client)
	if err != nil {
		log.Fatal("DeployRUPXListing", err)
	}
	rupxListing.TransactOpts.GasPrice = big.NewInt(250000000000000)

	fmt.Println("===> rupx listing address", rupxListtingAddr.Hex())
	fmt.Println("wait 10s to execute init smart contract : rupx listing ")
	time.Sleep(2 * time.Second)

	// init Relayer Registration
	auth.Nonce = big.NewInt(int64(nonce + 2))
	relayerRegistrationAddr, relayerRegistration, err := rupx.DeployRelayerRegistration(auth, client, rupxListtingAddr, simulation.MaxRelayers, simulation.MaxTokenList, simulation.MinDeposit)
	if err != nil {
		log.Fatal("DeployRelayerRegistration", err)
	}
	relayerRegistration.TransactOpts.GasPrice = big.NewInt(250000000000000)

	fmt.Println("===> relayer registration address", relayerRegistrationAddr.Hex())
	fmt.Println("wait 2s to execute init smart contract : relayer registration ")
	time.Sleep(2 * time.Second)

	auth.Nonce = big.NewInt(int64(nonce + 3))
	lendingRelayerRegistrationAddr, lendingRelayerRegistration, err := rupx.DeployLendingRelayerRegistration(auth, client, relayerRegistrationAddr, rupxListtingAddr)
	if err != nil {
		log.Fatal("DeployLendingRelayerRegistration", err)
	}
	lendingRelayerRegistration.TransactOpts.GasPrice = big.NewInt(250000000000000)

	fmt.Println("===> lending relayer registration address", lendingRelayerRegistrationAddr.Hex())
	fmt.Println("wait 2s to execute init smart contract : lending relayer registration ")
	time.Sleep(2 * time.Second)

	currentNonce := nonce + 4
	tokenList := initRRC21(auth, client, currentNonce, simulation.TokenNameList)

	currentNonce = currentNonce + uint64(len(simulation.TokenNameList)) // init smartcontract

	applyIssuer(rrc21Issuer, tokenList, currentNonce)

	currentNonce = currentNonce + uint64(len(simulation.TokenNameList))
	applyRupXListing(rupxListing, tokenList, currentNonce)

	// BTC Collateral
	nonce = currentNonce + uint64(len(simulation.TokenNameList))
	lendingRelayerRegistration.TransactOpts.Nonce = big.NewInt(int64(nonce))
	_, err = lendingRelayerRegistration.AddCollateral(tokenList[0]["address"].(common.Address), simulation.CollateralDepositRate, simulation.CollateralLiquidationRate, simulation.CollateralRecallRate)
	if err != nil {
		log.Fatal("Lending add collateral", err)
	}

	// ETH Collateral
	nonce = nonce + 1
	lendingRelayerRegistration.TransactOpts.Nonce = big.NewInt(int64(nonce))
	_, err = lendingRelayerRegistration.AddCollateral(tokenList[1]["address"].(common.Address), simulation.CollateralDepositRate, simulation.CollateralLiquidationRate, simulation.CollateralRecallRate)
	if err != nil {
		log.Fatal("Lending add collateral", err)
	}

	// RUPX Collateral
	nonce = nonce + 1
	lendingRelayerRegistration.TransactOpts.Nonce = big.NewInt(int64(nonce))
	_, err = lendingRelayerRegistration.AddCollateral(simulation.RUPXNative, simulation.CollateralDepositRate, simulation.CollateralLiquidationRate, simulation.CollateralRecallRate)

	if err != nil {
		log.Fatal("Lending add collateral", err)
	}

	// XRP ILO Collateral
	nonce = nonce + 1
	lendingRelayerRegistration.TransactOpts.Nonce = big.NewInt(int64(nonce))
	_, err = lendingRelayerRegistration.AddILOCollateral(tokenList[2]["address"].(common.Address), simulation.CollateralDepositRate, simulation.CollateralLiquidationRate, simulation.CollateralRecallRate)
	if err != nil {
		log.Fatal("Lending add ILO collateral", err)
	}

	// USD lending base
	nonce = nonce + 1
	lendingRelayerRegistration.TransactOpts.Nonce = big.NewInt(int64(nonce))
	_, err = lendingRelayerRegistration.AddBaseToken(tokenList[9]["address"].(common.Address))
	if err != nil {
		log.Fatal("Lending add base token USD", err)
	}

	// RUPX lending base
	nonce = nonce + 1
	lendingRelayerRegistration.TransactOpts.Nonce = big.NewInt(int64(nonce))
	_, err = lendingRelayerRegistration.AddBaseToken(simulation.RUPXNative)
	if err != nil {
		log.Fatal("Lending add base token RUPX", err)
	}

	// BTC lending base
	nonce = nonce + 1
	lendingRelayerRegistration.TransactOpts.Nonce = big.NewInt(int64(nonce))
	_, err = lendingRelayerRegistration.AddBaseToken(tokenList[0]["address"].(common.Address))
	if err != nil {
		log.Fatal("Lending add base token BTC", err)
	}

	// add term 1 minute for testing
	for i := 0; i < len(simulation.Terms); i++ {
		nonce = nonce + 1
		lendingRelayerRegistration.TransactOpts.Nonce = big.NewInt(int64(nonce))
		_, err = lendingRelayerRegistration.AddTerm(simulation.Terms[i])
		if err != nil {
			log.Fatal("Lending add terms", err)
		}
		fmt.Println("wait 2s to add term lending lending contract", simulation.Terms[i])
		time.Sleep(2 * time.Second)
	}
	fmt.Println("wait 2s to setup lending contract")
	time.Sleep(2 * time.Second)

	currentNonce = nonce + 1
	airdrop(auth, client, tokenList, simulation.TeamAddresses, currentNonce)

	// relayer registration
	ownerRelayer := bind.NewKeyedTransactor(simulation.OwnerRelayerKey)
	nonce, _ = client.NonceAt(context.Background(), simulation.OwnerRelayerAddr, nil)
	relayerRegistration, err = rupx.NewRelayerRegistration(ownerRelayer, relayerRegistrationAddr, client)
	if err != nil {
		log.Fatal("NewRelayerRegistration", err)
	}
	relayerRegistration.TransactOpts.Nonce = big.NewInt(int64(nonce))
	relayerRegistration.TransactOpts.Value = simulation.MinDeposit
	relayerRegistration.TransactOpts.GasPrice = big.NewInt(250000000000000)

	fromTokens := []common.Address{}
	toTokens := []common.Address{}

	/*
		for _, token := range tokenList {
			fromTokens = append(fromTokens, token["address"].(common.Address))
			toTokens = append(toTokens, simulation.RUPXNative)
		}
	*/

	// RUPX/BTC
	fromTokens = append(fromTokens, simulation.RUPXNative)
	toTokens = append(toTokens, tokenList[0]["address"].(common.Address))

	// RUPX/USDT
	fromTokens = append(fromTokens, simulation.RUPXNative)
	toTokens = append(toTokens, tokenList[9]["address"].(common.Address))

	// ETH/RUPX
	fromTokens = append(fromTokens, tokenList[1]["address"].(common.Address))
	toTokens = append(toTokens, simulation.RUPXNative)

	fromTokens = append(fromTokens, tokenList[2]["address"].(common.Address))
	toTokens = append(toTokens, simulation.RUPXNative)

	fromTokens = append(fromTokens, tokenList[3]["address"].(common.Address))
	toTokens = append(toTokens, simulation.RUPXNative)

	fromTokens = append(fromTokens, tokenList[4]["address"].(common.Address))
	toTokens = append(toTokens, simulation.RUPXNative)

	fromTokens = append(fromTokens, tokenList[5]["address"].(common.Address))
	toTokens = append(toTokens, simulation.RUPXNative)

	fromTokens = append(fromTokens, tokenList[6]["address"].(common.Address))
	toTokens = append(toTokens, simulation.RUPXNative)

	fromTokens = append(fromTokens, tokenList[7]["address"].(common.Address))
	toTokens = append(toTokens, simulation.RUPXNative)

	fromTokens = append(fromTokens, tokenList[8]["address"].(common.Address))
	toTokens = append(toTokens, simulation.RUPXNative)

	// ETH/BTC
	fromTokens = append(fromTokens, tokenList[1]["address"].(common.Address))
	toTokens = append(toTokens, tokenList[0]["address"].(common.Address))

	// XRP/BTC
	fromTokens = append(fromTokens, tokenList[2]["address"].(common.Address))
	toTokens = append(toTokens, tokenList[0]["address"].(common.Address))

	// BTC/USDT
	fromTokens = append(fromTokens, tokenList[0]["address"].(common.Address))
	toTokens = append(toTokens, tokenList[9]["address"].(common.Address))

	// ETH/USDT
	fromTokens = append(fromTokens, tokenList[1]["address"].(common.Address))
	toTokens = append(toTokens, tokenList[9]["address"].(common.Address))

	_, err = relayerRegistration.Register(simulation.RelayerCoinbaseAddr, simulation.TradeFee, fromTokens, toTokens)
	if err != nil {
		log.Fatal("relayerRegistration Register ", err)
	}
	fmt.Println("wait 2s to apply token to list rupx")
	time.Sleep(2 * time.Second)

	// Lending apply
	nonce = nonce + 1
	lendingRelayerRegistration, err = rupx.NewLendingRelayerRegistration(ownerRelayer, lendingRelayerRegistrationAddr, client)
	if err != nil {
		log.Fatal("NewRelayerRegistration", err)
	}
	lendingRelayerRegistration.TransactOpts.Nonce = big.NewInt(int64(nonce))
	lendingRelayerRegistration.TransactOpts.Value = big.NewInt(0)
	lendingRelayerRegistration.TransactOpts.GasPrice = big.NewInt(250000000000000)
	lendingRelayerRegistration.TransactOpts.GasLimit = uint64(10000000)

	baseTokens := []common.Address{}
	terms := []*big.Int{}
	collaterals := []common.Address{}

	// USD 1 minute for testing
	baseTokens = append(baseTokens, tokenList[9]["address"].(common.Address))
	terms = append(terms, big.NewInt(60))
	collaterals = append(collaterals, common.HexToAddress("0x0"))

	baseTokens = append(baseTokens, tokenList[9]["address"].(common.Address))
	terms = append(terms, big.NewInt(60))
	collaterals = append(collaterals, tokenList[2]["address"].(common.Address))

	// USD 1 days
	baseTokens = append(baseTokens, tokenList[9]["address"].(common.Address))
	terms = append(terms, big.NewInt(86400))
	collaterals = append(collaterals, common.HexToAddress("0x0"))

	baseTokens = append(baseTokens, tokenList[9]["address"].(common.Address))
	terms = append(terms, big.NewInt(86400))
	collaterals = append(collaterals, tokenList[2]["address"].(common.Address))

	// USD 7 days
	baseTokens = append(baseTokens, tokenList[9]["address"].(common.Address))
	terms = append(terms, big.NewInt(7*86400))
	collaterals = append(collaterals, common.HexToAddress("0x0"))

	// USD 30 days
	baseTokens = append(baseTokens, tokenList[9]["address"].(common.Address))
	terms = append(terms, big.NewInt(30*86400))
	collaterals = append(collaterals, common.HexToAddress("0x0"))

	// RUPX 1 min
	baseTokens = append(baseTokens, simulation.RUPXNative)
	terms = append(terms, big.NewInt(60))
	collaterals = append(collaterals, common.HexToAddress("0x0"))

	// RUPX 1 day
	baseTokens = append(baseTokens, simulation.RUPXNative)
	terms = append(terms, big.NewInt(86400))
	collaterals = append(collaterals, common.HexToAddress("0x0"))

	// RUPX 7 days
	baseTokens = append(baseTokens, simulation.RUPXNative)
	terms = append(terms, big.NewInt(7*86400))
	collaterals = append(collaterals, common.HexToAddress("0x0"))

	// RUPX 30 days
	baseTokens = append(baseTokens, simulation.RUPXNative)
	terms = append(terms, big.NewInt(30*86400))
	collaterals = append(collaterals, common.HexToAddress("0x0"))

	// BTC 1 min
	baseTokens = append(baseTokens, tokenList[0]["address"].(common.Address))
	terms = append(terms, big.NewInt(60))
	collaterals = append(collaterals, common.HexToAddress("0x0"))

	// BTC 1 day
	baseTokens = append(baseTokens, tokenList[0]["address"].(common.Address))
	terms = append(terms, big.NewInt(86400))
	collaterals = append(collaterals, common.HexToAddress("0x0"))

	// BTC 7 days
	baseTokens = append(baseTokens, tokenList[0]["address"].(common.Address))
	terms = append(terms, big.NewInt(7*86400))
	collaterals = append(collaterals, common.HexToAddress("0x0"))

	// BTC 30 days
	baseTokens = append(baseTokens, tokenList[0]["address"].(common.Address))
	terms = append(terms, big.NewInt(30*86400))
	collaterals = append(collaterals, common.HexToAddress("0x0"))

	_, err = lendingRelayerRegistration.Update(simulation.RelayerCoinbaseAddr, simulation.LendingTradeFee, baseTokens, terms, collaterals)
	if err != nil {
		log.Fatal("lendingRelayerRegistration Update", err)
	}

	fmt.Println("wait 2s to update lending contract")
	time.Sleep(2 * time.Second)
}

func initRRC21(auth *bind.TransactOpts, client *ethclient.Client, nonce uint64, tokenNameList []string) []map[string]interface{} {
	tokenListResult := []map[string]interface{}{}
	for _, tokenName := range tokenNameList {
		auth.Nonce = big.NewInt(int64(nonce))
		d := uint8(18)
		depositFee := big.NewInt(0)
		withdrawFee := big.NewInt(0)
		tokenCap := simulation.RRC21TokenCap
		if tokenName == "ADA" {
			d = 0
			tokenCap = new(big.Int).Div(simulation.RRC21TokenCap, simulation.BaseRUPX)
		}
		if tokenName == "USDT" {
			d = 6
			tokenCap = new(big.Int).Div(simulation.RRC21TokenCap, big.NewInt(1000000000000))
			withdrawFee = big.NewInt(970000)
		}
		if tokenName == "BTC" {
			d = 8
			tokenCap = new(big.Int).Div(simulation.RRC21TokenCap, big.NewInt(10000000000))
			withdrawFee = big.NewInt(40000)
		}
		if tokenName == "ETH" {
			withdrawFee = big.NewInt(3000000000000000)
		}
		tokenAddr, _, err := rupx.DeployRRC21(auth, client, simulation.Owners, simulation.Required, tokenName, tokenName, d, tokenCap, simulation.RRC21TokenFee, depositFee, withdrawFee)
		if err != nil {
			log.Fatal("DeployRRC21 ", tokenName, err)
		}

		fmt.Println(tokenName+" token address", tokenAddr.Hex(), "cap", tokenCap)

		tokenListResult = append(tokenListResult, map[string]interface{}{
			"name":     tokenName,
			"address":  tokenAddr,
			"decimals": d,
		})
		nonce = nonce + 1
	}
	time.Sleep(5 * time.Second)
	return tokenListResult
}

func applyIssuer(rrc21Issuer *rupx.RRC21Issuer, tokenList []map[string]interface{}, nonce uint64) {
	for _, token := range tokenList {
		rrc21Issuer.TransactOpts.Nonce = big.NewInt(int64(nonce))
		rrc21Issuer.TransactOpts.Value = simulation.MinRRC21Apply
		_, err := rrc21Issuer.Apply(token["address"].(common.Address))
		if err != nil {
			log.Fatal("rrc21Issuer Apply  ", token["name"].(string), err)
		}
		fmt.Println("applyIssuer ", token["name"].(string))
		nonce = nonce + 1
	}
	time.Sleep(5 * time.Second)
}

func applyRupXListing(rupxListing *rupx.RUPXListing, tokenList []map[string]interface{}, nonce uint64) {
	for _, token := range tokenList {
		rupxListing.TransactOpts.Nonce = big.NewInt(int64(nonce))
		rupxListing.TransactOpts.Value = simulation.RupXListingFee
		_, err := rupxListing.Apply(token["address"].(common.Address))
		if err != nil {
			log.Fatal("rupxListing Apply ", token["name"].(string), err)
		}
		fmt.Println("applyRupXListing ", token["name"].(string))
		nonce = nonce + 1
	}
	time.Sleep(5 * time.Second)
}

func airdrop(auth *bind.TransactOpts, client *ethclient.Client, tokenList []map[string]interface{}, addresses []common.Address, nonce uint64) {
	for _, token := range tokenList {
		for _, address := range addresses {
			rrc21Contract, _ := rupx.NewRRC21(auth, token["address"].(common.Address), client)
			rrc21Contract.TransactOpts.Nonce = big.NewInt(int64(nonce))
			baseAmount := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(token["decimals"].(uint8))), nil)
			amount := big.NewInt(0).Mul(baseAmount, big.NewInt(1000000))
			_, err := rrc21Contract.Transfer(address, amount)
			if err == nil {
				fmt.Printf("Transfer %v %v to %v successfully", amount.String(), token["name"].(string), address.String())
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
