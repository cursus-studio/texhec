# console
## Architecture
Flushes changes to the console while allowing to create temporary messages which are removed when next shows up.
Example usage is printing `fps` (frames per second)

## Lines of code
```
github.com/AlDanial/cloc
-------------------------------------------------------------------------------
Language                     files          blank        comment           code
-------------------------------------------------------------------------------
Go                               3             11              0             53
Markdown                         1              0              0              2
-------------------------------------------------------------------------------
SUM:                             4             11              0             55
-------------------------------------------------------------------------------
```
## Dependencies
`engine/modules/console`:
  - `engine/modules/console.Service`

### Third Party
- `github.com/ogiusek/ioc/v2`