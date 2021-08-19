// Code generated by gotdgen, DO NOT EDIT.

package td

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

// UserAuth represents TL type `user.auth#f4815592`.
//
// See https://localhost:80/doc/constructor/user.auth for reference.
type UserAuth struct {
	// Foo field of UserAuth.
	Foo string
}

// UserAuthTypeID is TL type id of UserAuth.
const UserAuthTypeID = 0xf4815592

// construct implements constructor of UserAuthClass.
func (a UserAuth) construct() UserAuthClass { return &a }

// Ensuring interfaces in compile-time for UserAuth.
var (
	_ bin.Encoder     = &UserAuth{}
	_ bin.Decoder     = &UserAuth{}
	_ bin.BareEncoder = &UserAuth{}
	_ bin.BareDecoder = &UserAuth{}

	_ UserAuthClass = &UserAuth{}
)

func (a *UserAuth) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Foo == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *UserAuth) String() string {
	if a == nil {
		return "UserAuth(nil)"
	}
	type Alias UserAuth
	return fmt.Sprintf("UserAuth%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UserAuth) TypeID() uint32 {
	return UserAuthTypeID
}

// TypeName returns name of type in TL schema.
func (*UserAuth) TypeName() string {
	return "user.auth"
}

// TypeInfo returns info about TL type.
func (a *UserAuth) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "user.auth",
		ID:   UserAuthTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Foo",
			SchemaName: "foo",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *UserAuth) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode user.auth#f4815592 as nil")
	}
	b.PutID(UserAuthTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *UserAuth) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode user.auth#f4815592 as nil")
	}
	b.PutString(a.Foo)
	return nil
}

// Decode implements bin.Decoder.
func (a *UserAuth) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode user.auth#f4815592 to nil")
	}
	if err := b.ConsumeID(UserAuthTypeID); err != nil {
		return fmt.Errorf("unable to decode user.auth#f4815592: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *UserAuth) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode user.auth#f4815592 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user.auth#f4815592: field foo: %w", err)
		}
		a.Foo = value
	}
	return nil
}

// GetFoo returns value of Foo field.
func (a *UserAuth) GetFoo() (value string) {
	return a.Foo
}

// UserAuthPassword represents TL type `user.authPassword#5981e317`.
//
// See https://localhost:80/doc/constructor/user.authPassword for reference.
type UserAuthPassword struct {
	// Pwd field of UserAuthPassword.
	Pwd string
}

// UserAuthPasswordTypeID is TL type id of UserAuthPassword.
const UserAuthPasswordTypeID = 0x5981e317

// construct implements constructor of UserAuthClass.
func (a UserAuthPassword) construct() UserAuthClass { return &a }

// Ensuring interfaces in compile-time for UserAuthPassword.
var (
	_ bin.Encoder     = &UserAuthPassword{}
	_ bin.Decoder     = &UserAuthPassword{}
	_ bin.BareEncoder = &UserAuthPassword{}
	_ bin.BareDecoder = &UserAuthPassword{}

	_ UserAuthClass = &UserAuthPassword{}
)

func (a *UserAuthPassword) Zero() bool {
	if a == nil {
		return true
	}
	if !(a.Pwd == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (a *UserAuthPassword) String() string {
	if a == nil {
		return "UserAuthPassword(nil)"
	}
	type Alias UserAuthPassword
	return fmt.Sprintf("UserAuthPassword%+v", Alias(*a))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*UserAuthPassword) TypeID() uint32 {
	return UserAuthPasswordTypeID
}

// TypeName returns name of type in TL schema.
func (*UserAuthPassword) TypeName() string {
	return "user.authPassword"
}

// TypeInfo returns info about TL type.
func (a *UserAuthPassword) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "user.authPassword",
		ID:   UserAuthPasswordTypeID,
	}
	if a == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Pwd",
			SchemaName: "pwd",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (a *UserAuthPassword) Encode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode user.authPassword#5981e317 as nil")
	}
	b.PutID(UserAuthPasswordTypeID)
	return a.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (a *UserAuthPassword) EncodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't encode user.authPassword#5981e317 as nil")
	}
	b.PutString(a.Pwd)
	return nil
}

// Decode implements bin.Decoder.
func (a *UserAuthPassword) Decode(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode user.authPassword#5981e317 to nil")
	}
	if err := b.ConsumeID(UserAuthPasswordTypeID); err != nil {
		return fmt.Errorf("unable to decode user.authPassword#5981e317: %w", err)
	}
	return a.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (a *UserAuthPassword) DecodeBare(b *bin.Buffer) error {
	if a == nil {
		return fmt.Errorf("can't decode user.authPassword#5981e317 to nil")
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode user.authPassword#5981e317: field pwd: %w", err)
		}
		a.Pwd = value
	}
	return nil
}

// GetPwd returns value of Pwd field.
func (a *UserAuthPassword) GetPwd() (value string) {
	return a.Pwd
}

// UserAuthClass represents user.Auth generic type.
//
// See https://localhost:80/doc/type/user.Auth for reference.
//
// Example:
//  g, err := td.DecodeUserAuth(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *td.UserAuth: // user.auth#f4815592
//  case *td.UserAuthPassword: // user.authPassword#5981e317
//  default: panic(v)
//  }
type UserAuthClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() UserAuthClass

	// TypeID returns type id in TL schema.
	//
	// See https://core.telegram.org/mtproto/TL-tl#remarks.
	TypeID() uint32
	// TypeName returns name of type in TL schema.
	TypeName() string
	// String implements fmt.Stringer.
	String() string
	// Zero returns true if current object has a zero value.
	Zero() bool
}

// DecodeUserAuth implements binary de-serialization for UserAuthClass.
func DecodeUserAuth(buf *bin.Buffer) (UserAuthClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case UserAuthTypeID:
		// Decoding user.auth#f4815592.
		v := UserAuth{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserAuthClass: %w", err)
		}
		return &v, nil
	case UserAuthPasswordTypeID:
		// Decoding user.authPassword#5981e317.
		v := UserAuthPassword{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode UserAuthClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode UserAuthClass: %w", bin.NewUnexpectedID(id))
	}
}

// UserAuth boxes the UserAuthClass providing a helper.
type UserAuthBox struct {
	Auth UserAuthClass
}

// Decode implements bin.Decoder for UserAuthBox.
func (b *UserAuthBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode UserAuthBox to nil")
	}
	v, err := DecodeUserAuth(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.Auth = v
	return nil
}

// Encode implements bin.Encode for UserAuthBox.
func (b *UserAuthBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.Auth == nil {
		return fmt.Errorf("unable to encode UserAuthClass as nil")
	}
	return b.Auth.Encode(buf)
}
