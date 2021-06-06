package goconsumer

// import (
// 	"errors"
// 	"fmt"
// 	"net/http"
// 	"net/http/httptest"
// 	"net/url"
// 	"os"
// 	"strings"
// 	"testing"

// 	ex "github.com/pact-foundation/pact-go/examples/types"
// 	dsl "github.com/pact-foundation/pact-go/v3"
// 	v3 "github.com/pact-foundation/pact-go/v3"
// )

// // Common test data
// var dir, _ = os.Getwd()
// var pactDir = fmt.Sprintf("%s/../../pacts", dir)
// var logDir = fmt.Sprintf("%s/log", dir)
// var form url.Values
// var rr http.ResponseWriter
// var req *http.Request
// var name = "jmarie"
// var password = "issilly"
// var like = dsl.Like
// var eachLike = dsl.EachLike
// var term = dsl.Term

// type s = dsl.String
// type request = dsl.Request

// var loginRequest = ex.LoginRequest{
// 	Username: name,
// 	Password: password,
// }
// var commonHeaders = dsl.MapMatcher{
// 	"Content-Type": term("application/json; charset=utf-8", `application\/json`),
// }

// var pending bool

// // Use this to control the setup and teardown of Pact
// func TestMain(m *testing.M) {
// 	// Setup Pact and related test stuff
// 	setup()

// 	// Run all the tests
// 	code := m.Run()

// 	os.Exit(code)
// }

// // Setup common test data
// func setup() {
// 	// Login form values
// 	form = url.Values{}
// 	form.Add("username", name)
// 	form.Add("password", password)

// 	// Create a request to pass to our handler.
// 	req, _ = http.NewRequest("POST", "/login", strings.NewReader(form.Encode()))
// 	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
// 	req.PostForm = form

// 	rr = httptest.NewRecorder()

// 	// Pending integration test
// 	if os.Getenv("PENDING") != "" {
// 		pending = true
// 	}
// }

// func TestExampleConsumerLoginHandler_UserExists(t *testing.T) {
// 	var testJmarieExists = func(config dsl.MockServerConfig) error {
// 		client := Client{
// 			Host: fmt.Sprintf("http://localhost:%d", config.Port),
// 			// This header will be dynamically replaced at verification time
// 			token: "Bearer 1234",
// 		}
// 		client.loginHandler(rr, req)

// 		// Expect User to be set on the Client
// 		if client.user == nil {
// 			return errors.New("Expected user not to be nil")
// 		}

// 		return nil
// 	}

// 	pact, _ := dsl.NewV2Pact(v3.MockHTTPProviderConfig{
// 		Consumer:                 "jmarie",
// 		Provider:                 "loginprovider",
// 		Host:                     "127.0.0.1",
// 		LogDir:                   logDir,
// 		PactDir:                  pactDir,
// 		DisableToolValidityCheck: true,
// 	})

// 	// Pull from pact broker, used in e2e/integrated tests for pact-go release
// 	// Setup interactions on the Mock Service. Note that you can have multiple
// 	// interactions
// 	pact.
// 		AddInteraction().
// 		Given("User jmarie exists").
// 		UponReceiving("A request to login with user 'jmarie'").
// 		WithRequest("GET", term("/login/10", "/login/[0-9]+")).
// 		Query(dsl.QueryMatcher{
// 			"foo": []dsl.Matcher{term("bar", "[a-zA-Z]+")},
// 		}).
// 		Headers(commonHeaders).
// 		JSON(loginRequest).
// 		WillRespondWith(200).
// 		BodyMatch(ex.LoginResponse{
// 			User: &ex.User{},
// 		}).
// 		Headers(dsl.MapMatcher{
// 			"X-Api-Correlation-Id": dsl.Like("100"),
// 			"Content-Type":         term("application/json; charset=utf-8", `application\/json`),
// 			"X-Auth-Token":         dsl.Like("1234"),
// 		})

// 	err := pact.ExecuteTest(testJmarieExists)
// 	if err != nil {
// 		t.Fatalf("Error on Verify: %v", err)
// 	}
// }

// // func TestExampleConsumerLoginHandler_UserDoesNotExist(t *testing.T) {
// // 	var testJmarieDoesNotExists = func() error {
// // 		client := Client{
// // 			Host: fmt.Sprintf("http://localhost:%d", pact.Server.Port),
// // 			// This header will be dynamically replaced at verification time
// // 			token: "Bearer 1234",
// // 		}
// // 		client.loginHandler(rr, req)

