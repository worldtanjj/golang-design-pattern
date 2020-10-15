package prototype

type Cloneable interface {
	Clone() Cloneable
}

type PrototypeManager struct {
	prototypes map[string]Cloneable
}

func NewPrototypeManager() *PrototypeManager {
	return &PrototypeManager{
		prototypes: make(map[string]Cloneable),
	}
}

func (this *PrototypeManager) Get(name string) Cloneable {
	return this.prototypes[name]
}

func (this *PrototypeManager) Set(name string, prototype Cloneable) {
	this.prototypes[name] = prototype
}
