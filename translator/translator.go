package translator

type Translator interface {
	Translate() (output string, err error)
}
