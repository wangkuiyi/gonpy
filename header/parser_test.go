package header

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	m := Parse(strings.NewReader("{ 'descr' : '<f4', 'fortran_order': False, 'shape': (51, 3000), }"))
	assert.Equal(t, "<f4", m["descr"])
	assert.Equal(t, false, m["fortran_order"])
	assert.Equal(t, 51, m["shape"].(*Dim).x)
	assert.Equal(t, 3000, m["shape"].(*Dim).y)
}
