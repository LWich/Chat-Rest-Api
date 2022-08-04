package store

import "database/sql"

type Store struct {
	db             *sql.DB
	userRepository *UserRepository
}

func New(db *sql.DB) *Store {
	return &Store{
		db: db,
	}
}

func (s *Store) User() *UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	repository := &UserRepository{
		store: s,
	}
	s.userRepository = repository

	return repository
}
