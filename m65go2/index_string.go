// generated by stringer -type=Index; DO NOT EDIT

package m65go2

import "fmt"

const _Index_name = "XY"

var _Index_index = [...]uint8{0, 1, 2}

func (i Index) String() string {
	if i+1 >= Index(len(_Index_index)) {
		return fmt.Sprintf("Index(%d)", i)
	}
	return _Index_name[_Index_index[i]:_Index_index[i+1]]
}
