package load

// Engine is the interface that powers the Load Simulations
type Engine interface {
	Run() (ResponseData, error)
	RunAfter() ([]ResponseData, error)
}
