# preprocess

Preprocess applies token normalisation rules to remove superficial differences between tokens.

All the functions take a channel of string and return a channel of string, this is so the functions can be one to many, or one to none (for exclusion functions).

This means that, depending on the functions used, a document containing `PIZZA` will appear when a user searches for `pizza`.
And a search for `operating` will appear when a user searches for `operational`.

## lowercase

This returns the token in lowercase.

## removeShort

Remove short ensures any tokens created are bigger than 1 character.

## language

Remove accents and diacratics.

I.e. `cliché` -> `cliche`

## removeStopWords

Remove common words which are not valuable for helping a user select a document. These will differ depending on the context but common examples are `and`, `from`, `has`, `that`, `were`.

## stemming 

In attempt to group families of words in different forms but with with the same meanings, `stemming` takes matches common variant suffixes and replaces them with one common suffix.

I.e. the following tokens would all return `operate`:
`operate`, `operating`, `operates`, `operation`, `operative`, `operatives`, `operational`
