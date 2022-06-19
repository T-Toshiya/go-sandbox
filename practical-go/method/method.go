package main

import "fmt"

type Struct struct {
	v int
}

type StructWithPointer struct {
	v *int
}

// SetValue インスタンスのレシーバー
// 値を変更してもインスタンスのフィールドは変更されない
func (s Struct) SetValue(v int) {
	s.v = v
}

// SetValueWithPointer ポインターのレシーバー
// 値が変更できる
func (s *Struct) SetValueWithPointer(v int) {
	s.v = v
}

func (sp StructWithPointer) SetValue() {
	*sp.v = 10
}

func main() {
	s := Struct{
		v: 5,
	}

	s.SetValue(3)
	fmt.Printf("s: %v \n", s.v)

	s.SetValueWithPointer(3)
	fmt.Printf("s: %v \n", s.v)

	sp := StructWithPointer{}
	i := 5
	sp.v = &i
	sp.SetValue()
	// 書き変わってしまう。紛らわしいのでこう言ったコードは書かない。
	fmt.Printf("s: %v \n", *sp.v)
}
