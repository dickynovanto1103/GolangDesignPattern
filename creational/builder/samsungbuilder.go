package main

type SamsungBuilder struct {
	samsung *Samsung
}

func (s *SamsungBuilder) Reset() {
	s.samsung = &Samsung{}
}

func (s *SamsungBuilder) SetMerk() {
	s.samsung.merk = "samsung"
}

func (s *SamsungBuilder) SetCountry() {
	s.samsung.country = "korea"
}

func (s *SamsungBuilder) GetHP() *Samsung {
	return s.samsung
}