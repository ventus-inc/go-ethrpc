package ethrpc

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
	. "github.com/smartystreets/goconvey/convey"
)

func init() {
	err := godotenv.Load()
	if err != nil && os.Getenv("GO_ENV") != "test" {
		log.Fatal("Error loading .env file")
	}
}

func TestRPCClient(t *testing.T) {
	Convey("Success", t, func() {
		c := NewRPCClient(os.Getenv("GETH_ENDPOINT"))
		_, err := c.GetBlockNumber()
		So(err, ShouldBeNil)
	})
}

func TestGetBlockNumber(t *testing.T) {
	Convey("Success", t, func() {
		c := NewRPCClient(os.Getenv("GETH_ENDPOINT"))
		_, err := c.GetBlockNumber()
		So(err, ShouldBeNil)
	})
}

func TestCall(t *testing.T) {
	mockRepo := MockEthClient{}
	mockRepo.On("Call", "0xfrom", "0xto", "0xdata").Return("0xtrue", nil)
	mockRepo.On("Call", "0x0", "0x0", "0x0").Return("", "Some error")
	c := mockRepo
	Convey("Success", t, func() {
		//		c := NewRPCClient(os.Getenv("GETH_ENDPOINT"))
		_, err := c.Call("0xfrom", "0xto", "0xdata")
		So(err, ShouldBeNil)
	})
}
