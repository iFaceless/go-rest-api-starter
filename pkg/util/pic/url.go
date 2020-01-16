package pic

import (
	"fmt"
	"math/rand"
	"net/url"
	"path"
	"strings"
)

const imageDir = "images"

func GetFullURL(token string) string {
	if isURL(token) {
		return token
	}

	return fmt.Sprintf("https://img%d.aizoo.com/%s", rand.Intn(5), token)
}

func parseURLToken(u string) string {
	u = strings.TrimSpace(u)
	if !isURL(u) {
		return u
	}

	pu, err := url.Parse(u)
	if err != nil {
		return u
	}

	subPaths := strings.Split(pu.Path, "/")
	if len(subPaths) == 0 {
		return u
	}

	subPath := subPaths[len(subPaths)-1]
	ext := path.Ext(subPath)
	token := strings.Split(strings.TrimRight(subPath, ext), "_")[0]
	return token + path.Ext(subPath)
}

func isURL(s string) bool {
	hasP := strings.HasPrefix
	return hasP(s, "https://") || hasP(s, "http://") || hasP(s, "://")
}
