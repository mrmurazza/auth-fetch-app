const http = require("axios");

const debug = process.env.CURRENCY_API_DEBUG || true;

function getCurrencyConversion(amount, from, to) {
    if (debug) {
        return Promise.resolve(0.000066571);
    }

    return http
        .get(
            `https://api.apilayer.com/currency_data/convert?to=${to}&from=${from}&amount=${amount}`,
            {
                headers: {
                    apikey: process.env.CURRENCY_API_KEY || "",
                },
            }
        )
        .then(function (response) {
            if (response.status != 200) {
                throw new Error("status code ", response.code);
            }

            return response.data.result;
        })
        .catch(function (error) {
            // handle error
            console.log(error);
            throw error;
        });
}

module.exports = getCurrencyConversion;
