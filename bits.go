package gotils

type Bits uint64

func Set(b, flag Bits) Bits {
	return (b | flag)
}

func Clear(b, flag Bits) Bits {
	return (b &^ flag)
}

func Toggle(b, flag Bits) Bits {
	return (b ^ flag)
}

func Has(b, flag Bits) bool {
	return (b & flag) != 0
}

func Check(b, flag Bits) bool {
	return Has(b, flag)
}

func (t *Bits) Set(flag Bits) {
	*t = ((*t) | flag)
}

func (t *Bits) Clear(flag Bits) {
	*t = ((*t) &^ flag)
}

func (t *Bits) Toggle(flag Bits) {
	*t = ((*t) ^ flag)
}

func (t *Bits) Has(flag Bits) bool {
	return ((*t) & flag) != 0
}

func (t *Bits) Check(flag Bits) bool {
	return t.Has(flag)
}
