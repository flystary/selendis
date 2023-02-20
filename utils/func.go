package utils

import "bytes"




func PK(endpoint, metric string, tags map[string]string)string {
	ret := bufferPool.Get().(*bytes.Buffer)
	ret.Reset()

	defer bufferPool.Put(ret)
	if tags == nil || len(tags) == 0 {
		ret.WriteString(endpoint)
		ret.WriteString("/")
		ret.WriteString(metric)

		return ret.String()
	}
	ret.WriteString(endpoint)
	ret.WriteString("/")
	ret.WriteString(metric)
	ret.WriteString("/")
	ret.WriteString(SortedTags(tags))
	return ret.String()

}