package normal

import "github.com/T-Toshiya/go-sandbox/interface/model"

type normalTaiyaki struct {
	Nakami string
}

// NewNormalTaiyaki は新しい普通の鯛焼きを生成する
func NewNormalTaiyaki() model.Taiyaki {
	return &normalTaiyaki{}
}

func (t *normalTaiyaki) GetNakami() string {
	return t.Nakami
}

func (t *normalTaiyaki) SetNakami(nakami string) {
	t.Nakami = nakami
}
