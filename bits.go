package gotils

type Bits uint64

func Set(b, flag Bits) Bits {
	return (b | flag)
}

func Clear(b, flag Bits) Bits {
	b.Clear(flag)
	return b
}

func Toggle(b, flag Bits) Bits {
	b.Toggle(flag)
	return b
}

func CheckAny(b, flags Bits) bool {
	return b.CheckAny(flags)
}

func IsSetAny(b, flags Bits) bool {
	return b.IsSetAny(flags)
}

func IsSetAll(b, flags Bits) bool {
	return b.IsSetAll(flags)
}

func CheckAll(b, flags Bits) bool {
	return (b & flags) == flags
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

func (t *Bits) CheckAny(flags Bits) bool {
	return ((*t) & flags) != 0
}

func (t *Bits) CheckAll(flags Bits) bool {
	return ((*t) & flags) == flags
}

func (t *Bits) IsSetAny(flags Bits) bool {
	return t.CheckAny(flags)
}

func (t *Bits) IsSetAll(flags Bits) bool {
	return t.CheckAll(flags)
}
