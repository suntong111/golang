package models
import (

	_"github.com/go-sql-driver/mysql"
	"time"
)
type Article struct {
	Id int `json:"id" orm:"pk;auto"`
	Create time.Time `json:"create"orm:"auto_now_add;type(datetime)"`
	Title string `json:"title"orm:"NULL;size(60)"`
	Star int	`json:"star"orm:"NULL"`
	Content string `json:"content"orm:"NULL"`
	Comment string `json:"comment"orm:"NULL"`
	ImgUrl string `json:"img"orm:"NULL"`
	Update time.Time `json:"update" orm:"auto_now_add;type(date)"`
	Author string `json:"author"orm:"NULL;size(60)"`
	Json string  `json:"json"orm:"NULL;type(text)"`
	Img string `json:"img"orm:"NULL"`
}
