package ip_hash

import (
	"crypto"
	mock_ip_hash "github.com/stockhut/hsfl-master-ai-cloud-engineering/load-balancer/strategies/ip-hash/_mocks"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"net/http"
	"testing"
)

func TestIpHash_New(t *testing.T) {

	t.Run("checks availability of the hash algorithm", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		mockAlgo := mock_ip_hash.NewMockHashAlgo(ctrl)
		mockAlgo.EXPECT().Available().Return(false).Times(1)

		_, err := New(mockAlgo)

		assert.Error(t, err)
	})

	t.Run("creates instance with algo", func(t *testing.T) {

		ctrl := gomock.NewController(t)
		mockAlgo := mock_ip_hash.NewMockHashAlgo(ctrl)
		mockAlgo.EXPECT().Available().Return(true).Times(1)

		strategy, err := New(mockAlgo)

		assert.Nil(t, err)
		assert.Equal(t, mockAlgo, strategy.algo)
	})
}

func TestIpHash_GetTarget(t *testing.T) {

	t.Run("uses ip to calculate hash", func(t *testing.T) {

		var hashInput []byte
		var h HashFunction = func(algo HashAlgo, input []byte) []byte {
			hashInput = input
			return []byte("hashvalue")
		}

		strategy := IpHash{
			algo:     crypto.SHA1,
			hashFunc: h,
		}

		req, err := http.NewRequest(http.MethodGet, "/", nil)
		assert.Nil(t, err, "Failed to create http request")

		req.RemoteAddr = "123.456.789:1234"

		strategy.GetTarget(req, []string{""}, func(t string) {
		})

		assert.Equal(t, []byte("123.456.789"), hashInput)

	})

	t.Run("uses hash to find target", func(t *testing.T) {

		replicas := []string{
			"first",
			"second",
		}
		for i, replica := range replicas {
			var h HashFunction = func(algo HashAlgo, input []byte) []byte {
				return []byte{byte(i)}
			}

			strategy := IpHash{
				algo:     crypto.SHA1,
				hashFunc: h,
			}

			req, err := http.NewRequest(http.MethodGet, "/", nil)

			assert.Nil(t, err, "Failed to create http request")

			var target string
			callback := func(t string) {
				target = t
			}
			strategy.GetTarget(req, replicas, callback)

			assert.Equal(t, replica, target)
		}

	})
}

func Test_defaultHashFunction(t *testing.T) {

	t.Run("hashes correctly", func(t *testing.T) {

		h := defaultHashFunction(crypto.SHA1, []byte("hello"))

		expH := crypto.SHA1.New()
		expH.Write([]byte("hello"))

		assert.Equal(t, expH.Sum(nil), h)
	})
}
