'use strict';

const Service = require('./HeatService');
const Transformer = require('./HeatTransformer');

class HeatController {

    constructor() {
        this.service = new Service;
        this.transformer = new Transformer;
    }

    async getAll(req, res) {
        const heats = await this.service.getAll(req.query)
        res.writeHead(200, {'Content-Type': 'application/json'});
        res.end(JSON.stringify(
            {'data': this.transformer.transformMany(heats)}
        ));
    }

    async getAllUserHasHeatZone(req, res) {
        const users = await this.service.getAllUserHasHeatZone(req.query);
        res.writeHead(200, {'Content-Type': 'application/json'});
        res.end(JSON.stringify(
            { 'data': users }
        ));
    }

    async store(req, res) {
        const heat = await this.service.store(req.body)
        res.writeHead(201, {'Content-Type': 'application/json'});
        res.end(JSON.stringify(
            { 'data': this.transformer.transform(heat)}
        ));
    }
}

module.exports = HeatController;