# Table of Contents

- [Backend](#backend)
   - [Development](#development)
   - [API](#api)
- [Frontend](#frontend)
   - [Setup](#setup)
   - [Development](#development-1)
   - [Production](#production)

# Backend

If you don't have the Golang runtime installed on your system, you'll need to install it by following the installation guide in the official [Golang documentation](https://go.dev/doc/install).

## Development
1. Go to [backend scraper](backend/scraper) directory.
2. Run following command start the development server:

```bash
go run main.go
```

It will automatically install all necessary dependencies and then run
the development server on host <http://localhost:8080>.

## API
<details>
<summary>
<code>GET</code>
<code>SSE</code>
<code><b>/posts/search</b></code>
</summary>

#### Parameters

| Name | Type | In | Require |Description |
| :--- | :--- | :--- | :--- | :--- |
| `query` | `string` | query | + | Post name to search |
| `sources` | `string[]` | query | - | Sources to search from |
| `pp_page` | `number` | query | - | pp.lv page number |
| `ss_page` | `number` | query | - | ss.lv page number |
| `facebook_page` | `number` | query | - | facebook.com page number |
| `banknote_page` | `number` | query | - | banknote.lv page number |

#### Events

<table>
<tr>
<td><b>Status</b></td>
<td><b>Description</b></td>
<td><b>Response</b></td>
</tr>
<tr>
<td>posts</td>
<td>
Partially send posts from one page.
</td>
<td>

```typescript
[
  {
    title: string,
    preview_img: string,
    price: string,
  },
  // ... more posts
]
```

</td>
</tr>
<tr>
<td>pagination</td>
<td>
Dispatches once after page was scarped.
</td>
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
Dispatches when scraping was finished.
</td>
<td>
Connection closed
</td>
</tr>
</table>
</details>

# Frontend

If you don't have the JavaScript runtime installed on your system, you'll need to install it. The [Framework](https://nuxt.com/) we use to run our frontend server may use one of the following JavaScript runtimes to work correctly:
- [NodeJS](https://nodejs.org/en/download)
- [Bun](https://bun.sh/docs/installation)

## Setup

1. Go to [frontend](frontend) directory.
2. Install the dependencies:

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

It will automatically install all necessary dependencies. You can then either run [development server](#development-server) or [bundle/build](#production) the production server.

## Development

Start the development server on `http://localhost:3000`:

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

## Production

Build the application for production:

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

Locally preview production build:

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
