// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package evm

import (
	"context"
	"fmt"
	"math/big"

	"github.com/nonceblox/chainbridge-core/config/chain"
	"github.com/nonceblox/chainbridge-core/relayer/message"
	"github.com/nonceblox/chainbridge-core/store"
	"github.com/rs/zerolog/log"
)

type EventListener interface {
	ListenToEvents(ctx context.Context, startBlock *big.Int, msgChan chan *message.Message, msgChan1 chan *message.Message2, errChan chan<- error)
}

type ProposalExecutor interface {
	Execute(message *message.Message) error
	Execute1(message *message.Message2) (bool, error)
	ExecuteSourceTransactiions(message *message.Message2) error
	ExecuteRemovefromdest(message *message.Message2) error
	FeeClaimByRelayer(p *message.Message) error
	IsFeeThresholdReached() bool
}

// EVMChain is struct that aggregates all data required for
type EVMChain struct {
	listener   EventListener
	writer     ProposalExecutor
	blockstore *store.BlockStore
	config     *chain.EVMConfig
}

func NewEVMChain(listener EventListener, writer ProposalExecutor, blockstore *store.BlockStore, config *chain.EVMConfig) *EVMChain {
	return &EVMChain{listener: listener, writer: writer, blockstore: blockstore, config: config}
}

// PollEvents is the goroutine that polls blocks and searches Deposit events in them.
// Events are then sent to eventsChan.
func (c *EVMChain) PollEvents(ctx context.Context, sysErr chan<- error, msgChan chan *message.Message, msgChan1 chan *message.Message2) {
	log.Info().Msg("Polling Blocks...")

	startBlock, err := c.blockstore.GetStartBlock(
		*c.config.GeneralChainConfig.Id,
		c.config.StartBlock,
		c.config.GeneralChainConfig.LatestBlock,
		c.config.GeneralChainConfig.FreshStart,
	)
	if err != nil {
		sysErr <- fmt.Errorf("error %w on getting last stored block", err)
		return
	}

	go c.listener.ListenToEvents(ctx, startBlock, msgChan, msgChan1, sysErr)
}

func (c *EVMChain) Write(msg *message.Message) error {
	return c.writer.Execute(msg)
}

func (c *EVMChain) Write1(msg1 *message.Message2) (bool, error) {
	a, err := c.writer.Execute1(msg1)
	return a, err
}
func (c *EVMChain) Write2(msg *message.Message2) error {
	return c.writer.ExecuteSourceTransactiions(msg)
}
func (c *EVMChain) WriteRemoval(msg *message.Message2) error {
	return c.writer.ExecuteRemovefromdest(msg)
}
func (c *EVMChain) DomainID() uint8 {
	return *c.config.GeneralChainConfig.Id
}
func (c *EVMChain) CheckFeeClaim() bool {
	return c.writer.IsFeeThresholdReached()
}
func (c *EVMChain) GetFeeClaim(msg *message.Message) error {
	return c.writer.FeeClaimByRelayer(msg)
}
