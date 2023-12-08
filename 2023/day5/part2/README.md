# Algorithm

**Now the issue with Part2 is what I already assumed would happen if we had too many seeds.**

> There are 2 ways:
>
>- Start from a seed, find the locations and then sort them
>- Start from the lowest location and go back the path trying to reach a seed
>
>Depending on how many seeds we have, the second approach might be more efficient.
>
>Let's take S as the number of seeds and L as the number of locations.
>
>- If S << L, we should opt for Option 1
>- If S >> L, we should opt for Option 2

## How to optimize it

Now, instead of having the seed creation as O(S), since we have ranges,
each seed can end up as O(R) being R the range of seeds.

The extract seeds operation then becomes O(S*R) where R is very big.

Instead of generating all the seeds and testing for each of them,
we would need to do the reverse path: start from the lowest locations
and try to reach a seed.

## OOP

How can we change the code so that it continues to be maintainable and easy to change?

Right now we would need to reverse all the maps to make this work. But what if for some
reason we need to keep both functionalities?

Well, we could create both paths in the map and create methods to do the reverse
operation. Another option is to create maps that receive a src and dst type (Seed, Soil)
and the map will do the job to understand if that is a reverse path.

The same way we can find Locations from Seeds, we need to be able to do the other way around.

Therefore, we need to have a map of Soil to Seeds and that map is a constructed range the same
way the other maps were created.

The locations need to be calculated from the map, but instead of calculating all of them
at once, we need to start from the lowest one and start traversing the graph.

To iterate over the locations, we can use the keys of the map.

## Thinking Again

Thinking in reverse looked better, but is still pretty slow. How can we optimize it?

We have different ranges (intervals) of locations, humidity, etc. and we want to find
the overlap between all of them.

**So, instead of a graph problem, this becomes an interval problem!**

  seeds
|--------------|
       soils
     |-----------------|
     fertilizers
  |-------------|
    ......

     locations
|-----------|

    seeds to locations intersection
     |------|
