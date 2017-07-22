package sequence

type Sequence interface {
	Query(id string) (uint64, error)
}

