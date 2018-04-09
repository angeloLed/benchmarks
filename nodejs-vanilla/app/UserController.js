'use strict';

const Service = require('./UserService');
const Transformer = require('./UserTransformer');

class UserController {

    constructor() {
        this.service = new Service;
        this.transformer = new Transformer;
    }

    getAll(req, res) {
        return this.service.getAll()
        .then( (users) => {
            res.writeHead(200, {'Content-Type': 'application/json'});
            res.end(JSON.stringify(
                this.transformer.transformMany(users)
            ));
        });
    }

    store(req, res) {
        return this.service.store(req.body)
        .then( (data) => {
            res.writeHead(201, {'Content-Type': 'application/json'});
            res.end(JSON.stringify(
                this.transformer.transform(data)
            ));
        });
    }
}

module.exports = UserController;