package seeweb

import "fmt"

type NetworkService service

/*===============================================================================================================================================================================================
 * [STRUCTS / DATA MODELS] - Strutture dati
 *==============================================================================================================================================================================================*/

type Network struct {
	ID            int     `json:"id"`
	NetworkID     int     `json:"network_id"`
	Name          string  `json:"name"`
	Active        bool    `json:"active"`
	CreatedAt     string  `json:"created_at"`
	DeactivatedAt *string `json:"deactivated_at"`
	User          string  `json:"user"`
}

type SeewebNetworkCreateResponse struct {
	Status  string   `json:"status"`
	Network *Network `json:"network"`
}

/*===============================================================================================================================================================================================
 * [CREATE / INSERT] - Inserisce una nuova risorsa nel sistema
 *==============================================================================================================================================================================================*/
type NetworkCreateRequest struct {
}

func (n *NetworkService) Create(req *NetworkCreateRequest) (*SeewebNetworkCreateResponse, *Response, error) {
	u := "/ecs/v2/networks"
	v := new(SeewebNetworkCreateResponse)

	resp, err := n.client.newRequestDo("POST", u, req, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [GET / FIND BY ID] - Recupera una singola risorsa tramite identificativo univoco
 *==============================================================================================================================================================================================*/

type SeewebNetworkResponse struct {
	Status  string   `json:"status"`
	Network *Network `json:"network"`
}

func (n *NetworkService) Get(name string) (*SeewebNetworkResponse, *Response, error) {
	u := fmt.Sprintf("/ecs/v2/networks/%s", name)
	v := new(SeewebNetworkResponse)

	resp, err := n.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [DELETE / REMOVE] - Rimuove una risorsa dal sistema
 *==============================================================================================================================================================================================*/

type SeewebNetworkDeleteResponse struct {
	Status string `json:"status"`
}

func (n *NetworkService) Delete(name string) (*SeewebNetworkDeleteResponse, *Response, error) {
	u := fmt.Sprintf("/ecs/v2/networks/%s", name)
	v := new(SeewebNetworkDeleteResponse)

	resp, err := n.client.newRequestDo("DELETE", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [LIST / READ ALL] - Recupera l'elenco completo delle risorse
 *==============================================================================================================================================================================================*/

type SeewebNetworkListResponse struct {
	Status   string     `json:"status"`
	Networks []*Network `json:"networks"`
}

func (n *NetworkService) List() (*SeewebNetworkListResponse, *Response, error) {
	u := "/ecs/v2/networks"
	v := new(SeewebNetworkListResponse)

	resp, err := n.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
