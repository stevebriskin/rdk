// Code generated by "stringer -type NodeState -trimprefix NodeState"; DO NOT EDIT.

package resource

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[NodeStateUnknown-0]
	_ = x[NodeStateUnconfigured-1]
	_ = x[NodeStateConfiguring-2]
	_ = x[NodeStateReady-3]
	_ = x[NodeStateRemoving-4]
	_ = x[NodeStateUnhealthy-5]
}

const _NodeState_name = "UnknownUnconfiguredConfiguringReadyRemovingUnhealthy"

var _NodeState_index = [...]uint8{0, 7, 19, 30, 35, 43, 52}

func (i NodeState) String() string {
	if i >= NodeState(len(_NodeState_index)-1) {
		return "NodeState(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _NodeState_name[_NodeState_index[i]:_NodeState_index[i+1]]
}
