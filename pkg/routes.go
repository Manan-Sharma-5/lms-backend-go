package routes

import (
	"backend/internal/auth"
	"backend/internal/classroom"
	fetchrequests "backend/internal/fetch-requests"
	file_upload "backend/internal/fileupload"
	middleware "backend/internal/middlewares"

	"github.com/gin-gonic/gin"
)

func Routes(r* gin.Engine){
	apiRouter := r.Group("/api/v1")
	r.POST("/signup", auth.SignUp)
	r.POST("/signin", auth.SignIn)
	apiRouter.Use(middleware.AuthMiddleware())
	{
		apiRouter.POST("/file-upload", file_upload.FileUpload)
		apiRouter.POST("/previous-year-upload", file_upload.PreviousYearUpload)
		apiRouter.POST("/view-notes", fetchrequests.FetchNotes)
		apiRouter.POST("/view-pyqs", fetchrequests.FetchPYQS)
		apiRouter.POST("/create-class", classroom.CreateClass)
		apiRouter.GET("/fetch-class", classroom.FetchClassForUser)
		apiRouter.POST("/fetch-class-subject", classroom.FetchClassBySubject)
		apiRouter.POST("/fetch-subjects-notes", fetchrequests.FetchSubjectsForNotes)
		apiRouter.POST("/fetch-subjects-pyqs", fetchrequests.FetchSubjectsForPapers)
	}
}