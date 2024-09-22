// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dal

import (
	"context"
	"fmt"
	"testing"

	"GoLib/GORM/model"

	"gorm.io/gen"
	"gorm.io/gen/field"
	"gorm.io/gorm/clause"
)

func init() {
	InitializeDB()
	err := _gen_test_db.AutoMigrate(&model.Dep{})
	if err != nil {
		fmt.Printf("Error: AutoMigrate(&model.Dep{}) fail: %s", err)
	}
}

func Test_depQuery(t *testing.T) {
	dep := newDep(_gen_test_db)
	dep = *dep.As(dep.TableName())
	_do := dep.WithContext(context.Background()).Debug()

	primaryKey := field.NewString(dep.TableName(), clause.PrimaryKey)
	_, err := _do.Unscoped().Where(primaryKey.IsNotNull()).Delete()
	if err != nil {
		t.Error("clean table <deps> fail:", err)
		return
	}

	_, ok := dep.GetFieldByName("")
	if ok {
		t.Error("GetFieldByName(\"\") from dep success")
	}

	err = _do.Create(&model.Dep{})
	if err != nil {
		t.Error("create item in table <deps> fail:", err)
	}

	err = _do.Save(&model.Dep{})
	if err != nil {
		t.Error("create item in table <deps> fail:", err)
	}

	err = _do.CreateInBatches([]*model.Dep{{}, {}}, 10)
	if err != nil {
		t.Error("create item in table <deps> fail:", err)
	}

	_, err = _do.Select(dep.ALL).Take()
	if err != nil {
		t.Error("Take() on table <deps> fail:", err)
	}

	_, err = _do.First()
	if err != nil {
		t.Error("First() on table <deps> fail:", err)
	}

	_, err = _do.Last()
	if err != nil {
		t.Error("First() on table <deps> fail:", err)
	}

	_, err = _do.Where(primaryKey.IsNotNull()).FindInBatch(10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatch() on table <deps> fail:", err)
	}

	err = _do.Where(primaryKey.IsNotNull()).FindInBatches(&[]*model.Dep{}, 10, func(tx gen.Dao, batch int) error { return nil })
	if err != nil {
		t.Error("FindInBatches() on table <deps> fail:", err)
	}

	_, err = _do.Select(dep.ALL).Where(primaryKey.IsNotNull()).Order(primaryKey.Desc()).Find()
	if err != nil {
		t.Error("Find() on table <deps> fail:", err)
	}

	_, err = _do.Distinct(primaryKey).Take()
	if err != nil {
		t.Error("select Distinct() on table <deps> fail:", err)
	}

	_, err = _do.Select(dep.ALL).Omit(primaryKey).Take()
	if err != nil {
		t.Error("Omit() on table <deps> fail:", err)
	}

	_, err = _do.Group(primaryKey).Find()
	if err != nil {
		t.Error("Group() on table <deps> fail:", err)
	}

	_, err = _do.Scopes(func(dao gen.Dao) gen.Dao { return dao.Where(primaryKey.IsNotNull()) }).Find()
	if err != nil {
		t.Error("Scopes() on table <deps> fail:", err)
	}

	_, _, err = _do.FindByPage(0, 1)
	if err != nil {
		t.Error("FindByPage() on table <deps> fail:", err)
	}

	_, err = _do.ScanByPage(&model.Dep{}, 0, 1)
	if err != nil {
		t.Error("ScanByPage() on table <deps> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrInit()
	if err != nil {
		t.Error("FirstOrInit() on table <deps> fail:", err)
	}

	_, err = _do.Attrs(primaryKey).Assign(primaryKey).FirstOrCreate()
	if err != nil {
		t.Error("FirstOrCreate() on table <deps> fail:", err)
	}

	var _a _another
	var _aPK = field.NewString(_a.TableName(), "id")

	err = _do.Join(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("Join() on table <deps> fail:", err)
	}

	err = _do.LeftJoin(&_a, primaryKey.EqCol(_aPK)).Scan(map[string]interface{}{})
	if err != nil {
		t.Error("LeftJoin() on table <deps> fail:", err)
	}

	_, err = _do.Not().Or().Clauses().Take()
	if err != nil {
		t.Error("Not/Or/Clauses on table <deps> fail:", err)
	}
}