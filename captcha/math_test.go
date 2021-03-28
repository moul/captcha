package captcha_test

import (
	"fmt"
	"math/rand"

	"moul.io/captcha/captcha"
)

func ExampleNewMathCaptcha() {
	rand.Seed(42)
	captcha := captcha.NewMathCaptcha()
	question, _ := captcha.Question()
	fmt.Println(question)

	valid, _ := captcha.Validate("42")
	fmt.Println(valid)

	valid, _ = captcha.Validate("lorem ipsum")
	fmt.Println(valid)

	valid, _ = captcha.Validate("1 + 3")
	fmt.Println(valid)

	valid, _ = captcha.Validate("4")
	fmt.Println(valid)

	valid, _ = captcha.Validate("\n 4\t \n\n \t")
	fmt.Println(valid)

	// Output:
	// 1 + 3
	// false
	// false
	// false
	// true
	// true
}
