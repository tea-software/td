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

// InputReportReasonSpam represents TL type `inputReportReasonSpam#58dbcab8`.
// Report for spam
//
// See https://core.telegram.org/constructor/inputReportReasonSpam for reference.
type InputReportReasonSpam struct {
}

// InputReportReasonSpamTypeID is TL type id of InputReportReasonSpam.
const InputReportReasonSpamTypeID = 0x58dbcab8

// construct implements constructor of ReportReasonClass.
func (i InputReportReasonSpam) construct() ReportReasonClass { return &i }

// Ensuring interfaces in compile-time for InputReportReasonSpam.
var (
	_ bin.Encoder     = &InputReportReasonSpam{}
	_ bin.Decoder     = &InputReportReasonSpam{}
	_ bin.BareEncoder = &InputReportReasonSpam{}
	_ bin.BareDecoder = &InputReportReasonSpam{}

	_ ReportReasonClass = &InputReportReasonSpam{}
)

func (i *InputReportReasonSpam) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputReportReasonSpam) String() string {
	if i == nil {
		return "InputReportReasonSpam(nil)"
	}
	type Alias InputReportReasonSpam
	return fmt.Sprintf("InputReportReasonSpam%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputReportReasonSpam) TypeID() uint32 {
	return InputReportReasonSpamTypeID
}

// TypeName returns name of type in TL schema.
func (*InputReportReasonSpam) TypeName() string {
	return "inputReportReasonSpam"
}

// TypeInfo returns info about TL type.
func (i *InputReportReasonSpam) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputReportReasonSpam",
		ID:   InputReportReasonSpamTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputReportReasonSpam) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReportReasonSpam#58dbcab8 as nil")
	}
	b.PutID(InputReportReasonSpamTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputReportReasonSpam) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReportReasonSpam#58dbcab8 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputReportReasonSpam) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReportReasonSpam#58dbcab8 to nil")
	}
	if err := b.ConsumeID(InputReportReasonSpamTypeID); err != nil {
		return fmt.Errorf("unable to decode inputReportReasonSpam#58dbcab8: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputReportReasonSpam) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReportReasonSpam#58dbcab8 to nil")
	}
	return nil
}

// InputReportReasonViolence represents TL type `inputReportReasonViolence#1e22c78d`.
// Report for violence
//
// See https://core.telegram.org/constructor/inputReportReasonViolence for reference.
type InputReportReasonViolence struct {
}

// InputReportReasonViolenceTypeID is TL type id of InputReportReasonViolence.
const InputReportReasonViolenceTypeID = 0x1e22c78d

// construct implements constructor of ReportReasonClass.
func (i InputReportReasonViolence) construct() ReportReasonClass { return &i }

// Ensuring interfaces in compile-time for InputReportReasonViolence.
var (
	_ bin.Encoder     = &InputReportReasonViolence{}
	_ bin.Decoder     = &InputReportReasonViolence{}
	_ bin.BareEncoder = &InputReportReasonViolence{}
	_ bin.BareDecoder = &InputReportReasonViolence{}

	_ ReportReasonClass = &InputReportReasonViolence{}
)

func (i *InputReportReasonViolence) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputReportReasonViolence) String() string {
	if i == nil {
		return "InputReportReasonViolence(nil)"
	}
	type Alias InputReportReasonViolence
	return fmt.Sprintf("InputReportReasonViolence%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputReportReasonViolence) TypeID() uint32 {
	return InputReportReasonViolenceTypeID
}

// TypeName returns name of type in TL schema.
func (*InputReportReasonViolence) TypeName() string {
	return "inputReportReasonViolence"
}

// TypeInfo returns info about TL type.
func (i *InputReportReasonViolence) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputReportReasonViolence",
		ID:   InputReportReasonViolenceTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputReportReasonViolence) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReportReasonViolence#1e22c78d as nil")
	}
	b.PutID(InputReportReasonViolenceTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputReportReasonViolence) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReportReasonViolence#1e22c78d as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputReportReasonViolence) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReportReasonViolence#1e22c78d to nil")
	}
	if err := b.ConsumeID(InputReportReasonViolenceTypeID); err != nil {
		return fmt.Errorf("unable to decode inputReportReasonViolence#1e22c78d: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputReportReasonViolence) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReportReasonViolence#1e22c78d to nil")
	}
	return nil
}

