// Code generated by "stringer -type=Note -linecomment=true"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[C-1]
	_ = x[Db-2]
	_ = x[D-3]
	_ = x[Eb-4]
	_ = x[E-5]
}

const _Note_name = "261.63277.18293.66311.13329.63"

var _Note_index = [...]uint8{0, 6, 12, 18, 24, 30}

func (i Note) String() string {
	i -= 1
	if i < 0 || i >= Note(len(_Note_index)-1) {
		return "Note(" + strconv.FormatInt(int64(i+1), 10) + ")"
	}
	return _Note_name[_Note_index[i]:_Note_index[i+1]]
}
