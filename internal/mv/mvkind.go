package mv

import "fmt"

type MVKind uint8

const (
	MVKindOriginal MVKind = iota
	MVKindSEKAI
)

func (k MVKind) String() string {
	switch k {
	case MVKindOriginal:
		return "original"
	case MVKindSEKAI:
		return "sekai"
	default:
		panic(fmt.Sprintf("invalid MVKind: %d", k))
	}
}

func (k *MVKind) Set(value string) error {
	switch value {
	case "original":
		*k = MVKindOriginal
	case "sekai":
		*k = MVKindSEKAI
	default:
		return fmt.Errorf("invalid MVKind: %s", value)
	}
	return nil
}

func (k MVKind) Type() string {
	return "MVKind"
}

func (k MVKind) Other() MVKind {
	switch k {
	case MVKindSEKAI:
		return MVKindOriginal
	case MVKindOriginal:
		return MVKindSEKAI
	default:
		return k // fallback not defined
	}
}
