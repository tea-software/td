// Code generated by gotdgen, DO NOT EDIT.

package tdapi

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"go.uber.org/multierr"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdjson"
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
	_ = tdjson.Encoder{}
)

// ColorReplacement represents TL type `colorReplacement#a8c0639c`.
type ColorReplacement struct {
	// Original animated emoji color in the RGB24 format
	OldColor int32
	// Replacement animated emoji color in the RGB24 format
	NewColor int32
}

// ColorReplacementTypeID is TL type id of ColorReplacement.
const ColorReplacementTypeID = 0xa8c0639c

// Ensuring interfaces in compile-time for ColorReplacement.
var (
	_ bin.Encoder     = &ColorReplacement{}
	_ bin.Decoder     = &ColorReplacement{}
	_ bin.BareEncoder = &ColorReplacement{}
	_ bin.BareDecoder = &ColorReplacement{}
)

func (c *ColorReplacement) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.OldColor == 0) {
		return false
	}
	if !(c.NewColor == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ColorReplacement) String() string {
	if c == nil {
		return "ColorReplacement(nil)"
	}
	type Alias ColorReplacement
	return fmt.Sprintf("ColorReplacement%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ColorReplacement) TypeID() uint32 {
	return ColorReplacementTypeID
}

// TypeName returns name of type in TL schema.
func (*ColorReplacement) TypeName() string {
	return "colorReplacement"
}

// TypeInfo returns info about TL type.
func (c *ColorReplacement) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "colorReplacement",
		ID:   ColorReplacementTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "OldColor",
			SchemaName: "old_color",
		},
		{
			Name:       "NewColor",
			SchemaName: "new_color",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ColorReplacement) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode colorReplacement#a8c0639c as nil")
	}
	b.PutID(ColorReplacementTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ColorReplacement) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode colorReplacement#a8c0639c as nil")
	}
	b.PutInt32(c.OldColor)
	b.PutInt32(c.NewColor)
	return nil
}

// Decode implements bin.Decoder.
func (c *ColorReplacement) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode colorReplacement#a8c0639c to nil")
	}
	if err := b.ConsumeID(ColorReplacementTypeID); err != nil {
		return fmt.Errorf("unable to decode colorReplacement#a8c0639c: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ColorReplacement) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode colorReplacement#a8c0639c to nil")
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode colorReplacement#a8c0639c: field old_color: %w", err)
		}
		c.OldColor = value
	}
	{
		value, err := b.Int32()
		if err != nil {
			return fmt.Errorf("unable to decode colorReplacement#a8c0639c: field new_color: %w", err)
		}
		c.NewColor = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ColorReplacement) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode colorReplacement#a8c0639c as nil")
	}
	b.ObjStart()
	b.PutID("colorReplacement")
	b.FieldStart("old_color")
	b.PutInt32(c.OldColor)
	b.FieldStart("new_color")
	b.PutInt32(c.NewColor)
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ColorReplacement) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode colorReplacement#a8c0639c to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("colorReplacement"); err != nil {
				return fmt.Errorf("unable to decode colorReplacement#a8c0639c: %w", err)
			}
		case "old_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode colorReplacement#a8c0639c: field old_color: %w", err)
			}
			c.OldColor = value
		case "new_color":
			value, err := b.Int32()
			if err != nil {
				return fmt.Errorf("unable to decode colorReplacement#a8c0639c: field new_color: %w", err)
			}
			c.NewColor = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetOldColor returns value of OldColor field.
func (c *ColorReplacement) GetOldColor() (value int32) {
	return c.OldColor
}

// GetNewColor returns value of NewColor field.
func (c *ColorReplacement) GetNewColor() (value int32) {
	return c.NewColor
}