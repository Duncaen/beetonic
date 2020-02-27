package spec

import (
	"encoding/xml"
	"time"
	"fmt"
)

func NewResponse() *SubsonicResponse {
	resp := &SubsonicResponse{
		XMLNS: "http://subsonic.org/restapi",
	}
	resp.Version = "1.16.1"
	return resp
}

type DateTime time.Time

func (dt DateTime) MarshalXMLAttr(name xml.Name) (xml.Attr, error) {
	return xml.Attr{name, time.Time(dt).Format(time.RFC3339)}, nil
}

func (dt *DateTime) UnmarshalXMLAttr(attr xml.Attr) error {
	t, err := time.Parse(time.RFC3339, attr.Value)
	if err != nil {
		return err
	}
	*dt = DateTime(t)
	return nil
}

func (rs ResponseStatus) MarshalText() (text []byte, err error) {
	return []byte(ResponseStatusValues[rs]), nil
}

type Validate interface {
	Validate() error
}

func validateFloat32(val, min, max float32) error {
	if val < min || val > max {
		return fmt.Errorf("Number out of range")
	}
	return nil
}

func validateInt(val, min, max int) error {
	if val < min || val > max {
		return fmt.Errorf("Number out of range")
	}
	return nil
}

func (ur UserRating) Validate() error {
	return validateInt(
		int(ur),
		int(UserRatingMin),
		int(UserRatingMax),
	)
}

func (ar AverageRating) Validate() error {
	return validateFloat32(
		float32(ar),
		float32(AverageRatingMin),
		float32(AverageRatingMax),
	)
}

func (rs ResponseStatus) Validate() error {
	if len(ResponseStatusValues) < int(rs) {
		return fmt.Errorf("Invalid ResponseStatus")
	}
	return nil
}

func (mt MediaType) Validate() error {
	if len(MediaTypeValues) < int(mt) {
		return fmt.Errorf("Invalid MediaType")
	}
	return nil
}

func (ps PodcastStatus) Validate() error {
	if len(PodcastStatusValues) < int(ps) {
		return fmt.Errorf("Invalid PodcastStatus")
	}
	return nil
}

func (v Version) Validate() error {
	return nil
}

func (e *Error) Error() string {
	return e.Message
}
