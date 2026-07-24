package seeweb

type PlanService service

/*===============================================================================================================================================================================================
 * [STRUCTS / DATA MODELS] - Strutture dati
 *==============================================================================================================================================================================================*/

type AvailableRegions struct {
	ID          int    `json:"id"`
	Location    string `json:"location"`
	Description string `json:"description"`
}

type Plan struct {
	ID               int                 `json:"id"`
	Name             string              `json:"name"`
	CPU              string              `json:"cpu"`
	GPU              string              `json:"gpu"`
	GPU_LABEL        string              `json:"gpu_label"`
	RAM              string              `json:"ram"`
	Disk             string              `json:"disk"`
	HourlyPrice      float64             `json:"hourly_price"`
	MontlyPrice      float64             `json:"montly_price"`
	Windows          bool                `json:"windows"`
	Available        bool                `json:"available"`
	AvailableRegions []*AvailableRegions `json:"available_regions"`
}

type PlanAvailable struct {
	ID               int                 `json:"id"`
	Name             string              `json:"name"`
	CPU              string              `json:"cpu"`
	GPU              string              `json:"gpu"`
	GPU_LABEL        string              `json:"gpu_label"`
	RAM              string              `json:"ram"`
	Disk             string              `json:"disk"`
	HourlyPrice      float64             `json:"hourly_price"`
	MontlyPrice      float64             `json:"montly_price"`
	Windows          bool                `json:"windows"`
	HostType         string              `json:"host_type"`
	Available        bool                `json:"available"`
	OsAvailables     []*Image            `json:"os_availables"`
	AvailableRegions []*AvailableRegions `json:"available_regions"`
}

type PlanSpot struct {
	Piano    Plan    `json:"plan"`
	Discount float64 `json:"discount"`
}

/*===============================================================================================================================================================================================
 * [LIST PLANS/ READ ALL PLANS] - Recupera l'elenco completo delle risorse PLANS
 *==============================================================================================================================================================================================*/

type SeewebPlanListResponse struct {
	Status string  `json:"status"`
	Plans  []*Plan `json:"plans"`
}

// List lists all existing plans.
func (a *PlanService) List() (*SeewebPlanListResponse, *Response, error) {
	u := "/plans"
	v := new(SeewebPlanListResponse)

	resp, err := a.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [LIST PLANS AVAILABLE / READ ALL PLANS AVAILABLE] - Recupera l'elenco completo delle risorse PLANS AVAILABLE
 *==============================================================================================================================================================================================*/

type SeewebPlanListAvailableResponse struct {
	Status string           `json:"status"`
	Plans  []*PlanAvailable `json:"plans"`
}

// ListAvailables recupera l'elenco di tutti i piani hardware disponibili con relative immagini e regioni.
func (p *PlanService) ListAvailables() (*SeewebPlanListAvailableResponse, *Response, error) {
	u := "/plans/availables"
	v := new(SeewebPlanListAvailableResponse)

	resp, err := p.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [LIST PLANS-SPOT/ READ ALL PLANS-SPOT] - Recupera l'elenco completo delle risorse PLANS-SPOT
 *==============================================================================================================================================================================================*/

type SeewebPlanListSpotResponse struct {
	Status string      `json:"status"`
	Plans  []*PlanSpot `json:"plans"`
}

func (a *PlanService) ListSpot() (*SeewebPlanListSpotResponse, *Response, error) {
	u := "/plans/spot"
	v := new(SeewebPlanListSpotResponse)

	resp, err := a.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
