package main

import "testing"

func TestGcd(t *testing.T) {
	testCases := []struct {
		desc string
		a    int
		b    int
		want int
	}{
		{
			desc: "TC1",
			a:    2,
			b:    4,
			want: 2,
		},
		{
			desc: "TC2",
			a:    4,
			b:    2,
			want: 2,
		},
		{
			desc: "TC3",
			a:    1111111,
			b:    1234567,
			want: 1,
		},
		{
			desc: "TC4",
			b:    1111111,
			a:    1234567,
			want: 1,
		},
		{
			desc: "TC5",
			a:    100,
			b:    10,
			want: 10,
		},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			got := gcd(tC.a, tC.b)
			if got != tC.want {
				t.Errorf("got %d, want %d", got, tC.want)
			}
		})
	}
}
