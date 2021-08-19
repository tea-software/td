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

// LangpackGetLangPackRequest represents TL type `langpack.getLangPack#f2f2330a`.
// Get localization pack strings
//
// See https://core.telegram.org/method/langpack.getLangPack for reference.
type LangpackGetLangPackRequest struct {
	// Language pack name
	LangPack string
	// Language code
	LangCode string
}

// LangpackGetLangPackRequestTypeID is TL type id of LangpackGetLangPackRequest.
const LangpackGetLangPackRequestTypeID = 0xf2f2330a

// Ensuring interfaces in compile-time for LangpackGetLangPackRequest.
var (
	_ bin.Encoder     = &LangpackGetLangPackRequest{}
	_ bin.Decoder     = &LangpackGetLangPackRequest{}
	_ bin.BareEncoder = &LangpackGetLangPackRequest{}
	_ bin.BareDecoder = &LangpackGetLangPackRequest{}
)

func (g *LangpackGetLangPackRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.LangPack == "") {
		return false
	}
	if !(g.LangCode == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *LangpackGetLangPackRequest) String() string {
	if g == nil {
		return "LangpackGetLangPackRequest(nil)"
	}
	type Alias LangpackGetLangPackRequest
	return fmt.Sprintf("LangpackGetLangPackRequest%+v", Alias(*g))
}

// FillFrom fills LangpackGetLangPackRequest from given interface.
func (g *LangpackGetLangPackRequest) FillFrom(from interface {
	GetLangPack() (value string)
	GetLangCode() (value string)
}) {
	g.LangPack = from.GetLangPack()
	g.LangCode = from.GetLangCode()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*LangpackGetLangPackRequest) TypeID() uint32 {
	return LangpackGetLangPackRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*LangpackGetLangPackRequest) TypeName() string {
	return "langpack.getLangPack"
}

// TypeInfo returns info about TL type.
func (g *LangpackGetLangPackRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "langpack.getLangPack",
		ID:   LangpackGetLangPackRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "LangPack",
			SchemaName: "lang_pack",
		},
		{
			Name:       "LangCode",
			SchemaName: "lang_code",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *LangpackGetLangPackRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode langpack.getLangPack#f2f2330a as nil")
	}
	b.PutID(LangpackGetLangPackRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *LangpackGetLangPackRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode langpack.getLangPack#f2f2330a as nil")
	}
	b.PutString(g.LangPack)
	b.PutString(g.LangCode)
	return nil
}

// Decode implements bin.Decoder.
func (g *LangpackGetLangPackRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode langpack.getLangPack#f2f2330a to nil")
	}
	if err := b.ConsumeID(LangpackGetLangPackRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode langpack.getLangPack#f2f2330a: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *LangpackGetLangPackRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode langpack.getLangPack#f2f2330a to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langpack.getLangPack#f2f2330a: field lang_pack: %w", err)
		}
		g.LangPack = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode langpack.getLangPack#f2f2330a: field lang_code: %w", err)
		}
		g.LangCode = value
	}
	return nil
}

// GetLangPack returns value of LangPack field.
func (g *LangpackGetLangPackRequest) GetLangPack() (value string) {
	return g.LangPack
}

// GetLangCode returns value of LangCode field.
func (g *LangpackGetLangPackRequest) GetLangCode() (value string) {
	return g.LangCode
}

// LangpackGetLangPack invokes method langpack.getLangPack#f2f2330a returning error if any.
// Get localization pack strings
//
// Possible errors:
//  400 LANG_PACK_INVALID: The provided language pack is invalid
//
// See https://core.telegram.org/method/langpack.getLangPack for reference.
func (c *Client) LangpackGetLangPack(ctx context.Context, request *LangpackGetLangPackRequest) (*LangPackDifference, error) {
	var result LangPackDifference

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
