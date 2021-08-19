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

// ChannelsChannelParticipant represents TL type `channels.channelParticipant#dfb80317`.
// Represents a channel participant
//
// See https://core.telegram.org/constructor/channels.channelParticipant for reference.
type ChannelsChannelParticipant struct {
	// The channel participant
	Participant ChannelParticipantClass
	// Chats field of ChannelsChannelParticipant.
	Chats []ChatClass
	// Users
	Users []UserClass
}

// ChannelsChannelParticipantTypeID is TL type id of ChannelsChannelParticipant.
const ChannelsChannelParticipantTypeID = 0xdfb80317

// Ensuring interfaces in compile-time for ChannelsChannelParticipant.
var (
	_ bin.Encoder     = &ChannelsChannelParticipant{}
	_ bin.Decoder     = &ChannelsChannelParticipant{}
	_ bin.BareEncoder = &ChannelsChannelParticipant{}
	_ bin.BareDecoder = &ChannelsChannelParticipant{}
)

func (c *ChannelsChannelParticipant) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.Participant == nil) {
		return false
	}
	if !(c.Chats == nil) {
		return false
	}
	if !(c.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChannelsChannelParticipant) String() string {
	if c == nil {
		return "ChannelsChannelParticipant(nil)"
	}
	type Alias ChannelsChannelParticipant
	return fmt.Sprintf("ChannelsChannelParticipant%+v", Alias(*c))
}

// FillFrom fills ChannelsChannelParticipant from given interface.
func (c *ChannelsChannelParticipant) FillFrom(from interface {
	GetParticipant() (value ChannelParticipantClass)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
}) {
	c.Participant = from.GetParticipant()
	c.Chats = from.GetChats()
	c.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsChannelParticipant) TypeID() uint32 {
	return ChannelsChannelParticipantTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsChannelParticipant) TypeName() string {
	return "channels.channelParticipant"
}

// TypeInfo returns info about TL type.
func (c *ChannelsChannelParticipant) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.channelParticipant",
		ID:   ChannelsChannelParticipantTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Participant",
			SchemaName: "participant",
		},
		{
			Name:       "Chats",
			SchemaName: "chats",
		},
		{
			Name:       "Users",
			SchemaName: "users",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChannelsChannelParticipant) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channels.channelParticipant#dfb80317 as nil")
	}
	b.PutID(ChannelsChannelParticipantTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChannelsChannelParticipant) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode channels.channelParticipant#dfb80317 as nil")
	}
	if c.Participant == nil {
		return fmt.Errorf("unable to encode channels.channelParticipant#dfb80317: field participant is nil")
	}
	if err := c.Participant.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.channelParticipant#dfb80317: field participant: %w", err)
	}
	b.PutVectorHeader(len(c.Chats))
	for idx, v := range c.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode channels.channelParticipant#dfb80317: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channels.channelParticipant#dfb80317: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(c.Users))
	for idx, v := range c.Users {
		if v == nil {
			return fmt.Errorf("unable to encode channels.channelParticipant#dfb80317: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channels.channelParticipant#dfb80317: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (c *ChannelsChannelParticipant) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channels.channelParticipant#dfb80317 to nil")
	}
	if err := b.ConsumeID(ChannelsChannelParticipantTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.channelParticipant#dfb80317: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChannelsChannelParticipant) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode channels.channelParticipant#dfb80317 to nil")
	}
	{
		value, err := DecodeChannelParticipant(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.channelParticipant#dfb80317: field participant: %w", err)
		}
		c.Participant = value
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.channelParticipant#dfb80317: field chats: %w", err)
		}

		if headerLen > 0 {
			c.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode channels.channelParticipant#dfb80317: field chats: %w", err)
			}
			c.Chats = append(c.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.channelParticipant#dfb80317: field users: %w", err)
		}

		if headerLen > 0 {
			c.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode channels.channelParticipant#dfb80317: field users: %w", err)
			}
			c.Users = append(c.Users, value)
		}
	}
	return nil
}

// GetParticipant returns value of Participant field.
func (c *ChannelsChannelParticipant) GetParticipant() (value ChannelParticipantClass) {
	return c.Participant
}

// GetChats returns value of Chats field.
func (c *ChannelsChannelParticipant) GetChats() (value []ChatClass) {
	return c.Chats
}

// GetUsers returns value of Users field.
func (c *ChannelsChannelParticipant) GetUsers() (value []UserClass) {
	return c.Users
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (c *ChannelsChannelParticipant) MapChats() (value ChatClassArray) {
	return ChatClassArray(c.Chats)
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (c *ChannelsChannelParticipant) MapUsers() (value UserClassArray) {
	return UserClassArray(c.Users)
}
