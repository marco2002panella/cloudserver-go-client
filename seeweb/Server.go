package seeweb

import (
	"fmt"
	"time"
)

type ServerService service

/*===============================================================================================================================================================================================
 * [STRUCTS / DATA MODELS] - Strutture dati
 *==============================================================================================================================================================================================*/

type PlanSize struct {
	Core      string `json:"core"`
	RAM       string `json:"ram"`
	Disk      string `json:"disk"`
	GPU       string `json:"gpu"`
	Gpu_label string `json:"gpu_label"`
	Host_type string `json:"host_type"`
}

type ServerDetails struct {
	Name          string        `json:"name"`
	Ipv4          string        `json:"ipv4"`
	Ipv6          string        `json:"ipv6"`
	Group         *Group        `json:"group"`
	Plan          string        `json:"plan"`
	PlanSize      *PlanSize     `json:"plan_size"`
	ReservedPlans []interface{} `json:"reserved_plans"`
	IsReserved    bool          `json:"is_reserved"`
	ReservedUntil string        `json:"reserved_until"`
	Support       *string       `json:"support"`
	Location      string        `json:"location"`
	LocationLabel string        `json:"location_label"`
	Notes         string        `json:"notes"`
	So            string        `json:"so"`
	SoLabel       string        `json:"so_label"`
	CreationDate  string        `json:"creation_date"` // Può essere convertito in time.Time se usi un parser personalizzato
	DeletionDate  *string       `json:"deletion_date"`
	ActiveFlag    bool          `json:"active_flag"`
	Status        string        `json:"status"`
	Progress      int           `json:"progress"`
	ApiVersion    string        `json:"api_version"`
	User          string        `json:"user"`
}

type Server struct {
	Name           string    `json:"name"`
	Ipv4           string    `json:"ipv4"`
	Ipv6           string    `json:"ipv6"`
	Plan           string    `json:"plan"`
	PlanSize       *PlanSize `json:"plan_size"`
	Location       string    `json:"location"`
	Location_label string    `json:"location_label"`
	Notes          string    `json:"notes"`
	So             string    `json:"so"`
	So_label       string    `so_label`
	support        string    `json:"support"`
	CreationDate   time.Time `json:"creation_date"`
	DeletionDate   time.Time `json:"deletion_date"`
	ActiveFlag     bool      `json:"active_flag"`
	Status         string    `json:"status"`
	APIVersion     string    `json:"api_version"`
	User           string    `json:"user"`
	Progress       int       `json:"progress"`
	Group          *Group    `json:"group"`
	SSHKey         string    `json:"ssh_key,omitempty"`
}

/*===============================================================================================================================================================================================
 * [CREATE / INSERT] - Inserisce una nuova risorsa nel sistema
 *==============================================================================================================================================================================================*/

type SeewebServerCreateRequest struct {
	Plan        string   `json:"plan"`
	Location    string   `json:"location"`
	Image       string   `json:"image"`
	Notes       string   `json:"notes"`
	SSHKey      string   `json:"ssh_key,omitempty"`
	Group       string   `json:"group,omitempty"`
	IsolateFrom []string `json:"isolate_from,omitempty"`
	Spot        bool     `json:"spot,omitempty"`
}

type SeewebServerCreateResponse struct {
	Status   string  `json:"status"`
	ActionID int     `json:"action_id"`
	Server   *Server `json:"server"`
}

func (s *ServerService) Create(createServerRequest *SeewebServerCreateRequest) (*SeewebServerCreateResponse, *Response, error) {
	u := "/servers"
	v := new(SeewebServerCreateResponse)

	resp, err := s.client.newRequestDo("POST", u, &createServerRequest, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [DELETE / REMOVE] - Rimuove una risorsa dal sistema
 *==============================================================================================================================================================================================*/

type SeewebServerDeleteResponse struct {
	Status string  `json:"status"`
	Action *Action `json:"action"`
}

func (s *ServerService) Delete(name string) (*SeewebServerDeleteResponse, *Response, error) {
	u := fmt.Sprintf("/servers/%s", name)
	v := new(SeewebServerDeleteResponse)

	resp, err := s.client.newRequestDo("DELETE", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [LIST / READ ALL] - Recupera l'elenco completo delle risorse
 *==============================================================================================================================================================================================*/

type SeewebServerListResponse struct {
	Status string    `json:"status"`
	Count  int       `json:"count"`
	Server []*Server `json:"server"`
}

func (s *ServerService) List() (*SeewebServerListResponse, *Response, error) {
	u := "/servers"
	v := new(SeewebServerListResponse)

	resp, err := s.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [UPDATE / EDIT] - Modifica e aggiorna una risorsa esistente
 *==============================================================================================================================================================================================*/

type SeewebServerUpdateRequest struct {
	Note  string `json:"note,omitempty"`
	Group string `json:"group,omitempty"`
}

type SeewebServerUpdateResponse struct {
	Status string `json:"status"`
}

func (s *ServerService) Update(name string, updateServerRequest *SeewebServerUpdateRequest) (*SeewebServerUpdateResponse, *Response, error) {
	u := fmt.Sprintf("/servers/%s", name)
	v := new(SeewebServerUpdateResponse)

	resp, err := s.client.newRequestDo("PUT", u, &updateServerRequest, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [GET / FIND BY ID] - Recupera una singola risorsa tramite identificativo univoco
 *==============================================================================================================================================================================================*/

type SeewebServerResponse struct {
	Status string         `json:"status"`
	Server *ServerDetails `json:"server"`
}

func (s *ServerService) Get(name string) (*SeewebServerResponse, *Response, error) {
	u := fmt.Sprintf("/servers/%s", name)
	v := new(SeewebServerResponse)

	resp, err := s.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
