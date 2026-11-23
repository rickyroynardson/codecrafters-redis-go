package store

import "time"

type Entry struct {
	Value string
	Exp   time.Time
}

func (s *Store) Set(key, value string, exp time.Time) {
	s.String[key] = Entry{
		Value: value,
		Exp:   exp,
	}
}

func (s *Store) Get(key string) (string, bool) {
	entry, ok := s.String[key]
	if !ok {
		return "", false
	}

	if entry.Exp.IsZero() || entry.Exp.After(time.Now()) {
		return entry.Value, true
	}

	delete(s.String, key)
	return "", false
}
