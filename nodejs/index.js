'use strict';

//express
const express = require('express');
const bodyParser = require('body-parser');
const app = express();
app.use(bodyParser.json())

// Controller
const heatController = new (require('./app/HeatController'));

//middleware log
// app.use(function (req, res, next) {
//     console.log( req.method + ' ' + req.originalUrl);
//     next();
// });

//routes
app.get('/heats', (req,res) => { heatController.getAll(req, res) });
app.post('/heats', (req,res) => { heatController.store(req, res) });
app.get('/userHeats', (req,res) => { heatController.getAllUserHasHeatZone(req, res) });

//run
app.listen(80, () => {
    console.log('Server nodejs run');
});