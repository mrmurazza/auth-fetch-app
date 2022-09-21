const http = require("axios");

function getCurrencyConversion(amount, from, to) {
    return http
        .get(
            `https://api.apilayer.com/exchangerates_data/convert?to=${to}&from=${from}&amount=${amount}`,
            {
                headers: {
                    apikey: "EdpCarnBHwXo8kSFutq7DOV6dEvBzlfq",
                },
            }
        )
        .then(function (response) {
            console.log(response);
            if (response.status != 200) {
                throw new Error("status code ", response.code);
            }

            return response.data.result;
        })
        .catch(function (error) {
            // handle error
            console.log(error);
        });
    // return Promise.resolve(0.000066571);
}

module.exports = getCurrencyConversion;
