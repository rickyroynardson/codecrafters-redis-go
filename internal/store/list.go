package store

func (s *Store) RPush(key string, values ...string) int {
	for _, v := range values {
		s.List[key] = append(s.List[key], v)
	}
	return len(s.List[key])
}

func (s *Store) LRange(key string, start, stop int) []string {
	list, ok := s.List[key]
	if !ok {
		return []string{}
	}

	if start >= len(list) {
		return []string{}
	}

	if stop >= len(list) {
		stop = len(list) - 1
	}

	if start > stop && start > 0 && stop > 0 {
		return []string{}
	}

	if start < 0 {
		start = len(list) + start
	}

	if stop < 0 {
		stop = len(list) + stop
	}

	return list[start : stop+1]
}
