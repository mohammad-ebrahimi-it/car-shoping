package logging

func mapToZapParams(extra map[ExtraKey]interface{}) []interface{} {
	params := make([]interface{}, 0)
	for key, value := range extra {
		params = append(params, string(key))
		params = append(params, value)
	}

	return params
}

func prepareZepParams(cat Category, sub SubCategory, extra map[ExtraKey]interface{}) []interface{} {
	if extra == nil {
		extra = make(map[ExtraKey]interface{}, 0)
	}
	extra["Category"] = cat
	extra["SubCategory"] = sub
	params := mapToZapParams(extra)

	return params
}

func prepareZeroParams(extra map[ExtraKey]interface{}) map[string]interface{} {
	params := map[string]interface{}{}

	for key, value := range extra {
		params[string(key)] = value
	}

	return params
}
