// Code generated from Pkl module `nobot.AppConfig`. DO NOT EDIT.
package appconfig

type Relays interface {
	GetUrls() []string
}

var _ Relays = (*RelaysImpl)(nil)

// The relay URLs to use.
type RelaysImpl struct {
	Urls []string `pkl:"urls"`
}

func (rcv *RelaysImpl) GetUrls() []string {
	return rcv.Urls
}
