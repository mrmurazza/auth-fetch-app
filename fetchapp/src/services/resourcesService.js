const getCommodities = require("../clients/commodities");
const { convertCurrency } = require("./currencyService");

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
module.exports = getResources;
