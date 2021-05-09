package function

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
)

// Twitter Client shared between function calls
var client *twitter.Client

func init() {

	apiKey, err := getAPISecret("secret-twitter-api-key")
	if err != nil {
		log.Fatalln("apiKey not found.")
	}

	apiSecret, err := getAPISecret("secret-twitter-api-secret")
	if err != nil {
		log.Fatalln("apiSecret not found.")
	}

	accessToken, err := getAPISecret("secret-twitter-access-token")
	if err != nil {
		log.Fatalln("accessToken not found.")
	}

	accessSecret, err := getAPISecret("secret-twitter-access-secret")
	if err != nil {
		log.Fatalln("accessSecret not found.")
	}

	config := oauth1.NewConfig(string(apiKey), string(apiSecret))
	token := oauth1.NewToken(string(accessToken), string(accessSecret))
	httpClient := config.Client(oauth1.NoContext, token)

	// Twitter client
	client = twitter.NewClient(httpClient)
}

func Handle(w http.ResponseWriter, r *http.Request) {

	fmt.Println("Received HTTP request.")

	var input []byte

	if r.Body != nil {
		defer r.Body.Close()

		body, _ := ioutil.ReadAll(r.Body)

		input = body
	}

	fmt.Printf("%s\n", string(input))

	// TODO: Parse temp, humdity & pressure from payload

	// Send a Tweet
	tweet, resp, err := client.Statuses.Update(".", nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Fatalf("Failed to send tweet. Error: %s", err.Error())
	}

	fmt.Printf("Tweet sent: %s", tweet.FullText)

	respBody, _ := ioutil.ReadAll(resp.Body)

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Response body: %s", string(respBody))))
}

func getAPISecret(secretName string) (secretBytes []byte, err error) {

	// read from the faasd-provider secrets folder
	secretBytes, err = ioutil.ReadFile("/var/lib/faasd-provider/secrets/" + secretName)

	if err != nil {
		// read from the original location for backwards compatibility with openfaas <= 0.8.2
		secretBytes, err = ioutil.ReadFile("/run/secrets/" + secretName)
	}

	return secretBytes, err
}
