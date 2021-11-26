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

// SetCustomLanguagePackRequest represents TL type `setCustomLanguagePack#234b0607`.
type SetCustomLanguagePackRequest struct {
	// Information about the language pack. Language pack ID must start with 'X', consist
	// only of English letters, digits and hyphens, and must not exceed 64 characters. Can be
	// called before authorization
	Info LanguagePackInfo
	// Strings of the new language pack
	Strings []LanguagePackString
}

// SetCustomLanguagePackRequestTypeID is TL type id of SetCustomLanguagePackRequest.
const SetCustomLanguagePackRequestTypeID = 0x234b0607

// Ensuring interfaces in compile-time for SetCustomLanguagePackRequest.
var (
	_ bin.Encoder     = &SetCustomLanguagePackRequest{}
	_ bin.Decoder     = &SetCustomLanguagePackRequest{}
	_ bin.BareEncoder = &SetCustomLanguagePackRequest{}
	_ bin.BareDecoder = &SetCustomLanguagePackRequest{}
)

func (s *SetCustomLanguagePackRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Info.Zero()) {
		return false
	}
	if !(s.Strings == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetCustomLanguagePackRequest) String() string {
	if s == nil {
		return "SetCustomLanguagePackRequest(nil)"
	}
	type Alias SetCustomLanguagePackRequest
	return fmt.Sprintf("SetCustomLanguagePackRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetCustomLanguagePackRequest) TypeID() uint32 {
	return SetCustomLanguagePackRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetCustomLanguagePackRequest) TypeName() string {
	return "setCustomLanguagePack"
}

// TypeInfo returns info about TL type.
func (s *SetCustomLanguagePackRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setCustomLanguagePack",
		ID:   SetCustomLanguagePackRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Info",
			SchemaName: "info",
		},
		{
			Name:       "Strings",
			SchemaName: "strings",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetCustomLanguagePackRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setCustomLanguagePack#234b0607 as nil")
	}
	b.PutID(SetCustomLanguagePackRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetCustomLanguagePackRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setCustomLanguagePack#234b0607 as nil")
	}
	if err := s.Info.Encode(b); err != nil {
		return fmt.Errorf("unable to encode setCustomLanguagePack#234b0607: field info: %w", err)
	}
	b.PutInt(len(s.Strings))
	for idx, v := range s.Strings {
		if err := v.EncodeBare(b); err != nil {
			return fmt.Errorf("unable to encode bare setCustomLanguagePack#234b0607: field strings element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *SetCustomLanguagePackRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setCustomLanguagePack#234b0607 to nil")
	}
	if err := b.ConsumeID(SetCustomLanguagePackRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setCustomLanguagePack#234b0607: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetCustomLanguagePackRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setCustomLanguagePack#234b0607 to nil")
	}
	{
		if err := s.Info.Decode(b); err != nil {
			return fmt.Errorf("unable to decode setCustomLanguagePack#234b0607: field info: %w", err)
		}
	}
	{
		headerLen, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode setCustomLanguagePack#234b0607: field strings: %w", err)
		}

		if headerLen > 0 {
			s.Strings = make([]LanguagePackString, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value LanguagePackString
			if err := value.DecodeBare(b); err != nil {
				return fmt.Errorf("unable to decode bare setCustomLanguagePack#234b0607: field strings: %w", err)
			}
			s.Strings = append(s.Strings, value)
		}
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetCustomLanguagePackRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setCustomLanguagePack#234b0607 as nil")
	}
	b.ObjStart()
	b.PutID("setCustomLanguagePack")
	b.FieldStart("info")
	if err := s.Info.EncodeTDLibJSON(b); err != nil {
		return fmt.Errorf("unable to encode setCustomLanguagePack#234b0607: field info: %w", err)
	}
	b.FieldStart("strings")
	b.ArrStart()
	for idx, v := range s.Strings {
		if err := v.EncodeTDLibJSON(b); err != nil {
			return fmt.Errorf("unable to encode setCustomLanguagePack#234b0607: field strings element with index %d: %w", idx, err)
		}
	}
	b.ArrEnd()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetCustomLanguagePackRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setCustomLanguagePack#234b0607 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setCustomLanguagePack"); err != nil {
				return fmt.Errorf("unable to decode setCustomLanguagePack#234b0607: %w", err)
			}
		case "info":
			if err := s.Info.DecodeTDLibJSON(b); err != nil {
				return fmt.Errorf("unable to decode setCustomLanguagePack#234b0607: field info: %w", err)
			}
		case "strings":
			if err := b.Arr(func(b tdjson.Decoder) error {
				var value LanguagePackString
				if err := value.DecodeTDLibJSON(b); err != nil {
					return fmt.Errorf("unable to decode setCustomLanguagePack#234b0607: field strings: %w", err)
				}
				s.Strings = append(s.Strings, value)
				return nil
			}); err != nil {
				return fmt.Errorf("unable to decode setCustomLanguagePack#234b0607: field strings: %w", err)
			}
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetInfo returns value of Info field.
func (s *SetCustomLanguagePackRequest) GetInfo() (value LanguagePackInfo) {
	return s.Info
}

// GetStrings returns value of Strings field.
func (s *SetCustomLanguagePackRequest) GetStrings() (value []LanguagePackString) {
	return s.Strings
}

// SetCustomLanguagePack invokes method setCustomLanguagePack#234b0607 returning error if any.
func (c *Client) SetCustomLanguagePack(ctx context.Context, request *SetCustomLanguagePackRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}