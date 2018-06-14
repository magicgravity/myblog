package model


type User struct {
	Uid uint32
	UserName string
	Password string
	Email string
	HomeUrl string
	ScreenName string
	Created uint32
	Activated uint32
	Logged uint32
	GroupName string
}


type Log struct {
	Id uint32
	Action string
	Data string
	AuthorId uint32
	Ip string
	Created uint32
}


type Attach struct {
	Id uint32
	FName string
	FType string
	FKey string
	AuthorId uint32
	Created uint32
}

type Comments struct {
	CoId uint32
	Cid uint32
	Created uint32
	Author string
	AuthorId uint32
	OwnerId uint32
	Mail string
	Url string
	Ip string
	Agent string
	Content []byte
	Type string
	Status string
	Parent uint32
}


type Contents struct {
	Cid uint32
	Title string
	Slug string
	Created uint32
	Modified uint32
	Content []byte
	AuthorId uint32
	Type string
	Status string
	Tags string
	Categories string
	Hits uint32
	CommentsNum uint32
	AllowComment bool
	AllowPing bool
	AllowFeed bool
}

type Meta struct {
	Mid uint32
	Name string
	Slug string
	Type string
	Description string
	Sort uint32
	Parent uint32
}

type Options struct {
	Name string
	Value string
	Description string
}


type RelationShip struct {
	Cid uint32
	Mid uint32
}