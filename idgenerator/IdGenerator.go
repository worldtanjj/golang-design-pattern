package idgenerator

type IdGenerator interface {
	Generate() (string, error)
}
