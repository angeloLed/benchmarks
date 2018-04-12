'use strict';

const Heat = require('./Heat');

class HeatRepo {

    constructor() {
        this.heat = new Heat;
    }

    getAll(filters = {}) {
        return this.heat.getAll(filters);
    }

    store(data) {
        return this.heat.store(data);
    }
}

module.exports = HeatRepo;