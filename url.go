package marshaler

import (
	"encoding/json"
	"net/url"
)

type URL struct {
	*url.URL
}

func (u *URL) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	url, err := url.Parse(s)
	if err != nil {
		return err
	}

	u.URL = url
	return nil
}

func (u URL) MarshalJSON() ([]byte, error) {
	if u.URL == nil {
		return json.Marshal(nil)
	}

	return json.Marshal(u.URL.String())
}
