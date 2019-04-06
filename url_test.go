package marshaler_test

import (
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"testing"

	marshaler "github.com/xescugc/url-marshaler"
)

type TestStruct struct {
	URL marshaler.URL `json:"url"`
}

func TestMarshalJSON(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		u, err := url.Parse("http://example.com")
		if err != nil {
			t.Fatalf("error parsing the url %s", err)
		}

		eu := fmt.Sprintf(`{"url":"%s"}`, u)
		ts := TestStruct{URL: marshaler.URL{URL: u}}

		b, err := json.Marshal(ts)
		if err != nil {
			t.Fatalf("error marshalling the struct %s", err)
		}

		if string(b) != eu {
			t.Errorf("expected %q to be equal to %q", b, eu)
		}
	})
}

func TestUnmarshalJSON(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		u, err := url.Parse("http://example.com")
		if err != nil {
			t.Fatalf("error parsing the url %s", err)
		}

		eu := []byte(fmt.Sprintf(`{"url": "%s"}`, u))
		ets := TestStruct{URL: marshaler.URL{URL: u}}
		var ts TestStruct

		err = json.Unmarshal(eu, &ts)
		if err != nil {
			t.Fatalf("error unmarshaling the struct %s", err)
		}

		if !reflect.DeepEqual(ets, ts) {
			t.Errorf("expected %+v to be equal to %+v", ets, ts)
		}
	})
}
