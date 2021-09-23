package responses

type AddressResponse struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Name    string      `json:"nama"`
	Data    interface{} `json:"data"`
}
