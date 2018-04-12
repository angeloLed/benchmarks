'use strict';

//mongoose
const mongoose = require('mongoose');
mongoose.connect('mongodb://mongodb347/heats');


class User {

    constructor() {
        this.model = mongoose.model('usersheats', { 
            user: String,
            x: Number,
            y: Number
        });
    }

    getAll(filters = {}) {
        return this.model.find(filters);
    }
    
    store(data) {
        const user = new this.model(data);
        return user.save();
    }
}

module.exports = User;