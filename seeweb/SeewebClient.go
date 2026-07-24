package seeweb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

/*===============================================================================================================================================================================================
 * [STRUCTS / DATA MODELS] - Strutture dati
 *==============================================================================================================================================================================================*/
const (
	defaultBaseURL   = "https://api.seeweb.it/ecs/v2"
	defaultUserAgent = "Seeweb/cloudserver-go-client(terraform)"
)

type service struct {
	client *Client
}

type Config struct {
	BaseURL    string
	HTTPClient *http.Client
	Timeout    int
	XAPITOKEN  string
	JWTtoken   string
	UserAgent  string
	DEBUG      bool
}

type Client struct {
	baseURL      *url.URL
	client       *http.Client
	Config       *Config
	Server       *ServerService
	Action       *ActionService
	Template     *TemplateService
	Group        *GroupService
	Region       *RegionService
	Plan         *PlanService
	Network      *NetworkService
	CloudScripts *ScriptService
	Image        *ImageService
	Vlan         *VlanService
	Snapshot     *SnapshotService
	Cron         *CronService
}

type Response struct {
	Response  *http.Response
	BodyBytes []byte
}

type RequestOptions struct {
	Type  string
	Label string
	Value string
}

/*===============================================================================================================================================================================================
 * [CREATE CLIENT / NEW CLIENT] - Creazione di un nuovo client
 *==============================================================================================================================================================================================*/

func NewClient(config *Config) (*Client, error) {

	durata := 30 * time.Second
	if config.Timeout > 0 {
		durata = time.Duration(config.Timeout) * time.Second
	}

	if config.HTTPClient == nil {
		config.HTTPClient = &http.Client{
			Timeout: durata,
		}
	}

	if config.BaseURL == "" {
		config.BaseURL = defaultBaseURL
	}

	baseURL, err := url.Parse(config.BaseURL)
	if err != nil {
		return nil, err
	}

	c := &Client{
		baseURL: baseURL,
		client:  config.HTTPClient,
		Config:  config,
	}

	c.Server = &ServerService{c}
	c.Action = &ActionService{c}
	c.Template = &TemplateService{c}
	c.Group = &GroupService{c}
	c.Region = &RegionService{c}
	c.Plan = &PlanService{c}
	c.Network = &NetworkService{c}
	c.Image = &ImageService{c}
	c.CloudScripts = &ScriptService{c}
	c.Snapshot = &SnapshotService{c}
	c.Cron = &CronService{c}
	return c, nil
}

/*===============================================================================================================================================================================================
 * [CLIENT REQUEST ] - da qui in poi il client usa sempre la funzione newRequestDo per fare le richieste agli endpoint, porre DEBUG=true per attivare il debug nei file di log false altrimenti.
 *==============================================================================================================================================================================================*/

func (c *Client) newRequestDo(method, url string, body, v interface{}) (*Response, error) {
	req, err := c.newRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	return c.do(req, v)
}

/*===============================================================================================================================================================================================
 * [UTILITY FUNCTIONS/ IGNORE ] - funzioni che si possono ignorare, aiutano soltanto newRequestDO a fare il proprio lavoro.
 *==============================================================================================================================================================================================*/

func (c *Client) newRequest(method, url string, body interface{}, options ...RequestOptions) (*http.Request, error) {
	var buf io.ReadWriter
	var bodyString string
	if body != nil {
		buf = new(bytes.Buffer)
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
		bodyString = buf.(*bytes.Buffer).String()
	}

	u := c.baseURL.String() + url

	req, err := http.NewRequest(method, u, buf)
	if err != nil {
		return nil, err
	}

	if len(options) > 0 {
		for _, o := range options {
			if o.Type == "header" {
				req.Header.Add(o.Label, o.Value)
			}
		}
	}

	if c.Config.XAPITOKEN != "" {
		req.Header.Add("X-APITOKEN", c.Config.XAPITOKEN)
	} else if c.Config.JWTtoken != "" {
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.Config.JWTtoken))
	}

	req.Header.Add("Content-Type", "application/json")

	if c.Config.UserAgent != "" {
		req.Header.Add("User-Agent", c.Config.UserAgent)
	} else {
		req.Header.Add("User-Agent", defaultUserAgent)
	}

	if c.Config.DEBUG {
		// Scrive la richiesta nel file log_REQUESTES_log.txt
		reqLog := fmt.Sprintf("[%s] METHOD: %s | URL: %s\nBODY: %s",
			time.Now().Format("2006-01-02 15:04:05"), method, u, bodyString)
		WriteLog("log_REQUESTES_log.txt", reqLog)

		if c.Config.DEBUG {
			log.Printf("[DEBUG] Seeweb - Preparing %s request to %s", method, url)
		}
	}

	return req, nil
}

func (c *Client) do(req *http.Request, v interface{}) (*Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}
	bodyBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// Ripristina il body nel caso serva ad altri lettori
	resp.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	response := &Response{
		Response:  resp,
		BodyBytes: bodyBytes,
	}

	var prettyBody bytes.Buffer
	if err := json.Indent(&prettyBody, bodyBytes, "", "    "); err == nil {
		bodyBytes = prettyBody.Bytes()
	}
	if c.Config.DEBUG {
		// Scrive la risposta nel file log_RESPONSES_log.txt
		respLog := fmt.Sprintf("[%s] URL: %s | STATUS: %s\nBODY:\n%s",
			time.Now().Format("2006-01-02 15:04:05"), req.URL.String(), resp.Status, string(bodyBytes))

		WriteLog("log_RESPONSES_log.txt", respLog)
	}

	if err := c.checkResponse(response); err != nil {
		return response, err
	}

	if v != nil {
		if err := c.DecodeJSON(response, v); err != nil {
			return response, err
		}
	}

	return response, nil
}

func (c *Client) DecodeJSON(res *Response, v interface{}) error {
	return json.Unmarshal(res.BodyBytes, v)
}

func (c *Client) checkResponse(res *Response) error {
	if res.Response.StatusCode >= 200 && res.Response.StatusCode <= 299 {
		return nil
	}

	return c.decodeErrorResponse(res)
}

func (c *Client) decodeErrorResponse(res *Response) error {
	// Try to decode error response or fallback with standard error
	v := &Error{ErrorResponse: res}
	if err := c.DecodeJSON(res, v); err != nil {
		return fmt.Errorf("%s API call to %s failed: %v", res.Response.Request.Method, res.Response.Request.URL.String(), res.Response.Status)
	}

	return v
}
