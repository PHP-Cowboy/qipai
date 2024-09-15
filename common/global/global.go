package global

import (
	"go.uber.org/zap"
)

var (
	Logger = make(map[string]*zap.SugaredLogger, 0)
)
