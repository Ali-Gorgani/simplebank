package api

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/Ali-Gorgani/simplebank/db/mock"
	db "github.com/Ali-Gorgani/simplebank/db/sqlc"
	"github.com/Ali-Gorgani/simplebank/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

// func TestCreateAccountAPI(t *testing.T) {
// 	server := newTestServer(t)
// 	// Add account
// 	// Get account
// 	// Delete account
// }

func TestGetAccountAPI(t *testing.T) {
	account := randomAccount()
	
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	// Build stub
	store.EXPECT().
		GetAccount(gomock.Any(), gomock.Eq(account.ID)).
		Times(1).
		Return(account, nil)

	// Start test server and send request
	server := NewServer(store)
	recorder := httptest.NewRecorder()

	url := fmt.Sprintf("/accounts/%d", account.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)
	// Check response
	require.Equal(t, http.StatusOK, recorder.Code)
}

// func TestListAccountsAPI(t *testing.T) {
// 	server := newTestServer(t)
// 	// Add 3 accounts
// 	// List accounts
// 	// Check that accounts have been created
// }

// func TestUpdateAccountAPI(t *testing.T) {
// 	server := newTestServer(t)
// 	// Add account
// 	// Update account
// 	// Check that account has been updated
// }

// func TestDeleteAccountAPI(t *testing.T) {
// 	server := newTestServer(t)
// 	// Add account
// 	// Delete account
// 	// Check that account has been deleted
// }

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomInt(1, 1000),
		Owner:    util.RandomOwner(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}


