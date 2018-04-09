'use strict';

const User = require('./User');

class UserRepo {

    constructor() {
        this.user = new User;
    }

    getAll() {
        return this.user.getAll();
    }

    store(data) {
        return this.user.store(data);
    }
}

module.exports = UserRepo;