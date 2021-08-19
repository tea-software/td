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

// UsersGetFullUserRequest represents TL type `users.getFullUser#ca30a5b1`.
// Returns extended user info by ID.
//
// See https://core.telegram.org/method/users.getFullUser for reference.
type UsersGetFullUserRequest struct {
	// User ID
	ID InputUserClass
}

// UsersGetFullUserRequestTypeID is TL type id of UsersGetFullUserRequest.
const UsersGetFullUserRequestTypeID = 0xca30a5b1

// Ensuring interfaces in compile-time for UsersGetFullUserRequest.
var (
	_ bin.Encoder     = &UsersGetFullUserRequest{}
	_ bin.Decoder     = &UsersGetFullUserRequest{}
	_ bin.BareEncoder = &UsersGetFullUserRequest{}
	_ bin.BareDecoder = &UsersGetFullUserRequest{}
)

func (g *UsersGetFullUserRequest) Zero() bool {
	if g == nil {
		return true
	}
	if !(g.ID == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (g *UsersGetFullUserRequest) String() string {
	if g == nil {
		return "UsersGetFullUserRequest(nil)"
	}
	type Alias UsersGetFullUserRequest
	return fmt.Sprintf("UsersGetFullUserRequest%+v", Alias(*g))
}

// FillFrom fills UsersGetFullUserRequest from given interface.
func (g *UsersGetFullUserRequest) FillFrom(from interface {
	GetID() (value InputUserClass)
}) {
	g.ID = from.GetID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UsersGetFullUserRequest) TypeID() uint32 {
	return UsersGetFullUserRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*UsersGetFullUserRequest) TypeName() string {
	return "users.getFullUser"
}

// TypeInfo returns info about TL type.
func (g *UsersGetFullUserRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "users.getFullUser",
		ID:   UsersGetFullUserRequestTypeID,
	}
	if g == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (g *UsersGetFullUserRequest) Encode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode users.getFullUser#ca30a5b1 as nil")
	}
	b.PutID(UsersGetFullUserRequestTypeID)
	return g.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (g *UsersGetFullUserRequest) EncodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't encode users.getFullUser#ca30a5b1 as nil")
	}
	if g.ID == nil {
		return fmt.Errorf("unable to encode users.getFullUser#ca30a5b1: field id is nil")
	}
	if err := g.ID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode users.getFullUser#ca30a5b1: field id: %w", err)
	}
	return nil
}

// Decode implements bin.Decoder.
func (g *UsersGetFullUserRequest) Decode(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode users.getFullUser#ca30a5b1 to nil")
	}
	if err := b.ConsumeID(UsersGetFullUserRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode users.getFullUser#ca30a5b1: %w", err)
	}
	return g.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (g *UsersGetFullUserRequest) DecodeBare(b *bin.Buffer) error {
	if g == nil {
		return fmt.Errorf("can't decode users.getFullUser#ca30a5b1 to nil")
	}
	{
		value, err := DecodeInputUser(b)
		if err != nil {
			return fmt.Errorf("unable to decode users.getFullUser#ca30a5b1: field id: %w", err)
		}
		g.ID = value
	}
	return nil
}

// GetID returns value of ID field.
func (g *UsersGetFullUserRequest) GetID() (value InputUserClass) {
	return g.ID
}

// UsersGetFullUser invokes method users.getFullUser#ca30a5b1 returning error if any.
// Returns extended user info by ID.
//
// Possible errors:
//  400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup
//  400 MSG_ID_INVALID: Invalid message ID provided
//  400 USER_ID_INVALID: The provided user ID is invalid
//
// See https://core.telegram.org/method/users.getFullUser for reference.
// Can be used by bots.
func (c *Client) UsersGetFullUser(ctx context.Context, id InputUserClass) (*UserFull, error) {
	var result UserFull

	request := &UsersGetFullUserRequest{
		ID: id,
	}
	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return &result, nil
}
