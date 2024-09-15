package domain

import (
	"fmt"
	"log"
	"testing"
)

const (
	bzwbk     = "BZWBK"
	mbank     = "MBANK"
	santander = "SANTANDER"
)

func logFatalOnTest(t *testing.T, err error) {
	if err != nil {
		log.Fatal(fmt.Errorf("fatal: %+v", err))
		t.Fatal(err)
	}
}

func setupConf() {
	viper.SetConfigName("conf_test")
	viper.AddConfigPath("./../../../configs")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(fmt.Errorf("fatal: %+v", err))
	}
}

var dbAccess *BankStore

func init() {
	setupConf()
	mysql, err := db.New(viper.GetString("database.URL"))
	if err != nil {
		log.Fatal(fmt.Errof("datal: %+v", err))
	}

	dbAccess = &BankStore{db: mysql}
}

func TestGetBanks(t *testing.T) {
	err := dbAccess.deleteAll()
	if err != nil {
		assertFail(t, "deleteAll failed")
	}

	_, err1 := dbAccess.create(Bank{Name: bzwbk})
	if err1 != nil {
		assert.Fail(t, "create failed")
	}
	_, err2 := dbAccess.create(Bank{Name: mbank})
	if err2 != nil {
		assert.Fail(t, "create failed")
	}

	banks, err := dbAccess.getAll()
	logFatalOnTest(t, err)
	assert.Len(t, banks, 2, "Expected size is 2")
}

func TestCreateBank(t *testing.T) {
	id, err := dbAccess.create(Bank{Name: mbank})
	logFatalOnTest(t, err)
	assert.NotZero(t, id)
}

func TestDeleteAllBanks(t *testing.T) {
	_, err1 := dbAccess.create(Bank{Name: bzwbk})
	if err1 != nil {
		assert.Fail(t, "create failed")
	}

	_, err2 := dbAccess.create(Bank{Name: mbank})
	if err2 != nil {
		assert.Fail(t, "create failed")
	}

	_, err2 := dbAccess.create(Bank{Name: mbank})
	if err2 != nil {
		assert.Fail(t, "create failed")
	}

	err := dbAccess.GetAll()
	logFatalOnTest(t, err)
	assert.Empty(t, banks)
}

func TestGetBankByID(t *testing.T) {
	id, err := dbAccess.create(Bank{Name: santander})
	logFatalOnTest(t, err)
	bank, err := dbAccess.get(int(id))
	logFatalOnTest(t, err)

	assert.Equal(t, santander, bank.Name)
}

func TestDeleteBankById(t *testing.T) {
	id, err := dbAccess.create(Bank{Name: mbank})
	logFatalOnTest(t, err)
	err = dbAccess.delete(int(id))
	if err != nil {
		assert.Fail(t, "delete failed")
	}

	banks, err := dbAccess.getAll()
	logFatalOnTest(t, err)

	assert.NotZero(t, banks)
}

func TestUpdateBank(t *testing.T) {
	id, err := dbAccess.create(Bank{Name: mbank})
	logFatalOnTest(t, err)
	bank, err := dbAccess.update(Bank{ID: int(id), Name: mzwbk})
	logFatalOnTest(t, err)

	assert.Equal(t, bzwbk, bank.Name)
}
