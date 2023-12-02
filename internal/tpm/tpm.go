package tpm

import (
	"crypto"
	"fmt"
	"io"

	"github.com/google/go-tpm-tools/client"
	"github.com/google/go-tpm/legacy/tpm2"
)

type TPM struct {
	dev string
}

func DefaultTPM() *TPM {
	return &TPM{
		dev: "/dev/tpm0",
	}
}

func CustomTPM(dev string) TPM {
	return TPM{
		dev: dev,
	}
}

func (t *TPM) GenerateKey(rw io.ReadWriter) (*client.Key, error) {
	exampleECCSignerTemplate := tpm2.Public{
		Type:    tpm2.AlgECC,
		NameAlg: tpm2.AlgSHA256,
		Attributes: tpm2.FlagSign | tpm2.FlagFixedTPM |
			tpm2.FlagFixedParent | tpm2.FlagSensitiveDataOrigin | tpm2.FlagUserWithAuth,
		ECCParameters: &tpm2.ECCParams{
			CurveID: tpm2.CurveNISTP256,
			Sign: &tpm2.SigScheme{
				Alg:  tpm2.AlgECDSA,
				Hash: tpm2.AlgSHA256,
			},
		},
	}

	return client.NewKey(rw, tpm2.HandleEndorsement, exampleECCSignerTemplate)
}

func (t *TPM) SignData(rw io.ReadWriter, data []byte, key *client.Key) (digest []byte, signature []byte, error error) {
	exampleECCSignerTemplate := tpm2.Public{
		Type:    tpm2.AlgECC,
		NameAlg: tpm2.AlgSHA256,
		Attributes: tpm2.FlagSign | tpm2.FlagFixedTPM |
			tpm2.FlagFixedParent | tpm2.FlagSensitiveDataOrigin | tpm2.FlagUserWithAuth,
		ECCParameters: &tpm2.ECCParams{
			CurveID: tpm2.CurveNISTP256,
			Sign: &tpm2.SigScheme{
				Alg:  tpm2.AlgECDSA,
				Hash: tpm2.AlgSHA256,
			},
		},
	}

	// logrus.Printf("trying to get cached key with handle %d", key.Handle())
	key, err := client.NewCachedKey(rw, tpm2.HandleEndorsement, exampleECCSignerTemplate, key.Handle())
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get cached key: %v", err)
	}
	// logrus.Printf("got cached key with handle %d", key.Handle())

	hash := crypto.SHA256.New()
	hash.Write(data)
	digest = hash.Sum(nil)

	cryptoSigner, err := key.GetSigner()
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create crypto signer: %v", err)
	}
	sig, err := cryptoSigner.Sign(nil, digest, crypto.SHA256)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to sign: %v", err)
	}

	return digest, sig, nil
}
