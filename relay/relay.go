package main

type Relay struct {
	s *SQUIC
	c *CQUIC
}

func NewRelay() *Relay {
	return &Relay{
		s: NewSQUIC(),
		c: NewCQUIC(),
	}
}

func (r *Relay) Run() {

}
