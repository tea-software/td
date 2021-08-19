// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
	"github.com/gotd/td/tgerr"
)

// No-op definition for keeping imports.
var (
	_ = bin.Buffer{}
	_ = context.Background()
	_ = fmt.Stringer(nil)
	_ = strings.Builder{}
	_ = errors.Is
	_ = multierr.AppendInto
	_ = sort.Ints
	_ = tdp.Format
	_ = tgerr.Error{}
)

// TestUseConfigSimpleRequest represents TL type `test.useConfigSimple#f9b7b23d`.
//
// See https://core.telegram.org/method/test.useConfigSimple for reference.
type TestUseConfigSimpleRequest struct {
}

// TestUseConfigSimpleRequestTypeID is TL type id of TestUseConfigSimpleRequest.
const TestUseConfigSimpleRequestTypeID = 0xf9b7b23d

// Ensuring interfaces in compile-time for TestUseConfigSimpleRequest.
var (
	_ bin.Encoder     = &TestUseConfigSimpleRequest{}
	_ bin.Decoder     = &TestUseConfigSimpleRequest{}
	_ bin.BareEncoder = &TestUseConfigSimpleRequest{}
	_ bin.BareDecoder = &TestUseConfigSimpleRequest{}
)

func (u *TestUseConfigSimpleRequest) Zero() bool {
	if u == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (u *TestUseConfigSimpleRequest) String() string {
	if u == nil {
		return "TestUseConfigSimpleRequest(nil)"
	}
	type Alias TestUseConfigSimpleRequest
	return fmt.Sprintf("TestUseConfigSimpleRequest%+v", Alias(*u))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*TestUseConfigSimpleRequest) TypeID() uint32 {
	return TestUseConfigSimpleRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*TestUseConfigSimpleRequest) TypeName() string {
	return "test.useConfigSimple"
}

// TypeInfo returns info about TL type.
func (u *TestUseConfigSimpleRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "test.useConfigSimple",
		ID:   TestUseConfigSimpleRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (u *TestUseConfigSimpleRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode test.useConfigSimple#f9b7b23d as nil")
	}
	b.PutID(TestUseConfigSimpleRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *TestUseConfigSimpleRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode test.useConfigSimple#f9b7b23d as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *TestUseConfigSimpleRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode test.useConfigSimple#f9b7b23d to nil")
	}
	if err := b.ConsumeID(TestUseConfigSimpleRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode test.useConfigSimple#f9b7b23d: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *TestUseConfigSimpleRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode test.useConfigSimple#f9b7b23d to nil")
	}
	return nil
}

// TestUseConfigSimple invokes method test.useConfigSimple#f9b7b23d returning error if any.
//
// See https://core.telegram.org/method/test.useConfigSimple for reference.
func (c *Client) TestUseConfigSimple(ctx context.Context) (*HelpConfigSimple, error) {
	var result HelpConfigSimple

	request := &TestUseConfigSimpleRequest{}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
