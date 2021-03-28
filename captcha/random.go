package captcha

import (
	"math/rand"
)

func NewRandomCaptcha() Captcha {
	availables := []func() Captcha{
		NewBannerCaptcha,
		NewMathCaptcha,
	}
	return availables[rand.Intn(len(availables))]() // nolint:gosec
}
