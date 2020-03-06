package main

type XiaomiBuilder struct {
	xiaomi *Xiaomi
}

func (x *XiaomiBuilder) Reset() {
	x.xiaomi = &Xiaomi{}
}

func (x *XiaomiBuilder) SetMerk() {
	x.xiaomi.merk = "xiaomi"
}

func (x *XiaomiBuilder) SetCountry() {
	x.xiaomi.country = "china"
}

func (x *XiaomiBuilder) GetHP() *Xiaomi {
	return x.xiaomi
}