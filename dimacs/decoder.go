// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

package dimacs

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strconv"
	"strings"
	"unicode"
	"unsafe"
)

// Decoder reads a DIMACS format stream of bytes from an io.Decoder.
type Decoder struct {
	s     *bufio.Scanner
	h     *Header
	htext []byte
	n     int
	c     []Lit
	err   error
}

// NewDecoder returns
func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{s: bufio.NewScanner(r)}
}

// Err returns any error that encountered while decoding the input bytes.
func (r *Decoder) Err() error {
	return r.err
}

// Header returns the header decoded from the input stream.  Header returns nil
// if no header could be decoded from the input.  If Header returns nil then
// r.Err() will return the encountered error.
func (r *Decoder) Header() *Header {
	if r.h == nil {
		r.readHeader()
		if r.err != nil {
			return nil
		}
	}
	h := &Header{}
	*h = *r.h
	return h
}

func (r *Decoder) skipComments() (ok bool) {
	for r.s.Scan() {
		b := r.s.Bytes()
		if len(b) == 0 {
			return
		}
		if b[0] == 'c' {
			continue
		}
		r.htext = b
		return true
	}
	r.err = r.s.Err()
	return false
}

func (r *Decoder) readHeader() {
	if r.h != nil || r.err != nil {
		return
	}

	if !r.skipComments() && r.err != nil {
		return
	}
	if len(r.htext) == 0 || r.htext[0] != 'p' {
		r.err = fmt.Errorf("missing problem header: %q", r.htext)
		return
	}
	fields := strings.Fields(string(r.htext))
	nextfield := func() (f string) {
		if len(fields) == 0 {
			return ""
		}
		f, fields = fields[0], fields[1:]
		return f
	}
	if nextfield() != "p" {
		r.err = fmt.Errorf("missing problem header")
		return
	}
	switch format := nextfield(); format {
	case "cnf":
	case "":
		r.err = fmt.Errorf("missing instance format in header")
		return
	default:
		r.err = fmt.Errorf("invalid instance format in header: %q", format)
		return
	}

	h := &Header{}
	switch _numvar := nextfield(); _numvar {
	case "":
		r.err = fmt.Errorf("missing instance nbvar")
		return
	default:
		h.NumVar, r.err = strconv.Atoi(_numvar)
		if r.err != nil {
			return
		}
	}

	switch _numclause := nextfield(); _numclause {
	case "":
		r.err = fmt.Errorf("missing instance nbclause")
		return
	default:
		h.NumClause, r.err = strconv.Atoi(_numclause)
		if r.err != nil {
			return
		}
	}

	if len(fields) > 0 {
		r.err = fmt.Errorf("too many fields in header")
		return
	}

	r.h = h
	r.c = make([]Lit, r.h.NumVar)
}

// Clause returns the last clause decoded from the input stream.  The
// underlying storage of the returned slice is part of an internal buffer.  Any
// attempt to presist the clause must store the literal values in a new slice.
func (r *Decoder) Clause() []Lit {
	if len(r.c) == 0 || r.err != nil {
		return nil
	}
	return r.c
}

// Decode decodes a clause from the input stream.  If r can decode a clause
// true is returned and the clause can be inspected or copied using r.Clause().
// If no clause can be decoded false is returned and r.Err() will return any
// encountered error.  If the stream was fully consumed than false will be
// returned and r.Err() will return nil.
func (r *Decoder) Decode() bool {
	r.readHeader()
	if r.err != nil {
		return false
	}

	if !r.skipComments() {
		r.err = r.s.Err()
		return false
	}

	if r.n >= r.h.NumClause {
		r.err = fmt.Errorf("too many clauses")
		return false
	}
	r.n++

	r.c = r.c[:0]

	lineFull := r.s.Bytes()
	line := lineFull
	foundTerm := false
	for {
		iNotSpace := bytes.IndexFunc(line, func(c rune) bool { return !unicode.IsSpace(c) })
		if iNotSpace < 0 {
			if !foundTerm && len(r.c) > 0 {
				r.err = fmt.Errorf("invalid clause line: missing terminating null")
				return false
			}
			return true
		}
		if foundTerm {
			r.err = fmt.Errorf("invalid clause line: unexpected null literal")
			return false
		}

		line = line[iNotSpace:]
		iSpace := bytes.IndexFunc(line, unicode.IsSpace)
		var field []byte
		if iSpace < 0 {
			field = line
			line = nil
		} else {
			field = line[:iSpace]
			line = line[iSpace:]
		}

		x, err := strconv.Atoi(*(*string)(unsafe.Pointer(&field)))
		if err != nil {
			r.err = fmt.Errorf("invalid clause line: failed to parse literal: %v", err)
			return false
		}
		if x == 0 {
			foundTerm = true
			continue
		}

		lit := Lit(x)
		v := lit.Var()

		if v > r.h.NumVar {
			r.err = fmt.Errorf("invalid clause line: variable outside of range")
			return false
		}

		r.c = append(r.c, lit)
	}
}
