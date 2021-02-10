package dto

//Request body definition
type Request struct {
	TextoATraducir string `json:"texto_a_traducir"`
	FormatoOrigen  string `json:"formato_origen"`
	FormatoDestino string `json:"formato_destino"`
}

// Response body definition
type Response struct {
	TextoTraducido string `json:"texto_traducido,omitempty"`
}

type ErrorResponse struct {
	Description string `json:"description,omitempty"`
}