// InputReportReasonPornography represents TL type `inputReportReasonPornography#2e59d922`.
// Report for pornography
//
// See https://core.telegram.org/constructor/inputReportReasonPornography for reference.
type InputReportReasonPornography struct {
}

// InputReportReasonPornographyTypeID is TL type id of InputReportReasonPornography.
const InputReportReasonPornographyTypeID = 0x2e59d922

// construct implements constructor of ReportReasonClass.
func (i InputReportReasonPornography) construct() ReportReasonClass { return &i }

// Ensuring interfaces in compile-time for InputReportReasonPornography.
var (
	_ bin.Encoder     = &InputReportReasonPornography{}
	_ bin.Decoder     = &InputReportReasonPornography{}
	_ bin.BareEncoder = &InputReportReasonPornography{}
	_ bin.BareDecoder = &InputReportReasonPornography{}

	_ ReportReasonClass = &InputReportReasonPornography{}
)

func (i *InputReportReasonPornography) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputReportReasonPornography) String() string {
	if i == nil {
		return "InputReportReasonPornography(nil)"
	}
	type Alias InputReportReasonPornography
	return fmt.Sprintf("InputReportReasonPornography%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputReportReasonPornography) TypeID() uint32 {
	return InputReportReasonPornographyTypeID
}

// TypeName returns name of type in TL schema.
func (*InputReportReasonPornography) TypeName() string {
	return "inputReportReasonPornography"
}

// TypeInfo returns info about TL type.
func (i *InputReportReasonPornography) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputReportReasonPornography",
		ID:   InputReportReasonPornographyTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputReportReasonPornography) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReportReasonPornography#2e59d922 as nil")
	}
	b.PutID(InputReportReasonPornographyTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputReportReasonPornography) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReportReasonPornography#2e59d922 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputReportReasonPornography) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReportReasonPornography#2e59d922 to nil")
	}
	if err := b.ConsumeID(InputReportReasonPornographyTypeID); err != nil {
		return fmt.Errorf("unable to decode inputReportReasonPornography#2e59d922: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputReportReasonPornography) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReportReasonPornography#2e59d922 to nil")
	}
	return nil
}

// InputReportReasonChildAbuse represents TL type `inputReportReasonChildAbuse#adf44ee3`.
// Report for child abuse
//
// See https://core.telegram.org/constructor/inputReportReasonChildAbuse for reference.
type InputReportReasonChildAbuse struct {
}

// InputReportReasonChildAbuseTypeID is TL type id of InputReportReasonChildAbuse.
const InputReportReasonChildAbuseTypeID = 0xadf44ee3

// construct implements constructor of ReportReasonClass.
func (i InputReportReasonChildAbuse) construct() ReportReasonClass { return &i }

// Ensuring interfaces in compile-time for InputReportReasonChildAbuse.
var (
	_ bin.Encoder     = &InputReportReasonChildAbuse{}
	_ bin.Decoder     = &InputReportReasonChildAbuse{}
	_ bin.BareEncoder = &InputReportReasonChildAbuse{}
	_ bin.BareDecoder = &InputReportReasonChildAbuse{}

	_ ReportReasonClass = &InputReportReasonChildAbuse{}
)

