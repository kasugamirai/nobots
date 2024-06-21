// Code generated from Pkl module `nobot.AppConfig`. DO NOT EDIT.
package appconfig

type Report interface {
	GetPublicKeys() []string

	GetUsernames() []string
}

var _ Report = (*ReportImpl)(nil)

type ReportImpl struct {
	PublicKeys []string `pkl:"publicKeys"`

	Usernames []string `pkl:"usernames"`
}

func (rcv *ReportImpl) GetPublicKeys() []string {
	return rcv.PublicKeys
}

func (rcv *ReportImpl) GetUsernames() []string {
	return rcv.Usernames
}
