package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Recharge struct {
	Type        int `orm:"column(Type);null" description:"充值类型 （分为豪华首冲，至尊首冲，特惠礼包）"`
	Qualify     int `orm:"column(Qualify);null" description:"是否达到领取礼包的资格  1： 达到； 0：未达到"`
	Remain      int `orm:"column(Remain);null" description:"是否还未领取礼 1：还未领取礼包；0：已经领取礼包"`
	LeftTime    int `orm:"column(LeftTime);null" description:"剩余购买次数"`
	Version     int `orm:"column(Version);null" description:"档位（例如648礼包分为四挡，每次刷新限时活动随机从四挡中取一档，具体可以查看活动模块）"`
	RefreshTime int `orm:"column(RefreshTime);null" description:"刷新时间，每隔72小时刷新一次"`
	AlreadyBuy  int `orm:"column(AlreadyBuy);null" description:"是否购买礼包  1：是； 0：否"`
	Id          int `orm:"column(id);auto"`
	UserId      int `orm:"column(user_id)"`
	ChongZhi    int `orm:"column(ChongZhi)"`
}

func (t *Recharge) TableName() string {
	return "recharge"
}

func init() {
	orm.RegisterModel(new(Recharge))
}

// AddRecharge insert a new Recharge into database and returns
// last inserted Id on success.
func AddRecharge(m *Recharge) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetRechargeById retrieves Recharge by Id. Returns error if
// Id doesn't exist
func GetRechargeById(id int) (v *Recharge, err error) {
	o := orm.NewOrm()
	v = &Recharge{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllRecharge retrieves all Recharge matches certain condition. Returns empty list if
// no records exist
func GetAllRecharge(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Recharge))
	// query k=v
	for k, v := range query {
		// rewrite dot-notation to Object__Attribute
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		} else {
			qs = qs.Filter(k, v)
		}
	}
	// order by:
	var sortFields []string
	if len(sortby) != 0 {
		if len(sortby) == len(order) {
			// 1) for each sort field, there is an associated order
			for i, v := range sortby {
				orderby := ""
				if order[i] == "desc" {
					orderby = "-" + v
				} else if order[i] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
			qs = qs.OrderBy(sortFields...)
		} else if len(sortby) != len(order) && len(order) == 1 {
			// 2) there is exactly one order, all the sorted fields will be sorted by this order
			for _, v := range sortby {
				orderby := ""
				if order[0] == "desc" {
					orderby = "-" + v
				} else if order[0] == "asc" {
					orderby = v
				} else {
					return nil, errors.New("Error: Invalid order. Must be either [asc|desc]")
				}
				sortFields = append(sortFields, orderby)
			}
		} else if len(sortby) != len(order) && len(order) != 1 {
			return nil, errors.New("Error: 'sortby', 'order' sizes mismatch or 'order' size is not 1")
		}
	} else {
		if len(order) != 0 {
			return nil, errors.New("Error: unused 'order' fields")
		}
	}

	var l []Recharge
	qs = qs.OrderBy(sortFields...)
	if _, err = qs.Limit(limit, offset).All(&l, fields...); err == nil {
		if len(fields) == 0 {
			for _, v := range l {
				ml = append(ml, v)
			}
		} else {
			// trim unused fields
			for _, v := range l {
				m := make(map[string]interface{})
				val := reflect.ValueOf(v)
				for _, fname := range fields {
					m[fname] = val.FieldByName(fname).Interface()
				}
				ml = append(ml, m)
			}
		}
		return ml, nil
	}
	return nil, err
}

// UpdateRecharge updates Recharge by Id and returns error if
// the record to be updated doesn't exist
func UpdateRechargeById(m *Recharge) (err error) {
	o := orm.NewOrm()
	v := Recharge{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteRecharge deletes Recharge by Id and returns error if
// the record to be deleted doesn't exist
func DeleteRecharge(id int) (err error) {
	o := orm.NewOrm()
	v := Recharge{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Recharge{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
