package workouts

// Workout is a set of intervals
type Workout struct {
	Name     string `json:"name"`
	Duration uint32 `json:"duration"`
}

// GetCollection returns an array or something
func GetCollection() []Workout {
	return make([]Workout, 3)
}

// Get return a single workout
func Get() Workout {
	workout := Workout{
		Name:     "super",
		Duration: 3200}

	return workout
}
