package main

// Network is the current state of the graph at a given point in time.
type Network struct {
	// these values will remain constant throughout the life of the network:
	v           map[ValveID]*Valve // all valves
	routes      Graph              // distances between each valve
	maxFlowRate int                // flow if all valves were open

	// these values will all change over time:
	closed       map[ValveID]struct{} // valves which are still closed
	curr         ValveID              // current position
	minute       int                  // elapsed time in minutes
	totalFlow    int                  // how much flow has accumulated so far
	currFlowRate int                  // current flow rate based on open valves
}

// NewNetwork intialises a new network with the given set of valves, which are
// all assumed to be closed, and the given start position.
func NewNetwork(valves map[ValveID]*Valve) *Network {
	net := &Network{
		v:      valves,
		curr:   ID("AA"),
		closed: make(map[ValveID]struct{}, len(valves)/2),
	}

	net.routes = NewGraph(valves)

	for k, v := range valves {
		if v.Flow != 0 {
			net.maxFlowRate += v.Flow
			net.closed[k] = struct{}{}
		}
	}

	return net
}

// MissedFlow calculates how much flow was *not* achieved vs all valves being
// open from the start.
func (n *Network) MissedFlow() int {
	return n.minute*n.maxFlowRate - n.totalFlow
}

func copymap[K comparable, V any](m map[K]V) map[K]V {
	res := make(map[K]V, len(m))
	for k, v := range m {
		res[k] = v
	}

	return res
}

// costPerMinute finds the sum of the opportunity cost for each of the given
// routes, based on the given state of the network.
// Uses Dijkstra's algorithm.
// func TotalCosts(n *Network, routes []Path) (cost map[Path]OppCost) {
// 	return nil
// }

// OppCost is the opportunity cost of activating a valve.  This includes the
// time (in minutes) to move to the valve and open it, and the potential
// flow that is *not* achieved due to other valves being closed during that time.
type OppCost struct {
	minutes int
	nonFlow int
}

// less returns true iff the opportunity cost a (in flow per minute) is less than b
func less(a, b OppCost) bool {
	adiv, amod := a.nonFlow/a.minutes, a.nonFlow%a.minutes
	bdiv, bmod := b.nonFlow/b.minutes, b.nonFlow%b.minutes

	diff := adiv - bdiv
	if diff != 0 {
		return diff < 0
	}

	diff = amod - bmod
	return diff < 0
}
