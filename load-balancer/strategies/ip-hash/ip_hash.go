package ip_hash

import (
	"crypto"
	"errors"
	"fmt"
	"hash"
	"net/http"
	"strings"
)

// HashAlgo models crypto.Hash as interface for easy mocking
type HashAlgo interface {
	HashFunc() crypto.Hash
	String() string
	Size() int
	New() hash.Hash
	Available() bool
}

// HashFunction calculates the hash of input using the specified algorithm.
// It does not have to check for the availability of the hash function.
type HashFunction func(hash HashAlgo, input []byte) []byte

func defaultHashFunction(hash HashAlgo, input []byte) []byte {
	h := hash.New()
	h.Write(input)
	return h.Sum(nil)
}

// The IpHash strategy picks the target host based in the clients IP Address
type IpHash struct {
	algo     HashAlgo
	hashFunc HashFunction
}

// New creates an IpHash instance using the given hash function
func New(hashAlgo HashAlgo) (*IpHash, error) {

	if !hashAlgo.Available() {
		return nil, errors.New("hash algorithm is not available")
	}
	return &IpHash{
		algo:     hashAlgo,
		hashFunc: defaultHashFunction,
	}, nil
}

func (ipHash *IpHash) GetTarget(r *http.Request, replicas []string, f func(host string)) {

	remoteIp := strings.Split(r.RemoteAddr, ":")[0]
	sum := ipHash.hashFunc(ipHash.algo, []byte(remoteIp))

	i := int32(sum[len(sum)-1]) % int32(len(replicas))

	host := replicas[i]

	fmt.Printf("Picked %s from healthy list: %v\n", host, replicas)
	f(host)
}
