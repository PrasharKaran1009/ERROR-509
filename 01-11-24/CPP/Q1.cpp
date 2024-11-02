#include <bits/stdc++.h>
using namespace std;

int solution()
{
    int n;
    cin>>n;
    vector <vector <int>> num (n,vector<int> (n,0)); 

    for(int i = 0; i < n; i++) //input
    {
        for(int j = 0; j < n; j++) 
        {
            cin >> num[i][j];
        }
    }

    int sum = 0; //read at last of code
    for(int i = 0;i<2*n-1;++i)
    {
        int mini = 0,j=0,k=0;

        if(i<n) k = (n-1)-i;
        else j = (i-n)+1;

        while (j<n && k<n) 
        {
            mini = min(mini,num[j][k]);
            ++j;
            ++k;
        }

        sum += abs(mini);
    }

    return sum;
}

int main()
{
    ios::sync_with_stdio(false);
    cin.tie(nullptr);

    int t;
    cin >> t;
    while(t--)
    {
        cout<<solution()<<endl;
    }
    return 0;
}

/*
   04
   03 14
   02 13 24
   01 12 23 34
   00 11 22 33 44
   10 21 32 43
   20 31 42
   30 41
   40

take minimum element and add there absolute and return that sum
*/