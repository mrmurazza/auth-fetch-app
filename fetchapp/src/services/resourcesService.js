const moment = require("moment");
const getCommodities = require("../clients/commodities");
const { convertCurrency } = require("./currencyService");

async function aggregateResources() {
    return getResources().then(function (res) {
        const reduced = res.reduce(function (rv, resource) {
            if (resource.uuid == null) {
                return rv;
            }

            const date = moment(parseInt(resource.timestamp));
            const week = date.format("YYYY-WW");

            const province = resource.area_provinsi;
            const key = `${province}-${week}`;
            (rv[key] = rv[key] || []).push(resource);
            return rv;
        }, {});

        const aggregated = [];
        Object.entries(reduced).forEach(([key, resources]) => {
            var price = {
                max: 0,
                min: Number.MAX_SAFE_INTEGER,
                sum: 0,
            };
            var size = {
                max: 0,
                min: Number.MAX_SAFE_INTEGER,
                sum: 0,
            };

            resources.forEach((resource) => {
                price = fillAggregateObj(resource, "price", price);
                size = fillAggregateObj(resource, "size", size);
            });

            const firstData = resources[0];
            const date = moment(parseInt(firstData.timestamp)).locale("id");
            const week = date
                .day(0)
                .format("YYYY-MMMM (DD-")
                .concat(date.day(6).format("DD)"));

            aggregated.push({
                area_provinsi: firstData.area_provinsi,
                week: week,
                price: {
                    min: price.min,
                    max: price.max,
                    average: (price.sum / resources.length).toFixed(2),
                    median: getMedianValue(resources, "price"),
                },
                size: {
                    min: size.min,
                    max: size.max,
                    average: (size.sum / resources.length).toFixed(2),
                    median: getMedianValue(resources, "size"),
                },
            });
        });

        return aggregated;
    });
}

function fillAggregateObj(source, key, aggregateObj) {
    const value = parseInt(source[key]);
    if (aggregateObj.max < value) {
        aggregateObj.max = value;
    }
    if (aggregateObj.min > value) {
        aggregateObj.min = value;
    }
    aggregateObj.sum += value;

    return aggregateObj;
}

function getMedianValue(arr, key) {
    const length = arr.length;
    const halfIndex = length / 2;

    // if length is odd get floor of half
    if (length % 2 != 0) {
        return parseInt(arr[Math.floor(halfIndex)][key]);
    }

    const firstIndex = halfIndex - 1;
    const secondIndex = halfIndex;
    return parseInt(arr[firstIndex][key] + arr[secondIndex][key]) / 2;
}

async function getResources() {
    const resource = getCommodities().then(function (commodities) {
        const result = Promise.all(
            commodities.map(function (commodity) {
                const newCom = fillUSDPrice(commodity);
                return newCom;
            })
        );

        return result;
    });

    return resource;
}

async function fillUSDPrice(commodity) {
    return convertCurrency(commodity.price, "IDR", "USD").then(function (
        price_usd
    ) {
        try {
            price_usd = price_usd.toFixed(2).toString();
            if (commodity.uuid == null) {
                price_usd = null;
            }
            commodity.price_usd = price_usd;
            return commodity;
        } catch (error) {
            next(error);
        }
    });
}
module.exports = {
    getResources,
    aggregateResources,
};
