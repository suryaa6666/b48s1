package main

import (
	"context"
	"database/sql"
	"fmt"
	"html/template"
	"net/http"
	"personal-web/connection"
	"personal-web/middleware"
	"strconv"
	"time"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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
	Id             int
	Name           string
	Email          string
	HashedPassword string
	Experience     []string
	Year           []string
}

type UserLoginSession struct {
	IsLogin bool
	Name    string
}

var userLoginSession = UserLoginSession{}

// var dataBlogs = []Blog{ // Blog -> struct biasa, []Blog -> slice of struc, mirip array of object
// 	{
// 		Title:    "Title 1",
// 		Content:  "Content 1",
// 		Author:   "Surya Elidanto".(sql.NullString),
// 		PostDate: time.Now(),
// 	},
// 	{
// 		Title:    "Title 2",
// 		Content:  "Content 2",
// 		Author:   "Angga Nur",
// 		PostDate: time.Now(),
// 	},
// }

func main() {
	e := echo.New()

	connection.DatabaseConnect()

	e.Use(session.Middleware(sessions.NewCookieStore([]byte("suryaganteng"))))

	e.Static("/assets", "assets")
	e.Static("/uploads", "uploads")

	e.GET("/", home)
	e.GET("/hello", helloWorld)
	e.GET("/about", aboutMe)
	e.GET("/contact", contact)
	e.GET("/form-blog", formBlog)
	e.GET("/blog", blog)
	e.GET("/testimonial", testimonial)
	e.GET("/blog-detail/:id", blogDetail)
	e.POST("/add-blog", middleware.UploadFile(addBlog))
	e.POST("/delete-blog/:id", deleteBlog)
	e.GET("/update-blog-form/:id", updateBlogForm)
	e.POST("/update-blog", updateBlog)

	// auth
	e.GET("/form-login", formLogin)
	e.POST("/login", login)

	e.GET("/form-register", formRegister)
	e.POST("/register", register)

	e.POST("/logout", logout)

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

	sess, errSess := session.Get("session", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	flash := map[string]interface{}{
		"FlashMessage": sess.Values["message"], // "Register berhasil"
		"FlashStatus":  sess.Values["status"],  // true
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	dataReponse := map[string]interface{}{
		"User":  dataUser,
		"Flash": flash,
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

	dataBlogs, errBlogs := connection.Conn.Query(context.Background(), "SELECT tb_blog.id, tb_user.name, tb_blog.title, tb_blog.content, tb_blog.image, tb_blog.post_date FROM tb_blog LEFT JOIN tb_user ON tb_blog.author_id = tb_user.id")

	if errBlogs != nil {
		return c.JSON(http.StatusInternalServerError, errBlogs.Error())
	}

	var resultBlogs []Blog
	for dataBlogs.Next() {
		var each = Blog{}

		// each.Author = "Surya Elidanto" // udah otomatis, kita matiin
		var tempAuthor sql.NullString // temp -> temporary -> sementara

		err := dataBlogs.Scan(&each.Id, &tempAuthor, &each.Title, &each.Content, &each.Image, &each.PostDate)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, err.Error())
		}
		fmt.Println("ini datamu bos : ", tempAuthor.String)

		each.Author = tempAuthor.String

		resultBlogs = append(resultBlogs, each)
	}

	sess, _ := session.Get("session", c)

	if sess.Values["isLogin"] != true {
		userLoginSession.IsLogin = false
	} else {
		userLoginSession.IsLogin = true
		userLoginSession.Name = sess.Values["name"].(string)
	}

	data := map[string]interface{}{
		"Blogs":            resultBlogs,
		"UserLoginSession": userLoginSession,
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

	var tempAuthor sql.NullString

	errQuery := connection.Conn.QueryRow(context.Background(), "SELECT tb_blog.id, tb_user.name, tb_blog.title, tb_blog.content, tb_blog.image, tb_blog.post_date FROM tb_blog LEFT JOIN tb_user ON tb_blog.author_id = tb_user.id WHERE tb_blog.id=$1", idToInt).Scan(&blogDetail.Id, &tempAuthor, &blogDetail.Title, &blogDetail.Content, &blogDetail.Image, &blogDetail.PostDate)

	blogDetail.Author = tempAuthor.String

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

	image := c.Get("dataFile").(string) // image-12930812039.png

	// append

	// newBlog := Blog{
	// 	Title:    title,
	// 	Author:   "Surya Elidanto",
	// 	Content:  content,
	// 	PostDate: time.Now(),
	// }

	sess, _ := session.Get("session", c)

	test, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_blog (title, content, image, post_date, author_id) VALUES ($1, $2, $3, $4, $5)", title, content, image, time.Now(), sess.Values["id"].(int))
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

func formLogin(c echo.Context) error {
	// bikin pengecekan
	// ngambil dari session datanya, misalnya isLogin -> false
	// sess, _ := session.Get("session", c)

	// if sess.Values["isLogin"] != true {
	// 	return c.Redirect(http.StatusMovedPermanently, "/")
	// }

	tmpl, err := template.ParseFiles("views/form-login.html")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	sess, errSess := session.Get("session", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	// fmt.Println("message:", sess.Values["message"])
	// fmt.Println("status:", sess.Values["status"])

	flash := map[string]interface{}{
		"FlashMessage": sess.Values["message"], // "Register berhasil"
		"FlashStatus":  sess.Values["status"],  // true
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	return tmpl.Execute(c.Response(), flash)
}

func login(c echo.Context) error {
	inputEmail := c.FormValue("inputEmail")
	inputPassword := c.FormValue("inputPassword") //

	user := User{}

	// check apakah ada emailnya di db
	err := connection.Conn.QueryRow(context.Background(), "SELECT id, name, email, password FROM tb_user WHERE email=$1", inputEmail).Scan(&user.Id, &user.Name, &user.Email, &user.HashedPassword)

	if err != nil {
		return redirectWithMessage(c, "Login gagal!", false, "/form-login")
	}

	errPassword := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(inputPassword))

	if errPassword != nil {
		return redirectWithMessage(c, "Login gagal!", false, "/form-login")
	}

	// return c.JSON(http.StatusOK, "Berhasil login!")

	// set session login (berhasil login)
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = 10800 // 3 JAM -> berapa lama expired
	sess.Values["message"] = "Login success!"
	sess.Values["status"] = true
	sess.Values["name"] = user.Name
	sess.Values["email"] = user.Email
	sess.Values["id"] = user.Id
	sess.Values["isLogin"] = true
	sess.Save(c.Request(), c.Response())

	return c.Redirect(http.StatusMovedPermanently, "/")
}

func formRegister(c echo.Context) error {
	tmpl, err := template.ParseFiles("views/form-register.html")

	sess, errSess := session.Get("session", c)
	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	// fmt.Println("message:", sess.Values["message"])
	// fmt.Println("status:", sess.Values["status"])

	flash := map[string]interface{}{
		"FlashMessage": sess.Values["message"], // "Register berhasil"
		"FlashStatus":  sess.Values["status"],  // true
	}

	delete(sess.Values, "message")
	delete(sess.Values, "status")
	sess.Save(c.Request(), c.Response())

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	return tmpl.Execute(c.Response(), flash)
}

func register(c echo.Context) error {

	inputName := c.FormValue("inputName")
	inputEmail := c.FormValue("inputEmail") // harus valid email
	inputPassword := c.FormValue("inputPassword")
	// min : 4 character, minimal harus ada 1 special character, 1 huruf besar

	// validasi (trim, validasi valid email)

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(inputPassword), 10)

	if err != nil {
		// fmt.Println("masuk sini")
		return c.JSON(http.StatusInternalServerError, err.Error())
	}

	fmt.Println(inputName, inputEmail, inputPassword)

	query, err := connection.Conn.Exec(context.Background(), "INSERT INTO tb_user (name, email, password) VALUES($1, $2, $3)", inputName, inputEmail, hashedPassword)

	fmt.Println("affected row : ", query.RowsAffected())

	if err != nil {
		return redirectWithMessage(c, "Register gagal!", false, "/form-register")
	}

	return redirectWithMessage(c, "Register berhasil!", true, "/form-login")
}

func logout(c echo.Context) error {
	sess, _ := session.Get("session", c)
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())

	return redirectWithMessage(c, "Logout berhasil!", true, "/")
}

func redirectWithMessage(c echo.Context, message string, status bool, redirectPath string) error {
	sess, errSess := session.Get("session", c)

	if errSess != nil {
		return c.JSON(http.StatusInternalServerError, errSess.Error())
	}

	sess.Values["message"] = message
	sess.Values["status"] = status
	sess.Save(c.Request(), c.Response())
	return c.Redirect(http.StatusMovedPermanently, redirectPath)
}
