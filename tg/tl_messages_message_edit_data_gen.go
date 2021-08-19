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

// MessagesMessageEditData represents TL type `messages.messageEditData#26b5dde6`.
// Message edit data for media
//
// See https://core.telegram.org/constructor/messages.messageEditData for reference.
type MessagesMessageEditData struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Media caption, if the specified media's caption can be edited
	Caption bool
}

// MessagesMessageEditDataTypeID is TL type id of MessagesMessageEditData.
const MessagesMessageEditDataTypeID = 0x26b5dde6

// Ensuring interfaces in compile-time for MessagesMessageEditData.
var (
	_ bin.Encoder     = &MessagesMessageEditData{}
	_ bin.Decoder     = &MessagesMessageEditData{}
	_ bin.BareEncoder = &MessagesMessageEditData{}
	_ bin.BareDecoder = &MessagesMessageEditData{}
)

func (m *MessagesMessageEditData) Zero() bool {
	if m == nil {
		return true
	}
	if !(m.Flags.Zero()) {
		return false
	}
	if !(m.Caption == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (m *MessagesMessageEditData) String() string {
	if m == nil {
		return "MessagesMessageEditData(nil)"
	}
	type Alias MessagesMessageEditData
	return fmt.Sprintf("MessagesMessageEditData%+v", Alias(*m))
}

// FillFrom fills MessagesMessageEditData from given interface.
func (m *MessagesMessageEditData) FillFrom(from interface {
	GetCaption() (value bool)
}) {
	m.Caption = from.GetCaption()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*MessagesMessageEditData) TypeID() uint32 {
	return MessagesMessageEditDataTypeID
}

// TypeName returns name of type in TL schema.
func (*MessagesMessageEditData) TypeName() string {
	return "messages.messageEditData"
}

// TypeInfo returns info about TL type.
func (m *MessagesMessageEditData) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "messages.messageEditData",
		ID:   MessagesMessageEditDataTypeID,
	}
	if m == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Caption",
			SchemaName: "caption",
			Null:       !m.Flags.Has(0),
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (m *MessagesMessageEditData) Encode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messages.messageEditData#26b5dde6 as nil")
	}
	b.PutID(MessagesMessageEditDataTypeID)
	return m.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (m *MessagesMessageEditData) EncodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't encode messages.messageEditData#26b5dde6 as nil")
	}
	if !(m.Caption == false) {
		m.Flags.Set(0)
	}
	if err := m.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode messages.messageEditData#26b5dde6: field flags: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (m *MessagesMessageEditData) Decode(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messages.messageEditData#26b5dde6 to nil")
	}
	if err := b.ConsumeID(MessagesMessageEditDataTypeID); err != nil {
		return fmt.Errorf("unable to decode messages.messageEditData#26b5dde6: %w", err)
	}
	return m.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (m *MessagesMessageEditData) DecodeBare(b *bin.Buffer) error {
	if m == nil {
		return fmt.Errorf("can't decode messages.messageEditData#26b5dde6 to nil")
	}
	{
		if err := m.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode messages.messageEditData#26b5dde6: field flags: %w", err)
		}
	}
	m.Caption = m.Flags.Has(0)
	return nil
}

// SetCaption sets value of Caption conditional field.
func (m *MessagesMessageEditData) SetCaption(value bool) {
	if value {
		m.Flags.Set(0)
		m.Caption = true
	} else {
		m.Flags.Unset(0)
		m.Caption = false
	}
}

// GetCaption returns value of Caption conditional field.
func (m *MessagesMessageEditData) GetCaption() (value bool) {
	return m.Flags.Has(0)
}
