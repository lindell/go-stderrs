go-stderrs
----

This package contains commonly used error types. All errors have a function that will define that error so that you can check for it without depending on this library. Eg. `NotFound() bool`.

All errors does also indicate if it's a temporary or not, and what HTTP status code it should be converted to.

## Errors

| Name | Temporary | HTTP Status |
|------|-----------|-------------|
| BadData | No | 400|
| NotFound | No | 404|
| ServiceUnavailable | Yes | 503|

