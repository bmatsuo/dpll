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
	33: mkClauseLocal33,
	34: mkClauseLocal34,
	35: mkClauseLocal35,
	36: mkClauseLocal36,
	37: mkClauseLocal37,
	38: mkClauseLocal38,
	39: mkClauseLocal39,
	40: mkClauseLocal40,
	41: mkClauseLocal41,
	42: mkClauseLocal42,
	43: mkClauseLocal43,
	44: mkClauseLocal44,
	45: mkClauseLocal45,
	46: mkClauseLocal46,
	47: mkClauseLocal47,
	48: mkClauseLocal48,
	49: mkClauseLocal49,
	50: mkClauseLocal50,
	51: mkClauseLocal51,
	52: mkClauseLocal52,
	53: mkClauseLocal53,
	54: mkClauseLocal54,
	55: mkClauseLocal55,
	56: mkClauseLocal56,
	57: mkClauseLocal57,
	58: mkClauseLocal58,
	59: mkClauseLocal59,
	60: mkClauseLocal60,
	61: mkClauseLocal61,
	62: mkClauseLocal62,
	63: mkClauseLocal63,
	64: mkClauseLocal64,
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
	33: mkClauseExtraLocal33,
	34: mkClauseExtraLocal34,
	35: mkClauseExtraLocal35,
	36: mkClauseExtraLocal36,
	37: mkClauseExtraLocal37,
	38: mkClauseExtraLocal38,
	39: mkClauseExtraLocal39,
	40: mkClauseExtraLocal40,
	41: mkClauseExtraLocal41,
	42: mkClauseExtraLocal42,
	43: mkClauseExtraLocal43,
	44: mkClauseExtraLocal44,
	45: mkClauseExtraLocal45,
	46: mkClauseExtraLocal46,
	47: mkClauseExtraLocal47,
	48: mkClauseExtraLocal48,
	49: mkClauseExtraLocal49,
	50: mkClauseExtraLocal50,
	51: mkClauseExtraLocal51,
	52: mkClauseExtraLocal52,
	53: mkClauseExtraLocal53,
	54: mkClauseExtraLocal54,
	55: mkClauseExtraLocal55,
	56: mkClauseExtraLocal56,
	57: mkClauseExtraLocal57,
	58: mkClauseExtraLocal58,
	59: mkClauseExtraLocal59,
	60: mkClauseExtraLocal60,
	61: mkClauseExtraLocal61,
	62: mkClauseExtraLocal62,
	63: mkClauseExtraLocal63,
	64: mkClauseExtraLocal64,
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

type clauseLocal33 struct {
	Clause
	Lit [33]Lit
}
type clauseExtraLocal33 struct {
	Clause
	ClauseExtra
	Lit [33]Lit
}

type clauseLocal34 struct {
	Clause
	Lit [34]Lit
}
type clauseExtraLocal34 struct {
	Clause
	ClauseExtra
	Lit [34]Lit
}

type clauseLocal35 struct {
	Clause
	Lit [35]Lit
}
type clauseExtraLocal35 struct {
	Clause
	ClauseExtra
	Lit [35]Lit
}

type clauseLocal36 struct {
	Clause
	Lit [36]Lit
}
type clauseExtraLocal36 struct {
	Clause
	ClauseExtra
	Lit [36]Lit
}

type clauseLocal37 struct {
	Clause
	Lit [37]Lit
}
type clauseExtraLocal37 struct {
	Clause
	ClauseExtra
	Lit [37]Lit
}

type clauseLocal38 struct {
	Clause
	Lit [38]Lit
}
type clauseExtraLocal38 struct {
	Clause
	ClauseExtra
	Lit [38]Lit
}

type clauseLocal39 struct {
	Clause
	Lit [39]Lit
}
type clauseExtraLocal39 struct {
	Clause
	ClauseExtra
	Lit [39]Lit
}

type clauseLocal40 struct {
	Clause
	Lit [40]Lit
}
type clauseExtraLocal40 struct {
	Clause
	ClauseExtra
	Lit [40]Lit
}

type clauseLocal41 struct {
	Clause
	Lit [41]Lit
}
type clauseExtraLocal41 struct {
	Clause
	ClauseExtra
	Lit [41]Lit
}

type clauseLocal42 struct {
	Clause
	Lit [42]Lit
}
type clauseExtraLocal42 struct {
	Clause
	ClauseExtra
	Lit [42]Lit
}

type clauseLocal43 struct {
	Clause
	Lit [43]Lit
}
type clauseExtraLocal43 struct {
	Clause
	ClauseExtra
	Lit [43]Lit
}

