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

// PopularContact represents TL type `popularContact#5ce14175`.
// Popular contact
//
// See https://core.telegram.org/constructor/popularContact for reference.
type PopularContact struct {
	// Contact identifier
	ClientID int64
	// How many people imported this contact
	Importers int
}

// PopularContactTypeID is TL type id of PopularContact.
const PopularContactTypeID = 0x5ce14175

// Ensuring interfaces in compile-time for PopularContact.
var (
	_ bin.Encoder     = &PopularContact{}
	_ bin.Decoder     = &PopularContact{}
	_ bin.BareEncoder = &PopularContact{}
	_ bin.BareDecoder = &PopularContact{}
)

func (p *PopularContact) Zero() bool {
	if p == nil {
		return true
	}
	if !(p.ClientID == 0) {
		return false
	}
	if !(p.Importers == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (p *PopularContact) String() string {
	if p == nil {
		return "PopularContact(nil)"
	}
	type Alias PopularContact
	return fmt.Sprintf("PopularContact%+v", Alias(*p))
}

// FillFrom fills PopularContact from given interface.
func (p *PopularContact) FillFrom(from interface {
	GetClientID() (value int64)
	GetImporters() (value int)
}) {
	p.ClientID = from.GetClientID()
	p.Importers = from.GetImporters()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*PopularContact) TypeID() uint32 {
	return PopularContactTypeID
}

// TypeName returns name of type in TL schema.
func (*PopularContact) TypeName() string {
	return "popularContact"
}

// TypeInfo returns info about TL type.
func (p *PopularContact) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "popularContact",
		ID:   PopularContactTypeID,
	}
	if p == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ClientID",
			SchemaName: "client_id",
		},
		{
			Name:       "Importers",
			SchemaName: "importers",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (p *PopularContact) Encode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode popularContact#5ce14175 as nil")
	}
	b.PutID(PopularContactTypeID)
	return p.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (p *PopularContact) EncodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't encode popularContact#5ce14175 as nil")
	}
	b.PutLong(p.ClientID)
	b.PutInt(p.Importers)
	return nil
}

// Decode implements bin.Decoder.
func (p *PopularContact) Decode(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode popularContact#5ce14175 to nil")
	}
	if err := b.ConsumeID(PopularContactTypeID); err != nil {
		return fmt.Errorf("unable to decode popularContact#5ce14175: %w", err)
	}
	return p.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (p *PopularContact) DecodeBare(b *bin.Buffer) error {
	if p == nil {
		return fmt.Errorf("can't decode popularContact#5ce14175 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode popularContact#5ce14175: field client_id: %w", err)
		}
		p.ClientID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode popularContact#5ce14175: field importers: %w", err)
		}
		p.Importers = value
	}
	return nil
}

// GetClientID returns value of ClientID field.
func (p *PopularContact) GetClientID() (value int64) {
	return p.ClientID
}

// GetImporters returns value of Importers field.
func (p *PopularContact) GetImporters() (value int) {
	return p.Importers
}
