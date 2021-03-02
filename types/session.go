package types

type Session interface {
	Send(string, string) error
}
