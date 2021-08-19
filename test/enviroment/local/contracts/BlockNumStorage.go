// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ContractsABI is the input ABI used to generate the binding from.
const ContractsABI = "[{\"inputs\":[],\"name\":\"store\",\"outputs\":[{\"internalType\":\"contractUintStorage\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"key\",\"type\":\"string\"}],\"name\":\"sync\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// ContractsBin is the compiled bytecode used for deploying new contracts.
var ContractsBin = "0x60806040526040516100109061007e565b604051809103906000f08015801561002c573d6000803e3d6000fd5b506000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555034801561007857600080fd5b5061008b565b6103ef806106c283390190565b6106288061009a6000396000f3fe6080604052600436106100295760003560e01c8063975057e71461002e578063fd620be114610059575b600080fd5b34801561003a57600080fd5b50610043610089565b60405161005091906103e0565b60405180910390f35b610073600480360381019061006e919061031f565b6100ad565b604051610080919061044d565b60405180910390f35b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e846040518263ffffffff1660e01b815260040161010991906103fb565b60206040518083038186803b15801561012157600080fd5b505afa158015610135573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906101599190610360565b90508043146101f15760008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638a42ebe93485436040518463ffffffff1660e01b81526004016101be92919061041d565b6000604051808303818588803b1580156101d757600080fd5b505af11580156101eb573d6000803e3d6000fd5b50505050505b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663693ec85e846040518263ffffffff1660e01b815260040161024a91906103fb565b60206040518083038186803b15801561026257600080fd5b505afa158015610276573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061029a9190610360565b915050919050565b60006102b56102b08461048d565b610468565b9050828152602081018484840111156102cd57600080fd5b6102d8848285610528565b509392505050565b600082601f8301126102f157600080fd5b81356103018482602086016102a2565b91505092915050565b600081519050610319816105db565b92915050565b60006020828403121561033157600080fd5b600082013567ffffffffffffffff81111561034b57600080fd5b610357848285016102e0565b91505092915050565b60006020828403121561037257600080fd5b60006103808482850161030a565b91505092915050565b61039281610504565b82525050565b60006103a3826104be565b6103ad81856104c9565b93506103bd818560208601610537565b6103c6816105ca565b840191505092915050565b6103da816104fa565b82525050565b60006020820190506103f56000830184610389565b92915050565b600060208201905081810360008301526104158184610398565b905092915050565b600060408201905081810360008301526104378185610398565b905061044660208301846103d1565b9392505050565b600060208201905061046260008301846103d1565b92915050565b6000610472610483565b905061047e828261056a565b919050565b6000604051905090565b600067ffffffffffffffff8211156104a8576104a761059b565b5b6104b1826105ca565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061050f82610516565b9050919050565b6000610521826104da565b9050919050565b82818337600083830152505050565b60005b8381101561055557808201518184015260208101905061053a565b83811115610564576000848401525b50505050565b610573826105ca565b810181811067ffffffffffffffff821117156105925761059161059b565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b6105e4816104fa565b81146105ef57600080fd5b5056fea264697066735822122022613cf7fdb3c20aa0c44f02035f9eeb0947a906e70b235428e6b3d6f370c3f364736f6c63430008010033608060405234801561001057600080fd5b506103cf806100206000396000f3fe6080604052600436106100295760003560e01c8063693ec85e1461002e5780638a42ebe91461006b575b600080fd5b34801561003a57600080fd5b5061005560048036038101906100509190610152565b610087565b604051610062919061023e565b60405180910390f35b61008560048036038101906100809190610193565b6100ae565b005b600080826040516100989190610227565b9081526020016040518091039020549050919050565b806000836040516100bf9190610227565b9081526020016040518091039020819055505050565b60006100e86100e38461027e565b610259565b90508281526020810184848401111561010057600080fd5b61010b8482856102cf565b509392505050565b600082601f83011261012457600080fd5b81356101348482602086016100d5565b91505092915050565b60008135905061014c81610382565b92915050565b60006020828403121561016457600080fd5b600082013567ffffffffffffffff81111561017e57600080fd5b61018a84828501610113565b91505092915050565b600080604083850312156101a657600080fd5b600083013567ffffffffffffffff8111156101c057600080fd5b6101cc85828601610113565b92505060206101dd8582860161013d565b9150509250929050565b60006101f2826102af565b6101fc81856102ba565b935061020c8185602086016102de565b80840191505092915050565b610221816102c5565b82525050565b600061023382846101e7565b915081905092915050565b60006020820190506102536000830184610218565b92915050565b6000610263610274565b905061026f8282610311565b919050565b6000604051905090565b600067ffffffffffffffff82111561029957610298610342565b5b6102a282610371565b9050602081019050919050565b600081519050919050565b600081905092915050565b6000819050919050565b82818337600083830152505050565b60005b838110156102fc5780820151818401526020810190506102e1565b8381111561030b576000848401525b50505050565b61031a82610371565b810181811067ffffffffffffffff8211171561033957610338610342565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b61038b816102c5565b811461039657600080fd5b5056fea26469706673582212202ee47e81a6c678b620e8b11565954c4b3857502f3e7cfc1f9e844083d1add93064736f6c63430008010033"

