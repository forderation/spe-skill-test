# SPE-SKILL-TEST

## How to run
- open terminal
- navigate to project folder
- type 'go mod vendor' to make sure depedency downloaded if any
- type 'go run .' or 'go build . && ./spe-skill-test'
- example are at below
```console
kharismamuzaki@kharismas-MacBook-Pro spe-skill-test % go run .
2023/09/04 17:24:11 narcissistic(153) true
2023/09/04 17:24:11 narcissistic(111) false
2023/09/04 17:24:11 [2, 4, 0, 100, 4, 11, 2602, 36]: 11
2023/09/04 17:24:11 [160, 3, 1719, 19, 11, 13, -21]: 160
2023/09/04 17:24:11 [11, 13, 15, 19, 9, 13, -21]: false
2023/09/04 17:24:11 findNeedle(["red", "blue", "yellow", "black", "grey"], "blue"): 1
2023/09/04 17:24:11 blueOcean([1,2,3,4,6,10], [1]): [2 3 4 6 10]
2023/09/04 17:24:11 blueOcean([1,5,5,5,5,3], [5]): [1 3]
```

## Unit Test
- same as how to run, but change typing to 'go test -v'
- example are at below
```console
kharismamuzaki@kharismas-MacBook-Pro spe-skill-test % go test -v                    
=== RUN   TestSpeSkillTest_narcissistic
=== RUN   TestSpeSkillTest_narcissistic/narcissistic(153)
=== RUN   TestSpeSkillTest_narcissistic/narcissistic(111)
--- PASS: TestSpeSkillTest_narcissistic (0.00s)
    --- PASS: TestSpeSkillTest_narcissistic/narcissistic(153) (0.00s)
    --- PASS: TestSpeSkillTest_narcissistic/narcissistic(111) (0.00s)
=== RUN   TestSpeSkillTest_findOutlier
=== RUN   TestSpeSkillTest_findOutlier/[2,_4,_0,_100,_4,_11,_2602,_36]
=== RUN   TestSpeSkillTest_findOutlier/[160,_3,_1719,_19,_11,_13,_-21]
=== RUN   TestSpeSkillTest_findOutlier/[11,_13,_15,_19,_9,_13,_-21]
--- PASS: TestSpeSkillTest_findOutlier (0.00s)
    --- PASS: TestSpeSkillTest_findOutlier/[2,_4,_0,_100,_4,_11,_2602,_36] (0.00s)
    --- PASS: TestSpeSkillTest_findOutlier/[160,_3,_1719,_19,_11,_13,_-21] (0.00s)
    --- PASS: TestSpeSkillTest_findOutlier/[11,_13,_15,_19,_9,_13,_-21] (0.00s)
=== RUN   TestSpeSkillTest_findNeedle
=== RUN   TestSpeSkillTest_findNeedle/findNeedle(["red",_"blue",_"yellow",_"black",_"grey"],_"blue")
=== RUN   TestSpeSkillTest_findNeedle/findNeedle(["red",_"blue",_"yellow",_"black",_"grey"],_"silver")
--- PASS: TestSpeSkillTest_findNeedle (0.00s)
    --- PASS: TestSpeSkillTest_findNeedle/findNeedle(["red",_"blue",_"yellow",_"black",_"grey"],_"blue") (0.00s)
    --- PASS: TestSpeSkillTest_findNeedle/findNeedle(["red",_"blue",_"yellow",_"black",_"grey"],_"silver") (0.00s)
=== RUN   TestSpeSkillTest_blueOcean
=== RUN   TestSpeSkillTest_blueOcean/blueOcean([1,2,3,4,6,10],_[1]):_
=== RUN   TestSpeSkillTest_blueOcean/blueOcean([1,5,5,5,5,3],_[5]):_
--- PASS: TestSpeSkillTest_blueOcean (0.00s)
    --- PASS: TestSpeSkillTest_blueOcean/blueOcean([1,2,3,4,6,10],_[1]):_ (0.00s)
    --- PASS: TestSpeSkillTest_blueOcean/blueOcean([1,5,5,5,5,3],_[5]):_ (0.00s)
PASS
ok      github.com/forderation/spe-skill-test   0.104s
```