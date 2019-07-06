package models

import (
	"errors"
	"fmt"
	"reflect"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

type StatisticsUser struct {
	Id                  int       `orm:"column(id);auto" description:"用户id"`
	UserExperience      float32   `orm:"column(user_experience);null" description:"经验"`
	UserTransfer        float32   `orm:"column(user_transfer);null" description:"转职"`
	UserRefining        float32   `orm:"column(user_refining);null" description:"肉身等级"`
	UserState           float32   `orm:"column(user_state);null" description:"境界"`
	UserPotential       float32   `orm:"column(user_potential);null" description:"潜力"`
	UserRootBone        float32   `orm:"column(user_root_bone);null" description:"根骨"`
	UserBody            float32   `orm:"column(user_body);null" description:"肉身"`
	UserEndurance       float32   `orm:"column(user_endurance);null" description:"耐力"`
	UserMind            float32   `orm:"column(user_mind);null" description:"心智"`
	UserToken           string    `orm:"column(user_token);size(225);null" description:"用户的token"`
	UserName            string    `orm:"column(user_name);size(225);null" description:"用户姓名"`
	ZhiyeId             int       `orm:"column(zhiye_id);null" description:"职业id"`
	ChushengId          int       `orm:"column(chusheng_id);null" description:"出身id"`
	MubiaoId            int       `orm:"column(mubiao_id);null" description:"目标id"`
	ChallengeNum        int       `orm:"column(challenge_num);null" description:"挑战次数"`
	VipLevel            int       `orm:"column(vip_level);null" description:"用户vip"`
	StoreHp             int       `orm:"column(store_hp);null" description:"用来存储HP的恢复值的"`
	PlayoffTime         int64     `orm:"column(playoff_time);null" description:"下线时间"`
	RemoveBlackMistStep int       `orm:"column(remove_black_mist_step);null" description:"去掉黑雾的步数"`
	XiuweiRate          float32   `orm:"column(xiuwei_rate);null" description:"xiuweixiaolv"`
	XiuJiaSunTime       int       `orm:"column(xiu_jia_sun_time);null" description:"xiuweijiashudan"`
	XiuTime             int       `orm:"column(xiu_time);null" description:"diyici"`
	QuWuPathCount       int       `orm:"column(qu_wu_path_count);null" description:"quwubushu"`
	Money               int       `orm:"column(money);null" description:"钱"`
	Yuanbao             int       `orm:"column(yuanbao);null" description:"元宝"`
	FirstChou           int       `orm:"column(firstChou);null"`
	Kajin               int       `orm:"column(kajin);null"`
	Username            string    `orm:"column(username);size(25)" description:"用户名"`
	Password            string    `orm:"column(password);size(25)" description:"密码"`
	CreatedTime         time.Time `orm:"column(created_time);type(timestamp)" description:"创建时间"`
	RewardTime          time.Time `orm:"column(reward_time);type(timestamp)"`
	Status              int       `orm:"column(status)" description:"1正常2，封号"`
	Chongzhinum         int       `orm:"column(chongzhinum);null"`
	Lingyunum           int       `orm:"column(lingyunum);null"`
	Yuanbaonum          int       `orm:"column(yuanbaonum);null"`
	Yinliangnum         int       `orm:"column(yinliangnum);null"`
	Shiwunum            int       `orm:"column(shiwunum);null"`
	Mucainum            int       `orm:"column(mucainum);null"`
	Caoyaonum           int       `orm:"column(caoyaonum);null"`
	Jingtienum          int       `orm:"column(jingtienum);null"`
	Fengyaota           int       `orm:"column(fengyaota);null"`
	UserId              int       `orm:"column(user_id)" description:"用户id "`
	JueLv               int       `orm:"column(JueLv)"`
}

func (t *StatisticsUser) TableName() string {
	return "statistics_user"
}

func init() {
	orm.RegisterModel(new(StatisticsUser))
}

// AddStatisticsUser insert a new StatisticsUser into database and returns
// last inserted Id on success.
func AddStatisticsUser(m *StatisticsUser) (id int64, err error) {
	o := orm.NewOrm()
	id, err = o.Insert(m)
	return
}

// GetStatisticsUserById retrieves StatisticsUser by Id. Returns error if
// Id doesn't exist
func GetStatisticsUserById(id int) (v *StatisticsUser, err error) {
	o := orm.NewOrm()
	v = &StatisticsUser{Id: id}
	if err = o.Read(v); err == nil {
		return v, nil
	}
	return nil, err
}

// GetAllStatisticsUser retrieves all StatisticsUser matches certain condition. Returns empty list if
// no records exist
func GetAllStatisticsUser(query map[string]string, fields []string, sortby []string, order []string,
	offset int64, limit int64) (ml []interface{}, err error) {
	o := orm.NewOrm()
	qs := o.QueryTable(new(StatisticsUser))
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

	var l []StatisticsUser
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

// UpdateStatisticsUser updates StatisticsUser by Id and returns error if
// the record to be updated doesn't exist
func UpdateStatisticsUserById(m *StatisticsUser) (err error) {
	o := orm.NewOrm()
	v := StatisticsUser{Id: m.Id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Update(m); err == nil {
			fmt.Println("Number of records updated in database:", num)
		}
	}
	return
}

// DeleteStatisticsUser deletes StatisticsUser by Id and returns error if
// the record to be deleted doesn't exist
func DeleteStatisticsUser(id int) (err error) {
	o := orm.NewOrm()
	v := StatisticsUser{Id: id}
	// ascertain id exists in the database
	if err = o.Read(&v); err == nil {
		var num int64
		if num, err = o.Delete(&StatisticsUser{Id: id}); err == nil {
			fmt.Println("Number of records deleted in database:", num)
		}
	}
	return
}
