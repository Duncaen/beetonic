package library

import (
	"io"
)

type Album struct {
	Id         string
	Title      string
	Artist     string
	TrackTotal int
	Year       int
	Genre      string
	CoverArt   string
	Length     float64
}

type AlbumArtist struct {
	Id         string
	Name       string
	ImageURL   string
	AlbumCount int
}

type Artist struct {
	Id   string
	Name string
}

type Song struct {
	Id     string
	Title  string
	Artist string
	Track  int
	Year   int
	Genre  string
	Length float64
}

type Genre struct {
	Name       string
	SongCount  uint
	AlbumCount uint
}

type Order int

const (
	OrderNewest Order = iota
	OrderRandom
	OrderRecent
	OrderHighest
	OrderFrequent
	OrderAlbum
	OrderArtist
	OrderYear
	OrderGenre
)

type Library interface {
	Album(string) (*Album, error)
	AlbumArtist(string) (*AlbumArtist, error)
	AlbumArtistAlbums(string) ([]Album, error)
	AlbumArtists() ([]AlbumArtist, error)
	AlbumSearch(string, uint, uint) ([]Album, error)
	Albums(uint, uint, Order) ([]Album, error)
	AlbumsByGenre(string, uint, uint, Order) ([]Album, error)
	AlbumsByYear(uint, uint, uint, uint) ([]Album, error)

	// ArtistSearch(string) (*Artist, error)
	CoverArt(string) (io.Reader, error)
	Songs(string) ([]Song, error)

	Genres() ([]Genre, error)
}
