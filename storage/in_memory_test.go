package storage_test

import (
	urlstorage "linkly/storage"
	"testing"
)

func TestSaveUrl(t *testing.T) {
	expectedId := "1"
	id := urlstorage.SaveUrl("http://www.google.com/")
	if expectedId != id {
		t.Errorf("expected id %s, got %s", expectedId, id)
	}
}

func TestGetUrl(t *testing.T) {
	url := "https://example.com"
	id := urlstorage.SaveUrl(url)

	retrievedUrl := urlstorage.GetUrl(id)
	if retrievedUrl != url {
		t.Errorf("expected url %s, got %s", url, retrievedUrl)
	}
}
