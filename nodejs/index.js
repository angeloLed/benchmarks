'use strict';

//express
const express = require('express');
const bodyParser = require('body-parser');
const app = express();
app.use(bodyParser.json())

// Controller
const userController = new (require('./app/UserController'));

//middleware log
app.use(function (req, res, next) {
    console.log( req.method + ' ' + req.originalUrl);
    next();
});

//routes
app.get('/users', (req,res) => { userController.getAll(req, res) });
app.post('/users', (req,res) => { userController.store(req, res) });

//run
app.listen(80, () => {
    console.log('Server nodejs run');
});