package state

type NoItemState struct {

}

func (s *NoItemState) RequestItem() error {
	panic("implement me")
}

func (s *NoItemState) AddItem(count int) error {
	panic("implement me")
}

func (s *NoItemState) InsertMoney(money int) error {
	panic("implement me")
}

func (s *NoItemState) DispenseItem() error {
	panic("implement me")
}