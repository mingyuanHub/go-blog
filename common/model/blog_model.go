package model

import (
	"go-blog/config"
	"go-blog/common/core"
	// "fmt"

)

type Blog struct {
	Id 			int
	Title 		string
	Content 	string
	CreateTime 	string
	UpdateTime 	string
}

func ListBlog() (list []Blog) {
	rows, err := config.Db().Query("select id, title, content, create_time, update_time from blog")

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

