package iland

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type client struct {
	username        string
	password        string
	clientID        string
	clientSecret    string
	Token           Token
	tokenExpiration time.Time
	http            *http.Client
}

func NewClient(Username, Password, clientID, clientSecret string) (ConsoleService, error) {
	client := client{
		username:     Username,
		password:     Password,
		clientID:     clientID,
		clientSecret: clientSecret,
		http:         &http.Client{},
	}
	return &client, nil
}

func (c *client) Location() LocationService {
	return &locationService{c}
}

func (c *client) Company() CompanyService {
	return &companyService{c}
}

func (c *client) User() UserService {
	return &userService{c}
}

func (c *client) Org() OrgService {
	return &orgService{c}
}

func (c *client) Catalog() CatalogService {
	return &catalogService{c}
}

func (c *client) VAppTemplate() VAppTemplateService {
	return &vappTemplateService{c}
}

func (c *client) Vdc() VdcService {
	return &vdcService{c}
}

func (c *client) Edge() EdgeService {
	return &edgeService{c}
}

func (c *client) OrgVdcNetwork() OrgVdcNetworkService {
	return &orgVdcNetworkService{c}
}

func (c *client) VApp() VAppService {
	return &vappService{c}

}

func (c *client) VAppNetwork() VAppNetworkService {
	return &vappNetworkService{c}

}

func (c *client) VirtualMachine() VirtualMachineService {
	return &virtualMachineService{c}

}

func (c *client) VCCBackupTenant() VCCBackupTenantService {
	return &vccBackupTenantService{c}

}

func (c *client) Vpg() VpgService {
	return &vpgService{c}

}

func (c *client) Task() TaskService {
	return &taskService{c}

}

func (c *client) GetOperatingSystems() ([]OperatingSystem, error) {
	schema := struct {
		OperatingSystems []OperatingSystem `json:"data"`
	}{}
	err := c.getObject("/v1/constants/operating-systems", &schema)
	return schema.OperatingSystems, err
}

func (c *client) GetLocations() []Location {
	locations := []Location{}
	for _, locationID := range LocationIDs {
		location := Location{
			ID: locationID,
		}
		locations = append(locations, location)
	}
	return locations
}

func (c *client) GetCompanies() ([]Company, error) {
	schema := struct {
		Companies []Company `json:"data"`
	}{}
	err := c.getObject(fmt.Sprintf("/v1/users/%s/companies", c.username), &schema)
	if err != nil {
		return []Company{}, err
	}
	for i, company := range schema.Companies {
		schema.Companies[i] = company
	}
	return schema.Companies, nil
}

func (c *client) GetOrgs() ([]Org, error) {
	schema := struct {
		Orgs []Org `json:"data"`
	}{}
	err := c.getObject(fmt.Sprintf("/v1/users/%s/orgs", c.username), &schema)
	if err != nil {
		return []Org{}, err
	}
	return schema.Orgs, nil
}

type SocketData struct {
	Type string          `json:"type"`
	Data json.RawMessage `json:"data"`
}

func (c *client) StreamEvents(companyID string) (chan Event, error) {
	events := make(chan Event, 100)
	conn, err := c.dialEventStream(companyID)
	if err != nil {
		return nil, err
	}
	go c.streamEvents(companyID, conn, events)
	return events, nil
}

func (c *client) dialEventStream(companyID string) (*websocket.Conn, error) {
	err := c.RefreshTokenIfNecessary()
	if err != nil {
		return nil, err
	}
	conn, _, err := websocket.DefaultDialer.Dial("wss://api.ilandcloud.com/v1/event-websocket", nil)
	if err != nil {
		return nil, err
	}
	_, message, err := conn.ReadMessage()
	if err != nil {
		return nil, err
	}
	auth := fmt.Sprintf("Bearer %s", c.Token.AccessToken)
	if companyID != "" {
		auth = fmt.Sprintf("companyId=%s,Bearer %s", companyID, c.Token.AccessToken)
	}
	if string(message) == "AUTHORIZATION" {
		err = conn.WriteMessage(websocket.TextMessage, []byte(auth))
		if err != nil {
			return nil, err
		}
		return conn, nil
	}
	return nil, errors.New(string(message))
}

func (c *client) streamEvents(companyID string, conn *websocket.Conn, events chan Event) {
	defer conn.Close()
	for {
		_, message, err := conn.ReadMessage()
		if err != nil {
			break
		}
		msg := SocketData{}
		err = json.Unmarshal(message, &msg)
		if err != nil {
			break
		}
		switch msg.Type {
		case "EVENT":
			event := Event{}
			err = json.Unmarshal(msg.Data, &event)
			if err == nil {
				events <- event
			}
		}
	}
	for {
		time.Sleep(time.Second * 5)
		newConn, err := c.dialEventStream(companyID)
		if err != nil {
			continue
		}
		go c.streamEvents(companyID, newConn, events)
		return
	}
}
