package model

type Return struct {
	Structure *ReturnStructure
}

type ReturnStructure struct {
	ArgumentBindings map[string]string // связка аргументов i: itemID, s: saver
}
