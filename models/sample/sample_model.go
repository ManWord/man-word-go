package sample

import "github.com/astaxie/beego/orm"

type SampleModel struct {
	Id      int
	Content string
}

func (sample *SampleModel) TableName() string {
	return "sample"
}

func init() {
	orm.RegisterModel(new(SampleModel))
}
