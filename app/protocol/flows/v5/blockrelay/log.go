package blockrelay

import (
	"github.com/karlsen-network/karlsend/infrastructure/logger"
	"github.com/karlsen-network/karlsend/util/panics"
)

var log = logger.RegisterSubSystem("PROT")
var spawn = panics.GoroutineWrapperFunc(log)
