package factory

import (
	"sync"
)

var one = &sync.Once{}
