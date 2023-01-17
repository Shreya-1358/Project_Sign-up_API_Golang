package model

type NewInterface interface {
	f1()
}

type Check struct{}

func (s *Check) f1() {

}
