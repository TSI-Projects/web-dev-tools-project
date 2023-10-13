# Table of Contents

- [About Project](#about-project)
- [Backend](#backend)
   - [Server](#server)
   - [API](#api)
- [Frontend](#frontend)
   - [Setup](#setup)
   - [Development](#development)
   - [Production](#production)
        - [Build the Application](#build-the-application)
        - [Local Production Preview](#local-production-preview)
- [License](#license)

# About Project

This GitHub project, created by [Vadim](https://github.com/SkinonikS) and [Andrej](https://github.com/AndrejsPon00) as part of our coursework in the "WEB Application Development Tools" subject at TSI, is designed to showcase our skills and knowledge in web application development.

This project reflects our teamwork and collaboration skills, as we've worked together to bring our ideas to life and overcome challenges along the way.

We hope that this repository serves as a valuable resource for fellow students and anyone interested in web application development. Feel free to explore our code, documentation, and share your feedback or questions. Thank you for visiting our project!

# Backend

The backend of our application is written in Golang and serves as the foundation for data retrieval. To ensure the backend functions properly, you'll need to have the Golang runtime installed on your system. If you haven't installed it yet, please refer to the official [Golang documentation](https://go.dev/doc/install) for installation instructions.

## Server

To initiate the backend server navigate to [backend](backend/scraper) page and run following command:

```bash
go run main.go
```

This command handles the automatic installation of all required dependencies and launches the server. The server will be accessible at <http://localhost:8080>.

## API

<details>
<summary>
<code>GET</code>
<code>SSE</code>
<code><b>/posts/search</b></code> - Search for posts
</summary>

#### Parameters

| Name | Type | In | Require |Description |
| :--- | :--- | :--- | :--- | :--- |
| `query` | `string` | query | + | The name of the post you want to search for. |
| `sources` | `string[]` | query | - | Specify sources to search from. |
| `pp_page` | `number` | query | - | Page number for pp.lv. |
| `ss_page` | `number` | query | - | Page number for ss.lv. |
| `facebook_page` | `number` | query | - | Page number for facebook.com. |
| `banknote_page` | `number` | query | - | Page number for banknote.lv. |

#### Events

<table>
<tr>
<td><b>Name</b></td>
<td><b>Description</b></td>
<td><b>Response</b></td>
</tr>
<tr>
<td>posts</td>
<td>This event partially sends posts from a single page.</td>
<td>

```typescript
[
  {
    title: string,
    preview_img: string,
    price: string,
    url: string,
  },
  // ... more posts
]
```

</td>
</tr>
<tr>
<td>pagination</td>
<td>This event dispatches once after a page has been scraped.</td>
<td>

```typescript
{
  source: string,
  has_next: boolean,
}
```

</td>
</tr>
<tr>
<td>close</td>
<td>
This event dispatches when the scraping process is complete. The connection will be closed.
</td>
<td>

```typescript
"Connection closed"
```

</td>
</tr>
</table>
</details>

# Frontend

The frontend of our application is powered by JavaScript and utilizes the [Nuxt.js](https://nuxt.com/) framework to run the frontend server. To ensure seamless operation, you'll need to have a JavaScript runtime installed on your system. Depending on your preference, you can choose one of the following JavaScript runtimes:
- [NodeJS](https://nodejs.org/en/download)
- [Bun](https://bun.sh/docs/installation)

## Setup

Before you can get started with the frontend, you must install the necessary dependencies. You can do this using your preferred package manager:

```bash
# npm
npm install
# pnpm
pnpm install
# yarn
yarn install
# bun
bun install
```

This command will install all the required dependencies for the frontend portion of the application. You can then either run [development](#development) server or [build](#production) the production server.

## Development

To start the development server and work on your frontend code, run the following command:

```bash
# npm
npm run dev
# pnpm
pnpm run dev
# yarn
yarn dev
# bun
bun run dev
```

This will launch a development server at http://localhost:3000, allowing you to make changes and see them in real-time.

## Production

When you're ready to prepare your application for production deployment, follow these steps:

### Build the Application

To build the application for production, execute the following command:

```bash
# npm
npm run build
# pnpm
pnpm run build
# yarn
yarn build
# bun
bun run build
```

This command will compile and optimize your code for production use.

### Local Production Preview

If you want to preview the production build locally before deployment, use the following command:

```bash
# npm
npm run preview
# pnpm
pnpm run preview
# yarn
yarn preview
# bun
bun run preview
```

This will allow you to view and test your production-ready application on your local machine, ensuring everything is working as expected before deploying it to a production server.

# License

This project is open-source and available under the [MIT License](LICENSE). You are free to use, modify, and distribute the code as per the terms of the license. Please review the full license text for more details.