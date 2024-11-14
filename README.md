# Better Evidence Project Mock using HTMX
An HTMX mock for the [Better evidence project](https://gmustaging.savepointsoftware.com/#/search) website being remade for DSSD.

## Running the web server
To start the web server, run the following command:
```bash
go run .
```

This will start the Echo web server on port `8080`.

Features displayed:
- Search page url displays both "home" search page and results page
- filter button to remove some results
- "dynamic" links and pages for each result
- queries remain inside the search bar

Todo:
- Filter results based on search query
- Improve Go Code and split `index.html` into multiple files
- Add more unique results
- Add reading from CSV file `./data/data.csv`
- More definitely

For development, I recommend using [Air](https://github.com/air-verse/air), a live reloader for Go apps.
