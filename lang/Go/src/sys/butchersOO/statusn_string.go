// Code generated by "stringer -type=statusN -linecomment=true"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Initialized-0]
	_ = x[Positioned-1]
	_ = x[Collecting-2]
	_ = x[Working-3]
	_ = x[Resting-4]
	_ = x[Done-5]
	_ = x[Left-6]
}

const _statusN_name = "_PCWRD."

var _statusN_index = [...]uint8{0, 1, 2, 3, 4, 5, 6, 7}

func (i statusN) String() string {
	if i < 0 || i >= statusN(len(_statusN_index)-1) {
		return "statusN(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _statusN_name[_statusN_index[i]:_statusN_index[i+1]]
}
