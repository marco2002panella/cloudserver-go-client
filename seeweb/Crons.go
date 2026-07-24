package seeweb

import "fmt"

type CronService service

/*===============================================================================================================================================================================================
 * [CREATE / INSERT] - Inserisce una nuova risorsa nel sistema
 *==============================================================================================================================================================================================*/

type CronConfig struct {
	Typology     string `json:"typology"`
	MaxRetention int    `json:"max_retention"`
	DayOfWeek    int    `json:"day_of_week"`
}

type CronInfos struct {
	MaxRetention int `json:"max_retention"`
	DayOfWeek    int `json:"day_of_week"`
}

type Cron struct {
	Name          string     `json:"name"`
	Schedule      string     `json:"schedule"`
	Enabled       bool       `json:"enabled"`
	LastRunAt     *string    `json:"last_run_at"`
	Infos         *CronInfos `json:"infos"`
	LastRunStatus *string    `json:"last_run_status"`
}

/*===============================================================================================================================================================================================
 * [CREATE / INSERT] - Inserisce una nuova risorsa nel sistema
 *==============================================================================================================================================================================================*/

type CronCreateRequest struct {
	Type   string      `json:"type"`
	Config *CronConfig `json:"config"`
}

type SeewebCronCreateResponse struct {
	Status string  `json:"status"`
	Cron   *Cron   `json:"cron"`
	Action *Action `json:"action"`
}

func (c *CronService) Create(servername string, req *CronCreateRequest) (*SeewebCronCreateResponse, *Response, error) {
	u := fmt.Sprintf("/servers/%s/crons", servername)
	v := new(SeewebCronCreateResponse)
	
	resp, err := c.client.newRequestDo("POST", u, req, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [DELETE / REMOVE] - Rimuove una risorsa dal sistema
 *==============================================================================================================================================================================================*/

type SeewebCronDeleteResponse struct {
	Status string `json:"status"`
	Cron   *Cron  `json:"cron"`
}

func (c *CronService) Delete(servername string) (*SeewebCronDeleteResponse, *Response, error) {
	cronName := fmt.Sprintf("snapshot_schedule_%s", servername)
	u := fmt.Sprintf("/servers/%s/crons/%s", servername, cronName)
	v := new(SeewebCronDeleteResponse)

	resp, err := c.client.newRequestDo("DELETE", u, nil, &v)
	if err != nil {
		// Se il server restituisce un corpo vuoto ma lo status HTTP è di successo (es. 200 o 204),
		// possiamo ignorare l'errore di decodifica JSON.
		if resp != nil && resp.Response != nil && (resp.Response.StatusCode == 200 || resp.Response.StatusCode == 204) {
			return &SeewebCronDeleteResponse{Status: "ok"}, resp, nil
		}
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [GET / FIND BY ID] - Recupera una singola risorsa tramite identificativo univoco
 *==============================================================================================================================================================================================*/

type SeewebCronGetResponse struct {
	Status string `json:"status"`
	Cron   *Cron  `json:"cron"`
}

func (c *CronService) Get(servername string) (*SeewebCronGetResponse, *Response, error) {
	cronName := fmt.Sprintf("snapshot_schedule_%s", servername)
	u := fmt.Sprintf("/servers/%s/crons/%s", servername, cronName)
	v := new(SeewebCronGetResponse)

	resp, err := c.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [UPDATE / EDIT] - Modifica e aggiorna una risorsa esistente
 *==============================================================================================================================================================================================*/

type CronUpdateRequest struct {
	Enabled      *bool `json:"enabled,omitempty"`
	MaxRetention *int  `json:"max_retention,omitempty"`
	DayOfWeek    *int  `json:"day_of_week,omitempty"`
}

func (c *CronService) Update(servername string, req *CronUpdateRequest) (*SeewebCronGetResponse, *Response, error) {
	cronName := fmt.Sprintf("snapshot_schedule_%s", servername)
	u := fmt.Sprintf("/servers/%s/crons/%s", servername, cronName)
	v := new(SeewebCronGetResponse)

	resp, err := c.client.newRequestDo("PATCH", u, req, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
