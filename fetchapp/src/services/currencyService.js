const getCurrencyConversion = require("../clients/currency");

var currencyCache = [];

function convertCurrency(amount, from, to) {
    return getConversionFromCache(from, to).then(function (conversion) {
        return conversion.value * amount;
    });
}

function initCache() {
    getConversionFromCache("IDR", "USD");
}

async function getConversionFromCache(from, to) {
    const key = `${from}-${to}`;
    const now = Date.now();
    var conversion = currencyCache[key];

    if (conversion === undefined || conversion["timestamp"] < now) {
        const expiry = new Date();
        const expiryDurationDay =
            process.env.CURRENCY_CONVERSION_CACHE_DAY || 1;
        expiry.setDate(expiry.getDate() + expiryDurationDay);

        await setToCache(expiry, from, to);
    }

    return conversion;
}

async function setToCache(tomorrow, from, to) {
    const key = `${from}-${to}`;
    var conversion = await getCurrencyConversion(1, from, to);

    conversion = {
        timestamp: tomorrow.getTime(),
        value: conversion,
    };

    currencyCache[key] = conversion;
}

module.exports = {
    convertCurrency,
    initCache,
};
