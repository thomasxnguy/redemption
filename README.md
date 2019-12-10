# Redemption Server

> Note: README.md is a work in progress

This is an open source project for running a crypto gift card or airdrop program.

Security features will be added in over time, especially private key management.

[![Deploy](https://www.herokucdn.com/deploy/button.svg)](https://heroku.com/deploy?template=https://github.com/trustwallet/redemption/tree/pantani/initial-commit-draft)

## Features

**Creation**

-   [ ] A simple dashboard GUI allows the creation of _n_ one-use redemption codes
-   [ ] Each one-use redemption code is tied to a specified amount of cryptocurrency

**Redemption**

-   [x] Redemption POST API sends cryptocurrency to the user if the redemption code is correct and valid
-   [x] Checks that redemption codes are unclaimed
-   [x] Race condition prevention (double-spend)
-   [x] Tiebreaking algorithm if two users scan at the same time

**Dashboard and Admin**

-   [ ] Visual dashboard to track all redemption codes
-   [ ] User can invalidate unclaimed redemption codes from dashboard

**Deployment**

-   [x] Heroku Autodeploy script
-   [ ] `npm start` at root should start both backend and frontend (docker?)

## Structure

-   [ ] Document use of [react-admin](https://github.com/marmelab/react-admin)
-   [x] Document go backend API

## License

MIT licensed
