package zip2latlong

import (
	"reflect"
	"testing"
)

func TestLookupZip(t *testing.T) {
	l, err := NewLocator()
	if err != nil {
		t.Fatal(err)
	}

	gotZip, err := l.LookupZip("44120")
	if err != nil {
		t.Fatal(err)
	}

	expectedZip := ZipCode{Code: "44120", City: "SHAKER HTS", Lat: 41.47, Long: -81.67}

	if !reflect.DeepEqual(expectedZip, gotZip) {
		t.Errorf("expected: %+v, got: %+v", expectedZip, gotZip)
	}
}
