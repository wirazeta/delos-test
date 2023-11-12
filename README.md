<h1>DELOS TEST INTERNSHIP</h1>

<h3>Before you run this application, you need install mysql and follow this step :</h3>

 1 After clone this repository, go to your clone folder
 2 Run go mod tidy
 3 Go to folder ./models/setup to setup your database
 4 Change "root:@tcp(localhost:3306)/delos-internship?parseTime=true" to your configuration, For format : "your_user:your_pass@tcp(your_host:your_port)/your_dbname?charset=utf8mb4&parseTime=True&loc=Local"
 5 Run application with command "go run index.go"
 6 For Testing, you can run command "go test ./unit-test/index_test.go"

<a href="https://documenter.getpostman.com/view/13766213/2s9YXk4Lym">API Documentation</a>