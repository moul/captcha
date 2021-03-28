package captcha

type Captcha interface {
	Question() (string, error)
	Validate(input string) (bool, error)
}
