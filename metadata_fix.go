package ledgerclient

import (
	"fmt"
	_neturl "net/url"
)

func addMetadataParams(queryParams _neturl.Values, metadata map[string]interface{}) {
	for k, v := range metadata {
		key, value := walkMeta(k, v)
		queryParams.Add(fmt.Sprintf("metadata[%s]", key), parameterToString(value, ""))
	}
}

func walkMeta(key string, value interface{}) (string, interface{}) {
	if nested, ok := value.(map[string]interface{}); ok {
		for k, v := range nested {
			return walkMeta(fmt.Sprintf("%s.%s", key, k), v)
		}
	}
	return key, value
}
