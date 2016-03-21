package dpll

var mkClauseLocal = []func([]Lit) *Clause{
	1:  mkClauseLocal1,
	2:  mkClauseLocal2,
	3:  mkClauseLocal3,
	4:  mkClauseLocal4,
	5:  mkClauseLocal5,
	6:  mkClauseLocal6,
	7:  mkClauseLocal7,
	8:  mkClauseLocal8,
	9:  mkClauseLocal9,
	10: mkClauseLocal10,
	11: mkClauseLocal11,
	12: mkClauseLocal12,
	13: mkClauseLocal13,
	14: mkClauseLocal14,
	15: mkClauseLocal15,
	16: mkClauseLocal16,
	17: mkClauseLocal17,
	18: mkClauseLocal18,
	19: mkClauseLocal19,
	20: mkClauseLocal20,
	21: mkClauseLocal21,
	22: mkClauseLocal22,
	23: mkClauseLocal23,
	24: mkClauseLocal24,
	25: mkClauseLocal25,
	26: mkClauseLocal26,
	27: mkClauseLocal27,
	28: mkClauseLocal28,
	29: mkClauseLocal29,
	30: mkClauseLocal30,
	31: mkClauseLocal31,
	32: mkClauseLocal32,
}
var mkClauseExtraLocal = []func([]Lit) *Clause{
	1:  mkClauseExtraLocal1,
	2:  mkClauseExtraLocal2,
	3:  mkClauseExtraLocal3,
	4:  mkClauseExtraLocal4,
	5:  mkClauseExtraLocal5,
	6:  mkClauseExtraLocal6,
	7:  mkClauseExtraLocal7,
	8:  mkClauseExtraLocal8,
	9:  mkClauseExtraLocal9,
	10: mkClauseExtraLocal10,
	11: mkClauseExtraLocal11,
	12: mkClauseExtraLocal12,
	13: mkClauseExtraLocal13,
	14: mkClauseExtraLocal14,
	15: mkClauseExtraLocal15,
	16: mkClauseExtraLocal16,
	17: mkClauseExtraLocal17,
	18: mkClauseExtraLocal18,
	19: mkClauseExtraLocal19,
	20: mkClauseExtraLocal20,
	21: mkClauseExtraLocal21,
	22: mkClauseExtraLocal22,
	23: mkClauseExtraLocal23,
	24: mkClauseExtraLocal24,
	25: mkClauseExtraLocal25,
	26: mkClauseExtraLocal26,
	27: mkClauseExtraLocal27,
	28: mkClauseExtraLocal28,
	29: mkClauseExtraLocal29,
	30: mkClauseExtraLocal30,
	31: mkClauseExtraLocal31,
	32: mkClauseExtraLocal32,
}

type clauseLocal1 struct {
	Clause
	Lit [1]Lit
}
type clauseExtraLocal1 struct {
	Clause
	ClauseExtra
	Lit [1]Lit
}

type clauseLocal2 struct {
	Clause
	Lit [2]Lit
}
type clauseExtraLocal2 struct {
	Clause
	ClauseExtra
	Lit [2]Lit
}

type clauseLocal3 struct {
	Clause
	Lit [3]Lit
}
type clauseExtraLocal3 struct {
	Clause
	ClauseExtra
	Lit [3]Lit
}

type clauseLocal4 struct {
	Clause
	Lit [4]Lit
}
type clauseExtraLocal4 struct {
	Clause
	ClauseExtra
	Lit [4]Lit
}

type clauseLocal5 struct {
	Clause
	Lit [5]Lit
}
type clauseExtraLocal5 struct {
	Clause
	ClauseExtra
	Lit [5]Lit
}

type clauseLocal6 struct {
	Clause
	Lit [6]Lit
}
type clauseExtraLocal6 struct {
	Clause
	ClauseExtra
	Lit [6]Lit
}

type clauseLocal7 struct {
	Clause
	Lit [7]Lit
}
type clauseExtraLocal7 struct {
	Clause
	ClauseExtra
	Lit [7]Lit
}

type clauseLocal8 struct {
	Clause
	Lit [8]Lit
}
type clauseExtraLocal8 struct {
	Clause
	ClauseExtra
	Lit [8]Lit
}

type clauseLocal9 struct {
	Clause
	Lit [9]Lit
}
type clauseExtraLocal9 struct {
	Clause
	ClauseExtra
	Lit [9]Lit
}

type clauseLocal10 struct {
	Clause
	Lit [10]Lit
}
type clauseExtraLocal10 struct {
	Clause
	ClauseExtra
	Lit [10]Lit
}

