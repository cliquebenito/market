package store

import "project/config"

type Store struct {
	config *config.Config
}

func New(config *config.Config) *Store {
	return &Store{
		config: config,
	}
}
func (s *Store) Open() error {

	return nil
}

func (s *Store) Close() error {
	return nil
}
