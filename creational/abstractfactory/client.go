package main

type Client struct {
	furnitureFactory FurnitureFactory
}

func NewClient(factory FurnitureFactory) *Client {
	return &Client{furnitureFactory:factory}
}

func (c *Client) SetFurnitureFactory(factory FurnitureFactory) {
	c.furnitureFactory = factory
}