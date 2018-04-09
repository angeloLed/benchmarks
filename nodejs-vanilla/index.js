'use strict';

//express
const http = require('http'); 

// Controller
const userController = new (require('./app/UserController'));

const server = http.createServer().listen(80);
server.on('request', function (req, res) {
    
    if (req.method == 'POST') {
        var body = '';
    }

    req.on('data', function (data) {
        body += data;
    });

    req.on('end', function () {
        //log
        console.log( req.method + ' ' + req.url);

        //body to json
        req.body = JSON.parse(body);

        //sorting route
        if(req.method == 'GET' && req.url == '/users') {
            userController.getAll(req, res);
        }
        if(req.method == 'POST' && req.url == '/users') {
            userController.store(req, res);
        }
    });
});