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

// ChannelsDeleteHistoryRequest represents TL type `channels.deleteHistory#9baa9647`.
// Delete the history of a supergroup¹
//
// Links:
//  1. https://core.telegram.org/api/channel
//
// See https://core.telegram.org/method/channels.deleteHistory for reference.
type ChannelsDeleteHistoryRequest struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Whether the history should be deleted for everyone
	ForEveryone bool
	// Supergroup¹ whose history must be deleted
	//
	// Links:
	//  1) https://core.telegram.org/api/channel
	Channel InputChannelClass
	// ID of message up to which the history must be deleted
	MaxID int
}

// ChannelsDeleteHistoryRequestTypeID is TL type id of ChannelsDeleteHistoryRequest.
const ChannelsDeleteHistoryRequestTypeID = 0x9baa9647

// Ensuring interfaces in compile-time for ChannelsDeleteHistoryRequest.
var (
	_ bin.Encoder     = &ChannelsDeleteHistoryRequest{}
	_ bin.Decoder     = &ChannelsDeleteHistoryRequest{}
	_ bin.BareEncoder = &ChannelsDeleteHistoryRequest{}
	_ bin.BareDecoder = &ChannelsDeleteHistoryRequest{}
)

func (d *ChannelsDeleteHistoryRequest) Zero() bool {
	if d == nil {
		return true
	}
	if !(d.Flags.Zero()) {
		return false
	}
	if !(d.ForEveryone == false) {
		return false
	}
	if !(d.Channel == nil) {
		return false
	}
	if !(d.MaxID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (d *ChannelsDeleteHistoryRequest) String() string {
	if d == nil {
		return "ChannelsDeleteHistoryRequest(nil)"
	}
	type Alias ChannelsDeleteHistoryRequest
	return fmt.Sprintf("ChannelsDeleteHistoryRequest%+v", Alias(*d))
}

// FillFrom fills ChannelsDeleteHistoryRequest from given interface.
func (d *ChannelsDeleteHistoryRequest) FillFrom(from interface {
	GetForEveryone() (value bool)
	GetChannel() (value InputChannelClass)
	GetMaxID() (value int)
}) {
	d.ForEveryone = from.GetForEveryone()
	d.Channel = from.GetChannel()
	d.MaxID = from.GetMaxID()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChannelsDeleteHistoryRequest) TypeID() uint32 {
	return ChannelsDeleteHistoryRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*ChannelsDeleteHistoryRequest) TypeName() string {
	return "channels.deleteHistory"
}

// TypeInfo returns info about TL type.
func (d *ChannelsDeleteHistoryRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "channels.deleteHistory",
		ID:   ChannelsDeleteHistoryRequestTypeID,
	}
	if d == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ForEveryone",
			SchemaName: "for_everyone",
			Null:       !d.Flags.Has(0),
		},
		{
			Name:       "Channel",
			SchemaName: "channel",
		},
		{
			Name:       "MaxID",
			SchemaName: "max_id",
		},
	}
	return typ
}

// SetFlags sets flags for non-zero fields.
func (d *ChannelsDeleteHistoryRequest) SetFlags() {
	if !(d.ForEveryone == false) {
		d.Flags.Set(0)
	}
}

// Encode implements bin.Encoder.
func (d *ChannelsDeleteHistoryRequest) Encode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode channels.deleteHistory#9baa9647 as nil")
	}
	b.PutID(ChannelsDeleteHistoryRequestTypeID)
	return d.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (d *ChannelsDeleteHistoryRequest) EncodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't encode channels.deleteHistory#9baa9647 as nil")
	}
	d.SetFlags()
	if err := d.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.deleteHistory#9baa9647: field flags: %w", err)
	}
	if d.Channel == nil {
		return fmt.Errorf("unable to encode channels.deleteHistory#9baa9647: field channel is nil")
	}
	if err := d.Channel.Encode(b); err != nil {
		return fmt.Errorf("unable to encode channels.deleteHistory#9baa9647: field channel: %w", err)
	}
	b.PutInt(d.MaxID)
	return nil
}

// Decode implements bin.Decoder.
func (d *ChannelsDeleteHistoryRequest) Decode(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode channels.deleteHistory#9baa9647 to nil")
	}
	if err := b.ConsumeID(ChannelsDeleteHistoryRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode channels.deleteHistory#9baa9647: %w", err)
	}
	return d.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (d *ChannelsDeleteHistoryRequest) DecodeBare(b *bin.Buffer) error {
	if d == nil {
		return fmt.Errorf("can't decode channels.deleteHistory#9baa9647 to nil")
	}
	{
		if err := d.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode channels.deleteHistory#9baa9647: field flags: %w", err)
		}
	}
	d.ForEveryone = d.Flags.Has(0)
	{
		value, err := DecodeInputChannel(b)
		if err != nil {
			return fmt.Errorf("unable to decode channels.deleteHistory#9baa9647: field channel: %w", err)
		}
		d.Channel = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode channels.deleteHistory#9baa9647: field max_id: %w", err)
		}
		d.MaxID = value
	}
	return nil
}

// SetForEveryone sets value of ForEveryone conditional field.
func (d *ChannelsDeleteHistoryRequest) SetForEveryone(value bool) {
	if value {
		d.Flags.Set(0)
		d.ForEveryone = true
	} else {
		d.Flags.Unset(0)
		d.ForEveryone = false
	}
}

// GetForEveryone returns value of ForEveryone conditional field.
func (d *ChannelsDeleteHistoryRequest) GetForEveryone() (value bool) {
	if d == nil {
		return
	}
	return d.Flags.Has(0)
}

// GetChannel returns value of Channel field.
func (d *ChannelsDeleteHistoryRequest) GetChannel() (value InputChannelClass) {
	if d == nil {
		return
	}
	return d.Channel
}

// GetMaxID returns value of MaxID field.
func (d *ChannelsDeleteHistoryRequest) GetMaxID() (value int) {
	if d == nil {
		return
	}
	return d.MaxID
}

// GetChannelAsNotEmpty returns mapped value of Channel field.
func (d *ChannelsDeleteHistoryRequest) GetChannelAsNotEmpty() (NotEmptyInputChannel, bool) {
	return d.Channel.AsNotEmpty()
}

// ChannelsDeleteHistory invokes method channels.deleteHistory#9baa9647 returning error if any.
// Delete the history of a supergroup¹
//
// Links:
//  1. https://core.telegram.org/api/channel
//
// Possible errors:
//
//	400 CHANNEL_INVALID: The provided channel is invalid.
//	400 CHANNEL_PARICIPANT_MISSING: The current user is not in the channel.
//	400 CHANNEL_PRIVATE: You haven't joined this channel/supergroup.
//	400 CHAT_ADMIN_REQUIRED: You must be an admin in this chat to do this.
//
// See https://core.telegram.org/method/channels.deleteHistory for reference.
func (c *Client) ChannelsDeleteHistory(ctx context.Context, request *ChannelsDeleteHistoryRequest) (UpdatesClass, error) {
	var result UpdatesBox

	if err := c.rpc.Invoke(ctx, request, &result); err != nil {
		return nil, err
	}
	return result.Updates, nil
}
