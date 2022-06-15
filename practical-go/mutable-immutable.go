package main

import "math/big"

type Currency string

type MutableMoney struct {
	currency Currency
	amount   *big.Int
}

func (m MutableMoney) Currency() Currency {
	return m.currency
}

func (m *MutableMoney) SetCurrency(c Currency) {
	m.currency = c
}

type ImmutableMoney struct {
	currency Currency
	amount   *big.Int
}

func (im ImmutableMoney) Currency() Currency {
	return im.currency
}

func (im ImmutableMoney) SetCurrency(c Currency) ImmutableMoney {
	// レシーバーを変更しないので、レシーバーは値型でいい
	// 値を変えず、変更を加えた新しいインスタンスを返す
	return ImmutableMoney{
		currency: c,
		amount:   im.amount,
	}
}
