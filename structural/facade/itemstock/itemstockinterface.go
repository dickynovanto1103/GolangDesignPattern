package itemstock

type ItemStockInterface interface {
	CheckItemStock(item string) int
}
