func isWinner(player1 []int, player2 []int) int {
    var winner int
    p1 := 0
    p2 := 0
    double := 0
    for i:=0; i < len(player1); i++ {
        if double > 0 {
            doubleNum := player1[i] * 2
            p1 += doubleNum
            double--
            if player1[i] == 10 {
                double = 2
            }
            continue
        }
        if player1[i] == 10 {
            double = 2
        }
        p1 += player1[i]
    } 

    double = 0

    for i:=0; i < len(player2); i++ {
        if double > 0 {
            doubleNum := player2[i] * 2
            p2 += doubleNum
            double--
            if player2[i] == 10 {
                double = 2
            }
            continue
        }
        if player2[i] == 10 {
            double = 2
        }
        p2 += player2[i]
    } 

    if p1 > p2 {
        winner = 1
    } else if p1 < p2 {
        winner = 2
    }

    return winner
}