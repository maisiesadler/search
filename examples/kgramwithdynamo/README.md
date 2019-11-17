# Kgram with dynamo

This example creates a dynamo dictionary and uses it as the dictinonary for the kgram index.

Running the code will output:

```
Found 2 result
- 'were' = doc-id,
- 'weve' = doc-id,
```

(It will fail the first time, since it doesn't wait for the table to be created before attempting to add to it. Once the tokens have been added the line adding them (`kgramIdx.Add("doc-id", tokens)`) can be removed.)