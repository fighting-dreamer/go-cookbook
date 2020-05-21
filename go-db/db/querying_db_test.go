package db

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"errors"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"nipun.io/go-cookbook/go-db/domain"
	"testing"

	"github.com/stretchr/testify/suite"
)

type QueryingDBTestSuite struct {
	suite.Suite
	mock sqlmock.Sqlmock
	dl *DAO_Layer
	context context.Context

}

type CustomConverter struct{}

func (s CustomConverter) ConvertValue(v interface{}) (driver.Value, error) {
	switch v.(type) {
	case string:
		return v.(string), nil
	case []string:
		//return v.([]string), nil
		return v.(pq.StringArray), nil
	case int:
		return v.(int), nil
	default:
		return nil, errors.New(fmt.Sprintf("cannot convert %T with value %v", v, v))
	}
}


func mockConnect() (*sql.DB, sqlmock.Sqlmock) {
	//db, mock, err := sqlmock.New(sqlmock.ValueConverterOption(CustomConverter{}))
	db, mock, err := sqlmock.New()

	if err != nil {
		fmt.Println(err)
	}
	return db, mock
}

func mockClose(conn *sqlx.DB) {
	conn.Close()
}

func mockPQArray(a interface{}) interface { Scan(interface {}) error; Value() (driver.Value, error) } {
//interface { Scan(interface {}) error; Value() (driver.Value, error) }
	switch a := a.(type) {
		case []string:
			return (*pq.StringArray)(&a)
	}
	return nil
}

func (suite *QueryingDBTestSuite) SetupSuite() {
	suite.context = context.Background()
	db, mock := mockConnect()
	suite.mock = mock
	suite.dl = &DAO_Layer{
		Conn : sqlx.NewDb(db, "sqlmock"),
	}
}

func (suite *QueryingDBTestSuite) SetupTest() {
	//suite.orderRepository = NewOrderRepository()
}

func TestQueryingDBTestSuite(t *testing.T) {
	suite.Run(t, new(QueryingDBTestSuite))
}

func (suite *QueryingDBTestSuite) TestSingleLevelDataQuery() {
	dl := suite.dl
	mock := suite.mock
	expected := []domain.DBUser{
		{
			Id:       1,
			Name:     "user-1",
		},
		{
			Id:       2,
			Name:     "user-2",
		},
		{
			Id:       3,
			Name:     "user-3",
		},
		{
			Id:       4,
			Name:     "user-4",
		},
	};
	rows := sqlmock.NewRows([]string{"id", "name"}).
					AddRow(1, "user-1").
					AddRow(2, "user-2").
					AddRow(3, "user-3").
					AddRow(4, "user-4")
	mock.ExpectQuery("^SELECT id, name FROM db_user").WillReturnRows(rows)
	output := dl.SingleLevelDataQuery()
	suite.Equal(expected, output)
}

func (suite *QueryingDBTestSuite) TestSingleLevelWithArrayDataQuery() {
	dl := suite.dl
	mock := suite.mock

	expected := []domain.DBUser{
		//{
		//	Id:       1,
		//	Name:     "user-1",
		//	Contacts: []string{"1234567890","2345678901"},
		//},
		//{
		//	Id:       2,
		//	Name:     "user-2",
		//	Contacts: []string{"3456789012","4567890123"},
		//},
		//{
		//	Id:       3,
		//	Name:     "user-3",
		//	Contacts: []string{"1111111111","0000000000"},
		//},
		{
			Id:       4,
			Name:     "user-4",
			Contacts: domain.PqStringArray([]string{"4444444444","5555555555"}),
		},
	};

	// NOTE : Array need to be stored in native format of the driver using, here we are using postgres, specifically the function pq.Array(), due to that!!!
	rows := sqlmock.NewRows([]string{"id", "name", "contacts"}).
		//AddRow(1, "user-1", []string{"1234567890","2345678901"},).
		//AddRow(2, "user-2", []string{"3456789012","4567890123"}).
		//AddRow(3, "user-3", []string{"1111111111","0000000000"}).
		AddRow(4, "user-4", domain.PqStringArray([]string{"4444444444","5555555555"}))
	mock.ExpectQuery("SELECT id, name, contacts FROM db_user;").WillReturnRows(rows)
	output := dl.SingleLevelWithArrayDataQuery()
	suite.Equal(expected, output)
}

func (suite *QueryingDBTestSuite) TestSingleLevelWithArrayDataInsertQuery() {
	dl := suite.dl
	mock := suite.mock

	expectedRowsEffected := int64(1)


	mock.ExpectExec("INSERT INTO db_user").WillReturnResult(sqlmock.NewResult(-1, 1))
	outputRowsEffected, err := dl.SingleLevelWithArrayDataInsertQuery().RowsAffected()
	suite.NoError(err)
	suite.Nil(err)
	suite.Equal(expectedRowsEffected, outputRowsEffected)
}

func (suite *QueryingDBTestSuite) TestSingleLevelWithArrayDataUpdateQuery() {
	dl := suite.dl
	mock := suite.mock

	expectedRowsEffected := int64(1)


	mock.ExpectExec("UPDATE db_user").WillReturnResult(sqlmock.NewResult(-1, 1))
	outputRowsEffected, err := dl.SingleLevelWithArrayDataUpdateQuery().RowsAffected()
	suite.NoError(err)
	suite.Nil(err)
	suite.Equal(expectedRowsEffected, outputRowsEffected)
}

func (suite *QueryingDBTestSuite) TestSingleLevelWithArrayDataUpdateQueryNoUPDATE() {
	dl := suite.dl
	mock := suite.mock

	expectedRowsEffected := int64(0)


	mock.ExpectExec("UPDATE db_user").WillReturnResult(sqlmock.NewResult(-1, 0))
	outputRowsEffected, err := dl.SingleLevelWithArrayDataUpdateQueryNoUPDATE().RowsAffected()
	suite.NoError(err)
	suite.Nil(err)
	suite.Equal(expectedRowsEffected, outputRowsEffected)
}

func (suite *QueryingDBTestSuite) TestSingleLevelWithArrayDataUpdateQueryGetError() {
	dl := suite.dl
	mock := suite.mock

	expectedRowsEffected := int64(0)


	mock.ExpectExec("UPDATE db_user").WillReturnResult(sqlmock.NewResult(-1, 0))
	outputRowsEffected, err := dl.SingleLevelWithArrayDataUpdateQueryNoUPDATE().RowsAffected()
	suite.NoError(err)
	suite.Nil(err)
	suite.Equal(expectedRowsEffected, outputRowsEffected)
}

//func (suite *QueryingDBTestSuite) TestUpdateInTransaction() {
//	dl := suite.dl
//	mock := suite.mock
//
//	mock.ExpectBegin()
//	mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
//	mock.ExpectCommit()
//	//mock.ExpectExec("INSERT INTO product_viewers").
//	//	WithArgs(2, 3).
//	//	WillReturnError(fmt.Errorf("some error"))
//
//	//mock.ExpectRollback()
//
//	mock.ExpectationsWereMet() != nil
//	expected := nil
//
//}