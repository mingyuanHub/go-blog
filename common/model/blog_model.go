package model

import (
	"go-blog/config"
	"go-blog/common/core"
	"fmt"

)

type Blog struct {
	Id 			int
	Title 		string
	Content 	string
	CreateTime 	string
	UpdateTime 	string
}

var column = "id, title, content, create_time, update_time"

func ListBlog() (list []Blog) {
	rows, err := config.Db().Query("select " + column + " from blog")

	if err != nil {
		panic(err)
	}

	blog := Blog{}
	for rows.Next() {
		var createTime, updateTime string
		rows.Scan(&blog.Id, &blog.Title, &blog.Content, &createTime, &updateTime)
		blog.CreateTime = core.TimeToDate(createTime)
		blog.UpdateTime = core.TimeToDate(updateTime)
		list = append(list, blog)
	}

	return
}

func GetBlog(id int) (blog Blog) {
	blog = Blog{}
	err := config.Db().QueryRow("select " + column + " from blog where id = ?", id).Scan(&blog.Id, &blog.Title, &blog.Content, &blog.CreateTime, &blog.UpdateTime)
	if err != nil {
		fmt.Println(err)
	}
	return 
}

