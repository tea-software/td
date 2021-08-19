// Code generated by gotdgen, DO NOT EDIT.

package e2e

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

// Long represents TL type `long#22076cba`.
//
// See https://core.telegram.org/constructor/long for reference.
type Long struct {
}

// LongTypeID is TL type id of Long.
const LongTypeID = 0x22076cba

// Ensuring interfaces in compile-time for Long.
var (
	_ bin.Encoder     = &Long{}
	_ bin.Decoder     = &Long{}
	_ bin.BareEncoder = &Long{}
	_ bin.BareDecoder = &Long{}
)

func (l *Long) Zero() bool {
	if l == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (l *Long) String() string {
	if l == nil {
		return "Long(nil)"
	}
	type Alias Long
	return fmt.Sprintf("Long%+v", Alias(*l))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*Long) TypeID() uint32 {
	return LongTypeID
}

// TypeName returns name of type in TL schema.
func (*Long) TypeName() string {
	return "long"
}

// TypeInfo returns info about TL type.
func (l *Long) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "long",
		ID:   LongTypeID,
	}
	if l == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (l *Long) Encode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode long#22076cba as nil")
	}
	b.PutID(LongTypeID)
	return l.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (l *Long) EncodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't encode long#22076cba as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (l *Long) Decode(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode long#22076cba to nil")
	}
	if err := b.ConsumeID(LongTypeID); err != nil {
		return fmt.Errorf("unable to decode long#22076cba: %w", err)
	}
	return l.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (l *Long) DecodeBare(b *bin.Buffer) error {
	if l == nil {
		return fmt.Errorf("can't decode long#22076cba to nil")
	}
	return nil
}
