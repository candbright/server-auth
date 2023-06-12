package domain

import (
	"fmt"
	"gorm.io/datatypes"
)

type ExternalIds struct {
	ExternalIds datatypes.JSONMap `gorm:"column:external_ids" json:"external_ids"`
}

func (ext *ExternalIds) SetExternal(key string, value interface{}) {
	if ext.ExternalIds == nil {
		ext.ExternalIds = make(map[string]interface{})
	}
	ext.ExternalIds[key] = value
}

func (ext *ExternalIds) GetExternal(key string) interface{} {
	return ext.ExternalIds[key]
}

func (ext *ExternalIds) DeleteExternal(key string) {
	delete(ext.ExternalIds, key)
}

func (ext *ExternalIds) GetString(key string) string {
	value := ext.GetExternal(key)
	if value == nil {
		return ""
	}
	return fmt.Sprint(value)
}
