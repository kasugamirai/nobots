// Code generated from Pkl module `nobot.AppConfig`. DO NOT EDIT.
package appconfig

type Nostr interface {
	GetPublicKeys() string

	GetEventID() string
}

var _ Nostr = (*NostrImpl)(nil)

type NostrImpl struct {
	PublicKeys string `pkl:"publicKeys"`

	EventID string `pkl:"eventID"`
}

func (rcv *NostrImpl) GetPublicKeys() string {
	return rcv.PublicKeys
}

func (rcv *NostrImpl) GetEventID() string {
	return rcv.EventID
}
