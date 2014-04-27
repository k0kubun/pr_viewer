// GENERATED CODE - DO NOT EDIT
package routes

import "github.com/revel/revel"


type tApplication struct {}
var Application tApplication


func (_ tApplication) Index(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Application.Index", args).Url
}


type tGorpController struct {}
var GorpController tGorpController


func (_ tGorpController) Begin(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GorpController.Begin", args).Url
}

func (_ tGorpController) Commit(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GorpController.Commit", args).Url
}

func (_ tGorpController) Rollback(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("GorpController.Rollback", args).Url
}


type tStatic struct {}
var Static tStatic


func (_ tStatic) Serve(
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.Serve", args).Url
}

func (_ tStatic) ServeModule(
		moduleName string,
		prefix string,
		filepath string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "moduleName", moduleName)
	revel.Unbind(args, "prefix", prefix)
	revel.Unbind(args, "filepath", filepath)
	return revel.MainRouter.Reverse("Static.ServeModule", args).Url
}


type tSessions struct {}
var Sessions tSessions


func (_ tSessions) Create(
		code string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "code", code)
	return revel.MainRouter.Reverse("Sessions.Create", args).Url
}

func (_ tSessions) Destroy(
		) string {
	args := make(map[string]string)
	
	return revel.MainRouter.Reverse("Sessions.Destroy", args).Url
}


type tUsers struct {}
var Users tUsers


func (_ tUsers) Show(
		login string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "login", login)
	return revel.MainRouter.Reverse("Users.Show", args).Url
}

func (_ tUsers) Update(
		login string,
		) string {
	args := make(map[string]string)
	
	revel.Unbind(args, "login", login)
	return revel.MainRouter.Reverse("Users.Update", args).Url
}


