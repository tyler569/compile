// Code generated by "stringer -type nodeType"; DO NOT EDIT.

package main

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[number-0]
	_ = x[str-1]
	_ = x[ident-2]
	_ = x[binop-3]
}

const _nodeType_name = "numberstridentbinop"

var _nodeType_index = [...]uint8{0, 6, 9, 14, 19}

func (i nodeType) String() string {
	if i < 0 || i >= nodeType(len(_nodeType_index)-1) {
		return "nodeType(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _nodeType_name[_nodeType_index[i]:_nodeType_index[i+1]]
}