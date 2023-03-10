package state

import (
	"math/big"
	"time"

	"github.com/0xPolygonHermez/zkevm-node/state/runtime/instrumentation"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// ProcessBatchResponse represents the response of a batch process.
type ProcessBatchResponse struct {
	NewStateRoot       common.Hash
	NewAccInputHash    common.Hash
	NewLocalExitRoot   common.Hash
	NewBatchNumber     uint64
	UsedZkCounters     ZKCounters
	Responses          []*ProcessTransactionResponse
	ExecutorError      error
	IsBatchProcessed   bool
	ReadWriteAddresses []*InfoReadWrite
}

// ProcessTransactionResponse represents the response of a tx process.
type ProcessTransactionResponse struct {
	// TxHash is the hash of the transaction
	TxHash common.Hash
	// Type indicates legacy transaction
	// It will be always 0 (legacy) in the executor
	Type uint32
	// ReturnValue is the returned data from the runtime (function result or data supplied with revert opcode)
	ReturnValue []byte
	// GasLeft is the total gas left as result of execution
	GasLeft uint64
	// GasUsed is the total gas used as result of execution or gas estimation
	GasUsed uint64
	// GasRefunded is the total gas refunded as result of execution
	GasRefunded uint64
	// RomError represents any error encountered during the execution
	RomError error
	// CreateAddress is the new SC Address in case of SC creation
	CreateAddress common.Address
	// StateRoot is the State Root
	StateRoot common.Hash
	// Logs emitted by LOG opcode
	Logs []*types.Log
	// IsProcessed indicates if this tx didn't fit into the batch
	IsProcessed bool
	// Tx is the whole transaction object
	Tx types.Transaction
	// ExecutionTrace contains the traces produced in the execution
	ExecutionTrace []instrumentation.StructLog
	// CallTrace contains the call trace.
	CallTrace instrumentation.ExecutorTrace
}

// ZKCounters counters for the tx
type ZKCounters struct {
	CumulativeGasUsed    uint64
	UsedKeccakHashes     uint32
	UsedPoseidonHashes   uint32
	UsedPoseidonPaddings uint32
	UsedMemAligns        uint32
	UsedArithmetics      uint32
	UsedBinaries         uint32
	UsedSteps            uint32
}

// InfoReadWrite has information about modified addresses during the execution
type InfoReadWrite struct {
	Address common.Address
	Nonce   *uint64
	Balance *big.Int
}

const (
	// DebugInfoErrorType_EXECUTOR_ERROR indicates a error happened in the executor
	DebugInfoErrorType_EXECUTOR_ERROR = "EXECUTOR ERROR"
)

// DebugInfo allows handling runtime debug info
type DebugInfo struct {
	ErrorType string
	Timestamp time.Time
	Payload   string
}
