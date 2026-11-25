package store

type Store struct {
	String         map[string]Entry
	List           map[string][]string
	BlockedClients map[string][]chan []string
}

func New() *Store {
	return &Store{
		String:         make(map[string]Entry),
		List:           make(map[string][]string),
		BlockedClients: make(map[string][]chan []string),
	}
}

func (s *Store) Type(key string) string {
	return "none"
}
