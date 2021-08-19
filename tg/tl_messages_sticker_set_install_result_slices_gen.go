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

// MessagesStickerSetInstallResultClassArray is adapter for slice of MessagesStickerSetInstallResultClass.
type MessagesStickerSetInstallResultClassArray []MessagesStickerSetInstallResultClass

// Sort sorts slice of MessagesStickerSetInstallResultClass.
func (s MessagesStickerSetInstallResultClassArray) Sort(less func(a, b MessagesStickerSetInstallResultClass) bool) MessagesStickerSetInstallResultClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesStickerSetInstallResultClass.
func (s MessagesStickerSetInstallResultClassArray) SortStable(less func(a, b MessagesStickerSetInstallResultClass) bool) MessagesStickerSetInstallResultClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesStickerSetInstallResultClass.
func (s MessagesStickerSetInstallResultClassArray) Retain(keep func(x MessagesStickerSetInstallResultClass) bool) MessagesStickerSetInstallResultClassArray {
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
func (s MessagesStickerSetInstallResultClassArray) First() (v MessagesStickerSetInstallResultClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesStickerSetInstallResultClassArray) Last() (v MessagesStickerSetInstallResultClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesStickerSetInstallResultClassArray) PopFirst() (v MessagesStickerSetInstallResultClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesStickerSetInstallResultClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesStickerSetInstallResultClassArray) Pop() (v MessagesStickerSetInstallResultClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsMessagesStickerSetInstallResultArchive returns copy with only MessagesStickerSetInstallResultArchive constructors.
func (s MessagesStickerSetInstallResultClassArray) AsMessagesStickerSetInstallResultArchive() (to MessagesStickerSetInstallResultArchiveArray) {
	for _, elem := range s {
		value, ok := elem.(*MessagesStickerSetInstallResultArchive)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// MessagesStickerSetInstallResultArchiveArray is adapter for slice of MessagesStickerSetInstallResultArchive.
type MessagesStickerSetInstallResultArchiveArray []MessagesStickerSetInstallResultArchive

// Sort sorts slice of MessagesStickerSetInstallResultArchive.
func (s MessagesStickerSetInstallResultArchiveArray) Sort(less func(a, b MessagesStickerSetInstallResultArchive) bool) MessagesStickerSetInstallResultArchiveArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of MessagesStickerSetInstallResultArchive.
func (s MessagesStickerSetInstallResultArchiveArray) SortStable(less func(a, b MessagesStickerSetInstallResultArchive) bool) MessagesStickerSetInstallResultArchiveArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of MessagesStickerSetInstallResultArchive.
func (s MessagesStickerSetInstallResultArchiveArray) Retain(keep func(x MessagesStickerSetInstallResultArchive) bool) MessagesStickerSetInstallResultArchiveArray {
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
func (s MessagesStickerSetInstallResultArchiveArray) First() (v MessagesStickerSetInstallResultArchive, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s MessagesStickerSetInstallResultArchiveArray) Last() (v MessagesStickerSetInstallResultArchive, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *MessagesStickerSetInstallResultArchiveArray) PopFirst() (v MessagesStickerSetInstallResultArchive, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero MessagesStickerSetInstallResultArchive
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *MessagesStickerSetInstallResultArchiveArray) Pop() (v MessagesStickerSetInstallResultArchive, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
