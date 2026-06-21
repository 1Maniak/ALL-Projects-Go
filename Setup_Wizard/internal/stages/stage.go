package stages

// Интерфейс
type Stage interface {
	Name() string
	IsCompleted() (bool, error)
	Run() error
	Status() string
}
