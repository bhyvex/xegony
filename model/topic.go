package model

import (
	"fmt"

	"github.com/xeipuuv/gojsonschema"
)

type Topic struct {
	Id      int64  `json:"id"`
	Title   string `json:"title"`
	ForumId int64  `json:"forumId" db:"forum_id"`
	Icon    string `json:"icon"`
}

func (c *Topic) NewSchema(requiredFields []string, optionalFields []string) (schema *gojsonschema.Schema, err error) {
	s := Schema{}
	s.Type = "object"
	s.Required = requiredFields
	s.Properties = make(map[string]Schema)
	var field string
	var prop Schema
	for _, field = range requiredFields {
		if prop, err = c.getSchemaProperty(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	for _, field := range optionalFields {
		if prop, err = c.getSchemaProperty(field); err != nil {
			return
		}
		s.Properties[field] = prop
	}
	jsRef := gojsonschema.NewGoLoader(s)
	schema, err = gojsonschema.NewSchema(jsRef)
	if err != nil {
		return
	}
	return
}

func (c *Topic) getSchemaProperty(field string) (prop Schema, err error) {
	switch field {
	case "id":
		prop.Type = "integer"
		prop.Minimum = 1
	case "body":
		prop.Type = "string"
		prop.MinLength = 3
		prop.MaxLength = 1024
		prop.Pattern = "^[a-zA-Z' ]*$"
	default:
		err = fmt.Errorf("Invalid field passed: %s", field)
	}

	return
}
