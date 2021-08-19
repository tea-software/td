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

// AccountSendVerifyPhoneCodeRequest represents TL type `account.sendVerifyPhoneCode#a5a356f9`.
// Send the verification phone code for telegram passport¹.
//
// Links:
//  1) https://core.telegram.org/passport
//
// See https://core.telegram.org/method/account.sendVerifyPhoneCode for reference.
type AccountSendVerifyPhoneCodeRequest struct {
	// The phone number to verify
	PhoneNumber string
	// Phone code settings
	Settings CodeSettings
}

// AccountSendVerifyPhoneCodeRequestTypeID is TL type id of AccountSendVerifyPhoneCodeRequest.
const AccountSendVerifyPhoneCodeRequestTypeID = 0xa5a356f9

// Ensuring interfaces in compile-time for AccountSendVerifyPhoneCodeRequest.
var (
	_ bin.Encoder     = &AccountSendVerifyPhoneCodeRequest{}
	_ bin.Decoder     = &AccountSendVerifyPhoneCodeRequest{}
	_ bin.BareEncoder = &AccountSendVerifyPhoneCodeRequest{}
	_ bin.BareDecoder = &AccountSendVerifyPhoneCodeRequest{}
)

func (s *AccountSendVerifyPhoneCodeRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.PhoneNumber == "") {
		return false
	}
	if !(s.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSendVerifyPhoneCodeRequest) String() string {
	if s == nil {
		return "AccountSendVerifyPhoneCodeRequest(nil)"
	}
	type Alias AccountSendVerifyPhoneCodeRequest
	return fmt.Sprintf("AccountSendVerifyPhoneCodeRequest%+v", Alias(*s))
}

// FillFrom fills AccountSendVerifyPhoneCodeRequest from given interface.
func (s *AccountSendVerifyPhoneCodeRequest) FillFrom(from interface {
	GetPhoneNumber() (value string)
	GetSettings() (value CodeSettings)
}) {
	s.PhoneNumber = from.GetPhoneNumber()
	s.Settings = from.GetSettings()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountSendVerifyPhoneCodeRequest) TypeID() uint32 {
	return AccountSendVerifyPhoneCodeRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountSendVerifyPhoneCodeRequest) TypeName() string {
	return "account.sendVerifyPhoneCode"
}

// TypeInfo returns info about TL type.
func (s *AccountSendVerifyPhoneCodeRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.sendVerifyPhoneCode",
		ID:   AccountSendVerifyPhoneCodeRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "PhoneNumber",
			SchemaName: "phone_number",
		},
		{
			Name:       "Settings",
			SchemaName: "settings",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *AccountSendVerifyPhoneCodeRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.sendVerifyPhoneCode#a5a356f9 as nil")
	}
	b.PutID(AccountSendVerifyPhoneCodeRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *AccountSendVerifyPhoneCodeRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.sendVerifyPhoneCode#a5a356f9 as nil")
	}
	b.PutString(s.PhoneNumber)
	if err := s.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.sendVerifyPhoneCode#a5a356f9: field settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (s *AccountSendVerifyPhoneCodeRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.sendVerifyPhoneCode#a5a356f9 to nil")
	}
	if err := b.ConsumeID(AccountSendVerifyPhoneCodeRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.sendVerifyPhoneCode#a5a356f9: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *AccountSendVerifyPhoneCodeRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.sendVerifyPhoneCode#a5a356f9 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.sendVerifyPhoneCode#a5a356f9: field phone_number: %w", err)
		}
		s.PhoneNumber = value
	}
	{
		if err := s.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.sendVerifyPhoneCode#a5a356f9: field settings: %w", err)
		}
	}
	return nil
}

// GetPhoneNumber returns value of PhoneNumber field.
func (s *AccountSendVerifyPhoneCodeRequest) GetPhoneNumber() (value string) {
	return s.PhoneNumber
}

// GetSettings returns value of Settings field.
func (s *AccountSendVerifyPhoneCodeRequest) GetSettings() (value CodeSettings) {
	return s.Settings
}

// AccountSendVerifyPhoneCode invokes method account.sendVerifyPhoneCode#a5a356f9 returning error if any.
// Send the verification phone code for telegram passport¹.
//
// Links:
//  1) https://core.telegram.org/passport
//
// See https://core.telegram.org/method/account.sendVerifyPhoneCode for reference.
func (c *Client) AccountSendVerifyPhoneCode(ctx context.Context, request *AccountSendVerifyPhoneCodeRequest) (*AuthSentCode, error) {
	var result AuthSentCode

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
