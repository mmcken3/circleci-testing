package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIncrNum(t *testing.T) {
	testCases := []struct {
		desc string
		n    int
		incr int
	}{
		{
			desc: "one",
			n:    1,
			incr: 2,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			v := incrNum(tC.n)
			assert.Equal(t, tC.incr, v)
		})
	}
}
