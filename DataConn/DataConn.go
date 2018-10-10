package DataConn

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

//用户
type User struct {
	Id        int     `gorm:"auto_increment"` //会员编号
	Number    string  `gorm:"not null"`       //账号
	Password  string  `gorm:"not null"`       //密码
	Username  string  `gorm:"not null"`       //用户名
	Tel       string  `gorm:"not null"`       //电话
	Address   string  `gorm:"not null"`       //地址
	Grade     string  `gorm:"default:'普通用户'"` //等级
	Totalcost float64 `gorm:"not null"`       //累计消费
}

//产品类别
type Category struct {
	Id    int    `gorm:"auto_increment"`
	Name  string `gorm:"not null"` //类别名称
	Descr string `gorm:"not null"` //类别描述
}

//产品
type Pruduct struct {
	Id           int       `gorm:"auto_increment"`
	Name         string    `gorm:"not null"`        //产品名
	Descr        string    `gorm:"not null"`        //产品描述
	Normalprice  string    `gorm:"not null"`        //普通价
	Memberprice  string    `gorm:"not null"`        //会员价
	Uppercabinet string    `gorm:"default:'false'"` //是否上柜
	Pdate        time.Time //上柜时间
	Category     string    `gorm:"not null"` //类别名
}

//订单情况
type Salesorder struct {
	Id          int       `gorm:"auto_increment"` //订单编号
	Username    string    `gorm:"not null"`       //用户名
	Pruductid   string    `gorm:"not null"`       //产品id
	Pruductname string    `gorm:"not null"`       //产品名称
	Unitprice   string    `gorm:"not null"`       //单价
	Pcount      int       `gorm:"not null"`       //数量
	Totalprice  string    `gorm:"not null"`       //总价
	Address     string    `gorm:"not null"`       //送货地址
	Ordertime   time.Time `gorm:"not null"`       //下单时间
	Status      int       `gorm:"not null"`       //订单状态
}
