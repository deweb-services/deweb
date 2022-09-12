# DWS Docs

This repo contains the documentation of Decentralized Web Services hosted on https://docs.deweb.services.
Feel free to open issues/PRs to contribute!

### Installation

```
$ npm
```

### Local Development

```
$ npm start
```

This command starts a local development server and opens up a browser window. Most changes are reflected live without having to restart the server.

### Build

```
$ npm run build
```

This command generates static content into the `build` directory and can be served using any static contents hosting service.

### Deployment

Using SSH:

```
$ USE_SSH=true npm run deploy
```

Not using SSH:

```
$ GIT_USER=<Your GitHub username> npm run deploy
```

If you are using GitHub pages for hosting, this command is a convenient way to build the website and push to the `gh-pages` branch.

### Credits

> Docs powered by [Docusaurus 2](https://docusaurus.io/), a modern static website generator.
