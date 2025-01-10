// Semi generic stack that contain data of any type that implements the fmt.Stringer interface
package sgstaxlib

import (
	"fmt"

	"github.com/pluckynumbat/linked-list-stuff-go/sdlistlib"
)

var stackNilError = fmt.Errorf("The stack is nil")
var stackEmptyError = fmt.Errorf("The stack is empty")

type SemiGenericStack[T fmt.Stringer] struct {
	sdlist *sdlistlib.SemiGenericList[T]
}

