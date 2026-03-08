package models
type StartMenuState struct{
	Choices []string
	Cursor int
	Selected map[int]struct{}
	Exiting bool
}