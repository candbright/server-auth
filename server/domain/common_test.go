package domain

import "testing"

func TestExternalIds_GetExternal(t *testing.T) {
	tests := []struct {
		key   string
		value interface{}
	}{
		{"key1", 10},
		{"key2", "value2"},
		{"key3", true},
	}
	ext := ExternalIds{}
	for _, data := range tests {
		ext.SetExternal(data.key, data.value)
	}
	t.Log("key0:", ext.GetExternal("key0"))
	for _, data := range tests {
		t.Log(data.key+":", ext.GetExternal(data.key))
	}
}

func TestExternalIds_DeleteExternal(t *testing.T) {
	tests := []struct {
		key   string
		value interface{}
	}{
		{"key1", 10},
		{"key2", "value2"},
		{"key3", true},
	}
	ext := ExternalIds{}
	for _, data := range tests {
		ext.SetExternal(data.key, data.value)
	}
	t.Log(ext)
	ext.DeleteExternal("key1")
	t.Log(ext)
}

func TestExternalIds_GetString(t *testing.T) {
	tests := []struct {
		key   string
		value interface{}
	}{
		{"key1", 10},
		{"key2", "value2"},
		{"key3", true},
	}
	ext := ExternalIds{}
	for _, data := range tests {
		ext.SetExternal(data.key, data.value)
	}
	t.Log("key0:", ext.GetString("key0"))
	for _, data := range tests {
		t.Log(data.key+":", ext.GetString(data.key))
	}
}
