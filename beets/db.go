package beets

import (
	_ "database/sql"
	"encoding/base64"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/Duncaen/beetonic/library"

	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type Beets struct {
	db *sqlx.DB
}

// disctotal INTEGER,
// albumstatus TEXT
// month INTEGER,
// original_day INTEGER,
// albumartist TEXT,
// year INTEGER
// albumdisambig TEXT
// label TEXT
// id INTEGER PRIMARY KEY
// album TEXT
// asin TEXT
// albumartist_sort TEXT
// script TEXT
// mb_albumid TEXT
// tracktotal INTEGER
// rg_album_gain REAL
// mb_releasegroupid TEXT
// artpath BLOB
// rg_album_peak REAL
// albumartist_credit TEXT
// catalognum TEXT
// added REAL
// original_month INTEGER
// comp INTEGER
// genre TEXT
// day INTEGER
// original_year INTEGER
// language TEXT
// mb_albumartistid TEXT
// country TEXT
// albumtype TEXT
// releasegroupdisambig TEXT
// r128_album_gain INTEGER
type album struct {
	disctotal            int
	albumstatus          string
	month                int
	original_day         int
	albumartist          string
	year                 string
	albumdisambig        string
	label                string
	id                   int
	album                string
	asin                 string
	albumartist_sort     string
	script               string
	mb_albumid           string
	tracktotal           int
	rg_album_gain        float64
	mb_releasegroupid    string
	artpath              []byte
	rg_album_peak        float64
	albumartist_credit   string
	catalognum           string
	added                float64
	original_month       int
	comp                 int
	genre                string
	day                  int
	original_year        int
	language             string
	country              string
	albumtype            string
	releasegroupdisambig string
	r128_track_gain      int
}

type item struct {
}

// lyrics TEXT
// disctitle TEXT
// month INTEGER
// channels INTEGER
// disc INTEGER
// mb_trackid TEXT
// composer TEXT
// albumartist_sort TEXT
// bitdepth INTEGER
// title TEXT
// mb_albumid TEXT
// acoustid_fingerprint TEXT
// rg_album_gain REAL
// mb_releasegroupid TEXT
// rg_album_peak REAL
// albumartist_credit TEXT
// acoustid_id TEXT
// format TEXT
// encoder TEXT
// rg_track_gain REAL
// day INTEGER
// original_year INTEGER
// artist TEXT
// mb_albumartistid TEXT
// bpm INTEGER
// artist_credit TEXT
// grouping TEXT
// disctotal INTEGER
// album_id INTEGER
// albumstatus TEXT
// mtime REAL
// original_day INTEGER
// albumartist TEXT
// year INTEGER
// albumdisambig TEXT
// samplerate INTEGER
// id INTEGER PRIMARY KEY
// album TEXT
// mb_artistid TEXT
// media TEXT
// artist_sort TEXT
// comments TEXT
// tracktotal INTEGER
// rg_track_peak REAL
// catalognum TEXT
// added REAL
// original_month INTEGER
// asin TEXT
// track INTEGER
// comp INTEGER
// initial_key TEXT
// genre TEXT
// path BLOB
// bitrate INTEGER
// language TEXT
// country TEXT
// script TEXT
// label TEXT
// length REAL
// albumtype TEXT
// lyricist TEXT
// composer_sort TEXT
// arranger TEXT
// mb_releasetrackid TEXT
// releasegroupdisambig TEXT
// r128_track_gain INTEGER
// r128_album_gain INTEGER

func Open(path string) (*Beets, error) {
	var err error
	beets := &Beets{}
	beets.db, err = sqlx.Connect("sqlite3", path)
	if err != nil {
		return nil, err
	}
	return beets, nil
}

func (beets *Beets) Close() error {
	return beets.db.Close()
}

const albumsQuery = `
SELECT id, album, albumartist, length(artpath)
FROM albums
%s
LIMIT ? OFFSET ?;
`

var albumsOrderBy = map[library.Order]string{
	library.OrderRandom: `ORDER BY RANDOM() DESC`,
	library.OrderNewest: `ORDER BY added DESC`,
	library.OrderYear: `ORDER BY original_year DESC`,
	library.OrderAlbum: `ORDER BY album ASC`,
	library.OrderArtist: `ORDER BY albumartist_sort DESC`,
}

func (beets *Beets) Albums(limit uint, offset uint, order library.Order) ([]library.Album, error) {
	orderBy, ok := albumsOrderBy[order]
	if !ok {
		log.Println("Order By for: %d not implemented.", order)
	}

	query := fmt.Sprintf(albumsQuery, orderBy)
	rows, err := beets.db.Query(query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := []library.Album{}
	for rows.Next() {
		var album library.Album
		var id int64
		var artlen int
		err := rows.Scan(&id, &album.Title, &album.Artist, &artlen)
		if err != nil {
			panic(err)
		}
		album.Id = fmt.Sprintf("album:%d", id)
		if artlen > 0 {
			album.CoverArt = album.Id
		}
		res = append(res, album)
	}
	return res, nil
}

const albumsByGenreQuery = `
SELECT id, album, albumartist, length(artpath)
FROM albums
WHERE genre = ?
%s
LIMIT ? OFFSET ?;
`

func (beets *Beets) AlbumsByGenre(genre string, limit, offset uint, order library.Order) ([]library.Album, error) {
	orderBy, ok := albumsOrderBy[order]
	if !ok {
		log.Println("Order By for: %d not implemented.", order)
	}

	query := fmt.Sprintf(albumsByGenreQuery, orderBy)
	rows, err := beets.db.Query(query, genre, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := []library.Album{}
	for rows.Next() {
		var album library.Album
		var id int64
		var artlen int
		err := rows.Scan(&id, &album.Title, &album.Artist, &artlen)
		if err != nil {
			panic(err)
		}
		album.Id = fmt.Sprintf("album:%d", id)
		if artlen > 0 {
			album.CoverArt = album.Id
		}
		res = append(res, album)
	}
	return res, nil
}

const albumsByYearQuery = `
SELECT id, album, albumartist, length(artpath)
FROM albums
WHERE year >= ? AND year <= ?
ORDER BY year %s
LIMIT ? OFFSET ?;
`

func (beets *Beets) AlbumsByYear(fromYear, toYear, limit, offset uint) ([]library.Album, error) {
	orderBy := "ASC"
	if fromYear > toYear {
		fromYear, toYear = toYear, fromYear
		orderBy = "DESC"
	}
	query := fmt.Sprintf(albumsByYearQuery, orderBy)
	rows, err := beets.db.Query(query, fromYear, toYear, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	res := []library.Album{}
	for rows.Next() {
		var album library.Album
		var id int64
		var artlen int
		err := rows.Scan(&id, &album.Title, &album.Artist, &artlen)
		if err != nil {
			panic(err)
		}
		album.Id = fmt.Sprintf("album:%d", id)
		if artlen > 0 {
			album.CoverArt = album.Id
		}
		res = append(res, album)
	}
	return res, nil
}

func parseId(id string) (string, int, error) {
	strs := strings.SplitN(id, ":", 2)
	prefix, s := strs[0], strs[1]
	var i int64
	var err error
	if i, err = strconv.ParseInt(s, 10, 32); err != nil {
		return "", 0, err
	}
	return prefix, int(i), nil
}

func albumId(sid string) (int, error) {
	t, id, err := parseId(sid)
	if err != nil {
		return 0, err
	}
	if t != "album" {
		return 0, fmt.Errorf("bad album id")
	}
	return id, nil
}

// SELECT
//  albums.id,
//  albums.album,
//  albums.albumartist,
//  count(items.id),
//  albums.year,
//  albums.genre,
//  length(albums.artpath)
// FROM albums
// LEFT JOIN items
// ON albums.id = items.album_id
// WHERE albums.albumartist = ?
// GROUP BY albums.id;
var albumQuery = `
SELECT
 albums.album,
 albums.albumartist,
 count(items.id),
 albums.year,
 albums.genre,
 sum(items.length),
 length(albums.artpath)
FROM albums
LEFT JOIN items
ON albums.id = items.album_id
WHERE albums.id = ?
LIMIT 1;
`

func (beets *Beets) Album(sid string) (*library.Album, error) {
	id, err := albumId(sid)
	if err != nil {
		return nil, err
	}
	album := library.Album{Id: sid}
	row := beets.db.QueryRow(albumQuery, id)
	var artlen int
	err = row.Scan(&album.Title, &album.Artist, &album.TrackTotal, &album.Year, &album.Genre, &album.Length, &artlen)
	if err != nil {
		return nil, err
	}
	album.Id = fmt.Sprintf("album:%d", id)
	if artlen > 0 {
		album.CoverArt = album.Id
	}
	return &album, nil
}

var songsQuery = `
SELECT id, title, artist, track, year, genre, length
FROM items
WHERE album_id = ?
ORDER BY track ASC;
`

func (beets *Beets) Songs(sid string) ([]library.Song, error) {
	id, err := albumId(sid)
	if err != nil {
		return nil, err
	}
	rows, err := beets.db.Query(songsQuery, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	songs := []library.Song{}
	for rows.Next() {
		var song library.Song
		var id int64
		err := rows.Scan(&id, &song.Title, &song.Artist, &song.Track, &song.Year, &song.Genre, &song.Length)
		if err != nil {
			return nil, err
		}
		song.Id = fmt.Sprintf("track:%d", id)
		songs = append(songs, song)
	}
	return songs, nil
}

var albumArtQuery = `
SELECT artpath FROM albums
WHERE id = ?
LIMIT 1;
`

var trackArtQuery = `
SELECT artpath FROM albums
INNER JOIN items ON albums.id = items.album_id
WHERE items.id = ?
LIMIT 1;
`

func (beets *Beets) CoverArt(sid string) (io.Reader, error) {
	t, id, err := parseId(sid)
	if err != nil {
		return nil, err
	}
	var query string
	switch t {
	case "album":
		query = albumArtQuery
	case "track":
		query = trackArtQuery
	default:
		panic("bad id")
	}
	var path string
	if err := beets.db.Get(&path, query, id); err != nil {
		return nil, err
	}
	return os.Open(path)
}

var albumArtistsQuery = `
SELECT albumartist, count(*)
FROM albums
GROUP BY albums.albumartist
`

func formatArtistId(name string) string {
	return fmt.Sprintf("artist:%s", base64.StdEncoding.EncodeToString([]byte(name)))
}

func (beets *Beets) AlbumArtists() ([]library.AlbumArtist, error) {
	artists := []library.AlbumArtist{}
	rows, err := beets.db.Query(albumArtistsQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var artist library.AlbumArtist
		err := rows.Scan(&artist.Name, &artist.AlbumCount)
		if err != nil {
			return nil, err
		}
		artist.Id = formatArtistId(artist.Name)
		artists = append(artists, artist)
	}
	return artists, nil
}

func artistFromId(id string) (string, error) {
	strs := strings.SplitN(id, ":", 2)
	if strs[0] != "artist" || len(strs) != 2 {
		return "", fmt.Errorf("bad artist id")
	}
	data, err := base64.StdEncoding.DecodeString(strs[1])
	if err != nil {
		return "", err
	}
	return string(data), nil
}

var albumArtistQuery = `
SELECT albumartist, count(*)
FROM albums
WHERE albumartist = ?
GROUP BY albums.albumartist
`

func (beets *Beets) AlbumArtist(sid string) (*library.AlbumArtist, error) {
	artist, err := artistFromId(sid)
	if err != nil {
		return nil, err
	}
	a := library.AlbumArtist{Id: sid}
	row := beets.db.QueryRow(albumArtistQuery, artist)
	err = row.Scan(&a.Name, &a.AlbumCount)
	if err != nil {
		return nil, err
	}
	return &a, nil
}

var albumArtistAlbumsQuery = `
SELECT
 albums.id,
 albums.album,
 albums.albumartist,
 count(items.id),
 albums.year,
 albums.genre,
 length(albums.artpath)
FROM albums
LEFT JOIN items ON albums.id = items.album_id
WHERE albums.albumartist = ?
GROUP BY albums.id;
`

func (beets *Beets) AlbumArtistAlbums(sid string) ([]library.Album, error) {
	artist, err := artistFromId(sid)
	if err != nil {
		return nil, err
	}

	rows, err := beets.db.Query(albumArtistAlbumsQuery, artist)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	albums := []library.Album{}
	for rows.Next() {
		var album library.Album
		var id int64
		var artlen int
		err := rows.Scan(&id, &album.Title, &album.Artist, &album.TrackTotal,
			&album.Year, &album.Genre, &artlen)
		if err != nil {
			return nil, err
		}
		album.Id = fmt.Sprintf("album:%d", id)
		if artlen > 0 {
			album.CoverArt = album.Id
		}
		albums = append(albums, album)
	}
	return albums, nil
}

var albumSearchQuery = `
SELECT
 albums.id,
 albums.album,
 albums.albumartist,
 count(items.id),
 albums.year,
 albums.genre,
 sum(items.length),
 length(albums.artpath)
FROM albums
LEFT JOIN items ON albums.id = items.album_id
WHERE
 albums.album LIKE '%' || ? || '%'
OR
 albums.albumdisambig LIKE '%' || ? || '%'
GROUP BY albums.id
LIMIT ? OFFSET ?;
`

func (beets *Beets) AlbumSearch(query string, limit uint, offset uint) ([]library.Album, error) {
	log.Print("AlbumSearch:", query, limit, offset)
	rows, err := beets.db.Query(albumSearchQuery, query, query, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	albums := []library.Album{}
	for rows.Next() {
		var album library.Album
		var id int64
		var artlen int
		err := rows.Scan(&id, &album.Title, &album.Artist, &album.TrackTotal,
			&album.Year, &album.Genre, &album.Length, &artlen)
		if err != nil {
			return nil, err
		}
		album.Id = fmt.Sprintf("album:%d", id)
		if artlen > 0 {
			album.CoverArt = album.Id
		}
		albums = append(albums, album)
	}
	return albums, nil
}

var genresQuery = `
SELECT genre, SUM(items_count), SUM(albums_count)
FROM (
 SELECT genre, count(id) AS items_count, 0 AS albums_count FROM items
 GROUP BY genre
 UNION ALL
 SELECT genre, 0, count(id) FROM albums GROUP BY genre
 ORDER BY genre
)
GROUP BY genre
`

func (beets *Beets) Genres() ([]library.Genre, error) {
	rows, err := beets.db.Query(genresQuery)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	genres := []library.Genre{}
	for rows.Next() {
		var genre library.Genre
		err := rows.Scan(&genre.Name, &genre.SongCount, &genre.AlbumCount)
		if err != nil {
			return nil, err
		}
		genres = append(genres, genre)
	}
	return genres, nil
}
