'use strict';

//express
const http = require('http');
const url = require('url');

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
        // console.log( req.method + ' ' + req.url);

        //body to json
        req.body = JSON.parse(body || '{}');

        //create url infos
        const urlInfo = url.parse(req.url, true);

        //build querystring (to simulate express)
        req.query = urlInfo.query;

        //sorting route ( if you want an router engine, maybe is not a vanilla test that you looking for )
        if(req.method == 'GET' && urlInfo.pathname == '/heats') {
            userController.getAll(req, res);
        }
        else if(req.method == 'POST' && urlInfo.pathname == '/heats') {
            userController.store(req, res);
        }
        else if(req.method == 'GET' && urlInfo.pathname == '/userHeats') {
            userController.getAllUserHasHeatZone(req, res);
        }
        else {
            res.writeHead(404, {'Content-Type': 'application/json'});
            res.end(JSON.stringify({'error': 'route not found'}));
        }
    });
});