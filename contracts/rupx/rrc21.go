package rupx

import (
	"math/big"

	"github.com/rupayaproject/rupaya/accounts/abi/bind"
	"github.com/rupayaproject/rupaya/common"
	"github.com/rupayaproject/rupaya/contracts/rupx/contract"
)

type MyRRC21 struct {
	*contract.MyRRC21Session
	contractBackend bind.ContractBackend
}

func NewRRC21(transactOpts *bind.TransactOpts, contractAddr common.Address, contractBackend bind.ContractBackend) (*MyRRC21, error) {
	smartContract, err := contract.NewMyRRC21(contractAddr, contractBackend)
	if err != nil {
		return nil, err
	}

	return &MyRRC21{
		&contract.MyRRC21Session{
			Contract:     smartContract,
			TransactOpts: *transactOpts,
		},
		contractBackend,
	}, nil
}

func DeployRRC21(transactOpts *bind.TransactOpts, contractBackend bind.ContractBackend, owners []common.Address, required *big.Int, name string, symbol string, decimals uint8, cap, fee, depositFee, withdrawFee *big.Int) (common.Address, *MyRRC21, error) {
	contractAddr, _, _, err := contract.DeployMyRRC21(transactOpts, contractBackend, owners, required, name, symbol, decimals, cap, fee, depositFee, withdrawFee)
	if err != nil {
		return contractAddr, nil, err
	}
	smartContract, err := NewRRC21(transactOpts, contractAddr, contractBackend)
	if err != nil {
		return contractAddr, nil, err
	}

	return contractAddr, smartContract, nil
}
