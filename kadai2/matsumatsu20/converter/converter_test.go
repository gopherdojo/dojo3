package converter

import "testing"

func TestValidateFormat(t *testing.T) {
	noError := func(err error) {
		if err != nil {
			t.Helper()
			t.Errorf("expected no error")
		}
	}

	withError := func(err error) {
		if err == nil {
			t.Helper()
			t.Errorf("expected returning error")
		}
	}

	cases := []struct {
		format   string
		assertFn func(err error)
	}{
		{format: "jpeg", assertFn: noError},
		{format: "jpg", assertFn: noError},
		{format: "png", assertFn: noError},
		{format: "gif", assertFn: noError},
		{format: "other", assertFn: withError},
	}

	for _, v := range cases {
		err := ValidateFormat(v.format)
		v.assertFn(err)
	}
}

func TestIsConvertTargetFormat(t *testing.T) {
	cases := []struct {
		converter Converter
		format    string
		expected  bool
	}{
		{converter: Converter{FilePath: "dir", InputFormat: "jpeg", OutputFormat: "png"}, format: "jpeg", expected: true},
		{converter: Converter{FilePath: "dir", InputFormat: "jpg", OutputFormat: "png"}, format: "jpeg", expected: true},
		{converter: Converter{FilePath: "dir", InputFormat: "jpeg", OutputFormat: "png"}, format: "gif", expected: false},
	}

	for _, v := range cases {
		if v.converter.isConvertTargetFormat(v.format) != v.expected {
			t.Errorf("expected %v", v.expected)
		}
	}

}

func TestGetDistFileName(t *testing.T) {
	cases := []struct {
		converter Converter
		expected  string
	}{
		{converter: Converter{FilePath: "dir/image.png", InputFormat: "png", OutputFormat: "gif"}, expected: "dir/image.gif"},
		{converter: Converter{FilePath: "dir/image", InputFormat: "png", OutputFormat: "gif"}, expected: "dir/image.gif"},
	}

	for _, v := range cases {
		if v.converter.GetDistFileName() != v.expected {
			t.Errorf("expected %v", v.expected)
		}
	}
}