// // 		if client.user != nil {
// // 			return fmt.Errorf("Expected user to be nil but in stead got: %v", client.user)
// // 		}

// // 		return nil
// // 	}

// // 	pact.
// // 		AddInteraction().
// // 		Given("User jmarie does not exist").
// // 		UponReceiving("A request to login with user 'jmarie'").
// // 		WithRequest(request{
// // 			Method:  "POST",
// // 			Path:    s("/login/10"),
// // 			Body:    loginRequest,
// // 			Headers: commonHeaders,
// // 			Query: dsl.MapMatcher{
// // 				"foo": s("anything"),
// // 			},
// // 		}).
// // 		WillRespondWith(dsl.Response{
// // 			Status:  404,
// // 			Headers: commonHeaders,
// // 		})

// // 	err := pact.Verify(testJmarieDoesNotExists)
// // 	if err != nil {
// // 		t.Fatalf("Error on Verify: %v", err)
// // 	}
// // }

// // func TestExampleConsumerLoginHandler_UserUnauthorised(t *testing.T) {
// // 	var testJmarieUnauthorized = func() error {
// // 		client := Client{
// // 			Host: fmt.Sprintf("http://localhost:%d", pact.Server.Port),
// // 		}
// // 		client.loginHandler(rr, req)

// // 		if client.user != nil {
// // 			return fmt.Errorf("Expected user to be nil but got: %v", client.user)
// // 		}

// // 		return nil
// // 	}

// // 	pact.
// // 		AddInteraction().
// // 		Given("User jmarie is unauthorized").
// // 		UponReceiving("A request to login with user 'jmarie'").
// // 		WithRequest(request{
// // 			Method:  "POST",
// // 			Path:    s("/login/10"),
// // 			Body:    loginRequest,
// // 			Headers: commonHeaders,
// // 		}).
// // 		WillRespondWith(dsl.Response{
// // 			Status:  401,
// // 			Headers: commonHeaders,
// // 		})

// // 	err := pact.Verify(testJmarieUnauthorized)
// // 	if err != nil {
// // 		t.Fatalf("Error on Verify: %v", err)
// // 	}
// // }

// // func TestExampleConsumerGetUser_Authenticated(t *testing.T) {
// // 	var testJmarieUnauthenticated = func() error {
// // 		client := Client{
// // 			Host:  fmt.Sprintf("http://localhost:%d", pact.Server.Port),
// // 			token: "Bearer 1234",
// // 		}
// // 		client.getUser("10")

// // 		if client.user != nil {
// // 			return fmt.Errorf("Expected user to be nil but got: %v", client.user)
// // 		}

// // 		return nil
// // 	}

// // 	pact.
// // 		AddInteraction().
// // 		Given("User jmarie is authenticated").
// // 		UponReceiving("A request to get user 'jmarie'").
// // 		WithRequest(request{
// // 			Method: "GET",
// // 			Path:   s("/users/10"),
// // 			Headers: dsl.MapMatcher{
// // 				"Authorization": s("Bearer 1234"),
// // 			},
// // 		}).
// // 		WillRespondWith(dsl.Response{
// // 			Status:  200,
// // 			Headers: commonHeaders,
// // 			Body:    dsl.Match(ex.User{}),
// // 		})

// // 	err := pact.Verify(testJmarieUnauthenticated)
// // 	if err != nil {
// // 		t.Fatalf("Error on Verify: %v", err)
// // 	}

// // }
// // func TestExampleConsumerGetUser_Unauthenticated(t *testing.T) {
// // 	var testJmarieUnauthenticated = func() error {
// // 		client := Client{
// // 			Host: fmt.Sprintf("http://localhost:%d", pact.Server.Port),
// // 		}
// // 		client.getUser("10")

// // 		if client.user != nil {
// // 			return fmt.Errorf("Expected user to be nil but got: %v", client.user)
// // 		}

// // 		return nil
// // 	}

// // 	pact.
// // 		AddInteraction().
// // 		Given("User jmarie is unauthenticated").
// // 		UponReceiving("A request to get with user 'jmarie'").
// // 		WithRequest(request{
// // 			Method:  "GET",
// // 			Path:    s("/users/10"),
// // 			Headers: commonHeaders,
// // 		}).
// // 		WillRespondWith(dsl.Response{
// // 			Status:  401,
// // 			Headers: commonHeaders,
// // 		})

// // 	err := pact.Verify(testJmarieUnauthenticated)
// // 	if err != nil {
// // 		t.Fatalf("Error on Verify: %v", err)
// // 	}

// // }
