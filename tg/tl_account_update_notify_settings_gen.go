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

// AccountUpdateNotifySettingsRequest represents TL type `account.updateNotifySettings#84be5b93`.
// Edits notification settings from a given user/group, from all users/all groups.
//
// See https://core.telegram.org/method/account.updateNotifySettings for reference.
type AccountUpdateNotifySettingsRequest struct {
	// Notification source
	Peer InputNotifyPeerClass
	// Notification settings
	Settings InputPeerNotifySettings
}

// AccountUpdateNotifySettingsRequestTypeID is TL type id of AccountUpdateNotifySettingsRequest.
const AccountUpdateNotifySettingsRequestTypeID = 0x84be5b93

// Ensuring interfaces in compile-time for AccountUpdateNotifySettingsRequest.
var (
	_ bin.Encoder     = &AccountUpdateNotifySettingsRequest{}
	_ bin.Decoder     = &AccountUpdateNotifySettingsRequest{}
	_ bin.BareEncoder = &AccountUpdateNotifySettingsRequest{}
	_ bin.BareDecoder = &AccountUpdateNotifySettingsRequest{}
)

func (u *AccountUpdateNotifySettingsRequest) Zero() bool {
	if u == nil {
		return true
	}
	if !(u.Peer == nil) {
		return false
	}
	if !(u.Settings.Zero()) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (u *AccountUpdateNotifySettingsRequest) String() string {
	if u == nil {
		return "AccountUpdateNotifySettingsRequest(nil)"
	}
	type Alias AccountUpdateNotifySettingsRequest
	return fmt.Sprintf("AccountUpdateNotifySettingsRequest%+v", Alias(*u))
}

// FillFrom fills AccountUpdateNotifySettingsRequest from given interface.
func (u *AccountUpdateNotifySettingsRequest) FillFrom(from interface {
	GetPeer() (value InputNotifyPeerClass)
	GetSettings() (value InputPeerNotifySettings)
}) {
	u.Peer = from.GetPeer()
	u.Settings = from.GetSettings()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*AccountUpdateNotifySettingsRequest) TypeID() uint32 {
	return AccountUpdateNotifySettingsRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*AccountUpdateNotifySettingsRequest) TypeName() string {
	return "account.updateNotifySettings"
}

// TypeInfo returns info about TL type.
func (u *AccountUpdateNotifySettingsRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "account.updateNotifySettings",
		ID:   AccountUpdateNotifySettingsRequestTypeID,
	}
	if u == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Peer",
			SchemaName: "peer",
		},
		{
			Name:       "Settings",
			SchemaName: "settings",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (u *AccountUpdateNotifySettingsRequest) Encode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateNotifySettings#84be5b93 as nil")
	}
	b.PutID(AccountUpdateNotifySettingsRequestTypeID)
	return u.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (u *AccountUpdateNotifySettingsRequest) EncodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't encode account.updateNotifySettings#84be5b93 as nil")
	}
	if u.Peer == nil {
		return fmt.Errorf("unable to encode account.updateNotifySettings#84be5b93: field peer is nil")
	}
	if err := u.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.updateNotifySettings#84be5b93: field peer: %w", err)
	}
	if err := u.Settings.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.updateNotifySettings#84be5b93: field settings: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (u *AccountUpdateNotifySettingsRequest) Decode(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateNotifySettings#84be5b93 to nil")
	}
	if err := b.ConsumeID(AccountUpdateNotifySettingsRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.updateNotifySettings#84be5b93: %w", err)
	}
	return u.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (u *AccountUpdateNotifySettingsRequest) DecodeBare(b *bin.Buffer) error {
	if u == nil {
		return fmt.Errorf("can't decode account.updateNotifySettings#84be5b93 to nil")
	}
	{
		value, err := DecodeInputNotifyPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.updateNotifySettings#84be5b93: field peer: %w", err)
		}
		u.Peer = value
	}
	{
		if err := u.Settings.Decode(b); err != nil {
			return fmt.Errorf("unable to decode account.updateNotifySettings#84be5b93: field settings: %w", err)
		}
	}
	return nil
}

// GetPeer returns value of Peer field.
func (u *AccountUpdateNotifySettingsRequest) GetPeer() (value InputNotifyPeerClass) {
	return u.Peer
}

// GetSettings returns value of Settings field.
func (u *AccountUpdateNotifySettingsRequest) GetSettings() (value InputPeerNotifySettings) {
	return u.Settings
}

// AccountUpdateNotifySettings invokes method account.updateNotifySettings#84be5b93 returning error if any.
// Edits notification settings from a given user/group, from all users/all groups.
//
// Possible errors:
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 MSG_ID_INVALID: Invalid message ID provided
//  400 PEER_ID_INVALID: The provided peer id is invalid
//  400 SETTINGS_INVALID: Invalid settings were provided
//
// See https://core.telegram.org/method/account.updateNotifySettings for reference.
func (c *Client) AccountUpdateNotifySettings(ctx context.Context, request *AccountUpdateNotifySettingsRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}
