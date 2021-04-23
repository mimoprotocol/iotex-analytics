// Copyright (c) 2019 IoTeX
// This is an alpha (internal) release and is not suitable for production. This source code is provided 'as is' and no
// warranties are given as to title or non-infringement, merchantability or fitness for purpose and, to the extent
// permitted by law, all liability for your use of the code is disclaimed. This source code is governed by Apache
// License 2.0 that can be found in the LICENSE file.

package indexprotocol

import (
	"context"
	"encoding/hex"
	"math/big"

	"github.com/pkg/errors"

	"github.com/iotexproject/go-pkgs/hash"
	"github.com/iotexproject/iotex-address/address"
	"github.com/iotexproject/iotex-core/action"
	"github.com/iotexproject/iotex-proto/golang/iotexapi"
)

var (
	// ErrNotExist indicates certain item does not exist in Blockchain database
	ErrNotExist = errors.New("not exist in DB")
	// ErrAlreadyExist indicates certain item already exists in Blockchain database
	ErrAlreadyExist = errors.New("already exist in DB")
	// ErrUnimplemented indicates a method is not implemented yet
	ErrUnimplemented = errors.New("method is unimplemented")
)

// ConvertTopicToAddress converts topic in log to address
func ConvertTopicToAddress(topic hash.Hash256) (address.Address, error) {
	return address.FromBytes(topic[12:])
}

// ReadContract reads contract
func ReadContract(cli iotexapi.APIServiceClient, addr string, callData []byte) ([]byte, error) {
	execution, err := action.NewExecution(addr, uint64(1), big.NewInt(0), uint64(3000000), big.NewInt(1), callData)
	if err != nil {
		return nil, err
	}
	request := &iotexapi.ReadContractRequest{
		Execution:     execution.Proto(),
		CallerAddress: address.ZeroAddress,
	}

	res, err := cli.ReadContract(context.Background(), request)
	if err != nil {
		return nil, err
	}
	if res.Receipt.Status != uint64(1) {
		return nil, errors.Wrap(ErrUnimplemented, "failed to read contract")
	}
	return hex.DecodeString(res.Data)
}
