package ngingo

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"hash"
	"strconv"
	"time"
)

const (
	// DefaultSecretKey is key of secret
	DefaultSecretKey = "ushiosngingodefaultsecret"
)

// SecureLinkSecret return secret string for ngx_http_secure_link_module
// see: http://nginx.org/en/docs/http/ngx_http_secure_link_module.html
func SecureLinkSecret(path string, expire time.Time) string {
	return secureLinkSecret(path, md5.New(), expire)
}

// secureLinkSecret return secret string
// not base64 encoded.
func secureLinkSecret(path string, h hash.Hash, expire time.Time) string {
	var b bytes.Buffer

	b.WriteString(DefaultSecretKey)
	b.WriteString(path)
	b.WriteString(strconv.FormatInt(expire.Unix(), 10))

	h.Write(b.Bytes())
	return hex.EncodeToString(h.Sum(nil))
}
