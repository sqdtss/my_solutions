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
    int nums[100000 + 100];
    int t;
    cin >> t;
    while(t--) {
        int n, k;
        cin >> n >> k;
        for (int i = 0; i < n; i++) {
            cin >> nums[i];
        }
        int ans = 0, left = 0;
        for (int right = 1; right < n; right++) {
            if (nums[right] * 2 > nums[right-1]) {
                if (right - left + 1 > k) {
                    ans++;
                }
            } else {
                left = right;
            }
        }
        cout << ans << endl;
    }
    return 0;
}