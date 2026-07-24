package seeweb

import (
	"fmt"
	"strconv"
)

type ScriptService service

/*===============================================================================================================================================================================================
 * [STRUCTS / DATA MODELS] - Strutture dati
 *==============================================================================================================================================================================================*/

type Script struct {
	ID       int     `json:"id"`
	User     string  `json:"user"`
	Title    string  `json:"title"`
	Content  string  `json:"content"`
	Windows  bool    `json:"windows"`
	Public   bool    `json:"public"`
	Category *string `json:"category"`
}

/*===============================================================================================================================================================================================
 * [CREATE / INSERT] - Inserisce una nuova risorsa nel sistema
 *==============================================================================================================================================================================================*/

type ScriptCreateRequest struct {
	Title   string `json:"title"`
	Content string `json:"content"`
	Windows bool   `json:"windows"`
}

type SeewebScriptCreateResponse struct {
	ID       int     `json:"id"`
	User     string  `json:"user"`
	Title    string  `json:"title"`
	Content  string  `json:"content"`
	Windows  bool    `json:"windows"`
	Public   bool    `json:"public"`
	Category *string `json:"category"`
}

func (s *ScriptService) Create(req *ScriptCreateRequest) (*SeewebScriptCreateResponse, *Response, error) {
	u := "/scripts"
	v := new(SeewebScriptCreateResponse)

	resp, err := s.client.newRequestDo("POST", u, req, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [DELETE / REMOVE] - Rimuove una risorsa dal sistema
 *==============================================================================================================================================================================================*/

func (s *ScriptService) Delete(scriptID int) (*Response, error) {
	u := fmt.Sprintf("/scripts/%s", strconv.Itoa(scriptID))

	resp, err := s.client.newRequestDo("DELETE", u, nil, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

/*===============================================================================================================================================================================================
 * [GET / FIND BY ID] - Recupera una singola risorsa tramite identificativo univoco
 *==============================================================================================================================================================================================*/

type SeewebScriptResponse struct {
	Status string  `json:"status"`
	Script *Script `json:"script"`
}

func (s *ScriptService) Get(id int) (*SeewebScriptResponse, *Response, error) {

	u := fmt.Sprintf("/scripts/%s", strconv.Itoa(id))
	v := new(SeewebScriptResponse)

	resp, err := s.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil

}

/*===============================================================================================================================================================================================
 * [LIST / READ ALL] - Recupera l'elenco completo delle risorse
 *==============================================================================================================================================================================================*/

type SeewebScriptListResponse struct {
	Status  string    `json:"status"`
	Scripts []*Script `json:"scripts"`
}

func (s *ScriptService) List() (*SeewebScriptListResponse, *Response, error) {
	u := "/scripts"
	v := new(SeewebScriptListResponse)

	resp, err := s.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [UPDATE / EDIT] - Modifica e aggiorna una risorsa esistente
 *==============================================================================================================================================================================================*/

type ScriptUpdateRequest struct {
	Title   *string `json:"title,omitempty"`
	Content *string `json:"content,omitempty"`
	Windows *bool   `json:"windows,omitempty"`
}

func (s *ScriptService) Update(id int, req *ScriptUpdateRequest) (*SeewebScriptResponse, *Response, error) {
	u := fmt.Sprintf("/scripts/%s", strconv.Itoa(id))
	v := new(SeewebScriptResponse)

	resp, err := s.client.newRequestDo("PUT", u, req, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
