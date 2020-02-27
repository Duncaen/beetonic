package spec

import "encoding/xml"

// go:generate goyacc -o gopher.go -p parser gopher.y
type SubsonicResponse struct {
	XMLName  xml.Name `json:"-" xml:"subsonic-response"`
	XMLNS    string   `json:"-" xml:"xmlns,attr"`
	Response `json:"subsonic-response"`
}
type ResponseStatus int

var ResponseStatusValues = []string{"ok", "failed"}

const (
	ResponseStatusOk     ResponseStatus = 0
	ResponseStatusFailed ResponseStatus = 1
)

// make sure ResponseStatus implements Validate
var _ Validate = (*ResponseStatus)(nil)

type Version string

// make sure Version implements Validate
var _ Validate = (*Version)(nil)

type MediaType int

var MediaTypeValues = []string{"music", "podcast", "audiobook", "video"}

const (
	MediaTypeMusic     MediaType = 0
	MediaTypePodcast   MediaType = 1
	MediaTypeAudiobook MediaType = 2
	MediaTypeVideo     MediaType = 3
)

// make sure MediaType implements Validate
var _ Validate = (*MediaType)(nil)

type UserRating int

const (
	UserRatingMin UserRating = 1
	UserRatingMax UserRating = 5
)

// make sure UserRating implements Validate
var _ Validate = (*UserRating)(nil)

type AverageRating float32

const (
	AverageRatingMin AverageRating = 1.0
	AverageRatingMax AverageRating = 5.0
)

// make sure AverageRating implements Validate
var _ Validate = (*AverageRating)(nil)

type PodcastStatus int

var PodcastStatusValues = []string{"new", "downloading", "completed", "error", "deleted", "skipped"}

const (
	PodcastStatusNew         PodcastStatus = 0
	PodcastStatusDownloading PodcastStatus = 1
	PodcastStatusCompleted   PodcastStatus = 2
	PodcastStatusError       PodcastStatus = 3
	PodcastStatusDeleted     PodcastStatus = 4
	PodcastStatusSkipped     PodcastStatus = 5
)

// make sure PodcastStatus implements Validate
var _ Validate = (*PodcastStatus)(nil)

