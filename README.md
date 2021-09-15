## checkdomain
Takes a list of domains and checks the availability status based on the [Domainr](https://domainr.com/docs/api) API. Outputs a list of domains available for registration. Useful if you're parsing large lists of domains, such as IOC's related to malicious infrastructure.

## Recommended Usage

`$ cat domains | checkdomain`

## Install

You need to have [Go installed](https://golang.org/doc/install) and configured (i.e. with $GOPATH/bin in your $PATH):

`go get -u github.com/cybercdh/checkdomain`

## Configuration

You need to obtain a *free* API key from [Domainr](https://rapidapi.com/domainr/api/Domainr). 

The basic key will allow for 10,000 queries per month. Once obtained, export this as an environment variable:

`$ export MASHAPE_KEY=<your_key>`

or, to make this persist, add the following to your `~/.bashrc` file:

`export MASHAPE_KEY=<your_key>`

## Contributing
Pull requests are welcome. 

For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://choosealicense.com/licenses/mit/)