package store

func (s *Store) RPush(key string, values ...string) int {
	for _, v := range values {
		s.List[key] = append(s.List[key], v)
	}
	return len(s.List[key])
}
