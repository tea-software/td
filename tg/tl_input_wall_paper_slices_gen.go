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

// InputWallPaperClassArray is adapter for slice of InputWallPaperClass.
type InputWallPaperClassArray []InputWallPaperClass

// Sort sorts slice of InputWallPaperClass.
func (s InputWallPaperClassArray) Sort(less func(a, b InputWallPaperClass) bool) InputWallPaperClassArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputWallPaperClass.
func (s InputWallPaperClassArray) SortStable(less func(a, b InputWallPaperClass) bool) InputWallPaperClassArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputWallPaperClass.
func (s InputWallPaperClassArray) Retain(keep func(x InputWallPaperClass) bool) InputWallPaperClassArray {
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
func (s InputWallPaperClassArray) First() (v InputWallPaperClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputWallPaperClassArray) Last() (v InputWallPaperClass, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputWallPaperClassArray) PopFirst() (v InputWallPaperClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputWallPaperClass
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputWallPaperClassArray) Pop() (v InputWallPaperClass, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// AsInputWallPaper returns copy with only InputWallPaper constructors.
func (s InputWallPaperClassArray) AsInputWallPaper() (to InputWallPaperArray) {
	for _, elem := range s {
		value, ok := elem.(*InputWallPaper)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputWallPaperSlug returns copy with only InputWallPaperSlug constructors.
func (s InputWallPaperClassArray) AsInputWallPaperSlug() (to InputWallPaperSlugArray) {
	for _, elem := range s {
		value, ok := elem.(*InputWallPaperSlug)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// AsInputWallPaperNoFile returns copy with only InputWallPaperNoFile constructors.
func (s InputWallPaperClassArray) AsInputWallPaperNoFile() (to InputWallPaperNoFileArray) {
	for _, elem := range s {
		value, ok := elem.(*InputWallPaperNoFile)
		if !ok {
			continue
		}
		to = append(to, *value)
	}

	return to
}

// InputWallPaperArray is adapter for slice of InputWallPaper.
type InputWallPaperArray []InputWallPaper

// Sort sorts slice of InputWallPaper.
func (s InputWallPaperArray) Sort(less func(a, b InputWallPaper) bool) InputWallPaperArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputWallPaper.
func (s InputWallPaperArray) SortStable(less func(a, b InputWallPaper) bool) InputWallPaperArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputWallPaper.
func (s InputWallPaperArray) Retain(keep func(x InputWallPaper) bool) InputWallPaperArray {
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
func (s InputWallPaperArray) First() (v InputWallPaper, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputWallPaperArray) Last() (v InputWallPaper, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputWallPaperArray) PopFirst() (v InputWallPaper, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputWallPaper
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputWallPaperArray) Pop() (v InputWallPaper, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputWallPaperSlugArray is adapter for slice of InputWallPaperSlug.
type InputWallPaperSlugArray []InputWallPaperSlug

// Sort sorts slice of InputWallPaperSlug.
func (s InputWallPaperSlugArray) Sort(less func(a, b InputWallPaperSlug) bool) InputWallPaperSlugArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputWallPaperSlug.
func (s InputWallPaperSlugArray) SortStable(less func(a, b InputWallPaperSlug) bool) InputWallPaperSlugArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputWallPaperSlug.
func (s InputWallPaperSlugArray) Retain(keep func(x InputWallPaperSlug) bool) InputWallPaperSlugArray {
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
func (s InputWallPaperSlugArray) First() (v InputWallPaperSlug, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputWallPaperSlugArray) Last() (v InputWallPaperSlug, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputWallPaperSlugArray) PopFirst() (v InputWallPaperSlug, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputWallPaperSlug
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputWallPaperSlugArray) Pop() (v InputWallPaperSlug, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// InputWallPaperNoFileArray is adapter for slice of InputWallPaperNoFile.
type InputWallPaperNoFileArray []InputWallPaperNoFile

// Sort sorts slice of InputWallPaperNoFile.
func (s InputWallPaperNoFileArray) Sort(less func(a, b InputWallPaperNoFile) bool) InputWallPaperNoFileArray {
	sort.Slice(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// SortStable sorts slice of InputWallPaperNoFile.
func (s InputWallPaperNoFileArray) SortStable(less func(a, b InputWallPaperNoFile) bool) InputWallPaperNoFileArray {
	sort.SliceStable(s, func(i, j int) bool {
		return less(s[i], s[j])
	})
	return s
}

// Retain filters in-place slice of InputWallPaperNoFile.
func (s InputWallPaperNoFileArray) Retain(keep func(x InputWallPaperNoFile) bool) InputWallPaperNoFileArray {
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
func (s InputWallPaperNoFileArray) First() (v InputWallPaperNoFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[0], true
}

// Last returns last element of slice (if exists).
func (s InputWallPaperNoFileArray) Last() (v InputWallPaperNoFile, ok bool) {
	if len(s) < 1 {
		return
	}
	return s[len(s)-1], true
}

// PopFirst returns first element of slice (if exists) and deletes it.
func (s *InputWallPaperNoFileArray) PopFirst() (v InputWallPaperNoFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[0]

	// Delete by index from SliceTricks.
	copy(a[0:], a[1:])
	var zero InputWallPaperNoFile
	a[len(a)-1] = zero
	a = a[:len(a)-1]
	*s = a

	return v, true
}

// Pop returns last element of slice (if exists) and deletes it.
func (s *InputWallPaperNoFileArray) Pop() (v InputWallPaperNoFile, ok bool) {
	if s == nil || len(*s) < 1 {
		return
	}

	a := *s
	v = a[len(a)-1]
	a = a[:len(a)-1]
	*s = a

	return v, true
}
