'use strict';

class UserTransformer {

    constructor() {}

    transform(item) {
        return {
            id: item._id,
            name: item.name
        }
    }

    transformMany(items) {
		for (var i = 0, len = items.length; i < len; i++) {
			items[i] = this.transform(items[i]);
		}
		return items;
	}
}

module.exports = UserTransformer;