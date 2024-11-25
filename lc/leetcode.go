package main

import "fmt"

type timeValue struct {
	value     string
	timeStamp int
}
type TimeMap struct {
	timeMap map[string][]timeValue
}

func Constructor() TimeMap {
	timeMap := make(map[string][]timeValue)
	return TimeMap{
		timeMap: timeMap,
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {

	timeValueObj := timeValue{
		value,
		timestamp,
	}
	this.timeMap[key] = append(this.timeMap[key], timeValueObj)

}

func (this *TimeMap) Get(key string, timestamp int) string {
	val, ok := this.timeMap[key]
	if !ok {
		return ""
	}
	var vals string
	l, r := 0, len(val)-1
	for l <= r {
		m := l + (r-l)/2
		if val[m].timeStamp <= timestamp {
			vals = val[m].value
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return vals
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
func lc() {
	obj := Constructor()
	fmt.Println(obj)
}
