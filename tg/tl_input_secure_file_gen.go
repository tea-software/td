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

// InputSecureFileUploaded represents TL type `inputSecureFileUploaded#3334b0f0`.
// Uploaded secure file, for more info see the passport docs »¹
//
// Links:
//  1) https://core.telegram.org/passport/encryption#inputsecurefile
//
// See https://core.telegram.org/constructor/inputSecureFileUploaded for reference.
type InputSecureFileUploaded struct {
	// Secure file ID
	ID int64
	// Secure file part count
	Parts int
	// MD5 hash of encrypted uploaded file, to be checked server-side
	MD5Checksum string
	// File hash
	FileHash []byte
	// Secret
	Secret []byte
}

// InputSecureFileUploadedTypeID is TL type id of InputSecureFileUploaded.
const InputSecureFileUploadedTypeID = 0x3334b0f0

// construct implements constructor of InputSecureFileClass.
func (i InputSecureFileUploaded) construct() InputSecureFileClass { return &i }

// Ensuring interfaces in compile-time for InputSecureFileUploaded.
var (
	_ bin.Encoder     = &InputSecureFileUploaded{}
	_ bin.Decoder     = &InputSecureFileUploaded{}
	_ bin.BareEncoder = &InputSecureFileUploaded{}
	_ bin.BareDecoder = &InputSecureFileUploaded{}

	_ InputSecureFileClass = &InputSecureFileUploaded{}
)

func (i *InputSecureFileUploaded) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == 0) {
		return false
	}
	if !(i.Parts == 0) {
		return false
	}
	if !(i.MD5Checksum == "") {
		return false
	}
	if !(i.FileHash == nil) {
		return false
	}
	if !(i.Secret == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputSecureFileUploaded) String() string {
	if i == nil {
		return "InputSecureFileUploaded(nil)"
	}
	type Alias InputSecureFileUploaded
	return fmt.Sprintf("InputSecureFileUploaded%+v", Alias(*i))
}

// FillFrom fills InputSecureFileUploaded from given interface.
func (i *InputSecureFileUploaded) FillFrom(from interface {
	GetID() (value int64)
	GetParts() (value int)
	GetMD5Checksum() (value string)
	GetFileHash() (value []byte)
	GetSecret() (value []byte)
}) {
	i.ID = from.GetID()
	i.Parts = from.GetParts()
	i.MD5Checksum = from.GetMD5Checksum()
	i.FileHash = from.GetFileHash()
	i.Secret = from.GetSecret()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputSecureFileUploaded) TypeID() uint32 {
	return InputSecureFileUploadedTypeID
}

// TypeName returns name of type in TL schema.
func (*InputSecureFileUploaded) TypeName() string {
	return "inputSecureFileUploaded"
}

// TypeInfo returns info about TL type.
func (i *InputSecureFileUploaded) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputSecureFileUploaded",
		ID:   InputSecureFileUploadedTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "Parts",
			SchemaName: "parts",
		},
		{
			Name:       "MD5Checksum",
			SchemaName: "md5_checksum",
		},
		{
			Name:       "FileHash",
			SchemaName: "file_hash",
		},
		{
			Name:       "Secret",
			SchemaName: "secret",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputSecureFileUploaded) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputSecureFileUploaded#3334b0f0 as nil")
	}
	b.PutID(InputSecureFileUploadedTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputSecureFileUploaded) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputSecureFileUploaded#3334b0f0 as nil")
	}
	b.PutLong(i.ID)
	b.PutInt(i.Parts)
	b.PutString(i.MD5Checksum)
	b.PutBytes(i.FileHash)
	b.PutBytes(i.Secret)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputSecureFileUploaded) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputSecureFileUploaded#3334b0f0 to nil")
	}
	if err := b.ConsumeID(InputSecureFileUploadedTypeID); err != nil {
		return fmt.Errorf("unable to decode inputSecureFileUploaded#3334b0f0: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputSecureFileUploaded) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputSecureFileUploaded#3334b0f0 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureFileUploaded#3334b0f0: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureFileUploaded#3334b0f0: field parts: %w", err)
		}
		i.Parts = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureFileUploaded#3334b0f0: field md5_checksum: %w", err)
		}
		i.MD5Checksum = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureFileUploaded#3334b0f0: field file_hash: %w", err)
		}
		i.FileHash = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureFileUploaded#3334b0f0: field secret: %w", err)
		}
		i.Secret = value
	}
	return nil
}

// GetID returns value of ID field.
func (i *InputSecureFileUploaded) GetID() (value int64) {
	return i.ID
}

// GetParts returns value of Parts field.
func (i *InputSecureFileUploaded) GetParts() (value int) {
	return i.Parts
}

// GetMD5Checksum returns value of MD5Checksum field.
func (i *InputSecureFileUploaded) GetMD5Checksum() (value string) {
	return i.MD5Checksum
}

// GetFileHash returns value of FileHash field.
func (i *InputSecureFileUploaded) GetFileHash() (value []byte) {
	return i.FileHash
}

