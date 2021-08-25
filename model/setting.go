package setting

const SettingKind = "Setting"

// Model for an application setting
type Setting struct {
	AppId string
	Key   string
	Value string
}
