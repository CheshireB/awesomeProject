package pair

type Builder struct {
	pairType int
}

type pair interface {
	Do()
}

func NewBuilder(pairType int) *Builder {
	return &Builder{pairType: pairType}
}

func (b Builder) Build() pair {
	switch b.pairType {
	case 1:
		return NewPairOne()
	case 2:
		return NewPairTwo()
	default:
		return nil
	}
}
