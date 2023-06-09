This project is [FOSS](https://en.wikipedia.org/wiki/Free_and_open-source_software), available under the [BSD-3 Clause license](https://opensource.org/licenses/BSD-3-Clause).

# Krushr

Host and manage routes in your area.

## Description

Krushr is a full-stuck application for managing routes.
It consists of a JSON web API built in [Go](https://go.dev/) and a web interface
built with [Astro](https://astro.build/) and [Svelte](https://svelte.dev/).

There is no central instance, so we encourage you to run your own instance for your area.

However, we do run a ["demo" instance](https://krushr.hoenson.xyz).

## Requirements

- Go 1.20, or above
- Node 18.16.0, or above
- GNU Make 3.81, or above

## API
      
### Use without building

If you don't want to build the binary or use Docker, you can just run the application.

Make sure to configure your `.env` file and that it's in the root directory.

```sh
cp .env.example .env
go run ./cmd/api
```
      
### Building from source

// **TODO** fix grammar.
If you want to build the binary and run it, you can optionally deploy this binary to a preffered location.

Make sure to configure the `.env` file and that it's in the same directory as your binary.

```sh
cp .env.example .env
go build -o krushr ./cmd/api
./krushr
```
      
### Docker

You can also use our Docker image, just make sure you're using the corresponding backend and frontend versions.
Also make sure to configure your `.env` file.

```sh
cp .env.example .env
git pull stanofsteel/krushr
docker run -p 8080:8080 -v "/data:/data" --env-file=.env -d stanofsteel/krushr
```

## Web interface

### Building

Make sure to install dependencies and to configure the `.env` file.

```sh
cd ./ui
npm install
cp .env.example .env
npm run build # creates a dist folder containing the bundled application
```

Change the newly created `.env` file to your needs.

### Serving

Now serve the `dist` folder with your preffered file serving tool.
This is commonly done to `/var/www/html` with something like [Apache]() or [NGINX]().

---

## Configuration

The `.env` files in both the root directory and the `./ui` directory contain comments explaining the variables.

Refer to them to configure your application.

## Contributing

-

## Testing

### API

Go provides us with great built-in testing capabilities and tooling.
We've used a Makefile to simplify it even further.

- `make audit`

(Note: for this command to work, you need to have [`gofumpt`](https://github.com/mvdan/gofumpt) and [`staticcheck`](https://github.com/dominikh/go-tools) installed.
Also, you need to be in `./backend/krushr` to run this command.)

This command does multiple things. It tidies and verifies dependencies, it formats the code, it "vets" the code and finally, it runs the tests.

The vetting and linter will tell you about problems with your code and where they're located, so you can fix them.
Only when all the problems are solved, will the tests be run.

When the tests fail, you should see the expected value and the actual value.
Also, you'll see the test coverage.

- `make coverage`

This command will test your code, create a coverage profile and open said profile in your browser.

This shows you very precisely which parts of your code are covered and which aren't.

### Web interface

-

## Licensing

This project is FOSS, available under the [BSD-3 Clause license](https://opensource.org/licenses/BSD-3-Clause).

## Further reading

Below are some resources we would recommend for getting started with Go, Astro and Svelte.

We would also highly recommend you check out the documentation for each tool.

### Go

- [Go by Example](https://gobyexample.com/): a website with examples on how to many different things in Go
- [Boldly Go](https://boldlygo.tech/): a blog about Go programming
- [For the Love of Go](https://bitfieldconsulting.com/books/love): a book for beginner programmers
- [Learning Go](https://www.oreilly.com/library/view/learning-go/9781492077206/): a book for competent programmers, but new to Go
- [Go is a great programming language](https://drewdevault.com/2021/04/02/Go-is-a-great-language.html): a short article praising Go

### Astro and Svelte

- [Astro in 100 seconds](https://www.youtube.com/watch?v=dsTXcSeAZq80): a short introductory video for Astro
- [Svelte in 100 seconds](https://www.youtube.com/watch?v=rv3Yq-B8qp4): a short introductory video for Svelte

### Miscellaneous

- [LandChad.net](https://landchad.net/): a website for setting up your server
- [The transitional web](https://gomakethings.com/the-transitional-web/): a short article about the state of the web
- [Consider SQLite](https://blog.wesleyac.com/posts/consider-sqlite): an article pointing out some misconceptions about SQLite
- [The small web is beautiful](https://benhoyt.com/writings/the-small-web-is-beautiful/): an essay advocating for using "smaller" tools
