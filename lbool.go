// Copyright 2016 Bryan Matsuo
//
// Use of this software is governed by the MIT license.  A copy of the license
// agreement can be found in the LICENSE file distributed with this software.

package dpll

// LBool constants
const (
	LTrue  = LBool(0)
	LFalse = LBool(1)
	LUndef = LBool(2)
)

// LBool is a lifted boolean value which may be undefined.
type LBool uint8

// Bool returns LTrue or LFalse if b is true or false respectively.
func Bool(b bool) LBool {
	if b {
		return LTrue
	}
	return LFalse
}

var lboolStrings = []string{
	LTrue:  "LTrue",
	LFalse: "LFalse",
	LUndef: "LUndef",
}

// String returns the string representation of an LBool
func (b LBool) String() string {
	if int(b) >= len(lboolStrings) {
		return "LInvalid"
	}
	return lboolStrings[b]
}

// IsTrue returns true if b is equivalent to LTrue
func (b LBool) IsTrue() bool {
	return b.Equal(LTrue)
}

// IsFalse returns true if b is equivalent to LFalse
func (b LBool) IsFalse() bool {
	return b.Equal(LFalse)
}

// IsUndef returns true if b is equivalent to LUndef
func (b LBool) IsUndef() bool {
	return b.Equal(LUndef)
}

// Equal returns true iff b and b2 are equivalent.
func (b LBool) Equal(b2 LBool) bool {
	// NOTE: the b&2 == 0 check is a little pedantic. it seems to presurpose
	// that values other than the three enumerated constants can be used.
	return (b&2 != 0 && b2&2 != 0) || (b&2 == 0 && b == b2)
}

// And returns the conjunction of b and b2.
func (b LBool) And(b2 LBool) LBool {
	const magic = 0xF7F755F4
	return LBool(uint(magic)>>((b<<1)|(b2<<3))) & 3
}

// Or returns the disjunction of b and b2.
func (b LBool) Or(b2 LBool) LBool {
	const magic = 0xFCFCF400
	return LBool(uint(magic)>>((b<<1)|(b2<<3))) & 3
}

// Xor returns the exclusive or of b and b2.
func (b LBool) Xor(b2 bool) LBool {
	if b2 {
		return b&1 ^ 1
	}
	return b
}
