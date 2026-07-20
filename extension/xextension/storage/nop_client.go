package storage

import "context"

type nopClient struct{}

var nopClientInstance Client = &nopClient{}

func NewNopClient() Client { _ = "STUB: not implemented"; return *new(Client) }

func (c nopClient) Get(context.Context, string) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (c nopClient) Set(context.Context, string, []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (c nopClient) Delete(context.Context, string) error { _ = "STUB: not implemented"; return nil }

func (c nopClient) Close(context.Context) error { _ = "STUB: not implemented"; return nil }

func (c nopClient) Batch(context.Context, ...*Operation) error {
	_ = "STUB: not implemented"
	return nil
}
