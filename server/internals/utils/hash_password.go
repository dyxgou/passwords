package utils

import (
	"context"
	"crypto/sha256"
	"encoding/hex"
)

const (
	ROT    = "ROT"
	SHA256 = "SHA256"
)

type HashResults struct {
	HashType string
	Hash     string
	Err      error
}

func HashRot(ctx context.Context, password string, ch chan<- HashResults) {
	hashed := make([]byte, len(password))
	const ROT_STEPS = 2

	select {
	case <-ctx.Done():
		ch <- HashResults{ROT, "", ctx.Err()}
	default:
		for i, chr := range password {
			hashed[i] = byte(chr) + ROT_STEPS
		}

		hr := HashResults{
			HashType: ROT,
			Hash:     string(hashed),
			Err:      nil,
		}

		ch <- hr
	}
}

func HashSHA256(ctx context.Context, password string, ch chan<- HashResults) {
	hash := sha256.New()

	select {
	case <-ctx.Done():
		ch <- HashResults{SHA256, "", ctx.Err()}
		return
	default:
		_, err := hash.Write([]byte(password))
		if err != nil {
			ch <- HashResults{SHA256, "", err}
			return
		}

		ch <- HashResults{
			HashType: SHA256,
			Hash:     hex.EncodeToString(hash.Sum(nil)),
			Err:      nil,
		}
	}
}
