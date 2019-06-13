package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type LinYuChuang struct {
	Id                int    `orm:"column(id);auto" description:"id"`
	LinYuChuangId     int    `orm:"column(lin_yu_chuang_id)" description:"前端id"`
	IsXiuLian         int    `orm:"column(is_xiu_lian)" description:"是否再修炼"`
	ShangZhenSlot     string `orm:"column(shang_zhen_slot);size(255)" description:"一个存放5个的数组，里面有SlotLock（bool值）和ID"`
	IsLock            int    `orm:"column(is_lock)" description:"是否锁"`
	IsTiaoZhan        int    `orm:"column(is_tiao_zhan)" description:"是否是挑战"`
	IsUsed            int    `orm:"column(is_used)" description:"是否可用"`
	UserId            int    `orm:"column(user_id)" description:"用户id"`
	SlotLock          int    `orm:"column(slot_lock);null"`
	LingYuChuangXwChu int    `orm:"column(ling_yu_chuang_xw_chu);null"`
}

func (t *LinYuChuang) TableName() string {
	return "lin_yu_chuang"
}

func init() {
	orm.RegisterModel(new(LinYuChuang))
}

// AddLinYuChuang insert a new LinYuChuang into database and returns
// last inserted Id on success.
func AddLinYuChuang(m *LinYuChuang) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetLinYuChuangById retrieves LinYuChuang by Id. Returns error if
// Id doesn't exist
func GetLinYuChuangById(id int) (v *LinYuChuang, err error) {
	o := orm.NewOrm()
	v = &LinYuChuang{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllLinYuChuang retrieves all LinYuChuang matches certain condition. Returns empty list if
// no records exist
func GetAllLinYuChuang(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(LinYuChuang))
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

	var l []LinYuChuang
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

// UpdateLinYuChuang updates LinYuChuang by Id and returns error if
// the record to be updated doesn't exist
func UpdateLinYuChuangById(m *LinYuChuang) (err error) {
	o := orm.NewOrm()
	v := LinYuChuang{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteLinYuChuang deletes LinYuChuang by Id and returns error if
// the record to be deleted doesn't exist
func DeleteLinYuChuang(id int) (err error) {
	o := orm.NewOrm()
	v := LinYuChuang{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&LinYuChuang{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
