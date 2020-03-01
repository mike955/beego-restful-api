package models

import (
	"github.com/astaxie/beego/orm"
	// "fmt"
)

type Account struct {
	Id int64 `json:"id" orm:"column(id);auto"`
	AccountName string `json:"accountName" orm:"column(account_name)"`
	Password string `json:"password"; orm:"column(password)"`
	Email string `json:"email" orm:"column(email)"`
}

func init() {
	orm.RegisterModel(new(Account))
}

func NewAccount() *Account {
	return &Account{}
}

func AddAccount(t *Account) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(t)
	return id, err
}