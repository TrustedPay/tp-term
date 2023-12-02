package tpterm

import (
	"context"
	"sync"

	"github.com/sirupsen/logrus"
)

type TPTermState int

const (
	State_READY             TPTermState = 0
	State_UNVERIFIED        TPTermState = 1
	State_AWAITING_RESPONSE TPTermState = 2
	State_DECLINED          TPTermState = 3
	State_TRANSACTION_OK    TPTermState = 4
	State_TRY_ANOTHER       TPTermState = 5
	State_TERMINAL_CLOSED   TPTermState = 6
)

type TPTerm struct {
	UnimplementedTPTermServer

	lock *sync.Mutex

	state TPTermState
}

func NewTPTerm() *TPTerm {
	return &TPTerm{
		lock:  new(sync.Mutex),
		state: State_READY,
	}
}

func (term *TPTerm) SignRequest(ctx context.Context, data *Transaction) (*TransactionSignature, error) {

	// Wait to acquire mutex
	term.lock.Lock()
	defer term.lock.Unlock()

	// Sign the transaction
	sig := []byte("signature here") // TODO: sign the transaction
	logrus.Printf("generated signature: %x", sig)

	return &TransactionSignature{
		TransactionSignature: sig,
	}, nil
}
