package conv

func StringMap(value interface{}) map[string]interface{} {
	if mapValue, ok := value.(map[interface{}]interface{}); ok {
		rv := map[string]interface{}{}
		for key, value := range mapValue {
			if keyString, ok := key.(string); ok {
				rv[keyString] = value
			}
		}
		return rv
	}
	return nil
}

// TODO implement StringMapDeep(value interface{}) map[string]interface{}
