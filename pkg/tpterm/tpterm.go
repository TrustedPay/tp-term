package tpterm

import (
	"context"
	"crypto/x509"
	"encoding/json"
	"encoding/pem"
	"fmt"
	"log"
	"sync"

	"github.com/TrustedPay/tp-term/internal/tpm"
	"github.com/google/go-tpm-tools/client"
	"github.com/google/go-tpm-tools/simulator"
	"github.com/sirupsen/logrus"
)

type TPTerm struct {
	UnimplementedTPTermServer

	lock *sync.Mutex
	tpm  *tpm.TPM
	key  *client.Key
	rw   *simulator.Simulator
}

func NewTPTerm() *TPTerm {
	t := tpm.DefaultTPM()

	// TODO: use real TPM.
	simulator, err := simulator.Get()
	if err != nil {
		log.Fatalf("failed to initialize simulator: %v", err)
	}

	key, err := t.GenerateKey(simulator)
	if err != nil {
		logrus.Fatalf("failed to generate key")
	}

	keyBytes, err := x509.MarshalPKIXPublicKey(key.PublicKey())
	if err != nil {
		log.Fatalf("failed to marshal public key")
	}
	pemBlock := pem.EncodeToMemory(&pem.Block{
		Type:    "PUBLIC KEY",
		Headers: map[string]string{},
		Bytes:   keyBytes,
	})
	logrus.Printf("Copy this public key:\n%s", pemBlock)
	parsedBlock, _ := pem.Decode(pemBlock)
	if parsedBlock == nil {
		logrus.Fatalf("failed to parse PEM block back")
	}

	return &TPTerm{
		lock: new(sync.Mutex),
		tpm:  t,
		key:  key,
		rw:   simulator,
	}
}

func (term *TPTerm) Shutdown() error {
	return term.rw.Close()
}

func (term *TPTerm) SignRequest(ctx context.Context, data *Transaction) (*TransactionSignature, error) {
	// Wait to acquire mutex
	term.lock.Lock()
	defer term.lock.Unlock()

	// Sign the transaction
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal transaction: %v", err)
	}
	digest, sig, err := term.tpm.SignData(term.rw, dataBytes, term.key)
	if err != nil {
		return nil, fmt.Errorf("failed to sign transaction: %v", err)
	}
	logrus.Printf("generated signature: %x", sig)

	return &TransactionSignature{
		TransactionDigest:    digest,
		TransactionSignature: sig,
	}, nil
}
