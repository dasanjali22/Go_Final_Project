# Go_Final_Project
Parallel Sorting of Enumeration and Merge Sort using Channels and Memory Sharing

## Authors
"A1 Level" Anjali K Das (21124701)  and Sruthi Nivetha Kennedy (21140693)

# Parallel Sort Algorithms
1.Enumeration Sort

2.Merge Sort

## References
http://user.it.uu.se/~hesc0353/paral/2016-03-07b-EnumsortMPI.pdf
https://www.geeksforgeeks.org/merge-sort/


## Build
Install Golang locally (https://golang.org/doc/install) 

## Output Results
1).First Time Result:-

->The number of threads on this machine are =  1

a).Enumeration Sort:-
    [1 2 3 4 6 7 8 9 10 11 12 14 15 16 18 22 25 31 42 88]
    Enum Sort took :  154.542µs

b).Merge Sort:-
    [1 2 3 4 6 7 8 9 10 11 12 14 15 16 18 22 25 31 42 88]
    Merge Sort took :  42.625µs
            ***************************

->The number of threads on this machine are =  8

a).Enumeration Sort:-
[1 2 3 4 6 7 8 9 10 11 12 14 15 16 18 22 25 31 42 88]
After increase of cores Enum Sort took :  32.25µs

b).Merge Sort:-
[1 2 3 4 6 7 8 9 10 11 12 14 15 16 18 22 25 31 42 88]
After increase of cores Merge Sort took :  410.958µs
            ***************************

2).Second Time Result:-

->The number of threads on this machine are =  1

a).Enumeration Sort:-
    [1 2 3 4 6 7 8 9 10 11 12 14 15 16 18 22 25 31 42 88]
    Enum Sort took :  148.708µs

b).Merge Sort:-
    [1 2 3 4 6 7 8 9 10 11 12 14 15 16 18 22 25 31 42 88]
    Merge Sort took :  42.458µs
            ***************************

->The number of threads on this machine are =  8

a).Enumeration Sort:-
    [1 2 3 4 6 7 8 9 10 11 12 14 15 16 18 22 25 31 42 88]
    Enum Sort took :  9.5µs

b).Merge Sort:-
    [1 2 3 4 6 7 8 9 10 11 12 14 15 16 18 22 25 31 42 88]
    Merge Sort took :  412.667µs
            ***************************

