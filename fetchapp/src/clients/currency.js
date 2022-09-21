const http = require("axios");

function getCurrencyConversion(amount, from, to) {
    return http
        .get(
            `https://api.apilayer.com/currency_data/convert?to=${to}&from=${from}&amount=${amount}`,
            {
                headers: {
                    apikey: "EdpCarnBHwXo8kSFutq7DOV6dEvBzlfq",
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
