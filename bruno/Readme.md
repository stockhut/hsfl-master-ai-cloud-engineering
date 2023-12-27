# Bruno collection 🐶

This folder contains a collection of API requests that can be opened in [Bruno](https://www.usebruno.com/)

## Authentication

The JWT received from the login endpoint will automatically be stored in the environment using a [post response script](https://docs.usebruno.com/scripting/javascript-reference.html#javascript-api-reference) and included in other requests.
The token is stored as a secret outside the collection directory, so it is safe to share via git.

## Environments:

There are two environments included, `Default` and `HTMX`.
The `HTMX` envrionnment sets the `HX-Request` header to `true`, so use this setting if you want to test hypermedia endpoints.
