package subsonic

import (
	"testing"
	"encoding/xml"

	. "github.com/Duncaen/beetonic/subsonic/spec"
)

func TestPing(t *testing.T) {
	resp := NewResponse()
	bytes, err := xml.Marshal(&resp)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s\n", bytes)
	var res SubsonicResponse
	err = xml.Unmarshal(bytes, &res)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%#v\n", res)
}

func TestLicense(t *testing.T) {
	resp := NewResponse()
	resp.License = &License{Valid: true}
	bytes, err := xml.Marshal(&resp)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s\n", bytes)
	var res SubsonicResponse
	err = xml.Unmarshal(bytes, &res)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%+v\n", res)
}

func TestMediaTypeValidate(t *testing.T) {
	if err := MediaTypeMusic.Validate(); err != nil {
		t.Error(err)
	}
	if err := MediaType(100).Validate(); err == nil {
		t.Error("expected error")
	}
}

func TestUserRatingValidate(t *testing.T) {
	if err := UserRating(1).Validate(); err != nil {
		t.Error(err)
	}
	if err := UserRating(5).Validate(); err != nil {
		t.Error(err)
	}
	if err := UserRating(0).Validate(); err == nil {
		t.Error("expected error")
	}
	if err := UserRating(6).Validate(); err == nil {
		t.Error("expected error")
	}
}

func TestAverageRatingValidate(t *testing.T) {
	if err := AverageRating(1.0).Validate(); err != nil {
		t.Error(err)
	}
	if err := AverageRating(3.5).Validate(); err != nil {
		t.Error(err)
	}
	if err := AverageRating(5.0).Validate(); err != nil {
		t.Error(err)
	}
	if err := AverageRating(0).Validate(); err == nil {
		t.Error("expected error")
	}
	if err := AverageRating(0.1).Validate(); err == nil {
		t.Error("expected error")
	}
	if err := AverageRating(6).Validate(); err == nil {
		t.Error("expected error")
	}
	if err := AverageRating(5.1).Validate(); err == nil {
		t.Error("expected error")
	}
}
