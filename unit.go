package unit

type Unit struct {
	LeftInput  int32
	RightInput int32
	Output     int32
}

func (u Unit) SetInput(l int32, r int32) {
	u.LeftInput = l
	u.RightInput = r
}
