'use strict';

//mongoose
const mongoose = require('mongoose');
mongoose.connect('mongodb://mongodb347/nodejs');


class User {

    constructor() {
        this.model = mongoose.model('Users', { name: String });
    }

    getAll() {
        return this.model.find();
    }
    
    store(data) {
        const user = new this.model(data);
        return user.save();
    }
}

module.exports = User;