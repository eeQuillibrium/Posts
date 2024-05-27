package storage

import (
	"errors"
	"strconv"
	"strings"
)

const (
	eventSeparator   = " "
	eventIsUserExist = "a"
	eventIsPostExist = "b"
)

type Mediator interface {
	Notify(event string) error
}

type storageMediator struct {
	cs *commentsStorage
	ps *postsStorage
	us *usersStorage
}

func NewStorageMediator(
	cs *commentsStorage,
	ps *postsStorage,
	us *usersStorage,
) Mediator {
	mediator := &storageMediator{
		cs: cs,
		ps: ps,
		us: us,
	}

	cs.SetMediator(mediator)
	ps.SetMediator(mediator)
	us.SetMediator(mediator)
	return mediator
}

type eventParts struct {
	Event string
	Value string
}

func (ep *eventParts) InitEvent(eventSeparator string) error {
	var parts []string = strings.Split(eventSeparator, " ")
	if len(parts) != 2 {
		return errors.New("InitEvent(): doesn't have 2 parts in event: event and value")
	}
	ep.Event = parts[0]
	ep.Value = parts[1]
	return nil
}

func (sm *storageMediator) Notify(getEvent string) error {
	parts := &eventParts{}
	if err := parts.InitEvent(getEvent); err != nil {
		return errors.New("storageMediator.Notify():\n" + err.Error())
	}

	switch parts.Event {
	case eventIsUserExist:
		userID, err := strconv.Atoi(parts.Value)
		if err != nil {
			return errors.New("storageMediator.Notify():" + err.Error())
		}
		return sm.us.isUserExist(userID)
	case eventIsPostExist:
		postID, err := strconv.Atoi(parts.Value)
		if err != nil {
			return errors.New("storageMediator.Notify():" + err.Error())
		}
		return sm.ps.isPostExist(postID)
	}
	return nil
}
