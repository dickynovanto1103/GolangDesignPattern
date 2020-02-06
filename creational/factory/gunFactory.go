package main

import "errors"

type gunFactory struct {

}

func (f *gunFactory) createGun(name string) (gunInterface, error) {
	if name == "ak47" {
		return NewAK47(), nil
	} else if name == "maverick" {
		return NewMaverick(), nil
	} else {
		return nil, errors.New("gun not available")
	}
}