func (i *InputReportReasonChildAbuse) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputReportReasonChildAbuse) String() string {
	if i == nil {
		return "InputReportReasonChildAbuse(nil)"
	}
	type Alias InputReportReasonChildAbuse
	return fmt.Sprintf("InputReportReasonChildAbuse%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputReportReasonChildAbuse) TypeID() uint32 {
	return InputReportReasonChildAbuseTypeID
}

// TypeName returns name of type in TL schema.
func (*InputReportReasonChildAbuse) TypeName() string {
	return "inputReportReasonChildAbuse"
}

// TypeInfo returns info about TL type.
func (i *InputReportReasonChildAbuse) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputReportReasonChildAbuse",
		ID:   InputReportReasonChildAbuseTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputReportReasonChildAbuse) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReportReasonChildAbuse#adf44ee3 as nil")
	}
	b.PutID(InputReportReasonChildAbuseTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputReportReasonChildAbuse) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReportReasonChildAbuse#adf44ee3 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputReportReasonChildAbuse) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReportReasonChildAbuse#adf44ee3 to nil")
	}
	if err := b.ConsumeID(InputReportReasonChildAbuseTypeID); err != nil {
		return fmt.Errorf("unable to decode inputReportReasonChildAbuse#adf44ee3: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputReportReasonChildAbuse) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReportReasonChildAbuse#adf44ee3 to nil")
	}
	return nil
}

// InputReportReasonOther represents TL type `inputReportReasonOther#c1e4a2b1`.
// Other
//
// See https://core.telegram.org/constructor/inputReportReasonOther for reference.
type InputReportReasonOther struct {
}

// InputReportReasonOtherTypeID is TL type id of InputReportReasonOther.
const InputReportReasonOtherTypeID = 0xc1e4a2b1

// construct implements constructor of ReportReasonClass.
func (i InputReportReasonOther) construct() ReportReasonClass { return &i }

// Ensuring interfaces in compile-time for InputReportReasonOther.
var (
	_ bin.Encoder     = &InputReportReasonOther{}
	_ bin.Decoder     = &InputReportReasonOther{}
	_ bin.BareEncoder = &InputReportReasonOther{}
	_ bin.BareDecoder = &InputReportReasonOther{}

	_ ReportReasonClass = &InputReportReasonOther{}
)

func (i *InputReportReasonOther) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputReportReasonOther) String() string {
	if i == nil {
		return "InputReportReasonOther(nil)"
	}
	type Alias InputReportReasonOther
	return fmt.Sprintf("InputReportReasonOther%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputReportReasonOther) TypeID() uint32 {
	return InputReportReasonOtherTypeID
}

// TypeName returns name of type in TL schema.
func (*InputReportReasonOther) TypeName() string {
	return "inputReportReasonOther"
}

// TypeInfo returns info about TL type.
func (i *InputReportReasonOther) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputReportReasonOther",
		ID:   InputReportReasonOtherTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputReportReasonOther) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReportReasonOther#c1e4a2b1 as nil")
	}
	b.PutID(InputReportReasonOtherTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputReportReasonOther) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReportReasonOther#c1e4a2b1 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputReportReasonOther) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReportReasonOther#c1e4a2b1 to nil")
	}
	if err := b.ConsumeID(InputReportReasonOtherTypeID); err != nil {
		return fmt.Errorf("unable to decode inputReportReasonOther#c1e4a2b1: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputReportReasonOther) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReportReasonOther#c1e4a2b1 to nil")
	}
	return nil
}

// InputReportReasonCopyright represents TL type `inputReportReasonCopyright#9b89f93a`.
// Report for copyrighted content
//
// See https://core.telegram.org/constructor/inputReportReasonCopyright for reference.
type InputReportReasonCopyright struct {
}

// InputReportReasonCopyrightTypeID is TL type id of InputReportReasonCopyright.
const InputReportReasonCopyrightTypeID = 0x9b89f93a

// construct implements constructor of ReportReasonClass.
func (i InputReportReasonCopyright) construct() ReportReasonClass { return &i }

