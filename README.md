# Redemption Server

> Note: This is a work in progress

This is an open source backend for a crpyto gift card or airdrop program.

## Features

**Creation**

-   A simple dashboard GUI allows the creation of _n_ one-use redemption codes
-   Each one-use redemption code is tied to a specified amount of cryptocurrency

**Redemption**

-   Redemption POST API sends cryptocurrency to the user if the redemption code is correct and valid
-   Keeps track of claimed redemption codes

*   [ ] Race condition prevention

**Deployment**

-   [ ] Heroku Autodeploy script
-   [ ] `npm start` at root should start both backend and frontend

## Structure

-   [ ] Document use of [react-admin](https://github.com/marmelab/react-admin)
-   [ ] Document go backend API

## License
