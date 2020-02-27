package subsonic

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"net/http"

	"github.com/Duncaen/beetonic/subsonic/spec"

	"github.com/gorilla/mux"
)

type HandlerFunc func(s *Subsonic, r *http.Request) *spec.SubsonicResponse
type BinaryHandlerFunc func(s *Subsonic, w http.ResponseWriter, r *http.Request)

func (s *Subsonic) H(fn HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		resp := fn(s, r)
		if resp.Error != nil {
			log.Println(resp.Error)
		}
		f := r.FormValue("f")
		switch f {
		case "", "xml":
			bytes, err := xml.MarshalIndent(resp, "", "    ")
			if err != nil {
				resp.Status = spec.ResponseStatusFailed
				resp.Error = &spec.Error{0, err.Error()}
			}
			w.Header().Set("Content-Type", "application/xml")
			w.WriteHeader(http.StatusOK)
			w.Write(bytes)
		case "json":
			bytes, err := json.MarshalIndent(resp, "", "    ")
			if err != nil {
				resp.Status = spec.ResponseStatusFailed
				resp.Error = &spec.Error{0, err.Error()}
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(bytes)
		case "jsonp":
			cb := r.FormValue("callback")
			if cb == "" {
				panic("XXX: handle error")
			}
			bytes, err := json.MarshalIndent(resp, "", "    ")
			if err != nil {
				resp.Status = spec.ResponseStatusFailed
				resp.Error = &spec.Error{0, err.Error()}
			}
			w.Header().Set("Content-Type", "application/javascript")
			w.WriteHeader(http.StatusOK)
			fmt.Fprintf(w, "%s(%s);", cb, bytes)
		}
	}
}

func (s *Subsonic) B(fn BinaryHandlerFunc) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fn(s, w, r)
	}
}

func CORSOrigin(origin string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Access-Control-Allow-Origin", origin)
			w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
			next.ServeHTTP(w, r)
		})
	}
}

var handlers = map[string]HandlerFunc {
	"getAlbum": GetAlbumHandler,
	"getAlbumList": GetAlbumListHandler,
	"getArtist": GetArtistHandler,
	"getArtists": GetArtistsHandler,
	"getLicense": GetLicenseHandler,
	"getGenres": GetGenresHandler,
	"search2": Search2Handler,
	"ping": PingHandler,
}

var binaryHandlers = map[string]BinaryHandlerFunc {
	"getCoverArt": GetCoverArtHandler,
}


func (s *Subsonic) RegisterRoutes(r *mux.Router) {
	sr := r.PathPrefix("/rest").Subrouter()


	for k, fn := range handlers {
		h := s.H(fn)
		sr.HandleFunc(fmt.Sprintf("/%s.view", k), h).
			Methods(http.MethodOptions, http.MethodGet, http.MethodPost)
		sr.HandleFunc(fmt.Sprintf("/%s", k), h).
			Methods(http.MethodOptions, http.MethodGet, http.MethodPost)
	}
	for k, fn := range binaryHandlers {
		h := s.B(fn)
		sr.HandleFunc(fmt.Sprintf("/%s.view", k), h).
			Methods(http.MethodOptions, http.MethodGet, http.MethodPost)
		sr.HandleFunc(fmt.Sprintf("/%s", k), h).
			Methods(http.MethodOptions, http.MethodGet, http.MethodPost)
	}

	// sr.HandleFunc("/ping", s.H(pingHandler)).
	// 	Methods(http.MethodOptions, http.MethodGet, http.MethodPost)
	// sr.HandleFunc("/ping.view", s.H(pingHandler)).
	// 	Methods(http.MethodOptions, http.MethodGet, http.MethodPost)

	// sr.HandleFunc("/getLicense", s.H(licenseHandler)).
	// 	Methods(http.MethodOptions, http.MethodGet, http.MethodPost)
	// sr.HandleFunc("/getLicense.view", s.H(licenseHandler)).
	// 	Methods(http.MethodOptions, http.MethodGet, http.MethodPost)

	// sr.HandleFunc("/getAlbumList", s.H(albumListHandler)).
	// 	Methods(http.MethodOptions, http.MethodGet, http.MethodPost)
	// sr.HandleFunc("/getAlbumList.view", s.H(albumListHandler)).
	// 	Methods(http.MethodOptions, http.MethodGet, http.MethodPost)

	// sr.HandleFunc("/getAlbum", s.H(albumHandler)).
	// 	Methods(http.MethodOptions, http.MethodGet, http.MethodPost)
	// sr.HandleFunc("/getAlbum.view", s.H(albumHandler)).
	// 	Methods(http.MethodOptions, http.MethodGet, http.MethodPost)

	// sr.HandleFunc("/getArtist", s.H(artistHandler)).
	// 	Methods(http.MethodOptions, http.MethodGet, http.MethodPost)
	// sr.HandleFunc("/getArtist.view", s.H(artistHandler)).
	// 	Methods(http.MethodOptions, http.MethodGet, http.MethodPost)

	// sr.HandleFunc("/getArtists", s.H(artistsHandler)).
	// 	Methods(http.MethodOptions, http.MethodGet, http.MethodPost)
	// sr.HandleFunc("/getArtists.view", s.H(artistsHandler)).
	// 	Methods(http.MethodOptions, http.MethodGet, http.MethodPost)

	// sr.HandleFunc("/getCoverArt", s.B(coverArtHandler)).
	// 	Methods(http.MethodOptions, http.MethodGet, http.MethodPost)
	// sr.HandleFunc("/getCoverArt.view", s.B(coverArtHandler)).
	// 	Methods(http.MethodOptions, http.MethodGet, http.MethodPost)

	sr.Use(mux.CORSMethodMiddleware(sr))
	sr.Use(CORSOrigin("*"))
	// r.HandleFunc("/rest/getCoverArt", coverArtHandler).
	// Methods(http.MethodOptions, http.MethodGet)
	// r.HandleFunc("/rest/getCoverArt.view", coverArtHandler).
	// Methods(http.MethodOptions, http.MethodGet)
}
