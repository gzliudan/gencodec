// Code generated by github.com/fjl/gencodec. DO NOT EDIT.

package ftypes

import (
	"encoding/json"
)

func (x X) MarshalJSON() ([]byte, error) {
	type X struct {
		Int int
		Err error
		If  interface{}
	}
	var enc X
	enc.Int = x.Int
	enc.Err = x.Err
	enc.If = x.If
	return json.Marshal(&enc)
}

func (x *X) UnmarshalJSON(input []byte) error {
	type X struct {
		Int *int
		Err *error
		If  *interface{}
	}
	var dec X
	if err := json.Unmarshal(input, &dec); err != nil {
		return err
	}
	if dec.Int != nil {
		x.Int = *dec.Int
	}
	if dec.Err != nil {
		x.Err = *dec.Err
	}
	if dec.If != nil {
		x.If = *dec.If
	}
	return nil
}
