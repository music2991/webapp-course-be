package db

import "webapp-course-be/db/schema"

func GetConnection() *connection {
	if dbConnection == nil {
		dbConnection = &connection{
			Product: schema.NewProducts(),
		}
		return dbConnection
	}
	return dbConnection
}

var dbConnection *connection

type connection struct {
	Product *schema.Products
}
