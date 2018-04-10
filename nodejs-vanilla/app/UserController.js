'use strict';

const Service = require('./UserService');
const Transformer = require('./UserTransformer');

class UserController {

    constructor() {
        this.service = new Service;
        this.transformer = new Transformer;
    }

    async getAll(req, res) {
        const users = await this.service.getAll(req.query)
        res.writeHead(200, {'Content-Type': 'application/json'});
        res.end(JSON.stringify(
            {'data': this.transformer.transformMany(users)}
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
        const user = await this.service.store(req.body)
        res.writeHead(201, {'Content-Type': 'application/json'});
        res.end(JSON.stringify(
            { 'data': this.transformer.transform(user)}
        ));
    }
}

module.exports = UserController;