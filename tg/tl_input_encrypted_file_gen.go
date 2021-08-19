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

// InputEncryptedFileEmpty represents TL type `inputEncryptedFileEmpty#1837c364`.
// Empty constructor.
//
// See https://core.telegram.org/constructor/inputEncryptedFileEmpty for reference.
type InputEncryptedFileEmpty struct {
}

// InputEncryptedFileEmptyTypeID is TL type id of InputEncryptedFileEmpty.
const InputEncryptedFileEmptyTypeID = 0x1837c364

// construct implements constructor of InputEncryptedFileClass.
func (i InputEncryptedFileEmpty) construct() InputEncryptedFileClass { return &i }

// Ensuring interfaces in compile-time for InputEncryptedFileEmpty.
var (
	_ bin.Encoder     = &InputEncryptedFileEmpty{}
	_ bin.Decoder     = &InputEncryptedFileEmpty{}
	_ bin.BareEncoder = &InputEncryptedFileEmpty{}
	_ bin.BareDecoder = &InputEncryptedFileEmpty{}

	_ InputEncryptedFileClass = &InputEncryptedFileEmpty{}
)

func (i *InputEncryptedFileEmpty) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputEncryptedFileEmpty) String() string {
	if i == nil {
		return "InputEncryptedFileEmpty(nil)"
	}
	type Alias InputEncryptedFileEmpty
	return fmt.Sprintf("InputEncryptedFileEmpty%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputEncryptedFileEmpty) TypeID() uint32 {
	return InputEncryptedFileEmptyTypeID
}

// TypeName returns name of type in TL schema.
func (*InputEncryptedFileEmpty) TypeName() string {
	return "inputEncryptedFileEmpty"
}

// TypeInfo returns info about TL type.
func (i *InputEncryptedFileEmpty) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputEncryptedFileEmpty",
		ID:   InputEncryptedFileEmptyTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputEncryptedFileEmpty) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputEncryptedFileEmpty#1837c364 as nil")
	}
	b.PutID(InputEncryptedFileEmptyTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputEncryptedFileEmpty) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputEncryptedFileEmpty#1837c364 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputEncryptedFileEmpty) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputEncryptedFileEmpty#1837c364 to nil")
	}
	if err := b.ConsumeID(InputEncryptedFileEmptyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputEncryptedFileEmpty#1837c364: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputEncryptedFileEmpty) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputEncryptedFileEmpty#1837c364 to nil")
	}
	return nil
}

// InputEncryptedFileUploaded represents TL type `inputEncryptedFileUploaded#64bd0306`.
// Sets new encrypted file saved by parts using upload.saveFilePart method.
//
// See https://core.telegram.org/constructor/inputEncryptedFileUploaded for reference.
type InputEncryptedFileUploaded struct {
	// Random file ID created by clien
	ID int64
	// Number of saved parts
	Parts int
	// In case md5-HASH¹ of the (already encrypted) file was transmitted, file content will
	// be checked prior to use
	//
	// Links:
	//  1) https://en.wikipedia.org/wiki/MD5
	MD5Checksum string
	// 32-bit fingerprint of the key used to encrypt a file
	KeyFingerprint int
}

// InputEncryptedFileUploadedTypeID is TL type id of InputEncryptedFileUploaded.
const InputEncryptedFileUploadedTypeID = 0x64bd0306

// construct implements constructor of InputEncryptedFileClass.
func (i InputEncryptedFileUploaded) construct() InputEncryptedFileClass { return &i }

// Ensuring interfaces in compile-time for InputEncryptedFileUploaded.
var (
	_ bin.Encoder     = &InputEncryptedFileUploaded{}
	_ bin.Decoder     = &InputEncryptedFileUploaded{}
	_ bin.BareEncoder = &InputEncryptedFileUploaded{}
	_ bin.BareDecoder = &InputEncryptedFileUploaded{}

	_ InputEncryptedFileClass = &InputEncryptedFileUploaded{}
)

