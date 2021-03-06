// Code generated by msgraph-generate.go DO NOT EDIT.

package msgraph

// ConfigurationManagerActionType undocumented
type ConfigurationManagerActionType string

const (
	// ConfigurationManagerActionTypeVRefreshMachinePolicy undocumented
	ConfigurationManagerActionTypeVRefreshMachinePolicy ConfigurationManagerActionType = "refreshMachinePolicy"
	// ConfigurationManagerActionTypeVRefreshUserPolicy undocumented
	ConfigurationManagerActionTypeVRefreshUserPolicy ConfigurationManagerActionType = "refreshUserPolicy"
	// ConfigurationManagerActionTypeVWakeUpClient undocumented
	ConfigurationManagerActionTypeVWakeUpClient ConfigurationManagerActionType = "wakeUpClient"
	// ConfigurationManagerActionTypeVAppEvaluation undocumented
	ConfigurationManagerActionTypeVAppEvaluation ConfigurationManagerActionType = "appEvaluation"
)

var (
	// ConfigurationManagerActionTypePRefreshMachinePolicy is a pointer to ConfigurationManagerActionTypeVRefreshMachinePolicy
	ConfigurationManagerActionTypePRefreshMachinePolicy = &_ConfigurationManagerActionTypePRefreshMachinePolicy
	// ConfigurationManagerActionTypePRefreshUserPolicy is a pointer to ConfigurationManagerActionTypeVRefreshUserPolicy
	ConfigurationManagerActionTypePRefreshUserPolicy = &_ConfigurationManagerActionTypePRefreshUserPolicy
	// ConfigurationManagerActionTypePWakeUpClient is a pointer to ConfigurationManagerActionTypeVWakeUpClient
	ConfigurationManagerActionTypePWakeUpClient = &_ConfigurationManagerActionTypePWakeUpClient
	// ConfigurationManagerActionTypePAppEvaluation is a pointer to ConfigurationManagerActionTypeVAppEvaluation
	ConfigurationManagerActionTypePAppEvaluation = &_ConfigurationManagerActionTypePAppEvaluation
)

var (
	_ConfigurationManagerActionTypePRefreshMachinePolicy = ConfigurationManagerActionTypeVRefreshMachinePolicy
	_ConfigurationManagerActionTypePRefreshUserPolicy    = ConfigurationManagerActionTypeVRefreshUserPolicy
	_ConfigurationManagerActionTypePWakeUpClient         = ConfigurationManagerActionTypeVWakeUpClient
	_ConfigurationManagerActionTypePAppEvaluation        = ConfigurationManagerActionTypeVAppEvaluation
)

// ConfigurationManagerClientState undocumented
type ConfigurationManagerClientState string

const (
	// ConfigurationManagerClientStateVUnknown undocumented
	ConfigurationManagerClientStateVUnknown ConfigurationManagerClientState = "unknown"
	// ConfigurationManagerClientStateVInstalled undocumented
	ConfigurationManagerClientStateVInstalled ConfigurationManagerClientState = "installed"
	// ConfigurationManagerClientStateVHealthy undocumented
	ConfigurationManagerClientStateVHealthy ConfigurationManagerClientState = "healthy"
	// ConfigurationManagerClientStateVInstallFailed undocumented
	ConfigurationManagerClientStateVInstallFailed ConfigurationManagerClientState = "installFailed"
	// ConfigurationManagerClientStateVUpdateFailed undocumented
	ConfigurationManagerClientStateVUpdateFailed ConfigurationManagerClientState = "updateFailed"
	// ConfigurationManagerClientStateVCommunicationError undocumented
	ConfigurationManagerClientStateVCommunicationError ConfigurationManagerClientState = "communicationError"
)

var (
	// ConfigurationManagerClientStatePUnknown is a pointer to ConfigurationManagerClientStateVUnknown
	ConfigurationManagerClientStatePUnknown = &_ConfigurationManagerClientStatePUnknown
	// ConfigurationManagerClientStatePInstalled is a pointer to ConfigurationManagerClientStateVInstalled
	ConfigurationManagerClientStatePInstalled = &_ConfigurationManagerClientStatePInstalled
	// ConfigurationManagerClientStatePHealthy is a pointer to ConfigurationManagerClientStateVHealthy
	ConfigurationManagerClientStatePHealthy = &_ConfigurationManagerClientStatePHealthy
	// ConfigurationManagerClientStatePInstallFailed is a pointer to ConfigurationManagerClientStateVInstallFailed
	ConfigurationManagerClientStatePInstallFailed = &_ConfigurationManagerClientStatePInstallFailed
	// ConfigurationManagerClientStatePUpdateFailed is a pointer to ConfigurationManagerClientStateVUpdateFailed
	ConfigurationManagerClientStatePUpdateFailed = &_ConfigurationManagerClientStatePUpdateFailed
	// ConfigurationManagerClientStatePCommunicationError is a pointer to ConfigurationManagerClientStateVCommunicationError
	ConfigurationManagerClientStatePCommunicationError = &_ConfigurationManagerClientStatePCommunicationError
)

var (
	_ConfigurationManagerClientStatePUnknown            = ConfigurationManagerClientStateVUnknown
	_ConfigurationManagerClientStatePInstalled          = ConfigurationManagerClientStateVInstalled
	_ConfigurationManagerClientStatePHealthy            = ConfigurationManagerClientStateVHealthy
	_ConfigurationManagerClientStatePInstallFailed      = ConfigurationManagerClientStateVInstallFailed
	_ConfigurationManagerClientStatePUpdateFailed       = ConfigurationManagerClientStateVUpdateFailed
	_ConfigurationManagerClientStatePCommunicationError = ConfigurationManagerClientStateVCommunicationError
)

// ConfigurationUsage undocumented
type ConfigurationUsage string

const (
	// ConfigurationUsageVBlocked undocumented
	ConfigurationUsageVBlocked ConfigurationUsage = "blocked"
	// ConfigurationUsageVRequired undocumented
	ConfigurationUsageVRequired ConfigurationUsage = "required"
	// ConfigurationUsageVAllowed undocumented
	ConfigurationUsageVAllowed ConfigurationUsage = "allowed"
)

var (
	// ConfigurationUsagePBlocked is a pointer to ConfigurationUsageVBlocked
	ConfigurationUsagePBlocked = &_ConfigurationUsagePBlocked
	// ConfigurationUsagePRequired is a pointer to ConfigurationUsageVRequired
	ConfigurationUsagePRequired = &_ConfigurationUsagePRequired
	// ConfigurationUsagePAllowed is a pointer to ConfigurationUsageVAllowed
	ConfigurationUsagePAllowed = &_ConfigurationUsagePAllowed
)

var (
	_ConfigurationUsagePBlocked  = ConfigurationUsageVBlocked
	_ConfigurationUsagePRequired = ConfigurationUsageVRequired
	_ConfigurationUsagePAllowed  = ConfigurationUsageVAllowed
)
