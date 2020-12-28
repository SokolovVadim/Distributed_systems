package counter

import (
	"github.com/dati-mipt/distributed-storage-algorithms/network"
)

type BroadcastCounter struct {
	current  int64
	replicas map[int64]network.Link
}

func NewBroadcastCounter() *BroadcastCounter {
	return &BroadcastCounter{
		replicas: map[int64]network.Link{},
	}
}

func (c *BroadcastCounter) Inc() bool {
	c.current++
	for _, p := range c.replicas {
		p.AsyncMessage(struct{}{})
	}
	return true
}

func (c *BroadcastCounter) Read() int64 {
	return c.current
}

func (c *BroadcastCounter) Introduce(rid int64, link network.Link) {
	if link != nil {
		c.replicas[rid] = link
	}
}

func (c *BroadcastCounter) Receive(rid int64, msg interface{}) interface{} {
	c.current++
	return nil
}