func (i *InputEncryptedFileUploaded) Zero() bool {
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
	if !(i.KeyFingerprint == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputEncryptedFileUploaded) String() string {
	if i == nil {
		return "InputEncryptedFileUploaded(nil)"
	}
	type Alias InputEncryptedFileUploaded
	return fmt.Sprintf("InputEncryptedFileUploaded%+v", Alias(*i))
}

// FillFrom fills InputEncryptedFileUploaded from given interface.
func (i *InputEncryptedFileUploaded) FillFrom(from interface {
	GetID() (value int64)
	GetParts() (value int)
	GetMD5Checksum() (value string)
	GetKeyFingerprint() (value int)
}) {
	i.ID = from.GetID()
	i.Parts = from.GetParts()
	i.MD5Checksum = from.GetMD5Checksum()
	i.KeyFingerprint = from.GetKeyFingerprint()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputEncryptedFileUploaded) TypeID() uint32 {
	return InputEncryptedFileUploadedTypeID
}

// TypeName returns name of type in TL schema.
func (*InputEncryptedFileUploaded) TypeName() string {
	return "inputEncryptedFileUploaded"
}

// TypeInfo returns info about TL type.
func (i *InputEncryptedFileUploaded) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputEncryptedFileUploaded",
		ID:   InputEncryptedFileUploadedTypeID,
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
			Name:       "KeyFingerprint",
			SchemaName: "key_fingerprint",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputEncryptedFileUploaded) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputEncryptedFileUploaded#64bd0306 as nil")
	}
	b.PutID(InputEncryptedFileUploadedTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputEncryptedFileUploaded) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputEncryptedFileUploaded#64bd0306 as nil")
	}
	b.PutLong(i.ID)
	b.PutInt(i.Parts)
	b.PutString(i.MD5Checksum)
	b.PutInt(i.KeyFingerprint)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputEncryptedFileUploaded) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputEncryptedFileUploaded#64bd0306 to nil")
	}
	if err := b.ConsumeID(InputEncryptedFileUploadedTypeID); err != nil {
		return fmt.Errorf("unable to decode inputEncryptedFileUploaded#64bd0306: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputEncryptedFileUploaded) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputEncryptedFileUploaded#64bd0306 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFileUploaded#64bd0306: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFileUploaded#64bd0306: field parts: %w", err)
		}
		i.Parts = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFileUploaded#64bd0306: field md5_checksum: %w", err)
		}
		i.MD5Checksum = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFileUploaded#64bd0306: field key_fingerprint: %w", err)
		}
		i.KeyFingerprint = value
	}
	return nil
}

// GetID returns value of ID field.
func (i *InputEncryptedFileUploaded) GetID() (value int64) {
	return i.ID
}

// GetParts returns value of Parts field.
func (i *InputEncryptedFileUploaded) GetParts() (value int) {
	return i.Parts
}

// GetMD5Checksum returns value of MD5Checksum field.
func (i *InputEncryptedFileUploaded) GetMD5Checksum() (value string) {
	return i.MD5Checksum
}

// GetKeyFingerprint returns value of KeyFingerprint field.
func (i *InputEncryptedFileUploaded) GetKeyFingerprint() (value int) {
	return i.KeyFingerprint
}

// InputEncryptedFile represents TL type `inputEncryptedFile#5a17b5e5`.
// Sets forwarded encrypted file for attachment.
//
// See https://core.telegram.org/constructor/inputEncryptedFile for reference.
type InputEncryptedFile struct {
	// File ID, value of id parameter from encryptedFile¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/encryptedFile
	ID int64
	// Checking sum, value of access_hash parameter from encryptedFile¹
	//
	// Links:
	//  1) https://core.telegram.org/constructor/encryptedFile
	AccessHash int64
}

// InputEncryptedFileTypeID is TL type id of InputEncryptedFile.
const InputEncryptedFileTypeID = 0x5a17b5e5

// construct implements constructor of InputEncryptedFileClass.
func (i InputEncryptedFile) construct() InputEncryptedFileClass { return &i }

