'use strict';

const Service = require('./UserService');
const Transformer = require('./UserTransformer');

class UserController {

    constructor() {
        this.service = new Service;
        this.transformer = new Transformer;
    }

    async getAll(req, res) {
        const heats = await this.service.getAll(req.query);
        return res.status(200).json({
            'data': this.transformer.transformMany(heats)
        });
    }

    async getAllUserHasHeatZone(req, res) {
        const users = await this.service.getAllUserHasHeatZone(req.query);
        return res.status(200).json({
            'data': users
        });
    }

    async store(req, res) {
        const heat = await this.service.store(req.body)
        return res.status(201).json({
            'data': this.transformer.transform(heat)
        });
    }
}

module.exports = UserController;