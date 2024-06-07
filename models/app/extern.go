package app

type Extern App

func (e *Extern) Export() App {
	return App(*e)
}
