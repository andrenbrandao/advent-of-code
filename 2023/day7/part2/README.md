
# Algorithm

## Part 2

This change makes the Joker be able to transform into any other card.

QJJQ2

- QAJQ2
- QKJQ2
- QQJQ2

and so on..

Option 1:
For every J, we would have 13 possibilities.

Option 2:
But, we could find a more efficient way of trying to transform J into the closest values to the left or the right.

Let's code up option 1 first and then we optimize to Option 2.

    Q
  / | \
 A  K  Q
 / \
 A  K ...

We will recursively try all possibilities every time we find a new J, generating different cards.

Then, we compare the Hand with the highest type and use it as comparison.

TC: O(12^N) where N is the number of cards, which is 5 and 12 is the number of possible values of J.
