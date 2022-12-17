package main

// Network is the state of the graph which remains constant.
type Network struct {
	v           map[ValveID]*Valve // all valves
	graph       Graph              // distances between each valve
	maxFlowRate int                // flow if all valves were open
}

// NewNetwork intialises a new network with the given set of valves, which are
// all assumed to be closed, and the given start position.
func NewNetwork(valves map[ValveID]*Valve) Network {
	n := Network{
		v:     valves,
		graph: NewGraph(valves),
	}

	for _, v := range valves {
		if v.Flow != 0 {
			n.maxFlowRate += v.Flow
		}
	}

	return n
}

// TransitionMany moves along the given route.
//
// The start point (not included in the route) is assumed to be 'AA'
func (n Network) TransitionMany(route string) state {
	s := state{
		curr: ID("AA"),
	}
	b := []byte(route)

	for len(b) > 0 {
		next := ID(string(b[:2]))
		b = b[2:]
		s = n.TransitionOne(s, next)
	}

	return s
}

// state holds one set of the variable information as a Network changes over time.
type state struct {
	curr         ValveID // current position
	mins         int     // elapsed time in minutes
	totalFlow    int     // total accumulated flow
	missedFlow   int     // total accumulated opportunity cost
	currFlowRate int     // current flow rate based on open valves
}

// TransitionOne transitions the network from the given state by moving to the
// given valve and opening it. This produces another state.
func (n Network) TransitionOne(s state, next ValveID) state {
	mins, missRate := n.OppCost(s, next)
	s.curr = next
	s.mins += mins
	s.totalFlow += mins * s.currFlowRate
	s.missedFlow += mins * missRate
	s.currFlowRate += n.v[next].Flow
	return s
}

// OppCost measures the opportunity cost of starting in state s, then
// moving to the given valve and opening it.
func (n Network) OppCost(s state, to ValveID) (mins int, rate int) {
	mins = n.graph[s.curr][to].Dist + 1
	rate = n.maxFlowRate - s.currFlowRate
	return mins, rate
}
