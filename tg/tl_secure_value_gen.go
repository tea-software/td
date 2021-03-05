// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"sort"
	"strings"

	"github.com/gotd/td/bin"
	"github.com/gotd/td/tdp"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is
var _ = sort.Ints
var _ = tdp.Format

// SecureValue represents TL type `secureValue#187fa0ca`.
// Secure value
//
// See https://core.telegram.org/constructor/secureValue for reference.
type SecureValue struct {
	// Flags, see TL conditional fields¹
	//
	// Links:
	//  1) https://core.telegram.org/mtproto/TL-combinators#conditional-fields
	Flags bin.Fields
	// Secure passport¹ value type
	//
	// Links:
	//  1) https://core.telegram.org/passport
	Type SecureValueTypeClass
	// Encrypted Telegram Passport¹ element data
	//
	// Links:
	//  1) https://core.telegram.org/passport
	//
	// Use SetData and GetData helpers.
	Data SecureData
	// Encrypted passport¹ file with the front side of the document
	//
	// Links:
	//  1) https://core.telegram.org/passport
	//
	// Use SetFrontSide and GetFrontSide helpers.
	FrontSide SecureFileClass
	// Encrypted passport¹ file with the reverse side of the document
	//
	// Links:
	//  1) https://core.telegram.org/passport
	//
	// Use SetReverseSide and GetReverseSide helpers.
	ReverseSide SecureFileClass
	// Encrypted passport¹ file with a selfie of the user holding the document
	//
	// Links:
	//  1) https://core.telegram.org/passport
	//
	// Use SetSelfie and GetSelfie helpers.
	Selfie SecureFileClass
	// Array of encrypted passport¹ files with translated versions of the provided documents
	//
	// Links:
	//  1) https://core.telegram.org/passport
	//
	// Use SetTranslation and GetTranslation helpers.
	Translation []SecureFileClass
	// Array of encrypted passport¹ files with photos the of the documents
	//
	// Links:
	//  1) https://core.telegram.org/passport
	//
	// Use SetFiles and GetFiles helpers.
	Files []SecureFileClass
	// Plaintext verified passport¹ data
	//
	// Links:
	//  1) https://core.telegram.org/passport
	//
	// Use SetPlainData and GetPlainData helpers.
	PlainData SecurePlainDataClass
	// Data hash
	Hash []byte
}

// SecureValueTypeID is TL type id of SecureValue.
const SecureValueTypeID = 0x187fa0ca

func (s *SecureValue) Zero() bool {
	if s == nil {
		return true
	}
	if !(s.Flags.Zero()) {
		return false
	}
	if !(s.Type == nil) {
		return false
	}
	if !(s.Data.Zero()) {
		return false
	}
	if !(s.FrontSide == nil) {
		return false
	}
	if !(s.ReverseSide == nil) {
		return false
	}
	if !(s.Selfie == nil) {
		return false
	}
	if !(s.Translation == nil) {
		return false
	}
	if !(s.Files == nil) {
		return false
	}
	if !(s.PlainData == nil) {
		return false
	}
	if !(s.Hash == nil) {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (s *SecureValue) String() string {
	if s == nil {
		return "SecureValue(nil)"
	}
	type Alias SecureValue
	return fmt.Sprintf("SecureValue%+v", Alias(*s))
}

// FillFrom fills SecureValue from given interface.
func (s *SecureValue) FillFrom(from interface {
	GetType() (value SecureValueTypeClass)
	GetData() (value SecureData, ok bool)
	GetFrontSide() (value SecureFileClass, ok bool)
	GetReverseSide() (value SecureFileClass, ok bool)
	GetSelfie() (value SecureFileClass, ok bool)
	GetTranslation() (value []SecureFileClass, ok bool)
	GetFiles() (value []SecureFileClass, ok bool)
	GetPlainData() (value SecurePlainDataClass, ok bool)
	GetHash() (value []byte)
}) {
	s.Type = from.GetType()
	if val, ok := from.GetData(); ok {
		s.Data = val
	}

	if val, ok := from.GetFrontSide(); ok {
		s.FrontSide = val
	}

	if val, ok := from.GetReverseSide(); ok {
		s.ReverseSide = val
	}

	if val, ok := from.GetSelfie(); ok {
		s.Selfie = val
	}

	if val, ok := from.GetTranslation(); ok {
		s.Translation = val
	}

	if val, ok := from.GetFiles(); ok {
		s.Files = val
	}

	if val, ok := from.GetPlainData(); ok {
		s.PlainData = val
	}

	s.Hash = from.GetHash()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*SecureValue) TypeID() uint32 {
	return SecureValueTypeID
}

// TypeName returns name of type in TL schema.
func (*SecureValue) TypeName() string {
	return "secureValue"
}

// TypeInfo returns info about TL type.
func (s *SecureValue) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "secureValue",
		ID:   SecureValueTypeID,
	}
	if s == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{
		{
			Name:       "Type",
			SchemaName: "type",
		},
		{
			Name:       "Data",
			SchemaName: "data",
			Null:       !s.Flags.Has(0),
		},
		{
			Name:       "FrontSide",
			SchemaName: "front_side",
			Null:       !s.Flags.Has(1),
		},
		{
			Name:       "ReverseSide",
			SchemaName: "reverse_side",
			Null:       !s.Flags.Has(2),
		},
		{
			Name:       "Selfie",
			SchemaName: "selfie",
			Null:       !s.Flags.Has(3),
		},
		{
			Name:       "Translation",
			SchemaName: "translation",
			Null:       !s.Flags.Has(6),
		},
		{
			Name:       "Files",
			SchemaName: "files",
			Null:       !s.Flags.Has(4),
		},
		{
			Name:       "PlainData",
			SchemaName: "plain_data",
			Null:       !s.Flags.Has(5),
		},
		{
			Name:       "Hash",
			SchemaName: "hash",
		},
	}
	return typ
}

