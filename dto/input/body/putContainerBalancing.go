// Generated ZEROPS sdk

package body

import (
	"encoding/json"
	"strconv"

	"github.com/zeropsio/zerops-go/types"
	"github.com/zeropsio/zerops-go/validator"
)

var _ strconv.NumError
var _ json.Unmarshaler = (*PutContainerBalancing)(nil)

type PutContainerBalancing struct {
	L3  types.BoolNull `json:"l3"`
	L7  types.BoolNull `json:"l7"`
	Dns types.BoolNull `json:"dns"`
}

func (dto PutContainerBalancing) GetL3() types.BoolNull {
	return dto.L3
}
func (dto PutContainerBalancing) GetL7() types.BoolNull {
	return dto.L7
}
func (dto PutContainerBalancing) GetDns() types.BoolNull {
	return dto.Dns
}

func (dto *PutContainerBalancing) UnmarshalJSON(b []byte) error {
	var aux = struct {
		L3  types.BoolNull
		L7  types.BoolNull
		Dns types.BoolNull
	}{}
	err := json.Unmarshal(b, &aux)
	if err != nil {
		return validator.JsonValidation("PutContainerBalancing", err)
	}
	var errorList validator.ErrorList

	if errorList != nil {
		return errorList.GetError()
	}
	dto.L3 = aux.L3
	dto.L7 = aux.L7
	dto.Dns = aux.Dns

	return nil
}
