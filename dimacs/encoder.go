package dimacs

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

// Encoder encodes Lit clauses as a DIMACS format byte stream.
type Encoder struct {
	w    *bufio.Writer
	h    *Header
	seen []bool
	n    int
}

// NewEncoder initializes a new Encoder.  The returned encoder must be closed
// before it is discarded to avoid a corrupt output stream.
//		enc := NewEncoder(f)
//		defer enc.Close()
func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{w: bufio.NewWriter(w)}
}

// WriteHeader encodes and writes h to the output stream.  WriteHeader must be
// called only once, before any clauses have been written.
func (enc *Encoder) WriteHeader(h *Header) error {
	enc.h = h
	enc.seen = make([]bool, h.NumVar+1)
	_, err := fmt.Fprintf(enc.w, "p cnf %d %d\n", h.NumVar, h.NumClause)
	return err
}

// Encode encodes clause and writes it to the output stream.
func (enc *Encoder) Encode(clause []Lit) error {
	if enc.h == nil {
		return fmt.Errorf("no header")
	}
	if len(clause) > enc.h.NumVar {
		return fmt.Errorf("too many literals")
	}
	if enc.n >= enc.h.NumClause {
		return fmt.Errorf("too many clauses supplied")
	}
	for i := range enc.seen {
		enc.seen[i] = false
	}
	for _, lit := range clause {
		v := lit.Var()
		if v == 0 || v > enc.h.NumVar {
			return fmt.Errorf("invalid literal: %d", lit)
		}
		if enc.seen[v] {
			return fmt.Errorf("duplicate variable: %d", v)
		}
		enc.seen[v] = true
	}
	for _, lit := range clause {
		s := strconv.Itoa(int(lit))
		err := enc.writeString(s)
		if err != nil {
			return err
		}
		err = enc.writeString(" ")
		if err != nil {
			return err
		}
	}
	err := enc.writeString("0\n")
	if err != nil {
		return err
	}
	enc.n++
	return nil
}

func (enc *Encoder) writeString(s string) error {
	for len(s) > 0 {
		n, err := enc.w.WriteString(s)
		if err != nil {
			return err
		}
		s = s[n:]
	}
	return nil
}

// Close writes any buffered output to the underlying io.Writer.  If the number
// of encoded clauses is less than the number specified in the header an error
// is returned.
func (enc *Encoder) Close() error {
	if enc.h == nil {
		return fmt.Errorf("no output written")
	}
	if enc.n != enc.h.NumClause {
		return fmt.Errorf("not enough clauses encoded")
	}
	return enc.w.Flush()
}
