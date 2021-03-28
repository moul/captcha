package captcha

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
)

type mathCaptcha struct {
	equation string
	result   int
}

func NewMathCaptcha() Captcha {
	a := rand.Intn(5) + 1 // nolint:gosec,gomnd
	b := rand.Intn(5) + 1 // nolint:gosec,gomnd
	equation := fmt.Sprintf("%d + %d", a, b)
	result := a + b
	return mathCaptcha{equation, result}
}

func (c mathCaptcha) Question() (string, error) {
	return c.equation, nil
}

func (c mathCaptcha) Validate(input string) (bool, error) {
	input = strings.TrimSpace(input)
	if input == "" {
		return false, nil
	}
	nb, err := strconv.Atoi(input)
	if err != nil {
		return false, fmt.Errorf("invalid input: %w", err)
	}
	return nb == c.result, nil
}

var _ Captcha = (*mathCaptcha)(nil)
