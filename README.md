Aim :- Creating an Rest API where an user can sign-up the required fields - id, name, mobile number, address, username, password and creating a gRPC API for the user login that accepts username and password, returns login failed or success.

Task:- Firstly storing the data in MySql database. For Rest API I have database configuration, model, routes, controller and main file. For gRPC API I have proto, client and server file.

Process:- 

Rest API: 
config : I have added the database configuration in config.go file under config directory.

model : The model for the user consisting of the following fields - id, name, mobile number, address, username, password, is created in model.go under model directory.

routes : I have created a group to adjust my routing related to users and have to take care of the group while consuming the api. This is created in routes.go file under routes directory.

model, controllers : http requests coming from front end is handled in controller. I have created different functions that handles specific requests routed to controller by the router. I have made user.go file inside model directory to interact with the database. Response to the user is according to the data that is received from my database. If no error is encountered, I have supplied response as StatusOK and if we get an error, error status is supplied. I have created profile_handler.go file inside controllers directory.

main : The starter function of my project is the main.go file. I connect mysql and setup router form here. This is created in the root of the project.

gRPC API:
proto : I have created a proto file for the protocol buffer. The Go code generated from the rpc_login_user.proto file is also kept in the proto directory.

client : The client directory contains the main code for client in client.go file. From here we are sending the request to the server.

server : The server folder contains the server.go code. Here we're connecting with the database and checking if the username and password entered by the user is correct and showing the output accordingly.
