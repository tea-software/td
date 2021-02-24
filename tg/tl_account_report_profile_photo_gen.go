// Code generated by gotdgen, DO NOT EDIT.

package tg

import (
	"context"
	"errors"
	"fmt"
	"strings"

	"github.com/gotd/td/bin"
)

// No-op definition for keeping imports.
var _ = bin.Buffer{}
var _ = context.Background()
var _ = fmt.Stringer(nil)
var _ = strings.Builder{}
var _ = errors.Is

// AccountReportProfilePhotoRequest represents TL type `account.reportProfilePhoto#fa8cc6f5`.
//
// See https://core.telegram.org/method/account.reportProfilePhoto for reference.
type AccountReportProfilePhotoRequest struct {
	// Peer field of AccountReportProfilePhotoRequest.
	Peer InputPeerClass `tl:"peer"`
	// PhotoID field of AccountReportProfilePhotoRequest.
	PhotoID InputPhotoClass `tl:"photo_id"`
	// Reason field of AccountReportProfilePhotoRequest.
	Reason ReportReasonClass `tl:"reason"`
	// Message field of AccountReportProfilePhotoRequest.
	Message string `tl:"message"`
}

// AccountReportProfilePhotoRequestTypeID is TL type id of AccountReportProfilePhotoRequest.
const AccountReportProfilePhotoRequestTypeID = 0xfa8cc6f5

func (r *AccountReportProfilePhotoRequest) Zero() bool {
	if r == nil {
		return true
	}
	if !(r.Peer == nil) {
		return false
	}
	if !(r.PhotoID == nil) {
		return false
	}
	if !(r.Reason == nil) {
		return false
	}
	if !(r.Message == "") {
		return false
	}

	return true
}

// String implements fmt.Stringer.
func (r *AccountReportProfilePhotoRequest) String() string {
	if r == nil {
		return "AccountReportProfilePhotoRequest(nil)"
	}
	type Alias AccountReportProfilePhotoRequest
	return fmt.Sprintf("AccountReportProfilePhotoRequest%+v", Alias(*r))
}

// FillFrom fills AccountReportProfilePhotoRequest from given interface.
func (r *AccountReportProfilePhotoRequest) FillFrom(from interface {
	GetPeer() (value InputPeerClass)
	GetPhotoID() (value InputPhotoClass)
	GetReason() (value ReportReasonClass)
	GetMessage() (value string)
}) {
	r.Peer = from.GetPeer()
	r.PhotoID = from.GetPhotoID()
	r.Reason = from.GetReason()
	r.Message = from.GetMessage()
}

// TypeID returns type id in TL schema.
//
// See https://core.telegram.org/mtproto/TL-tl#remarks.
func (r *AccountReportProfilePhotoRequest) TypeID() uint32 {
	return AccountReportProfilePhotoRequestTypeID
}

// TypeName returns name of type in TL schema.
func (r *AccountReportProfilePhotoRequest) TypeName() string {
	return "account.reportProfilePhoto"
}

// Encode implements bin.Encoder.
func (r *AccountReportProfilePhotoRequest) Encode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't encode account.reportProfilePhoto#fa8cc6f5 as nil")
	}
	b.PutID(AccountReportProfilePhotoRequestTypeID)
	if r.Peer == nil {
		return fmt.Errorf("unable to encode account.reportProfilePhoto#fa8cc6f5: field peer is nil")
	}
	if err := r.Peer.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.reportProfilePhoto#fa8cc6f5: field peer: %w", err)
	}
	if r.PhotoID == nil {
		return fmt.Errorf("unable to encode account.reportProfilePhoto#fa8cc6f5: field photo_id is nil")
	}
	if err := r.PhotoID.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.reportProfilePhoto#fa8cc6f5: field photo_id: %w", err)
	}
	if r.Reason == nil {
		return fmt.Errorf("unable to encode account.reportProfilePhoto#fa8cc6f5: field reason is nil")
	}
	if err := r.Reason.Encode(b); err != nil {
		return fmt.Errorf("unable to encode account.reportProfilePhoto#fa8cc6f5: field reason: %w", err)
	}
	b.PutString(r.Message)
	return nil
}

// GetPeer returns value of Peer field.
func (r *AccountReportProfilePhotoRequest) GetPeer() (value InputPeerClass) {
	return r.Peer
}

// GetPhotoID returns value of PhotoID field.
func (r *AccountReportProfilePhotoRequest) GetPhotoID() (value InputPhotoClass) {
	return r.PhotoID
}

// GetPhotoIDAsNotEmpty returns mapped value of PhotoID field.
func (r *AccountReportProfilePhotoRequest) GetPhotoIDAsNotEmpty() (*InputPhoto, bool) {
	return r.PhotoID.AsNotEmpty()
}

// GetReason returns value of Reason field.
func (r *AccountReportProfilePhotoRequest) GetReason() (value ReportReasonClass) {
	return r.Reason
}

// GetMessage returns value of Message field.
func (r *AccountReportProfilePhotoRequest) GetMessage() (value string) {
	return r.Message
}

// Decode implements bin.Decoder.
func (r *AccountReportProfilePhotoRequest) Decode(b *bin.Buffer) error {
	if r == nil {
		return fmt.Errorf("can't decode account.reportProfilePhoto#fa8cc6f5 to nil")
	}
	if err := b.ConsumeID(AccountReportProfilePhotoRequestTypeID); err != nil {
		return fmt.Errorf("unable to decode account.reportProfilePhoto#fa8cc6f5: %w", err)
	}
	{
		value, err := DecodeInputPeer(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.reportProfilePhoto#fa8cc6f5: field peer: %w", err)
		}
		r.Peer = value
	}
	{
		value, err := DecodeInputPhoto(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.reportProfilePhoto#fa8cc6f5: field photo_id: %w", err)
		}
		r.PhotoID = value
	}
	{
		value, err := DecodeReportReason(b)
		if err != nil {
			return fmt.Errorf("unable to decode account.reportProfilePhoto#fa8cc6f5: field reason: %w", err)
		}
		r.Reason = value
	}
	{
		value, err := b.String()
		if err != nil {
			return fmt.Errorf("unable to decode account.reportProfilePhoto#fa8cc6f5: field message: %w", err)
		}
		r.Message = value
	}
	return nil
}

// Ensuring interfaces in compile-time for AccountReportProfilePhotoRequest.
var (
	_ bin.Encoder = &AccountReportProfilePhotoRequest{}
	_ bin.Decoder = &AccountReportProfilePhotoRequest{}
)

// AccountReportProfilePhoto invokes method account.reportProfilePhoto#fa8cc6f5 returning error if any.
//
// See https://core.telegram.org/method/account.reportProfilePhoto for reference.
func (c *Client) AccountReportProfilePhoto(ctx context.Context, request *AccountReportProfilePhotoRequest) (bool, error) {
	var result BoolBox

	if err := c.rpc.InvokeRaw(ctx, request, &result); err != nil {
		return false, err
	}
	_, ok := result.Bool.(*BoolTrue)
	return ok, nil
}