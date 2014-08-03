package hcl

import (
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"
)

func TestDecode(t *testing.T) {
	cases := []struct {
		File string
		Err  bool
		Out  interface{}
	}{
		{
			"basic.hcl",
			false,
			map[string]interface{}{
				"foo": "bar",
			},
		},
		{
			"structure.hcl",
			false,
			map[string]interface{}{
				"foo": []interface{}{
					map[string]interface{}{
						"baz": []interface{}{
							map[string]interface{}{
								"key": 7,
								"foo": "bar",
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range cases {
		d, err := ioutil.ReadFile(filepath.Join(fixtureDir, tc.File))
		if err != nil {
			t.Fatalf("err: %s", err)
		}

		var out map[string]interface{}
		err = Decode(&out, string(d))
		if (err != nil) != tc.Err {
			t.Fatalf("Input: %s\n\nError: %s", tc.File, err)
		}

		if !reflect.DeepEqual(out, tc.Out) {
			t.Fatalf("Input: %s\n\n%#v", tc.File, out)
		}
	}
}

func TestDecode_equal(t *testing.T) {
	cases := []struct {
		One, Two string
	}{
		{
			"basic.hcl",
			"basic.json",
		},
		{
			"structure.hcl",
			"structure.json",
		},
		{
			"structure.hcl",
			"structure_flat.json",
		},
	}

	for _, tc := range cases {
		p1 := filepath.Join(fixtureDir, tc.One)
		p2 := filepath.Join(fixtureDir, tc.Two)

		d1, err := ioutil.ReadFile(p1)
		if err != nil {
			t.Fatalf("err: %s", err)
		}

		d2, err := ioutil.ReadFile(p2)
		if err != nil {
			t.Fatalf("err: %s", err)
		}

		var i1, i2 interface{}
		err = Decode(&i1, string(d1))
		if err != nil {
			t.Fatalf("err: %s", err)
		}

		err = Decode(&i2, string(d2))
		if err != nil {
			t.Fatalf("err: %s", err)
		}

		if !reflect.DeepEqual(i1, i2) {
			t.Fatalf(
				"%s != %s\n\n%#v\n\n%#v",
				tc.One, tc.Two,
				i1, i2)
		}
	}
}
