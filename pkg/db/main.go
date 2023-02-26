package db

import (
	"log"
	"my-website/pkg/db/schema"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open(sqlite.Open("build/db/sqlite.db"), &gorm.Config{})
	if err != nil {
		log.Fatalln("Failed to connect to the database")
	}
	db.AutoMigrate(&schema.Blog{})
	DB = db
}

func GetAllBlogs() (blogs []schema.Blog, err error) {
	err = DB.Find(&blogs).Error
	return
}

func InsertBlog(input schema.Blog) (err error) {
	err = DB.Create(&input).Error
	return
}

func GetBlogBySlug(slug string) (blog schema.Blog) {
	DB.Where("slug = ?", slug).First(&blog)
	return
}