// Ensuring interfaces in compile-time for InputReportReasonCopyright.
var (
	_ bin.Encoder     = &InputReportReasonCopyright{}
	_ bin.Decoder     = &InputReportReasonCopyright{}
	_ bin.BareEncoder = &InputReportReasonCopyright{}
	_ bin.BareDecoder = &InputReportReasonCopyright{}

	_ ReportReasonClass = &InputReportReasonCopyright{}
)

func (i *InputReportReasonCopyright) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputReportReasonCopyright) String() string {
	if i == nil {
		return "InputReportReasonCopyright(nil)"
	}
	type Alias InputReportReasonCopyright
	return fmt.Sprintf("InputReportReasonCopyright%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputReportReasonCopyright) TypeID() uint32 {
	return InputReportReasonCopyrightTypeID
}

// TypeName returns name of type in TL schema.
func (*InputReportReasonCopyright) TypeName() string {
	return "inputReportReasonCopyright"
}

// TypeInfo returns info about TL type.
func (i *InputReportReasonCopyright) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputReportReasonCopyright",
		ID:   InputReportReasonCopyrightTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputReportReasonCopyright) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReportReasonCopyright#9b89f93a as nil")
	}
	b.PutID(InputReportReasonCopyrightTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputReportReasonCopyright) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReportReasonCopyright#9b89f93a as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputReportReasonCopyright) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReportReasonCopyright#9b89f93a to nil")
	}
	if err := b.ConsumeID(InputReportReasonCopyrightTypeID); err != nil {
		return fmt.Errorf("unable to decode inputReportReasonCopyright#9b89f93a: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputReportReasonCopyright) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReportReasonCopyright#9b89f93a to nil")
	}
	return nil
}

// InputReportReasonGeoIrrelevant represents TL type `inputReportReasonGeoIrrelevant#dbd4feed`.
// Report an irrelevant geogroup
//
// See https://core.telegram.org/constructor/inputReportReasonGeoIrrelevant for reference.
type InputReportReasonGeoIrrelevant struct {
}

// InputReportReasonGeoIrrelevantTypeID is TL type id of InputReportReasonGeoIrrelevant.
const InputReportReasonGeoIrrelevantTypeID = 0xdbd4feed

// construct implements constructor of ReportReasonClass.
func (i InputReportReasonGeoIrrelevant) construct() ReportReasonClass { return &i }

// Ensuring interfaces in compile-time for InputReportReasonGeoIrrelevant.
var (
	_ bin.Encoder     = &InputReportReasonGeoIrrelevant{}
	_ bin.Decoder     = &InputReportReasonGeoIrrelevant{}
	_ bin.BareEncoder = &InputReportReasonGeoIrrelevant{}
	_ bin.BareDecoder = &InputReportReasonGeoIrrelevant{}

	_ ReportReasonClass = &InputReportReasonGeoIrrelevant{}
)

func (i *InputReportReasonGeoIrrelevant) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputReportReasonGeoIrrelevant) String() string {
	if i == nil {
		return "InputReportReasonGeoIrrelevant(nil)"
	}
	type Alias InputReportReasonGeoIrrelevant
	return fmt.Sprintf("InputReportReasonGeoIrrelevant%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputReportReasonGeoIrrelevant) TypeID() uint32 {
	return InputReportReasonGeoIrrelevantTypeID
}

// TypeName returns name of type in TL schema.
func (*InputReportReasonGeoIrrelevant) TypeName() string {
	return "inputReportReasonGeoIrrelevant"
}

// TypeInfo returns info about TL type.
func (i *InputReportReasonGeoIrrelevant) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputReportReasonGeoIrrelevant",
		ID:   InputReportReasonGeoIrrelevantTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputReportReasonGeoIrrelevant) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReportReasonGeoIrrelevant#dbd4feed as nil")
	}
	b.PutID(InputReportReasonGeoIrrelevantTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputReportReasonGeoIrrelevant) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReportReasonGeoIrrelevant#dbd4feed as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputReportReasonGeoIrrelevant) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReportReasonGeoIrrelevant#dbd4feed to nil")
	}
	if err := b.ConsumeID(InputReportReasonGeoIrrelevantTypeID); err != nil {
		return fmt.Errorf("unable to decode inputReportReasonGeoIrrelevant#dbd4feed: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputReportReasonGeoIrrelevant) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReportReasonGeoIrrelevant#dbd4feed to nil")
	}
	return nil
}

