# Invite Token service

### 1. Deployment instruction
##### Setting the database (Postgres)
Please ensure that there is postgres installed and a databse is created with the name `invite_token`. If you want to use a different database. Please update the DatabaseURL in dev.env file.
##### Updating the basic auth credentials
Basic auth is used to authenticate all the admin related endpoints. You can change these credentials using dev.env file `BasicAuthUserName` and `BasicAuthPassword`. 

#### 1.1 Without docker
Use the below commands to run the application:
```
cd invite-token
go mod tidy
go mod vendor
go build
source dev.env
./invite-token
```
#### 1.1 Using docker
Use the below commands to run the application:
```
cd invite-token
docker build -t invite-token .
docker run -e Port='3000' \ 
-e MaxDBConn='5' \
-e DatabaseURL='postgres://postgres@localhost:5432/invite_token?sslmode=disable' \
-e BasicAuthUserName='test' \
-e BasicAuthPassword='test123' \
-p 3000:3000 invite-token
```
Ensure that the database is accessilble inside the docker container.