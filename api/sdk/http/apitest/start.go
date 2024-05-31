package apitest

import (
	"net/http/httptest"
	"testing"

	authbuild "github.com/tiberzus/goservice/api/cmd/services/auth/build/all"
	salesbuild "github.com/tiberzus/goservice/api/cmd/services/sales/build/all"
	"github.com/tiberzus/goservice/api/sdk/http/mux"
	"github.com/tiberzus/goservice/app/sdk/auth"
	"github.com/tiberzus/goservice/app/sdk/authclient"
	"github.com/tiberzus/goservice/business/sdk/dbtest"
)

// StartTest initialized the system to run a test.
func StartTest(t *testing.T, testName string) *Test {
	db := dbtest.NewDatabase(t, testName)

	// -------------------------------------------------------------------------

	auth, err := auth.New(auth.Config{
		Log:       db.Log,
		DB:        db.DB,
		KeyLookup: &KeyStore{},
	})
	if err != nil {
		t.Fatal(err)
	}

	// -------------------------------------------------------------------------

	server := httptest.NewServer(mux.WebAPI(mux.Config{
		Log:  db.Log,
		Auth: auth,
		DB:   db.DB,
	}, authbuild.Routes()))

	authClient := authclient.New(db.Log, server.URL)

	// -------------------------------------------------------------------------

	mux := mux.WebAPI(mux.Config{
		Log:        db.Log,
		AuthClient: authClient,
		DB:         db.DB,
	}, salesbuild.Routes())

	return New(db, auth, mux)
}