type clauseLocal11 struct {
	Clause
	Lit [11]Lit
}
type clauseExtraLocal11 struct {
	Clause
	ClauseExtra
	Lit [11]Lit
}

type clauseLocal12 struct {
	Clause
	Lit [12]Lit
}
type clauseExtraLocal12 struct {
	Clause
	ClauseExtra
	Lit [12]Lit
}

type clauseLocal13 struct {
	Clause
	Lit [13]Lit
}
type clauseExtraLocal13 struct {
	Clause
	ClauseExtra
	Lit [13]Lit
}

type clauseLocal14 struct {
	Clause
	Lit [14]Lit
}
type clauseExtraLocal14 struct {
	Clause
	ClauseExtra
	Lit [14]Lit
}

type clauseLocal15 struct {
	Clause
	Lit [15]Lit
}
type clauseExtraLocal15 struct {
	Clause
	ClauseExtra
	Lit [15]Lit
}

type clauseLocal16 struct {
	Clause
	Lit [16]Lit
}
type clauseExtraLocal16 struct {
	Clause
	ClauseExtra
	Lit [16]Lit
}

type clauseLocal17 struct {
	Clause
	Lit [17]Lit
}
type clauseExtraLocal17 struct {
	Clause
	ClauseExtra
	Lit [17]Lit
}

type clauseLocal18 struct {
	Clause
	Lit [18]Lit
}
type clauseExtraLocal18 struct {
	Clause
	ClauseExtra
	Lit [18]Lit
}

type clauseLocal19 struct {
	Clause
	Lit [19]Lit
}
type clauseExtraLocal19 struct {
	Clause
	ClauseExtra
	Lit [19]Lit
}

type clauseLocal20 struct {
	Clause
	Lit [20]Lit
}
type clauseExtraLocal20 struct {
	Clause
	ClauseExtra
	Lit [20]Lit
}

type clauseLocal21 struct {
	Clause
	Lit [21]Lit
}
type clauseExtraLocal21 struct {
	Clause
	ClauseExtra
	Lit [21]Lit
}

type clauseLocal22 struct {
	Clause
	Lit [22]Lit
}
type clauseExtraLocal22 struct {
	Clause
	ClauseExtra
	Lit [22]Lit
}

type clauseLocal23 struct {
	Clause
	Lit [23]Lit
}
type clauseExtraLocal23 struct {
	Clause
	ClauseExtra
	Lit [23]Lit
}

type clauseLocal24 struct {
	Clause
	Lit [24]Lit
}
type clauseExtraLocal24 struct {
	Clause
	ClauseExtra
	Lit [24]Lit
}

type clauseLocal25 struct {
	Clause
	Lit [25]Lit
}
type clauseExtraLocal25 struct {
	Clause
	ClauseExtra
	Lit [25]Lit
}

type clauseLocal26 struct {
	Clause
	Lit [26]Lit
}
type clauseExtraLocal26 struct {
	Clause
	ClauseExtra
	Lit [26]Lit
}

type clauseLocal27 struct {
	Clause
	Lit [27]Lit
}
type clauseExtraLocal27 struct {
	Clause
	ClauseExtra
	Lit [27]Lit
}

type clauseLocal28 struct {
	Clause
	Lit [28]Lit
}
type clauseExtraLocal28 struct {
	Clause
	ClauseExtra
	Lit [28]Lit
}

type clauseLocal29 struct {
	Clause
	Lit [29]Lit
}
type clauseExtraLocal29 struct {
	Clause
	ClauseExtra
	Lit [29]Lit
}

type clauseLocal30 struct {
	Clause
	Lit [30]Lit
}
type clauseExtraLocal30 struct {
	Clause
	ClauseExtra
	Lit [30]Lit
}

type clauseLocal31 struct {
	Clause
	Lit [31]Lit
}
type clauseExtraLocal31 struct {
	Clause
	ClauseExtra
	Lit [31]Lit
}

type clauseLocal32 struct {
	Clause
	Lit [32]Lit
}
type clauseExtraLocal32 struct {
	Clause
	ClauseExtra
	Lit [32]Lit
}

