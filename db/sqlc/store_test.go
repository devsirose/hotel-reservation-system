package db

import (
	"testing"
)

func TestExecTx(t *testing.T) {
	//conn, _ := sql.Open(dbDriver, dbSource)
	//store := db.NewStore(conn)
	//
	//ctx := context.Background()
	////lambda function giong nhu ban huong dan va sau do thuc thi boi fn(params)
	//// lambda function nay can tham so la *db.Queries de thuc thi cac cau query
	//err := store.execTx(ctx, func(q *db.Queries) error {
	//	acc1, err := q.CreateAccount(ctx, db.CreateAccountParams{
	//		Owner:    "sirose",
	//		Balance:  0,
	//		Currency: "EUR",
	//	})
	//	if err != nil {
	//		return err
	//	}
	//	acc1, err = q.AddAccountBalance(ctx, db.AddAccountBalanceParams{
	//		Amount: 100,
	//		ID:     acc1.ID,
	//	})
	//	return nil
	//})
	//require.NoError(t, err)
}
