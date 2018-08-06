package fourcc

import (
	"fmt"
)

// Scan tries to interpret ‘value’ in a way that it can convert it into a Type.
//
// Scan makes Type fit the sql/database.Scanner interface.
func (receiver *Type) Scan(value interface{}) error {

	switch casted := value.(type) {
	case []byte:
		const expectedLength = 4
		if expectedLength != len(casted) {
			return fmt.Errorf("fourcc: scanning from type %T is supported, but expected length be %d, but actually is %d for %q", value, expectedLength, len(casted), casted)
		}

		receiver.value = Uint32(casted[0], casted[1], casted[2], casted[3])

	case string:
		const expectedLength = 4
		if expectedLength != len(casted) {
			return fmt.Errorf("fourcc: scanning from type %T is supported, but expected length be %d, but actually is %d for %q", value, expectedLength, len(casted), casted)
		}

		receiver.value = Uint32(casted[0], casted[1], casted[2], casted[3])

	case int64:
		const minUint32 = 0
		const maxUint32 = 4294967295

		if casted < minUint32 {
			return fmt.Errorf("fourcc: scanning from type %T is supported, but expected value to not be less than %d, but actually was %d.", value, minUint32, casted)
		}
		if maxUint32 < casted  {
			return fmt.Errorf("fourcc: scanning from type %T is supported, but expected value to not be greater than %d, but actually was %d.", value, maxUint32, casted)
		}
		receiver.value = uint32(casted)

	case uint32:
		receiver.value = casted

	case uint64:
		const maxUint32 = 4294967295

		if maxUint32 < casted  {
			return fmt.Errorf("fourcc: scanning from type %T is supported, but expected value to not be greater than %d, but actually was %d.", value, maxUint32, casted)
		}
		receiver.value = uint32(casted)

	default:
		return fmt.Errorf("fourcc: scanning from type %T not supported.", value)
	}

	return nil
}
