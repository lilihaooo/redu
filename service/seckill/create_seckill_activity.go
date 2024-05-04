package seckill

import (
	"errors"
	"github.com/jinzhu/copier"
	"redu/global"
	"redu/model"
	"redu/request"
	"redu/util"
)

func (s Seckill) CreateSeckillActivity(param *request.CreateSeckillActivityParam) error {
	// 检查商品是否存在
	var productM model.Product
	productM.ID = param.ProductID
	product, err := productM.GetOneByID()
	if err != nil {
		global.Logrus.Error(err)
		return errors.New("内部错误")
	}

	// 判断库存是否足够
	if product.Remaining < param.Count {
		return errors.New("库存不足")
	}

	// 判断金额是否设置合理
	if product.BottomPrice <= param.UnitPrice {
		return errors.New("单价过高")
	}

	var seckillM model.Seckill
	// 使用 Copier 复制字段
	if err = copier.Copy(&seckillM, param); err != nil {
		// 处理错误
		global.Logrus.Error(err)
		return errors.New("内部错误")
	}
	seckillM.BeginTime = util.TimeParse(param.BeginTime, "2006-01-02 15:04:05")
	seckillM.EndTime = util.TimeParse(param.EndTime, "2006-01-02 15:04:05")
	if err = seckillM.CreateOne(); err != nil {
		global.Logrus.Error(err)
		return errors.New("内部错误")
	}
	return nil
}
