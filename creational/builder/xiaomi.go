package main

type Xiaomi struct {
	merk, country string
}

func (x *Xiaomi) SetMerk() {
	x.merk = "xiaomi"
}

func (x *Xiaomi) SetCountry() {
	x.country = "china"
}

func (x *Xiaomi) GetHP() *HP {
	return &HP{
		merk: x.merk,
		country: x.country,
	}
}