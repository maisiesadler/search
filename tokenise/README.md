# tokenise

Processes a block of characters into tokens to be processed.

All the functions take a channel of string and return a channel of string, this is so the functions can be one to many, or one to none (for exclusion functions).

## toWord

To word removes punctuation, and splits the text on whitespace.

For example, this sentence would become:
`For`, `example`, `this`, `sentence`, `would`, `become`

The current implementation doesn't remove angle brackets used by html `<>` so that can be removed later.

## removeHTML

Remove HTML removes any tokens that include `<>`, this is so we don't index the html tags.

## removeExcludedTerms

Given a list of excluded terms, do not return any tokens that include those values.
