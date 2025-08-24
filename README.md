# enigma

A simple cli password generator. Inspired from [Bitwarden's Free Password Generator](https://bitwarden.com/password-generator/#password-generator). Not suitable for production. More of a thought experiment. Entirely non-vibe-coded.

## Usage

Build:

```bash
go build -o enigma
```

Command-line usage:

```bash
$ ./enigma -h
Usage of ./enigma:
  -N	include number in passphrase?
  -c	capitalize passphrase?
  -n int
    	number of characters (password) / words (passphrase) (default 14)
  -s string
    	comma-separated set of characters to pick from. values: a,A,0,! (default "a,A,0")
  -t string
    	type of secret. values: password, passphrase (default "password")
  -w string
    	word separator for passphrase (default " ")
```

To generate a password with 14 characters with a-z, A-Z, 0-9, and special characters,

```bash
$ ./enigma -n 14 -t password -s 'a,A,0,!'
4Yhyuy$4$UtsVQ
```

To generate a passphrase with 5 words,

```bash
$ ./enigma -n 5 -t passphrase -w "-" -c -N
Clinic-Carefully-Trimming-Dislodge-Surplus2
```

## Future Plans

- [ ] Add a quiet mode switch for piping into system clipboard using `xclip -selection clipboard` or `pbcopy`
- [ ] Write a Makefile
- [ ] Improve error handling
- [ ] Clean up names
- [ ] Use those fancy cli libraries like urfave/cli or spf13/cobra
