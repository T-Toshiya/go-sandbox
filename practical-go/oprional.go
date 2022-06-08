package main

import (
	"fmt"
	"time"
)

type Portion int

const (
	Regular Portion = iota
	Small
	Large
)

type Option struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

type Udon struct {
	men      Portion
	aburaage bool
	ebiten   uint
}

func NewUdon(opt Option) *Udon {
	// ゼロ値に対するデフォルト値処理は関数 / メソッド内部で行う
	if opt.ebiten == 0 && time.Now().Hour() < 10 {
		opt.ebiten = 1
	}

	return &Udon{
		men:      opt.men,
		aburaage: opt.aburaage,
		ebiten:   opt.ebiten,
	}
}

func main() {
	tempuraUdon := NewUdon(Option{
		Large, false, 2,
	})

	fmt.Printf("%v\n", tempuraUdon)
	fmt.Printf("%+v\n", tempuraUdon)
}
