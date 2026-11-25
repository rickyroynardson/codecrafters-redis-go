package store

import "math"

func (s *Store) RPush(key string, values ...string) int {
	for _, v := range values {
		s.List[key] = append(s.List[key], v)
	}
	return len(s.List[key])
}

func (s *Store) LPush(key string, values ...string) int {
	for _, v := range values {
		s.List[key] = append([]string{v}, s.List[key]...)
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

	if start < 0 && int(math.Abs(float64(start))) > len(list) {
		start = 0
	}

	if stop < 0 && int(math.Abs(float64(stop))) > len(list) {
		stop = 0
	}

	if start < 0 {
		start = len(list) + start
	}

	if stop < 0 {
		stop = len(list) + stop
	}

	return list[start : stop+1]
}

func (s *Store) LLen(key string) int {
	list, ok := s.List[key]
	if !ok {
		return 0
	}

	return len(list)
}

func (s *Store) LPop(key string, count int) []string {
	list, ok := s.List[key]
	if !ok || len(list) == 0 {
		return []string{}
	}

	if count > len(list) {
		count = len(list)
	}

	values := list[:count]
	s.List[key] = list[count:]
	return values
}
