# Fleet wordgame server

This repository provides an implementation of the Fleet wordgame server.

## API

A live version of the API is hosted at [https://fleet-wordgame.herokuapp.com/](https://fleet-wordgame.herokuapp.com/).

The server stores state in-memory, and can handle multiple concurrent games. Games that are played to a win or loss state are automatically cleared from the in-memory store.
