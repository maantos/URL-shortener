package msgpck

import (
	"github.com/maantos/urlShortener/shortener"
	"github.com/pkg/errors"
	"github.com/vmihailenco/msgpack"
)

// Redirect implements RedirectSerializer for msgpck
type Redirect struct{}

// Decode deocde bute input into Redirect struct
func (r *Redirect) Decode(input []byte) (*shortener.Redirect, error) {
	redirect := &shortener.Redirect{}
	if err := msgpack.Unmarshal(input, redirect); err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Decode")
	}
	return redirect, nil
}

// Encode encode Redirect struct into slice of bytes
func (r *Redirect) Encode(input *shortener.Redirect) ([]byte, error) {
	rawMsg, err := msgpack.Marshal(input)
	if err != nil {
		return nil, errors.Wrap(err, "serializer.Redirect.Encode")
	}
	return rawMsg, nil
}
