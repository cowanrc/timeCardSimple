package dto

import "timeCardSimple/app/domain/id"

func Id(v interface{}) interface{} {
	return v.(id.ID).String()
}

func Ids(v interface{}) interface{} {
	ids := v.([]id.ID)
	result := make([]string, len(ids))
	for i, id := range ids {
		result[i] = id.String()
	}
	return result
}
