package ledgerclient

import (
	"testing"

	"net/url"
	_neturl "net/url"

	"github.com/stretchr/testify/assert"
)

func Test_addMetadataParams(t *testing.T) {
	child := map[string]interface{}{"child": "James"}
	meta := map[string]interface{}{"parent": child, "ref": "ref01"}

	params := _neturl.Values{}
	addMetadataParams(params, meta)

	assert.Equal(t, "ref01", params.Get("metadata[ref]"))
	assert.Equal(t, "James", params.Get("metadata[parent.child]"))
	v, _ := url.QueryUnescape(params.Encode())
	assert.Equal(t, "metadata[parent.child]=James&metadata[ref]=ref01", v)
}