// Encode implements bin.Encoder.
func (s *SecureValue) Encode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't encode secureValue#187fa0ca as nil")
	}
	b.PutID(SecureValueTypeID)
	if !(s.Data.Zero()) {
		s.Flags.Set(0)
	}
	if !(s.FrontSide == nil) {
		s.Flags.Set(1)
	}
	if !(s.ReverseSide == nil) {
		s.Flags.Set(2)
	}
	if !(s.Selfie == nil) {
		s.Flags.Set(3)
	}
	if !(s.Translation == nil) {
		s.Flags.Set(6)
	}
	if !(s.Files == nil) {
		s.Flags.Set(4)
	}
	if !(s.PlainData == nil) {
		s.Flags.Set(5)
	}
	if err := s.Flags.Encode(b); err != nil {
		return fmt.Errorf("unable to encode secureValue#187fa0ca: field flags: %w", err)
	}
	if s.Type == nil {
		return fmt.Errorf("unable to encode secureValue#187fa0ca: field type is nil")
	}
	if err := s.Type.Encode(b); err != nil {
		return fmt.Errorf("unable to encode secureValue#187fa0ca: field type: %w", err)
	}
	if s.Flags.Has(0) {
		if err := s.Data.Encode(b); err != nil {
			return fmt.Errorf("unable to encode secureValue#187fa0ca: field data: %w", err)
		}
	}
	if s.Flags.Has(1) {
		if s.FrontSide == nil {
			return fmt.Errorf("unable to encode secureValue#187fa0ca: field front_side is nil")
		}
		if err := s.FrontSide.Encode(b); err != nil {
			return fmt.Errorf("unable to encode secureValue#187fa0ca: field front_side: %w", err)
		}
	}
	if s.Flags.Has(2) {
		if s.ReverseSide == nil {
			return fmt.Errorf("unable to encode secureValue#187fa0ca: field reverse_side is nil")
		}
		if err := s.ReverseSide.Encode(b); err != nil {
			return fmt.Errorf("unable to encode secureValue#187fa0ca: field reverse_side: %w", err)
		}
	}
	if s.Flags.Has(3) {
		if s.Selfie == nil {
			return fmt.Errorf("unable to encode secureValue#187fa0ca: field selfie is nil")
		}
		if err := s.Selfie.Encode(b); err != nil {
			return fmt.Errorf("unable to encode secureValue#187fa0ca: field selfie: %w", err)
		}
	}
	if s.Flags.Has(6) {
		b.PutVectorHeader(len(s.Translation))
		for idx, v := range s.Translation {
			if v == nil {
				return fmt.Errorf("unable to encode secureValue#187fa0ca: field translation element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode secureValue#187fa0ca: field translation element with index %d: %w", idx, err)
			}
		}
	}
	if s.Flags.Has(4) {
		b.PutVectorHeader(len(s.Files))
		for idx, v := range s.Files {
			if v == nil {
				return fmt.Errorf("unable to encode secureValue#187fa0ca: field files element with index %d is nil", idx)
			}
			if err := v.Encode(b); err != nil {
				return fmt.Errorf("unable to encode secureValue#187fa0ca: field files element with index %d: %w", idx, err)
			}
		}
	}
	if s.Flags.Has(5) {
		if s.PlainData == nil {
			return fmt.Errorf("unable to encode secureValue#187fa0ca: field plain_data is nil")
		}
		if err := s.PlainData.Encode(b); err != nil {
			return fmt.Errorf("unable to encode secureValue#187fa0ca: field plain_data: %w", err)
		}
	}
	b.PutBytes(s.Hash)
	return nil
}

