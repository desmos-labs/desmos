package types

// DONTCOVER

// SubspacesHooksWrapper is a wrapper for modules to inject SubspacesHooks using depinject.
type SubspacesHooksWrapper struct{ Hooks SubspacesHooks }

// IsOnePerModuleType implements the depinject.OnePerModuleType interface.
func (SubspacesHooksWrapper) IsOnePerModuleType() {}
