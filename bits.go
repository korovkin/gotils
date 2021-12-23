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

func CheckAll(b, flags Bits) bool {
	return (b & flags) == flags
}

func CheckAny(b, flags Bits) bool {
	return (b & flags) != 0
}

func HasAll(b, flags Bits) bool {
	return CheckAll(b, flags)
}

func HasAny(b, flags Bits) bool {
	return CheckAny(b, flags)
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
