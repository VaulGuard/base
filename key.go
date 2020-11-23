package base

type KeyGenerator interface {
	Generate(out ...interface{}) error
}