func mkClauseLocal1(ps []Lit) *Clause {
	c := &clauseLocal1{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal1(ps []Lit) *Clause {
	c := &clauseExtraLocal1{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal2(ps []Lit) *Clause {
	c := &clauseLocal2{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal2(ps []Lit) *Clause {
	c := &clauseExtraLocal2{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal3(ps []Lit) *Clause {
	c := &clauseLocal3{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal3(ps []Lit) *Clause {
	c := &clauseExtraLocal3{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal4(ps []Lit) *Clause {
	c := &clauseLocal4{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal4(ps []Lit) *Clause {
	c := &clauseExtraLocal4{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal5(ps []Lit) *Clause {
	c := &clauseLocal5{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal5(ps []Lit) *Clause {
	c := &clauseExtraLocal5{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal6(ps []Lit) *Clause {
	c := &clauseLocal6{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal6(ps []Lit) *Clause {
	c := &clauseExtraLocal6{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal7(ps []Lit) *Clause {
	c := &clauseLocal7{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal7(ps []Lit) *Clause {
	c := &clauseExtraLocal7{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal8(ps []Lit) *Clause {
	c := &clauseLocal8{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal8(ps []Lit) *Clause {
	c := &clauseExtraLocal8{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal9(ps []Lit) *Clause {
	c := &clauseLocal9{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal9(ps []Lit) *Clause {
	c := &clauseExtraLocal9{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal10(ps []Lit) *Clause {
	c := &clauseLocal10{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal10(ps []Lit) *Clause {
	c := &clauseExtraLocal10{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal11(ps []Lit) *Clause {
	c := &clauseLocal11{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal11(ps []Lit) *Clause {
	c := &clauseExtraLocal11{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal12(ps []Lit) *Clause {
	c := &clauseLocal12{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal12(ps []Lit) *Clause {
	c := &clauseExtraLocal12{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal13(ps []Lit) *Clause {
	c := &clauseLocal13{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal13(ps []Lit) *Clause {
	c := &clauseExtraLocal13{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal14(ps []Lit) *Clause {
	c := &clauseLocal14{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal14(ps []Lit) *Clause {
	c := &clauseExtraLocal14{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal15(ps []Lit) *Clause {
	c := &clauseLocal15{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal15(ps []Lit) *Clause {
	c := &clauseExtraLocal15{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal16(ps []Lit) *Clause {
	c := &clauseLocal16{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal16(ps []Lit) *Clause {
	c := &clauseExtraLocal16{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal17(ps []Lit) *Clause {
	c := &clauseLocal17{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal17(ps []Lit) *Clause {
	c := &clauseExtraLocal17{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal18(ps []Lit) *Clause {
	c := &clauseLocal18{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal18(ps []Lit) *Clause {
	c := &clauseExtraLocal18{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal19(ps []Lit) *Clause {
	c := &clauseLocal19{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal19(ps []Lit) *Clause {
	c := &clauseExtraLocal19{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal20(ps []Lit) *Clause {
	c := &clauseLocal20{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal20(ps []Lit) *Clause {
	c := &clauseExtraLocal20{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal21(ps []Lit) *Clause {
	c := &clauseLocal21{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal21(ps []Lit) *Clause {
	c := &clauseExtraLocal21{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal22(ps []Lit) *Clause {
	c := &clauseLocal22{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal22(ps []Lit) *Clause {
	c := &clauseExtraLocal22{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal23(ps []Lit) *Clause {
	c := &clauseLocal23{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal23(ps []Lit) *Clause {
	c := &clauseExtraLocal23{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal24(ps []Lit) *Clause {
	c := &clauseLocal24{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal24(ps []Lit) *Clause {
	c := &clauseExtraLocal24{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal25(ps []Lit) *Clause {
	c := &clauseLocal25{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal25(ps []Lit) *Clause {
	c := &clauseExtraLocal25{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal26(ps []Lit) *Clause {
	c := &clauseLocal26{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal26(ps []Lit) *Clause {
	c := &clauseExtraLocal26{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal27(ps []Lit) *Clause {
	c := &clauseLocal27{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal27(ps []Lit) *Clause {
	c := &clauseExtraLocal27{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal28(ps []Lit) *Clause {
	c := &clauseLocal28{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal28(ps []Lit) *Clause {
	c := &clauseExtraLocal28{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal29(ps []Lit) *Clause {
	c := &clauseLocal29{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal29(ps []Lit) *Clause {
	c := &clauseExtraLocal29{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal30(ps []Lit) *Clause {
	c := &clauseLocal30{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal30(ps []Lit) *Clause {
	c := &clauseExtraLocal30{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal31(ps []Lit) *Clause {
	c := &clauseLocal31{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal31(ps []Lit) *Clause {
	c := &clauseExtraLocal31{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal32(ps []Lit) *Clause {
	c := &clauseLocal32{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal32(ps []Lit) *Clause {
	c := &clauseExtraLocal32{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}
