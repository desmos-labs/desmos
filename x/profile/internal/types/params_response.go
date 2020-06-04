package types

import "encoding/json"

// ParamsQueryResponse represents the profile params
// that are returned to user upon a query
type ParamsQueryResponse struct {
	NameSurnameLenParams NameSurnameLenParams `json:"name_surname_len_params"`
	MonikerLenParams     MonikerLenParams     `json:"moniker_len_params"`
	BioLenParams         BioLenParams         `json:"bio_len_params"`
}

func NewParamsQueryResponse(nsParams NameSurnameLenParams, mParams MonikerLenParams, bParams BioLenParams) ParamsQueryResponse {
	return ParamsQueryResponse{
		NameSurnameLenParams: nsParams,
		MonikerLenParams:     mParams,
		BioLenParams:         bParams,
	}
}

// MarshalJSON implements json.Marshaler as Amino does
// not respect default json composition
func (response ParamsQueryResponse) MarshalJSON() ([]byte, error) {
	type temp ParamsQueryResponse
	return json.Marshal(temp(response))
}

// UnmarshalJSON implements json.Unmarshaler as Amino does
// not respect default json composition
func (response *ParamsQueryResponse) UnmarshalJSON(data []byte) error {
	type postResponse ParamsQueryResponse
	var temp postResponse
	if err := json.Unmarshal(data, &temp); err != nil {
		return err
	}

	*response = ParamsQueryResponse(temp)
	return nil
}
