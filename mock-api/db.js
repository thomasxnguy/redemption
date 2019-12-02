var faker = require("faker");
var _ = require("lodash");

module.exports = () => {
    const data = { codes: [], campaigns: [] };

    // Creates 1 campaign (initial MVP will only have 1 campaign)
    var q = _.random(10, 1000);
    for (let i = 0; i < 1; i++) {
        data.campaigns.push({
            id: i, // autoincrement
            assets: [
                {
                    coin: 714,
                    token_id: "BUSD-BD1",
                    amount: _.random(10000, 1000000000) // needs to be BigNum?
                },
                {
                    coin: 232,
                    token_id: "BNB",
                    amount: _.random(10000, 1000000000) // need to be BigNum?
                }
            ],
            quantity: q,
            redeemed: _.random(0, q), // initialized
            invalidated: _.random(0, q * 0.2), // initialized
            provider: faker.internet.url(), // Campaign POST API URL
            address: faker.finance.bitcoinAddress(), // Campaign wallet public address
            seed: faker.random.words(12) // Campaign wallet seed phrase
        });
    }

    // Creates 1000 codes
    for (let j = 0; j < 1000; j++) {
        data.codes.push({
            id: faker.random.uuid(),
            serial: j, // autoincrements
            campaign: 1, // MVP only has 1 campaign
            code: faker.random.words(12),
            claimed: faker.random.boolean(),
            valid: faker.random.boolean(),
            claimerAddress: faker.random.boolean()
                ? ""
                : faker.finance.bitcoinAddress(),
            createdAt: faker.date.past(), // date timestamp
            updatedAt: faker.date.future(), // date timestamp
            claimedAt: faker.date.future(), // date timestamp (null for now)
            claimIPAddress: faker.internet.ip(), // claimer IP address
            claimDevice: faker.internet.userAgent(),
            newInstall: faker.random.boolean() // track whether it resulted in install of Trust Wallet
        });
    }

    return data;
};
