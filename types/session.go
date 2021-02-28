package types

type any interface{}

type Session interface {
	Send(string, string) error
}
