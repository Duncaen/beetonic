package subsonic

import (
	"unicode/utf8"
	"log"

	"github.com/Duncaen/beetonic/subsonic/spec"
)

// GetArtists returns a indexed list of Artists, organized according to ID3 tags.
func (s *Subsonic) GetArtists(musicFolderId string) *spec.SubsonicResponse {
	resp := spec.NewResponse()

	// TODO: support musicFolderId

	artists, err := s.lib.AlbumArtists()
	if err != nil {
		resp.Status = spec.ResponseStatusFailed
		resp.Error = &spec.Error{0, err.Error()}
		return resp
	}

	// articles := []string{"The", "El", "La", "Los", "Las", "Le", "Les"}

	index := make(map[rune][]spec.ArtistID3)

	for _, artist := range artists {
		str := artist.Name
		r, _ := utf8.DecodeRuneInString(str)
		if r == utf8.RuneError {
			r = rune(str[0])
		}
		idx, ok := index[r]
		if !ok {
			idx = []spec.ArtistID3{}
		}
		idx = append(idx, spec.ArtistID3{
			Id: artist.Id,
			Name: artist.Name,
			AlbumCount: artist.AlbumCount,
		})
		index[r] = idx
	}

	list := make([]spec.IndexID3, len(index))
	i := 0
	for r, chunk := range index {
		list[i] = spec.IndexID3{string(r), chunk}
		i++
	}
	resp.Artists = &spec.ArtistsID3{
		IgnoredArticles: "",
		Index: list,
	}
	return resp
}

func (s *Subsonic) GetArtist(id string) *spec.SubsonicResponse {
	resp := spec.NewResponse()

	if id == "" {
		resp.Status = spec.ResponseStatusFailed
		resp.Error = &spec.Error{10, "missing \"id\" parameter"}
		return resp
	}

	artist, err := s.lib.AlbumArtist(id)
	if err != nil {
		log.Println("AlbumArtist:", err)
		resp.Status = spec.ResponseStatusFailed
		resp.Error = &spec.Error{0, err.Error()}
		return resp
	}

	albums, err := s.lib.AlbumArtistAlbums(id)
	if err != nil {
		log.Println("AlbumArtistAlbums:", err)
		resp.Status = spec.ResponseStatusFailed
		resp.Error = &spec.Error{0, err.Error()}
		return resp
	}

	list := make([]spec.AlbumID3, len(albums))
	for i, album := range albums {
		list[i] = spec.AlbumID3{
			Id: album.Id,
			Name: album.Title,
			Artist: album.Artist,
			ArtistId: artist.Id,
			CoverArt: album.CoverArt,
			Year: album.Year,
			Genre: album.Genre,
			SongCount: album.TrackTotal,
		}
	}

	resp.Artist = &spec.ArtistWithAlbumsID3{
		Id: artist.Id,
		Name: artist.Name,
		// CoverArt: artist.CoverArt,
		AlbumCount: artist.AlbumCount,
		Album: list,
	}
	return resp
}
