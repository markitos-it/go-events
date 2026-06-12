package types

type Logger interface {
	Info(message string)
	Error(message string)
	Fatal(message string)
	OpenGroup(name string)
	CloseGroup(name string)
}
