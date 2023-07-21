package main

import (
	"context"
	"html/template"
	"net/http"
	"personal-web/connection"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Blog struct {
	Title    string // title -> unexported, Title -> exported -> bisa diakses di package lain
	Content  string
	Image    string
	Author   string
	PostDate time.Time
}

var dataBlogs = []Blog{ // Blog -> struct biasa, []Blog -> slice of struc, mirip array of object
	{
		Title:    "Title 1",
		Content:  "Content 1",
		Author:   "Surya Elidanto",
		PostDate: time.Now(),
	},
	{
		Title:    "Title 2",
		Content:  "Content 2",
		Author:   "Angga Nur",
		PostDate: time.Now(),
	},
}

func main() {
	e := echo.New()

	connection.DatabaseConnect()

	e.Static("/assets", "assets")

	e.GET("/", home)
	e.GET("/hello", helloWorld)
	e.GET("/about", aboutMe)
	e.GET("/contact", contact)
	e.GET("/form-blog", formBlog)
	e.GET("/blog", blog)
	e.GET("/testimonial", testimonial)
	e.GET("/blog-detail/:id", blogDetail)
	e.POST("/add-blog", addBlog)
	e.POST("/delete-blog/:id", deleteBlog)

	e.Logger.Fatal(e.Start("localhost:5000"))
}

// handler / controller (di php)
func helloWorld(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"name":    "123",
		"address": "3",
	})
}

func aboutMe(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{
		"message": "Halo nama saya Angga",
	})
}

func home(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/index.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func contact(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/contact.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func formBlog(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/form-blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func blog(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	dataBlogs, errBlogs := connection.Conn.Query(context.Background(), "SELECT title, content, image, post_date FROM tb_blog")

	if errBlogs != nil {
		return c.JSON(http.StatusInternalServerError, errBlogs.Error())
	}

	var resultBlogs []Blog
	for dataBlogs.Next() {
		var each = Blog{}

		err := dataBlogs.Scan(&each.Title, &each.Content, &each.Image, &each.PostDate)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}

		each.Author = "Surya Elidanto"

		resultBlogs = append(resultBlogs, each)
	}

	data := map[string]interface{}{
		"Blogs": resultBlogs,
	}

	return tmpl.Execute(c.Response(), data)
}

func testimonial(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func blogDetail(c echo.Context) error {
	id := c.Param("id") // misal : 0

	tmpl, err := template.ParseFiles("views/blog-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	idToInt, _ := strconv.Atoi(id)

	blogDetail := Blog{}

	for index, data := range dataBlogs {
		// index += 1
		if index == idToInt { // 1 == 0
			blogDetail = Blog{
				Title:    data.Title,
				Author:   data.Author,
				Content:  data.Content,
				PostDate: data.PostDate,
			}
		}
	}

	data := map[string]interface{}{ // interface -> tipe data apapun
		"Id":   id,
		"Blog": blogDetail,
	}

	return tmpl.Execute(c.Response(), data)
}

func addBlog(c echo.Context) error {
	title := c.FormValue("title")
	content := c.FormValue("content")

	// append

	newBlog := Blog{
		Title:    title,
		Author:   "Surya Elidanto",
		Content:  content,
		PostDate: time.Now(),
	}

	dataBlogs = append(dataBlogs, newBlog) // reassign / timpa

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}

func deleteBlog(c echo.Context) error {
	id := c.Param("id")

	idToInt, _ := strconv.Atoi(id)
	// append

	// slice -> 3 struct (+ 1 struct)

	// slice = append(slice, structlagi)

	// fmt.Println("persiapan delete index : ", id)
	dataBlogs = append(dataBlogs[:idToInt], dataBlogs[idToInt+1:]...)

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}