// DeployContracts deploys a new Ethereum contract, binding an instance of Contracts to it.
func DeployContracts(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Contracts, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ContractsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// Contracts is an auto generated Go binding around an Ethereum contract.
type Contracts struct {
	ContractsCaller     // Read-only binding to the contract
	ContractsTransactor // Write-only binding to the contract
	ContractsFilterer   // Log filterer for contract events
}

// ContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractsSession struct {
	Contract     *Contracts        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractsCallerSession struct {
	Contract *ContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractsTransactorSession struct {
	Contract     *ContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractsRaw struct {
	Contract *Contracts // Generic contract binding to access the raw methods on
}

// ContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractsCallerRaw struct {
	Contract *ContractsCaller // Generic read-only contract binding to access the raw methods on
}

// ContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractsTransactorRaw struct {
	Contract *ContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContracts creates a new instance of Contracts, bound to a specific deployed contract.
func NewContracts(address common.Address, backend bind.ContractBackend) (*Contracts, error) {
	contract, err := bindContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Contracts{ContractsCaller: ContractsCaller{contract: contract}, ContractsTransactor: ContractsTransactor{contract: contract}, ContractsFilterer: ContractsFilterer{contract: contract}}, nil
}

// NewContractsCaller creates a new read-only instance of Contracts, bound to a specific deployed contract.
func NewContractsCaller(address common.Address, caller bind.ContractCaller) (*ContractsCaller, error) {
	contract, err := bindContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsCaller{contract: contract}, nil
}

// NewContractsTransactor creates a new write-only instance of Contracts, bound to a specific deployed contract.
func NewContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractsTransactor, error) {
	contract, err := bindContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractsTransactor{contract: contract}, nil
}

// NewContractsFilterer creates a new log filterer instance of Contracts, bound to a specific deployed contract.
func NewContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractsFilterer, error) {
	contract, err := bindContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractsFilterer{contract: contract}, nil
}

// bindContracts binds a generic wrapper to an already deployed contract.
func bindContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.ContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.ContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Contracts *ContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Contracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Contracts *ContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Contracts *ContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Contracts.Contract.contract.Transact(opts, method, params...)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() view returns(address)
func (_Contracts *ContractsCaller) Store(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "store")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() view returns(address)
func (_Contracts *ContractsSession) Store() (common.Address, error) {
	return _Contracts.Contract.Store(&_Contracts.CallOpts)
}

// Store is a free data retrieval call binding the contract method 0x975057e7.
//
// Solidity: function store() view returns(address)
func (_Contracts *ContractsCallerSession) Store() (common.Address, error) {
	return _Contracts.Contract.Store(&_Contracts.CallOpts)
}

// Sync is a paid mutator transaction binding the contract method 0xfd620be1.
//
// Solidity: function sync(string key) payable returns(uint256)
func (_Contracts *ContractsTransactor) Sync(opts *bind.TransactOpts, key string) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "sync", key)
}

// Sync is a paid mutator transaction binding the contract method 0xfd620be1.
//
// Solidity: function sync(string key) payable returns(uint256)
func (_Contracts *ContractsSession) Sync(key string) (*types.Transaction, error) {
	return _Contracts.Contract.Sync(&_Contracts.TransactOpts, key)
}

// Sync is a paid mutator transaction binding the contract method 0xfd620be1.
//
// Solidity: function sync(string key) payable returns(uint256)
func (_Contracts *ContractsTransactorSession) Sync(key string) (*types.Transaction, error) {
	return _Contracts.Contract.Sync(&_Contracts.TransactOpts, key)
}
