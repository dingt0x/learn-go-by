package main

func Sum(numbers []int) int {
    sum := 0
    for _, number := range numbers {
        sum += number
    }

    return sum
}

func SumAll(numbersToSum ...[]int) []int {
    lengthOfNumbers := len(numbersToSum)
    sums := make([]int, lengthOfNumbers)
    for i, numbers := range numbersToSum {
        sum := Sum(numbers)
        sums[i] = sum
    }
    return sums
}

func SumAllTails(numbersToSumTail ...[]int) []int {
    lengthOfNumbers := len(numbersToSumTail)
    sums := make([]int, lengthOfNumbers)
    for i, numbers := range numbersToSumTail {

        if len(numbers) < 2 {
            sums[i] = 0
        } else {
            tailNumbers := numbers[1:]
            sum := Sum(tailNumbers)
            sums[i] = sum
        }

    }
    return sums

}
