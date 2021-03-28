package captcha_test

import (
	"fmt"
	"math/rand"

	"moul.io/captcha/captcha"
)

func ExampleNewRandomCaptcha() {
	rand.Seed(42)

	// first captcha
	c := captcha.NewRandomCaptcha()
	question, _ := c.Question()
	fmt.Println(question)

	// second captcha
	c = captcha.NewRandomCaptcha()
	question, _ = c.Question()
	fmt.Println(question)

	// Output:
	// 3 + 4
	//        _    _
	//  _ __ | |_ | |_  _  _  ___
	// | '_ \|  _||  _|| || |/ -_)
	// | .__/ \__| \__| \_,_|\___|
	// |_|
}
