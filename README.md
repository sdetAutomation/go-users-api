```                                                                      
#             .___      __     _____          __                         __  .__               
#    ______ __| _/_____/  |_  /  _  \  __ ___/  |_  ____   _____ _____ _/  |_|__| ____   ____  
#   /  ___// __ |/ __ \   __\/  /_\  \|  |  \   __\/  _ \ /     \\__  \\   __\  |/  _ \ /    \ 
#   \___ \/ /_/ \  ___/|  | /    |    \  |  /|  | (  <_> )  Y Y  \/ __ \|  | |  (  <_> )   |  \
#  /____  >____ |\___  >__| \____|__  /____/ |__|  \____/|__|_|  (____  /__| |__|\____/|___|  /
#       \/     \/    \/             \/                         \/     \/                    \/ 
```
# go-users-api
sample project building an api using [gin-gonic](https://github.com/gin-gonic/gin) and [golang](https://golang.org/)


Introduction
------------
This project is made for anyone who is looking for an example of how to create a rest endpoint using gin-gonic and go.

This service calls a local sqlite database. Please see sqlite directory for more details. 

This projet was written using Visual Studio Code.   


Running the application
-----
From the root of this project enter the following terminal command:

`go run main.py`  


Project Database
-----
This project uses a local sqlite for a repository.  


Rest Api 
-----

#### Users Api

POST - createUser: [users](http://localhost:8080/users) + include a json body with first_name, last_name, email, and password.

GET - getUserById: [users/1](http://localhost:8080/users/1)

PUT - updateUser: [users/1](http://localhost:8080/users/1) + include a json body with first_name, last_name, and email.

DELETE - deleteUser: [users/1](http://localhost:8080/users/1)

    
Docker
-----
This application can be run in Docker.  Please see Dockerfile for image setup.  Steps to create an image & how to run 
the app in a container list below. (must have docker installed)

Create a docker image: `docker build -t go-api .`

Run docker container: `docker run -it --rm -p 8080:8080 --name mygoapi go-api`

__*** other docker commands ***__

View docker images: `docker images`

View docker containers: `docker ps -a`

Remove docker images: `docker rmi $(docker images -q)`

Remove docker containers: `docker rm $(docker ps -aq)`

[Click here for more information regarding docker](https://docs.docker.com/)


Questions / Contact / Contribute
------------
Feel free to fork this repo, add to it, and create a pull request if you like to contribute.  

If you have any questions, you can contact me via email: `sdet.testautomation@gmail.com`
