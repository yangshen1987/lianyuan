package controllers

import (
	"crypto/md5"
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
	"lianyun/helpers"
	"lianyun/models"

	"github.com/astaxie/beego"
)

// Operations about Users
type UserController struct {
	beego.Controller
}
// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router /user/list [options]
func (u *UserController)test1()  {

	u.Data["json"] = map[string]string{"uid": "yangshen"}
	u.ServeJSON()
}
// @Title GetAll
// @Description fengyonghu
// @Success 200 {object} models.User
// @router /user/feng [options]
func (u *UserController)test2()  {
	u.Data["json"] = map[string]string{"uid": "yangshen"}
	u.ServeJSON()
}
// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [options]
func (u *UserController)test3()  {
	u.Data["json"] = map[string]string{"uid": "yangshen"}
	u.ServeJSON()
}

// @Title CreateUser
// @Description create users
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} models.User.Id
// @Failure 403 body is empty
// @router / [post]
func (u *UserController) Post() {
	//var user models.User
	//json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	//uid := models.AddUser(user)
	u.Data["json"] = map[string]string{"uid": "yangshen"}
	u.ServeJSON()
}

// @Title GetAll
// @Description get all Users
// @Success 200 {object} models.User
// @router /user/list [get]
func (u *UserController) GetAll() {
	page, err := u.GetInt("page", 1)
	if err != nil {
		u.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"登录失败"}
		u.ServeJSON()
	}
	if page == 0 {
		 page = 1
	}
	start := (page-1) *15
	end := start + 15
	var users []*models.User
	o := orm.NewOrm()
	qs := o.QueryTable("user")
	count,err := qs.All(&users)
	if err != nil {
		logs.Error(err)
		u.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"登录失败"}
		u.ServeJSON()
	}
	users = users[start:end]
	list := helpers.PageUtil(int(count), page, 15, users)
	u.Data["json"] = map[string]interface{}{"code":200, "data":list, "msg":"登录失败"}
	u.ServeJSON()
}
// @Title GetAll
// @Description fengyonghu
// @Success 200 {object} models.User
// @router /user/feng [get]
func (u *UserController)Feng()  {
	id,_ := u.GetInt("id")
	userInfo := &models.User{}
	userInfo.Id = id
	o := orm.NewOrm()
	err := o.Read(userInfo)
	if err == orm.ErrNoRows {
		u.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"登录失败"}
		u.ServeJSON()
	}
	if userInfo.Status == 1 {
		userInfo.Status = 2
	}else{
		userInfo.Status = 1
	}

	num, err := o.Update(userInfo)
	if err != nil{
		u.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"登录失败"}
		u.ServeJSON()
	}
	u.Data["json"] = map[string]interface{}{"code":200, "data":num, "msg":"登录失败"}
	u.ServeJSON()
}


// @Title Get
// @Description get user by uid
// @Param	uid		path 	string	true		"The key for staticblock"
// @Success 200 {object} models.User
// @Failure 403 :uid is empty
// @router /:uid [get]
func (u *UserController) Get() {
	//uid := u.GetString(":uid")
	//if uid != "" {
	//	user, err := models.GetUser(uid)
	//	if err != nil {
	//		u.Data["json"] = err.Error()
	//	} else {
	//		u.Data["json"] = user
	//	}
	//}
	//u.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
// @router /:uid [put]
func (u *UserController) Put() {
	//uid := u.GetString(":uid")
	//if uid != "" {
	//	var user models.User
	//	json.Unmarshal(u.Ctx.Input.RequestBody, &user)
	//	uu, err := models.UpdateUser(uid, &user)
	//	if err != nil {
	//		u.Data["json"] = err.Error()
	//	} else {
	//		u.Data["json"] = uu
	//	}
	//}
	//u.ServeJSON()
}

// @Title Delete
// @Description delete the user
// @Param	uid		path 	string	true		"The uid you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 uid is empty
// @router /:uid [delete]
func (u *UserController) Delete() {
	//uid := u.GetString(":uid")
	//models.DeleteUser(uid)
	//u.Data["json"] = "delete success!"
	//u.ServeJSON()
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [post]
func (u *UserController) Login() {
	o := orm.NewOrm()
	err := o.Using("lianyun")
	if err != nil {
		u.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"登录失败"}
		u.ServeJSON()
	}
	user := new(models.Users)
	user.Name = u.GetString("username")
	user.Password = u.GetString("password")
	logs.Info(user)
	data := []byte(user.Password)
	has := md5.Sum(data)
	user.Password = fmt.Sprintf("%x", has)
	logs.Debug(user.Password)
	qs :=o.QueryTable(user)
	err = qs.Filter("name", user.Name).Filter("password", user.Password).One(user)
	if err != nil {
		logs.Error(err)
		u.Data["json"] = map[string]interface{}{"code":201, "data":"", "msg":"登录失败"}
		u.ServeJSON()
	}
	u.Data["json"] = map[string]interface{}{"code":200, "data":user, "msg":"登录失败"}
	u.ServeJSON()
}

// @Title logout
// @Description Logs out current logged in user session
// @Success 200 {string} logout success
// @router /logout [get]
func (u *UserController) Logout() {
	u.Data["json"] = "logout success"
	u.ServeJSON()
}
