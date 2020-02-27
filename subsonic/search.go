package subsonic

import (
	// "log"
	// "unicode/utf8"

	"github.com/Duncaen/beetonic/subsonic/spec"
)

func (s *Subsonic) Search2(query string, artistCount uint, artistOffset uint,
	albumCount uint, albumOffset uint, songCount uint, songOffset uint,
	musicFolderId string) *spec.SubsonicResponse {
	resp := spec.NewResponse()

	if query == "" {
		resp.Status = spec.ResponseStatusFailed
		resp.Error = &spec.Error{10, "missing \"query\" parameter"}
		return resp
	}
	if artistCount == 0 {
		artistCount = 20
	}
	if albumCount == 0 {
		albumCount = 20
	}
	if songCount == 0 {
		songCount = 20
	}

	albums, err := s.lib.AlbumSearch(query, artistCount, artistOffset)
	if err != nil {
		resp.Status = spec.ResponseStatusFailed
		resp.Error = &spec.Error{70, err.Error()}
		return resp
	}

	albumChilds := make([]spec.Child, len(albums))
	for i, a := range albums {
		albumChilds[i] = spec.Child{
			Id:       a.Id,
			Title:    a.Title,
			Artist:   a.Artist,
			CoverArt: a.CoverArt,
		}
	}

	resp.SearchResult2 = &spec.SearchResult2{
		// Artist: artists,
		Album: albumChilds,
		// Song: songs,
	}

	return resp
}
