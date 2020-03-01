package itemstock

type ItemStock struct{}

func (s *ItemStock) CheckItemStock(item string) int {
	if item == "book" {
		return 4
	}
	return 0
}
