package store

type Store struct {
	String map[string]Entry
	List   map[string][]string
}

func New() *Store {
	return &Store{
		String: make(map[string]Entry),
		List:   make(map[string][]string),
	}
}

func (s *Store) Type(key string) string {
	return "none"
}
