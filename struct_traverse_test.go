package gotils_test

import (
	"testing"

	. "github.com/onsi/gomega"

	"github.com/korovkin/gotils"
)

type MyTraversedObject struct {
	Public  string `json:"public"`
	Private string `json:"private"`
}

type MyTraversedObjectCollection struct {
	A []*MyTraversedObject          `json:"a"`
	M map[string]*MyTraversedObject `json:"m"`
	O *MyTraversedObject            `json:"o"`
}

func (t *MyTraversedObject) StructTraverseVisitor(ctx *gotils.StructTraverseTraverser) error {
	t.Private = "hidden"
	return nil
}

func TestStructTraverse(t *testing.T) {
	RegisterTestingT(t)

	t.Run("traverse_modify", func(_ *testing.T) {
		// a few copes of the object to be manipulated:
		c := &MyTraversedObjectCollection{
			A: []*MyTraversedObject{
				{Public: "pub", Private: "private"},
				{Public: "pub", Private: "private"},
			},
			M: map[string]*MyTraversedObject{
				"001": {Public: "pub", Private: "private"},
				"002": {Public: "pub", Private: "private"},
			},
			O: &MyTraversedObject{Public: "pub", Private: "private"},
		}

		// "private" fields are still there:
		parsed := &MyTraversedObjectCollection{}
		gotils.FromJSONString(gotils.ToJSONString(c), parsed)
		Expect(parsed.O.Private).To(BeEquivalentTo("private"))
		Expect(parsed.O.Public).To(BeEquivalentTo("pub"))
		Expect(parsed.A[0].Private).To(BeEquivalentTo("private"))
		Expect(parsed.A[1].Private).To(BeEquivalentTo("private"))
		Expect(parsed.M["001"].Private).To(BeEquivalentTo("private"))
		Expect(parsed.M["002"].Private).To(BeEquivalentTo("private"))

		// traverse the struct recursively and "hide" the private field:
		ctx := gotils.CreateStructTraverseContext(map[string]interface{}{})
		err := ctx.Traverse(parsed)
		Expect(err).To(BeNil())
		Expect(parsed.O.Public).To(BeEquivalentTo("pub"))
		Expect(parsed.O.Private).To(BeEquivalentTo("hidden"))
		Expect(parsed.A[0].Private).To(BeEquivalentTo("hidden"))
		Expect(parsed.A[1].Private).To(BeEquivalentTo("hidden"))
		Expect(parsed.M["001"].Private).To(BeEquivalentTo("hidden"))
		Expect(parsed.M["002"].Private).To(BeEquivalentTo("hidden"))
	})
}
