//go:build !no_gotd_slices
// +build !no_gotd_slices

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

// SendMessageActionClassArray is adapter for slice of SendMessageActionClass.
type SendMessageActionClassArray []SendMessageActionClass

// Sort sorts slice of SendMessageActionClass.
func (s SendMessageActionClassArray) Sort(less func(a, b SendMessageActionClass) bool) SendMessageActionClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SendMessageActionClass.
func (s SendMessageActionClassArray) SortStable(less func(a, b SendMessageActionClass) bool) SendMessageActionClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SendMessageActionClass.
func (s SendMessageActionClassArray) Retain(keep func(x SendMessageActionClass) bool) SendMessageActionClassArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s SendMessageActionClassArray) First() (v SendMessageActionClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SendMessageActionClassArray) Last() (v SendMessageActionClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SendMessageActionClassArray) PopFirst() (v SendMessageActionClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SendMessageActionClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SendMessageActionClassArray) Pop() (v SendMessageActionClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsSendMessageUploadVideoAction returns copy with only SendMessageUploadVideoAction constructors.
func (s SendMessageActionClassArray) AsSendMessageUploadVideoAction() (to SendMessageUploadVideoActionArray) {
	for _, elem := range s {
		value, ok := elem.(*SendMessageUploadVideoAction)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsSendMessageUploadAudioAction returns copy with only SendMessageUploadAudioAction constructors.
func (s SendMessageActionClassArray) AsSendMessageUploadAudioAction() (to SendMessageUploadAudioActionArray) {
	for _, elem := range s {
		value, ok := elem.(*SendMessageUploadAudioAction)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsSendMessageUploadPhotoAction returns copy with only SendMessageUploadPhotoAction constructors.
func (s SendMessageActionClassArray) AsSendMessageUploadPhotoAction() (to SendMessageUploadPhotoActionArray) {
	for _, elem := range s {
		value, ok := elem.(*SendMessageUploadPhotoAction)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsSendMessageUploadDocumentAction returns copy with only SendMessageUploadDocumentAction constructors.
func (s SendMessageActionClassArray) AsSendMessageUploadDocumentAction() (to SendMessageUploadDocumentActionArray) {
	for _, elem := range s {
		value, ok := elem.(*SendMessageUploadDocumentAction)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsSendMessageUploadRoundAction returns copy with only SendMessageUploadRoundAction constructors.
func (s SendMessageActionClassArray) AsSendMessageUploadRoundAction() (to SendMessageUploadRoundActionArray) {
	for _, elem := range s {
		value, ok := elem.(*SendMessageUploadRoundAction)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsSendMessageHistoryImportAction returns copy with only SendMessageHistoryImportAction constructors.
func (s SendMessageActionClassArray) AsSendMessageHistoryImportAction() (to SendMessageHistoryImportActionArray) {
	for _, elem := range s {
		value, ok := elem.(*SendMessageHistoryImportAction)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// SendMessageUploadVideoActionArray is adapter for slice of SendMessageUploadVideoAction.
type SendMessageUploadVideoActionArray []SendMessageUploadVideoAction

// Sort sorts slice of SendMessageUploadVideoAction.
func (s SendMessageUploadVideoActionArray) Sort(less func(a, b SendMessageUploadVideoAction) bool) SendMessageUploadVideoActionArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SendMessageUploadVideoAction.
func (s SendMessageUploadVideoActionArray) SortStable(less func(a, b SendMessageUploadVideoAction) bool) SendMessageUploadVideoActionArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SendMessageUploadVideoAction.
func (s SendMessageUploadVideoActionArray) Retain(keep func(x SendMessageUploadVideoAction) bool) SendMessageUploadVideoActionArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s SendMessageUploadVideoActionArray) First() (v SendMessageUploadVideoAction, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SendMessageUploadVideoActionArray) Last() (v SendMessageUploadVideoAction, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SendMessageUploadVideoActionArray) PopFirst() (v SendMessageUploadVideoAction, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SendMessageUploadVideoAction
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SendMessageUploadVideoActionArray) Pop() (v SendMessageUploadVideoAction, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SendMessageUploadAudioActionArray is adapter for slice of SendMessageUploadAudioAction.
type SendMessageUploadAudioActionArray []SendMessageUploadAudioAction

// Sort sorts slice of SendMessageUploadAudioAction.
func (s SendMessageUploadAudioActionArray) Sort(less func(a, b SendMessageUploadAudioAction) bool) SendMessageUploadAudioActionArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SendMessageUploadAudioAction.
func (s SendMessageUploadAudioActionArray) SortStable(less func(a, b SendMessageUploadAudioAction) bool) SendMessageUploadAudioActionArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SendMessageUploadAudioAction.
func (s SendMessageUploadAudioActionArray) Retain(keep func(x SendMessageUploadAudioAction) bool) SendMessageUploadAudioActionArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s SendMessageUploadAudioActionArray) First() (v SendMessageUploadAudioAction, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SendMessageUploadAudioActionArray) Last() (v SendMessageUploadAudioAction, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SendMessageUploadAudioActionArray) PopFirst() (v SendMessageUploadAudioAction, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SendMessageUploadAudioAction
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SendMessageUploadAudioActionArray) Pop() (v SendMessageUploadAudioAction, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SendMessageUploadPhotoActionArray is adapter for slice of SendMessageUploadPhotoAction.
type SendMessageUploadPhotoActionArray []SendMessageUploadPhotoAction

// Sort sorts slice of SendMessageUploadPhotoAction.
func (s SendMessageUploadPhotoActionArray) Sort(less func(a, b SendMessageUploadPhotoAction) bool) SendMessageUploadPhotoActionArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SendMessageUploadPhotoAction.
func (s SendMessageUploadPhotoActionArray) SortStable(less func(a, b SendMessageUploadPhotoAction) bool) SendMessageUploadPhotoActionArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SendMessageUploadPhotoAction.
func (s SendMessageUploadPhotoActionArray) Retain(keep func(x SendMessageUploadPhotoAction) bool) SendMessageUploadPhotoActionArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s SendMessageUploadPhotoActionArray) First() (v SendMessageUploadPhotoAction, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SendMessageUploadPhotoActionArray) Last() (v SendMessageUploadPhotoAction, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SendMessageUploadPhotoActionArray) PopFirst() (v SendMessageUploadPhotoAction, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SendMessageUploadPhotoAction
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SendMessageUploadPhotoActionArray) Pop() (v SendMessageUploadPhotoAction, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SendMessageUploadDocumentActionArray is adapter for slice of SendMessageUploadDocumentAction.
type SendMessageUploadDocumentActionArray []SendMessageUploadDocumentAction

// Sort sorts slice of SendMessageUploadDocumentAction.
func (s SendMessageUploadDocumentActionArray) Sort(less func(a, b SendMessageUploadDocumentAction) bool) SendMessageUploadDocumentActionArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SendMessageUploadDocumentAction.
func (s SendMessageUploadDocumentActionArray) SortStable(less func(a, b SendMessageUploadDocumentAction) bool) SendMessageUploadDocumentActionArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SendMessageUploadDocumentAction.
func (s SendMessageUploadDocumentActionArray) Retain(keep func(x SendMessageUploadDocumentAction) bool) SendMessageUploadDocumentActionArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s SendMessageUploadDocumentActionArray) First() (v SendMessageUploadDocumentAction, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SendMessageUploadDocumentActionArray) Last() (v SendMessageUploadDocumentAction, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SendMessageUploadDocumentActionArray) PopFirst() (v SendMessageUploadDocumentAction, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SendMessageUploadDocumentAction
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SendMessageUploadDocumentActionArray) Pop() (v SendMessageUploadDocumentAction, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SendMessageUploadRoundActionArray is adapter for slice of SendMessageUploadRoundAction.
type SendMessageUploadRoundActionArray []SendMessageUploadRoundAction

// Sort sorts slice of SendMessageUploadRoundAction.
func (s SendMessageUploadRoundActionArray) Sort(less func(a, b SendMessageUploadRoundAction) bool) SendMessageUploadRoundActionArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SendMessageUploadRoundAction.
func (s SendMessageUploadRoundActionArray) SortStable(less func(a, b SendMessageUploadRoundAction) bool) SendMessageUploadRoundActionArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SendMessageUploadRoundAction.
func (s SendMessageUploadRoundActionArray) Retain(keep func(x SendMessageUploadRoundAction) bool) SendMessageUploadRoundActionArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s SendMessageUploadRoundActionArray) First() (v SendMessageUploadRoundAction, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SendMessageUploadRoundActionArray) Last() (v SendMessageUploadRoundAction, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SendMessageUploadRoundActionArray) PopFirst() (v SendMessageUploadRoundAction, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SendMessageUploadRoundAction
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SendMessageUploadRoundActionArray) Pop() (v SendMessageUploadRoundAction, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// SendMessageHistoryImportActionArray is adapter for slice of SendMessageHistoryImportAction.
type SendMessageHistoryImportActionArray []SendMessageHistoryImportAction

// Sort sorts slice of SendMessageHistoryImportAction.
func (s SendMessageHistoryImportActionArray) Sort(less func(a, b SendMessageHistoryImportAction) bool) SendMessageHistoryImportActionArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of SendMessageHistoryImportAction.
func (s SendMessageHistoryImportActionArray) SortStable(less func(a, b SendMessageHistoryImportAction) bool) SendMessageHistoryImportActionArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of SendMessageHistoryImportAction.
func (s SendMessageHistoryImportActionArray) Retain(keep func(x SendMessageHistoryImportAction) bool) SendMessageHistoryImportActionArray {
	n := 0
	for _, x := range s {
		if keep(x) {
			s[n] = x
			n++
		}
	}
	s = s[:n]

	return s
}

// First returns first element of slice (if exists).
func (s SendMessageHistoryImportActionArray) First() (v SendMessageHistoryImportAction, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s SendMessageHistoryImportActionArray) Last() (v SendMessageHistoryImportAction, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *SendMessageHistoryImportActionArray) PopFirst() (v SendMessageHistoryImportAction, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero SendMessageHistoryImportAction
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *SendMessageHistoryImportActionArray) Pop() (v SendMessageHistoryImportAction, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
