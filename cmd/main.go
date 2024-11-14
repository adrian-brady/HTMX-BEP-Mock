package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

var ipsum = "Lorem ipsum odor amet, consectetuer adipiscing elit. Lorem felis mi ex senectus, ultricies pharetra. Vel eros est quisque nascetur blandit gravida lectus? Pulvinar vehicula nec nisi ridiculus consequat. Vestibulum pharetra primis justo sapien dictum blandit erat convallis. Quam et pharetra curabitur gravida efficitur conubia augue nisi. Leo posuere elementum imperdiet facilisis vel lobortis? Lorem cubilia nascetur hendrerit fames elit nisl. Est ultricies eleifend ipsum adipiscing magna inceptos mauris. Dis per nascetur morbi consectetur habitasse?"
var data = newData()
var id = 0

type Templates struct {
	templates *template.Template
}

func (t *Templates) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}

func newTemplate() *Templates {
	return &Templates{
		templates: template.Must(template.ParseGlob("views/*.html")),
	}
}

type Result struct {
	ImgPath     string
	Description string
	Title       string
	Id          int
}

func newResult(title string, desc string, path string) Result {
	id++
	return Result{
		Title:       title,
		Description: desc,
		ImgPath:     path,
		Id:          id,
	}
}

type Filter struct {
	Name        string
	ID          int
	Description string
	Active      bool
}

func newFilter(id int, name string, description string) Filter {
	return Filter{
		Name:        name,
		ID:          id,
		Description: description,
		Active:      false,
	}
}

type Results = []Result
type Filters = []Filter

type Data struct {
	Results      Results
	Filters      Filters
	Query        string
	ResultsCount int
	Active       bool
}

type Book struct {
	Title       string
	Author      string
	Description string
	img         string
	Id          int
}

func newBook() Book {
	return Book{
		Title:       "Title",
		Author:      "Author",
		Description: ipsum,
		img:         "images/cptt.jpg",
	}
}

func newData() Data {
	return Data{
		Results: []Result{
			newResult("MyTitle2", ipsum, "images/cptt.jpg"),
			newResult("MyTitle3", ipsum, "images/crafting_interpreters.jpg"),
			newResult("MyTitle2", ipsum, "images/cptt.jpg"),
			newResult("MyTitle3", ipsum, "images/crafting_interpreters.jpg"),
			newResult("MyTitle2", ipsum, "images/cptt.jpg"),
			newResult("MyTitle3", ipsum, "images/crafting_interpreters.jpg"),
			newResult("MyTitle2", ipsum, "images/cptt.jpg"),
			newResult("MyTitle3", ipsum, "images/crafting_interpreters.jpg"),
			newResult("MyTitle2", ipsum, "images/cptt.jpg"),
			newResult("MyTitle3", ipsum, "images/crafting_interpreters.jpg"),
			newResult("MyTitle2", ipsum, "images/cptt.jpg"),
			newResult("MyTitle3", ipsum, "images/crafting_interpreters.jpg"),
			newResult("MyTitle2", ipsum, "images/cptt.jpg"),
			newResult("MyTitle3", ipsum, "images/crafting_interpreters.jpg"),
		},
		Filters: []Filter{
			newFilter(0, "Filter1", "Description1"),
		},
		ResultsCount: 0,
		Query:        "",
	}
}

func home(c echo.Context) error {
	return c.Render(http.StatusOK, "index", 5)
}

func search(c echo.Context) error {
	data.ResultsCount = len(data.Results)
	data.Query = c.FormValue("query")
	return c.Render(http.StatusOK, "results-page", data)
}

func book(c echo.Context) error {
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.String(400, "ERROR")
	}
	book := newBook()
	book.Id = id
	return c.Render(http.StatusOK, "book-page", book)
}

var active = false

func filter(c echo.Context) error {
	newData := data
	idstr := c.Param("id")
	id, err := strconv.Atoi(idstr)
	if err != nil {
		c.String(400, "Error, ID invalid")
	}
	for index, filter := range newData.Filters {
		if filter.ID == id {
			newData.Filters[index].Active = !filter.Active
		}
		fmt.Fprintf(c.Logger().Output(), "%t", newData.Filters[index].Active)
	}
	res := []Result{}
	active = !active
	if active {
		for index, result := range newData.Results {
			if result.Id%2 == 0 {
				res = append(res, newData.Results[index])
			}
		}
		newData.Results = res
	}
	newData.ResultsCount = len(newData.Results)
	newData.Query = c.FormValue("query")
	c.Render(http.StatusOK, "sidecolumn", newData)
	return c.Render(http.StatusOK, "results", newData)
}

func main() {
	e := echo.New()
	e.Use(middleware.Logger())

	e.Static("/images", "images")
	e.Static("/css", "css")

	e.Renderer = newTemplate()

	e.GET("/", home)

	e.POST("/search", search)
	e.GET("/page/:id", book)

	e.POST("/filter/:id", filter)

	e.Logger.Fatal(e.Start(":8080"))
}
