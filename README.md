<h1>DELOS TEST INTERNSHIP</h1>

<h3>Before you run this application, you need install mysql and follow this step :</h3>

<ol>
<li>After clone this repository, go to your clone folder</li>
<li>Run go mod tidy</li>
<li>Go to folder ./models/setup to setup your database</li>
<li>Change "root:@tcp(localhost:3306)/delos-internship?parseTime=true" to your configuration, For format : "your_user:your_pass@tcp(your_host:your_port)/your_dbname?charset=utf8mb4&parseTime=True&loc=Local"</li>
<li>Run application with command "go run index.go"</li>
<li>For Testing, you can run command "go test ./unit-test/index_test.go"</li>
</ol>

<a href="https://documenter.getpostman.com/view/13766213/2s9YXk4Lym">API Documentation</a>