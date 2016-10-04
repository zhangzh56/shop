package configuration

const (
	MONGO_DB          string = "shop"
	MONGO_DB_PRODUCTS string = "products"

	LocalUnitTests = "dockerhost/" + MONGO_DB
	WithDocker     = "mongo/" + MONGO_DB

	ProductsLocal  = "dockerhost/" + MONGO_DB_PRODUCTS
	ProductsDocker = "mongo/" + MONGO_DB_PRODUCTS

	// MONGO_BASE_URL = "mongodb://dockerhost/"
	// //MONGO_BASE_URL = "mongodb://mongo/"

	//MONGO_URL string = "mongodb://" + LocalUnitTests
	MONGO_URL          string = "mongodb://" + WithDocker
	MONGO_URL_PRODUCTS string = "mongodb://" + ProductsLocal

	MONGO_COLLECTION_CREDENTIALS       string = "credentials"
	MONGO_COLLECTION_SHOP_EVENT_LOG    string = "shop_event_log"
	MONGO_COLLECTION_ORDERS            string = "orders"
	MONGO_COLLECTION_ORDERS_HISTORY    string = "orders_history"
	MONGO_COLLECTION_CUSTOMERS_HISTORY string = "customers_history"
	MONGO_COLLECTION_CUSTOMERS         string = "customers"
	MONGO_COLLECTION_WATCHLISTS        string = "watchlists"
)

// AllowedLanguages contains language codes for all allowed languages
var AllowedLanguages = [...]string{"de", "fr"}
