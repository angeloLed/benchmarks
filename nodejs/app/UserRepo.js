'use strict';

const User = require('./User');

class UserRepo {

    constructor() {
        this.user = new User;
    }

    getAll(filters = {}) {
        return this.user.getAll(filters);
    }

    store(data) {
        return this.user.store(data);
    }
}

module.exports = UserRepo;