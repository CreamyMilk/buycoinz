# Buycoins Backend Challange

A live instance of the api is running at 
[heroku]https://buycoinz.herokuapp.com/graphql<br/>


> What's a good reason why the pure Levenshtein Distance algorithm might be a more effective solution than t
> Damerau–Levenshtein Distance algorithm in this specific scenario ?

The *[Damerau Levenshtein](https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance)* algorithim differs from the *[Classical Levenshtein](https://en.wikipedia.org/wiki/Damerau%E2%80%93Levenshtein_distance)* algorithim by including __transpositions__
among its allowable operations in addition to the three classical single-character 
edit operations (insertions, deletions and substitutions).


Because each operation has a score as shown in the table
| Operation     | Classical Cost  | Demerau Cost|
| :--------:    |:---:            |:---: |
| Insertion     | 1               | 1    |
| Deletion      | 1               | 1    |
| Substitution  | 1               | 1    |
| Transposition | 2 (computed)    | 1    |


> For the Classical Levenshtein Algorithim   
> ``Transposition = Deletion + Insertion``


Introduction of transpositions can lead to the creation of some non-gramatical matches
that can arise in cases where two consecutive letters are swapped in
a word examples of such cases can be viewed below


##### Minor Examples

|Actual    | User Specified  | Damerau-Levenshtein Score | Classical Levenshtein Score |
|:-------: | :-------------: |      :-----------:        |       :---------:      |
|Paystack  | Paysatkc        | 2                         | 3                      |
|Buycoins  | Bucyoisn        | 2                         | 4                      |
|Etherum   | Ehterum         | 2                         | 3                      |
|Buycoins  | uBcyoins        | 2                         | 3                      |
|Bitcoin   | Bictoi          | 2                         | 3                      |
|ABCDEF    | ABDCFE          | 2                         | 3                      |

> *The values where obtained using this custom program that can be found [here](https://gist.github.com/CreamyMilk/d3cf06bc380411e1ab8fd87cd425c25b)*


<br />

With a threshold of  ` (2) ` used to qualify a valid match.
Its quite observable that using Damerau-Levenshtein to match account names is not quite appropriate in this case
as compared to Levenshtein algrothim which maintains an appropriate score to organic results .


>Damarau-Levenshtein algorithim also introduces an extra compartive check for cases where transpoition is to be applied
>but it has a negligeble effect on the efficieny of execution of the algorithim



<br />
<br />
<br />
✌️ Was Quite fun 
