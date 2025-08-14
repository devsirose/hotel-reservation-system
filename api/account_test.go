package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	mockdb "github.com/devsirose/simplebank/db/mock"
	db "github.com/devsirose/simplebank/db/sqlc"
	"github.com/devsirose/simplebank/util"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestGetAccount(t *testing.T) {
	account := randomAccount()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	store := mockdb.NewMockStore(ctrl)
	//Given
	store.EXPECT().GetAccount(gomock.Any(), gomock.Eq(account.ID)).Return(account, nil).Times(1)
	//When action occurs in service or bussiness logic
	//Example test result of function (GetAccountBy in api) whether match account
	server := NewServer(store)
	recorder := httptest.NewRecorder()

	// When: gọi API GET /accounts/:id
	url := fmt.Sprintf("/api/v1/accounts/%d", account.ID)
	request, err := http.NewRequest(http.MethodGet, url, nil)
	require.NoError(t, err)

	server.router.ServeHTTP(recorder, request)

	// Then: HTTP 200 và body khớp account
	require.Equal(t, http.StatusOK, recorder.Code)

	var gotAccount db.Account
	err = json.Unmarshal(recorder.Body.Bytes(), &gotAccount)
	require.NoError(t, err)
	require.Equal(t, account, gotAccount)

}

func randomAccount() db.Account {
	return db.Account{
		ID:       util.RandomID(),
		Owner:    util.RandomName(),
		Balance:  util.RandomMoney(),
		Currency: util.RandomCurrency(),
	}
}
