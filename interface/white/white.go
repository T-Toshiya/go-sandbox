package white

import "github.com/T-Toshiya/go-sandbox/interface/model"

type whiteTaiyaki struct {
	Nakami string
}

// NewWhiteTaiyaki は新しいクリームの鯛焼きを生成する
func NewWhiteTaiyaki() model.Taiyaki {
	return &whiteTaiyaki{}
}

func (t *whiteTaiyaki) GetNakami() string {
	return t.Nakami
}

func (t *whiteTaiyaki) SetNakami(nakami string) {
	t.Nakami = nakami
}
