package currency

import (
	"testing"
)

func TestConverterCreation(t *testing.T) {
	c, _ := NewConverter()
	c.Age()
}
