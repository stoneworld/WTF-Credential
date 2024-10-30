package handle

import (
	"github.com/gin-gonic/gin"
	"wtf-credential/errors"
	"wtf-credential/middleware"
	"wtf-credential/request"
	"wtf-credential/response"
	"wtf-credential/service"
)

func GetAllCourse(ctx *gin.Context) {

	data, err := service.GetAllCourse(ctx)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	response.JsonSuccess(ctx, data)
}

func GetCoursesByType(ctx *gin.Context) {
	data, err := service.GetCoursesByType(ctx)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	response.JsonSuccess(ctx, data)
}

func GetCourseInfo(ctx *gin.Context) {
	req, err := request.BinGetCourseInfo(ctx)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}

	data, err := service.GetCourseInfo(ctx, req)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}

	response.JsonSuccess(ctx, data)
}

func GetCourseQuizzes(ctx *gin.Context) {
	loginUid, ok := middleware.GetUuidFromContext(ctx)
	if !ok {
		ctx.JSON(200, errors.Entity("failed to retrieve user from context"))
		return
	}
	req, err := request.BinGetCourseQuizzes(ctx)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	data, err := service.GetCourseQuizzes(ctx, req, loginUid)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	response.JsonSuccess(ctx, data)
}

func GetUserCourseLesson(ctx *gin.Context) {
	loginUid, ok := middleware.GetUuidFromContext(ctx)
	if !ok {
		ctx.JSON(200, errors.Entity("failed to retrieve user from context"))
		return
	}

	req, err := request.BinGetUserCourseLesson(ctx)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	data, err := service.GetUserCourseLesson(ctx, req, loginUid)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	response.JsonSuccess(ctx, data)
}

func GetStatistics(ctx *gin.Context) {
	data, err := service.GetStatistics(ctx)
	if err != nil {
		ctx.JSON(200, errors.Unknown(err))
		return
	}
	response.JsonSuccess(ctx, data)
}
