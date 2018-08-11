package board

// Board interface for use in main package
type Board interface {
	printState()
	updateState()
	resetState()
}
