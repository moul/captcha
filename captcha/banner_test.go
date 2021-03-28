package captcha_test

import (
	"fmt"
	"math/rand"

	"moul.io/captcha/captcha"
)

func ExampleNewBannerCaptcha() {
	rand.Seed(42)
	captcha := captcha.NewBannerCaptcha()
	question, _ := captcha.Question()
	fmt.Println(question)

	valid, _ := captcha.Validate("lorem")
	fmt.Println(valid)

	valid, _ = captcha.Validate("ipsum")
	fmt.Println(valid)

	valid, _ = captcha.Validate("HRUKP")
	fmt.Println(valid)

	valid, _ = captcha.Validate("hrukp")
	fmt.Println(valid)

	valid, _ = captcha.Validate("\n hRuKP\t \n\n \t")
	fmt.Println(valid)

	valid, _ = captcha.Validate("dolor")
	fmt.Println(valid)

	// Output:
	// _                _
	// | |_   _ _  _  _ | |__ _ __
	// | ' \ | '_|| || || / /| '_ \
	// |_||_||_|   \_,_||_\_\| .__/
	//                       |_|
	// false
	// false
	// true
	// true
	// true
	// false
}
