const http = require("axios");

function getCommodities() {
    return http
        .get(
            "https://stein.efishery.com/v1/storages/5e1edf521073e315924ceab4/list"
        )
        .then(function (response) {
            // console.log(response);
            if (response.status != 200) {
                throw new Error(`status code ${response.code}`)
            }

            return response.data;
        })
        .catch(function (error) {
            // handle error
            console.log(error);
        });
}

module.exports = getCommodities;
