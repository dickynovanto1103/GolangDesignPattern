package state

type HasItemState struct {

}

func (s *HasItemState) RequestItem() error {
	panic("implement me")
}

func (s *HasItemState) AddItem(count int) error {
	panic("implement me")
}

func (s *HasItemState) InsertMoney(money int) error {
	panic("implement me")
}

func (s *HasItemState) DispenseItem() error {
	panic("implement me")
}