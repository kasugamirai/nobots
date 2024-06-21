// Code generated from Pkl module `nobot.AppConfig`. DO NOT EDIT.
package appconfig

type Twitter interface {
	GetUsers() []string

	GetVipUsers() []string
}

var _ Twitter = (*TwitterImpl)(nil)

type TwitterImpl struct {
	Users []string `pkl:"users"`

	VipUsers []string `pkl:"vipUsers"`
}

func (rcv *TwitterImpl) GetUsers() []string {
	return rcv.Users
}

func (rcv *TwitterImpl) GetVipUsers() []string {
	return rcv.VipUsers
}
