package subsonic

import (
	"io"
	"log"
	"image"
	"image/png"
	_ "image/jpeg"
	"net/http"

	"github.com/disintegration/imaging"
)


func (s *Subsonic) GetCoverArt(id string, size int) image.Image {
	if id == "" {
		return nil
	}

	var rd io.Reader
	var err error
	if rd, err = s.lib.CoverArt(id); err != nil {
		log.Println("coverArtHandler", err)
		return nil
	}
	if size <= 0 || size > 512 {
		size = 512
	}

	img, _, err := image.Decode(rd)
	if err != nil {
		log.Println("coverArtHandler", err)
		return nil
	}

	// scale down images
	rec := img.Bounds()
	if rec.Max.X - rec.Min.X > size {
		img = imaging.Resize(img, size, 0, imaging.Lanczos)
	} else  if rec.Max.X - rec.Min.X > size {
		img = imaging.Resize(img, 0, size, imaging.Lanczos)
	}

	return img
}

func GetCoverArtHandler(s *Subsonic, w http.ResponseWriter, r *http.Request) {
	img := s.GetCoverArt(r.FormValue("id"), parseInt(r.FormValue("size")))
	if img == nil {
		w.WriteHeader(http.StatusNotFound)
	}
	w.WriteHeader(http.StatusOK)
	if err := png.Encode(w, img); err != nil {
		log.Printf("GetCoverArtHandler:", err)
	}
}
