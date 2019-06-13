package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Hero struct {
	Id              int `orm:"column(id);auto" description:"自增id"`
	HeroId          int `orm:"column(hero_id)" description:"英雄id"`
	HeroLevel       int `orm:"column(hero_level)" description:"侠客等级"`
	HeroStart       int `orm:"column(hero_start)" description:"侠客星级"`
	HeroState       int `orm:"column(hero_state)" description:"侠客境界"`
	SkillId         int `orm:"column(skill_id)" description:"技能id"`
	Exp             int `orm:"column(exp)" description:"经验"`
	SkillOneLevel   int `orm:"column(skill_one_level)" description:"技能1等级"`
	SkillTwoLevel   int `orm:"column(skill_two_level)" description:"技能2等级"`
	SkillThreeLevel int `orm:"column(skill_three_level)" description:"技能3等级"`
	SkillFourLevel  int `orm:"column(skill_four_level)" description:"技能4等级"`
	Fragment        int `orm:"column(fragment)" description:"碎片"`
	IsAttach        int `orm:"column(is_attach)" description:"是否在阵上"`
	UserId          int `orm:"column(user_id)" description:"用户id"`
	MapId           int `orm:"column(map_id)" description:"驻守id"`
	GarrisonOrder   int `orm:"column(garrison_order)" description:"驻守次序"`
}

func (t *Hero) TableName() string {
	return "hero"
}

func init() {
	orm.RegisterModel(new(Hero))
}

// AddHero insert a new Hero into database and returns
// last inserted Id on success.
func AddHero(m *Hero) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetHeroById retrieves Hero by Id. Returns error if
// Id doesn't exist
func GetHeroById(id int) (v *Hero, err error) {
	o := orm.NewOrm()
	v = &Hero{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllHero retrieves all Hero matches certain condition. Returns empty list if
// no records exist
func GetAllHero(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Hero))
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

	var l []Hero
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

// UpdateHero updates Hero by Id and returns error if
// the record to be updated doesn't exist
func UpdateHeroById(m *Hero) (err error) {
	o := orm.NewOrm()
	v := Hero{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteHero deletes Hero by Id and returns error if
// the record to be deleted doesn't exist
func DeleteHero(id int) (err error) {
	o := orm.NewOrm()
	v := Hero{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Hero{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
