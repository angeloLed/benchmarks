# Benchmark Dockerized Microservices #

## The purpose

This project give to you a clearer idea of what are the real performances of microservices backend with ODM / ORM. Usually, the benchmarks that you find around the web are limited to a simple "hello word" or something else of "easy".

Obviously, there is no absolutely perfect language, there is only the most suitable according to the needs.

 

## The Project

This is a "real-life" benchmarks of some dockerized CRUD backend applications in some different languages. 

There are two routes for every CRUD, one to insert the records of user heatmap and the other one to read them with a query and code logic.

  

All is under container docker, just run "docker-composer up" and all the needs will be resolve automatically. 

  

### Candidates
- PHP 7.2 + Lumen v5.6 + Jenssegers MongoDB ODM v3.4.1 
- Nodejs 8.9.4 + Express v4.16.3 + Mongoose v5.0.13 
- Nodejs 8.9.4 + Mongoose v5.0.13 
- Golang 1.10.1 + Gin 1.2
- DotNetCore 2.0.6 with Kestrel config

### Database
- Mongodb 3.4.7

### Note
PHP Image: https://github.com/OsLab/docker-php-nginx with mongo driver 

### Configuration of JMeter 3.3 r1808647
 - randomize parameters
 - Thears: 200
 - RUP: 20

# GO WITH BENCHMARK

### Docker Configuration
- 2 CORE
- 2 GB RAM
- 1 GB SWAP

DotNetCore: POST 
![picture](https://github.com/angeloLed/benchmarks/blob/master/other/DN-POST.png "")

DotNetCore: GET 
![picture](https://github.com/angeloLed/benchmarks/blob/master/other/DN-GET.png "")

NodeJs + express: POST 
![picture](https://github.com/angeloLed/benchmarks/blob/master/other/NODEJS-POST.png "")

NodeJs + express: GET 
![picture](https://github.com/angeloLed/benchmarks/blob/master/other/NODEJS-GET.png "")

NodeJs Vanilla: POST 
![picture](https://github.com/angeloLed/benchmarks/blob/master/other/NODEJSV-POST.png "")

NodeJs Vanilla: GET 
![picture](https://github.com/angeloLed/benchmarks/blob/master/other/NODEJSV-GET.png "")

PHP: POST 
![picture](https://github.com/angeloLed/benchmarks/blob/master/other/PHP-POST.png "")

PHP: GET 
![picture](https://github.com/angeloLed/benchmarks/blob/master/other/PHP-GET.png "")

GO: POST 
![picture](https://github.com/angeloLed/benchmarks/blob/master/other/GO-POST.png "")

GO: GET 
![picture](https://github.com/angeloLed/benchmarks/blob/master/other/GO-GET.png "")

### Resume
![picture](https://github.com/angeloLed/benchmarks/blob/master/other/resume-rs-1.png "")
![picture](https://github.com/angeloLed/benchmarks/blob/master/other/resume-avg-1.png "")

## NEXT
- other languages with other frameworks
- other docker configuration