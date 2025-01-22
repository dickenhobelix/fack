# fack

This is the file-age-check tool,
or fack for short.

Ever been tired of comparing timestamps of last-modification-date against the
current date,
just in order to check if the compiler *really* did modify the artifact?

Well these times are over now thanks to fack.

## Usage

Just run `fack <filename>` to get a human-readable representation of the
file-last-modification timestamp.
Need more information? No problem, run `fack --detailed <filename>` or
`fack -d <filename>` for short.

Want to use fack in a non-interactive way for checking if the file has been
changed in the last, say, 12h? Use `fack --non-interactive --since 12h <filename>`
or `fack -n -s 12h <filename>` for short.

## Why

Well, as stated above, I sometimes have to give a fack about when a file
was last modified. And it was a nice little sideproject to write in Go.

Also, sometimes it can be quite entertaining to type 'fack' in a screen-shared
terminal session, when your colleagues and employers are watching.

## Bugreports and Security

This is just a fun sideproject. I will do my best to respond to issues reported via
the project's github page, but please be well aware that this is not at all well-maintained software.

## License

fack is licensed under [The 3-Clause BSD License](https://opensource.org/license/bsd-3-clause).

## Last words

If you want to support me, you could contribute some features or maybe [buy me a coffee](https://ko-fi.com/dickenhobelix)