// GetType returns value of Type field.
func (s *SecureValue) GetType() (value SecureValueTypeClass) {
	return s.Type
}

// SetData sets value of Data conditional field.
func (s *SecureValue) SetData(value SecureData) {
	s.Flags.Set(0)
	s.Data = value
}

// GetData returns value of Data conditional field and
// boolean which is true if field was set.
func (s *SecureValue) GetData() (value SecureData, ok bool) {
	if !s.Flags.Has(0) {
		return value, false
	}
	return s.Data, true
}

// SetFrontSide sets value of FrontSide conditional field.
func (s *SecureValue) SetFrontSide(value SecureFileClass) {
	s.Flags.Set(1)
	s.FrontSide = value
}

// GetFrontSide returns value of FrontSide conditional field and
// boolean which is true if field was set.
func (s *SecureValue) GetFrontSide() (value SecureFileClass, ok bool) {
	if !s.Flags.Has(1) {
		return value, false
	}
	return s.FrontSide, true
}

// GetFrontSideAsNotEmpty returns mapped value of FrontSide conditional field and
// boolean which is true if field was set.
func (s *SecureValue) GetFrontSideAsNotEmpty() (*SecureFile, bool) {
	if value, ok := s.GetFrontSide(); ok {
		return value.AsNotEmpty()
	}
	return nil, false
}

// SetReverseSide sets value of ReverseSide conditional field.
func (s *SecureValue) SetReverseSide(value SecureFileClass) {
	s.Flags.Set(2)
	s.ReverseSide = value
}

// GetReverseSide returns value of ReverseSide conditional field and
// boolean which is true if field was set.
func (s *SecureValue) GetReverseSide() (value SecureFileClass, ok bool) {
	if !s.Flags.Has(2) {
		return value, false
	}
	return s.ReverseSide, true
}

// GetReverseSideAsNotEmpty returns mapped value of ReverseSide conditional field and
// boolean which is true if field was set.
func (s *SecureValue) GetReverseSideAsNotEmpty() (*SecureFile, bool) {
	if value, ok := s.GetReverseSide(); ok {
		return value.AsNotEmpty()
	}
	return nil, false
}

// SetSelfie sets value of Selfie conditional field.
func (s *SecureValue) SetSelfie(value SecureFileClass) {
	s.Flags.Set(3)
	s.Selfie = value
}

// GetSelfie returns value of Selfie conditional field and
// boolean which is true if field was set.
func (s *SecureValue) GetSelfie() (value SecureFileClass, ok bool) {
	if !s.Flags.Has(3) {
		return value, false
	}
	return s.Selfie, true
}

// GetSelfieAsNotEmpty returns mapped value of Selfie conditional field and
// boolean which is true if field was set.
func (s *SecureValue) GetSelfieAsNotEmpty() (*SecureFile, bool) {
	if value, ok := s.GetSelfie(); ok {
		return value.AsNotEmpty()
	}
	return nil, false
}

// SetTranslation sets value of Translation conditional field.
func (s *SecureValue) SetTranslation(value []SecureFileClass) {
	s.Flags.Set(6)
	s.Translation = value
}

// GetTranslation returns value of Translation conditional field and
// boolean which is true if field was set.
func (s *SecureValue) GetTranslation() (value []SecureFileClass, ok bool) {
	if !s.Flags.Has(6) {
		return value, false
	}
	return s.Translation, true
}

