'use strict';

class HeatTransformer {

    constructor() {}

    transform(item) {
        return {
            heatId: item._id,
            username: item.user,
            x: item.x,
            y: item.y,
        }
    }

    transformMany(items) {
		for (var i = 0, len = items.length; i < len; i++) {
			items[i] = this.transform(items[i]);
		}
		return items;
	}
}

module.exports = HeatTransformer;