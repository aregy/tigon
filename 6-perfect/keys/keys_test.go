package keys

import (
	"testing"
)

func TestRepresentation(t *testing.T) {
	t.Run("Tagged unions lol", func(t *testing.T) {
		descAll := Descriptor1 | Descriptor2 | Descriptor3
		if descAll.String() != "" {
			t.Errorf("Ouch not a repr: %s", descAll.String())
		}
	})
}