// Ensuring interfaces in compile-time for InputEncryptedFile.
var (
	_ bin.Encoder     = &InputEncryptedFile{}
	_ bin.Decoder     = &InputEncryptedFile{}
	_ bin.BareEncoder = &InputEncryptedFile{}
	_ bin.BareDecoder = &InputEncryptedFile{}

	_ InputEncryptedFileClass = &InputEncryptedFile{}
)

func (i *InputEncryptedFile) Zero() bool {
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
func (i *InputEncryptedFile) String() string {
	if i == nil {
		return "InputEncryptedFile(nil)"
	}
	type Alias InputEncryptedFile
	return fmt.Sprintf("InputEncryptedFile%+v", Alias(*i))
}

// FillFrom fills InputEncryptedFile from given interface.
func (i *InputEncryptedFile) FillFrom(from interface {
	GetID() (value int64)
	GetAccessHash() (value int64)
}) {
	i.ID = from.GetID()
	i.AccessHash = from.GetAccessHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputEncryptedFile) TypeID() uint32 {
	return InputEncryptedFileTypeID
}

// TypeName returns name of type in TL schema.
func (*InputEncryptedFile) TypeName() string {
	return "inputEncryptedFile"
}

// TypeInfo returns info about TL type.
func (i *InputEncryptedFile) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputEncryptedFile",
		ID:   InputEncryptedFileTypeID,
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
func (i *InputEncryptedFile) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputEncryptedFile#5a17b5e5 as nil")
	}
	b.PutID(InputEncryptedFileTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputEncryptedFile) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputEncryptedFile#5a17b5e5 as nil")
	}
	b.PutLong(i.ID)
	b.PutLong(i.AccessHash)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputEncryptedFile) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputEncryptedFile#5a17b5e5 to nil")
	}
	if err := b.ConsumeID(InputEncryptedFileTypeID); err != nil {
		return fmt.Errorf("unable to decode inputEncryptedFile#5a17b5e5: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputEncryptedFile) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputEncryptedFile#5a17b5e5 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFile#5a17b5e5: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFile#5a17b5e5: field access_hash: %w", err)
		}
		i.AccessHash = value
	}
	return nil
}

// GetID returns value of ID field.
func (i *InputEncryptedFile) GetID() (value int64) {
	return i.ID
}

// GetAccessHash returns value of AccessHash field.
func (i *InputEncryptedFile) GetAccessHash() (value int64) {
	return i.AccessHash
}

// InputEncryptedFileBigUploaded represents TL type `inputEncryptedFileBigUploaded#2dc173c8`.
// Assigns a new big encrypted file (over 10Mb in size), saved in parts using the method
// upload.saveBigFilePart¹.
//
// Links:
//  1) https://core.telegram.org/method/upload.saveBigFilePart
//
// See https://core.telegram.org/constructor/inputEncryptedFileBigUploaded for reference.
type InputEncryptedFileBigUploaded struct {
	// Random file id, created by the client
	ID int64
	// Number of saved parts
	Parts int
	// 32-bit imprint of the key used to encrypt the file
	KeyFingerprint int
}

// InputEncryptedFileBigUploadedTypeID is TL type id of InputEncryptedFileBigUploaded.
const InputEncryptedFileBigUploadedTypeID = 0x2dc173c8

// construct implements constructor of InputEncryptedFileClass.
func (i InputEncryptedFileBigUploaded) construct() InputEncryptedFileClass { return &i }

// Ensuring interfaces in compile-time for InputEncryptedFileBigUploaded.
var (
	_ bin.Encoder     = &InputEncryptedFileBigUploaded{}
	_ bin.Decoder     = &InputEncryptedFileBigUploaded{}
	_ bin.BareEncoder = &InputEncryptedFileBigUploaded{}
	_ bin.BareDecoder = &InputEncryptedFileBigUploaded{}

	_ InputEncryptedFileClass = &InputEncryptedFileBigUploaded{}
)

