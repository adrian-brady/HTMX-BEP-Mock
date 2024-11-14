# Better Evidence Project Mock using HTMX
An HTMX mock for the [Better evidence project](https://gmustaging.savepointsoftware.com/#/search) website being remade for DSSD.

## Running the web server
Run this command
```bash
go run .
```

This will start the Echo web server.

Features displayed:
- Search page url displays both "home" search page and results page
- filter button to remove some results
- "dynamic" links and pages for each result
- queries remain inside the search bar

For development, I recommend using [Air](https://github.com/air-verse/air), a live reloader for Go apps.
