package seeweb

import (
	"fmt"
)

type SnapshotService service

/*===============================================================================================================================================================================================
 * [STRUCTS / DATA MODELS] - Strutture dati
 *==============================================================================================================================================================================================*/

type Snapshot struct {
	ID                 int     `json:"id"`
	Name               string  `json:"name"`
	User               string  `json:"user"`
	SnapshotParent     *int    `json:"snapshot_parent"`
	SnapshotParentName *string `json:"snapshot_parent_name"`
	IsLastRestored     bool    `json:"is_last_restored"`
	Protected          bool    `json:"protected"`
	Restoring          bool    `json:"restoring"`
	SourceServer       string  `json:"source_server"`
	Status             string  `json:"status"`
	StatusLabel        string  `json:"status_label"`
	UID                string  `json:"uid"`
	Description        string  `json:"description"`
	Notes              string  `json:"notes"`
	ActiveFlag         bool    `json:"active_flag"`
	SizeOnDisk         *int    `json:"size_on_disk"`
	CreatedAt          string  `json:"created_at"`
	UpdatedAt          string  `json:"updated_at"`
	DeletedAt          *string `json:"deleted_at"`
	ApiVersionValue    int     `json:"api_version_value"`
	ApiVersion         string  `json:"api_version"`
}

/*===============================================================================================================================================================================================
 * [GET / FIND BY ID] - Recupera una singola risorsa tramite identificativo univoco
 *==============================================================================================================================================================================================*/

type SnapshotCreateRequest struct {
	Description string `json:"description,omitempty"`
	Notes       string `json:"notes,omitempty"`
}

type SeewebSnapshotCreateResponse struct {
	Status   string    `json:"status"`
	Snapshot *Snapshot `json:"snapshot"`
	Action   *Action   `json:"action"`
}

func (s *SnapshotService) Create(servername string, req *SnapshotCreateRequest) (*SeewebSnapshotCreateResponse, *Response, error) {
	u := fmt.Sprintf("/servers/%s/snapshots", servername)
	v := new(SeewebSnapshotCreateResponse)

	resp, err := s.client.newRequestDo("POST", u, req, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [GET / FIND BY ID] - Recupera una singola risorsa tramite identificativo univoco
 *==============================================================================================================================================================================================*/

type SeewebSnapshotResponse struct {
	Status   string    `json:"status"`
	Snapshot *Snapshot `json:"snapshot"`
}

func (s *SnapshotService) Get(servername string, snapshotID string) (*SeewebSnapshotResponse, *Response, error) {
	u := fmt.Sprintf("/servers/%s/snapshots/%s", servername, snapshotID)
	v := new(SeewebSnapshotResponse)

	resp, err := s.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [DELETE / REMOVE] - Rimuove una risorsa dal sistema
 *==============================================================================================================================================================================================*/

type SeewebSnapshotDeleteResponse struct {
	Status   string    `json:"status"`
	Snapshot *Snapshot `json:"snapshot"`
	Action   *Action   `json:"action"`
}

func (s *SnapshotService) Delete(servername string, snapshotID string) (*SeewebSnapshotDeleteResponse, *Response, error) {
	u := fmt.Sprintf("/servers/%s/snapshots/%s", servername, snapshotID)
	v := new(SeewebSnapshotDeleteResponse)

	resp, err := s.client.newRequestDo("DELETE", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [LIST / READ ALL] - Recupera l'elenco completo delle risorse
 *==============================================================================================================================================================================================*/

type SeewebSnapshotListResponse struct {
	Status    string      `json:"status"`
	Snapshots []*Snapshot `json:"snapshots"`
}

func (s *SnapshotService) List(servername string) (*SeewebSnapshotListResponse, *Response, error) {
	u := fmt.Sprintf("/servers/%s/snapshots", servername)
	v := new(SeewebSnapshotListResponse)

	resp, err := s.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [UPDATE / EDIT] - Modifica e aggiorna una risorsa esistente
 *==============================================================================================================================================================================================*/

type SnapshotUpdateRequest struct {
	Description *string `json:"description,omitempty"`
	Notes       *string `json:"notes,omitempty"`
	Protected   *bool   `json:"protected,omitempty"`
}

func (s *SnapshotService) Update(servername string, snapshotID string, req *SnapshotUpdateRequest) (*SeewebSnapshotResponse, *Response, error) {
	u := fmt.Sprintf("/servers/%s/snapshots/%s", servername, snapshotID)
	v := new(SeewebSnapshotResponse)

	resp, err := s.client.newRequestDo("PATCH", u, req, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
