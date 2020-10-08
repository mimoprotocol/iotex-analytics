// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package indexprotocol

import (
	"context"
	"database/sql"

	"github.com/iotexproject/iotex-core/blockchain/block"
	"github.com/iotexproject/iotex-proto/golang/iotextypes"
)

type (
	// BlockData defines the block data of a specific height
	BlockData struct {
		Block           *block.Block
		TransactionLogs []*iotextypes.TransactionLog
	}

	// ProtocolV2 defines the protocol interfaces for block indexer
	ProtocolV2 interface {
		BlockDataHandler
		Initialize(context.Context, *sql.Tx) error
	}

	// BlockDataHandler is the interface of handling block data
	BlockDataHandler interface {
		HandleBlockData(context.Context, *sql.Tx, *BlockData) error
	}
)
