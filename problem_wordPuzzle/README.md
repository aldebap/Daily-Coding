# Problem of word puzzle game

Given a two dimensional array of letters, find a given word written in any of the 8 directions.

EXAMPLE

Input: UBER

Input:

```
A  U  I  K  F  W  N
W  Q  B  O  L  X  P
T  L  A  E  R  E  S
Y  Z  X  E  R  L  W
```

Output: true

---

In this case, in the first row with the second column (arr[0][1]) we have a “U” which, in diagonal to bottom-right direction,
we can find entirely “UBER”, so the output should be true.
Return false if “UBER” (the input word) could not be found.

The possible directions are:

1. left to right (and reverted, right to left)
2. top to bottom (and reverted, bottom to top)
3. diagonal top-left to bottom-right like in the example (and reverted, bottom-right to top-left)
4. diagonal top-right to bottom-left (and reverted, bottom-left to top-right)