type clauseLocal44 struct {
	Clause
	Lit [44]Lit
}
type clauseExtraLocal44 struct {
	Clause
	ClauseExtra
	Lit [44]Lit
}

type clauseLocal45 struct {
	Clause
	Lit [45]Lit
}
type clauseExtraLocal45 struct {
	Clause
	ClauseExtra
	Lit [45]Lit
}

type clauseLocal46 struct {
	Clause
	Lit [46]Lit
}
type clauseExtraLocal46 struct {
	Clause
	ClauseExtra
	Lit [46]Lit
}

type clauseLocal47 struct {
	Clause
	Lit [47]Lit
}
type clauseExtraLocal47 struct {
	Clause
	ClauseExtra
	Lit [47]Lit
}

type clauseLocal48 struct {
	Clause
	Lit [48]Lit
}
type clauseExtraLocal48 struct {
	Clause
	ClauseExtra
	Lit [48]Lit
}

type clauseLocal49 struct {
	Clause
	Lit [49]Lit
}
type clauseExtraLocal49 struct {
	Clause
	ClauseExtra
	Lit [49]Lit
}

type clauseLocal50 struct {
	Clause
	Lit [50]Lit
}
type clauseExtraLocal50 struct {
	Clause
	ClauseExtra
	Lit [50]Lit
}

type clauseLocal51 struct {
	Clause
	Lit [51]Lit
}
type clauseExtraLocal51 struct {
	Clause
	ClauseExtra
	Lit [51]Lit
}

type clauseLocal52 struct {
	Clause
	Lit [52]Lit
}
type clauseExtraLocal52 struct {
	Clause
	ClauseExtra
	Lit [52]Lit
}

type clauseLocal53 struct {
	Clause
	Lit [53]Lit
}
type clauseExtraLocal53 struct {
	Clause
	ClauseExtra
	Lit [53]Lit
}

type clauseLocal54 struct {
	Clause
	Lit [54]Lit
}
type clauseExtraLocal54 struct {
	Clause
	ClauseExtra
	Lit [54]Lit
}

type clauseLocal55 struct {
	Clause
	Lit [55]Lit
}
type clauseExtraLocal55 struct {
	Clause
	ClauseExtra
	Lit [55]Lit
}

type clauseLocal56 struct {
	Clause
	Lit [56]Lit
}
type clauseExtraLocal56 struct {
	Clause
	ClauseExtra
	Lit [56]Lit
}

type clauseLocal57 struct {
	Clause
	Lit [57]Lit
}
type clauseExtraLocal57 struct {
	Clause
	ClauseExtra
	Lit [57]Lit
}

type clauseLocal58 struct {
	Clause
	Lit [58]Lit
}
type clauseExtraLocal58 struct {
	Clause
	ClauseExtra
	Lit [58]Lit
}

type clauseLocal59 struct {
	Clause
	Lit [59]Lit
}
type clauseExtraLocal59 struct {
	Clause
	ClauseExtra
	Lit [59]Lit
}

type clauseLocal60 struct {
	Clause
	Lit [60]Lit
}
type clauseExtraLocal60 struct {
	Clause
	ClauseExtra
	Lit [60]Lit
}

type clauseLocal61 struct {
	Clause
	Lit [61]Lit
}
type clauseExtraLocal61 struct {
	Clause
	ClauseExtra
	Lit [61]Lit
}

type clauseLocal62 struct {
	Clause
	Lit [62]Lit
}
type clauseExtraLocal62 struct {
	Clause
	ClauseExtra
	Lit [62]Lit
}

type clauseLocal63 struct {
	Clause
	Lit [63]Lit
}
type clauseExtraLocal63 struct {
	Clause
	ClauseExtra
	Lit [63]Lit
}

