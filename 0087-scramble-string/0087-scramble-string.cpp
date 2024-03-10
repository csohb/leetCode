class Solution {
public:
    bool isScramble(string s1, string s2) {
        // IMPORTANT: Please reset any member data you declared, as
        // the same Solution instance will be reused for each test case.
        
        //return isScramble1(s1, s2);
        
        return isScramble250(s1, s2);
    }
    
    // http://discuss.leetcode.com/questions/262/scramble-string
    // 这个是DP解法，没看懂。 250的是常规递归法，比较容易理解。
    bool isScramble1(string s1, string s2) {
        int n = s1.size();
        if (s2.size() != n) return false;

        bool f[n][n][n];

        for (int i=0; i<n; i++)
            for (int j=0; j<n; j++)
                f[i][j][0] = s1[i] == s2[j];

        for (int l=1; l<n; l++) {
            for (int i=0; i+l<n; i++) {
                for (int j=0; j+l<n; j++) {
                    f[i][j][l] = false;
                    for (int k=0; k<l; k++) {
                        if (f[i][j][k] && f[i+k+1][j+k+1][l-1-k]
                            || f[i][j+l-k][k] && f[i+k+1][j][l-1-k]) {
                            f[i][j][l] = true;
                            break;
                        }
                    }
                }
            }
        }

        return f[0][0][n-1];
    }
    
    
    bool isScramble250(string s1, string s2) 
    {
        if (s1.empty() && s2.empty()) {
            return true;
        }
        
        if (s1 == s2) {
            return true;
        }
        
        return scramble(s1, 0, s1.size() - 1, s2, 0, s2.size() - 1);
    }
    
    bool scramble(string &s1, int b1, int e1, string &s2, int b2, int e2) 
    {
        
        if (!same(s1, b1, e1, s2, b2, e2))
        {
            return false;
        }
        
        int len = e1 - b1 + 1;
        
        if (len == 0) 
        {
            return true;
        }
        
        if (len == 1) 
        {
            return (s1[b1] == s2[b2]);
        }
        
        for (int i = 0; i < len - 1; ++i) 
        {
            if ((scramble(s1, b1, b1 + i, s2, b2, b2 + i) && scramble(s1, b1 + i + 1, e1, s2, b2 + i + 1, e2)) || 
                (scramble(s1, b1, b1 + i, s2, e2 - i, e2)) && scramble(s1, b1 + i + 1, e1, s2, b2, e2 - i - 1)) 
            {
                return true;
            }
        }
        
        return false;
    }
    
    // To detect if they are the same collection.
    bool same(string &s1, int b1, int e1, string &s2, int b2, int e2) {
        vector<int> count(256, 0);
        
        for (int i = b1; i <= e1; ++i) {
            ++count[s1[i]];
        }
        
        for (int i = b2; i <= e2; ++i) {
            --count[s2[i]];
        }
        
        for (int i = 0; i < 256; ++i) {
            if (count[i] != 0) {
                return false;
            }
        }
        
        return true;
    }
};