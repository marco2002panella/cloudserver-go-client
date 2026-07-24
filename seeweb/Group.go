package seeweb

import (
	"fmt"
	"strconv"
)

type GroupService service

/*===============================================================================================================================================================================================
 * [STRUCTS / DATA MODELS] - Strutture dati
 *==============================================================================================================================================================================================*/

type Group struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Notes   string `json:"notes"`
	Enabled bool   `json:"enabled"`
}

/*===============================================================================================================================================================================================
 * [CREATE / INSERT] - Inserisce una nuova risorsa nel sistema
 *==============================================================================================================================================================================================*/

type SeewebGroupCreateRequest struct {
	Notes    string `json:"notes"`
	Password string `json:"password"` //deprecated
}

type SeewebGroupCreateResponse struct {
	Status string `json:"status"`
	Group  *Group `json:"group"`
}

func (s *GroupService) Create(createGroupRequest *SeewebGroupCreateRequest) (*SeewebGroupCreateResponse, *Response, error) {
	u := "/groups"
	v := new(SeewebGroupCreateResponse)

	resp, err := s.client.newRequestDo("POST", u, &createGroupRequest, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [GET / FIND BY ID] - Recupera una singola risorsa tramite identificativo univoco
 *==============================================================================================================================================================================================*/

type SeewebGroupResponse struct {
	Status string `json:"status"`
	Group  *Group `json:"group"`
}

func (a *GroupService) Get(groupID string) (*SeewebGroupResponse, *Response, error) {
	u := fmt.Sprintf("/groups/%s", groupID)
	v := new(SeewebGroupResponse)

	resp, err := a.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [DELETE / REMOVE] - Rimuove una risorsa dal sistema
 *==============================================================================================================================================================================================*/

type SeewebGroupDeleteResponse struct {
	Status string `json:"status"`
}

func (s *GroupService) Delete(id int) (*SeewebGroupDeleteResponse, *Response, error) {
	u := fmt.Sprintf("/groups/%s", strconv.Itoa(id))
	v := new(SeewebGroupDeleteResponse)

	resp, err := s.client.newRequestDo("DELETE", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

type SeewebGroupListResponse struct {
	Status string   `json:"status"`
	Groups []*Group `json:"groups"`
}

/*===============================================================================================================================================================================================
 * [LIST / READ ALL] - Recupera l'elenco completo delle risorse
 *==============================================================================================================================================================================================*/

func (a *GroupService) List() (*SeewebGroupListResponse, *Response, error) {
	u := "/groups"
	v := new(SeewebGroupListResponse)

	resp, err := a.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
