package droptypedeclaration

type Dropped struct {
	Name string
	Age  int32
}

func (d *Dropped) Drop() bool {
	return true
}
