package main

import (
	"flag"
	"testing"
)

var (
	sysTest *bool
)

func init() {
	sysTest = flag.Bool("systemTest", false, "set to true when run system test")
}

func TestScore(t *testing.T) {
	cases := []struct {
		Score int
	}{
		{Score: 50},
		{Score: 60},
		{Score: 80},
		{Score: 90},
	}
	for _, c := range cases {
		res := Score(c.Score)
		t.Log(res)
	}
}

func TestSystem(t *testing.T) {
	if *sysTest {
		main()
	}
}
