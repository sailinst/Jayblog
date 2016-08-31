package models

import (
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql"

	"os"
	"path"
	"strconv"
	"time"
)

const (
	_DB_NAME      = "data/Jayblog"
	_MYSQL_DRIVER = "mysql"
)

type User struct {
	Id       int64
	Name     string
	Password string
	Email    string
}

type Category struct {
	Id              int64
	Title           string
	Created         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	TopicTime       time.Time `orm:"index"`
	TopicCount      int64
	TopicLastUserId int64
}
type Topic struct {
	Id              int64
	Uid             int64
	Title           string
	Content         string `orm:"size(5000)"`
	Attachment      string
	Created         time.Time `orm:"index"`
	Updated         time.Time `orm:"index"`
	Views           int64     `orm:"index"`
	Author          string
	ReplyTime       time.Time `orm:"index"`
	ReplyCount      int64
	ReplyLastUserId int64
}
type Photoalbum struct {
	Id             int64
	PhotoalbumName string
	PhotoalbumDate time.Time `orm:"index"`
}

type Photograph struct {
	Id             int64
	photographName string
	photographDate time.Time `orm:"index"`
	photoAlbumName string
	photoURL       string
}

func RegisterDB() {
	if !com.IsExist(_DB_NAME) {
		os.MkdirAll(path.Dir(_DB_NAME), os.ModePerm)
		os.Create(_DB_NAME)
	}
	orm.RegisterModel(new(User), new(Category), new(Topic), new(Photoalbum), new(Photograph))
	orm.RegisterDriver(_MYSQL_DRIVER, orm.DRMySQL)
	orm.RegisterDataBase("default", _MYSQL_DRIVER, "root:Aa1234@/Jayblog?charset=utf8", 20, 20)

}
func AddUser(username, password, email string) error {
	o := orm.NewOrm()
	/*
		user := &User{
			Name:     username,
			Password: password,
			Email:    email,
		} */
	_, err := o.Insert(&User{Name: username, Password: password, Email: email})
	return err
}
func AddTopic(title, content string) error {
	o := orm.NewOrm()

	topic := &Topic{
		Title:   title,
		Content: content,
		Created: time.Now(),
		Updated: time.Now(),
	}
	_, err := o.Insert(topic)

	return err

}
func AddCategory(name string) error {
	o := orm.NewOrm()

	cate := &Category{
		Title:     name,
		Created:   time.Now(),
		TopicTime: time.Now(),
	}

	// 查询数据
	qs := o.QueryTable("category")
	err := qs.Filter("title", name).One(cate)
	if err == nil {
		return err
	}

	// 插入数据
	_, err = o.Insert(cate)
	if err != nil {
		return err
	}

	return nil
}

func DelCategory(id string) error {
	cid, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		return err
	}

	o := orm.NewOrm()

	cate := &Category{Id: cid}
	_, err = o.Delete(cate)
	return err
}
func GetAllTopics() ([]*Topic, error) {
	o := orm.NewOrm()
	topics := make([]*Topic, 0)
	qs := o.QueryTable("topic")
	_, err := qs.All(&topics)
	return topics, err
}
func GetAllCategories() ([]*Category, error) {
	o := orm.NewOrm()
	cates := make([]*Category, 0)
	qs := o.QueryTable("category")
	_, err := qs.All(&cates)
	return cates, err
}
