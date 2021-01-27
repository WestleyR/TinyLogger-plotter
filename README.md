# csv-plotter

Simple CLI command to convert CSV data files into nice looking PNG/SVGs.

This program is intended for use with the TinyLogger (coming soon) which
will log a voltage/time, and save it on to a SD card.

## Getting started

First, you csv file should look like this:

```csv
x,y
0.01,12.31
0.02,12.30
0.03,12.28
0.04,12.29
0.05,12.25
[...]
```

Where `x` is the time, and `y` is the data, in this case, its voltage.

## Install `csv-plotter`

Since this software is still in development, the easiest way to install
`csv-plotter` is by either `go get`:

```
go get -u github.com/WestleyR/csv-plotter
```

Or by cloning the repo, and compiling/installing manually:

```
git clone https://github.com/WestleyR/csv-plotter
cd csv-plotter/
go build ./...

# then copy the binary to your preferred bin dir
cp -f csv-plotter ~/.local/bin
```

## Getting help

Please open an issue in this repository for help/suggestions.

