#include <bits/stdc++.h>
using namespace std;

int solution()
{
    string s, t;
    cin >> s >> t;

    int count_sequence = 0;
    int i = 0;

    while (i < s.size() && i < t.size())
    {
        if (s[i] == t[i])
        {
            count_sequence++;
            ++i;
        }
        else
        {
            break;
        }
    }

    if (count_sequence == 0)
        return s.size() + t.size();

    return (s.size() - count_sequence) + (t.size() - count_sequence) + (count_sequence + 1);
}

int main()
{
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int t;
    cin >> t;
    while (t--)
    {
        cout << solution() << "\n";
    }
    return 0;
}