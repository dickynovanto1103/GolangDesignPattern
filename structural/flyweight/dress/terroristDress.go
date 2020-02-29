package dress

type TerroristDress struct {
	color string
}

func NewTerroristDress(color string) *TerroristDress {
	return &TerroristDress{color: color}
}

func (d *TerroristDress) GetColor() string {
	return d.color
}