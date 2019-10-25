package signaler

type empty struct{}

type Signaler struct {
	chn chan empty
}

// New create new signaler construct
func New() *Signaler {
	return &Signaler{
		chn: make(chan empty, 1),
	}
}

// Wait block until signaler is ready
func (s *Signaler) Wait() {
	<-s.chn
}

// Signal release and allow other blocked threads to continue
func (s *Signaler) Signal() {
	s.chn <- empty{}
}
