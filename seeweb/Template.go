package seeweb

import (
	"fmt"
	"time"
)

type TemplateService service

/*===============================================================================================================================================================================================
 * [STRUCTS / DATA MODELS] - Strutture dati
 *==============================================================================================================================================================================================*/
type Template struct {
	ID           int       `json:"id"`
	Name         string    `json:"name"`
	CreationDate time.Time `json:"creation_date"`
	ActiveFlag   bool      `json:"active_flag"`
	Status       string    `json:"status"`
	UUID         string    `json:"uuid"`
	Description  string    `json:"description"`
	Notes        string    `json:"notes"`
	Public       bool      `json:"public"`
	Cloud_image  bool      `json:"cloud_image"`
	So_base      string    `json:"so_base"`
	api_version  string    `json:"api_version"`
	Version      string    `json:"version"`
}

/*===============================================================================================================================================================================================
 * [CREATE / INSERT] - Inserisce una nuova risorsa nel sistema
 *==============================================================================================================================================================================================*/

type SeewebTemplateCreateRequest struct {
	Server      string `json:"server,omitempty"`
	Snapshot    int    `json:"snapshot,omitempty"`
	Description string `json:"description"`
}

type SeewebTemplateGetResponseAct struct {
	Status   string    `json:"status"`
	ActionId int       `json:"action_id"`
	Template *Template `json:"template"`
}

func (a *TemplateService) Create(id string, req *SeewebTemplateCreateRequest) (*SeewebTemplateGetResponseAct, *Response, error) {
	u := "/templates"
	v := new(SeewebTemplateGetResponseAct)

	resp, err := a.client.newRequestDo("POST", u, req, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [DELETE / REMOVE] - Rimuove una risorsa dal sistema
 *==============================================================================================================================================================================================*/

type SeewebTemplateDeleteResponse struct {
	Status string  `json:"status"`
	Action *Action `json:"action"`
}

// Delete deletes a single template
func (a *TemplateService) Delete(name string) (*SeewebTemplateDeleteResponse, *Response, error) {
	u := fmt.Sprintf("/templates/%s", name)
	v := new(SeewebTemplateDeleteResponse)

	resp, err := a.client.newRequestDo("DELETE", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [GET / FIND BY ID] - Recupera una singola risorsa tramite identificativo univoco
 *==============================================================================================================================================================================================*/

type SeewebTemplateGetResponse struct {
	Status   string    `json:"status"`
	Template *Template `json:"template"`
}

func (a *TemplateService) Get(name string) (*SeewebTemplateGetResponse, *Response, error) {
	u := fmt.Sprintf("/templates/%s", name)
	v := new(SeewebTemplateGetResponse)

	resp, err := a.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [LIST / READ ALL] - Recupera l'elenco completo delle risorse
 *==============================================================================================================================================================================================*/
type SeewebTemplateListResponse struct {
	Status    string      `json:"status"`
	Templates []*Template `json:"templates"`
}

func (a *TemplateService) List() (*SeewebTemplateListResponse, *Response, error) {
	u := "/templates"
	v := new(SeewebTemplateListResponse)

	resp, err := a.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [UPDATE / EDIT] - Modifica e aggiorna una risorsa esistente
 *==============================================================================================================================================================================================*/
type SeewebTemplateUpdateRequest struct {
	Description string `json:"description"`
}

func (a *TemplateService) Update(name string, req *SeewebTemplateUpdateRequest) (*SeewebTemplateGetResponse, *Response, error) {
	u := fmt.Sprintf("/templates/%s", name)
	v := new(SeewebTemplateGetResponse)

	resp, err := a.client.newRequestDo("POST", u, req, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
