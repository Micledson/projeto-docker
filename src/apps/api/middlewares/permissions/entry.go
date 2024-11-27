package permissions

type Entry interface {
	Subject() string
	Objects() []string
}

type entry struct {
	subject string
	objects []string
}

func NewEntry(sub string, objects []string) Entry {
	return &entry{sub, objects}
}

func (e *entry) Subject() string {
	return e.subject
}

func (e *entry) Objects() []string {
	return e.objects
}
