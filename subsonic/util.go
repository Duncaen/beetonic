package subsonic

import (
	"strconv"
)

type SubsonicCtxKey struct{}

func parseUint(s string) uint {
	if u, err := strconv.ParseUint(s, 10, 32); err == nil {
		return uint(u)
	}
	return 0
}

func parseInt(s string) int {
	if i, err := strconv.ParseInt(s, 10, 32); err == nil {
		return int(i)
	}
	return 0
}
