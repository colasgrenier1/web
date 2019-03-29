package main

import (
	"fmt"
	"regexp"
	"time"
)

func main() {
	reg,_ := regexp.Compile(`/static/(\w*)`)
 	t1 := time.Now()
	e := reg.FindStringSubmatch("/static/allo")
	delta := time.Now().Sub(t1)
	fmt.Printf("%v in %f\n", e, delta.Seconds())
	reg,_ = regexp.Compile(`(/([0-9]{4})(?:/([0-9]{2})(?:/(\w+))?)?)$`)
	t1 = time.Now()
	e = reg.FindStringSubmatch("/2019/04/allo-comment-ca-va")
	delta = time.Now().Sub(t1)
	fmt.Printf("%v in %f\n", e, delta.Seconds())
}
