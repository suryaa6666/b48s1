package main

import (
	"context"
	"fmt"
	"html/template"
	"net/http"
	"personal-web/connection"
	"strconv"
	"time"

	"github.com/labstack/echo/v4"
)

type Blog struct {
	Id       int
	Title    string // title -> unexported, Title -> exported -> bisa diakses di package lain
	Content  string
	Image    string
	Author   string
	PostDate time.Time
}

type User struct {
	Id         int
	Name       string
	Email      string
	Experience []string
	Year       []string
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

	e.GET("/update-blog-form/:id", updateBlogForm)
	e.POST("/update-blog", updateBlog)

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
	userId := 7

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	var dataUser User

	errQuery := connection.Conn.QueryRow(context.Background(), "SELECT id, name, email, experience, year FROM tb_user WHERE id=$1", userId).Scan(&dataUser.Id, &dataUser.Name, &dataUser.Email, &dataUser.Experience, &dataUser.Year)

	// id nantinya dapet dari user login

	if errQuery != nil {
		fmt.Println("masuk sini")
		return c.JSON(http.StatusInternalServerError, errQuery.Error())
	}

	dataReponse := map[string]interface{}{
		"User": dataUser,
	}

	// fmt.Println("takutnya ga masuk datanya : ", dataUser)

	return tmpl.Execute(c.Response(), dataReponse)
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

	dataBlogs, errBlogs := connection.Conn.Query(context.Background(), "SELECT id, title, content, image, post_date FROM tb_blog")

	if errBlogs != nil {
		return c.JSON(http.StatusInternalServerError, errBlogs.Error())
	}

	var resultBlogs []Blog
	for dataBlogs.Next() {
		var each = Blog{}

		err := dataBlogs.Scan(&each.Id, &each.Title, &each.Content, &each.Image, &each.PostDate)
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

	// query get 1 data

	errQuery := connection.Conn.QueryRow(context.Background(), "SELECT id, title, content, image, post_date FROM tb_blog WHERE id=$1", idToInt).Scan(&blogDetail.Id, &blogDetail.Title, &blogDetail.Content, &blogDetail.Image, &blogDetail.PostDate)

	fmt.Println("ini data blog detail: ", errQuery)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// for index, data := range dataBlogs {
	// 	// index += 1
	// 	if index == idToInt { // 1 == 0
	// 		blogDetail = Blog{
	// 			Title:    data.Title,
	// 			Author:   data.Author,
	// 			Content:  data.Content,
	// 			PostDate: data.PostDate,
	// 		}
	// 	}
	// }

	data := map[string]interface{}{ // interface -> tipe data apapun
		"Id":   id,
		"Blog": blogDetail,
	}

	return tmpl.Execute(c.Response(), data)
}

func addBlog(c echo.Context) error {
	title := c.FormValue("title") // surya
	content := c.FormValue("content")

	// append

	// newBlog := Blog{
	// 	Title:    title,
	// 	Author:   "Surya Elidanto",
	// 	Content:  content,
	// 	PostDate: time.Now(),
	// }

	test, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_blog (title, content, image, post_date) VALUES ($1, $2, $3, $4)", title, content, "default.jpg", time.Now())
	// test, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_blog (id, title, content, image, post_date) VALUES ($5, $1, $2, $3, $4)", title, content, "default.jpg", time.Now(), 100)

	fmt.Println("row affected:", test.RowsAffected())

	if err != nil {
		fmt.Println("error guys")
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	// dataBlogs = append(dataBlogs, newBlog) // reassign / timpa

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}

func deleteBlog(c echo.Context) error {
	id := c.Param("id") // ID : 1

	idToInt, _ := strconv.Atoi(id)
	// append

	// slice -> 3 struct (+ 1 struct)

	// slice = append(slice, structlagi)

	// fmt.Println("persiapan delete index : ", id)
	// dataBlogs = append(dataBlogs[:idToInt], dataBlogs[idToInt+1:]...)

	connection.Conn.Exec(context.Background(), "DELETE FROM tb_blog WHERE id=$1", idToInt)

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}

// render tampilan
func updateBlogForm(c echo.Context) error {
	id := c.Param("id") // ID : 1
	tmpl, err := template.ParseFiles("views/update-blog.html")

	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}

	data := map[string]interface{}{
		"Id": id,
	}

	return tmpl.Execute(c.Response(), data)
}

// processing
func updateBlog(c echo.Context) error {
	id := c.FormValue("id")       // ID : 1
	title := c.FormValue("title") // surya
	content := c.FormValue("content")

	fmt.Println("id", id)
	fmt.Println("title", title)
	fmt.Println("content", content)

	idToInt, err := strconv.Atoi(id) // id => "13 "

	if err != nil {
		fmt.Println("gagal conversion ke int")
		return c.JSON(http.StatusInternalServerError, "gagal conversion ke int")
	}

	fmt.Println(idToInt)

	dataUpdate, err := connection.Conn.Exec(context.Background(), "UPDATE tb_blog SET title=$1, content=$2 WHERE id=$3", title, content, id)

	if err != nil {
		fmt.Println("error guys", err.Error())
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	fmt.Println("halo bang", dataUpdate.RowsAffected())

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}
