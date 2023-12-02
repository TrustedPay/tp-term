package tpm

import (
	"crypto/rsa"
	"fmt"

	"github.com/google/go-tpm-tools/client"
	"github.com/google/go-tpm/legacy/tpm2"
	"github.com/google/go-tpm/tpmutil"
)

type TPM struct {
	dev string
}

func DefaultTPM() TPM {
	return TPM{
		dev: "/dev/tpm0",
	}
}

func CustomTPM(dev string) TPM {
	return TPM{
		dev: dev,
	}
}

// func (t *TPM) GenerateKey() (*tpmutil.Handle, *rsa.PublicKey, error) {
func (t *TPM) GenerateKey() (*client.Key, error) {
	// Create a new TPM context
	rw, err := tpm2.OpenTPM(t.dev)
	if err != nil {
		return nil, err
	}
	defer rw.Close()

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

	// // Create a new TPM context
	// rw, err := tpm2.OpenTPM(t.dev)
	// if err != nil {
	// 	return nil, nil, err
	// }
	// defer rw.Close()

	// // Generate a new ECC key pair in the TPM
	// keyHandle, pubKey, err := tpm2.CreatePrimary(rw, tpm2.HandleOwner, tpm2.PCRSelection{}, "", "", tpm2.Public{
	// 	Type:    tpm2.AlgECC,
	// 	NameAlg: tpm2.AlgSHA256,
	// 	Attributes: tpm2.FlagSign | tpm2.FlagFixedTPM |
	// 		tpm2.FlagFixedParent | tpm2.FlagSensitiveDataOrigin | tpm2.FlagUserWithAuth,
	// 	ECCParameters: &tpm2.ECCParams{
	// 		CurveID: tpm2.CurveNISTP256,
	// 		Sign: &tpm2.SigScheme{
	// 			Alg:  tpm2.AlgECDSA,
	// 			Hash: tpm2.AlgSHA256,
	// 		},
	// 	},
	// },
	// )
	// if err != nil {
	// 	return nil, nil, err
	// }

	// // var rsaPubKey *rsa.PublicKey
	// var ecdsaPubKey *ecdsa.PublicKey
	// // Marshal the public key to an RSA public key
	// switch pubKey := pubKey.(type) {
	// // case *rsa.PublicKey:
	// // 	rsaPubKey = pubKey
	// case *ecdsa.PublicKey:
	// 	ecdsaPubKey = pubKey
	// default:
	// 	return nil, nil, fmt.Errorf("pub key is not ECDSA, of type %T", pubKey)
	// }

	// // return &keyHandle, rsaPubKey, nil
	// return &keyHandle, ecdsaPubKey, nil
}

func (t *TPM) RecallKey(keyHandle tpmutil.Handle) (*rsa.PublicKey, error) {
	// Create a new TPM context
	rw, err := tpm2.OpenTPM(t.dev)
	if err != nil {
		return nil, err
	}
	defer rw.Close()

	// Load the key from handle
	rawPubKey, _, _, err := tpm2.ReadPublic(rw, keyHandle)
	if err != nil {
		return nil, err
	}

	pubKey, err := rawPubKey.Key()
	if err != nil {
		return nil, err
	}

	var rsaPubKey *rsa.PublicKey
	// Marshal the public key to an RSA public key
	switch pubKey := pubKey.(type) {
	case *rsa.PublicKey:
		rsaPubKey = pubKey
	default:
		return nil, fmt.Errorf("pub key is not RSA, of type %T", pubKey)
	}

	return rsaPubKey, nil
}
