package handler

import (
	"github.com/vitelabs/go-vite/protocols"
	"github.com/vitelabs/go-vite/ledger/access"
	"github.com/vitelabs/go-vite/common/types"
	"github.com/vitelabs/go-vite/ledger"
)

type AccountChain struct {
	vite Vite
	// Handle block
	acAccess *access.AccountChainAccess
}

func NewAccountChain (vite Vite) (*AccountChain) {
	return &AccountChain{
		vite: vite,
		acAccess: access.GetAccountChainAccess(),
	}
}

// HandleBlockHash
func (ac *AccountChain) HandleGetBlocks (msg *protocols.GetAccountBlocksMsg, peer *protocols.Peer) error {
	go func() {
		ac.acAccess.GetBlocksFromOrigin(&msg.Origin, msg.Count, msg.Forward)
		// send out
	}()
	return nil
}

// HandleBlockHash
func (ac *AccountChain) HandleSendBlocks (msg protocols.AccountBlocksMsg, peer *protocols.Peer) error {
	go func() {
		ac.acAccess.WriteBlockList(msg)
	}()
	return nil
}

func (ac *AccountChain) CreateTx (a types.Address, block *ledger.AccountBlock) {
	
}

func (ac *AccountChain) CreateTxWithPassphrase (a types.Address, passphrase string, block *ledger.AccountBlock) {

}