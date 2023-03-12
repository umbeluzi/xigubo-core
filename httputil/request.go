package httputil

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"hash"
	"io"
	"net/http"
)

type Request struct {
	*http.Request
}

type RequestOption func(*Request) error

type Signature string

const (
	signature256 = "X-Signature-256"
	signature512 = "X-Signature-512"
)

func withSignature(hashFn func() hash.Hash, header, secret string) RequestOption {
	return func(r *Request) error {
		data, err := io.ReadAll(r.Body)
		if err != nil {
			return err
		}

		hmac := hmac.New(hashFn, []byte(secret))
		if _, err := hmac.Write(data); err != nil {
			return err
		}

		r.Request.Header.Add(header, hex.EncodeToString(hmac.Sum(nil)))
		r.Body = io.NopCloser(bytes.NewReader(data))

		return nil
	}
}

func WithSignature256(secret string) RequestOption {
	return withSignature(sha256.New, signature256, secret)
}

func WithSignature512(secret string) RequestOption {
	return withSignature(sha256.New, signature512, secret)
}
