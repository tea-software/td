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

// ChannelsAdminLogResults represents TL type `channels.adminLogResults#ed8af74d`.
// Admin log events
//
// See https://core.telegram.org/constructor/channels.adminLogResults for reference.
type ChannelsAdminLogResults struct {
	// Admin log events
	Events []ChannelAdminLogEvent
	// Chats mentioned in events
	Chats []ChatClass
	// Users mentioned in events
	Users []UserClass
}

// ChannelsAdminLogResultsTypeID is TL type id of ChannelsAdminLogResults.
const ChannelsAdminLogResultsTypeID = 0xed8af74d

// Ensuring interfaces in compile-time for ChannelsAdminLogResults.
var (
	_ bin.Encoder     = &ChannelsAdminLogResults{}
	_ bin.Decoder     = &ChannelsAdminLogResults{}
	_ bin.BareEncoder = &ChannelsAdminLogResults{}
	_ bin.BareDecoder = &ChannelsAdminLogResults{}
)

func (a *ChannelsAdminLogResults) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Events == nil) {
		return false
	}
	if !(a.Chats == nil) {
		return false
	}
	if !(a.Users == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *ChannelsAdminLogResults) String() string {
	if a == nil {
		return "ChannelsAdminLogResults(nil)"
	}
	type Alias ChannelsAdminLogResults
	return fmt.Sprintf("ChannelsAdminLogResults%+v", Alias(*a))
}

// FillFrom fills ChannelsAdminLogResults from given interface.
func (a *ChannelsAdminLogResults) FillFrom(from interface {
	GetEvents() (value []ChannelAdminLogEvent)
	GetChats() (value []ChatClass)
	GetUsers() (value []UserClass)
}) {
	a.Events = from.GetEvents()
	a.Chats = from.GetChats()
	a.Users = from.GetUsers()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsAdminLogResults) TypeID() uint32 {
	return ChannelsAdminLogResultsTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsAdminLogResults) TypeName() string {
	return "channels.adminLogResults"
}

// TypeInfo returns info about TL type.
func (a *ChannelsAdminLogResults) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.adminLogResults",
		ID:   ChannelsAdminLogResultsTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Events",
			SchemaName: "events",
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
func (a *ChannelsAdminLogResults) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode channels.adminLogResults#ed8af74d as nil")
	}
	b.PutID(ChannelsAdminLogResultsTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *ChannelsAdminLogResults) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode channels.adminLogResults#ed8af74d as nil")
	}
	b.PutVectorHeader(len(a.Events))
	for idx, v := range a.Events {
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channels.adminLogResults#ed8af74d: field events element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(a.Chats))
	for idx, v := range a.Chats {
		if v == nil {
			return fmt.Errorf("unable to encode channels.adminLogResults#ed8af74d: field chats element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channels.adminLogResults#ed8af74d: field chats element with index %d: %w", idx, err)
		}
	}
	b.PutVectorHeader(len(a.Users))
	for idx, v := range a.Users {
		if v == nil {
			return fmt.Errorf("unable to encode channels.adminLogResults#ed8af74d: field users element with index %d is nil", idx)
		}
		if err := v.Encode(b); err != nil {
			return fmt.Errorf("unable to encode channels.adminLogResults#ed8af74d: field users element with index %d: %w", idx, err)
		}
	}
	return nil
}

// Decode implements bin.Decoder.
func (a *ChannelsAdminLogResults) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode channels.adminLogResults#ed8af74d to nil")
	}
	if err := b.ConsumeID(ChannelsAdminLogResultsTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.adminLogResults#ed8af74d: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *ChannelsAdminLogResults) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode channels.adminLogResults#ed8af74d to nil")
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.adminLogResults#ed8af74d: field events: %w", err)
		}

		if headerLen > 0 {
			a.Events = make([]ChannelAdminLogEvent, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			var value ChannelAdminLogEvent
			if err := value.Decode(b); err != nil {
				return fmt.Errorf("unable to decode channels.adminLogResults#ed8af74d: field events: %w", err)
			}
			a.Events = append(a.Events, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.adminLogResults#ed8af74d: field chats: %w", err)
		}

		if headerLen > 0 {
			a.Chats = make([]ChatClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeChat(b)
			if err != nil {
				return fmt.Errorf("unable to decode channels.adminLogResults#ed8af74d: field chats: %w", err)
			}
			a.Chats = append(a.Chats, value)
		}
	}
	{
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode channels.adminLogResults#ed8af74d: field users: %w", err)
		}

		if headerLen > 0 {
			a.Users = make([]UserClass, 0, headerLen%bin.PreallocateLimit)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeUser(b)
			if err != nil {
				return fmt.Errorf("unable to decode channels.adminLogResults#ed8af74d: field users: %w", err)
			}
			a.Users = append(a.Users, value)
		}
	}
	return nil
}

// GetEvents returns value of Events field.
func (a *ChannelsAdminLogResults) GetEvents() (value []ChannelAdminLogEvent) {
	return a.Events
}

// GetChats returns value of Chats field.
func (a *ChannelsAdminLogResults) GetChats() (value []ChatClass) {
	return a.Chats
}

// GetUsers returns value of Users field.
func (a *ChannelsAdminLogResults) GetUsers() (value []UserClass) {
	return a.Users
}

// MapChats returns field Chats wrapped in ChatClassArray helper.
func (a *ChannelsAdminLogResults) MapChats() (value ChatClassArray) {
	return ChatClassArray(a.Chats)
}

// MapUsers returns field Users wrapped in UserClassArray helper.
func (a *ChannelsAdminLogResults) MapUsers() (value UserClassArray) {
	return UserClassArray(a.Users)
}