type clauseLocal64 struct {
	Clause
	Lit [64]Lit
}
type clauseExtraLocal64 struct {
	Clause
	ClauseExtra
	Lit [64]Lit
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

func mkClauseLocal33(ps []Lit) *Clause {
	c := &clauseLocal33{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal33(ps []Lit) *Clause {
	c := &clauseExtraLocal33{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal34(ps []Lit) *Clause {
	c := &clauseLocal34{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal34(ps []Lit) *Clause {
	c := &clauseExtraLocal34{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal35(ps []Lit) *Clause {
	c := &clauseLocal35{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal35(ps []Lit) *Clause {
	c := &clauseExtraLocal35{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal36(ps []Lit) *Clause {
	c := &clauseLocal36{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal36(ps []Lit) *Clause {
	c := &clauseExtraLocal36{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal37(ps []Lit) *Clause {
	c := &clauseLocal37{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal37(ps []Lit) *Clause {
	c := &clauseExtraLocal37{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal38(ps []Lit) *Clause {
	c := &clauseLocal38{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal38(ps []Lit) *Clause {
	c := &clauseExtraLocal38{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal39(ps []Lit) *Clause {
	c := &clauseLocal39{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal39(ps []Lit) *Clause {
	c := &clauseExtraLocal39{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal40(ps []Lit) *Clause {
	c := &clauseLocal40{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal40(ps []Lit) *Clause {
	c := &clauseExtraLocal40{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal41(ps []Lit) *Clause {
	c := &clauseLocal41{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal41(ps []Lit) *Clause {
	c := &clauseExtraLocal41{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal42(ps []Lit) *Clause {
	c := &clauseLocal42{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal42(ps []Lit) *Clause {
	c := &clauseExtraLocal42{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal43(ps []Lit) *Clause {
	c := &clauseLocal43{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal43(ps []Lit) *Clause {
	c := &clauseExtraLocal43{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal44(ps []Lit) *Clause {
	c := &clauseLocal44{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal44(ps []Lit) *Clause {
	c := &clauseExtraLocal44{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal45(ps []Lit) *Clause {
	c := &clauseLocal45{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal45(ps []Lit) *Clause {
	c := &clauseExtraLocal45{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal46(ps []Lit) *Clause {
	c := &clauseLocal46{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal46(ps []Lit) *Clause {
	c := &clauseExtraLocal46{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal47(ps []Lit) *Clause {
	c := &clauseLocal47{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal47(ps []Lit) *Clause {
	c := &clauseExtraLocal47{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal48(ps []Lit) *Clause {
	c := &clauseLocal48{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal48(ps []Lit) *Clause {
	c := &clauseExtraLocal48{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal49(ps []Lit) *Clause {
	c := &clauseLocal49{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal49(ps []Lit) *Clause {
	c := &clauseExtraLocal49{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal50(ps []Lit) *Clause {
	c := &clauseLocal50{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal50(ps []Lit) *Clause {
	c := &clauseExtraLocal50{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal51(ps []Lit) *Clause {
	c := &clauseLocal51{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal51(ps []Lit) *Clause {
	c := &clauseExtraLocal51{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal52(ps []Lit) *Clause {
	c := &clauseLocal52{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal52(ps []Lit) *Clause {
	c := &clauseExtraLocal52{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal53(ps []Lit) *Clause {
	c := &clauseLocal53{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal53(ps []Lit) *Clause {
	c := &clauseExtraLocal53{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal54(ps []Lit) *Clause {
	c := &clauseLocal54{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal54(ps []Lit) *Clause {
	c := &clauseExtraLocal54{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal55(ps []Lit) *Clause {
	c := &clauseLocal55{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal55(ps []Lit) *Clause {
	c := &clauseExtraLocal55{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal56(ps []Lit) *Clause {
	c := &clauseLocal56{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal56(ps []Lit) *Clause {
	c := &clauseExtraLocal56{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal57(ps []Lit) *Clause {
	c := &clauseLocal57{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal57(ps []Lit) *Clause {
	c := &clauseExtraLocal57{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal58(ps []Lit) *Clause {
	c := &clauseLocal58{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal58(ps []Lit) *Clause {
	c := &clauseExtraLocal58{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal59(ps []Lit) *Clause {
	c := &clauseLocal59{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal59(ps []Lit) *Clause {
	c := &clauseExtraLocal59{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal60(ps []Lit) *Clause {
	c := &clauseLocal60{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal60(ps []Lit) *Clause {
	c := &clauseExtraLocal60{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal61(ps []Lit) *Clause {
	c := &clauseLocal61{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal61(ps []Lit) *Clause {
	c := &clauseExtraLocal61{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal62(ps []Lit) *Clause {
	c := &clauseLocal62{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal62(ps []Lit) *Clause {
	c := &clauseExtraLocal62{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal63(ps []Lit) *Clause {
	c := &clauseLocal63{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal63(ps []Lit) *Clause {
	c := &clauseExtraLocal63{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}

func mkClauseLocal64(ps []Lit) *Clause {
	c := &clauseLocal64{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	return &c.Clause
}
func mkClauseExtraLocal64(ps []Lit) *Clause {
	c := &clauseExtraLocal64{}
	copy(c.Lit[:], ps)
	c.Clause.Lit = c.Lit[:]
	c.Clause.ClauseExtra = &c.ClauseExtra
	return &c.Clause
}
