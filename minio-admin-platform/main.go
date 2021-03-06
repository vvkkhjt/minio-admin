package main

import (
	"fmt"
	"minio-admin/routers/bucket"
	"minio-admin/routers/user"
	"minio-admin/routers/policy"
	"minio-admin/util"
	"net/http"
)
func main() {
	router := util.Engine
	userRouter :=router.Group("/api/user")
	{
		userRouter.GET("/listuser",user.ListUser)
		userRouter.GET("/getuserinfo",user.GetUserInfo)
		userRouter.POST("/adduser",user.AddUser)
		userRouter.POST("/setuser",user.SetUser)
		userRouter.POST("/setuserstatus",user.SetUserStatus)
		userRouter.POST("/deluser",user.DelUser)
	}

	PolicyRouter :=router.Group("/api/policy")
	{
		PolicyRouter.GET("/listpolicy",policy.ListPolicy)
		PolicyRouter.GET("/getpolicyinfo",policy.GetPolicyInfo)
		PolicyRouter.POST("/addpolicy",policy.AddPolicy)
		PolicyRouter.POST("/setpolicy",policy.SetPolicy)
		PolicyRouter.POST("/delpolicy",policy.DelPolicy)
	}
	BucketRouter :=router.Group("/api/bucket")
	{
		BucketRouter.GET("/listbucket",bucket.ListBuckets)
	}
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", util.Setting.HTTPPort),
		Handler:        router,
		ReadTimeout:    util.Setting.ReadTimeout,
		WriteTimeout:   util.Setting.WriteTimeout,
		MaxHeaderBytes: 1 << 20,
	}
	err := s.ListenAndServe()
	if err != nil{
		util.Logger.Error(err)
		panic(err.Error())
	}
}
