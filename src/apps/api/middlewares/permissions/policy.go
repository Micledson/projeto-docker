package permissions

type Policy interface {
	Subject() string
	Object() string
	Action() string
}

type policy struct {
	subject string
	object  string
	action  string
}

func NewPolicy(sub, obj, act string) Policy {
	return &policy{sub, obj, act}
}

func (p *policy) Subject() string {
	return p.subject
}

func (p *policy) Object() string {
	return p.object
}

func (p *policy) Action() string {
	return p.action
}
