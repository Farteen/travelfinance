package testing

import (
	"encoding/hex"
	"github.com/mongodb/mongo-go-driver/x/mongo/driver/uuid"
	"github.com/onsi/gomega"
	"testing"
)

func TestUUIDGenerator(t *testing.T) {
	uuidBytes, err := uuid.New()
	if err != nil {
		t.Fatal("uuid err")
	}

	hexString := hex.EncodeToString(uuidBytes[0:])
	gomega.Expect(true, len(hexString) > 0)

}
