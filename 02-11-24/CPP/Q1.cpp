#include <bits/stdc++.h>
using namespace std;

//just print a binary string with exactly one 1 and rest 0
string solution()
{
    int n;
    cin >> n;
    string answer = "1";
    for (int i = 1; i < n; ++i)
    {
        answer += '0';
    }

    return answer;
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