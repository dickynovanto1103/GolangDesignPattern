package iteminfo

type ItemInfo struct{}

func (i *ItemInfo) CheckItemInfo(item string) bool {
	if item == "book" {
		return true
	}

	return false
}

