<!-- ![test](https://raw.githubusercontent.com/monosodiumg/recent-stack/badges/.badges/init/test.svg) -->
![coverage](https://raw.githubusercontent.com/monosodiumg/recent-stack/badges/.badges/init/coverage.svg)
# recent-stack

A stack that discards the oldest element when full. 

# TODO
## MUST
## SHOULD
## COULD
* opt: replace the deque with an array. 
  * eliminates one layer of mutex
  * fixed-size array avoids the reallocations in deque
## WONT