package task

import "github.com/google/wire"

// ProviderManagerSet is task providers.
var ProviderManagerSet = wire.NewSet(NewTaskManager)
