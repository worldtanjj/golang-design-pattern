package simplefactory

import "fmt"

type API interface {
	Say(name string) string
}

type hiAPI struct{}

func (i hiAPI) Say(name string) string {
	return fmt.Sprintf("Hi, %s", name)
}

type weiAPI struct{}

func (i weiAPI) Say(name string) string {
	return fmt.Sprintf("wei, %s", name)
}

func NewAPI(t int) API {
	if t == 1 {
		return &hiAPI{}
	} else if t == 2 {
		return &weiAPI{}
	}
	return nil
}
