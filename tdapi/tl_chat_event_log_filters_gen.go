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

// ChatEventLogFilters represents TL type `chatEventLogFilters#623e7a2d`.
type ChatEventLogFilters struct {
	// True, if message edits need to be returned
	MessageEdits bool
	// True, if message deletions need to be returned
	MessageDeletions bool
	// True, if pin/unpin events need to be returned
	MessagePins bool
	// True, if members joining events need to be returned
	MemberJoins bool
	// True, if members leaving events need to be returned
	MemberLeaves bool
	// True, if invited member events need to be returned
	MemberInvites bool
	// True, if member promotion/demotion events need to be returned
	MemberPromotions bool
	// True, if member restricted/unrestricted/banned/unbanned events need to be returned
	MemberRestrictions bool
	// True, if changes in chat information need to be returned
	InfoChanges bool
	// True, if changes in chat settings need to be returned
	SettingChanges bool
	// True, if changes to invite links need to be returned
	InviteLinkChanges bool
	// True, if video chat actions need to be returned
	VideoChatChanges bool
	// True, if forum-related actions need to be returned
	ForumChanges bool
}

// ChatEventLogFiltersTypeID is TL type id of ChatEventLogFilters.
const ChatEventLogFiltersTypeID = 0x623e7a2d

// Ensuring interfaces in compile-time for ChatEventLogFilters.
var (
	_ bin.Encoder     = &ChatEventLogFilters{}
	_ bin.Decoder     = &ChatEventLogFilters{}
	_ bin.BareEncoder = &ChatEventLogFilters{}
	_ bin.BareDecoder = &ChatEventLogFilters{}
)

func (c *ChatEventLogFilters) Zero() bool {
	if c == nil {
		return true
	}
	if !(c.MessageEdits == false) {
		return false
	}
	if !(c.MessageDeletions == false) {
		return false
	}
	if !(c.MessagePins == false) {
		return false
	}
	if !(c.MemberJoins == false) {
		return false
	}
	if !(c.MemberLeaves == false) {
		return false
	}
	if !(c.MemberInvites == false) {
		return false
	}
	if !(c.MemberPromotions == false) {
		return false
	}
	if !(c.MemberRestrictions == false) {
		return false
	}
	if !(c.InfoChanges == false) {
		return false
	}
	if !(c.SettingChanges == false) {
		return false
	}
	if !(c.InviteLinkChanges == false) {
		return false
	}
	if !(c.VideoChatChanges == false) {
		return false
	}
	if !(c.ForumChanges == false) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (c *ChatEventLogFilters) String() string {
	if c == nil {
		return "ChatEventLogFilters(nil)"
	}
	type Alias ChatEventLogFilters
	return fmt.Sprintf("ChatEventLogFilters%+v", Alias(*c))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*ChatEventLogFilters) TypeID() uint32 {
	return ChatEventLogFiltersTypeID
}

// TypeName returns name of type in TL schema.
func (*ChatEventLogFilters) TypeName() string {
	return "chatEventLogFilters"
}

// TypeInfo returns info about TL type.
func (c *ChatEventLogFilters) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "chatEventLogFilters",
		ID:   ChatEventLogFiltersTypeID,
	}
	if c == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "MessageEdits",
			SchemaName: "message_edits",
		},
		{
			Name:       "MessageDeletions",
			SchemaName: "message_deletions",
		},
		{
			Name:       "MessagePins",
			SchemaName: "message_pins",
		},
		{
			Name:       "MemberJoins",
			SchemaName: "member_joins",
		},
		{
			Name:       "MemberLeaves",
			SchemaName: "member_leaves",
		},
		{
			Name:       "MemberInvites",
			SchemaName: "member_invites",
		},
		{
			Name:       "MemberPromotions",
			SchemaName: "member_promotions",
		},
		{
			Name:       "MemberRestrictions",
			SchemaName: "member_restrictions",
		},
		{
			Name:       "InfoChanges",
			SchemaName: "info_changes",
		},
		{
			Name:       "SettingChanges",
			SchemaName: "setting_changes",
		},
		{
			Name:       "InviteLinkChanges",
			SchemaName: "invite_link_changes",
		},
		{
			Name:       "VideoChatChanges",
			SchemaName: "video_chat_changes",
		},
		{
			Name:       "ForumChanges",
			SchemaName: "forum_changes",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (c *ChatEventLogFilters) Encode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatEventLogFilters#623e7a2d as nil")
	}
	b.PutID(ChatEventLogFiltersTypeID)
	return c.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (c *ChatEventLogFilters) EncodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't encode chatEventLogFilters#623e7a2d as nil")
	}
	b.PutBool(c.MessageEdits)
	b.PutBool(c.MessageDeletions)
	b.PutBool(c.MessagePins)
	b.PutBool(c.MemberJoins)
	b.PutBool(c.MemberLeaves)
	b.PutBool(c.MemberInvites)
	b.PutBool(c.MemberPromotions)
	b.PutBool(c.MemberRestrictions)
	b.PutBool(c.InfoChanges)
	b.PutBool(c.SettingChanges)
	b.PutBool(c.InviteLinkChanges)
	b.PutBool(c.VideoChatChanges)
	b.PutBool(c.ForumChanges)
	return nil
}