// InputReportReasonFake represents TL type `inputReportReasonFake#f5ddd6e7`.
//
// See https://core.telegram.org/constructor/inputReportReasonFake for reference.
type InputReportReasonFake struct {
}

// InputReportReasonFakeTypeID is TL type id of InputReportReasonFake.
const InputReportReasonFakeTypeID = 0xf5ddd6e7

// construct implements constructor of ReportReasonClass.
func (i InputReportReasonFake) construct() ReportReasonClass { return &i }

// Ensuring interfaces in compile-time for InputReportReasonFake.
var (
	_ bin.Encoder     = &InputReportReasonFake{}
	_ bin.Decoder     = &InputReportReasonFake{}
	_ bin.BareEncoder = &InputReportReasonFake{}
	_ bin.BareDecoder = &InputReportReasonFake{}

	_ ReportReasonClass = &InputReportReasonFake{}
)

func (i *InputReportReasonFake) Zero() bool {
	if i == nil {
		return true
	}

	return true
}

// String implements fmt.Stringer.
func (i *InputReportReasonFake) String() string {
	if i == nil {
		return "InputReportReasonFake(nil)"
	}
	type Alias InputReportReasonFake
	return fmt.Sprintf("InputReportReasonFake%+v", Alias(*i))
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (*InputReportReasonFake) TypeID() uint32 {
	return InputReportReasonFakeTypeID
}

// TypeName returns name of type in TL schema.
func (*InputReportReasonFake) TypeName() string {
	return "inputReportReasonFake"
}

// TypeInfo returns info about TL type.
func (i *InputReportReasonFake) TypeInfo() tdp.Type {
	typ := tdp.Type{
		Name: "inputReportReasonFake",
		ID:   InputReportReasonFakeTypeID,
	}
	if i == nil {
		typ.Null = true
		return typ
	}
	typ.Fields = []tdp.Field{}
	return typ
}

// Encode implements bin.Encoder.
func (i *InputReportReasonFake) Encode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReportReasonFake#f5ddd6e7 as nil")
	}
	b.PutID(InputReportReasonFakeTypeID)
	return i.EncodeBare(b)
}

// EncodeBare implements bin.BareEncoder.
func (i *InputReportReasonFake) EncodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't encode inputReportReasonFake#f5ddd6e7 as nil")
	}
	return nil
}

// Decode implements bin.Decoder.
func (i *InputReportReasonFake) Decode(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReportReasonFake#f5ddd6e7 to nil")
	}
	if err := b.ConsumeID(InputReportReasonFakeTypeID); err != nil {
		return fmt.Errorf("unable to decode inputReportReasonFake#f5ddd6e7: %w", err)
	}
	return i.DecodeBare(b)
}

// DecodeBare implements bin.BareDecoder.
func (i *InputReportReasonFake) DecodeBare(b *bin.Buffer) error {
	if i == nil {
		return fmt.Errorf("can't decode inputReportReasonFake#f5ddd6e7 to nil")
	}
	return nil
}

