package subsonic

import (
	"net/http"
// 	"strconv"
// 	"io"
// 	"log"
// 	"image"
// 	_ "image/png"
// 	_ "image/jpeg"
// 	"fmt"
// 	"unicode/utf8"

	"github.com/Duncaen/beetonic/subsonic/spec"
// 	"github.com/Duncaen/beetonic/library"

// 	"github.com/disintegration/imaging"
)

func PingHandler(s *Subsonic, r *http.Request) *spec.SubsonicResponse {
	return spec.NewResponse()
}

func GetLicenseHandler(s *Subsonic, r *http.Request) *spec.SubsonicResponse {
	resp := spec.NewResponse()
	resp.License = &spec.License{Valid: true}
	return resp
}

func GetAlbumHandler(s *Subsonic, r *http.Request) *spec.SubsonicResponse {
	return s.GetAlbum(r.FormValue("id"))
}

func GetAlbumListHandler(s *Subsonic, r *http.Request) *spec.SubsonicResponse {
	req := &AlbumList{
		Type: r.FormValue("type"),
		Size: parseUint(r.FormValue("size")),
		Offset: parseUint(r.FormValue("offset")),
		FromYear: parseUint(r.FormValue("fromYear")),
		ToYear: parseUint(r.FormValue("toYear")),
		Genre: r.FormValue("genre"),
	}
	return s.GetAlbumList(req)
}

func GetArtistsHandler(s *Subsonic, r *http.Request) *spec.SubsonicResponse {
	return s.GetArtists(r.FormValue("musicFolderId"))
}

func GetArtistHandler(s *Subsonic, r *http.Request) *spec.SubsonicResponse {
	return s.GetArtist(r.FormValue("id"))
}

func Search2Handler(s *Subsonic, r *http.Request) *spec.SubsonicResponse {
	return s.Search2(
		r.FormValue("query"),
		parseUint(r.FormValue("artistCount")),
		parseUint(r.FormValue("artistOffset")),
		parseUint(r.FormValue("albumCount")),
		parseUint(r.FormValue("albumOffset")),
		parseUint(r.FormValue("songCount")),
		parseUint(r.FormValue("songOffset")),
		r.FormValue("musicFolderId"),
	)
}

func GetGenresHandler(s *Subsonic, r *http.Request) *spec.SubsonicResponse {
	return s.GetGenres()
}
