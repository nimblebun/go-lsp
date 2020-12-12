package lsp

import (
	"encoding/json"
	"strconv"
)

// ID is a JSON-RPC 2.0 request ID. It can be either a string or a number.
type ID struct {
	AsInteger uint64
	AsString  string

	IsString bool
}

func (id ID) String() string {
	if id.IsString {
		return strconv.Quote(id.AsString)
	}

	return strconv.FormatUint(id.AsInteger, 10)
}

// MarshalJSON will turn the ID into a JSON string.
func (id ID) MarshalJSON() ([]byte, error) {
	if id.IsString {
		return json.Marshal(id.AsString)
	}

	return json.Marshal(id.AsInteger)
}

// UnmarshalJSON will turn the passed data into an ID struct.
func (id *ID) UnmarshalJSON(data []byte) error {
	var asInteger uint64
	if err := json.Unmarshal(data, &asInteger); err == nil {
		*id = ID{AsInteger: asInteger}
		return nil
	}

	var asString string
	if err := json.Unmarshal(data, &asString); err != nil {
		return err
	}

	*id = ID{
		AsString: asString,
		IsString: true,
	}

	return nil
}
