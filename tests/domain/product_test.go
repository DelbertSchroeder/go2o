package domain

import (
	"encoding/json"
	"testing"

	"github.com/ixre/go2o/core/domain/interface/product"
	"github.com/ixre/go2o/tests/ti"
	"github.com/ixre/gof/types/typeconv"
)

func TestSaveProduct(t *testing.T) {
	var productId int64 = 1
	repo := ti.Factory.GetProductRepo()
	prod := repo.GetProduct(productId)
	if prod == nil {
		t.Error(product.ErrNoSuchProduct)
		t.FailNow()
	}
	value := prod.GetValue()
	value.Attrs = prod.Attr()
	t.Log(typeconv.MustJson(prod.Attr()))

	json.Unmarshal([]byte(`[{"Id":17,"ProductId":1,"AttrName":"","AttrId":6,"AttrData":"17","AttrWord":"电脑"},{"Id":16,"ProductId":1,"AttrName":"","AttrId":7,"AttrData":"20,21","AttrWord":"Type-C,标准USB"}]`),
		&value.Attrs)

	err := prod.SetAttr(value.Attrs)
	if err == nil {
		_, err = prod.Save()
	}
	if err != nil {
		t.Error(err)
		t.FailNow()
	}
	t.Log(typeconv.MustJson(prod.Attr()))
}