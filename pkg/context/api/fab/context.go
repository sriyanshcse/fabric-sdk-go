/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package fab

import "github.com/hyperledger/fabric-sdk-go/pkg/context"

// ChannelService supplies services related to a channel.
type ChannelService interface {
	Config() (ChannelConfig, error)
	Ledger() (ChannelLedger, error)
	Transactor() (Transactor, error)
	EventHub() (EventHub, error) // TODO support new event delivery
	Membership() (ChannelMembership, error)
}

// Transactor supplies methods for sending transaction proposals and transactions.
type Transactor interface {
	Sender
	ProposalSender
}

// ChannelProvider supplies Channel related-objects for the named channel.
type ChannelProvider interface {
	ChannelService(ic context.IdentityContext, channelID string) (ChannelService, error)
}
