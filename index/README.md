# index

Add searchable tokens for a document id.

Find returns a map of matches, where the key is the documentID and the value is the number of matches.

## dictionary index

The dictionary index is a map where the key is a token and the value is a map.
The inner map has a key as the documentID and the value as the number of times the token appears for that documentID.

Each token is added as a key in the map, each time a token is seen for a document the count is incremented.

Find will look for the token in the dictionary, and return the map of documentID to counts in the result.
Find will always have exactly 0 or 1 results in the array.

## kgram index

The kgram index has an inner dictionary index where the key is the documentID and the value is the token and a count of the number of times that token appears for the documentID.
It also has a map of kgrams to documentID.

Each token added to the index will be added to the inner dictionary (if it already exists, the documentID-token count will be incremented).
The token is then parsed into 3-grams, where we create 3 letter terms for each token. For example the word `index` would have the following 3-gram terms: `$in`, `ind`, `nde`, `dex`, `ex$`, `x$i`.
For each of these terms, we create a mapping back the documentID in the dictionary.

Find will process the query into 3-grams in the same way as when indexing the tokens.
For each kgram, we then look in the kgram index to find any matching documentIDs.
We then look to find any documentIDs that appear as a match for each kgram.
If there are any matches, we look them up in the dictionary of documentIDs to tokens. If it is indeed a match, we return the documentID-count result.

The kgram index allows for wildcard searches, it will process the word into kgrams around the wildcard and search for those. For example a search for `in*x` will create kgrams `$in` and `x$i`. The term `index` would appear in the results for the kgram search and be returned as a match.