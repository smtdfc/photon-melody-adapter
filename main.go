package photonMelodyAdapter

import (
	"github.com/google/uuid"
	"github.com/olahol/melody"
	"github.com/smtdfc/photon"
	"encoding/json"
	"log"
	"net/http"
	"errors"
)

type MelodyAdapter struct {
	Instance *melody.Melody
	handlers map[string]photon.SocketEventHandler
	name     string
	Clients  map[string]*photon.SocketSession
}

func Init() (*MelodyAdapter, *melody.Melody) {
	m := melody.New()
	return &MelodyAdapter{
		Instance: m,
		handlers: make(map[string]photon.SocketEventHandler),
		name:     "PhotonMelodyAdapter",
	}, m
}

func (m *MelodyAdapter) GetName() string {
	return m.name
}

func (m *MelodyAdapter) Init() error {
	m.Instance.HandleConnect(func(s *melody.Session) {
		clientID := uuid.New().String()
		m.Clients[clientID] = &photon.SocketSession{
			ClientID: clientID,
			Data:     map[string]any{},
			Instance: s,
		}
	})

	m.Instance.HandleDisconnect(func(s *melody.Session) {
		for id, client := range m.Clients {
			if client.Instance == s {
				delete(m.Clients, id)
				break
			}
		}
	})

	m.Instance.HandleMessage(func(s *melody.Session, msg []byte) {
		var event photon.SocketEventMessage
		if err := json.Unmarshal(msg, &event); err != nil {
			log.Println("Invalid message format:", err)
			return
		}

		var sender *photon.SocketSession
		for _, client := range m.Clients {
			if client.Instance == s {
				sender = client
				break
			}
		}

		if sender == nil {
			log.Println("Unknown session, ignoring message")
			return
		}

		if handler, ok := m.handlers[event.Event]; ok && handler != nil {
			handler(sender, &event)
		} else {
			log.Println("No handler for event:", event.Event)
		}
	})

	return nil
}

func (m *MelodyAdapter) Start() error {
	log.Println("Photon Melody Adapter started !")
	return nil
}

func (m *MelodyAdapter) Listen(port string) error {
	return nil
}

func (m *MelodyAdapter) On(event string, handler photon.SocketEventHandler) {
	m.handlers[event] = handler
}

func (m *MelodyAdapter) Emit(client *photon.SocketSession, msg *photon.SocketEventMessage) error {
	inst, ok := client.Instance.(*melody.Session)
	if !ok {
		return errors.New("invalid session instance")
	}
	
	payload, err := json.Marshal(msg)
if err != nil {
	log.Println("Marshal error:", err)
	return err
}


	return inst.Write(payload)

}

func (m *MelodyAdapter) Stop() error {
	return m.Instance.Close()
}

func (m *MelodyAdapter) HTTPHandler() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		m.Instance.HandleRequest(w, r)
	}
}
