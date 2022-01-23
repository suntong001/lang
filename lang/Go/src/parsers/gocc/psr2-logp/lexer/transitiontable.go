// Code generated by gocc; DO NOT EDIT.

package lexer

/*
Let s be the current state
Let r be the current input rune
transitionTable[s](r) returns the next state.
*/
type TransitionTable [NumStates]func(rune) int

var TransTab = TransitionTable{
	// S0
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 1
		case r == 69: // ['E','E']
			return 2
		case r == 72: // ['H','H']
			return 3
		case r == 83: // ['S','S']
			return 4
		}
		return NoState
	},
	// S1
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 5
		}
		return NoState
	},
	// S2
	func(r rune) int {
		switch {
		case r == 110: // ['n','n']
			return 6
		}
		return NoState
	},
	// S3
	func(r rune) int {
		switch {
		case r == 82: // ['R','R']
			return 7
		}
		return NoState
	},
	// S4
	func(r rune) int {
		switch {
		case r == 116: // ['t','t']
			return 8
		}
		return NoState
	},
	// S5
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 9
		}
		return NoState
	},
	// S6
	func(r rune) int {
		switch {
		case r == 100: // ['d','d']
			return 10
		}
		return NoState
	},
	// S7
	func(r rune) int {
		switch {
		case r == 73: // ['I','I']
			return 11
		}
		return NoState
	},
	// S8
	func(r rune) int {
		switch {
		case r == 97: // ['a','a']
			return 12
		}
		return NoState
	},
	// S9
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 13
		}
		return NoState
	},
	// S10
	func(r rune) int {
		switch {
		case r == 32: // [' ',' ']
			return 14
		}
		return NoState
	},
	// S11
	func(r rune) int {
		switch {
		case r == 109: // ['m','m']
			return 15
		}
		return NoState
	},
	// S12
	func(r rune) int {
		switch {
		case r == 114: // ['r','r']
			return 16
		}
		return NoState
	},
	// S13
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 17
		}
		return NoState
	},
	// S14
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 18
		case r == 13: // ['\r','\r']
			return 19
		case r == 111: // ['o','o']
			return 20
		default:
			return 21
		}
	},
	// S15
	func(r rune) int {
		switch {
		case r == 112: // ['p','p']
			return 22
		}
		return NoState
	},
	// S16
	func(r rune) int {
		switch {
		case r == 116: // ['t','t']
			return 23
		}
		return NoState
	},
	// S17
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 24
		}
		return NoState
	},
	// S18
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S19
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 18
		}
		return NoState
	},
	// S20
	func(r rune) int {
		switch {
		case r == 102: // ['f','f']
			return 25
		}
		return NoState
	},
	// S21
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 18
		case r == 13: // ['\r','\r']
			return 19
		default:
			return 21
		}
	},
	// S22
	func(r rune) int {
		switch {
		case r == 111: // ['o','o']
			return 26
		}
		return NoState
	},
	// S23
	func(r rune) int {
		switch {
		case r == 32: // [' ',' ']
			return 27
		}
		return NoState
	},
	// S24
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 28
		}
		return NoState
	},
	// S25
	func(r rune) int {
		switch {
		case r == 32: // [' ',' ']
			return 29
		}
		return NoState
	},
	// S26
	func(r rune) int {
		switch {
		case r == 114: // ['r','r']
			return 30
		}
		return NoState
	},
	// S27
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 31
		case r == 13: // ['\r','\r']
			return 32
		case r == 111: // ['o','o']
			return 33
		default:
			return 34
		}
	},
	// S28
	func(r rune) int {
		switch {
		case r == 45: // ['-','-']
			return 35
		}
		return NoState
	},
	// S29
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 36
		case r == 13: // ['\r','\r']
			return 37
		default:
			return 29
		}
	},
	// S30
	func(r rune) int {
		switch {
		case r == 116: // ['t','t']
			return 38
		}
		return NoState
	},
	// S31
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S32
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 31
		}
		return NoState
	},
	// S33
	func(r rune) int {
		switch {
		case r == 102: // ['f','f']
			return 39
		}
		return NoState
	},
	// S34
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 31
		case r == 13: // ['\r','\r']
			return 32
		default:
			return 34
		}
	},
	// S35
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 40
		}
		return NoState
	},
	// S36
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S37
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 36
		}
		return NoState
	},
	// S38
	func(r rune) int {
		switch {
		case r == 74: // ['J','J']
			return 41
		}
		return NoState
	},
	// S39
	func(r rune) int {
		switch {
		case r == 32: // [' ',' ']
			return 42
		}
		return NoState
	},
	// S40
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 43
		}
		return NoState
	},
	// S41
	func(r rune) int {
		switch {
		case r == 111: // ['o','o']
			return 44
		}
		return NoState
	},
	// S42
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 45
		case r == 13: // ['\r','\r']
			return 46
		default:
			return 42
		}
	},
	// S43
	func(r rune) int {
		switch {
		case r == 32: // [' ',' ']
			return 47
		}
		return NoState
	},
	// S44
	func(r rune) int {
		switch {
		case r == 98: // ['b','b']
			return 48
		}
		return NoState
	},
	// S45
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S46
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 45
		}
		return NoState
	},
	// S47
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 49
		}
		return NoState
	},
	// S48
	func(r rune) int {
		switch {
		case r == 32: // [' ',' ']
			return 50
		}
		return NoState
	},
	// S49
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 51
		}
		return NoState
	},
	// S50
	func(r rune) int {
		switch {
		case r == 74: // ['J','J']
			return 52
		case r == 82: // ['R','R']
			return 53
		}
		return NoState
	},
	// S51
	func(r rune) int {
		switch {
		case r == 58: // [':',':']
			return 54
		}
		return NoState
	},
	// S52
	func(r rune) int {
		switch {
		case r == 111: // ['o','o']
			return 55
		}
		return NoState
	},
	// S53
	func(r rune) int {
		switch {
		case r == 117: // ['u','u']
			return 56
		}
		return NoState
	},
	// S54
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 57
		}
		return NoState
	},
	// S55
	func(r rune) int {
		switch {
		case r == 98: // ['b','b']
			return 58
		}
		return NoState
	},
	// S56
	func(r rune) int {
		switch {
		case r == 110: // ['n','n']
			return 59
		}
		return NoState
	},
	// S57
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 60
		}
		return NoState
	},
	// S58
	func(r rune) int {
		switch {
		case r == 32: // [' ',' ']
			return 61
		}
		return NoState
	},
	// S59
	func(r rune) int {
		switch {
		case r == 110: // ['n','n']
			return 62
		}
		return NoState
	},
	// S60
	func(r rune) int {
		switch {
		case r == 58: // [':',':']
			return 63
		}
		return NoState
	},
	// S61
	func(r rune) int {
		switch {
		case r == 99: // ['c','c']
			return 64
		}
		return NoState
	},
	// S62
	func(r rune) int {
		switch {
		case r == 105: // ['i','i']
			return 65
		}
		return NoState
	},
	// S63
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 66
		}
		return NoState
	},
	// S64
	func(r rune) int {
		switch {
		case r == 111: // ['o','o']
			return 67
		}
		return NoState
	},
	// S65
	func(r rune) int {
		switch {
		case r == 110: // ['n','n']
			return 68
		}
		return NoState
	},
	// S66
	func(r rune) int {
		switch {
		case 48 <= r && r <= 57: // ['0','9']
			return 69
		}
		return NoState
	},
	// S67
	func(r rune) int {
		switch {
		case r == 109: // ['m','m']
			return 70
		}
		return NoState
	},
	// S68
	func(r rune) int {
		switch {
		case r == 103: // ['g','g']
			return 71
		}
		return NoState
	},
	// S69
	func(r rune) int {
		switch {
		case r == 46: // ['.','.']
			return 72
		}
		return NoState
	},
	// S70
	func(r rune) int {
		switch {
		case r == 112: // ['p','p']
			return 73
		}
		return NoState
	},
	// S71
	func(r rune) int {
		switch {
		case r == 32: // [' ',' ']
			return 74
		}
		return NoState
	},
	// S72
	func(r rune) int {
		switch {
		case r == 32: // [' ',' ']
			return 75
		case 48 <= r && r <= 57: // ['0','9']
			return 76
		}
		return NoState
	},
	// S73
	func(r rune) int {
		switch {
		case r == 108: // ['l','l']
			return 77
		}
		return NoState
	},
	// S74
	func(r rune) int {
		switch {
		case r == 106: // ['j','j']
			return 78
		}
		return NoState
	},
	// S75
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S76
	func(r rune) int {
		switch {
		case r == 32: // [' ',' ']
			return 75
		case 48 <= r && r <= 57: // ['0','9']
			return 76
		}
		return NoState
	},
	// S77
	func(r rune) int {
		switch {
		case r == 101: // ['e','e']
			return 79
		}
		return NoState
	},
	// S78
	func(r rune) int {
		switch {
		case r == 111: // ['o','o']
			return 80
		}
		return NoState
	},
	// S79
	func(r rune) int {
		switch {
		case r == 116: // ['t','t']
			return 81
		}
		return NoState
	},
	// S80
	func(r rune) int {
		switch {
		case r == 98: // ['b','b']
			return 82
		}
		return NoState
	},
	// S81
	func(r rune) int {
		switch {
		case r == 101: // ['e','e']
			return 83
		}
		return NoState
	},
	// S82
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 84
		case r == 13: // ['\r','\r']
			return 85
		}
		return NoState
	},
	// S83
	func(r rune) int {
		switch {
		case r == 100: // ['d','d']
			return 86
		}
		return NoState
	},
	// S84
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S85
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 84
		}
		return NoState
	},
	// S86
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 87
		case r == 13: // ['\r','\r']
			return 88
		}
		return NoState
	},
	// S87
	func(r rune) int {
		switch {
		}
		return NoState
	},
	// S88
	func(r rune) int {
		switch {
		case r == 10: // ['\n','\n']
			return 87
		}
		return NoState
	},
}
