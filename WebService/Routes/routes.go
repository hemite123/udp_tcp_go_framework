package routes

import "github.com/gin-gonic/gin"

func PostCollection(r *gin.Engine) {

}

func GetCollection(r *gin.Engine) {

}

func OptionCollection(r *gin.Engine) {

}

func DeleteCollection(r *gin.Engine) {

}

//Dont Change This
func Loader(r *gin.Engine) {
	PostCollection(r)
	GetCollection(r)
	OptionCollection(r)
	DeleteCollection(r)
}
