package seeweb

type RegionService service

/*===============================================================================================================================================================================================
 * [STRUCTS / DATA MODELS] - Strutture dati
 *==============================================================================================================================================================================================*/

type Region struct {
	ID          int    `json:"id"`
	Location    string `json:"location"`
	Description string `json:"description"`
}

/*===============================================================================================================================================================================================
 * [LIST REGIONS/ READ ALL REGIONS] - Recupera l'elenco completo delle risorse REGIONS
 *==============================================================================================================================================================================================*/

type SeewebRegionListResponse struct {
	Status  string    `json:"status"`
	Regions []*Region `json:"regions"`
}

func (a *RegionService) List() (*SeewebRegionListResponse, *Response, error) {
	u := "/regions"
	v := new(SeewebRegionListResponse)

	resp, err := a.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [LIST REGIONS AVAILABLE/ READ ALL REGIONS AVAILABLE] - Recupera l'elenco completo delle risorse REGIONS AVAILABLE
 *==============================================================================================================================================================================================*/

type SeewebRegionAvailabilityResponse struct {
	Status  string     `json:"status"`
	Regions [][]string `json:"regions"`
}

type RegionAvailabilityRequest struct {
	Plan string `json:"plan"`
}

func (r *RegionService) CheckAvailables(req *RegionAvailabilityRequest) (*SeewebRegionAvailabilityResponse, *Response, error) {
	u := "/regions/availables"
	v := new(SeewebRegionAvailabilityResponse)

	resp, err := r.client.newRequestDo("POST", u, req, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