// Decode implements bin.Decoder.
func (c *ChatEventLogFilters) Decode(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatEventLogFilters#623e7a2d to nil")
	}
	if err := b.ConsumeID(ChatEventLogFiltersTypeID); err != nil {
		return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: %w", err)
	}
	return c.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (c *ChatEventLogFilters) DecodeBare(b *bin.Buffer) error {
	if c == nil {
		return fmt.Errorf("can't decode chatEventLogFilters#623e7a2d to nil")
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field message_edits: %w", err)
		}
		c.MessageEdits = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field message_deletions: %w", err)
		}
		c.MessageDeletions = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field message_pins: %w", err)
		}
		c.MessagePins = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field member_joins: %w", err)
		}
		c.MemberJoins = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field member_leaves: %w", err)
		}
		c.MemberLeaves = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field member_invites: %w", err)
		}
		c.MemberInvites = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field member_promotions: %w", err)
		}
		c.MemberPromotions = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field member_restrictions: %w", err)
		}
		c.MemberRestrictions = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field info_changes: %w", err)
		}
		c.InfoChanges = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field setting_changes: %w", err)
		}
		c.SettingChanges = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field invite_link_changes: %w", err)
		}
		c.InviteLinkChanges = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field video_chat_changes: %w", err)
		}
		c.VideoChatChanges = value
	}
	{
		value, err := b.Bool()
		if err != nil {
			return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field forum_changes: %w", err)
		}
		c.ForumChanges = value
	}
	return nil
}

// EncodeTDLibJSON implements tdjson.TDLibEncoder.
func (c *ChatEventLogFilters) EncodeTDLibJSON(b tdjson.Encoder) error {
	if c == nil {
		return fmt.Errorf("can't encode chatEventLogFilters#623e7a2d as nil")
	}
	b.ObjStart()
	b.PutID("chatEventLogFilters")
	b.Comma()
	b.FieldStart("message_edits")
	b.PutBool(c.MessageEdits)
	b.Comma()
	b.FieldStart("message_deletions")
	b.PutBool(c.MessageDeletions)
	b.Comma()
	b.FieldStart("message_pins")
	b.PutBool(c.MessagePins)
	b.Comma()
	b.FieldStart("member_joins")
	b.PutBool(c.MemberJoins)
	b.Comma()
	b.FieldStart("member_leaves")
	b.PutBool(c.MemberLeaves)
	b.Comma()
	b.FieldStart("member_invites")
	b.PutBool(c.MemberInvites)
	b.Comma()
	b.FieldStart("member_promotions")
	b.PutBool(c.MemberPromotions)
	b.Comma()
	b.FieldStart("member_restrictions")
	b.PutBool(c.MemberRestrictions)
	b.Comma()
	b.FieldStart("info_changes")
	b.PutBool(c.InfoChanges)
	b.Comma()
	b.FieldStart("setting_changes")
	b.PutBool(c.SettingChanges)
	b.Comma()
	b.FieldStart("invite_link_changes")
	b.PutBool(c.InviteLinkChanges)
	b.Comma()
	b.FieldStart("video_chat_changes")
	b.PutBool(c.VideoChatChanges)
	b.Comma()
	b.FieldStart("forum_changes")
	b.PutBool(c.ForumChanges)
	b.Comma()
	b.StripComma()
	b.ObjEnd()
	return nil
}