func (i *InputEncryptedFileBigUploaded) Zero() bool {
	if i == nil {
		return true
	}
	if !(i.ID == 0) {
		return false
	}
	if !(i.Parts == 0) {
		return false
	}
	if !(i.KeyFingerprint == 0) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputEncryptedFileBigUploaded) String() string {
	if i == nil {
		return "InputEncryptedFileBigUploaded(nil)"
	}
	type Alias InputEncryptedFileBigUploaded
	return fmt.Sprintf("InputEncryptedFileBigUploaded%+v", Alias(*i))
}

// FillFrom fills InputEncryptedFileBigUploaded from given interface.
func (i *InputEncryptedFileBigUploaded) FillFrom(from interface {
	GetID() (value int64)
	GetParts() (value int)
	GetKeyFingerprint() (value int)
}) {
	i.ID = from.GetID()
	i.Parts = from.GetParts()
	i.KeyFingerprint = from.GetKeyFingerprint()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputEncryptedFileBigUploaded) TypeID() uint32 {
	return InputEncryptedFileBigUploadedTypeID
}

// TypeName returns name of type in TL schema.
func (*InputEncryptedFileBigUploaded) TypeName() string {
	return "inputEncryptedFileBigUploaded"
}

// TypeInfo returns info about TL type.
func (i *InputEncryptedFileBigUploaded) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputEncryptedFileBigUploaded",
		ID:   InputEncryptedFileBigUploadedTypeID,
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
			Name:       "KeyFingerprint",
			SchemaName: "key_fingerprint",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputEncryptedFileBigUploaded) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputEncryptedFileBigUploaded#2dc173c8 as nil")
	}
	b.PutID(InputEncryptedFileBigUploadedTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputEncryptedFileBigUploaded) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputEncryptedFileBigUploaded#2dc173c8 as nil")
	}
	b.PutLong(i.ID)
	b.PutInt(i.Parts)
	b.PutInt(i.KeyFingerprint)
	return nil
}

// Decode implements bin.Decoder.
func (i *InputEncryptedFileBigUploaded) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputEncryptedFileBigUploaded#2dc173c8 to nil")
	}
	if err := b.ConsumeID(InputEncryptedFileBigUploadedTypeID); err != nil {
		return fmt.Errorf("unable to decode inputEncryptedFileBigUploaded#2dc173c8: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputEncryptedFileBigUploaded) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputEncryptedFileBigUploaded#2dc173c8 to nil")
	}
	{
		value, err := b.Long()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFileBigUploaded#2dc173c8: field id: %w", err)
		}
		i.ID = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFileBigUploaded#2dc173c8: field parts: %w", err)
		}
		i.Parts = value
	}
	{
		value, err := b.Int()
		if err != nil {
			return fmt.Errorf("unable to decode inputEncryptedFileBigUploaded#2dc173c8: field key_fingerprint: %w", err)
		}
		i.KeyFingerprint = value
	}
	return nil
}

// GetID returns value of ID field.
func (i *InputEncryptedFileBigUploaded) GetID() (value int64) {
	return i.ID
}

// GetParts returns value of Parts field.
func (i *InputEncryptedFileBigUploaded) GetParts() (value int) {
	return i.Parts
}

// GetKeyFingerprint returns value of KeyFingerprint field.
func (i *InputEncryptedFileBigUploaded) GetKeyFingerprint() (value int) {
	return i.KeyFingerprint
}

// InputEncryptedFileClass represents InputEncryptedFile generic type.
//
// See https://core.telegram.org/type/InputEncryptedFile for reference.
//
// Example:
//  g, err := tg.DecodeInputEncryptedFile(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.InputEncryptedFileEmpty: // inputEncryptedFileEmpty#1837c364
//  case *tg.InputEncryptedFileUploaded: // inputEncryptedFileUploaded#64bd0306
//  case *tg.InputEncryptedFile: // inputEncryptedFile#5a17b5e5
//  case *tg.InputEncryptedFileBigUploaded: // inputEncryptedFileBigUploaded#2dc173c8
//  default: panic(v)
//  }
type InputEncryptedFileClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputEncryptedFileClass

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

	// AsNotEmpty tries to map InputEncryptedFileClass to NotEmptyInputEncryptedFile.
	AsNotEmpty() (NotEmptyInputEncryptedFile, bool)
}

