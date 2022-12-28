package gotils_test

import (
	"log"
	"testing"

	. "github.com/onsi/gomega"

	gotils "github.com/korovkin/gotils"
)

func func001() {
	func002()
}

func func002() {
	func003()
}

func func003() {
	name := gotils.DebugRuntimeCallerFuncion("func003")
	log.Println("=> DebugRuntimeCallerFuncion:", name)
	Expect(name).To(BeEquivalentTo("func002"))
}

func TestDebugRuntimeCallerFuncion(t *testing.T) {
	RegisterTestingT(t)

	t.Run("func_name", func(_ *testing.T) {
		func001()
	})
}
