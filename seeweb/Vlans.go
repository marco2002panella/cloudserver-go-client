package seeweb

import "fmt"

type VlanService service

/*===============================================================================================================================================================================================
 * [STRUCTS / DATA MODELS] - Strutture dati
 *==============================================================================================================================================================================================*/

type VLAN struct {
	ID            int     `json:"id"`
	VlanID        int     `json:"vlan_id"`
	Pvid          bool    `json:"pvid"`
	Status        string  `json:"status"`
	Active        bool    `json:"active"`
	CreatedAt     string  `json:"created_at"`
	DeactivatedAt *string `json:"deactivated_at"` // Puntatore per gestire i valori null
	MacAddress    string  `json:"mac_address"`
	Network       string  `json:"network"`
	Server        string  `json:"server"`
	User          *string `json:"user"`
}

/*===============================================================================================================================================================================================
 * [CREATE SINGLE/ INSERT SINGLE] - Inserisce una nuova risorsa nel sistema SINGLE
 *==============================================================================================================================================================================================*/

type SeewebVLANRequest struct {
	VlanID int  `json:"vlan_id"` // L'id della vlan
	Pvid   bool `json:"pvid"`    // Se la vlan è pvid dell'interfaccia su cui è ospite
}

func (s *VlanService) Create(servername string, networkname string, req *SeewebVLANRequest) (*SeewebActionGetResponse, *Response, error) {
	u := fmt.Sprintf("servers/%s/networks/%s/vlans", servername, networkname)
	v := new(SeewebActionGetResponse)
	resp, err := s.client.newRequestDo("POST", u, req, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [CREATE RANGE/ INSERT RANGE] - Inserisce una nuova risorsa nel sistema RANGE
 *==============================================================================================================================================================================================*/

type SeewebVLANRangeRequest struct {
	VLANS string `json:"vlans"`
}

func (s *VlanService) CreateRange(servername string, networkname string, req *SeewebVLANRequest) (*SeewebActionGetResponse, *Response, error) {
	u := fmt.Sprintf("servers/%s/networks/%s/vlans", servername, networkname)
	v := new(SeewebActionGetResponse)
	resp, err := s.client.newRequestDo("POST", u, req, v)
	if err != nil {
		return nil, nil, err
	}
	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [LIST / READ ALL] - Recupera l'elenco completo delle risorse
 *==============================================================================================================================================================================================*/

type VLANResponse struct {
	Status string `json:"status"`
	Vlans  []VLAN `json:"vlans"`
}

func (s *VlanService) List(servername string, networkname string) (*VLANResponse, *Response, error) {
	u := fmt.Sprintf("/ecs/v2/servers/%s/networks/%s/vlans", servername, networkname)
	v := new(VLANResponse)

	resp, err := s.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [DELETE SINGLE / REMOVE SINGLE] - Rimuove una risorsa dal sistema
 *==============================================================================================================================================================================================*/
type VLANDeleteResponse struct {
	Status string  `json:"status"`
	Action *Action `json:"action"`
}

func (s *VlanService) Delete(servername string, networkname string, vlanID int) (*VLANDeleteResponse, *Response, error) {
	u := fmt.Sprintf("/ecs/v2/servers/%s/networks/%s/vlans/%d", servername, networkname, vlanID)
	v := new(VLANDeleteResponse)

	resp, err := s.client.newRequestDo("DELETE", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [DELETE RANGE / REMOVE RANGE] - Rimuove una risorsa dal sistema
 *==============================================================================================================================================================================================*/

func (s *VlanService) DeleteRange(servername string, networkname string, vlanStartID, vlanEndID int) (*VLANDeleteResponse, *Response, error) {
	u := fmt.Sprintf("/ecs/v2/servers/%s/networks/%s/vlans/%d-%d", servername, networkname, vlanStartID, vlanEndID)
	v := new(VLANDeleteResponse)

	resp, err := s.client.newRequestDo("DELETE", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [UPDATE / EDIT] - Modifica e aggiorna una risorsa esistente
 *==============================================================================================================================================================================================*/

type VLANUpdateRequest struct {
	Pvid bool `json:"pvid"`
}

type SingleVLANResponse struct {
	Status string  `json:"status"`
	Vlan   VLAN    `json:"vlan"`
	Action *Action `json:"action"`
}

func (s *VlanService) Update(servername string, networkname string, vlanID int, req *VLANUpdateRequest) (*SingleVLANResponse, *Response, error) {
	u := fmt.Sprintf("/ecs/v2/servers/%s/networks/%s/vlans/%d", servername, networkname, vlanID)
	v := new(SingleVLANResponse)

	resp, err := s.client.newRequestDo("PATCH", u, req, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
