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
const ContractsABI = "[{\"inputs\":[],\"name\":\"watch\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"watcher\",\"outputs\":[{\"internalType\":\"contractBlockWatcher\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ContractsBin is the compiled bytecode used for deploying new contracts.
var ContractsBin = "0x60806040526040516100109061007e565b604051809103906000f08015801561002c573d6000803e3d6000fd5b506000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555034801561007857600080fd5b5061008b565b6106c28061030a83390190565b6102708061009a6000396000f3fe6080604052600436106100295760003560e01c80633489d8741461002e578063d6b748651461004c575b600080fd5b610036610077565b60405161004391906101ba565b60405180910390f35b34801561005857600080fd5b5061006161011f565b60405161006e919061019f565b60405180910390f35b60008060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16633489d8746040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156100e257600080fd5b505af11580156100f6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061011a9190610158565b905090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60008151905061015281610223565b92915050565b60006020828403121561016a57600080fd5b600061017884828501610143565b91505092915050565b61018a816101ff565b82525050565b610199816101f5565b82525050565b60006020820190506101b46000830184610181565b92915050565b60006020820190506101cf6000830184610190565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b600061020a82610211565b9050919050565b600061021c826101d5565b9050919050565b61022c816101f5565b811461023757600080fd5b5056fea26469706673582212204ff9c23edce3dbcc6e3188a3b4816370ca51a6f8eda4ac0851d08ce8fe84d20664736f6c6343000801003360806040526040516100109061007e565b604051809103906000f08015801561002c573d6000803e3d6000fd5b506000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555034801561007857600080fd5b5061008b565b61021f806104a383390190565b6104098061009a6000396000f3fe6080604052600436106100295760003560e01c80633489d8741461002e578063975057e71461004c575b600080fd5b610036610077565b6040516100439190610353565b60405180910390f35b34801561005857600080fd5b506100616102b8565b60405161006e9190610338565b60405180910390f35b60008060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636d4ce63c6040518163ffffffff1660e01b815260040160206040518083038186803b1580156100e057600080fd5b505afa1580156100f4573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061011891906102f1565b90508043146102145760008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166360fe47b1436040518263ffffffff1660e01b815260040161017a9190610353565b600060405180830381600087803b15801561019457600080fd5b505af11580156101a8573d6000803e3d6000fd5b5050505060008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166108fc349081150290604051600060405180830381858888f19350505050158015610212573d6000803e3d6000fd5b505b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16636d4ce63c6040518163ffffffff1660e01b815260040160206040518083038186803b15801561027a57600080fd5b505afa15801561028e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102b291906102f1565b91505090565b60008054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000815190506102eb816103bc565b92915050565b60006020828403121561030357600080fd5b6000610311848285016102dc565b91505092915050565b61032381610398565b82525050565b6103328161038e565b82525050565b600060208201905061034d600083018461031a565b92915050565b60006020820190506103686000830184610329565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006103a3826103aa565b9050919050565b60006103b58261036e565b9050919050565b6103c58161038e565b81146103d057600080fd5b5056fea2646970667358221220e4116533309233a7c464d195080494f443aa38ec94dc48ca3dcdca36a4e60e0c64736f6c63430008010033608060405234801561001057600080fd5b506101ff806100206000396000f3fe60806040526004361061002d5760003560e01c806360fe47b1146100395780636d4ce63c1461006257610034565b3661003457005b600080fd5b34801561004557600080fd5b50610060600480360381019061005b91906100e9565b61008d565b005b34801561006e57600080fd5b506100776100b1565b6040516100849190610159565b60405180910390f35b80600060405161009c90610144565b90815260200160405180910390208190555050565b6000806040516100c090610144565b908152602001604051809103902054905090565b6000813590506100e3816101b2565b92915050565b6000602082840312156100fb57600080fd5b6000610109848285016100d4565b91505092915050565b600061011f600383610174565b915061012a82610189565b600382019050919050565b61013e8161017f565b82525050565b600061014f82610112565b9150819050919050565b600060208201905061016e6000830184610135565b92915050565b600081905092915050565b6000819050919050565b7f636e740000000000000000000000000000000000000000000000000000000000600082015250565b6101bb8161017f565b81146101c657600080fd5b5056fea26469706673582212209f4fe48185273d5145b6d1efb94cc3c3ded77dfb3624cbf25853b252e9521f2864736f6c63430008010033"

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

// Watcher is a free data retrieval call binding the contract method 0xd6b74865.
//
// Solidity: function watcher() view returns(address)
func (_Contracts *ContractsCaller) Watcher(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Contracts.contract.Call(opts, &out, "watcher")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Watcher is a free data retrieval call binding the contract method 0xd6b74865.
//
// Solidity: function watcher() view returns(address)
func (_Contracts *ContractsSession) Watcher() (common.Address, error) {
	return _Contracts.Contract.Watcher(&_Contracts.CallOpts)
}

// Watcher is a free data retrieval call binding the contract method 0xd6b74865.
//
// Solidity: function watcher() view returns(address)
func (_Contracts *ContractsCallerSession) Watcher() (common.Address, error) {
	return _Contracts.Contract.Watcher(&_Contracts.CallOpts)
}

// Watch is a paid mutator transaction binding the contract method 0x3489d874.
//
// Solidity: function watch() payable returns(uint256)
func (_Contracts *ContractsTransactor) Watch(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Contracts.contract.Transact(opts, "watch")
}

// Watch is a paid mutator transaction binding the contract method 0x3489d874.
//
// Solidity: function watch() payable returns(uint256)
func (_Contracts *ContractsSession) Watch() (*types.Transaction, error) {
	return _Contracts.Contract.Watch(&_Contracts.TransactOpts)
}

// Watch is a paid mutator transaction binding the contract method 0x3489d874.
//
// Solidity: function watch() payable returns(uint256)
func (_Contracts *ContractsTransactorSession) Watch() (*types.Transaction, error) {
	return _Contracts.Contract.Watch(&_Contracts.TransactOpts)
}
