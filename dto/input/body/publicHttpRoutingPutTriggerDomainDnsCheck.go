// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PublicHttpRoutingPutTriggerDomainDnsCheck)(nil)

type PublicHttpRoutingPutTriggerDomainDnsCheck struct {
	Domains types.StringArrayNull `json:"domains"`
}

func (dto PublicHttpRoutingPutTriggerDomainDnsCheck) GetDomains() types.StringArrayNull {
	return dto.Domains
}

func (dto *PublicHttpRoutingPutTriggerDomainDnsCheck) UnmarshalJSON(b []byte) error {
	var aux = struct {
		Domains types.StringArrayNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PublicHttpRoutingPutTriggerDomainDnsCheck", err)
	}
	var errorList validator.ErrorList

	if errorList != nil {
		return errorList.GetError()
	}
	dto.Domains = aux.Domains

	return nil
}