// DecodeTDLibJSON implements tdjson.TDLibDecoder.
func (c *ChatEventLogFilters) DecodeTDLibJSON(b tdjson.Decoder) error {
	if c == nil {
		return fmt.Errorf("can't decode chatEventLogFilters#623e7a2d to nil")
	}

	return b.Obj(func(b tdjson.Decoder, key []byte) error {
		switch string(key) {
		case tdjson.TypeField:
			if err := b.ConsumeID("chatEventLogFilters"); err != nil {
				return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: %w", err)
			}
		case "message_edits":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field message_edits: %w", err)
			}
			c.MessageEdits = value
		case "message_deletions":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field message_deletions: %w", err)
			}
			c.MessageDeletions = value
		case "message_pins":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field message_pins: %w", err)
			}
			c.MessagePins = value
		case "member_joins":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field member_joins: %w", err)
			}
			c.MemberJoins = value
		case "member_leaves":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field member_leaves: %w", err)
			}
			c.MemberLeaves = value
		case "member_invites":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field member_invites: %w", err)
			}
			c.MemberInvites = value
		case "member_promotions":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field member_promotions: %w", err)
			}
			c.MemberPromotions = value
		case "member_restrictions":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field member_restrictions: %w", err)
			}
			c.MemberRestrictions = value
		case "info_changes":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field info_changes: %w", err)
			}
			c.InfoChanges = value
		case "setting_changes":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field setting_changes: %w", err)
			}
			c.SettingChanges = value
		case "invite_link_changes":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field invite_link_changes: %w", err)
			}
			c.InviteLinkChanges = value
		case "video_chat_changes":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field video_chat_changes: %w", err)
			}
			c.VideoChatChanges = value
		case "forum_changes":
			value, err := b.Bool()
			if err != nil {
				return fmt.Errorf("unable to decode chatEventLogFilters#623e7a2d: field forum_changes: %w", err)
			}
			c.ForumChanges = value
		default:
			return b.Skip()
		}
		return nil
	})
}

// GetMessageEdits returns value of MessageEdits field.
func (c *ChatEventLogFilters) GetMessageEdits() (value bool) {
	if c == nil {
		return
	}
	return c.MessageEdits
}

// GetMessageDeletions returns value of MessageDeletions field.
func (c *ChatEventLogFilters) GetMessageDeletions() (value bool) {
	if c == nil {
		return
	}
	return c.MessageDeletions
}

// GetMessagePins returns value of MessagePins field.
func (c *ChatEventLogFilters) GetMessagePins() (value bool) {
	if c == nil {
		return
	}
	return c.MessagePins
}

// GetMemberJoins returns value of MemberJoins field.
func (c *ChatEventLogFilters) GetMemberJoins() (value bool) {
	if c == nil {
		return
	}
	return c.MemberJoins
}

// GetMemberLeaves returns value of MemberLeaves field.
func (c *ChatEventLogFilters) GetMemberLeaves() (value bool) {
	if c == nil {
		return
	}
	return c.MemberLeaves
}

// GetMemberInvites returns value of MemberInvites field.
func (c *ChatEventLogFilters) GetMemberInvites() (value bool) {
	if c == nil {
		return
	}
	return c.MemberInvites
}

// GetMemberPromotions returns value of MemberPromotions field.
func (c *ChatEventLogFilters) GetMemberPromotions() (value bool) {
	if c == nil {
		return
	}
	return c.MemberPromotions
}

// GetMemberRestrictions returns value of MemberRestrictions field.
func (c *ChatEventLogFilters) GetMemberRestrictions() (value bool) {
	if c == nil {
		return
	}
	return c.MemberRestrictions
}

// GetInfoChanges returns value of InfoChanges field.
func (c *ChatEventLogFilters) GetInfoChanges() (value bool) {
	if c == nil {
		return
	}
	return c.InfoChanges
}

// GetSettingChanges returns value of SettingChanges field.
func (c *ChatEventLogFilters) GetSettingChanges() (value bool) {
	if c == nil {
		return
	}
	return c.SettingChanges
}

// GetInviteLinkChanges returns value of InviteLinkChanges field.
func (c *ChatEventLogFilters) GetInviteLinkChanges() (value bool) {
	if c == nil {
		return
	}
	return c.InviteLinkChanges
}

// GetVideoChatChanges returns value of VideoChatChanges field.
func (c *ChatEventLogFilters) GetVideoChatChanges() (value bool) {
	if c == nil {
		return
	}
	return c.VideoChatChanges
}

// GetForumChanges returns value of ForumChanges field.
func (c *ChatEventLogFilters) GetForumChanges() (value bool) {
	if c == nil {
		return
	}
	return c.ForumChanges
}
