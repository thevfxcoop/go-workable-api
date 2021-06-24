package client_test

import (
	"net/url"
	"os"
	"testing"

	// Modules
	workableclient "github.com/thevfxcoop/go-workable-api/pkg/client"
	schema "github.com/thevfxcoop/go-workable-api/pkg/schema"
)

func Test_Client_000(t *testing.T) {
	t.Log(t.Name())
}
func Test_Client_001(t *testing.T) {
	client := GetClient(t, false)
	t.Log(client)
}

func Test_Client_002(t *testing.T) {
	client := GetClient(t, true)
	if accounts, err := client.GetAccounts(); err != nil {
		t.Error(err)
	} else if len(accounts) == 0 {
		t.Error("No accounts returned")
	} else {
		for _, account := range accounts {
			t.Logf("account=%+v", account)
		}
	}
}

func Test_Client_003(t *testing.T) {
	client := GetClient(t, true)
	if jobs, err := client.GetJobs(); err != nil {
		t.Error(err)
	} else {
		for _, job := range jobs {
			t.Logf("job=%+v", job)
		}
	}
}

func Test_Client_004(t *testing.T) {
	client := GetClient(t, false)
	if jobs, err := client.GetJobs(); err != nil {
		t.Error(err)
	} else {
		for _, job := range jobs {
			if job, err := client.GetJobWithShortcode(job.Shortcode); err != nil {
				t.Error(err)
			} else {
				t.Logf("job=%+v", job)
			}
		}
	}
}

func Test_Client_005(t *testing.T) {
	client := GetClient(t, false)
	if candidates, err := client.GetCandidates(); err != nil {
		t.Error(err)
	} else {
		for _, candidate := range candidates {
			if candidate, err := client.GetCandidate(candidate.Id); err != nil {
				t.Error(err)
			} else {
				t.Logf("candidate=%+v", candidate)
			}
		}
	}
}

func Test_Client_006(t *testing.T) {
	client := GetClient(t, true)
	candidate := &schema.Candidate{}
	if shortcode := os.Getenv("WORKABLE_SHORTCODE"); shortcode == "" {
		t.Skip("Skipping test, set WORKABLE_SHORTCODE environment variable otherwise")
	} else if candidate_, err := client.CreateCandidate(shortcode, candidate); err != nil {
		t.Error(err)
	} else {
		t.Log("Created=", candidate_)
	}
}

///////////////////////////////////////////////////////////////////////////////

// To run some tests, use environment variables WORKABLE_ENDPOINT and
// WORKABLE_APIKEY

func GetClient(t *testing.T, verbose bool) *workableclient.Client {
	if endpoint := os.Getenv("WORKABLE_ENDPOINT"); endpoint == "" {
		t.Skip("Skipping test, set WORKABLE_ENDPOINT environment variable otherwise")
	} else if endpoint, err := url.Parse(endpoint); err != nil {
		t.Fatal(err)
	} else if apikey := os.Getenv("WORKABLE_APIKEY"); apikey == "" {
		t.Skip("Skipping test, set WORKABLE_APIKEY environment variable otherwise")
	} else if client, err := workableclient.NewClient(endpoint, workableclient.OptRateLimit(0.5), workableclient.OptTrace(os.Stderr, verbose), workableclient.OptAPIKey(apikey)); err != nil {
		t.Fatal(err)
	} else {
		return client
	}

	// Skip tests
	return nil
}
