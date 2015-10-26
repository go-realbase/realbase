package integrationtests

import (
	"github.com/go-neutrino/neutrino/client"
	"github.com/go-neutrino/neutrino/config"
	"github.com/go-neutrino/neutrino/log"
	"github.com/go-neutrino/neutrino/realtime-service/client"
	"math/rand"
	"strconv"
	"time"
)

var (
	ApiBaseUrl     = "http://localhost" + config.Get(config.KEY_API_PORT) + "/v1/"
	AppId          = ""
	Email          = ""
	Password       = ""
	ApiClient      *client.ApiClient
	RealtimeClient *neutrinoclient.NeutrinoClient
	RealtimeData   *neutrinoclient.NeutrinoData
)

func randomString() string {
	rand.Seed(time.Now().UnixNano())
	return "r" + strconv.Itoa(rand.Int())
}

func init() {
	//initialize the tests
	Email = randomString() + "@gmail.com"
	Password = randomString()

	ApiClient = client.NewApiClient(ApiBaseUrl, "")

	ApiClient.Register(Email, Password)
	ApiClient.Login(Email, Password)
	AppId, err := ApiClient.CreateApp(randomString())
	if err != nil {
		log.Error(err)
		return
	}

	ApiClient.AppId = AppId

	RealtimeClient = neutrinoclient.NewClient(AppId)
	RealtimeClient.Token = ApiClient.Token
	RealtimeData = RealtimeClient.Data("test")
}
