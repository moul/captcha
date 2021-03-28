package captcha

import (
	"math/rand"
	"strings"

	"moul.io/banner"
)

type bannerCaptcha struct {
	word string
}

func NewBannerCaptcha() Captcha {
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyz")
	b := make([]rune, 5)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))] // nolint:gosec
	}
	return bannerCaptcha{word: string(b)}
}

func (c bannerCaptcha) Question() (string, error) {
	return banner.Inline(c.word), nil
}

func (c bannerCaptcha) Validate(input string) (bool, error) {
	input = strings.ToLower(input)
	input = strings.TrimSpace(input)
	word := strings.ToLower(c.word)
	return strings.EqualFold(input, word), nil
}

var _ Captcha = (*bannerCaptcha)(nil)
