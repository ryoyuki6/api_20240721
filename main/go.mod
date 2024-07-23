module api_20240721/main

go 1.20

replace api_20240721/handler => ../handler

require (
	api_20240721/handler v0.0.0-00010101000000-000000000000
	github.com/gorilla/mux v1.8.1
)
