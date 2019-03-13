package components

// AI contains the current decision by the AI system.
type AI struct {
	Down  bool
	Up    bool
	Left  bool
	Right bool
}

// Name ...
func (i *AI) Name() string {
	return "ai"
}
