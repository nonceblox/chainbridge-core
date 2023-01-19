// Copyright 2021 ChainSafe Systems
// SPDX-License-Identifier: LGPL-3.0-only

package relayer

import (
	"context"
	"fmt"

	"github.com/VaivalGithub/chainbridge-core/relayer/message"
	"github.com/rs/zerolog/log"
)

type Metrics interface {
	TrackDepositMessage(m *message.Message)
}

type RelayedChain interface {
	PollEvents(ctx context.Context, sysErr chan<- error, msgChan chan *message.Message, msgChan1 chan *message.Message2)
	Write(message *message.Message) error
	Write1(message *message.Message2) (bool, error)
	Write2(message *message.Message2) error
	WriteRemoval(message *message.Message2) error
	DomainID() uint8
	CheckFeeClaim() bool
	GetFeeClaim(msg *message.Message) error
}

func NewRelayer(chains []RelayedChain, metrics Metrics, messageProcessors ...message.MessageProcessor) *Relayer {
	return &Relayer{relayedChains: chains, messageProcessors: messageProcessors, metrics: metrics}
}

type Relayer struct {
	metrics           Metrics
	relayedChains     []RelayedChain
	registry          map[uint8]RelayedChain
	messageProcessors []message.MessageProcessor
}

// Start function starts the relayer. Relayer routine is starting all the chains
// and passing them with a channel that accepts unified cross chain message format
func (r *Relayer) Start(ctx context.Context, sysErr chan error) {
	log.Debug().Msgf("Starting relayer")

	messagesChannel := make(chan *message.Message)
	messagesChannel1 := make(chan *message.Message2)
	for _, c := range r.relayedChains {
		log.Debug().Msgf("Starting chain %v", c.DomainID())
		r.addRelayedChain(c)
		go c.PollEvents(ctx, sysErr, messagesChannel, messagesChannel1)
	}

	for {
		select {
		case m := <-messagesChannel:
			go r.route(m)
			continue
		case n := <-messagesChannel1:
			go r.registerroute(n)
			continue
		case <-ctx.Done():
			return

		}
	}

}

// Route function winds destination writer by mapping DestinationID from message to registered writer.
func (r *Relayer) route(m *message.Message) {
	r.metrics.TrackDepositMessage(m)

	destChain, ok := r.registry[m.Destination]

	if !ok {
		log.Error().Msgf("no resolver for destID %v to send message registered", m.Destination)
		return
	}
	sorcChain, ok := r.registry[m.Source]
	if !ok {
		log.Error().Msgf("no resolver for source %v to send message registered", m.Source)
		return
	}

	for _, mp := range r.messageProcessors {
		if err := mp(m); err != nil {
			log.Error().Err(fmt.Errorf("error %w processing mesage %v", err, m))
			return
		}
	}

	log.Debug().Msgf("Sending message %+v to destination %v", m, m.Destination)
	// fee method here.
	boolVal := destChain.CheckFeeClaim()
	if boolVal {
		err := sorcChain.GetFeeClaim(m)
		if err != nil {
			log.Error().Err(err).Msgf("Claiming fees Error %+w", err)
		}
	}
	if err := destChain.Write(m); err != nil {
		log.Error().Err(err).Msgf("writing message %+v", m)
		return
	}
}

// Route function winds destination writer by mapping DestinationID from message to registered writer.
func (r *Relayer) registerroute(n *message.Message2) {
	destChain, ok := r.registry[n.Destination]
	if !ok {
		log.Error().Msgf("no resolver for destID %v to send message registered", n.Destination)
		return
	}
	sorcChain, ok := r.registry[n.Source]
	if !ok {
		log.Error().Msgf("no resolver for destID %v to send message registered", n.Source)
		return
	}
	a, err := destChain.Write1(n)
	if err != nil {
		log.Error().Err(err).Msgf("writing message %+v", n)
		return
	}
	if a {
		err := sorcChain.Write2(n)
		if err != nil {
			sorcChain.WriteRemoval(n)
		}

	}
}

func (r *Relayer) addRelayedChain(c RelayedChain) {
	if r.registry == nil {
		r.registry = make(map[uint8]RelayedChain)
	}
	domainID := c.DomainID()
	r.registry[domainID] = c
}
