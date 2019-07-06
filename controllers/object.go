package controllers

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"github.com/uniplaces/carbon"
	"lianyun/helpers"
	"lianyun/models"
	"math"
)

// Operations about object
type ObjectController struct {
	beego.Controller
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /setgonggao [options]
func (o *ObjectController) Post1() {
	//var user models.User
	//json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	//uid := models.AddUser(user)
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", allowHeaders)
	o.Data["json"] = map[string]string{"uid": "yangshen"}
	o.ServeJSON()
}
// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /getgonggao [options]
func (o *ObjectController) Post2() {
	//var user models.User
	//json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	//uid := models.AddUser(user)
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", allowHeaders)
	o.Data["json"] = map[string]string{"uid": "yangshen"}
	o.ServeJSON()
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /delgonggao [options]
func (o *ObjectController) Post6() {
	//var user models.User
	//json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	//uid := models.AddUser(user)
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", allowHeaders)
	o.Data["json"] = map[string]string{"uid": "yangshen"}
	o.ServeJSON()
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /getgonggaoone [options]
func (o *ObjectController) Post3() {
	//var user models.User
	//json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	//uid := models.AddUser(user)
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", allowHeaders)
	o.Data["json"] = map[string]string{"uid": "yangshen"}
	o.ServeJSON()
}
// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /updategonggaoone [options]
func (o *ObjectController) Post4() {
	//var user models.User
	//json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	//uid := models.AddUser(user)
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", allowHeaders)
	o.Data["json"] = map[string]string{"uid": "yangshen"}
	o.ServeJSON()
}
// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /updateStatus [options]
func (o *ObjectController) Post5() {
	//var user models.User
	//json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	//uid := models.AddUser(user)
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", allowHeaders)
	o.Data["json"] = map[string]string{"uid": "yangshen"}
	o.ServeJSON()
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /delgonggao [post]
func (o *ObjectController) Delgonggao() {
	//var user models.User
	//json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	//uid := models.AddUser(user)
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", allowHeaders)
	id, _ := o.GetInt("id")
	orm := orm.NewOrm()
	err := orm.Using("lianyun")
	if err != nil {
		logs.Error(err)
	}
	ret,err := orm.QueryTable("announcement").Filter("id", id).Delete()
	if err != nil {
		logs.Error(err)
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"数据库错误"}
		o.ServeJSON()
	}
	o.Data["json"] = map[string]interface{}{"code":200, "data":ret, "msg":"数据库错误"}
	o.ServeJSON()
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /getAnnouncement [post]
func (o *ObjectController) GetAnnouncements() {
	//var user models.User
	//json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	//uid := models.AddUser(user)
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", allowHeaders)
	orm := orm.NewOrm()
	err := orm.Using("lianyun")
	if err != nil {
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"数据库错误"}
		o.ServeJSON()
	}
	conect := &models.Announcement{}
	qs := orm.QueryTable("announcement")
	err = qs.Filter("start_time__lt", carbon.Now().DateTimeString()).
		Filter("end_time__gt", carbon.Now().DateTimeString()).
		Filter("status", 1).
		One(conect)
	if err != nil {
		logs.Error(err)
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"数据库错误"}
		o.ServeJSON()
	}
	o.Data["json"] = map[string]interface{}{"code":200, "data":conect, "msg":"数据库错误"}
	o.ServeJSON()
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /getAnnouncement [get]
func (o *ObjectController) GetAnnouncement() {
	//var user models.User
	//json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	//uid := models.AddUser(user)
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", allowHeaders)
	orm := orm.NewOrm()
	err := orm.Using("lianyun")
	if err != nil {
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"数据库错误"}
		o.ServeJSON()
	}
	conect := &models.Announcement{}
	qs := orm.QueryTable("announcement")
	err = qs.Filter("start_time__lt", carbon.Now().DateTimeString()).
		Filter("end_time__gt", carbon.Now().DateTimeString()).
		Filter("status", 1).
		One(conect)
	if err != nil {
		logs.Error(err)
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"数据库错误"}
		o.ServeJSON()
	}
	o.Data["json"] = map[string]interface{}{"code":200, "data":conect, "msg":"数据库错误"}
	o.ServeJSON()
}
// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /updateStatus [get]
func (o *ObjectController) UpdateStatus() {
	//var user models.User
	//json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	//uid := models.AddUser(user)
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", allowHeaders)
	o.Data["json"] = map[string]string{"uid": "yangshen"}
	conect := &models.Announcement{}
	id, err := o.GetInt("id", 1)
	if err != nil {
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"数据库错误"}
		o.ServeJSON()
	}
	orm := orm.NewOrm()
	err = orm.Using("lianyun")
	if err != nil {
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"数据库错误"}
		o.ServeJSON()
	}
	qs := orm.QueryTable("announcement")
	err  = qs.Filter("id", id).One(conect)
	if err != nil {
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"数据库错误"}
		o.ServeJSON()
	}
	if conect.Status == 1 {
		conect.Status = 0
	}else {
		conect.Status = 1
	}

	logs.Info(conect)
	insertId,err := orm.Update(conect)
	if err != nil {
		logs.Error(err)
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"数据库错误"}
		o.ServeJSON()
	}
	logs.Info(insertId)
	if err != nil {
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"数据库错误"}
		o.ServeJSON()
	}
	o.Data["json"] = map[string]interface{}{"code":200, "data":"", "msg":"数据库错误"}
	o.ServeJSON()
	o.ServeJSON()
}
// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /updategonggaoone [post]
func (o *ObjectController) Updategonggaoone() {

	var startTime, endTime float64
	var dsc string
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", allowHeaders)
	ob := make(map[string]interface{})
	if err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob); err == nil{
		logs.Info("http start_time is ",ob["start_time"])
		startTime = ob["start_time"].(float64)
		endTime =  ob["end_time"].(float64)
		dsc = ob["desc"].(string)
		logs.Info("startime" ,startTime)

	}else{
		logs.Error("https err is ",err)
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"数据库错误"}
		o.ServeJSON()
	}

	conect := &models.Announcement{}
	id, err := o.GetInt("id", 1)
	if err != nil {
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"数据库错误"}
		o.ServeJSON()
	}
	orm := orm.NewOrm()
	err = orm.Using("lianyun")
	if err != nil {
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"数据库错误"}
		o.ServeJSON()
	}
	qs := orm.QueryTable("announcement")
	err  = qs.Filter("id", id).One(conect)
	if err != nil {
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"数据库错误"}
		o.ServeJSON()
	}
	conect.Des = dsc
	logs.Info("end time is ",int64(endTime / math.Pow10(3)))
	stime, _ := carbon.CreateFromTimestamp( int64(startTime  / math.Pow10(3)), "Asia/Shanghai")
	etime, _ := carbon.CreateFromTimestamp( int64(endTime / math.Pow10(3)), "Asia/Shanghai")
	conect.StartTime = stime.Time
	conect.EndTime  = etime.Time
	conect.CreateTime = carbon.Now().Time
	conect.Status = 1
	logs.Info(conect)
	insertId,err := orm.Update(conect)
	if err != nil {
		logs.Error(err)
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"数据库错误"}
		o.ServeJSON()
	}
	logs.Info(insertId)
	if err != nil {
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"数据库错误"}
		o.ServeJSON()
	}
	o.Data["json"] = map[string]interface{}{"code":200, "data":"", "msg":"操作成功"}
	o.ServeJSON()
	o.ServeJSON()
}
// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /getgonggaoone [get]
func (o *ObjectController) Getgonggaoone() {
	//var user models.User
	//json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	//uid := models.AddUser(user)
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", allowHeaders)
	id, err := o.GetInt("id", 1)
	if err != nil {
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"数据库错误"}
		o.ServeJSON()
	}
	orm := orm.NewOrm()
	err = orm.Using("lianyun")
	if err != nil {
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"数据库错误"}
		o.ServeJSON()
	}
	qs := orm.QueryTable("announcement")
	one := &models.Announcement{}
	err  = qs.Filter("id", id).One(one)
	if err != nil {
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"数据库错误"}
		o.ServeJSON()
	}
	data := make(map[string]interface{})
	logs.Info(carbon.NewCarbon(one.StartTime).Time.UnixNano())
	data["Des"] = one.Des
	data["StartTime"] = carbon.NewCarbon(one.StartTime).Time.UnixNano() / 1e6
	logs.Info(data["StartTime"])
	data["EndTime"] = carbon.NewCarbon(one.EndTime).Time.UnixNano() / 1e6
	data["Id"]=one.Id
	o.Data["json"] = map[string]interface{}{"code":200, "data":data, "msg":"登录失败"}
	o.ServeJSON()
}
// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /getgonggao [get]
func (o *ObjectController) Getgonggao() {
	//var user models.User
	//json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	//uid := models.AddUser(user)
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", allowHeaders)
	orm := orm.NewOrm()
	err := orm.Using("lianyun")
	if err != nil {
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"数据库错误"}
		o.ServeJSON()
	}
	qs := orm.QueryTable("announcement")
	page, err := o.GetInt("page", 1)
	status, err := o.GetInt("status", 1)
	if err != nil {
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"登录失败"}
		o.ServeJSON()
	}
	if page == 0 {
		page = 1
	}
	start := (page-1) *15
	end := start + 15
	var users []*models.Announcement
	if err != nil{
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"登录失败"}
		o.ServeJSON()
	}

	switch status {
	case 1:
		qs = qs.OrderBy("-start_time")
		break;
	case 2:
		qs = qs.OrderBy("start_time")
		break;
	case 3:
		qs = qs.OrderBy("-end_time")
		break;
	case 4:
		qs = qs.OrderBy("end_time")
		break;
	case 5:
		qs = qs.OrderBy("create_time")
		break;
	case 6:
		qs = qs.OrderBy("-create_time")
		break;
	case 7:
		qs = qs.OrderBy("id")
		break;
	case 8:
		qs = qs.OrderBy("-id")
		break;
	}
	count,err := qs.All(&users)
	if err != nil {
		logs.Error(err)
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"登录失败"}
		o.ServeJSON()
	}
	if count > 15 {
		users = users[start:end]
	}
	list := helpers.PageUtil(int(count), page, 15, users)
	o.Data["json"] = map[string]interface{}{"code":200, "data":list, "msg":"登录失败"}
	o.ServeJSON()
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router /setgonggao [post]
func (o *ObjectController) Announcement() {
	//var user models.User
	//json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	//uid := models.AddUser(user)
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Origin", "*")
	o.Ctx.ResponseWriter.Header().Add("Access-Control-Allow-Headers", allowHeaders)
	ob := make(map[string]interface{})
	var startTime, endTime float64
	var dsc string
	if err := json.Unmarshal(o.Ctx.Input.RequestBody, &ob); err == nil{
		logs.Info(ob["start_time"])
		startTime = ob["start_time"].(float64)
		endTime =  ob["end_time"].(float64)
		dsc = ob["desc"].(string)
		logs.Info("startime" ,startTime)

	}else{
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"数据库错误"}
		o.ServeJSON()
	}
	conect := &models.Announcement{}
	conect.Des = dsc
	stime, _ := carbon.CreateFromTimestamp( int64(startTime  / math.Pow10(3)), "Asia/Shanghai")
	etime, _ := carbon.CreateFromTimestamp( int64(endTime / math.Pow10(3)), "Asia/Shanghai")
	conect.StartTime = stime.Time
	conect.EndTime  = etime.Time
	conect.CreateTime = carbon.Now().Time
	conect.Status = 1
	insertId,err := models.AddAnnouncement(conect)
	if err != nil {
		logs.Error(err)
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"数据库错误"}
		o.ServeJSON()
	}
	logs.Info(insertId)
	if err != nil {
		o.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"数据库错误"}
		o.ServeJSON()
	}
	o.Data["json"] = map[string]interface{}{"code":200, "data":"", "msg":"数据库错误"}
	o.ServeJSON()
}
// @Title Create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router / [post]
func (o *ObjectController) Post() {
	var ob models.Object
	//cron.Statistics()
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)
	objectid := models.AddOne(ob)
	o.Data["json"] = map[string]string{"ObjectId": objectid}
	o.ServeJSON()
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [get]
func (o *ObjectController) Get() {
	//cron.Statistics()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router / [get]
func (o *ObjectController) GetAll() {
	//cron.Statistics()
}

// @Title Update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /:objectId [put]
func (o *ObjectController) Put() {
	objectId := o.Ctx.Input.Param(":objectId")
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	err := models.Update(objectId, ob.Score)
	if err != nil {
		o.Data["json"] = err.Error()
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJSON()
}

// @Title Delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /:objectId [delete]
func (o *ObjectController) Delete() {
	objectId := o.Ctx.Input.Param(":objectId")
	models.Delete(objectId)
	o.Data["json"] = "delete success!"
	o.ServeJSON()
}

