package aggregate

func AggregateReceiver(tagoses []map[string]interface{}) map[string]int {
	resp := make(map[string]int, 0)
	for _, tagos := range tagoses {
		receiver := tagos["receiver"]
		str := receiver.(string)
		if _, ok := resp[str]; ok {
			resp[str] = resp[str] + 1
		} else {
			resp[str] = 1
		}
	}
	return resp
}