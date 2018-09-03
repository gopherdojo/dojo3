package sorcery

import (
	"testing"
	"fmt"
)

func TestImgExt(t *testing.T) {
	type checkFunc func(imgExt, error) error

	isSomeError := func() checkFunc {
		return func(ext imgExt, err error) error {
			if err == nil {
				return fmt.Errorf("expected error but has not occured")
			}
			return nil
		}
	}

	isValid := func(want imgExt) checkFunc {
		return func(ext imgExt, err error) error {
			if ext != want {
				return fmt.Errorf("unexpected imgExt %v", ext)
			}
			return nil
		}
	}

	tests := [...]struct{
		name string
		arg string
		check checkFunc
	}{
		{"JPEG should be accepted", "JPEG", isValid(Jpeg)},
		{"jpeg should be accepted", "jpeg", isValid(Jpeg)},
		{"JPG should be accepted", "JPG", isValid(Jpeg)},
		{"jpg should be accepted", "jpg", isValid(Jpeg)},
		{".JPEG should be accepted", ".JPEG", isValid(Jpeg)},
		{".jpeg should be accepted", ".jpeg", isValid(Jpeg)},
		{".JPG should be accepted", ".JPG", isValid(Jpeg)},
		{".jpg should be accepted", ".jpg", isValid(Jpeg)},
		{"PNG should be accepted", "PNG", isValid(Png)},
		{"png should be accepted", "png", isValid(Png)},
		{".PNG should be accepted", ".PNG", isValid(Png)},
		{"png should be accepted", ".png", isValid(Png)},
		{"GIF should be accepted", "GIF", isValid(Gif)},
		{"gif should be accepted", "gif", isValid(Gif)},
		{".GIF should be accepted", ".GIF", isValid(Gif)},
		{".gif should be accepted", ".gif", isValid(Gif)},
		{"TIFF should not be accepted", "TIFF", isSomeError()},
		{"TIFF should not be accepted", "svg", isSomeError()},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			imgExt, err := ImgExt(tt.arg)
			if err := tt.check(imgExt, err); err != nil {
				t.Error(err)
			}
		})
	}
}

func TestString(t *testing.T) {
	type checkFunc func(string) error

	isExpected := func(want string) checkFunc {
		return func(str string) error {
			if str != want {
				return fmt.Errorf("expected %s but returned %s", want, str)
			}
			return nil
		}
	}

	tests := [...]struct{
		name string
		arg imgExt
		check checkFunc
	}{
		{"Jpeg", Jpeg, isExpected("jpg")},
		{"Png", Png, isExpected("png")},
		{"Gif", Gif, isExpected("gif")},
		{"Unknown", end, isExpected("Unknown")},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			str := tt.arg.String()
			if err := tt.check(str); err != nil {
				t.Error(err)
			}
		})
	}
}

func testIsValid(t *testing.T) {
	type checkFunc func(bool) error

	isExpected := func(want bool) checkFunc {
		return func(b bool) error {
			if b != want {
				return fmt.Errorf("expected %v but returned %v", want, b)
			}
			return nil
		}
	}

	tests := [...]struct{
		name string
		arg imgExt
		check checkFunc
	}{
		{"Jpeg", Jpeg, isExpected(true)},
		{"Png", Png, isExpected(true)},
		{"Gif", Gif, isExpected(true)},
		{"Unknown", end, isExpected(false)},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b := tt.arg.isValid()
			if err := tt.check(b); err != nil {
				t.Error(err)
			}
		})
	}
}
