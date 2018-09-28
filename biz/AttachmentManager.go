package biz

import (
	"FFQATracking/models"
	"FFQATracking/utils"
	"errors"
	"log"
	"sync"
)

///////////////////////////////////////////////////////

// AttachmentSession session data type
type AttachmentSession struct {
	SessionID      int64
	AllAttachments []*models.AttachmentModel
}

// AttachmentWithIndex get attachment witn index
func (c *AttachmentSession) AttachmentWithIndex(idx int) (*models.AttachmentModel, error) {
	if idx < 0 || idx >= len(c.AllAttachments) {
		return nil, errors.New("out of range")
	}
	return c.AllAttachments[idx], nil
}

// AppendAttachement append a new attachment data
func (c *AttachmentSession) AppendAttachement(attachFN string) {
	pAttach := &models.AttachmentModel{
		Tm:       utils.TimeTickSince1970(),
		FileName: attachFN,
	}
	c.AllAttachments = append(c.AllAttachments, pAttach)
}

// RemoveAttachmentWithIndex remove element in array with idx
func (c *AttachmentSession) RemoveAttachmentWithIndex(idx int) error {

	if idx < 0 || idx >= len(c.AllAttachments) {
		return errors.New("out of range")
	}

	c.AllAttachments = append(c.AllAttachments[:idx], c.AllAttachments[idx+1:]...)
	return nil
}

///////////////////////////////////////////////////////

// AttachmentManager controller attachment
type AttachmentManager struct {
	Sessions map[int64]*AttachmentSession
}

var gAttachmentManagerInstance *AttachmentManager
var gAttachmentOnce sync.Once

// SharedAttachManager get attachment singleton object
func SharedAttachManager() *AttachmentManager {
	gAttachmentOnce.Do(func() {
		gAttachmentManagerInstance = &AttachmentManager{}
		gAttachmentManagerInstance.Sessions = make(map[int64]*AttachmentSession)
	})
	return gAttachmentManagerInstance
}

// SessionWithID fetch attachement with its id
func (c *AttachmentManager) SessionWithID(id int64) *AttachmentSession {

	if c != SharedAttachManager() {
		log.Fatal("Fatal: caller is not using singleton method.")
	}

	for key, value := range c.Sessions {
		if key == id {
			return value
		}
	}
	return nil
}

// NewSession generate new session for save attachment data temporary
func (c *AttachmentManager) NewSession() int64 {

	if c != SharedAttachManager() {
		log.Fatal("Fatal: caller is not using singleton method.")
	}
	newASKey := utils.TimeTickSince1970()
	newAS := &AttachmentSession{SessionID: newASKey}
	c.Sessions[newASKey] = newAS

	return newASKey
}

// RemoveSession remove session data with its session key
func (c *AttachmentManager) RemoveSession(sessionKey int64) error {

	if c != SharedAttachManager() {
		log.Fatal("Fatal: caller is not using singleton method.")
	}

	for key := range c.Sessions {
		if key == sessionKey {
			delete(c.Sessions, sessionKey)
			return nil
		}
	}
	return errors.New("Fails to delete attachment with section with id")
}
