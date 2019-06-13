package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"

	"github.com/astaxie/beego/orm"
)

type Act struct {
	Id                 int    `orm:"column(id);auto"`
	JiYuanSetNum       int    `orm:"column(JiYuanSetNum);null" description:"选择了第几个 int"`
	QiYuJiangLv        int    `orm:"column(QiYuJiangLv);null" description:"奇遇的奖励到了lv几"`
	PlayerIntime       int    `orm:"column(PlayerIntime);null" description:"奇遇领奖用的在线时间"`
	QiT                string `orm:"column(QiT);size(255);null" description:"一个10位的数组"`
	SelectT            string `orm:"column(SelectT);size(255);null" description:"一个5位的数组"`
	QiYuanLing         int    `orm:"column(QiYuanLing);null" description:"今天是否已经领取  （1或0）"`
	SureSet            int    `orm:"column(SureSet);null" description:"是否已经Set"`
	IsLing             string `orm:"column(IsLing);size(255);null" description:"每天VIP的奖励是否领取【3位的数组】"`
	DengLuDay          int    `orm:"column(DengLuDay);null" description:"累计登录的天数"`
	LeiDeng            int    `orm:"column(LeiDeng);null" description:"连续三天的累计登录，每隔3天刷新一次"`
	DailyLing          int    `orm:"column(dailyLing);null" description:"每日登陆奖励是否领取 （1，0）"`
	MeiRiZaiXianLing   string `orm:"column(MeiRiZaiXianLing);size(191);null" description:" 每日在线时间段是否已经领取 【3位的数组】"`
	ChengJiuDailyLing  string `orm:"column(ChengJiuDailyLing);size(199);null" description:"每日成就的完成与领取  【我传一个json】"`
	ChengZhangChengJiu string `orm:"column(ChengZhangChengJiu);size(199);null" description:"成长的成就等级 【4位的数组】"`
	JieYinMa           string `orm:"column(JieYinMa);size(192);null" description:"自身的接引码"`
	YiJieYin           int    `orm:"column(YiJieYin);null" description:"是否已经接引"`
	ShuRuJieYinMa      string `orm:"column(ShuRuJieYinMa);size(192);null" description:"输入的接引码"`
	DuiTable           string `orm:"column(DuiTable);size(255);null" description:"一个数组"`
	CanDouble          int    `orm:"column(canDouble);null" description:"判断今日起源领取是否可以双倍"`
	UserId             int    `orm:"column(user_id)"`
}

func (t *Act) TableName() string {
	return "act"
}

func init() {
	orm.RegisterModel(new(Act))
}

// AddAct insert a new Act into database and returns
// last inserted Id on success.
func AddAct(m *Act) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetActById retrieves Act by Id. Returns error if
// Id doesn't exist
func GetActById(id int) (v *Act, err error) {
	o := orm.NewOrm()
	v = &Act{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllAct retrieves all Act matches certain condition. Returns empty list if
// no records exist
func GetAllAct(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(Act))
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

	var l []Act
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

// UpdateAct updates Act by Id and returns error if
// the record to be updated doesn't exist
func UpdateActById(m *Act) (err error) {
	o := orm.NewOrm()
	v := Act{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteAct deletes Act by Id and returns error if
// the record to be deleted doesn't exist
func DeleteAct(id int) (err error) {
	o := orm.NewOrm()
	v := Act{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&Act{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
