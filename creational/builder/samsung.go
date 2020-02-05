package main

type Samsung struct {
	merk, country string
}

func (s *Samsung) SetMerk() {
	s.merk = "samsung"
}

func (s *Samsung) SetCountry() {
	s.country = "korea"
}

func (s *Samsung) GetHP() *HP {
	return &HP{
		merk: s.merk,
		country: s.country,
	}
}