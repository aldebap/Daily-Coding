# Problem of interval overlap

Givin an interval of dates, create a function to check if it overlaps with any interval of dates in a list of them.

As an example, givin the following list of intervals:

```
[ { start: "2020-01-01", end: "2020-01-05" }, { start: "2020-01-10", end: "2020-01-20" }, { start: "2020-01-07", end: "2020-01-08" } ]
```

and the interval ```{ start: "2020-01-06", end: "2020-01-09" }```, returns ```true``` since there's a overlap of this interval with one from the list.
