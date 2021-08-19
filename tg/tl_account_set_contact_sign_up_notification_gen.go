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

// AccountSetContactSignUpNotificationRequest represents TL type `account.setContactSignUpNotification#cff43f61`.
// Toggle contact sign up notifications
//
// See https://core.telegram.org/method/account.setContactSignUpNotification for reference.
type AccountSetContactSignUpNotificationRequest struct {
	// Whether to disable contact sign up notifications
	Silent bool
}

// AccountSetContactSignUpNotificationRequestTypeID is TL type id of AccountSetContactSignUpNotificationRequest.
const AccountSetContactSignUpNotificationRequestTypeID = 0xcff43f61

// Ensuring interfaces in compile-time for AccountSetContactSignUpNotificationRequest.
var (
	_ bin.Encoder     = &AccountSetContactSignUpNotificationRequest{}
	_ bin.Decoder     = &AccountSetContactSignUpNotificationRequest{}
	_ bin.BareEncoder = &AccountSetContactSignUpNotificationRequest{}
	_ bin.BareDecoder = &AccountSetContactSignUpNotificationRequest{}
)

func (s *AccountSetContactSignUpNotificationRequest) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Silent == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *AccountSetContactSignUpNotificationRequest) String() string {
	if s == nil {
		return "AccountSetContactSignUpNotificationRequest(nil)"
	}
	type Alias AccountSetContactSignUpNotificationRequest
	return fmt.Sprintf("AccountSetContactSignUpNotificationRequest%+v", Alias(*s))
}

// FillFrom fills AccountSetContactSignUpNotificationRequest from given interface.
func (s *AccountSetContactSignUpNotificationRequest) FillFrom(from interface {
	GetSilent() (value bool)
}) {
	s.Silent = from.GetSilent()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountSetContactSignUpNotificationRequest) TypeID() uint32 {
	return AccountSetContactSignUpNotificationRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountSetContactSignUpNotificationRequest) TypeName() string {
	return "account.setContactSignUpNotification"
}

// TypeInfo returns info about TL type.
func (s *AccountSetContactSignUpNotificationRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.setContactSignUpNotification",
		ID:   AccountSetContactSignUpNotificationRequestTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Silent",
			SchemaName: "silent",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *AccountSetContactSignUpNotificationRequest) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.setContactSignUpNotification#cff43f61 as nil")
	}
	b.PutID(AccountSetContactSignUpNotificationRequestTypeID)
	return s.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (s *AccountSetContactSignUpNotificationRequest) EncodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode account.setContactSignUpNotification#cff43f61 as nil")
	}
	b.PutBool(s.Silent)
	return nil
}

// Decode implements bin.Decoder.
func (s *AccountSetContactSignUpNotificationRequest) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.setContactSignUpNotification#cff43f61 to nil")
	}
	if err := b.ConsumeID(AccountSetContactSignUpNotificationRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.setContactSignUpNotification#cff43f61: %w", err)
	}
	return s.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (s *AccountSetContactSignUpNotificationRequest) DecodeBare(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode account.setContactSignUpNotification#cff43f61 to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode account.setContactSignUpNotification#cff43f61: field silent: %w", err)
		}
		s.Silent = value
	}
	return nil
}

// GetSilent returns value of Silent field.
func (s *AccountSetContactSignUpNotificationRequest) GetSilent() (value bool) {
	return s.Silent
}

// AccountSetContactSignUpNotification invokes method account.setContactSignUpNotification#cff43f61 returning error if any.
// Toggle contact sign up notifications
//
// See https://core.telegram.org/method/account.setContactSignUpNotification for reference.
func (c *Client) AccountSetContactSignUpNotification(ctx context.Context, silent bool) (bool, error) {
	var result BoolBox

	request := &AccountSetContactSignUpNotificationRequest{
		Silent: silent,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
