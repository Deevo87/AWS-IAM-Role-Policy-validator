# AWS-IAM-Role-Policy validator #
A web application written using the [gin-gonic framework](https://gin-gonic.com/) for Golang and html with JavaScript. The apllication checks whether a JSON file is compatible with the AWS::IAM::Role Policy format.

## Prerequisites ##
* Golang 1.21 or later
* Docker (optional)

## Instalation ##
### With git clone ###
Clone the repository on your local machine
```
git clone https://github.com/Deevo87/AWS-IAM-Role-Policy-validator
```
Go to the project directory
```
cd AWS-IAM-Role-Policy-validator
```
Build project
```
go build
```
And run it
```
go run zadanie_remitly
```
The application can be found at [`http://localhost:8080/`](http://localhost:8080/)

### With Docker ###
Go to the project directory and build Docker image
```
docker build --tag json-verifier .
```
Run the Docker container
```
docker -p 8080:8080 json-verifier
```
As above, the application can be found at [`http://localhost:8080/`](http://localhost:8080/)

### Hosted website ###
You can access the application on the website I hosted. Don't worry if the page doesn't render immediately, it will take a while.
[`https://json-verifier-latest-tvek.onrender.com/`](https://json-verifier-latest-tvek.onrender.com/)

## How does it work? ##
The web application is created in the main.go file. When the file is sent for validation, `AppController.go` will check if the file has a `.json` extension, and then a validator will be created. The Validate function in `JsonValidator.go` will decode the file into a preset AWS-IAM-policy structure. In the next step, the validator set is prepared and executed.

We only need to execute one validator thanks to the `chain of responsibility` design pattern. When we call the execute function on the first validator, there will be a chain reaction in which each validator will call its successor. If one of the validators returns an error during the validation process, the whole chain will be broken and we will have to upload the corrected file again.
