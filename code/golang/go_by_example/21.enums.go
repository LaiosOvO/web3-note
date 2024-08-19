package main

import "fmt"

type ServerState int

const (
	StateIdel = iota
	StateConnected
	StateError
	StateRetrying
)

var stateName = map[ServerState]string{
	StateIdel:      "idle",
	StateConnected: "connected",
	StateError:     "error",
	StateRetrying:  "retrying",
}

func (ss ServerState) String() string {
	return stateName[ss]
}

func main() {

	ns := transition(StateIdel)
	fmt.Println(ns)

	ns2 := transition(ns)
	fmt.Println(ns2)

	//ns3 := transition(5)
	//fmt.Println(ns3)

}

func transition(s ServerState) ServerState {
	switch s {
	case StateIdel:
		return StateConnected
	case StateConnected, StateRetrying:
		return StateIdel
	case StateError:
		return StateError
	default:
		panic(fmt.Errorf("unknow state: %s", s))
	}
}
