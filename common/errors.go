package common


type TipException struct {
	Reason string
}

func (t *TipException)Error()string {
	return "error happened,reason maybe >> "+t.Reason
}
