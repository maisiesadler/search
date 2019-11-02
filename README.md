# Search

Search is 3 packages that can be used to index and retreive documents by ID.

Each document is tokenised, preprocessed and added to the index by ID.
The index can then be searched to retreive documents matching by ID.

## [Tokenise](tokenise)

Takes a channel of strings and converts each string into tokens to be used in the preprocessor.

## [Preprocess](preprocess)

Preprocess takes each token and applies normalization rules to group tokens with superficial differences.

## [Index](index)

The normalized tokens can then be added to the index, with a reference to the documentID.
When the index is queried by a term it returns a list of documentIDs containing that term.