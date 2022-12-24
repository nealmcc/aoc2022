package main

import "fmt"

// Ice represents the direction(s) of travel for a whirlwind of snow and ice
// at some location.  There may be multiple bits of ice at the same location.
type Ice int

const (
	None  Ice = 0
	North Ice = 1 << (iota - 1) // ^
	East                        // >
	South                       // v
	West                        // <
)

func (i Ice) Render() byte {
	return [...]byte{
		'.',
		'^', // 1 = North
		'>', // 2 = East
		'2', // 3
		'v', // 4 = South
		'2', // 5
		'2', // 6
		'3', // 7
		'<', // 8 = West
		'2', // 9
		'2', // 10
		'3', // 11 = 8 + 2 + 1
		'2', // 12
		'3', // 13 = 8 + 4 + 1
		'3', // 14 = 8 + 4 + 2
		'4', // 15 = all four bits set
	}[i]
}

// String implements fmt.Stringer.
func (i Ice) String() string {
	return string(i.Render())
}

// ParseIce interprets the given byte as an Ice value.
// It is unable to parse more than once direction at the same location.
func ParseIce(b byte) (Ice, error) {
	switch b {
	case '.':
		return None, nil
	case '^':
		return North, nil
	case '>':
		return East, nil
	case 'v':
		return South, nil
	case '<':
		return West, nil
	default:
		return 0, fmt.Errorf("invalid byte value for Ice: %q", b)
	}
}

// Scan reads the given value and stores it in this Ice variable if possible.
func (i *Ice) Scan(val any) error {
	if val == nil {
		*i = None
		return nil
	}

	if val, ok := val.(int); ok {
		if val < 0 || 15 < val {
			return fmt.Errorf("invalid integer value for Ice: %d", val)
		}
		*i = Ice(val)
		return nil
	}

	if val, ok := val.(byte); ok {
		x, err := ParseIce(val)
		if err != nil {
			return err
		}
		*i = x
		return nil
	}

	if val, ok := val.(string); ok {
		if len(val) != 1 {
			return fmt.Errorf("invalid string value for Ice: %q", val)
		}
		x, err := ParseIce(val[0])
		if err != nil {
			return err
		}
		*i = x
		return nil
	}

	return fmt.Errorf("unsupported type for Ice value: %T", val)
}
