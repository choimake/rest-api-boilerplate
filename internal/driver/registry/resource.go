package registry

type Resource struct {
}

func Initialize() *Resource {
	return &Resource{}
}

func (r *Resource) Finalize() {
}