// GetSecret returns value of Secret field.
func (i *InputSecureFileUploaded) GetSecret() (value []byte) {
	return i.Secret
}

// InputSecureFile represents TL type `inputSecureFile#5367e5be`.
// Preuploaded passport¹ file, for more info see the passport docs »²
//
// Links:
//  1) https://core.telegram.org/passport
//  2) https://core.telegram.org/passport/encryption#inputsecurefile
//
// See https://core.telegram.org/constructor/inputSecureFile for reference.
type InputSecureFile struct {
	// Secure file ID
	ID int64
	// Secure file access hash
	AccessHash int64
}

// InputSecureFileTypeID is TL type id of InputSecureFile.
const InputSecureFileTypeID = 0x5367e5be

// construct implements constructor of InputSecureFileClass.
func (i InputSecureFile) construct() InputSecureFileClass { return &i }

// Ensuring interfaces in compile-time for InputSecureFile.
var (
	_ bin.Encoder     = &InputSecureFile{}
	_ bin.Decoder     = &InputSecureFile{}
	_ bin.BareEncoder = &InputSecureFile{}
	_ bin.BareDecoder = &InputSecureFile{}

	_ InputSecureFileClass = &InputSecureFile{}
)

func (i *InputSecureFile) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == 0) {
		return false
	}
	if !(i.AccessHash == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputSecureFile) String() string {
	if i == nil {
		return "InputSecureFile(nil)"
	}
	type Alias InputSecureFile
	return fmt.Sprintf("InputSecureFile%+v", Alias(*i))
}

// FillFrom fills InputSecureFile from given interface.
func (i *InputSecureFile) FillFrom(from interface {
	GetID() (value int64)
	GetAccessHash() (value int64)
}) {
	i.ID = from.GetID()
	i.AccessHash = from.GetAccessHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputSecureFile) TypeID() uint32 {
	return InputSecureFileTypeID
}

// TypeName returns name of type in TL schema.
func (*InputSecureFile) TypeName() string {
	return "inputSecureFile"
}

// TypeInfo returns info about TL type.
func (i *InputSecureFile) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputSecureFile",
		ID:   InputSecureFileTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "ID",
			SchemaName: "id",
		},
		{
			Name:       "AccessHash",
			SchemaName: "access_hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputSecureFile) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputSecureFile#5367e5be as nil")
	}
	b.PutID(InputSecureFileTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputSecureFile) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputSecureFile#5367e5be as nil")
	}
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputSecureFile) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputSecureFile#5367e5be to nil")
	}
	if err := b.ConsumeID(InputSecureFileTypeID); err != nil {
		return fmt.Errorf("unable to decode inputSecureFile#5367e5be: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputSecureFile) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputSecureFile#5367e5be to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureFile#5367e5be: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputSecureFile#5367e5be: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// GetID returns value of ID field.
func (i *InputSecureFile) GetID() (value int64) {
	return i.ID
}

// GetAccessHash returns value of AccessHash field.
func (i *InputSecureFile) GetAccessHash() (value int64) {
	return i.AccessHash
}

// InputSecureFileClass represents InputSecureFile generic type.
//
// See https://core.telegram.org/type/InputSecureFile for reference.
//
// Example:
//  g, err := tg.DecodeInputSecureFile(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.InputSecureFileUploaded: // inputSecureFileUploaded#3334b0f0
//  case *tg.InputSecureFile: // inputSecureFile#5367e5be
//  default: panic(v)
//  }
type InputSecureFileClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputSecureFileClass

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

	// Secure file ID
	GetID() (value int64)
}

// AsInputSecureFileLocation tries to map InputSecureFile to InputSecureFileLocation.
func (i *InputSecureFile) AsInputSecureFileLocation() *InputSecureFileLocation {
	value := new(InputSecureFileLocation)
	value.ID = i.GetID()
	value.AccessHash = i.GetAccessHash()

	return value
}

// DecodeInputSecureFile implements binary de-serialization for InputSecureFileClass.
func DecodeInputSecureFile(buf *bin.Buffer) (InputSecureFileClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputSecureFileUploadedTypeID:
		// Decoding inputSecureFileUploaded#3334b0f0.
		v := InputSecureFileUploaded{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputSecureFileClass: %w", err)
		}
		return &v, nil
	case InputSecureFileTypeID:
		// Decoding inputSecureFile#5367e5be.
		v := InputSecureFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputSecureFileClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputSecureFileClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputSecureFile boxes the InputSecureFileClass providing a helper.
type InputSecureFileBox struct {
	InputSecureFile InputSecureFileClass
}

// Decode implements bin.Decoder for InputSecureFileBox.
func (b *InputSecureFileBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputSecureFileBox to nil")
	}
	v, err := DecodeInputSecureFile(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputSecureFile = v
	return nil
}

// Encode implements bin.Encode for InputSecureFileBox.
func (b *InputSecureFileBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputSecureFile == nil {
		return fmt.Errorf("unable to encode InputSecureFileClass as nil")
	}
	return b.InputSecureFile.Encode(buf)
}
