package subsonic

import (
	"github.com/Duncaen/beetonic/library"
)

type Subsonic struct {
	lib library.Library
}

type Option func (*Subsonic) error

func Library(library library.Library) Option {
	return func(s *Subsonic) error {
		s.lib = library
		return nil
	}
}

func New(options ...Option) (*Subsonic, error) {
	s := &Subsonic{}
	for _, opt := range options {
		if err := opt(s); err != nil {
			return nil, err
		}
	}
	return s, nil
}
