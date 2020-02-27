package subsonic

import (
	"github.com/Duncaen/beetonic/subsonic/spec"
	"github.com/Duncaen/beetonic/library"
)

// GetAlbum Returns details for an album, including a list of songs. This method organizes music according to ID3 tags.
func (s *Subsonic) GetAlbum(id string) *spec.SubsonicResponse {
	resp := spec.NewResponse()
	if id == "" {
		// XXX: handle error
		resp.Status = spec.ResponseStatusFailed
		resp.Error = &spec.Error{10, "missing \"id\" parameter"}
		return resp
	}
	album, err := s.lib.Album(id)
	if err != nil {
		resp.Status = spec.ResponseStatusFailed
		resp.Error = &spec.Error{70, err.Error()}
		return resp
	}
	songs, err := s.lib.Songs(id)
	if err != nil {
		resp.Status = spec.ResponseStatusFailed
		resp.Error = &spec.Error{70, err.Error()}
		return resp
	}
	childs := make([]spec.Child, len(songs))
	for i, s := range songs {
		childs[i] = spec.Child{
			Id:       s.Id,
			Title:    s.Title,
			Artist:   s.Artist,
			Track:    s.Track,
			Album:    album.Title,
			CoverArt: album.CoverArt,
			Duration: int(s.Length),
		}
	}
	resp.Album = &spec.AlbumWithSongsID3{
		CoverArt:  album.CoverArt,
		Name:      album.Title,
		Artist:    album.Artist,
		Year:      album.Year,
		Genre:     album.Genre,
		Song:      childs,
		SongCount: album.TrackTotal,
		Duration:  int(album.Length),
	}
	return resp
}

type AlbumList struct {
	// Type is the list type.
	Type string
	// The number of albums to return.
	Size uint
	// The list offset.
	Offset uint
	// The first year in the range.
	FromYear uint
	// The last year in the range.
	ToYear uint
	// The genre.
	Genre string
	// (Since 1.11.0) Only return albums in the music folder with the given ID.
	MusicFolderId string
}

var albumListOrder = map[string]library.Order{
	"random": library.OrderRandom,
	"newest": library.OrderNewest,
	"highest": library.OrderHighest,
	"frequent": library.OrderFrequent,
	"recent": library.OrderRecent,
	"alphabeticalByName": library.OrderAlbum,
	"alphabeticalByArtist": library.OrderArtist,
	// TODO: what should the order for byGenre and byYear be?
	"byGenre": library.OrderNewest,
	"byYear": library.OrderNewest,
}

// GetAlbumList returns a list of random, newest, highest rated etc. albums.
func (s *Subsonic) GetAlbumList(al *AlbumList) (resp *spec.SubsonicResponse) {
	resp = spec.NewResponse()

	if al.Size == 0 {
		al.Size = 10
	} else if al.Size > 500 {
		al.Size = 500
	}

	order, ok := albumListOrder[al.Type]
	if !ok {
		resp.Error = &spec.Error{10, "missing \"type\" parameter"}
		return
	}

	switch al.Type {
	case "byYear":
		if al.FromYear == 0 {
			resp.Error = &spec.Error{10, "missing \"fromYear\" parameter"}
			return
		}
		if al.ToYear == 0 {
			resp.Error = &spec.Error{10, "missing \"toYear\" parameter"}
			return
		}
	case "byGenre":
		if al.Genre == "" {
			resp.Error = &spec.Error{10, "missing \"genre\" parameter"}
			return
		}
	}

	var albums []library.Album
	var err error

	switch al.Type {
	case "byYear":
		albums, err = s.lib.AlbumsByYear(al.FromYear, al.ToYear, al.Size, al.Offset)
	case "byGenre":
		albums, err = s.lib.AlbumsByGenre(al.Genre, al.Size, al.Offset, order)
	default:
		albums, err = s.lib.Albums(al.Size, al.Offset, order)
	}
	if err != nil {
		resp.Status = spec.ResponseStatusFailed
		resp.Error = &spec.Error{70, err.Error()}
		return resp
	}

	childs := make([]spec.Child, len(albums))
	for i, a := range albums {
		childs[i] = spec.Child{
			Id:       a.Id,
			Title:    a.Title,
			Artist:   a.Artist,
			CoverArt: a.CoverArt,
		}
	}

	resp.AlbumList = &spec.AlbumList{childs}
	return resp
}
