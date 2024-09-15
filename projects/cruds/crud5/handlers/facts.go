package handlers

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)

	if err := c.BodyParser(fact); err != nil {
		return NewFactView(c)
	}

	result := database.DB.Db.Create(&fact)
	if result.Error != nil {
		return NewFactView(c)
	}

	return ListFacts(c)
}

func setupRoutes(app *fiber.App) {
	app.Get("/", handlers.ListFacts)
	app.Get("/fact", handlers.NewFactView)
	app.Post("/fact", handlers.CreateFact)
	app.Get("/fact/:id", handlers.ShowFact)

	app.Get("/fact/:id/edit", handlers.EditFact)
	app.Patch("/fact/:id", handlers.UpdateFact)
}

func ShowFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Param("id")

	database.DB.Db.Where("id=?", id).First(&fact)

	return c.Render("show", fiber.Map{
		"Title": "Single Fact",
		"Fact":  fact,
	})
}

func ShowFact(c *fiber.Ctx) error {
	fact := models.Fact{}
	id := c.Params("id")

	result := database.DB.Db.Where("id=?", id).First(&fact)
	if result.Error != nil {
		return NotFound(c)
	}

	return c.Status(fiber.StatusOK).Render("show", fiber.Map{
		"Title": "Single fact",
		"Fact":  fact,
	})
}

func NotFound(c *fiber.Ctx) error {
	return c.Status(fiber.StatusNotFound).SendFile("./public/404.html")
}

func main() {
	database.ConnectDb()

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views:       engine,
		ViewsLayout: "layouts/main",
	})

	setupRoutes(app)

	app.Static("/", "./public")

	app.Use(handlers.NotFound)

	app.Listen(":3000")
}
