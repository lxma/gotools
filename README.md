# gotools

Functionality that ought to be part of the Go standard library and that are standard – particularly
to functional programming languages. Examples are:
- Concat to concatenate a number of slices into one big slice
- Map to apply a function to all elements of a slice and returns a slice of the results)
- Reduce to apply a function to elements of the slice, returning f(f(f(x1, x2), x3), x4)...
- Sets (with standard functions such as intersections and unions)

Other functions just have an awkward syntax (such as the sort.Slice function where the sorting function takes
indices instead of instances). So I added a Sort function that I found easier to consume.

Such functions are possible since generics are available in Go – i.e. version 1.18. My expectation is that some of
this functionality will be provided by the Go standard library in the long run.

The library is by no means complete. (I implemented what I needed rather than taking something like the Clojure standard
library as a blueprint.) So functionality you love might be missing without rational reason. Feel free to request such
functionality.

All functionality is backed with tests and used. However, some parts could be of better efficiency. (Particularly, the
implementation of sets is slower as the implementation in e.g. Java.) However, all functionality should be at least as
good as what you can hack down on your own in a cup of hours.

One goal of this library is also to provide *simple* functionality. You should have no issues forking and maintaining
this library by yourself in case you find my support no-efficient.

Still, the current state is sufficient for myself.

# Incompatible Changes in Version 2
- `Max()` and `Min()` have been removed. Instead, functions of the standard library `slices.Min()` and `slices.Max`
  should be used for slices and the primitives `max` and `min` for single values.
- `MakeSet()` now operates on single values instead of slices. I.e. `MakeSet([]int{1,2,3})` becomes `MakeSet(1,2,3)` and
  `MakeSet(slc)` becomes `MakeSet(slc...)`.
