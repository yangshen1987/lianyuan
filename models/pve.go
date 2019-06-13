package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Pve struct {
	Id              int    `orm:"column(id);auto" description:"自增id"`
	MapId           int    `orm:"column(map_id);null" description:"地图id"`
	IsPassed        int    `orm:"column(is_passed);null" description:"是否通关"`
	MapInfo         string `orm:"column(map_info);null" description:"地图信息"`
	UserId          int    `orm:"column(user_id);null" description:"用户id"`
	IsProduct       int    `orm:"column(is_product);null" description:"是否在生产"`
	NewProductCount int    `orm:"column(new_product_count);null" description:"新生产"`
	ConsumePerTime  int    `orm:"column(consume_per_time);null" description:"已消耗的单位时间（比如玩家在45分钟的时候下线，则该字段值就为45）"`
	ProductPer      int    `orm:"column(ProductPer);null" description:"每小时最少产出"`
	InsertTime      int    `orm:"column(insert_time);null" description:"进入次数"`
}

func (t *Pve) TableName() string {
	return "pve"
}

func init() {
	orm.RegisterModel(new(Pve))
}

// AddPve insert a new Pve into database and returns
// last inserted Id on success.
func AddPve(m *Pve) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetPveById retrieves Pve by Id. Returns error if
// Id doesn't exist
func GetPveById(id int) (v *Pve, err error) {
	o := orm.NewOrm()
	v = &Pve{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllPve retrieves all Pve matches certain condition. Returns empty list if
// no records exist
func GetAllPve(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Pve))
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

	var l []Pve
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

// UpdatePve updates Pve by Id and returns error if
// the record to be updated doesn't exist
func UpdatePveById(m *Pve) (err error) {
	o := orm.NewOrm()
	v := Pve{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeletePve deletes Pve by Id and returns error if
// the record to be deleted doesn't exist
func DeletePve(id int) (err error) {
	o := orm.NewOrm()
	v := Pve{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Pve{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
