#include <bits/stdc++.h>
using namespace std;

int solution()
{
    int n, r;
    cin >> n >> r;

    vector<int> family(n);
    int total_passenger = 0, unhappy = 0;
    for (auto it : family)
    {
        cin >> it;
        total_passenger += it;
        unhappy += (it % 2);
    }

    if (unhappy == 0)
        return total_passenger;

    int happy = total_passenger - unhappy;
    int remaining_seat = 2 * r - happy;

    if (remaining_seat / 2 >= unhappy)
        happy += unhappy;
    else
        happy += remaining_seat - unhappy;
    return happy;
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