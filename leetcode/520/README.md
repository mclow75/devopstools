# Один проход с флагами

```cpp
class Solution {
public:
    bool detectCapitalUse(string word) {
        bool isFirst = false, isSmall = false, isCapital = false;
        if(word[0] >= 'A' && word[0] <= 'Z')
            isFirst = true;
        for(int i = 1; i < word.size(); i++) 
            if(word[i] >= 'A' && word[i] <= 'Z')
                isCapital = true;
            else
                isSmall = true;
        if((isFirst && !isSmall) || (!isFirst && !isCapital) || (isFirst && !isCapital))
            return true;
        return false;
    }
};
```
