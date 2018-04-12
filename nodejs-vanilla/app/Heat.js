'use strict';

//mongoose
const mongoose = require('mongoose');
mongoose.connect('mongodb://mongodb347/heats');


class Heat {

    constructor() {
        this.model = mongoose.model('usersHeats', { 
            user: String,
            x: Number,
            y: Number
        });
    }

    getAll(filters = {}) {
        return this.model.find(filters);
    }
    
    store(data) {
        const heat = new this.model(data);
        return heat.save();
    }
}

module.exports = Heat;