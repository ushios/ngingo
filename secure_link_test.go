package ngingo

import (
	"crypto/md5"
	"hash"
	"testing"
	"time"
)

func TestSecureLinkTest(t *testing.T) {
	test := func(path, expire string, hasher hash.Hash, ans string) {
		eTime, err := time.Parse(time.RFC3339, expire)
		if err != nil {
			t.Fatal(err)
		}

		secret := secureLinkSecret(path, hasher, eTime)

		if secret != ans {
			t.Errorf("(path: %s) (expire: %s) (algorithm: %v) secret expected (%s) but (%s)",
				path, expire, hasher, ans, secret,
			)
		}
	}

	test("/icon/image.png", "2016-01-10T12:00:00+09:00", md5.New(), "fe54de580ca84cadcdd7c23715f2039d")
}
