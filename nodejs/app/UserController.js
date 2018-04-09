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
            return res.status(200).json( this.transformer.transformMany(users) );
        });
    }

    store(req, res) {
        return this.service.store(req.body)
        .then( (data) => {
            return res.status(201).json(this.transformer.transform(data));
        });
    }
}

module.exports = UserController;