package dao

import "rbac/db"

type MenuDao struct {

}

func (MenuDao) List() {
	db.Db.Find()
}