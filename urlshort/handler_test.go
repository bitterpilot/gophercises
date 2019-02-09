package urlshort

import (
	"reflect"
	"testing"
)

func TestMapHandler(t *testing.T) {

}

func TestYAMLHandler(t *testing.T) {

}

func TestParseYaml(t *testing.T) {

}

func TestBuildMap(t *testing.T) {
	for _, val := range buildMapCases {
		got := buildMap(val.input)

		eq := reflect.DeepEqual(got, val.expected)
		if !eq {
			t.Fail()
		}
	}
}
