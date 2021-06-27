package main

import (
	"fmt"
	"github.com/T-Toshiya/go-sandbox/interface/model"
	"github.com/T-Toshiya/go-sandbox/interface/normal"
	"github.com/T-Toshiya/go-sandbox/interface/white"
)

// ref:https://dokupe.hatenablog.com/entry/20181208/1544246322

func main() {
	t1 := normal.NewNormalTaiyaki()
	t2 := white.NewWhiteTaiyaki()

	makeTaiyaki(t1, "あんこ")
	makeTaiyaki(t2, "クリーム")

	fmt.Println(t1.GetNakami())
	fmt.Println(t2.GetNakami())
}

func makeTaiyaki(t model.Taiyaki, nakami string) model.Taiyaki {
	t.SetNakami(nakami)
	return t
}
