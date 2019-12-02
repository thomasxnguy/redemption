# Redemption Server

> Note: README.md is a work in progress

This is an open source project for running a crypto gift card or airdrop program.

## Features

**Creation**

-   [ ] A simple dashboard GUI allows the creation of _n_ one-use redemption codes
-   [ ] Each one-use redemption code is tied to a specified amount of cryptocurrency

**Redemption**

-   [ ] Redemption POST API sends cryptocurrency to the user if the redemption code is correct and valid
-   [ ] Checks that redemption codes are unclaimed
-   [ ] Race condition prevention (double-spend)
-   [ ] Tiebreaking algorithm if two users scan at the same time

**Dashboard and Admin**

-   [ ] Visual dashboard to track all redemption codes
-   [ ] User can invalidate unclaimed redemption codes from dashboard

**Deployment**

-   [ ] Heroku Autodeploy script
-   [ ] `npm start` at root should start both backend and frontend (docker?)

## Structure

-   [ ] Document use of [react-admin](https://github.com/marmelab/react-admin)
-   [ ] Document go backend API

## License

MIT licensed
