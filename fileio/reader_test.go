package fileio

import (
	"testing"
)

func Test_IsValidURL_Success(t *testing.T) {
	url_str := "https://github.com/facebook/react"
	valid := IsValidURL(url_str)
	if valid == false {
		t.Errorf("want %t, got %t.", true, valid)
	}
}

func Test_IsValidURl_Fail(t *testing.T) {
	bad_url_str := "https://google.com/fakebok/real"
	valid := IsValidURL(bad_url_str)
	if valid == true {
		t.Errorf("want %t, got %t.", false, valid)
	}
}

func Test_MakeUrlChannel_Success(t *testing.T) {
	result := MakeUrlChannel() 
	if result == nil {
		t.Errorf("MakeUrlChannel failed, returned nil")
	}
}