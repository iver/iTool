## iTool

**Work in progress**

[2016.02.20] This tool is not complete, please come back soon.
A command-line tool for iTunes Connect Reports. 

### Dependencies

* [go version >= go1.3](https://golang.org/doc/install/source)
* autoingestion.class file from [Apple](http://apple.com/itunesnews/docs/Autoingestion.class.zip)

### Installation

You can run `make` installation:

```$ make install```

After that you must download autoingestion.class into [vendor](itool/vendor) folder.

```
$ wget http://apple.com/itunesnews/docs/Autoingestion.class.zip -O itool/vendor/Autoingestion.class.zip
$ unzip -j itool/vendor/Autoingestion.class.zip Autoingestion/?utoingestion.* -d itool/vendor/
```

### Run iTool

You can run from make:

```$ make run ```

Or if you prefere from `go run`:

```
$ cd itool
$ go run main.go
```

### TODO:

* Execute autoingestion
* Validate date parameters
* Validate range dates
* Create examples