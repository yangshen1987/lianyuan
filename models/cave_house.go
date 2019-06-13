package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type CaveHouse struct {
	Id                            int `orm:"column(id);auto" description:"自增id"`
	UserId                        int `orm:"column(user_id)" description:"用户id"`
	HouseLeve                     int `orm:"column(house_leve)" description:"洞府等级"`
	SlaveServantUse               int `orm:"column(slave_servant_use);null" description:"奴仆使用数量"`
	SilverOreLevel                int `orm:"column(silver_ore_level);null" description:"银矿"`
	SilverOreSlaveServantUse      int `orm:"column(silver_ore_slave_servant_use);null"`
	LingTianLevel                 int `orm:"column(ling_tian_level);null" description:"灵田"`
	LingTianSlaveServantUse       int `orm:"column(ling_tian_slave_servant_use);null"`
	WoodYardLevel                 int `orm:"column(wood_yard_level);null" description:"木场"`
	WoodYardSlaveServantUse       int `orm:"column(wood_yard_ slave_servant_use);null"`
	MedicinalFieldLevel           int `orm:"column(medicinal_field_level);null" description:"药田"`
	MedicinalFieldSlaveServantUse int `orm:"column(medicinal_field_slave_servant_use);null"`
	IronOreLevel                  int `orm:"column(iron_ore_level);null" description:"铁矿"`
	IronOreLevelSlaveServantUse   int `orm:"column(iron_ore_level_ slave_servant_use);null"`
	PracticeDanRoomLevel          int `orm:"column(practice_dan_room_level);null" description:"练丹房"`
	LianQiGeLevel                 int `orm:"column(lian_qi_ge_level);null" description:"炼器阁等级"`
	IsDouble                      int `orm:"column(is_double);null" description:"是否双倍"`
	RemainDoubleTime              int `orm:"column(remain_double_time);null"`
	SilverOreChu                  int `orm:"column(silver_ore_chu);null" description:"银矿处"`
}

func (t *CaveHouse) TableName() string {
	return "cave_house"
}

func init() {
	orm.RegisterModel(new(CaveHouse))
}

// AddCaveHouse insert a new CaveHouse into database and returns
// last inserted Id on success.
func AddCaveHouse(m *CaveHouse) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetCaveHouseById retrieves CaveHouse by Id. Returns error if
// Id doesn't exist
func GetCaveHouseById(id int) (v *CaveHouse, err error) {
	o := orm.NewOrm()
	v = &CaveHouse{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllCaveHouse retrieves all CaveHouse matches certain condition. Returns empty list if
// no records exist
func GetAllCaveHouse(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(CaveHouse))
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

	var l []CaveHouse
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

// UpdateCaveHouse updates CaveHouse by Id and returns error if
// the record to be updated doesn't exist
func UpdateCaveHouseById(m *CaveHouse) (err error) {
	o := orm.NewOrm()
	v := CaveHouse{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteCaveHouse deletes CaveHouse by Id and returns error if
// the record to be deleted doesn't exist
func DeleteCaveHouse(id int) (err error) {
	o := orm.NewOrm()
	v := CaveHouse{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&CaveHouse{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
