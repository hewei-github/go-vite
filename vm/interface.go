package vm

import (
	"github.com/vitelabs/go-vite/common/types"
	"math/big"
)

const (
	BlockTypeSendCreate = iota
	BlockTypeSendCall
	BlockTypeSendMintage
	BlockTypeSendReward
	BlockTypeReceive
	BlockTypeReceiveError
)

type CreateAccountBlockFunc func(from, to types.Address, blockType, depth uint64) VmAccountBlock

type VmAccountBlock interface {
	// Account block height
	Height() *big.Int
	SetHeight(*big.Int)
	// Receiver account address
	ToAddress() types.Address
	SetToAddress(types.Address)
	// Sender account address
	AccountAddress() types.Address
	SetAccountAddress(types.Address)
	// Sender block hash, exists in receive block
	FromBlockHash() types.Hash
	SetFromBlockHash(types.Hash)
	// Transaction type of current block
	BlockType() uint64
	SetBlockType(uint64)
	// Last block hash
	PrevHash() types.Hash
	SetPrevHash(types.Hash)
	// Amount of this transaction
	Amount() *big.Int
	SetAmount(*big.Int)
	// Id of token received or sent
	TokenId() types.TokenTypeId
	SetTokenId(types.TokenTypeId)
	// Create contract fee of vite token
	CreateFee() *big.Int
	SetCreateFee(*big.Int)
	// Input data
	Data() []byte
	SetData([]byte)
	// State root hash of current account
	StateHash() types.Hash
	SetStateHash(types.Hash)
	// Send block summary hash list
	SendBlockHashList() []types.Hash
	AppendSendBlockHash(types.Hash)
	// Root hash of log list
	LogListHash() types.Hash
	SetLogListHash(types.Hash)
	// Snapshot block hash
	SnapshotHash() types.Hash
	SetSnapshotHash(types.Hash)
	// Call or create depth of current account block
	Depth() uint64
	SetDepth(uint64)
	// Quota used of current block
	Quota() uint64
	SetQuota(uint64)
	// Hash value of HEIGHT, AccountAddress, ToAddress, BlockType, Amount, TokenTypeId, Data, Depth
	SummaryHash() types.Hash
}

type VmSnapshotBlock interface {
	Height() *big.Int
	Timestamp() int64
	Hash() types.Hash
	PrevHash() types.Hash
	Producer() types.Address
}

type Gid [10]byte

type Log struct {
	// list of topics provided by the contract
	Topics []types.Hash
	// supplied by the contract, usually ABI-encoded
	Data []byte
}

type VmDatabase interface {
	Balance(addr types.Address, tokenId types.TokenTypeId) *big.Int
	SubBalance(addr types.Address, tokenId types.TokenTypeId, amount *big.Int)
	AddBalance(addr types.Address, tokenId types.TokenTypeId, amount *big.Int)

	SnapshotBlock(snapshotHash types.Hash) VmSnapshotBlock
	SnapshotBlockByHeight(height *big.Int) VmSnapshotBlock
	// forward=true return (startHeight, startHeight+count], forward=false return [startHeight-count, start)
	SnapshotBlockList(startHeight *big.Int, count uint64, forward bool) []VmSnapshotBlock

	AccountBlock(addr types.Address, blockHash types.Hash) VmAccountBlock

	Rollback()

	IsExistAddress(addr types.Address) bool

	IsExistToken(tokenId types.TokenTypeId) bool
	CreateToken(tokenId types.TokenTypeId, tokenName string, owner types.Address, totelSupply *big.Int, decimals uint64) bool

	SetContractGid(addr types.Address, gid Gid, open bool)
	SetContractCode(addr types.Address, gid Gid, code []byte)
	ContractCode(addr types.Address) []byte

	Storage(addr types.Address, loc types.Hash) []byte
	SetStorage(addr types.Address, loc types.Hash, value []byte)
	PrintStorage(addr types.Address) string
	StorageHash(addr types.Address) types.Hash

	AddLog(*Log)
	LogListHash() types.Hash

	IsExistGid(gid Gid) bool
}