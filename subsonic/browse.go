package subsonic

import (
	// "log"
	// "unicode/utf8"

	"github.com/Duncaen/beetonic/subsonic/spec"
)

func (s *Subsonic) GetGenres() *spec.SubsonicResponse {
	resp := spec.NewResponse()

	genres, err := s.lib.Genres()
	if err != nil {
		resp.Status = spec.ResponseStatusFailed
		resp.Error = &spec.Error{70, err.Error()}
		return resp
	}

	list := make([]spec.Genre, len(genres))
	for i, genre := range genres {
		list[i] = spec.Genre{
			SongCount: int(genre.SongCount),
			AlbumCount: int(genre.AlbumCount),
			Value: genre.Name,
		}
	}

	resp.Genres = &spec.Genres{
		Genre: list,
	}

	return resp
}