// AsInputEncryptedFileLocation tries to map InputEncryptedFile to InputEncryptedFileLocation.
func (i *InputEncryptedFile) AsInputEncryptedFileLocation() *InputEncryptedFileLocation {
	value := new(InputEncryptedFileLocation)
	value.ID = i.GetID()
	value.AccessHash = i.GetAccessHash()

	return value
}

// NotEmptyInputEncryptedFile represents NotEmpty subset of InputEncryptedFileClass.
type NotEmptyInputEncryptedFile interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() InputEncryptedFileClass

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

	// Random file ID created by clien
	GetID() (value int64)
}

// AsNotEmpty tries to map InputEncryptedFileEmpty to NotEmptyInputEncryptedFile.
func (i *InputEncryptedFileEmpty) AsNotEmpty() (NotEmptyInputEncryptedFile, bool) {
	value, ok := (InputEncryptedFileClass(i)).(NotEmptyInputEncryptedFile)
	return value, ok
}

// AsNotEmpty tries to map InputEncryptedFileUploaded to NotEmptyInputEncryptedFile.
func (i *InputEncryptedFileUploaded) AsNotEmpty() (NotEmptyInputEncryptedFile, bool) {
	value, ok := (InputEncryptedFileClass(i)).(NotEmptyInputEncryptedFile)
	return value, ok
}

// AsNotEmpty tries to map InputEncryptedFile to NotEmptyInputEncryptedFile.
func (i *InputEncryptedFile) AsNotEmpty() (NotEmptyInputEncryptedFile, bool) {
	value, ok := (InputEncryptedFileClass(i)).(NotEmptyInputEncryptedFile)
	return value, ok
}

// AsNotEmpty tries to map InputEncryptedFileBigUploaded to NotEmptyInputEncryptedFile.
func (i *InputEncryptedFileBigUploaded) AsNotEmpty() (NotEmptyInputEncryptedFile, bool) {
	value, ok := (InputEncryptedFileClass(i)).(NotEmptyInputEncryptedFile)
	return value, ok
}

// DecodeInputEncryptedFile implements binary de-serialization for InputEncryptedFileClass.
func DecodeInputEncryptedFile(buf *bin.Buffer) (InputEncryptedFileClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputEncryptedFileEmptyTypeID:
		// Decoding inputEncryptedFileEmpty#1837c364.
		v := InputEncryptedFileEmpty{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputEncryptedFileClass: %w", err)
		}
		return &v, nil
	case InputEncryptedFileUploadedTypeID:
		// Decoding inputEncryptedFileUploaded#64bd0306.
		v := InputEncryptedFileUploaded{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputEncryptedFileClass: %w", err)
		}
		return &v, nil
	case InputEncryptedFileTypeID:
		// Decoding inputEncryptedFile#5a17b5e5.
		v := InputEncryptedFile{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputEncryptedFileClass: %w", err)
		}
		return &v, nil
	case InputEncryptedFileBigUploadedTypeID:
		// Decoding inputEncryptedFileBigUploaded#2dc173c8.
		v := InputEncryptedFileBigUploaded{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode InputEncryptedFileClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode InputEncryptedFileClass: %w", bin.NewUnexpectedID(id))
	}
}

// InputEncryptedFile boxes the InputEncryptedFileClass providing a helper.
type InputEncryptedFileBox struct {
	InputEncryptedFile InputEncryptedFileClass
}

// Decode implements bin.Decoder for InputEncryptedFileBox.
func (b *InputEncryptedFileBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode InputEncryptedFileBox to nil")
	}
	v, err := DecodeInputEncryptedFile(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.InputEncryptedFile = v
	return nil
}

// Encode implements bin.Encode for InputEncryptedFileBox.
func (b *InputEncryptedFileBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.InputEncryptedFile == nil {
		return fmt.Errorf("unable to encode InputEncryptedFileClass as nil")
	}
	return b.InputEncryptedFile.Encode(buf)
}
