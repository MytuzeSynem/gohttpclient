gohttpclient

First project in golang

The core of the client (content of gohttp directory) is based on https://github.com/federicoleon/go-httpclient. Since it i my first conntact with language i just followed guy udemy course.

You stated in exercise doc that test should run aginst your fake api, so there is no test on client code itself.

As for the test itself i chose the godog (golang Cucumber implemetation) since you mention well readable test cases. There is some more popular frameworks in go but i have work earier with Cucumber and java, so it was easier start.

The Suite is not have all the posible tests, but enough to show how to write them. Also there was some cases, where api didnt respond according to doc, i.e. 

If no account number or IBAN is provided, Form3 generates a valid account number (see below). If supported by the country, an IBAN is also generated. with GB country

Also i didnt test response headers, to not expose some unwanted data. Also i have seen that sometimes resposne was in proper json format but with text/..  header. It should be tested

There are some cleanup methods that are not showed in the feature files. I cant manage it via before featrue steps, because there is noting like this. I think the better solution will be to initiate env (create some accounts for testng) 
via database query from go in beforeSuite step and clean everything in AfterSuite step, but it didnt have time to refector this.

Scanarios are independent from each other, you can see that in a way that test are run (Dockerfile). There are run conncurenty and in random order.

There are no network defined in docker-compose so i just used network_mode: host, to run aginst accountapi 
