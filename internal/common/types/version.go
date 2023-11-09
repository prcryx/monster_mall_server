package types

type IAPPVersion interface {
	ToString() string
}

type AppVersion string

var _ IAPPVersion = (*AppVersion)(nil)

func NewAppVersion(version string) AppVersion {
	return AppVersion(version)
}

func (av AppVersion) ToString() string {
	return string(av)
}
