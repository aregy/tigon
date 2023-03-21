package keys

import (
	"strings"
)

type Tag uint32

const (
	MaxValue        = ^Tag(0)
	Descriptor1 Tag = 1 << iota
	Descriptor2
	Descriptor3
)

func (tag Tag) String() string {
	switch tag {
	case Descriptor1:
		return "First descriptor"
	case Descriptor2:
		return "Second descriptor"
	case Descriptor3:
		return "Third descriptor"
		// default:
		// 	for i := 1; 1<<i < int(MaxValue); i++ {
		// 		if (1<<i)&tag == 1<<i {
		// 			return fmt.Sprintf("Descriptor #%d", i)
		// 		}
		// 	}
	}
	var description []string
	for i := Tag(1); i < MaxValue; i <<= 1 {
		if i&tag == i {
			description = append(description, Tag(1<<i).String())
		}
	}
	return strings.Join(description, " + ")
}
