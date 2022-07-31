#include <iostream>
#include <cstdio>
#include <cstdlib>
#include <cmath>
#include <algorithm>
#include <cstring>
#include <vector>
#include <stack>
#include <queue>
#include <string>
#include <unordered_map>
#include <unordered_set>
using namespace std;

int main() {
    int m, n;
    int nums[2000000 + 100];
    cin >> m;
    for (int i = 0; i < m; i++) {
      cin >> nums[i];
    }
    cin >> n;
    for (int i = 0; i < n; i++) {
        cin >> nums[m+i];
    }
    sort(nums, nums+m+n);
    for (int i = 0; i < m + n; i++) {
        cout << nums[i] << " ";
    }
    return 0;
}