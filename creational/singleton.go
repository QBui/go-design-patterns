package creational

// Singleton data structure
type Singleton struct {
	count int
}

var instance *Singleton

// GetInstance Retrieve a singleton object
func GetInstance() *Singleton {
	if instance == nil {
		instance = new(Singleton)
	}

	return instance
}

// AddOne demonstrate singleton operation
func (s *Singleton) AddOne() int {
	s.count++
	return s.count
}

// GetCount demonstrate singleton operation
func (s *Singleton) GetCount() int {
	return s.count
}
