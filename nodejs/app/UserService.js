'use strict';

const UserRepo = require('./UserRepo');

class UserService {

    constructor() {
        this.repo = new UserRepo;
    }

    getAll() {
        return this.repo.getAll();
    }

    store(data) {
        return this.repo.store(data);
    }
}

module.exports = UserService;