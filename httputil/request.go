package httputil

import (
	"bytes"
	"crypto/hmac"
	"encoding/hex"
	"hash"
	"io"
	"net/http"
)

type Request struct {
	*http.Request
}

type RequestOption func(*Request) error

func WithHMAC(hashFn func() hash.Hash, header, secret string) RequestOption {
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
