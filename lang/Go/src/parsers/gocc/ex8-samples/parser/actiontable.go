// Code generated by gocc; DO NOT EDIT.

package parser

type (
	actionTable [numStates]actionRow
	actionRow   struct {
		canRecover bool
		actions    [numSymbols]action
	}
)

var actionTab = actionTable{
	actionRow{ // S0
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // semicolon
			shift(4),  // if
			nil,       // else
			nil,       // plus_minus
			shift(6),  // digits
			shift(7),  // binary_seq
			shift(8),  // ident
			shift(9),  // abc
			shift(10), // def
		},
	},
	actionRow{ // S1
		canRecover: false,
		actions: [numSymbols]action{
			nil,          // INVALID
			accept(true), // $
			nil,          // semicolon
			nil,          // if
			nil,          // else
			nil,          // plus_minus
			nil,          // digits
			nil,          // binary_seq
			nil,          // ident
			nil,          // abc
			nil,          // def
		},
	},
	actionRow{ // S2
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(11), // semicolon
			nil,       // if
			nil,       // else
			shift(12), // plus_minus
			nil,       // digits
			nil,       // binary_seq
			nil,       // ident
			nil,       // abc
			nil,       // def
		},
	},
	actionRow{ // S3
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(2), // $, reduce: Stmt
			nil,       // semicolon
			nil,       // if
			nil,       // else
			nil,       // plus_minus
			nil,       // digits
			nil,       // binary_seq
			nil,       // ident
			nil,       // abc
			nil,       // def
		},
	},
	actionRow{ // S4
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // semicolon
			nil,       // if
			nil,       // else
			nil,       // plus_minus
			shift(15), // digits
			shift(16), // binary_seq
			shift(17), // ident
			shift(18), // abc
			shift(19), // def
		},
	},
	actionRow{ // S5
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(5), // semicolon, reduce: Expr
			nil,       // if
			nil,       // else
			reduce(5), // plus_minus, reduce: Expr
			nil,       // digits
			nil,       // binary_seq
			nil,       // ident
			nil,       // abc
			nil,       // def
		},
	},
	actionRow{ // S6
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(6), // semicolon, reduce: Term
			nil,       // if
			nil,       // else
			reduce(6), // plus_minus, reduce: Term
			nil,       // digits
			nil,       // binary_seq
			nil,       // ident
			nil,       // abc
			nil,       // def
		},
	},
	actionRow{ // S7
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(7), // semicolon, reduce: Term
			nil,       // if
			nil,       // else
			reduce(7), // plus_minus, reduce: Term
			nil,       // digits
			nil,       // binary_seq
			nil,       // ident
			nil,       // abc
			nil,       // def
		},
	},
	actionRow{ // S8
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(8), // semicolon, reduce: Term
			nil,       // if
			nil,       // else
			reduce(8), // plus_minus, reduce: Term
			nil,       // digits
			nil,       // binary_seq
			nil,       // ident
			nil,       // abc
			nil,       // def
		},
	},
	actionRow{ // S9
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(9), // semicolon, reduce: Term
			nil,       // if
			nil,       // else
			reduce(9), // plus_minus, reduce: Term
			nil,       // digits
			nil,       // binary_seq
			nil,       // ident
			nil,       // abc
			nil,       // def
		},
	},
	actionRow{ // S10
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			reduce(10), // semicolon, reduce: Term
			nil,        // if
			nil,        // else
			reduce(10), // plus_minus, reduce: Term
			nil,        // digits
			nil,        // binary_seq
			nil,        // ident
			nil,        // abc
			nil,        // def
		},
	},
	actionRow{ // S11
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(1), // $, reduce: Stmt
			nil,       // semicolon
			nil,       // if
			nil,       // else
			nil,       // plus_minus
			nil,       // digits
			nil,       // binary_seq
			nil,       // ident
			nil,       // abc
			nil,       // def
		},
	},
	actionRow{ // S12
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // semicolon
			nil,       // if
			nil,       // else
			nil,       // plus_minus
			shift(6),  // digits
			shift(7),  // binary_seq
			shift(8),  // ident
			shift(9),  // abc
			shift(10), // def
		},
	},
	actionRow{ // S13
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // semicolon
			shift(24), // if
			nil,       // else
			shift(25), // plus_minus
			shift(6),  // digits
			shift(7),  // binary_seq
			shift(8),  // ident
			shift(9),  // abc
			shift(10), // def
		},
	},
	actionRow{ // S14
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // semicolon
			reduce(5), // if, reduce: Expr
			nil,       // else
			reduce(5), // plus_minus, reduce: Expr
			reduce(5), // digits, reduce: Expr
			reduce(5), // binary_seq, reduce: Expr
			reduce(5), // ident, reduce: Expr
			reduce(5), // abc, reduce: Expr
			reduce(5), // def, reduce: Expr
		},
	},
	actionRow{ // S15
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // semicolon
			reduce(6), // if, reduce: Term
			nil,       // else
			reduce(6), // plus_minus, reduce: Term
			reduce(6), // digits, reduce: Term
			reduce(6), // binary_seq, reduce: Term
			reduce(6), // ident, reduce: Term
			reduce(6), // abc, reduce: Term
			reduce(6), // def, reduce: Term
		},
	},
	actionRow{ // S16
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // semicolon
			reduce(7), // if, reduce: Term
			nil,       // else
			reduce(7), // plus_minus, reduce: Term
			reduce(7), // digits, reduce: Term
			reduce(7), // binary_seq, reduce: Term
			reduce(7), // ident, reduce: Term
			reduce(7), // abc, reduce: Term
			reduce(7), // def, reduce: Term
		},
	},
	actionRow{ // S17
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // semicolon
			reduce(8), // if, reduce: Term
			nil,       // else
			reduce(8), // plus_minus, reduce: Term
			reduce(8), // digits, reduce: Term
			reduce(8), // binary_seq, reduce: Term
			reduce(8), // ident, reduce: Term
			reduce(8), // abc, reduce: Term
			reduce(8), // def, reduce: Term
		},
	},
	actionRow{ // S18
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // semicolon
			reduce(9), // if, reduce: Term
			nil,       // else
			reduce(9), // plus_minus, reduce: Term
			reduce(9), // digits, reduce: Term
			reduce(9), // binary_seq, reduce: Term
			reduce(9), // ident, reduce: Term
			reduce(9), // abc, reduce: Term
			reduce(9), // def, reduce: Term
		},
	},
	actionRow{ // S19
		canRecover: false,
		actions: [numSymbols]action{
			nil,        // INVALID
			nil,        // $
			nil,        // semicolon
			reduce(10), // if, reduce: Term
			nil,        // else
			reduce(10), // plus_minus, reduce: Term
			reduce(10), // digits, reduce: Term
			reduce(10), // binary_seq, reduce: Term
			reduce(10), // ident, reduce: Term
			reduce(10), // abc, reduce: Term
			reduce(10), // def, reduce: Term
		},
	},
	actionRow{ // S20
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(4), // semicolon, reduce: Expr
			nil,       // if
			nil,       // else
			reduce(4), // plus_minus, reduce: Expr
			nil,       // digits
			nil,       // binary_seq
			nil,       // ident
			nil,       // abc
			nil,       // def
		},
	},
	actionRow{ // S21
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // semicolon
			nil,       // if
			shift(26), // else
			nil,       // plus_minus
			nil,       // digits
			nil,       // binary_seq
			nil,       // ident
			nil,       // abc
			nil,       // def
		},
	},
	actionRow{ // S22
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(27), // semicolon
			nil,       // if
			nil,       // else
			shift(12), // plus_minus
			nil,       // digits
			nil,       // binary_seq
			nil,       // ident
			nil,       // abc
			nil,       // def
		},
	},
	actionRow{ // S23
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // semicolon
			nil,       // if
			reduce(2), // else, reduce: Stmt
			nil,       // plus_minus
			nil,       // digits
			nil,       // binary_seq
			nil,       // ident
			nil,       // abc
			nil,       // def
		},
	},
	actionRow{ // S24
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // semicolon
			nil,       // if
			nil,       // else
			nil,       // plus_minus
			shift(15), // digits
			shift(16), // binary_seq
			shift(17), // ident
			shift(18), // abc
			shift(19), // def
		},
	},
	actionRow{ // S25
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // semicolon
			nil,       // if
			nil,       // else
			nil,       // plus_minus
			shift(15), // digits
			shift(16), // binary_seq
			shift(17), // ident
			shift(18), // abc
			shift(19), // def
		},
	},
	actionRow{ // S26
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // semicolon
			shift(33), // if
			nil,       // else
			nil,       // plus_minus
			shift(6),  // digits
			shift(7),  // binary_seq
			shift(8),  // ident
			shift(9),  // abc
			shift(10), // def
		},
	},
	actionRow{ // S27
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // semicolon
			nil,       // if
			reduce(1), // else, reduce: Stmt
			nil,       // plus_minus
			nil,       // digits
			nil,       // binary_seq
			nil,       // ident
			nil,       // abc
			nil,       // def
		},
	},
	actionRow{ // S28
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // semicolon
			shift(24), // if
			nil,       // else
			shift(25), // plus_minus
			shift(6),  // digits
			shift(7),  // binary_seq
			shift(8),  // ident
			shift(9),  // abc
			shift(10), // def
		},
	},
	actionRow{ // S29
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // semicolon
			reduce(4), // if, reduce: Expr
			nil,       // else
			reduce(4), // plus_minus, reduce: Expr
			reduce(4), // digits, reduce: Expr
			reduce(4), // binary_seq, reduce: Expr
			reduce(4), // ident, reduce: Expr
			reduce(4), // abc, reduce: Expr
			reduce(4), // def, reduce: Expr
		},
	},
	actionRow{ // S30
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(35), // semicolon
			nil,       // if
			nil,       // else
			nil,       // plus_minus
			nil,       // digits
			nil,       // binary_seq
			nil,       // ident
			nil,       // abc
			nil,       // def
		},
	},
	actionRow{ // S31
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(36), // semicolon
			nil,       // if
			nil,       // else
			shift(12), // plus_minus
			nil,       // digits
			nil,       // binary_seq
			nil,       // ident
			nil,       // abc
			nil,       // def
		},
	},
	actionRow{ // S32
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(2), // semicolon, reduce: Stmt
			nil,       // if
			nil,       // else
			nil,       // plus_minus
			nil,       // digits
			nil,       // binary_seq
			nil,       // ident
			nil,       // abc
			nil,       // def
		},
	},
	actionRow{ // S33
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // semicolon
			nil,       // if
			nil,       // else
			nil,       // plus_minus
			shift(15), // digits
			shift(16), // binary_seq
			shift(17), // ident
			shift(18), // abc
			shift(19), // def
		},
	},
	actionRow{ // S34
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // semicolon
			nil,       // if
			shift(38), // else
			nil,       // plus_minus
			nil,       // digits
			nil,       // binary_seq
			nil,       // ident
			nil,       // abc
			nil,       // def
		},
	},
	actionRow{ // S35
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			reduce(3), // $, reduce: If_stmt
			nil,       // semicolon
			nil,       // if
			nil,       // else
			nil,       // plus_minus
			nil,       // digits
			nil,       // binary_seq
			nil,       // ident
			nil,       // abc
			nil,       // def
		},
	},
	actionRow{ // S36
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(1), // semicolon, reduce: Stmt
			nil,       // if
			nil,       // else
			nil,       // plus_minus
			nil,       // digits
			nil,       // binary_seq
			nil,       // ident
			nil,       // abc
			nil,       // def
		},
	},
	actionRow{ // S37
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // semicolon
			shift(24), // if
			nil,       // else
			shift(25), // plus_minus
			shift(6),  // digits
			shift(7),  // binary_seq
			shift(8),  // ident
			shift(9),  // abc
			shift(10), // def
		},
	},
	actionRow{ // S38
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // semicolon
			shift(33), // if
			nil,       // else
			nil,       // plus_minus
			shift(6),  // digits
			shift(7),  // binary_seq
			shift(8),  // ident
			shift(9),  // abc
			shift(10), // def
		},
	},
	actionRow{ // S39
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // semicolon
			nil,       // if
			shift(41), // else
			nil,       // plus_minus
			nil,       // digits
			nil,       // binary_seq
			nil,       // ident
			nil,       // abc
			nil,       // def
		},
	},
	actionRow{ // S40
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(42), // semicolon
			nil,       // if
			nil,       // else
			nil,       // plus_minus
			nil,       // digits
			nil,       // binary_seq
			nil,       // ident
			nil,       // abc
			nil,       // def
		},
	},
	actionRow{ // S41
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // semicolon
			shift(33), // if
			nil,       // else
			nil,       // plus_minus
			shift(6),  // digits
			shift(7),  // binary_seq
			shift(8),  // ident
			shift(9),  // abc
			shift(10), // def
		},
	},
	actionRow{ // S42
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			nil,       // semicolon
			nil,       // if
			reduce(3), // else, reduce: If_stmt
			nil,       // plus_minus
			nil,       // digits
			nil,       // binary_seq
			nil,       // ident
			nil,       // abc
			nil,       // def
		},
	},
	actionRow{ // S43
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			shift(44), // semicolon
			nil,       // if
			nil,       // else
			nil,       // plus_minus
			nil,       // digits
			nil,       // binary_seq
			nil,       // ident
			nil,       // abc
			nil,       // def
		},
	},
	actionRow{ // S44
		canRecover: false,
		actions: [numSymbols]action{
			nil,       // INVALID
			nil,       // $
			reduce(3), // semicolon, reduce: If_stmt
			nil,       // if
			nil,       // else
			nil,       // plus_minus
			nil,       // digits
			nil,       // binary_seq
			nil,       // ident
			nil,       // abc
			nil,       // def
		},
	},
}
