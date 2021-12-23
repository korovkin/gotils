package gotils_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	gotils "github.com/korovkin/gotils"
)

func TestImplode2(t *testing.T) {
	superman, batman := gotils.Explode2(gotils.Implode2("superman", "batman"))
	assert.Equal(t, superman, "superman", "")
	assert.Equal(t, batman, "batman", "")
}

func TestImplode3(t *testing.T) {
	superman, batman, spiderman := gotils.Explode3(gotils.Implode3("superman", "batman", "spiderman"))
	assert.Equal(t, superman, "superman", "")
	assert.Equal(t, batman, "batman", "")
	assert.Equal(t, spiderman, "spiderman", "")
}
