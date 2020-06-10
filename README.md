# acmetag

	go get git.omarpolo.com/acmetag

`acmetag` is a tool to programmatically interact with `acme(1)` tag.  It
provides two flag:

 * `-g` prints the tag content (mnemonic: get)
 * `-c` clears the tag content (mnemonic: clear)

any other arguments if passed will be appended to the tag.

Of course you can combine the flags:

	acmetag -g -c fmt

will first print the current tag content, than clears it and then append
`fmt`.


### Limitations

 * It cannot change the text *before* the `|` character.


### Tips

I'm using `acmetag` with [`autoacme`][autoacme] to automatically set
the tag content based on file type.

[autoacme]: https://github.com/mkhl/cmd/tree/master/acme/autoacme
