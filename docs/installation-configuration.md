---
title: Rootly Setup
meta_desc: Provides an overview on how to configure credentials for the Pulumi Rootly Provider.
layout: installation
---

## Installation

1. To use this package, please [install the Pulumi CLI first](https://www.pulumi.com/docs/get-started/install/).
2. This package is only available for JavaScript and TypeScript but support for other languages will be available soon.

### Node.js (JavaScript/TypeScript)

1. To use from JavaScript or TypeScript in Node.js, install using either `npm`:

```bash
npm install @rootly/pulumi
```

or `yarn`:

```bash
yarn add @rootly/pulumi
```

## Authentication

The Pulumi Rootly Provider needs to be configured with a Rootly `API Key`.

Once you generated the `API Key` there are two ways to communicate your authorization token to Pulumi:

1. Set the environment variables `ROOTLY_API_KEY`:
    ```bash
    $ export ROOTLY_API_KEY=cu_xxx
    ```

2. Set them using `pulumi config` command, if you prefer that they be stored alongside your Pulumi stack for easy multi-user access:
    ```bash
    $ pulumi config set rootly:apiKey cu_xxx --secret
    ```

> Remember to pass `--secret` when setting `rootly:apiKey` so it is properly encrypted.
