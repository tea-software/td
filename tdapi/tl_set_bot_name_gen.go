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

// SetBotNameRequest represents TL type `setBotName#d295fa71`.
type SetBotNameRequest struct {
	// Identifier of the target bot
	BotUserID int64
	// A two-letter ISO 639-1 language code. If empty, the name will be shown to all users
	// for whose languages there is no dedicated name
	LanguageCode string
	// New bot's name on the specified language; 0-64 characters; must be non-empty if
	// language code is empty
	Name string
}

// SetBotNameRequestTypeID is TL type id of SetBotNameRequest.
const SetBotNameRequestTypeID = 0xd295fa71

// Ensuring interfaces in compile-time for SetBotNameRequest.
var (
	_ bin.Encoder     = &SetBotNameRequest{}
	_ bin.Decoder     = &SetBotNameRequest{}
	_ bin.BareEncoder = &SetBotNameRequest{}
	_ bin.BareDecoder = &SetBotNameRequest{}
)

func (s *SetBotNameRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.BotUserID == 0) {
		return false
	}
	if !(s.LanguageCode == "") {
		return false
	}
	if !(s.Name == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SetBotNameRequest) String() string {
	if s == nil {
		return "SetBotNameRequest(nil)"
	}
	type Alias SetBotNameRequest
	return fmt.Sprintf("SetBotNameRequest%+v", Alias(*s))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SetBotNameRequest) TypeID() uint32 {
	return SetBotNameRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*SetBotNameRequest) TypeName() string {
	return "setBotName"
}

// TypeInfo returns info about TL type.
func (s *SetBotNameRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "setBotName",
		ID:   SetBotNameRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "BotUserID",
			SchemaName: "bot_user_id",
		},
		{
			Name:       "LanguageCode",
			SchemaName: "language_code",
		},
		{
			Name:       "Name",
			SchemaName: "name",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SetBotNameRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setBotName#d295fa71 as nil")
	}
	b.PutID(SetBotNameRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *SetBotNameRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode setBotName#d295fa71 as nil")
	}
	b.PutInt53(s.BotUserID)
	b.PutString(s.LanguageCode)
	b.PutString(s.Name)
	return nil
}

// Decode implements bin.Decoder.
func (s *SetBotNameRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setBotName#d295fa71 to nil")
	}
	if err := b.ConsumeID(SetBotNameRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode setBotName#d295fa71: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *SetBotNameRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode setBotName#d295fa71 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode setBotName#d295fa71: field bot_user_id: %w", err)
		}
		s.BotUserID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode setBotName#d295fa71: field language_code: %w", err)
		}
		s.LanguageCode = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode setBotName#d295fa71: field name: %w", err)
		}
		s.Name = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (s *SetBotNameRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if s == nil {
		return fmt.Errorf("can't encode setBotName#d295fa71 as nil")
	}
	b.ObjStart()
	b.PutID("setBotName")
	b.Comma()
	b.FieldStart("bot_user_id")
	b.PutInt53(s.BotUserID)
	b.Comma()
	b.FieldStart("language_code")
	b.PutString(s.LanguageCode)
	b.Comma()
	b.FieldStart("name")
	b.PutString(s.Name)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (s *SetBotNameRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if s == nil {
		return fmt.Errorf("can't decode setBotName#d295fa71 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("setBotName"); err != nil {
				return fmt.Errorf("unable to decode setBotName#d295fa71: %w", err)
			}
		case "bot_user_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode setBotName#d295fa71: field bot_user_id: %w", err)
			}
			s.BotUserID = value
		case "language_code":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode setBotName#d295fa71: field language_code: %w", err)
			}
			s.LanguageCode = value
		case "name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode setBotName#d295fa71: field name: %w", err)
			}
			s.Name = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetBotUserID returns value of BotUserID field.
func (s *SetBotNameRequest) GetBotUserID() (value int64) {
	if s == nil {
		return
	}
	return s.BotUserID
}

// GetLanguageCode returns value of LanguageCode field.
func (s *SetBotNameRequest) GetLanguageCode() (value string) {
	if s == nil {
		return
	}
	return s.LanguageCode
}

// GetName returns value of Name field.
func (s *SetBotNameRequest) GetName() (value string) {
	if s == nil {
		return
	}
	return s.Name
}

// SetBotName invokes method setBotName#d295fa71 returning error if any.
func (c *Client) SetBotName(ctx context.Context, request *SetBotNameRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}