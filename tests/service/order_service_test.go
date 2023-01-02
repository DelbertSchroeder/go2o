package service

import (
	"context"
	"strings"
	"testing"

	"github.com/ixre/go2o/core/service/impl"
	"github.com/ixre/go2o/core/service/proto"
)

func TestSubmitNormalOrder(t *testing.T) {
	var memberId int64 = 1
	_, err := impl.OrderService.SubmitNormalOrder_(
		context.TODO(),
		&proto.SubmitNormalOrderV2Request{
			BuyerId:         memberId,
			AddressId:       1,
			Subject:         "",
			CouponCode:      "",
			BalanceDiscount: false,
		})
	if err != nil {
		t.Error(err)
	}
}

func TestSubmitOrderSubjectPostgresInsert(t *testing.T) {
	s := "指间陶艺 精品宜兴紫砂茶宠 名家全手工小号座镇貔貅茶玩雕塑摆件精品 威震八方(貔貅) 公款"
	s2 := strings.Replace(s, " ", "", -1)
	s3 := s2[:15] + "..."
	t.Log("----", s3)
}