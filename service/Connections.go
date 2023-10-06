package service

import (
	"sync"
)

var Connections = sync.Map{}
