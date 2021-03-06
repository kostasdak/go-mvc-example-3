package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/kostasdak/gomvc"
)

var c gomvc.Controller

func main() {

	// Load Configuration file
	cfg := gomvc.ReadConfig("./config/config.yml")

	// Connect to database
	db, err := gomvc.ConnectDatabase(cfg.Database.Dbuser, cfg.Database.Dbpass, cfg.Database.Dbname)
	if err != nil {
		log.Fatal(err)
		return
	}
	defer db.Close()

	//Start Server
	srv := &http.Server{
		Addr:    ":" + strconv.FormatInt(int64(cfg.Server.Port), 10),
		Handler: AppHandler(db, cfg),
	}

	fmt.Println("Web app starting at port : ", cfg.Server.Port)

	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

// App handler ... Builds the structure of the app !!!
func AppHandler(db *sql.DB, cfg *gomvc.AppConfig) http.Handler {

	// initialize controller
	c.Initialize(db, cfg)

	// load template files ... path : /web/templates
	// required : homepagefile & template file
	// see [template names] for details
	c.CreateTemplateCache("home.view.tmpl", "base.layout.html")

	// *** Start registering urls, actions and models ***

	// RegisterAction(url, next, action, model)
	// url = url routing path
	// next = redirect after action complete, use in POST actions if necessary
	// model = database model object for CRUD operations

	// create model for Home page content - read content from table pages
	pHomeModel := gomvc.Model{DB: db, PKField: "id", TableName: "pages", OrderString: "ORDER BY id DESC"}
	pHomeModel.DefaultQuery = "SELECT content FROM pages WHERE name='home'"

	// home page : can have two urls "/" and "/home" -> home.view.tmpl must exist
	c.RegisterAction(gomvc.ActionRouting{URL: "/"}, gomvc.ActionView, &pHomeModel)
	c.RegisterAction(gomvc.ActionRouting{URL: "/home"}, gomvc.ActionView, &pHomeModel)

	// create model for [products] database table
	// use the same model for all action in this example
	pViewModel := gomvc.Model{DB: db, PKField: "id", TableName: "products", OrderString: "ORDER BY id DESC"}
	pViewModel.AddRelation(db, "colors", "id", gomvc.SQLKeyPair{LocalKey: "id", ForeignKey: "product_id"}, gomvc.ModelJoinLeft, gomvc.ResultStyleSubresult)

	// optionally assign labels for each table field. Can be used in template view code
	pViewModel.AssignLabels(map[string]string{
		"id":                "Id",
		"code":              "Code",
		"type":              "Veh. Type",
		"name":              "Product",
		"description":       "Description",
		"price":             "Price",
		"images":            "Photos",
		"status":            "Availability",
		"colors.id":         "Id",    //Related field in table [colors]
		"colors.product_id": "PId",   //Related field in table [colors]
		"colors.color":      "Color", //Related field in table [colors]
	})

	// view products ... / show all products || /products/view/{id} for one product
	c.RegisterAction(gomvc.ActionRouting{URL: "/products"}, gomvc.ActionView, &pViewModel)
	c.RegisterAction(gomvc.ActionRouting{URL: "/products/view/*"}, gomvc.ActionView, &pViewModel)

	// create edit model for table [products] -> used for create, update, delete actions only
	pEditModel := gomvc.Model{DB: db, PKField: "id", TableName: "products"}
	pEditModel.AddRelation(db, "colors", "id", gomvc.SQLKeyPair{LocalKey: "id", ForeignKey: "product_id"}, gomvc.ModelJoinLeft, gomvc.ResultStyleSubresult)

	// prepare create product action ... this url has two actions
	// #1 View page -> empty product form no redirect url (no next url required)
	// #2 Post form data to create a new record in table [products] -> then redirect to [next] url -> products page
	c.RegisterAction(gomvc.ActionRouting{URL: "/products/create", NeedsAuth: true}, gomvc.ActionView, &pEditModel)
	c.RegisterAction(gomvc.ActionRouting{URL: "/products/create", NeedsAuth: true, NextURL: "/products"}, gomvc.ActionCreate, &pEditModel)

	// prepare edit product actions ... this url has two actions
	// #1 View page with the product form -> edit form (no next url required)
	// #2 Post form data to update record in table [products] -> then redirect to [next] url -> products page
	c.RegisterAction(gomvc.ActionRouting{URL: "/products/edit/{id}", NeedsAuth: true}, gomvc.ActionView, &pEditModel)
	c.RegisterAction(gomvc.ActionRouting{URL: "/products/edit/{id}", NeedsAuth: true, NextURL: "/products"}, gomvc.ActionUpdate, &pEditModel)

	// prepare delete product actions ... this url has two actions
	// #1 View page with the product form -> edit form [locked] to confirm detetion (no next url required)
	// #2 Post form data to delete record in table [products] -> then redirect to [next] url -> products page
	c.RegisterAction(gomvc.ActionRouting{URL: "/products/delete/*", NeedsAuth: true}, gomvc.ActionView, &pEditModel)
	c.RegisterAction(gomvc.ActionRouting{URL: "/products/delete/*", NeedsAuth: true, NextURL: "/products"}, gomvc.ActionDelete, &pEditModel)

	// create product color model
	cModel := gomvc.Model{DB: db, PKField: "id", TableName: "colors"}
	cModel.AddRelation(db, "products", "id", gomvc.SQLKeyPair{LocalKey: "product_id", ForeignKey: "id"}, gomvc.ModelJoinInner, gomvc.ResultStyleFullresult)

	// prepare add product color page
	c.RegisterAction(gomvc.ActionRouting{URL: "/productcolor/add/*", NeedsAuth: true}, gomvc.ActionView, &cModel)
	c.RegisterAction(gomvc.ActionRouting{URL: "/productcolor/add/*", NeedsAuth: true, NextURL: "/products"}, gomvc.ActionCreate, &cModel)

	// prepare edit product color page
	c.RegisterAction(gomvc.ActionRouting{URL: "/productcolor/edit/*", NeedsAuth: true}, gomvc.ActionView, &cModel)
	c.RegisterAction(gomvc.ActionRouting{URL: "/productcolor/edit/*", NeedsAuth: true, NextURL: "/products"}, gomvc.ActionUpdate, &cModel)

	//prepare delete action / ONLY for post, no view file required
	//redirect will take action after delete action finish
	//see productcolor.edit.tmpl -> file has a form for post to -> productcolor/delete/{id}
	c.RegisterAction(gomvc.ActionRouting{URL: "/productcolor/delete/*", NeedsAuth: true, NextURL: "/products"}, gomvc.ActionDelete, &cModel)

	// prepare about page ... static page, no table/model, no [next] url
	c.RegisterAction(gomvc.ActionRouting{URL: "/about"}, gomvc.ActionView, nil)

	// prepare contact page ... static page, no table/model, no [next] url
	c.RegisterAction(gomvc.ActionRouting{URL: "/contact"}, gomvc.ActionView, nil)

	// create a model for contact page/action
	pContactModel := gomvc.Model{DB: db, PKField: "id", TableName: "messages"}

	// prepare contact page POST action ... static page, no table/model, no [next] url
	// Demostrating how to register a custom func to handle the http request/response using your oun code
	// and handle POST data and have access to database through the controller and model object
	c.RegisterCustomAction(gomvc.ActionRouting{URL: "/contact"}, gomvc.HttpPOST, &pContactModel, ContactPostForm)

	// 							*** Authentication Section ***
	// prepare User Model
	pUserModel := gomvc.Model{DB: db, PKField: "id", TableName: "users"}

	// create auth object
	auth := gomvc.AuthObject{
		UsernameFieldName: "username",
		PasswordFieldName: "password",
		HashCodeFieldName: "auth_hash",
		ExpTimeFieldName:  "expiration",
		SessionKey:        "mywebapplogin",
		ExpireAfterIdle:   5 * time.Minute,
		Model:             pUserModel,
		LoggedInMessage:   "Logged in successfully",
		LoginFailMessage:  "Login fail, please try again",
	}

	// register auth/login URL, login page then -> home page if successfull
	c.RegisterAuthAction("/login", "/", &pUserModel, auth)

	// register sign up page and form
	c.RegisterAction(gomvc.ActionRouting{URL: "/signup"}, gomvc.ActionView, nil)
	c.RegisterCustomAction(gomvc.ActionRouting{URL: "/signup"}, gomvc.HttpPOST, &pUserModel, SignUpPostForm)

	// register sign out url ... no page file required, example to use custom handler function to sign out current user
	c.RegisterCustomAction(gomvc.ActionRouting{URL: "/logout"}, gomvc.HttpGET, &pUserModel, SignOutPage)

	return c.Router
}

// Custom sign out page -> Auth object sign out -> reset cookie in browser
func SignOutPage(w http.ResponseWriter, r *http.Request) {
	// Get auth object
	auth := c.GetAuthObject()

	// Kill auth session -> sign out current user -> kill server if there is an error
	err := auth.KillAuthSession(w, r)
	if err != nil {
		log.Fatal(err)
		return
	}

	//redirect to homepage
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Custom handler for Sign Up page and action,
// this function handles the POST action from "Sign Up" page and uses custom SQL Query execution to create a new user
func SignUpPostForm(w http.ResponseWriter, r *http.Request) {
	// Get model object from controller
	m := c.Models["/signup"]

	// prepare your custom SQL query for execution
	q := "INSERT INTO users (firstname, lastname, username, password, auth_hash, expiration) VALUES (?, ?, ?, ?, ?, ?)"

	// Get auth object
	auth := c.GetAuthObject()

	pass, err := auth.HashPassword(r.Form.Get("password"))
	if err != nil {
		fmt.Println(err)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	// Build value slice
	v := make([]interface{}, 6)
	v[0] = r.Form.Get("firstname")
	v[1] = r.Form.Get("lastname")
	v[2] = r.Form.Get("username")
	v[3] = pass
	v[4] = auth.TokenGenerator()
	v[5] = auth.GetExpirationFromNow()

	// Execute query with values, no need for result value
	_, err = m.Execute(q, v...)
	if err != nil {
		fmt.Println(err)
	}

	//Set session message
	c.GetSession().Put(r.Context(), "flash", "Your signed up successfuly")

	//redirect to homepage
	http.Redirect(w, r, "/", http.StatusSeeOther)
}

// Custom handler for specific page and action,
// this function handles the POST action from "Contact Us" page
func ContactPostForm(w http.ResponseWriter, r *http.Request) {

	fmt.Print("\n")
	fmt.Println("********** ContactPostForm **********")

	// Get model object from controller
	m := c.Models["/contact"]

	// prepare your custom SQL query for execution
	q := "INSERT INTO messages (name, email, message) VALUES (?, ?, ?)"

	// Build value slice
	v := make([]interface{}, 3)
	v[0] = r.Form.Get("name")
	v[1] = r.Form.Get("email")
	v[2] = r.Form.Get("message")

	// Execute query with values, no need for result value
	_, err := m.Execute(q, v...)
	if err != nil {
		fmt.Println(err)
	}

	//Print Posted Form fields, just a test
	fmt.Print("\n")
	fmt.Println("Form fields : ", r.Form)
	fmt.Print("\n")

	//Set session message
	c.GetSession().Put(r.Context(), "flash", "Your message submited")

	//redirect to homepage
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
