func canCompleteCircuit(gas []int, cost []int) int {
    // jiral
    totalTank, totalCost := 0, 0
    start, tank := 0, 0
    for i:=0; i<len(gas); i++ {
        totalTank += gas[i]
        totalCost += cost[i]

        tank += gas[i] - cost[i]

        if tank < 0 {
            start = i + 1
            tank = 0
        }
    }

    if totalTank < totalCost {
        return -1
    }

    return start
}