// MapTranslation returns field Translation wrapped in SecureFileClassArray helper.
func (s *SecureValue) MapTranslation() (value SecureFileClassArray, ok bool) {
	if !s.Flags.Has(6) {
		return value, false
	}
	return SecureFileClassArray(s.Translation), true
}

// SetFiles sets value of Files conditional field.
func (s *SecureValue) SetFiles(value []SecureFileClass) {
	s.Flags.Set(4)
	s.Files = value
}

// GetFiles returns value of Files conditional field and
// boolean which is true if field was set.
func (s *SecureValue) GetFiles() (value []SecureFileClass, ok bool) {
	if !s.Flags.Has(4) {
		return value, false
	}
	return s.Files, true
}

// MapFiles returns field Files wrapped in SecureFileClassArray helper.
func (s *SecureValue) MapFiles() (value SecureFileClassArray, ok bool) {
	if !s.Flags.Has(4) {
		return value, false
	}
	return SecureFileClassArray(s.Files), true
}

// SetPlainData sets value of PlainData conditional field.
func (s *SecureValue) SetPlainData(value SecurePlainDataClass) {
	s.Flags.Set(5)
	s.PlainData = value
}

// GetPlainData returns value of PlainData conditional field and
// boolean which is true if field was set.
func (s *SecureValue) GetPlainData() (value SecurePlainDataClass, ok bool) {
	if !s.Flags.Has(5) {
		return value, false
	}
	return s.PlainData, true
}

// GetHash returns value of Hash field.
func (s *SecureValue) GetHash() (value []byte) {
	return s.Hash
}

// Decode implements bin.Decoder.
func (s *SecureValue) Decode(b *bin.Buffer) error {
	if s == nil {
		return fmt.Errorf("can't decode secureValue#187fa0ca to nil")
	}
	if err := b.ConsumeID(SecureValueTypeID); err != nil {
		return fmt.Errorf("unable to decode secureValue#187fa0ca: %w", err)
	}
	{
		if err := s.Flags.Decode(b); err != nil {
			return fmt.Errorf("unable to decode secureValue#187fa0ca: field flags: %w", err)
		}
	}
	{
		value, err := DecodeSecureValueType(b)
		if err != nil {
			return fmt.Errorf("unable to decode secureValue#187fa0ca: field type: %w", err)
		}
		s.Type = value
	}
	if s.Flags.Has(0) {
		if err := s.Data.Decode(b); err != nil {
			return fmt.Errorf("unable to decode secureValue#187fa0ca: field data: %w", err)
		}
	}
	if s.Flags.Has(1) {
		value, err := DecodeSecureFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode secureValue#187fa0ca: field front_side: %w", err)
		}
		s.FrontSide = value
	}
	if s.Flags.Has(2) {
		value, err := DecodeSecureFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode secureValue#187fa0ca: field reverse_side: %w", err)
		}
		s.ReverseSide = value
	}
	if s.Flags.Has(3) {
		value, err := DecodeSecureFile(b)
		if err != nil {
			return fmt.Errorf("unable to decode secureValue#187fa0ca: field selfie: %w", err)
		}
		s.Selfie = value
	}
	if s.Flags.Has(6) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode secureValue#187fa0ca: field translation: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeSecureFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode secureValue#187fa0ca: field translation: %w", err)
			}
			s.Translation = append(s.Translation, value)
		}
	}
	if s.Flags.Has(4) {
		headerLen, err := b.VectorHeader()
		if err != nil {
			return fmt.Errorf("unable to decode secureValue#187fa0ca: field files: %w", err)
		}
		for idx := 0; idx < headerLen; idx++ {
			value, err := DecodeSecureFile(b)
			if err != nil {
				return fmt.Errorf("unable to decode secureValue#187fa0ca: field files: %w", err)
			}
			s.Files = append(s.Files, value)
		}
	}
	if s.Flags.Has(5) {
		value, err := DecodeSecurePlainData(b)
		if err != nil {
			return fmt.Errorf("unable to decode secureValue#187fa0ca: field plain_data: %w", err)
		}
		s.PlainData = value
	}
	{
		value, err := b.Bytes()
		if err != nil {
			return fmt.Errorf("unable to decode secureValue#187fa0ca: field hash: %w", err)
		}
		s.Hash = value
	}
	return nil
}

// Ensuring interfaces in compile-time for SecureValue.
var (
	_ bin.Encoder = &SecureValue{}
	_ bin.Decoder = &SecureValue{}
)
