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

// EditForumTopicRequest represents TL type `editForumTopic#f2c261b0`.
type EditForumTopicRequest struct {
	// Identifier of the chat
	ChatID int64
	// Message thread identifier of the forum topic
	MessageThreadID int64
	// New name of the topic; 1-128 characters
	Name string
	// Identifier of the new custom emoji for topic icon. Telegram Premium users can use any
	// custom emoji, other users can use only a custom emoji returned by
	// getForumTopicDefaultIcons
	IconCustomEmojiID int64
}

// EditForumTopicRequestTypeID is TL type id of EditForumTopicRequest.
const EditForumTopicRequestTypeID = 0xf2c261b0

// Ensuring interfaces in compile-time for EditForumTopicRequest.
var (
	_ bin.Encoder     = &EditForumTopicRequest{}
	_ bin.Decoder     = &EditForumTopicRequest{}
	_ bin.BareEncoder = &EditForumTopicRequest{}
	_ bin.BareDecoder = &EditForumTopicRequest{}
)

func (e *EditForumTopicRequest) Zero() bool {
	if e == nil {
		return true
	}
	if !(e.ChatID == 0) {
		return false
	}
	if !(e.MessageThreadID == 0) {
		return false
	}
	if !(e.Name == "") {
		return false
	}
	if !(e.IconCustomEmojiID == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (e *EditForumTopicRequest) String() string {
	if e == nil {
		return "EditForumTopicRequest(nil)"
	}
	type Alias EditForumTopicRequest
	return fmt.Sprintf("EditForumTopicRequest%+v", Alias(*e))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*EditForumTopicRequest) TypeID() uint32 {
	return EditForumTopicRequestTypeID
}

// TypeName returns name of type in TL schema.
func (*EditForumTopicRequest) TypeName() string {
	return "editForumTopic"
}

// TypeInfo returns info about TL type.
func (e *EditForumTopicRequest) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "editForumTopic",
		ID:   EditForumTopicRequestTypeID,
	}
	if e == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ChatID",
			SchemaName: "chat_id",
		},
		{
			Name:       "MessageThreadID",
			SchemaName: "message_thread_id",
		},
		{
			Name:       "Name",
			SchemaName: "name",
		},
		{
			Name:       "IconCustomEmojiID",
			SchemaName: "icon_custom_emoji_id",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (e *EditForumTopicRequest) Encode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editForumTopic#f2c261b0 as nil")
	}
	b.PutID(EditForumTopicRequestTypeID)
	return e.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (e *EditForumTopicRequest) EncodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't encode editForumTopic#f2c261b0 as nil")
	}
	b.PutInt53(e.ChatID)
	b.PutInt53(e.MessageThreadID)
	b.PutString(e.Name)
	b.PutLong(e.IconCustomEmojiID)
	return nil
}

// Decode implements bin.Decoder.
func (e *EditForumTopicRequest) Decode(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editForumTopic#f2c261b0 to nil")
	}
	if err := b.ConsumeID(EditForumTopicRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode editForumTopic#f2c261b0: %w", err)
	}
	return e.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (e *EditForumTopicRequest) DecodeBare(b *bin.Buffer) error {
	if e == nil {
		return fmt.Errorf("can't decode editForumTopic#f2c261b0 to nil")
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode editForumTopic#f2c261b0: field chat_id: %w", err)
		}
		e.ChatID = value
	}
	{
		value, err := b.Int53()
		if err != nil {
			return fmt.Errorf("unable to decode editForumTopic#f2c261b0: field message_thread_id: %w", err)
		}
		e.MessageThreadID = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode editForumTopic#f2c261b0: field name: %w", err)
		}
		e.Name = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode editForumTopic#f2c261b0: field icon_custom_emoji_id: %w", err)
		}
		e.IconCustomEmojiID = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (e *EditForumTopicRequest) EncodeTDLibJSON(b tdjson.Encoder) error {
	if e == nil {
		return fmt.Errorf("can't encode editForumTopic#f2c261b0 as nil")
	}
	b.ObjStart()
	b.PutID("editForumTopic")
	b.Comma()
	b.FieldStart("chat_id")
	b.PutInt53(e.ChatID)
	b.Comma()
	b.FieldStart("message_thread_id")
	b.PutInt53(e.MessageThreadID)
	b.Comma()
	b.FieldStart("name")
	b.PutString(e.Name)
	b.Comma()
	b.FieldStart("icon_custom_emoji_id")
	b.PutLong(e.IconCustomEmojiID)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (e *EditForumTopicRequest) DecodeTDLibJSON(b tdjson.Decoder) error {
	if e == nil {
		return fmt.Errorf("can't decode editForumTopic#f2c261b0 to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("editForumTopic"); err != nil {
				return fmt.Errorf("unable to decode editForumTopic#f2c261b0: %w", err)
			}
		case "chat_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode editForumTopic#f2c261b0: field chat_id: %w", err)
			}
			e.ChatID = value
		case "message_thread_id":
			value, err := b.Int53()
			if err != nil {
				return fmt.Errorf("unable to decode editForumTopic#f2c261b0: field message_thread_id: %w", err)
			}
			e.MessageThreadID = value
		case "name":
			value, err := b.String()
			if err != nil {
				return fmt.Errorf("unable to decode editForumTopic#f2c261b0: field name: %w", err)
			}
			e.Name = value
		case "icon_custom_emoji_id":
			value, err := b.Long()
			if err != nil {
				return fmt.Errorf("unable to decode editForumTopic#f2c261b0: field icon_custom_emoji_id: %w", err)
			}
			e.IconCustomEmojiID = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetChatID returns value of ChatID field.
func (e *EditForumTopicRequest) GetChatID() (value int64) {
	if e == nil {
		return
	}
	return e.ChatID
}

// GetMessageThreadID returns value of MessageThreadID field.
func (e *EditForumTopicRequest) GetMessageThreadID() (value int64) {
	if e == nil {
		return
	}
	return e.MessageThreadID
}

// GetName returns value of Name field.
func (e *EditForumTopicRequest) GetName() (value string) {
	if e == nil {
		return
	}
	return e.Name
}

// GetIconCustomEmojiID returns value of IconCustomEmojiID field.
func (e *EditForumTopicRequest) GetIconCustomEmojiID() (value int64) {
	if e == nil {
		return
	}
	return e.IconCustomEmojiID
}

// EditForumTopic invokes method editForumTopic#f2c261b0 returning error if any.
func (c *Client) EditForumTopic(ctx context.Context, request *EditForumTopicRequest) error {
	var ok Ok

	if err := c.rpc.Invoke(ctx, request, &ok); err != nil {
		return err
	}
	return nil
}
