'use strict';

const HeatRepo = require('./HeatRepo');

class HeatService {

    constructor() {
        this.repo = new HeatRepo;
    }

    getAll(filters = {}) {
        return this.repo.getAll(filters);
    }

    getAllUserHasHeatZone(parameters) {
        let filters = {};

        if( Object.prototype.hasOwnProperty.call(parameters, 'x')) {
            const minx = Number(parameters.x) - Number( parameters.radius || 0 );
            const maxx = Number(parameters.x) + Number( parameters.radius || 0 );
            filters.x = { $gte: minx, $lte: maxx };
        }
        if( Object.prototype.hasOwnProperty.call(parameters, 'y')) {
            const miny = Number(parameters.y) - Number( parameters.radius || 0 );
            const maxy = Number(parameters.y) + Number( parameters.radius || 0 );
            filters.y = { $gte: miny, $lte: maxy };
        }

        return this.getAll(filters)
        .then( (heats) => {

            let users = [];
            heats.map((x) => {
                if(users.indexOf(x.user) == -1) {
                    users.push(x.user);
                }
             });

            return users;
        });
    }

    store(data) {
        return this.repo.store(data);
    }
}

module.exports = HeatService;