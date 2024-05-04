package model

import (
	"redu/global"
	"redu/model/common"
	"time"
)

type Seckill struct {
	BaseModel
	ProductID      uint `json:"product_id"`
	Product        Product
	Title          string               `json:"title"`
	BeginTime      time.Time            `json:"begin_time"`
	EndTime        time.Time            `json:"end_time"`
	Count          int                  `json:"count"`
	UnitPrice      float64              `json:"unit_price"`
	AllowUserGrade int                  `json:"allow_user_grade"`
	Status         common.SeckillStatus `json:"status"`
}

func (s *Seckill) CreateOne() error {
	return global.DB.Create(s).Error
}
