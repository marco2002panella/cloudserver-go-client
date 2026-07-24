package seeweb

type ImageService service

/*===============================================================================================================================================================================================
 * [STRUCTS / DATA MODELS] - Strutture dati
 *==============================================================================================================================================================================================*/

type Image struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	CreationDate string `json:"creation_date"`
	ActiveFlag   bool   `json:"active_flag"`
	Status       string `json:"status"`
	UUID         string `json:"uuid"`
	Description  string `json:"description"`
	Notes        string `json:"notes"`
	Public       bool   `json:"public"`
	CloudImage   bool   `json:"cloud_image"`
	SoBase       string `json:"so_base"`
	ApiVersion   string `json:"api_version"`
	Version      string `json:"version"`
}

/*===============================================================================================================================================================================================
 * [LIST BASICS/ READ ALL BASICS] - Recupera l'elenco completo delle risorse BASICS
 *==============================================================================================================================================================================================*/

type SeewebImageListResponse struct {
	Status string   `json:"status"`
	Images []*Image `json:"images"`
}

func (i *ImageService) ListBasics() (*SeewebImageListResponse, *Response, error) {
	u := "/images/basics"
	v := new(SeewebImageListResponse)

	resp, err := i.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}

/*===============================================================================================================================================================================================
 * [LIST CLOUD-IMAGES / READ ALL CLOUD-IMAGES] - Recupera l'elenco completo delle risorse CLOUD-IMAGES
 *==============================================================================================================================================================================================*/

func (i *ImageService) ListCloudImages() (*SeewebImageListResponse, *Response, error) {
	u := "/images/cloud-images"
	v := new(SeewebImageListResponse)

	resp, err := i.client.newRequestDo("GET", u, nil, &v)
	if err != nil {
		return nil, nil, err
	}

	return v, resp, nil
}
