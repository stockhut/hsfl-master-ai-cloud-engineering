# Bruno collection 🐶

This folder contains a collection of API requests that can be opened in [Bruno](https://www.usebruno.com/)

## Environments:

There are two environments included:

### User=test;

This environment contains a valid jwt token for a user called "test"

### User=test;HTMX

This environment contains a valid jwt token for a user called "test".
Is also sets the `HX-Request` header to `true`, so use this setting if you want to test hypermedia endpoints.