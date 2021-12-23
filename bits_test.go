package gotils

import (
	. "github.com/onsi/gomega"

	"github.com/korovkin/gotils"
	"testing"
)

const (
	FLAG_NONE gotils.Bits = 0
	FLAG_ONE  gotils.Bits = 1 << iota
	FLAG_TWO
)

func TestBits(t *testing.T) {
	RegisterTestingT(t)

	t.Run("bits", func(t *testing.T) {
		x := FLAG_NONE
		p := &x

		p.Set(FLAG_ONE)
		Expect(x.Check(FLAG_ONE)).Should(Equal(true))
		Expect(x.Check(FLAG_TWO)).Should(Equal(false))

		p.Toggle(FLAG_TWO)
		Expect(x.Check(FLAG_ONE)).Should(Equal(true))
		Expect(x.Check(FLAG_TWO)).Should(Equal(true))

		p.Clear(FLAG_TWO)
		Expect(x.Check(FLAG_ONE)).Should(Equal(true))
		Expect(x.Check(FLAG_TWO)).Should(Equal(false))

		p.Clear(FLAG_ONE)
		Expect(p.Check(FLAG_ONE)).Should(Equal(false))
		Expect(p.Check(FLAG_TWO)).Should(Equal(false))

		x.Set(FLAG_ONE | FLAG_TWO)
		Expect(p.Check(FLAG_ONE)).Should(Equal(true))
		Expect(p.Check(FLAG_TWO)).Should(Equal(true))
	})

	t.Run("bits_002", func(t *testing.T) {
		x := gotils.Toggle(FLAG_NONE, FLAG_ONE)
		Expect(x).Should(Equal(FLAG_ONE))

		x = gotils.Toggle(x, FLAG_ONE)
		Expect(x).Should(Equal(FLAG_NONE))

		x = FLAG_ONE
		x = gotils.Set(x, FLAG_ONE)
		x = gotils.Set(x, FLAG_TWO)
		Expect(x.Check((FLAG_ONE))).Should(Equal(true))
		Expect(x.Check((FLAG_TWO))).Should(Equal(true))
		Expect(x).Should(Equal(FLAG_TWO | FLAG_ONE))
	})
}
