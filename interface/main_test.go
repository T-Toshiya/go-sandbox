package main

import (
	"testing"
)

type testTaiyaki struct {
	Nakami string
}

func (t *testTaiyaki) GetNakami() string {
	return t.Nakami
}

func (t *testTaiyaki) SetNakami(nakami string) {
	t.Nakami = nakami
}

func Test_makeTaiyaki(t *testing.T) {
	tests := []struct {
		nakami string
	}{
		{nakami: "あんこ"},
		{nakami: "クリーム"},
		{nakami: "いちごジャム"},
		{nakami: "謎ジャム"},
	}
	for i, test := range tests {
		taiyaki := &testTaiyaki{}
		makeTaiyaki(taiyaki, test.nakami)
		if taiyaki.GetNakami() != test.nakami {
			t.Errorf("unexpected nakami(%d): %s != %s", i, taiyaki.GetNakami(), test.nakami)
		}
	}
}
