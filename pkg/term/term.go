package term

import (
	context "context"
	"fmt"
	sync "sync"

	proto "github.com/TrustedPay/tp-term/proto"
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

func (term *TPTerm) PaymentRequest(ctx context.Context, data *proto.Transaction) (*PaymentRequestReply, error) {

	// Wait to acquire mutex
	term.lock.Lock()
	defer term.lock.Unlock()

	// Check if in ready state
	if term.state != State_READY {
		return nil, fmt.Errorf("TP Term not in ready state")
	}

	// Move to UNVERIFIED state
	term.state = State_UNVERIFIED

	// Sign the transaction
	sig := []byte("signature here") // TODO: sign the transaction
	logrus.Printf("generated signature: %x", sig)

	// Move to AWAITING_RESPONSE state
	term.state = State_AWAITING_RESPONSE

	// Send the transaction to the bank for processing
	// TODO: send to bank server
	logrus.Printf("sending to bank for processing")

	// Move back to READY state
	term.state = State_READY

	return &PaymentRequestReply{
		Status:    proto.TransactionStatus_OK,
		PaymentId: "some payment id here",
	}, nil
}
