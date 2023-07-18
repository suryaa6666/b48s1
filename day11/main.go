package main

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.Static("/assets", "assets")

	e.GET("/hello", helloWorld)
	e.GET("/about", aboutMe)
	e.GET("/home", home)
	e.GET("/contact", contact)
	e.GET("/blog", blog)
	e.GET("/testimonial", testimonial)
	e.GET("/blog-detail/:id", blogDetail)
	e.POST("/add-blog", addBlog)

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

func blog(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/blog.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	// data := map[string]interface{}{
	// 	"Title":   "Title 1",
	// 	"isLogin": false,
	// }

	return tmpl.Execute(c.Response(), nil)
}

func testimonial(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/testimonial.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), nil)
}

func blogDetail(c echo.Context) error {
	id := c.Param("id") // misal : 1

	tmpl, err := template.ParseFiles("views/blog-detail.html")

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	blogDetail := map[string]interface{}{ // interface -> tipe data apapun
		"Id":      id,
		"Title":   "Dumbways ID memang keren",
		"Content": "Dumbways ID adalah bootcamp terbaik sedunia seakhirat!",
	}

	return tmpl.Execute(c.Response(), blogDetail)
}

func addBlog(c echo.Context) error {
	title := c.FormValue("title")
	content := c.FormValue("content")

	fmt.Println("title: ", title)
	fmt.Println("content: ", content)

	return c.Redirect(http.StatusMovedPermanently, "/blog")
}
