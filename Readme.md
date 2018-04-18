# Benchmark Microservices Project #

## The purpose

This project is to have a clearer idea of what are the real performances of microservices backend with framework and ODM / ORM. Usually, the benchmarks that you find around the web are limited to a simple "hello word" or something else of "easy".

Obviously, there is no absolutely perfect language, there is only the most suitable according to the needs. 

 

## The Project

This is a "real-life" benchmarks of some dockerized CRUD backend applications in some different languages. 

There are two routes for every CRUD, one to insert the records of user heatmap and the other one to read them with a query with code logic.

  

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

### Docker Configuration
- 2 CPU
- 2 GB RAM
- 1 GB SWAP

 

### Tool for benchmarking: 
 - Apache JMeter 3.3 r1808647 with randomized parameter

## RESULTS

DotNetCore - GET 
![picture](https://github.com/angeloLed/benchmarks/blob/master/other/DN-POST.png "")