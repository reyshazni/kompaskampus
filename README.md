# Kompas Kampus

## Setup and Installation

1. Clone repo
```
git clone https://github.com/reyshazni/kompaskampus.git
```
2. Install dependencies
```
yarn install
```
3. Start the app
```
yarn dev
```

## Branching and Commit Messages

Setiap membuat branch baru harus ambil base dari `main`. Untuk penamaan commit dan branch mengikuti format berikut.
Format branch: `<type>/<title>`
Format commit: `<type>: <subject>`
Penamaan menggunakan kebab case

Untuk type mengikuti semantic berikut.
- `feat`: (new feature for the user, not a new feature for build script)
- `fix`: (bug fix for the user, not a fix to a build script)
- `docs`: (changes to the documentation)
- `style`: (formatting, missing semi colons, etc; no production code change)
- `refactor`: (refactoring production code, eg. renaming a variable)
- `test`: (adding missing tests, refactoring tests; no production code change)
- `chore`: (updating grunt tasks etc; no production code change)

## Workflow
- Kerjakan semua dari subfolder masing-masing
ex: front end membuat semuanya di subfolder apps/front-end, back end di apps/back-end

## Pull Request

Untuk melakukan pull request perlu mengikuti flow berikut.
1. Saat selesai mengerjakan task, perlu melakukan PR ke development dengan membuat PR dari branch pengerjaan ke `dev`
2. Setelah dilakukan testing di `dev` dan dapat approval, bisa melakukan PR ke `main` dan memberitahukan lead/kadiv/wakadiv.
3. Merge untuk branch `main` dilakukan oleh lead/kadiv/wakadiv.
4. Penamaan PR mengikuti format penamaan branch.
5. Pada deskripsi PR mohon menyertakan format berikut.
  - What PR About?
  - What trello link this PR refers to?
  - Is there any problems in this PR?
6. Mohon utamakan komunikasi, terutama untuk approval.

### Apps and Packages

- `front-end`: aplikasi [Next.js](https://nextjs.org/)
- `back-end`: aplikasi [Golang](https://go.dev/doc/)

### Utilities

This turborepo has some additional tools already setup for you:

- [TypeScript](https://www.typescriptlang.org/) for static type checking
- [ESLint](https://eslint.org/) for code linting
- [Prettier](https://prettier.io) for code formatting

### Build

To build all apps and packages, run the following command:

```
cd my-turborepo
yarn run build
```

### Develop

To develop all apps and packages, run the following command:

```
cd my-turborepo
yarn run dev
```

### Remote Caching

Turborepo can use a technique known as [Remote Caching](https://turbo.build/repo/docs/core-concepts/remote-caching) to share cache artifacts across machines, enabling you to share build caches with your team and CI/CD pipelines.

By default, Turborepo will cache locally. To enable Remote Caching you will need an account with Vercel. If you don't have an account you can [create one](https://vercel.com/signup), then enter the following commands:

```
cd my-turborepo
npx turbo login
```

This will authenticate the Turborepo CLI with your [Vercel account](https://vercel.com/docs/concepts/personal-accounts/overview).
Next, you can link your Turborepo to your Remote Cache by running the following command from the root of your turborepo:

```
npx turbo link
```