// ReportReasonClass represents ReportReason generic type.
//
// See https://core.telegram.org/type/ReportReason for reference.
//
// Example:
//  g, err := tg.DecodeReportReason(buf)
//  if err != nil {
//      panic(err)
//  }
//  switch v := g.(type) {
//  case *tg.InputReportReasonSpam: // inputReportReasonSpam#58dbcab8
//  case *tg.InputReportReasonViolence: // inputReportReasonViolence#1e22c78d
//  case *tg.InputReportReasonPornography: // inputReportReasonPornography#2e59d922
//  case *tg.InputReportReasonChildAbuse: // inputReportReasonChildAbuse#adf44ee3
//  case *tg.InputReportReasonOther: // inputReportReasonOther#c1e4a2b1
//  case *tg.InputReportReasonCopyright: // inputReportReasonCopyright#9b89f93a
//  case *tg.InputReportReasonGeoIrrelevant: // inputReportReasonGeoIrrelevant#dbd4feed
//  case *tg.InputReportReasonFake: // inputReportReasonFake#f5ddd6e7
//  default: panic(v)
//  }
type ReportReasonClass interface {
	bin.Encoder
	bin.Decoder
	bin.BareEncoder
	bin.BareDecoder
	construct() ReportReasonClass

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

// DecodeReportReason implements binary de-serialization for ReportReasonClass.
func DecodeReportReason(buf *bin.Buffer) (ReportReasonClass, error) {
	id, err := buf.PeekID()
	if err != nil {
		return nil, err
	}
	switch id {
	case InputReportReasonSpamTypeID:
		// Decoding inputReportReasonSpam#58dbcab8.
		v := InputReportReasonSpam{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportReasonClass: %w", err)
		}
		return &v, nil
	case InputReportReasonViolenceTypeID:
		// Decoding inputReportReasonViolence#1e22c78d.
		v := InputReportReasonViolence{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportReasonClass: %w", err)
		}
		return &v, nil
	case InputReportReasonPornographyTypeID:
		// Decoding inputReportReasonPornography#2e59d922.
		v := InputReportReasonPornography{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportReasonClass: %w", err)
		}
		return &v, nil
	case InputReportReasonChildAbuseTypeID:
		// Decoding inputReportReasonChildAbuse#adf44ee3.
		v := InputReportReasonChildAbuse{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportReasonClass: %w", err)
		}
		return &v, nil
	case InputReportReasonOtherTypeID:
		// Decoding inputReportReasonOther#c1e4a2b1.
		v := InputReportReasonOther{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportReasonClass: %w", err)
		}
		return &v, nil
	case InputReportReasonCopyrightTypeID:
		// Decoding inputReportReasonCopyright#9b89f93a.
		v := InputReportReasonCopyright{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportReasonClass: %w", err)
		}
		return &v, nil
	case InputReportReasonGeoIrrelevantTypeID:
		// Decoding inputReportReasonGeoIrrelevant#dbd4feed.
		v := InputReportReasonGeoIrrelevant{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportReasonClass: %w", err)
		}
		return &v, nil
	case InputReportReasonFakeTypeID:
		// Decoding inputReportReasonFake#f5ddd6e7.
		v := InputReportReasonFake{}
		if err := v.Decode(buf); err != nil {
			return nil, fmt.Errorf("unable to decode ReportReasonClass: %w", err)
		}
		return &v, nil
	default:
		return nil, fmt.Errorf("unable to decode ReportReasonClass: %w", bin.NewUnexpectedID(id))
	}
}

// ReportReason boxes the ReportReasonClass providing a helper.
type ReportReasonBox struct {
	ReportReason ReportReasonClass
}

// Decode implements bin.Decoder for ReportReasonBox.
func (b *ReportReasonBox) Decode(buf *bin.Buffer) error {
	if b == nil {
		return fmt.Errorf("unable to decode ReportReasonBox to nil")
	}
	v, err := DecodeReportReason(buf)
	if err != nil {
		return fmt.Errorf("unable to decode boxed value: %w", err)
	}
	b.ReportReason = v
	return nil
}

// Encode implements bin.Encode for ReportReasonBox.
func (b *ReportReasonBox) Encode(buf *bin.Buffer) error {
	if b == nil || b.ReportReason == nil {
		return fmt.Errorf("unable to encode ReportReasonClass as nil")
	}
	return b.ReportReason.Encode(buf)
}
