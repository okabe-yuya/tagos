package aggregate

import (
	"strconv"
	"strings"
)

func ToResponse(tagoses []map[string]interface{}) string {
	aggregated := aggregateReceiver(tagoses)
	joined := joinAggregateResult(aggregated)
	return joined
}

func aggregateReceiver(tagoses []map[string]interface{}) map[string]int {
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

func joinAggregateResult(result map[string]int) string {
	resp := "- -=集計結果=- -\n"
	if len(result) > 0 {
		for receiver, count := range result {
			resp += receiver + " さん: 🌮を" + strconv.Itoa(count) + "つ頂いています\n"
		}
	} else {
		resp += "まだ🌮が送られていないようです👀"
	}
	rmMention := strings.Replace(resp, "@", "", -1)
	return rmMention
}