type Response struct {
	Status                ResponseStatus         `json:"status" xml:"status,attr"`
	Version               Version                `json:"version" xml:"version,attr"`
	MusicFolders          *MusicFolders          `json:"musicFolders,omitempty" xml:"musicFolders"`
	Indexes               *Indexes               `json:"indexes,omitempty" xml:"indexes"`
	Directory             *Directory             `json:"directory,omitempty" xml:"directory"`
	Genres                *Genres                `json:"genres,omitempty" xml:"genres"`
	Artists               *ArtistsID3            `json:"artists,omitempty" xml:"artists"`
	Artist                *ArtistWithAlbumsID3   `json:"artist,omitempty" xml:"artist"`
	Album                 *AlbumWithSongsID3     `json:"album,omitempty" xml:"album"`
	Song                  *Child                 `json:"song,omitempty" xml:"song"`
	Videos                *Videos                `json:"videos,omitempty" xml:"videos"`
	VideoInfo             *VideoInfo             `json:"videoInfo,omitempty" xml:"videoInfo"`
	NowPlaying            *NowPlaying            `json:"nowPlaying,omitempty" xml:"nowPlaying"`
	SearchResult          *SearchResult          `json:"searchResult,omitempty" xml:"searchResult"`
	SearchResult2         *SearchResult2         `json:"searchResult2,omitempty" xml:"searchResult2"`
	SearchResult3         *SearchResult3         `json:"searchResult3,omitempty" xml:"searchResult3"`
	Playlists             *Playlists             `json:"playlists,omitempty" xml:"playlists"`
	Playlist              *PlaylistWithSongs     `json:"playlist,omitempty" xml:"playlist"`
	JukeboxStatus         *JukeboxStatus         `json:"jukeboxStatus,omitempty" xml:"jukeboxStatus"`
	JukeboxPlaylist       *JukeboxPlaylist       `json:"jukeboxPlaylist,omitempty" xml:"jukeboxPlaylist"`
	License               *License               `json:"license,omitempty" xml:"license"`
	Users                 *Users                 `json:"users,omitempty" xml:"users"`
	User                  *User                  `json:"user,omitempty" xml:"user"`
	ChatMessages          *ChatMessages          `json:"chatMessages,omitempty" xml:"chatMessages"`
	AlbumList             *AlbumList             `json:"albumList,omitempty" xml:"albumList"`
	AlbumList2            *AlbumList2            `json:"albumList2,omitempty" xml:"albumList2"`
	RandomSongs           *Songs                 `json:"randomSongs,omitempty" xml:"randomSongs"`
	SongsByGenre          *Songs                 `json:"songsByGenre,omitempty" xml:"songsByGenre"`
	Lyrics                *Lyrics                `json:"lyrics,omitempty" xml:"lyrics"`
	Podcasts              *Podcasts              `json:"podcasts,omitempty" xml:"podcasts"`
	NewestPodcasts        *NewestPodcasts        `json:"newestPodcasts,omitempty" xml:"newestPodcasts"`
	InternetRadioStations *InternetRadioStations `json:"internetRadioStations,omitempty" xml:"internetRadioStations"`
	Bookmarks             *Bookmarks             `json:"bookmarks,omitempty" xml:"bookmarks"`
	PlayQueue             *PlayQueue             `json:"playQueue,omitempty" xml:"playQueue"`
	Shares                *Shares                `json:"shares,omitempty" xml:"shares"`
	Starred               *Starred               `json:"starred,omitempty" xml:"starred"`
	Starred2              *Starred2              `json:"starred2,omitempty" xml:"starred2"`
	AlbumInfo             *AlbumInfo             `json:"albumInfo,omitempty" xml:"albumInfo"`
	ArtistInfo            *ArtistInfo            `json:"artistInfo,omitempty" xml:"artistInfo"`
	ArtistInfo2           *ArtistInfo2           `json:"artistInfo2,omitempty" xml:"artistInfo2"`
	SimilarSongs          *SimilarSongs          `json:"similarSongs,omitempty" xml:"similarSongs"`
	SimilarSongs2         *SimilarSongs2         `json:"similarSongs2,omitempty" xml:"similarSongs2"`
	TopSongs              *TopSongs              `json:"topSongs,omitempty" xml:"topSongs"`
	ScanStatus            *ScanStatus            `json:"scanStatus,omitempty" xml:"scanStatus"`
	Error                 *Error                 `json:"error,omitempty" xml:"error"`
}
type MusicFolders struct {
	MusicFolder []MusicFolder `json:"musicFolder" xml:"musicFolder"`
}
type MusicFolder struct {
	Id   int    `json:"id" xml:"id,attr"`
	Name string `json:"name,omitempty" xml:"name,attr,omitempty"`
}
type Indexes struct {
	LastModified    int64    `json:"lastModified" xml:"lastModified,attr"`
	IgnoredArticles string   `json:"ignoredArticles" xml:"ignoredArticles,attr"`
	Shortcut        []Artist `json:"shortcut" xml:"shortcut"`
	Index           []Index  `json:"index" xml:"index"`
	Child           []Child  `json:"child" xml:"child"`
}
type Index struct {
	Name   string   `json:"name" xml:"name,attr"`
	Artist []Artist `json:"artist" xml:"artist"`
}
type Artist struct {
	Id             string        `json:"id" xml:"id,attr"`
	Name           string        `json:"name" xml:"name,attr"`
	ArtistImageUrl string        `json:"artistImageUrl,omitempty" xml:"artistImageUrl,attr,omitempty"`
	Starred        DateTime      `json:"starred,omitempty" xml:"starred,attr,omitempty"`
	UserRating     UserRating    `json:"userRating,omitempty" xml:"userRating,attr,omitempty"`
	AverageRating  AverageRating `json:"averageRating,omitempty" xml:"averageRating,attr,omitempty"`
}
type Genres struct {
	Genre []Genre `json:"genre" xml:"genre"`
}
type Genre struct {
	SongCount  int    `json:"songCount" xml:"songCount,attr"`
	AlbumCount int    `json:"albumCount" xml:"albumCount,attr"`
	Value      string `json:"value" xml:",chardata"`
}
type ArtistsID3 struct {
	IgnoredArticles string     `json:"ignoredArticles" xml:"ignoredArticles,attr"`
	Index           []IndexID3 `json:"index" xml:"index"`
}
type IndexID3 struct {
	Name   string      `json:"name" xml:"name,attr"`
	Artist []ArtistID3 `json:"artist" xml:"artist"`
}
type ArtistID3 struct {
	Id             string   `json:"id" xml:"id,attr"`
	Name           string   `json:"name" xml:"name,attr"`
	CoverArt       string   `json:"coverArt,omitempty" xml:"coverArt,attr,omitempty"`
	ArtistImageUrl string   `json:"artistImageUrl,omitempty" xml:"artistImageUrl,attr,omitempty"`
	AlbumCount     int      `json:"albumCount" xml:"albumCount,attr"`
	Starred        DateTime `json:"starred,omitempty" xml:"starred,attr,omitempty"`
}
type ArtistWithAlbumsID3 struct {
	Id             string     `json:"id" xml:"id,attr"`
	Name           string     `json:"name" xml:"name,attr"`
	CoverArt       string     `json:"coverArt,omitempty" xml:"coverArt,attr,omitempty"`
	ArtistImageUrl string     `json:"artistImageUrl,omitempty" xml:"artistImageUrl,attr,omitempty"`
	AlbumCount     int        `json:"albumCount" xml:"albumCount,attr"`
	Starred        DateTime   `json:"starred,omitempty" xml:"starred,attr,omitempty"`
	Album          []AlbumID3 `json:"album" xml:"album"`
}
type AlbumID3 struct {
	Id        string   `json:"id" xml:"id,attr"`
	Name      string   `json:"name" xml:"name,attr"`
	Artist    string   `json:"artist,omitempty" xml:"artist,attr,omitempty"`
	ArtistId  string   `json:"artistId,omitempty" xml:"artistId,attr,omitempty"`
	CoverArt  string   `json:"coverArt,omitempty" xml:"coverArt,attr,omitempty"`
	SongCount int      `json:"songCount" xml:"songCount,attr"`
	Duration  int      `json:"duration" xml:"duration,attr"`
	PlayCount int64    `json:"playCount,omitempty" xml:"playCount,attr,omitempty"`
	Created   DateTime `json:"created" xml:"created,attr"`
	Starred   DateTime `json:"starred,omitempty" xml:"starred,attr,omitempty"`
	Year      int      `json:"year,omitempty" xml:"year,attr,omitempty"`
	Genre     string   `json:"genre,omitempty" xml:"genre,attr,omitempty"`
}
type AlbumWithSongsID3 struct {
	Id        string   `json:"id" xml:"id,attr"`
	Name      string   `json:"name" xml:"name,attr"`
	Artist    string   `json:"artist,omitempty" xml:"artist,attr,omitempty"`
	ArtistId  string   `json:"artistId,omitempty" xml:"artistId,attr,omitempty"`
	CoverArt  string   `json:"coverArt,omitempty" xml:"coverArt,attr,omitempty"`
	SongCount int      `json:"songCount" xml:"songCount,attr"`
	Duration  int      `json:"duration" xml:"duration,attr"`
	PlayCount int64    `json:"playCount,omitempty" xml:"playCount,attr,omitempty"`
	Created   DateTime `json:"created" xml:"created,attr"`
	Starred   DateTime `json:"starred,omitempty" xml:"starred,attr,omitempty"`
	Year      int      `json:"year,omitempty" xml:"year,attr,omitempty"`
	Genre     string   `json:"genre,omitempty" xml:"genre,attr,omitempty"`
	Song      []Child  `json:"song" xml:"song"`
}
type Videos struct {
	Video []Child `json:"video" xml:"video"`
}
type VideoInfo struct {
	Id         string            `json:"id" xml:"id,attr"`
	Captions   []Captions        `json:"captions" xml:"captions"`
	AudioTrack []AudioTrack      `json:"audioTrack" xml:"audioTrack"`
	Conversion []VideoConversion `json:"conversion" xml:"conversion"`
}
type Captions struct {
	Id   string `json:"id" xml:"id,attr"`
	Name string `json:"name,omitempty" xml:"name,attr,omitempty"`
}
type AudioTrack struct {
	Id           string `json:"id" xml:"id,attr"`
	Name         string `json:"name,omitempty" xml:"name,attr,omitempty"`
	LanguageCode string `json:"languageCode,omitempty" xml:"languageCode,attr,omitempty"`
}
type VideoConversion struct {
	Id           string `json:"id" xml:"id,attr"`
	BitRate      int    `json:"bitRate,omitempty" xml:"bitRate,attr,omitempty"`
	AudioTrackId int    `json:"audioTrackId,omitempty" xml:"audioTrackId,attr,omitempty"`
}
type Directory struct {
	Id            string        `json:"id" xml:"id,attr"`
	Parent        string        `json:"parent,omitempty" xml:"parent,attr,omitempty"`
	Name          string        `json:"name" xml:"name,attr"`
	Starred       DateTime      `json:"starred,omitempty" xml:"starred,attr,omitempty"`
	UserRating    UserRating    `json:"userRating,omitempty" xml:"userRating,attr,omitempty"`
	AverageRating AverageRating `json:"averageRating,omitempty" xml:"averageRating,attr,omitempty"`
	PlayCount     int64         `json:"playCount,omitempty" xml:"playCount,attr,omitempty"`
	Child         []Child       `json:"child" xml:"child"`
}
type Child struct {
	Id                    string        `json:"id" xml:"id,attr"`
	Parent                string        `json:"parent,omitempty" xml:"parent,attr,omitempty"`
	IsDir                 bool          `json:"isDir" xml:"isDir,attr"`
	Title                 string        `json:"title" xml:"title,attr"`
	Album                 string        `json:"album,omitempty" xml:"album,attr,omitempty"`
	Artist                string        `json:"artist,omitempty" xml:"artist,attr,omitempty"`
	Track                 int           `json:"track,omitempty" xml:"track,attr,omitempty"`
	Year                  int           `json:"year,omitempty" xml:"year,attr,omitempty"`
	Genre                 string        `json:"genre,omitempty" xml:"genre,attr,omitempty"`
	CoverArt              string        `json:"coverArt,omitempty" xml:"coverArt,attr,omitempty"`
	Size                  int64         `json:"size,omitempty" xml:"size,attr,omitempty"`
	ContentType           string        `json:"contentType,omitempty" xml:"contentType,attr,omitempty"`
	Suffix                string        `json:"suffix,omitempty" xml:"suffix,attr,omitempty"`
	TranscodedContentType string        `json:"transcodedContentType,omitempty" xml:"transcodedContentType,attr,omitempty"`
	TranscodedSuffix      string        `json:"transcodedSuffix,omitempty" xml:"transcodedSuffix,attr,omitempty"`
	Duration              int           `json:"duration,omitempty" xml:"duration,attr,omitempty"`
	BitRate               int           `json:"bitRate,omitempty" xml:"bitRate,attr,omitempty"`
	Path                  string        `json:"path,omitempty" xml:"path,attr,omitempty"`
	IsVideo               bool          `json:"isVideo,omitempty" xml:"isVideo,attr,omitempty"`
	UserRating            UserRating    `json:"userRating,omitempty" xml:"userRating,attr,omitempty"`
	AverageRating         AverageRating `json:"averageRating,omitempty" xml:"averageRating,attr,omitempty"`
	PlayCount             int64         `json:"playCount,omitempty" xml:"playCount,attr,omitempty"`
	DiscNumber            int           `json:"discNumber,omitempty" xml:"discNumber,attr,omitempty"`
	Created               DateTime      `json:"created,omitempty" xml:"created,attr,omitempty"`
	Starred               DateTime      `json:"starred,omitempty" xml:"starred,attr,omitempty"`
	AlbumId               string        `json:"albumId,omitempty" xml:"albumId,attr,omitempty"`
	ArtistId              string        `json:"artistId,omitempty" xml:"artistId,attr,omitempty"`
	Type                  MediaType     `json:"type,omitempty" xml:"type,attr,omitempty"`
	BookmarkPosition      int64         `json:"bookmarkPosition,omitempty" xml:"bookmarkPosition,attr,omitempty"`
	OriginalWidth         int           `json:"originalWidth,omitempty" xml:"originalWidth,attr,omitempty"`
	OriginalHeight        int           `json:"originalHeight,omitempty" xml:"originalHeight,attr,omitempty"`
}
type NowPlaying struct {
	Entry []NowPlayingEntry `json:"entry" xml:"entry"`
}
type NowPlayingEntry struct {
	Id                    string        `json:"id" xml:"id,attr"`
	Parent                string        `json:"parent,omitempty" xml:"parent,attr,omitempty"`
	IsDir                 bool          `json:"isDir" xml:"isDir,attr"`
	Title                 string        `json:"title" xml:"title,attr"`
	Album                 string        `json:"album,omitempty" xml:"album,attr,omitempty"`
	Artist                string        `json:"artist,omitempty" xml:"artist,attr,omitempty"`
	Track                 int           `json:"track,omitempty" xml:"track,attr,omitempty"`
	Year                  int           `json:"year,omitempty" xml:"year,attr,omitempty"`
	Genre                 string        `json:"genre,omitempty" xml:"genre,attr,omitempty"`
	CoverArt              string        `json:"coverArt,omitempty" xml:"coverArt,attr,omitempty"`
	Size                  int64         `json:"size,omitempty" xml:"size,attr,omitempty"`
	ContentType           string        `json:"contentType,omitempty" xml:"contentType,attr,omitempty"`
	Suffix                string        `json:"suffix,omitempty" xml:"suffix,attr,omitempty"`
	TranscodedContentType string        `json:"transcodedContentType,omitempty" xml:"transcodedContentType,attr,omitempty"`
	TranscodedSuffix      string        `json:"transcodedSuffix,omitempty" xml:"transcodedSuffix,attr,omitempty"`
	Duration              int           `json:"duration,omitempty" xml:"duration,attr,omitempty"`
	BitRate               int           `json:"bitRate,omitempty" xml:"bitRate,attr,omitempty"`
	Path                  string        `json:"path,omitempty" xml:"path,attr,omitempty"`
	IsVideo               bool          `json:"isVideo,omitempty" xml:"isVideo,attr,omitempty"`
	UserRating            UserRating    `json:"userRating,omitempty" xml:"userRating,attr,omitempty"`
	AverageRating         AverageRating `json:"averageRating,omitempty" xml:"averageRating,attr,omitempty"`
	PlayCount             int64         `json:"playCount,omitempty" xml:"playCount,attr,omitempty"`
	DiscNumber            int           `json:"discNumber,omitempty" xml:"discNumber,attr,omitempty"`
	Created               DateTime      `json:"created,omitempty" xml:"created,attr,omitempty"`
	Starred               DateTime      `json:"starred,omitempty" xml:"starred,attr,omitempty"`
	AlbumId               string        `json:"albumId,omitempty" xml:"albumId,attr,omitempty"`
	ArtistId              string        `json:"artistId,omitempty" xml:"artistId,attr,omitempty"`
	Type                  MediaType     `json:"type,omitempty" xml:"type,attr,omitempty"`
	BookmarkPosition      int64         `json:"bookmarkPosition,omitempty" xml:"bookmarkPosition,attr,omitempty"`
	OriginalWidth         int           `json:"originalWidth,omitempty" xml:"originalWidth,attr,omitempty"`
	OriginalHeight        int           `json:"originalHeight,omitempty" xml:"originalHeight,attr,omitempty"`
	Username              string        `json:"username" xml:"username,attr"`
	MinutesAgo            int           `json:"minutesAgo" xml:"minutesAgo,attr"`
	PlayerId              int           `json:"playerId" xml:"playerId,attr"`
	PlayerName            string        `json:"playerName,omitempty" xml:"playerName,attr,omitempty"`
}
type SearchResult struct {
	Offset    int     `json:"offset" xml:"offset,attr"`
	TotalHits int     `json:"totalHits" xml:"totalHits,attr"`
	Match     []Child `json:"match" xml:"match"`
}
type SearchResult2 struct {
	Artist []Artist `json:"artist" xml:"artist"`
	Album  []Child  `json:"album" xml:"album"`
	Song   []Child  `json:"song" xml:"song"`
}
type SearchResult3 struct {
	Artist []ArtistID3 `json:"artist" xml:"artist"`
	Album  []AlbumID3  `json:"album" xml:"album"`
	Song   []Child     `json:"song" xml:"song"`
}
type Playlists struct {
	Playlist []Playlist `json:"playlist" xml:"playlist"`
}
type Playlist struct {
	Id          string   `json:"id" xml:"id,attr"`
	Name        string   `json:"name" xml:"name,attr"`
	Comment     string   `json:"comment,omitempty" xml:"comment,attr,omitempty"`
	Owner       string   `json:"owner,omitempty" xml:"owner,attr,omitempty"`
	Public      bool     `json:"public,omitempty" xml:"public,attr,omitempty"`
	SongCount   int      `json:"songCount" xml:"songCount,attr"`
	Duration    int      `json:"duration" xml:"duration,attr"`
	Created     DateTime `json:"created" xml:"created,attr"`
	Changed     DateTime `json:"changed" xml:"changed,attr"`
	CoverArt    string   `json:"coverArt,omitempty" xml:"coverArt,attr,omitempty"`
	AllowedUser []string `json:"allowedUser" xml:"allowedUser"`
}
type PlaylistWithSongs struct {
	Id          string   `json:"id" xml:"id,attr"`
	Name        string   `json:"name" xml:"name,attr"`
	Comment     string   `json:"comment,omitempty" xml:"comment,attr,omitempty"`
	Owner       string   `json:"owner,omitempty" xml:"owner,attr,omitempty"`
	Public      bool     `json:"public,omitempty" xml:"public,attr,omitempty"`
	SongCount   int      `json:"songCount" xml:"songCount,attr"`
	Duration    int      `json:"duration" xml:"duration,attr"`
	Created     DateTime `json:"created" xml:"created,attr"`
	Changed     DateTime `json:"changed" xml:"changed,attr"`
	CoverArt    string   `json:"coverArt,omitempty" xml:"coverArt,attr,omitempty"`
	AllowedUser []string `json:"allowedUser" xml:"allowedUser"`
	Entry       []Child  `json:"entry" xml:"entry"`
}
type JukeboxStatus struct {
	CurrentIndex int     `json:"currentIndex" xml:"currentIndex,attr"`
	Playing      bool    `json:"playing" xml:"playing,attr"`
	Gain         float64 `json:"gain" xml:"gain,attr"`
	Position     int     `json:"position,omitempty" xml:"position,attr,omitempty"`
}
type JukeboxPlaylist struct {
	CurrentIndex int     `json:"currentIndex" xml:"currentIndex,attr"`
	Playing      bool    `json:"playing" xml:"playing,attr"`
	Gain         float64 `json:"gain" xml:"gain,attr"`
	Position     int     `json:"position,omitempty" xml:"position,attr,omitempty"`
	Entry        []Child `json:"entry" xml:"entry"`
}
type ChatMessages struct {
	ChatMessage []ChatMessage `json:"chatMessage" xml:"chatMessage"`
}
type ChatMessage struct {
	Username string `json:"username" xml:"username,attr"`
	Time     int64  `json:"time" xml:"time,attr"`
	Message  string `json:"message" xml:"message,attr"`
}
type AlbumList struct {
	Album []Child `json:"album" xml:"album"`
}
type AlbumList2 struct {
	Album []AlbumID3 `json:"album" xml:"album"`
}
type Songs struct {
	Song []Child `json:"song" xml:"song"`
}
type Lyrics struct {
	Artist string `json:"artist,omitempty" xml:"artist,attr,omitempty"`
	Title  string `json:"title,omitempty" xml:"title,attr,omitempty"`
	Value  string `json:"value" xml:",chardata"`
}
type Podcasts struct {
	Channel []PodcastChannel `json:"channel" xml:"channel"`
}
type PodcastChannel struct {
	Id               string           `json:"id" xml:"id,attr"`
	Url              string           `json:"url" xml:"url,attr"`
	Title            string           `json:"title,omitempty" xml:"title,attr,omitempty"`
	Description      string           `json:"description,omitempty" xml:"description,attr,omitempty"`
	CoverArt         string           `json:"coverArt,omitempty" xml:"coverArt,attr,omitempty"`
	OriginalImageUrl string           `json:"originalImageUrl,omitempty" xml:"originalImageUrl,attr,omitempty"`
	Status           PodcastStatus    `json:"status" xml:"status,attr"`
	ErrorMessage     string           `json:"errorMessage,omitempty" xml:"errorMessage,attr,omitempty"`
	Episode          []PodcastEpisode `json:"episode" xml:"episode"`
}
type NewestPodcasts struct {
	Episode []PodcastEpisode `json:"episode" xml:"episode"`
}
type PodcastEpisode struct {
	Id                    string        `json:"id" xml:"id,attr"`
	Parent                string        `json:"parent,omitempty" xml:"parent,attr,omitempty"`
	IsDir                 bool          `json:"isDir" xml:"isDir,attr"`
	Title                 string        `json:"title" xml:"title,attr"`
	Album                 string        `json:"album,omitempty" xml:"album,attr,omitempty"`
	Artist                string        `json:"artist,omitempty" xml:"artist,attr,omitempty"`
	Track                 int           `json:"track,omitempty" xml:"track,attr,omitempty"`
	Year                  int           `json:"year,omitempty" xml:"year,attr,omitempty"`
	Genre                 string        `json:"genre,omitempty" xml:"genre,attr,omitempty"`
	CoverArt              string        `json:"coverArt,omitempty" xml:"coverArt,attr,omitempty"`
	Size                  int64         `json:"size,omitempty" xml:"size,attr,omitempty"`
	ContentType           string        `json:"contentType,omitempty" xml:"contentType,attr,omitempty"`
	Suffix                string        `json:"suffix,omitempty" xml:"suffix,attr,omitempty"`
	TranscodedContentType string        `json:"transcodedContentType,omitempty" xml:"transcodedContentType,attr,omitempty"`
	TranscodedSuffix      string        `json:"transcodedSuffix,omitempty" xml:"transcodedSuffix,attr,omitempty"`
	Duration              int           `json:"duration,omitempty" xml:"duration,attr,omitempty"`
	BitRate               int           `json:"bitRate,omitempty" xml:"bitRate,attr,omitempty"`
	Path                  string        `json:"path,omitempty" xml:"path,attr,omitempty"`
	IsVideo               bool          `json:"isVideo,omitempty" xml:"isVideo,attr,omitempty"`
	UserRating            UserRating    `json:"userRating,omitempty" xml:"userRating,attr,omitempty"`
	AverageRating         AverageRating `json:"averageRating,omitempty" xml:"averageRating,attr,omitempty"`
	PlayCount             int64         `json:"playCount,omitempty" xml:"playCount,attr,omitempty"`
	DiscNumber            int           `json:"discNumber,omitempty" xml:"discNumber,attr,omitempty"`
	Created               DateTime      `json:"created,omitempty" xml:"created,attr,omitempty"`
	Starred               DateTime      `json:"starred,omitempty" xml:"starred,attr,omitempty"`
	AlbumId               string        `json:"albumId,omitempty" xml:"albumId,attr,omitempty"`
	ArtistId              string        `json:"artistId,omitempty" xml:"artistId,attr,omitempty"`
	Type                  MediaType     `json:"type,omitempty" xml:"type,attr,omitempty"`
	BookmarkPosition      int64         `json:"bookmarkPosition,omitempty" xml:"bookmarkPosition,attr,omitempty"`
	OriginalWidth         int           `json:"originalWidth,omitempty" xml:"originalWidth,attr,omitempty"`
	OriginalHeight        int           `json:"originalHeight,omitempty" xml:"originalHeight,attr,omitempty"`
	StreamId              string        `json:"streamId,omitempty" xml:"streamId,attr,omitempty"`
	ChannelId             string        `json:"channelId" xml:"channelId,attr"`
	Description           string        `json:"description,omitempty" xml:"description,attr,omitempty"`
	Status                PodcastStatus `json:"status" xml:"status,attr"`
	PublishDate           DateTime      `json:"publishDate,omitempty" xml:"publishDate,attr,omitempty"`
}
type InternetRadioStations struct {
	InternetRadioStation []InternetRadioStation `json:"internetRadioStation" xml:"internetRadioStation"`
}
type InternetRadioStation struct {
	Id          string `json:"id" xml:"id,attr"`
	Name        string `json:"name" xml:"name,attr"`
	StreamUrl   string `json:"streamUrl" xml:"streamUrl,attr"`
	HomePageUrl string `json:"homePageUrl,omitempty" xml:"homePageUrl,attr,omitempty"`
}
type Bookmarks struct {
	Bookmark []Bookmark `json:"bookmark" xml:"bookmark"`
}
type Bookmark struct {
	Position int64    `json:"position" xml:"position,attr"`
	Username string   `json:"username" xml:"username,attr"`
	Comment  string   `json:"comment,omitempty" xml:"comment,attr,omitempty"`
	Created  DateTime `json:"created" xml:"created,attr"`
	Changed  DateTime `json:"changed" xml:"changed,attr"`
	Entry    Child    `json:"entry" xml:"entry"`
}
type PlayQueue struct {
	Current   int      `json:"current,omitempty" xml:"current,attr,omitempty"`
	Position  int64    `json:"position,omitempty" xml:"position,attr,omitempty"`
	Username  string   `json:"username" xml:"username,attr"`
	Changed   DateTime `json:"changed" xml:"changed,attr"`
	ChangedBy string   `json:"changedBy" xml:"changedBy,attr"`
	Entry     []Child  `json:"entry" xml:"entry"`
}
type Shares struct {
	Share []Share `json:"share" xml:"share"`
}
type Share struct {
	Id          string   `json:"id" xml:"id,attr"`
	Url         string   `json:"url" xml:"url,attr"`
	Description string   `json:"description,omitempty" xml:"description,attr,omitempty"`
	Username    string   `json:"username" xml:"username,attr"`
	Created     DateTime `json:"created" xml:"created,attr"`
	Expires     DateTime `json:"expires,omitempty" xml:"expires,attr,omitempty"`
	LastVisited DateTime `json:"lastVisited,omitempty" xml:"lastVisited,attr,omitempty"`
	VisitCount  int      `json:"visitCount" xml:"visitCount,attr"`
	Entry       []Child  `json:"entry" xml:"entry"`
}
type Starred struct {
	Artist []Artist `json:"artist" xml:"artist"`
	Album  []Child  `json:"album" xml:"album"`
	Song   []Child  `json:"song" xml:"song"`
}
type AlbumInfo struct {
	Notes          *string `json:"notes,omitempty" xml:"notes"`
	MusicBrainzId  *string `json:"musicBrainzId,omitempty" xml:"musicBrainzId"`
	LastFmUrl      *string `json:"lastFmUrl,omitempty" xml:"lastFmUrl"`
	SmallImageUrl  *string `json:"smallImageUrl,omitempty" xml:"smallImageUrl"`
	MediumImageUrl *string `json:"mediumImageUrl,omitempty" xml:"mediumImageUrl"`
	LargeImageUrl  *string `json:"largeImageUrl,omitempty" xml:"largeImageUrl"`
}
type ArtistInfoBase struct {
	Biography      *string `json:"biography,omitempty" xml:"biography"`
	MusicBrainzId  *string `json:"musicBrainzId,omitempty" xml:"musicBrainzId"`
	LastFmUrl      *string `json:"lastFmUrl,omitempty" xml:"lastFmUrl"`
	SmallImageUrl  *string `json:"smallImageUrl,omitempty" xml:"smallImageUrl"`
	MediumImageUrl *string `json:"mediumImageUrl,omitempty" xml:"mediumImageUrl"`
	LargeImageUrl  *string `json:"largeImageUrl,omitempty" xml:"largeImageUrl"`
}
type ArtistInfo struct {
	Biography      *string  `json:"biography,omitempty" xml:"biography"`
	MusicBrainzId  *string  `json:"musicBrainzId,omitempty" xml:"musicBrainzId"`
	LastFmUrl      *string  `json:"lastFmUrl,omitempty" xml:"lastFmUrl"`
	SmallImageUrl  *string  `json:"smallImageUrl,omitempty" xml:"smallImageUrl"`
	MediumImageUrl *string  `json:"mediumImageUrl,omitempty" xml:"mediumImageUrl"`
	LargeImageUrl  *string  `json:"largeImageUrl,omitempty" xml:"largeImageUrl"`
	SimilarArtist  []Artist `json:"similarArtist" xml:"similarArtist"`
}
type ArtistInfo2 struct {
	Biography      *string     `json:"biography,omitempty" xml:"biography"`
	MusicBrainzId  *string     `json:"musicBrainzId,omitempty" xml:"musicBrainzId"`
	LastFmUrl      *string     `json:"lastFmUrl,omitempty" xml:"lastFmUrl"`
	SmallImageUrl  *string     `json:"smallImageUrl,omitempty" xml:"smallImageUrl"`
	MediumImageUrl *string     `json:"mediumImageUrl,omitempty" xml:"mediumImageUrl"`
	LargeImageUrl  *string     `json:"largeImageUrl,omitempty" xml:"largeImageUrl"`
	SimilarArtist  []ArtistID3 `json:"similarArtist" xml:"similarArtist"`
}
type SimilarSongs struct {
	Song []Child `json:"song" xml:"song"`
}
type SimilarSongs2 struct {
	Song []Child `json:"song" xml:"song"`
}
type TopSongs struct {
	Song []Child `json:"song" xml:"song"`
}
type Starred2 struct {
	Artist []ArtistID3 `json:"artist" xml:"artist"`
	Album  []AlbumID3  `json:"album" xml:"album"`
	Song   []Child     `json:"song" xml:"song"`
}
type License struct {
	Valid          bool     `json:"valid" xml:"valid,attr"`
	Email          string   `json:"email,omitempty" xml:"email,attr,omitempty"`
	LicenseExpires DateTime `json:"licenseExpires,omitempty" xml:"licenseExpires,attr,omitempty"`
	TrialExpires   DateTime `json:"trialExpires,omitempty" xml:"trialExpires,attr,omitempty"`
}
type ScanStatus struct {
	Scanning bool  `json:"scanning" xml:"scanning,attr"`
	Count    int64 `json:"count,omitempty" xml:"count,attr,omitempty"`
}
type Users struct {
	User []User `json:"user" xml:"user"`
}
type User struct {
	Username            string   `json:"username" xml:"username,attr"`
	Email               string   `json:"email,omitempty" xml:"email,attr,omitempty"`
	ScrobblingEnabled   bool     `json:"scrobblingEnabled" xml:"scrobblingEnabled,attr"`
	MaxBitRate          int      `json:"maxBitRate,omitempty" xml:"maxBitRate,attr,omitempty"`
	AdminRole           bool     `json:"adminRole" xml:"adminRole,attr"`
	SettingsRole        bool     `json:"settingsRole" xml:"settingsRole,attr"`
	DownloadRole        bool     `json:"downloadRole" xml:"downloadRole,attr"`
	UploadRole          bool     `json:"uploadRole" xml:"uploadRole,attr"`
	PlaylistRole        bool     `json:"playlistRole" xml:"playlistRole,attr"`
	CoverArtRole        bool     `json:"coverArtRole" xml:"coverArtRole,attr"`
	CommentRole         bool     `json:"commentRole" xml:"commentRole,attr"`
	PodcastRole         bool     `json:"podcastRole" xml:"podcastRole,attr"`
	StreamRole          bool     `json:"streamRole" xml:"streamRole,attr"`
	JukeboxRole         bool     `json:"jukeboxRole" xml:"jukeboxRole,attr"`
	ShareRole           bool     `json:"shareRole" xml:"shareRole,attr"`
	VideoConversionRole bool     `json:"videoConversionRole" xml:"videoConversionRole,attr"`
	AvatarLastChanged   DateTime `json:"avatarLastChanged,omitempty" xml:"avatarLastChanged,attr,omitempty"`
	Folder              []int    `json:"folder" xml:"folder"`
}
type Error struct {
	Code    int    `json:"code" xml:"code,attr"`
	Message string `json:"message,omitempty" xml:"message,attr,omitempty"`
}
