package memory

import (
	"errors"

	"github.com/emersion/neutron/backend"
	"github.com/emersion/neutron/backend/util"
)

type Backend struct {
	backend.DomainsBackend
	backend.ContactsBackend
	backend.LabelsBackend
	backend.ConversationsBackend
	backend.SendBackend
	backend.EventsBackend

	users map[string]*user
}

func (b *Backend) Set(item interface{}) error {
	return errors.New("Unsupported backend")
}

type user struct {
	*backend.User
	password string
}

func New() backend.Backend {
	bkd := &Backend{
		users: map[string]*user{},
	}

	bkd.EventsBackend = NewEventsBackend()
	bkd.DomainsBackend = NewDomainsBackend()
	bkd.ContactsBackend = util.NewEventedContactsBackend(NewContactsBackend(), bkd.EventsBackend)
	bkd.LabelsBackend = util.NewEventedLabelsBackend(NewLabelsBackend(), bkd.EventsBackend)
	bkd.ConversationsBackend = util.NewEventedConversationsBackend(NewConversationsBackend(), bkd.EventsBackend)
	bkd.SendBackend = util.NewEchoSendBackend(bkd.ConversationsBackend)

	return bkd
}
