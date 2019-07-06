package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

    beego.GlobalControllerRouter["lianyun/controllers:ObjectController"] = append(beego.GlobalControllerRouter["lianyun/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:ObjectController"] = append(beego.GlobalControllerRouter["lianyun/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:ObjectController"] = append(beego.GlobalControllerRouter["lianyun/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:ObjectController"] = append(beego.GlobalControllerRouter["lianyun/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:ObjectController"] = append(beego.GlobalControllerRouter["lianyun/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:objectId`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:ObjectController"] = append(beego.GlobalControllerRouter["lianyun/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post6",
            Router: `/delgonggao`,
            AllowHTTPMethods: []string{"options"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:ObjectController"] = append(beego.GlobalControllerRouter["lianyun/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Delgonggao",
            Router: `/delgonggao`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:ObjectController"] = append(beego.GlobalControllerRouter["lianyun/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAnnouncements",
            Router: `/getAnnouncement`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:ObjectController"] = append(beego.GlobalControllerRouter["lianyun/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "GetAnnouncement",
            Router: `/getAnnouncement`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:ObjectController"] = append(beego.GlobalControllerRouter["lianyun/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post2",
            Router: `/getgonggao`,
            AllowHTTPMethods: []string{"options"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:ObjectController"] = append(beego.GlobalControllerRouter["lianyun/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Getgonggao",
            Router: `/getgonggao`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:ObjectController"] = append(beego.GlobalControllerRouter["lianyun/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post3",
            Router: `/getgonggaoone`,
            AllowHTTPMethods: []string{"options"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:ObjectController"] = append(beego.GlobalControllerRouter["lianyun/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Getgonggaoone",
            Router: `/getgonggaoone`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:ObjectController"] = append(beego.GlobalControllerRouter["lianyun/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post1",
            Router: `/setgonggao`,
            AllowHTTPMethods: []string{"options"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:ObjectController"] = append(beego.GlobalControllerRouter["lianyun/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Announcement",
            Router: `/setgonggao`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:ObjectController"] = append(beego.GlobalControllerRouter["lianyun/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post5",
            Router: `/updateStatus`,
            AllowHTTPMethods: []string{"options"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:ObjectController"] = append(beego.GlobalControllerRouter["lianyun/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "UpdateStatus",
            Router: `/updateStatus`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:ObjectController"] = append(beego.GlobalControllerRouter["lianyun/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Updategonggaoone",
            Router: `/updategonggaoone`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:ObjectController"] = append(beego.GlobalControllerRouter["lianyun/controllers:ObjectController"],
        beego.ControllerComments{
            Method: "Post4",
            Router: `/updategonggaoone`,
            AllowHTTPMethods: []string{"options"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:UserController"] = append(beego.GlobalControllerRouter["lianyun/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post",
            Router: `/`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:UserController"] = append(beego.GlobalControllerRouter["lianyun/controllers:UserController"],
        beego.ControllerComments{
            Method: "Put",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"put"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:UserController"] = append(beego.GlobalControllerRouter["lianyun/controllers:UserController"],
        beego.ControllerComments{
            Method: "Delete",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"delete"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:UserController"] = append(beego.GlobalControllerRouter["lianyun/controllers:UserController"],
        beego.ControllerComments{
            Method: "Get",
            Router: `/:uid`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:UserController"] = append(beego.GlobalControllerRouter["lianyun/controllers:UserController"],
        beego.ControllerComments{
            Method: "Getuser",
            Router: `/getuser`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:UserController"] = append(beego.GlobalControllerRouter["lianyun/controllers:UserController"],
        beego.ControllerComments{
            Method: "Post1",
            Router: `/getuser`,
            AllowHTTPMethods: []string{"options"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:UserController"] = append(beego.GlobalControllerRouter["lianyun/controllers:UserController"],
        beego.ControllerComments{
            Method: "Login",
            Router: `/login`,
            AllowHTTPMethods: []string{"post"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:UserController"] = append(beego.GlobalControllerRouter["lianyun/controllers:UserController"],
        beego.ControllerComments{
            Method: "Loginop",
            Router: `/login`,
            AllowHTTPMethods: []string{"options"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:UserController"] = append(beego.GlobalControllerRouter["lianyun/controllers:UserController"],
        beego.ControllerComments{
            Method: "Logout",
            Router: `/logout`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:UserController"] = append(beego.GlobalControllerRouter["lianyun/controllers:UserController"],
        beego.ControllerComments{
            Method: "Feng",
            Router: `/user/feng`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:UserController"] = append(beego.GlobalControllerRouter["lianyun/controllers:UserController"],
        beego.ControllerComments{
            Method: "Test2",
            Router: `/user/feng`,
            AllowHTTPMethods: []string{"options"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:UserController"] = append(beego.GlobalControllerRouter["lianyun/controllers:UserController"],
        beego.ControllerComments{
            Method: "Test1",
            Router: `/user/list`,
            AllowHTTPMethods: []string{"options"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

    beego.GlobalControllerRouter["lianyun/controllers:UserController"] = append(beego.GlobalControllerRouter["lianyun/controllers:UserController"],
        beego.ControllerComments{
            Method: "GetAll",
            Router: `/user/list`,
            AllowHTTPMethods: []string{"get"},
            MethodParams: param.Make(),
            Filters: nil,
            Params: nil})

}
